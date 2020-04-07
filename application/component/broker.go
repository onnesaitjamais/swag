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
	"github.com/arnumina/swag/broker/mongodb"
)

// Broker AFAIRE
type Broker struct {
	*builder.ComponentBuilder
}

// NewBroker AFAIRE
func NewBroker() *Broker {
	return &Broker{
		ComponentBuilder: builder.NewComponentBuilder(
			"broker",
			map[string]builder.Builder{
				"mongodb": mongodb.Build,
			},
		),
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
