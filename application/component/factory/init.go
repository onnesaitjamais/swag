/*
#######
##        ____    _____ ____ _
##       (_-< |/|/ / _ `/ _ `/
##      /___/__,__/\_,_/\_, /
##                     /___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package factory

import "github.com/arnumina/swag/application/component"

func init() {
	addComponentBuilder(component.NewBroker())
	addComponentBuilder(component.NewConfig())
	addComponentBuilder(component.NewLogger())
	addComponentBuilder(component.NewRegistry())
	addComponentBuilder(component.NewServer())
}

/*
######################################################################################################## @(°_°)@ #######
*/
