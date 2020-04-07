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
	"github.com/arnumina/swag/config"
)

// Config AFAIRE
type Config struct {
	*builder.ComponentBuilder
}

// NewConfig AFAIRE
func NewConfig() *Config {
	return &Config{
		ComponentBuilder: builder.NewComponentBuilder(
			"config",
			map[string]builder.Builder{
				"default": config.Build,
			},
		),
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
