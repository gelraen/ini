package ini

import (
	"reflect"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	cases := map[string]Document{
		"[section]\nkey=value\n\n": Document{"section": map[string]string{"key": "value"}},
	}

	for in, expected := range cases {
		actual, err := Parse(strings.NewReader(in))
		if err != nil || !reflect.DeepEqual(expected, actual) {
			t.Errorf("Parse(%q) = (%+v, %s), want (%+v, <nil>)", in, actual, err, expected)
		}
	}
}

func TestInvalidInputs(t *testing.T) {
	cases := []string{
		"[section\nkey=value\n\n",
		"no_section=value\n",
	}

	for _, in := range cases {
		v, err := Parse(strings.NewReader(in))
		if err == nil {
			t.Errorf("Parse(%q) = (%+v, nil), want error", in, v)
		} else {
			t.Logf("Parse(%q): %s", in, err)
		}
	}
}
