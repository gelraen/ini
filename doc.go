// Package ini implements a parser for INI files written as an exercise.
package ini

//go:generate ragel -Z ini.rl

import (
	"io"
	"io/ioutil"
)

// Document represents the content of the INI file.
// First key is section name, second - property name.
type Document map[string]map[string]string

// Parse reads data form r and returns a parsed document.
func Parse(r io.Reader) (Document, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	// Hack to avoid parser failure on missing newline at the EOF.
	if data[len(data)-1] != '\n' {
		data = append(data, '\n')
	}
	return ragel_machine(data)
}

// String returns a Document in INI format. No guarantees about ordering of fields and sections.
func (d Document) String() string {
	r := ""
	for section, values := range d {
		r += "[" + section + "]\n"
		for key, value := range values {
			r += key + "=" + value + "\n"
		}
		r += "\n"
	}
	return r
}
