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
	"log"

	"github.com/arnumina/swag/component"
)

// Write AFAIRE
func (l *logger) Write(p []byte) (int, error) {
	l.Notice(string(p))
	return len(p), nil
}

// NewStdLogger AFAIRE
func (l *logger) NewStdLogger(prefix string, flag int) *log.Logger {
	return log.New(l, prefix, flag)
}

type stdLogAdapter struct {
	lvl    Level
	logger *logger
	ctx    []interface{}
}

// NewAdapter AFAIRE
func (l *logger) NewStdLogAdapter(lvl string, ctx ...interface{}) component.StdLogAdapter {
	return &stdLogAdapter{
		lvl:    GetLevelFromString(lvl),
		logger: l,
		ctx:    append([]interface{}{}, ctx...),
	}
}

// Write AFAIRE
func (a *stdLogAdapter) Write(p []byte) (int, error) {
	a.logger.log(a.lvl, string(p), a.ctx...)
	return len(p), nil
}

// NewStdLogger AFAIRE
func (a *stdLogAdapter) NewStdLogger(prefix string, flag int) *log.Logger {
	return log.New(a, prefix, flag)
}

/*
######################################################################################################## @(°_°)@ #######
*/
