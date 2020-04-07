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

func init() {
	AddLoader("none", defaultLoader)
	AddLoader("json", defaultLoader)
	AddLoader("yaml", defaultLoader)
}

/*
######################################################################################################## @(°_°)@ #######
*/
