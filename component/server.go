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

// Server AFAIRE
type Server interface {
	// Start AFAIRE
	Start(port int) error
	// CheckHealth
	CheckHealth(fqdn string, port int) error
	// Stop AFAIRE
	Stop() error
	// Close AFAIRE
	Close() error
}

/*
######################################################################################################## @(°_°)@ #######
*/
