
//line ini.rl:1
package ini


//line ini.rl:36



//line ini.go:11
var _ini_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 1, 3, 
	2, 0, 3, 
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
	0, 2, 3, 0, 4, 1, 0, 2, 
	3, 0, 4, 5, 1, 2, 1, 2, 
	3, 4, 6, 6, 6, 6, 6, 1, 
	7, 8, 7, 7, 7, 7, 1, 9, 
	10, 11, 9, 12, 1, 10, 1, 13, 
	13, 14, 15, 14, 14, 14, 14, 1, 
	16, 16, 17, 1, 19, 20, 21, 19, 
	22, 18, 24, 25, 26, 24, 27, 23, 
	10, 11, 12, 0, 2, 3, 0, 4, 
	5, 1, 9, 10, 11, 9, 28, 12, 
	5, 28, 28, 28, 28, 1, 
}

var _ini_trans_targs []byte = []byte{
	1, 0, 2, 3, 4, 5, 6, 6, 
	7, 7, 15, 8, 13, 10, 9, 11, 
	10, 11, 12, 11, 15, 8, 13, 12, 
	12, 15, 8, 13, 9, 
}

var _ini_trans_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 1, 0, 
	3, 0, 0, 0, 0, 5, 0, 5, 
	0, 0, 1, 9, 9, 9, 9, 0, 
	7, 7, 7, 7, 1, 
}

const ini_start int = 14
const ini_first_final int = 14
const ini_error int = 0

const ini_en_main int = 14


//line ini.rl:39

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


//line ini.go:100
	{
	cs = ini_start
	}

//line ini.rl:52

//line ini.go:107
	{
	var _klen int
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
//line ini.rl:6
 start = p 
		case 1:
//line ini.rl:8

		section = string(data[start:p])
		// TODO(imax): detect duplicate sections.
		ret[section] = map[string]string{}
	
		case 2:
//line ini.rl:14

		key = string(data[start:p])
	
		case 3:
//line ini.rl:18

		ret[section][key] = string(data[start:p])
	
//line ini.go:206
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
	_out: {}
	}

//line ini.rl:53

	return ret, nil
}
