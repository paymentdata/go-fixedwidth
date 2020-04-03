package fixedwidth

import (
	"reflect"
	"testing"
)

func TestParseTag(t *testing.T) {
	for _, tt := range []struct {
		name     string
		tag      string
		startPos int
		endPos   int
		format   Format
		ok       bool
	}{
		{"Valid Tag", "0,10", 0, 10, DefaultFormat, true},
		{"Valid Tag Single position", "5,5", 5, 5, DefaultFormat, true},
		{"Valid Tag w/ Alignment", "0,10,right", 0, 10, Format{Right, DefaultPadChar}, true},
		{"Valid Tag w/ Padding Character", "0,10,default,0", 0, 10, Format{DefaultAlignment, '0'}, true},
		{"Tag Empty", "", 0, 0, DefaultFormat, false},
		{"Tag Too short", "0", 0, 0, DefaultFormat, false},
		{"Tag Too Long", "2,10,default,_,foo", 0, 0, DefaultFormat, false},
		{"StartPos Not Integer", "hello,3", 0, 0, DefaultFormat, false},
		{"EndPos Not Integer", "3,hello", 0, 0, DefaultFormat, false},
		{"Tag Contains a Space", "4, 11", 0, 0, DefaultFormat, false},
		{"Tag Interval Invalid", "14,5", 0, 0, DefaultFormat, false},
		{"Tag Both Positions Zero", "0,0", 0, 0, DefaultFormat, false},
		{"Space Padding Character", "0,0,default, ", 0, 0, DefaultFormat, false},
		{"Space Padding Character (_)", "0,0,default,_", 0, 0, DefaultFormat, false},
		{"Underscore Padding Character (__)", "0,0,default,__", 0, 0, DefaultFormat, false},
		{"Multi-byte Padding Character", "0,0,default,00", 0, 0, DefaultFormat, false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			startPos, endPos, format, ok := parseTag(tt.tag)
			if tt.ok != ok {
				t.Errorf("parseTagWithFormat() ok want %v, have %v", tt.ok, ok)
			}

			// only check startPos and endPos if valid tags are expected
			if tt.ok {
				if tt.startPos != startPos {
					t.Errorf("parseTagWithFormat() startPos want %v, have %v", tt.startPos, startPos)
				}
				if tt.endPos != endPos {
					t.Errorf("parseTagWithFormat() endPos want %v, have %v", tt.endPos, endPos)
				}
				if !reflect.DeepEqual(tt.format, format) {
					t.Errorf("parseTagWithFormat() format want %+v, have %+v", tt.format, format)
				}
			}
		})
	}
}
