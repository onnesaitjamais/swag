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
	"fmt"
	"io"
	"strings"
	"unicode/utf8"
)

const (
	_badKey   = "@key"
	_nilValue = "@nil"
)

// Encoder représente l'encodeur qui génère le format "logfmt".
type Encoder struct {
	writer    io.Writer
	buf       bytes.Buffer
	needSpace bool
}

// NewEncoder permet de créer une nouvelle instance d'un encodeur avec 'w' pour le flux de sortie.
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{writer: w}
}

// Reset permet de réinitialiser l'encodeur pour une nouvelle liste de couples clé/valeur.
func (e *Encoder) Reset() {
	e.needSpace = false
}

// Les caractères qui ne sont pas autorisés dans une clé sont supprimés.
func cleanKey(r rune) rune {
	if r <= ' ' || r == '=' || r == '"' || r == utf8.RuneError {
		return -1
	}

	return r
}

func (e *Encoder) writeKey(key interface{}) {
	switch v := key.(type) {
	case string:
		s := strings.Map(cleanKey, v)
		if s == "" {
			s = _badKey
		}

		e.buf.WriteString(s)
	// La clé est transformée en une chaîne de caractères si elle n'en est pas une.
	case nil:
		e.buf.WriteString(_nilValue)
	default:
		s := strings.Map(cleanKey, fmt.Sprintf("%#v", v))
		if s == "" {
			s = _badKey
		}

		e.buf.WriteString(s)
	}
}

func (e *Encoder) writeValue(value interface{}) {
	switch v := value.(type) {
	case int, int8, int16, int32, int64:
		e.buf.WriteString(fmt.Sprint(v))
	case uint, uint8, uint16, uint32, uint64:
		e.buf.WriteString(fmt.Sprint(v))
	case float32, float64:
		e.buf.WriteString(fmt.Sprint(v))
	case nil:
		e.buf.WriteString(_nilValue)
	default:
		e.buf.WriteString(fmt.Sprintf("%#v", v))
	}
}

func (e *Encoder) encodeKeyValue(key, value interface{}) error {
	// Si ce n'est pas le premier couple, il faut ajouter un espace.
	if e.needSpace {
		e.buf.WriteString(" ")
	}

	e.writeKey(key)
	e.buf.WriteString("=")
	e.writeValue(value)

	// Le résultat de l'encodage du couple est écrit dans le flux de sortie.
	_, err := e.writer.Write(e.buf.Bytes())

	e.buf.Reset()
	e.needSpace = true

	return err
}

// Encode effectue l'encodage de couples clé/valeur 'kv' au format "logfmt".
// Normalement, la clé est une chaîne de caractères identifiant la valeur.
func (e *Encoder) Encode(kv ...interface{}) error {
	if len(kv) == 0 {
		return nil
	}

	// Si la liste est impair, on ajoute arbitrairement la valeur 'nil'.
	if len(kv)%2 == 1 {
		kv = append(kv, nil)
	}

	for i := 0; i < len(kv); i += 2 {
		if err := e.encodeKeyValue(kv[i], kv[i+1]); err != nil {
			return err
		}
	}

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
