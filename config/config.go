/*
#######
##        ____    _____ ____ _
##       (_-< |/|/ / _ `/ _ `/
##      /___/__,__/\_,_/\_, /
##                     /___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package config

type config struct {
	port    int
	portMin int
	portMax int
}

// Port AFAIRE
func (c *config) Port() int {
	return c.port
}

// SetPort AFAIRE
func (c *config) SetPort(port int) {
	c.port = port
}

// PortMin AFAIRE
func (c *config) PortMin() int {
	return c.portMin
}

// PortMax AFAIRE
func (c *config) PortMax() int {
	return c.portMax
}

// Close AFAIRE
func (c *config) Close() error {
	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
