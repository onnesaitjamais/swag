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

// Level représente un niveau de log.
type Level int

// Les différents niveaux de log utilisables.
const (
	// Trace
	TLevel Level = iota
	// Debug
	DLevel
	// Info
	ILevel
	// Notice
	NLevel
	// Warning
	WLevel
	// Error
	ELevel
	// Fatal
	FLevel
)

// GetLevelFromString retourne le niveau de log correspondant à la chaîne de caractères
// passée en paramètre.
// Si celle-ci n'est pas valide, le niveau de log "TLevel" est utilisé
// par defaut.
func GetLevelFromString(level string) Level {
	switch level {
	case "debug":
		return DLevel
	case "info":
		return ILevel
	case "notice":
		return NLevel
	case "warning":
		return WLevel
	case "error":
		return ELevel
	case "fatal":
		return FLevel
	default:
		return TLevel
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
