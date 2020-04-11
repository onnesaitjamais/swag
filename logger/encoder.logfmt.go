/*
#######
##        ____    _____ ____ _
##       (_-< |/|/ / _ `/ _ `/
##      /___/__,__/\_,_/\_, /
##                     /___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package logger

import (
	"bytes"
	"sync"
	"time"

	"github.com/arnumina/swag/util/logfmt"
)

// LogFmtEncoder AFAIRE
type LogFmtEncoder struct {
	enc   *logfmt.Encoder
	buf   bytes.Buffer
	mutex sync.Mutex
}

// NewLogFmtEncoder permet de créer une nouvelle instance d'un encodeur de type "LogFmtEncoder".
func NewLogFmtEncoder() *LogFmtEncoder {
	encoder := &LogFmtEncoder{}
	encoder.enc = logfmt.NewEncoder(&encoder.buf)

	return encoder
}

// Encode encode le message de log.
func (e *LogFmtEncoder) Encode(
	runner string, lvl Level, msg string, ctx []interface{}, dt time.Time, out Output,
) ([]byte, error) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	e.buf.Reset()
	e.enc.Reset()

	// La sortie accepte-t-elle la date et l'heure ?
	if out.LogDateTime() {
		e.buf.WriteString(dt.Format("2006-01-02T15:04:05.000 "))
	}

	// La sortie accepte-t-elle le niveau de log ?
	if out.LogLevel() {
		switch lvl {
		case TLevel:
			e.buf.WriteString("{TRA} ")
		case DLevel:
			e.buf.WriteString("{DEB} ")
		case ILevel:
			e.buf.WriteString("{INF} ")
		case NLevel:
			e.buf.WriteString("{NOT} ")
		case WLevel:
			e.buf.WriteString("{WAR} ")
		case ELevel:
			e.buf.WriteString("{ERR} ")
		case CLevel:
			e.buf.WriteString("{CRI} ")
		default:
			e.buf.WriteString("{???} ")
		}
	}

	e.buf.WriteString(runner)
	e.buf.WriteString(" ")
	e.buf.WriteString(msg)

	if len(ctx) != 0 {
		e.buf.WriteString("> ")
		// Encodage au format "logfmt" des couples clé/valeur du contexte.
		err := e.enc.Encode(ctx...)
		if err != nil {
			return nil, err
		}
	}

	// La sortie accepte-t-elle un "\n" final ?
	if out.AddNewLine() {
		e.buf.WriteString("\n")
	}

	return e.buf.Bytes(), nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
