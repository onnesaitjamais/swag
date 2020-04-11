/*
#######
##        ____    _____ ____ _
##       (_-< |/|/ / _ `/ _ `/
##      /___/__,__/\_,_/\_, /
##                     /___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package component

import (
	"io"
	"log"
)

// StdLogAdapter AFAIRE
type StdLogAdapter interface {
	io.Writer
	// NewStdLogger AFAIRE
	NewStdLogger(prefix string, flag int) *log.Logger
}

// Logger AFAIRE
type Logger interface {
	io.Writer
	// New AFAIRE
	New(lvl string, ctx ...interface{}) Logger
	// Trace AFAIRE
	Trace(msg string, ctx ...interface{})
	// Debug AFAIRE
	Debug(msg string, ctx ...interface{})
	// Info AFAIRE
	Info(msg string, ctx ...interface{})
	// Notice AFAIRE
	Notice(msg string, ctx ...interface{})
	// Warning AFAIRE
	Warning(msg string, ctx ...interface{})
	// Error AFAIRE
	Error(msg string, ctx ...interface{})
	// Critical AFAIRE
	Critical(msg string, ctx ...interface{})
	// Close AFAIRE
	Close()
	// NewStdLogger AFAIRE
	NewStdLogger(prefix string, flag int) *log.Logger
	// NewStdLogAdapter AFAIRE
	NewStdLogAdapter(lvl string, ctx ...interface{}) StdLogAdapter
}

/*
######################################################################################################## @(°_°)@ #######
*/
