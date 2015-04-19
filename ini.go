
//line ini.rl:1
package ini

import (
	"fmt"
)


//line ini.rl:49



//line ini.go:15
var _ini_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 1, 3, 
	1, 4, 1, 5, 2, 0, 3, 2, 
	5, 0, 
}

var _ini_key_offsets []byte = []byte{
	0, 0, 5, 11, 12, 14, 22, 31, 
	36, 37, 48, 51, 56, 61, 63, 69, 
}

var _ini_trans_keys []byte = []byte{
	9, 10, 13, 32, 59, 9, 10, 13, 
	32, 59, 91, 10, 10, 13, 45, 95, 
	48, 57, 65, 90, 97, 122, 45, 93, 
	95, 48, 57, 65, 90, 97, 122, 9, 
	10, 13, 32, 59, 10, 9, 32, 45, 
	61, 95, 48, 57, 65, 90, 97, 122, 
	9, 32, 61, 9, 10, 13, 32, 59, 
	9, 10, 13, 32, 59, 10, 13, 9, 
	10, 13, 32, 59, 91, 9, 10, 13, 
	32, 45, 59, 91, 95, 48, 57, 65, 
	90, 97, 122, 
}

var _ini_single_lengths []byte = []byte{
	0, 5, 6, 1, 2, 2, 3, 5, 
	1, 5, 3, 5, 5, 2, 6, 8, 
}

var _ini_range_lengths []byte = []byte{
	0, 0, 0, 0, 0, 3, 3, 0, 
	0, 3, 0, 0, 0, 0, 0, 3, 
}

var _ini_index_offsets []byte = []byte{
	0, 0, 6, 13, 15, 18, 24, 31, 
	37, 39, 48, 52, 58, 64, 67, 74, 
}

var _ini_indicies []byte = []byte{
	1, 2, 3, 1, 4, 0, 5, 6, 
	7, 5, 8, 9, 0, 2, 0, 2, 
	3, 4, 10, 10, 10, 10, 10, 0, 
	11, 12, 11, 11, 11, 11, 0, 13, 
	14, 15, 13, 16, 0, 14, 0, 17, 
	17, 18, 19, 18, 18, 18, 18, 0, 
	20, 20, 21, 0, 23, 24, 25, 23, 
	26, 22, 28, 29, 30, 28, 31, 27, 
	14, 15, 16, 1, 2, 3, 1, 4, 
	32, 0, 33, 34, 35, 33, 36, 37, 
	9, 36, 36, 36, 36, 0, 
}

var _ini_trans_targs []byte = []byte{
	0, 1, 2, 3, 4, 1, 2, 3, 
	4, 5, 6, 6, 7, 7, 15, 8, 
	13, 10, 9, 11, 10, 11, 12, 11, 
	15, 8, 13, 12, 12, 15, 8, 13, 
	5, 7, 15, 8, 9, 13, 
}

var _ini_trans_actions []byte = []byte{
	9, 0, 0, 0, 0, 11, 11, 11, 
	11, 11, 1, 0, 3, 0, 0, 0, 
	0, 5, 0, 5, 0, 0, 1, 13, 
	13, 13, 13, 0, 7, 7, 7, 7, 
	0, 11, 11, 11, 16, 11, 
}

var _ini_eof_actions []byte = []byte{
	0, 9, 9, 9, 9, 9, 9, 9, 
	9, 9, 9, 9, 9, 9, 0, 11, 
}

const ini_start int = 14
const ini_first_final int = 14
const ini_error int = 0

const ini_en_main int = 14


//line ini.rl:52

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


//line ini.go:114
	{
	cs = ini_start
	}

//line ini.rl:67

//line ini.go:121
	{
	var _klen int
	var _ps int
	var _trans int
	var _acts int
	var _nacts uint
	var _keys int
	if p == pe {
		goto _test_eof
	}
	if cs == 0 {
		goto _out
	}
_resume:
	_keys = int(_ini_key_offsets[cs])
	_trans = int(_ini_index_offsets[cs])

	_klen = int(_ini_single_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + _klen - 1)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + ((_upper - _lower) >> 1)
			switch {
			case data[p] < _ini_trans_keys[_mid]:
				_upper = _mid - 1
			case data[p] > _ini_trans_keys[_mid]:
				_lower = _mid + 1
			default:
				_trans += int(_mid - int(_keys))
				goto _match
			}
		}
		_keys += _klen
		_trans += _klen
	}

	_klen = int(_ini_range_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + (_klen << 1) - 2)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + (((_upper - _lower) >> 1) & ^1)
			switch {
			case data[p] < _ini_trans_keys[_mid]:
				_upper = _mid - 2
			case data[p] > _ini_trans_keys[_mid + 1]:
				_lower = _mid + 2
			default:
				_trans += int((_mid - int(_keys)) >> 1)
				goto _match
			}
		}
		_trans += _klen
	}

_match:
	_trans = int(_ini_indicies[_trans])
	_ps = cs
	cs = int(_ini_trans_targs[_trans])

	if _ini_trans_actions[_trans] == 0 {
		goto _again
	}

	_acts = int(_ini_trans_actions[_trans])
	_nacts = uint(_ini_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _ini_actions[_acts-1] {
		case 0:
//line ini.rl:10
 start = p 
		case 1:
//line ini.rl:12

		section = string(data[start:p])
		// TODO(imax): detect duplicate sections.
		ret[section] = map[string]string{}
	
		case 2:
//line ini.rl:18

		key = string(data[start:p])
	
		case 3:
//line ini.rl:22

		ret[section][key] = string(data[start:p])
	
		case 4:
//line ini.rl:26

		return nil, fmt.Errorf("line %d char %d: parse failed in state %d", line, p - lineStart + 1, (_ps))
	
		case 5:
//line ini.rl:30

		line++
		lineStart = p
	
//line ini.go:233
		}
	}

_again:
	if cs == 0 {
		goto _out
	}
	p++
	if p != pe {
		goto _resume
	}
	_test_eof: {}
	if p == eof {
		__acts := _ini_eof_actions[cs]
		__nacts := uint(_ini_actions[__acts]); __acts++
		for ; __nacts > 0; __nacts-- {
			__acts++
			switch _ini_actions[__acts-1] {
			case 4:
//line ini.rl:26

		return nil, fmt.Errorf("line %d char %d: parse failed in state %d", line, p - lineStart + 1, (_ps))
	
			case 5:
//line ini.rl:30

		line++
		lineStart = p
	
//line ini.go:263
			}
		}
	}

	_out: {}
	}

//line ini.rl:68

	return ret, nil
}
