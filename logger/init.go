/*
#######
##        ____    _____ ____ _
##       (_-< |/|/ / _ `/ _ `/
##      /___/__,__/\_,_/\_, /
##                     /___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package logger

func init() {
	AddEncBuilder("logfmt", defaultEncBuilder)
	AddOutBuilder("file", defaultOutBuilder)
	AddOutBuilder("stderr", defaultOutBuilder)
	AddOutBuilder("stdout", defaultOutBuilder)
	AddOutBuilder("syslog", defaultOutBuilder)
}

/*
######################################################################################################## @(°_°)@ #######
*/
