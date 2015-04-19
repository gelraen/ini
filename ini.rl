package ini

import (
	"fmt"
)

%%{
	machine ini;

	action tokenStart { start = fpc }
	
	action sectionNameEnd {
		section = string(data[start:fpc])
		// TODO(imax): detect duplicate sections.
		ret[section] = map[string]string{}
	}

	action keyEnd {
		key = string(data[start:fpc])
	}

	action valueEnd {
		ret[section][key] = string(data[start:fpc])
	}

	action error {
		return nil, fmt.Errorf("line %d char %d: parse failed in state %d", line, fpc - lineStart + 1, fcurs)
	}

	action incLine {
		line++
		lineStart = fpc
	}

	eol = ("\n" | "\r\n") % incLine;
	comment = ";" . (^[\r\n])*;
	ws = [\t ]*;
	tail = ws . comment? . eol;
	ealnum = alnum | "-" | "_";
	key = ealnum+ > tokenStart % keyEnd;
	value = (^[;\r\n])* > tokenStart % valueEnd;
	kvpair = key . ws . "=" . ws . value;
	sectionName = ealnum+ > tokenStart % sectionNameEnd;
	header = "[" . sectionName . "]";
	section = tail* . header . tail . (kvpair? . tail)*;
	document = section*;

	main := document $!error;
}%%

%% write data;

func ragel_machine(data []byte) (Document, error) {
	var cs,p,pe,eof int
	eof = len(data)
	pe = eof
	var (
		start int = -1
		ret = Document{}
		section = ""
		key = ""
		line = 1
		lineStart = 0
	)

%% write init;
%% write exec;

	return ret, nil
}
