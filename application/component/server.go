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

// Server AFAIRE
type Server struct {
	*builder.ComponentBuilder
}

// NewServer AFAIRE
func NewServer() *Server {
	return &Server{
		ComponentBuilder: builder.NewComponentBuilder(
			"server",
			map[string]builder.Builder{
				"http": mongodb.Build,
			},
		),
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
