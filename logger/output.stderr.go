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

import "os"

type stderrOutput struct {
	*WriterOutput
}

func newStderrOutput() *stderrOutput {
	return &stderrOutput{
		WriterOutput: NewWriterOutput(os.Stderr, NewOutputProperties(true, true, true)),
	}
}

func (o *stderrOutput) Close() error {
	return nil
}

// Stderr correspond à la sortie standard "stderr".
var Stderr Output = newStderrOutput()

/*
######################################################################################################## @(°_°)@ #######
*/
