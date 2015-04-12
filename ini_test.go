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
