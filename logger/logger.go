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
	"fmt"
	"os"
	"time"

	"github.com/arnumina/swag/component"
)

type logger struct {
	lvl    Level
	runner string
	enc    Encoder
	out    Output
	ctx    []interface{}
}

// New AFAIRE
func (l *logger) New(lvl string, ctx ...interface{}) component.Logger {
	logger := &logger{
		lvl:    GetLevelFromString(lvl),
		runner: l.runner,
		enc:    l.enc,
		out:    l.out,
	}

	logger.ctx = append(logger.ctx, l.ctx...)
	if len(ctx) != 0 {
		logger.ctx = append(logger.ctx, ctx...)
	}

	return logger
}

func (l *logger) log(lvl Level, msg string, ctx ...interface{}) {
	if lvl < l.lvl {
		return
	}

	dt := time.Now()

	nc := append([]interface{}{}, l.ctx...)
	if len(ctx) != 0 {
		nc = append(nc, ctx...)
	}

	buf, err := l.enc.Encode(l.runner, lvl, msg, nc, dt, l.out)
	if err != nil {
		buf = []byte(fmt.Sprintf("ERROR [simple|log()] reason=%s", err)) ///////////////////////////////////////////////
	}

	if err = l.out.Log(lvl, buf); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR [simple|log()] reason=%s", err) //////////////////////////////////////////////////
	}
}

// Trace AFAIRE
func (l *logger) Trace(msg string, ctx ...interface{}) {
	l.log(TLevel, msg, ctx...)
}

// Debug AFAIRE
func (l *logger) Debug(msg string, ctx ...interface{}) {
	l.log(DLevel, msg, ctx...)
}

// Info AFAIRE
func (l *logger) Info(msg string, ctx ...interface{}) {
	l.log(ILevel, msg, ctx...)
}

// Notice AFAIRE
func (l *logger) Notice(msg string, ctx ...interface{}) {
	l.log(NLevel, msg, ctx...)
}

// Warning AFAIRE
func (l *logger) Warning(msg string, ctx ...interface{}) {
	l.log(WLevel, msg, ctx...)
}

// Error AFAIRE
func (l *logger) Error(msg string, ctx ...interface{}) {
	l.log(ELevel, msg, ctx...)
}

// Critical AFAIRE
func (l *logger) Critical(msg string, ctx ...interface{}) {
	l.log(CLevel, msg, ctx...)
}

// Close AFAIRE
func (l *logger) Close() {
	if err := l.out.Close(); err != nil {
		fmt.Fprintf( ///////////////////////////////////////////////////////////////////////////////////////////////////
			os.Stderr,
			"Error when closing the logger >>> %s\n",
			err,
		)
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
