/*
#######
##        ____    _____ ____ _
##       (_-< |/|/ / _ `/ _ `/
##      /___/__,__/\_,_/\_, /
##                     /___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package broker

import "time"

// Message AFAIRE
type Message struct {
	Event     string
	Payload   interface{}
	CreatedAt time.Time
}

/*
######################################################################################################## @(°_°)@ #######
*/
