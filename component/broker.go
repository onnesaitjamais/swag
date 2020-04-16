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
	"github.com/arnumina/swag/component/broker"
)

// Broker AFAIRE
type Broker interface {
	// Publish AFAIRE
	Publish(event string, data interface{})
	// Subscribe AFAIRE
	Subscribe(queue string, fn func(*broker.Message) bool)
	// Close AFAIRE
	Close() error
}

/*
######################################################################################################## @(°_°)@ #######
*/
