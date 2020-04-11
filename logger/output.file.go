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

// FileOutput AFAIRE
type FileOutput struct {
	*WriterOutput
	file *os.File
}

// NewFileOutput AFAIRE
func NewFileOutput(name string) (*FileOutput, error) {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	output := &FileOutput{
		WriterOutput: NewWriterOutput(file, NewOutputProperties(true, true, true)),
		file:         file,
	}

	return output, nil
}

// Close AFAIRE
func (o *FileOutput) Close() error {
	return o.file.Close()
}

/*
######################################################################################################## @(°_°)@ #######
*/
