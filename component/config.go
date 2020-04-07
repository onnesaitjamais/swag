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

// Config AFAIRE
type Config interface {
	// Port AFAIRE
	Port() int
	// SetPort AFAIRE
	SetPort(port int)
	// PortMin AFAIRE
	PortMin() int
	// PortMax AFAIRE
	PortMax() int
	// Close AFAIRE
	Close() error
}

/*
######################################################################################################## @(°_°)@ #######
*/
