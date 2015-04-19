package ini

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

	eol = ("\n" | "\r\n");
	comment = ";" . (^[\r\n])*;
	ws = [\t ]*;
	tail = ws . comment? . eol;
	ealnum = alnum | "-" | "_";
	key = ealnum+ > tokenStart % keyEnd;
	value = (^[;\r\n])* > tokenStart % valueEnd;
	kvpair = key . ws . "=" . ws . value . tail;
	sectionName = ealnum+ > tokenStart % sectionNameEnd;
	header = "[" . sectionName . "]" . tail;
	section = tail* header (kvpair | tail)*;
	document = section*;

	main := document;
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
	)

%% write init;
%% write exec;

	return ret, nil
}
