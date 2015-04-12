// Package ini implements a parser for INI files written as an excercise.
package ini

import (
	"bufio"
	"fmt"
	"io"
)

// Document represents the content of the INI file.
// First key is section name, second - property name.
type Document map[string]map[string]string

// Parse reads data form r and returns a parsed document.
func Parse(r io.Reader) (Document, error) {
	p := &parser{
		scanner: bufio.NewScanner(r),
		line:    1,
		pos:     1,
	}
	p.scanner.Split(bufio.ScanRunes)
	return p.run()
}

type parser struct {
	scanner   *bufio.Scanner
	buf       string
	line, pos int
}

func (p *parser) run() (Document, error) {
	return p.parseDocument()
}

func (p *parser) peek() (string, error) {
	if p.buf == "" {
		if !p.scanner.Scan() {
			if err := p.scanner.Err(); err != nil {
				return "", err
			}
			return "", io.EOF
		}
		p.buf = p.scanner.Text()
	}
	return p.buf, nil
}

func (p *parser) match(s string) error {
	v, err := p.peek()
	if err != nil {
		return err
	}
	if v != s {
		return p.Errorf("unexpected character %q (want %q)", v, s)
	}
	p.buf = ""
	p.pos++
	if s == "\n" {
		p.line++
		p.pos = 1
	}
	return nil
}

func (p *parser) Errorf(f string, args ...interface{}) error {
	return fmt.Errorf("line %d char %d: "+f, append([]interface{}{p.line, p.pos}, args...)...)
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\x09'
}

func isAlphaNumeric(ch byte) bool {
	return (ch >= 0x41 && ch <= 0x5A) || (ch >= 0x61 && ch <= 0x7A) || ch >= 0x80 || (ch >= '0' && ch <= '9') || ch == '-' || ch == '_'
}

func (p *parser) parseDocument() (Document, error) {
	r := Document{}
	for {
		name, values, err := p.parseSection()
		if err == io.EOF {
			r[name] = values
			return r, nil
		}
		if err != nil {
			return r, err
		}
		r[name] = values
	}
}

func (p *parser) parseSection() (string, map[string]string, error) {
	var name string
	values := map[string]string{}
	var err error

	for {
		s, err := p.peek()
		if err != nil {
			return name, values, err
		}
		if !isWhitespace(s[0]) && s != ";" && s != "\r" && s != "\n" {
			break
		}
		err = p.parseTail()
		if err != nil {
			return name, values, err
		}
	}

	name, err = p.parseHeader()
	if err != nil {
		return name, values, err
	}

	for {
		s, err := p.peek()
		if err != nil {
			return name, values, err
		}
		switch {
		case isAlphaNumeric(s[0]):
			k, v, err := p.parseKvpair()
			if err != nil {
				return name, values, err
			}
			values[k] = v
		case isWhitespace(s[0]) || s == ";" || s == "\r" || s == "\n":
			err := p.parseTail()
			if err != nil {
				return name, values, err
			}
		default:
			return name, values, nil
		}
	}
}

func (p *parser) parseHeader() (string, error) {
	err := p.match("[")
	if err != nil {
		return "", err
	}

	var name string
	for {
		s, err := p.peek()
		if err != nil {
			return "", err
		}
		if !isAlphaNumeric(s[0]) {
			break
		}
		if err := p.match(s); err != nil {
			return "", err
		}
		name += s
	}

	if err := p.match("]"); err != nil {
		return "", err
	}

	return name, p.parseTail()
}

func (p *parser) parseKvpair() (string, string, error) {
	key, err := p.parseKey()
	if err != nil {
		return "", "", err
	}

	if err = p.parseWhitespace(); err != nil {
		return "", "", err
	}
	if err = p.match("="); err != nil {
		return "", "", err
	}
	if err = p.parseWhitespace(); err != nil {
		return "", "", err
	}

	val, err := p.parseValue()
	if err != nil {
		return "", "", err
	}

	return key, val, p.parseTail()
}

func (p *parser) parseKey() (string, error) {
	var name string
	for {
		s, err := p.peek()
		if err != nil {
			return "", err
		}
		if !isAlphaNumeric(s[0]) {
			return name, nil
		}
		if err := p.match(s); err != nil {
			return "", err
		}
		name += s
	}
}

func (p *parser) parseValue() (string, error) {
	var v string
	for {
		s, err := p.peek()
		if err == io.EOF {
			return v, nil
		}
		if err != nil {
			return "", err
		}
		if s[0] < 0x21 && !isWhitespace(s[0]) {
			return v, nil
		}
		if err := p.match(s); err != nil {
			return "", err
		}
		v += s
	}
}

func (p *parser) parseTail() error {
	err := p.parseWhitespace()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	s, err := p.peek()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if s == ";" {
		err = p.parseComment()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
	}
	return p.parseEol()
}

func (p *parser) parseComment() error {
	for {
		s, err := p.peek()
		if err != nil {
			return err
		}
		if s == "\r" || s == "\n" {
			return nil
		}
		if err := p.match(s); err != nil {
			return err
		}
	}
}

func (p *parser) parseWhitespace() error {
	for {
		s, err := p.peek()
		if err != nil {
			return err
		}
		if !isWhitespace(s[0]) {
			return nil
		}
		if err := p.match(s); err != nil {
			return err
		}
	}
}

func (p *parser) parseEol() error {
	s, err := p.peek()
	if err != nil {
		return err
	}
	switch s {
	case "\r":
		if err := p.match(s); err != nil {
			return err
		}
		if err := p.match("\n"); err != nil {
			return err
		}
		return nil
	case "\n":
		if err := p.match(s); err != nil {
			return err
		}
		return nil
	}
	return p.Errorf("unexpected character %q (want \\r or \\n)", s)
}
