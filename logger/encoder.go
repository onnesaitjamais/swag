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

import "time"

// Encoder représente l'interface qui doit être implémentée par tous les encodeurs.
type Encoder interface {
	// AFAIRE
	Encode(runner string, lvl Level, msg string, ctx []interface{}, dt time.Time, out Output) ([]byte, error)
}

/*
######################################################################################################## @(°_°)@ #######
*/
