package decode

import (
	"testing"
)

var Base16Test = []struct {
	in       string
	eOut     string
	eIsPrint bool
}{
	{
		"666F6F626172",
		"foobar",
		true,
	},
	{
		"666F6F626172.!666F6F626172",
		"foobar" + genInvalid(2) + "foobar",
		false,
	},
}

func TestB16Decode(t *testing.T) {
	//invalid = '§'
	for _, tt := range Base16Test {
		d := NewBase16CodecC(tt.in)
		out, ip := d.Decode()
		if out != tt.eOut {
			t.Errorf("Expected decoded value: '%s' but got '%s'", tt.eOut, out)
		}
		if ip != tt.eIsPrint {
			t.Errorf("Expected printable: %v but got %v", tt.eIsPrint, ip)
		}
	}
}
