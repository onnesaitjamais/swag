/*
#######
##        ____    _____ ____ _
##       (_-< |/|/ / _ `/ _ `/
##      /___/__,__/\_,_/\_, /
##                     /___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package logfmt

import (
	"bytes"
	"testing"
)

func TestEncoder(t *testing.T) {
	buf := bytes.Buffer{}
	enc := NewEncoder(&buf)

	data := []struct {
		kv   []interface{}
		want string
	}{
		// 00
		{
			[]interface{}{},
			"",
		},
		// 01
		{
			[]interface{}{nil, nil},
			"@nil=@nil",
		},
		// 02
		{
			[]interface{}{"age", 53},
			"age=53",
		},
		// 03
		{
			[]interface{}{"a\tb\nc", "def"},
			"abc=\"def\"",
		},
		// 04
		{
			[]interface{}{[]byte("lsm"), "ceci est un message"},
			"[]byte{0x6c,0x73,0x6d}=\"ceci est un message\"",
		},
		// 05
		{
			[]interface{}{"", 789.456},
			"@key=789.456",
		},
		// 06
		{
			[]interface{}{"jour", 24, "mois", "décembre", "année", 2019},
			"jour=24 mois=\"décembre\" année=2019",
		},
		// 07
		{
			[]interface{}{"la valeur est manquante"},
			"lavaleurestmanquante=@nil",
		},
		// 08
		{
			[]interface{}{"message", "Joyeuses\tfêtes\n"},
			"message=\"Joyeuses\\tfêtes\\n\"",
		},
	}

	for i, d := range data {
		if err := enc.Encode(d.kv...); err == nil {
			if got := buf.String(); got != d.want {
				t.Errorf("Encode(): test %02d => got=[%v] want=[%v]", i, got, d.want)
			}
		} else {
			t.Error(err)
		}

		buf.Reset()
		enc.Reset()
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
