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
	"github.com/arnumina/swag/registry/mongodb"
)

// Registry AFAIRE
type Registry struct {
	*builder.ComponentBuilder
}

// NewRegistry AFAIRE
func NewRegistry() *Registry {
	return &Registry{
		ComponentBuilder: builder.NewComponentBuilder(
			"registry",
			map[string]builder.Builder{
				"mongodb": mongodb.Build,
			},
		),
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
