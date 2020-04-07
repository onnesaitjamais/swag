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
	"github.com/arnumina/swag/application/component/builder"
	"github.com/arnumina/swag/logger"
)

// Logger AFAIRE
type Logger struct {
	*builder.ComponentBuilder
}

// NewLogger AFAIRE
func NewLogger() *Logger {
	return &Logger{
		ComponentBuilder: builder.NewComponentBuilder(
			"logger",
			map[string]builder.Builder{
				"default": logger.Build,
			},
		),
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
