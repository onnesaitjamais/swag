/*
#######
##        ____    _____ ____ _
##       (_-< |/|/ / _ `/ _ `/
##      /___/__,__/\_,_/\_, /
##                     /___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package value

import (
	"testing"
)

func TestJSON(t *testing.T) {
	const json = `{"foo": ["bar", "baz"], "a": {"b": {"c": true}}}`

	value, err := FromJSON([]byte(json))
	if err != nil {
		t.Fatal(err)
	}

	_, v, err := value.Get("foo")
	if err != nil {
		t.Error(err)
	}

	_, err = value.Bool("a", "b", "c")
	if err != nil {
		t.Error(err)
	}

	_, err = value.DBool(true, "x", "y", "z")
	if err != nil {
		t.Error(err)
	}

	_, err = v.String("1")
	if err != nil {
		t.Error(err)
	}
}

func TestYAML(t *testing.T) {
	const yaml = "a: foo\nb: 456"

	value, err := FromYAML([]byte(yaml))
	if err != nil {
		t.Fatal(err)
	}

	s, err := value.String("a")
	if err != nil {
		t.Error(err)
	} else if s != "foo" {
		t.Errorf("String(): got=%s want=%s", s, "foo")
	}

	i, err := value.Int("b")
	if err != nil {
		t.Error(err)
	} else if i != 456 {
		t.Errorf("Int(): got=%d want=%d", i, 456)
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
