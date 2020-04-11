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

type stdoutOutput struct {
	*WriterOutput
}

func newStdoutOutput() *stdoutOutput {
	return &stdoutOutput{
		WriterOutput: NewWriterOutput(os.Stdout, NewOutputProperties(true, true, true)),
	}
}

func (o *stdoutOutput) Close() error {
	return nil
}

// Stdout correspond à la sortie standard "stdout".
var Stdout Output = newStdoutOutput()

/*
######################################################################################################## @(°_°)@ #######
*/
