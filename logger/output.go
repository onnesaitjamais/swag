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

import "io"

// Output décrit l'interface qui doit être implémentée par toutes les sorties.
type Output interface {
	// Est-ce que la date et l'heure doivent être écrites ?
	LogDateTime() bool
	// Est-ce que le niveau de log doit être écrit ?
	LogLevel() bool
	// Est-ce qu'un "\n" doit être ajouté à la fin du message ?
	AddNewLine() bool
	// Cette fonction écrit le message résultant de l'encodage.
	Log(level Level, buf []byte) error
	// Cette fonction, si nécessaire, ferme la sortie.
	Close() error
}

// OutputProperties permet de définir les propriétés d'une sortie.
type OutputProperties struct {
	flagDateTime bool
	flagLevel    bool
	flagNewLine  bool
}

// NewOutputProperties permet de créer une instance du type "OutputProperties".
func NewOutputProperties(fDateTime bool, fLevel bool, fNewLine bool) *OutputProperties {
	return &OutputProperties{
		flagDateTime: fDateTime,
		flagLevel:    fLevel,
		flagNewLine:  fNewLine,
	}
}

// LogDateTime permet d'indiquer à l'encodeur si il doit encodé la date et l'heure.
func (op *OutputProperties) LogDateTime() bool {
	return op.flagDateTime
}

// LogLevel permet d'indiquer à l'encodeur si il doit encodé le niveau de log.
func (op *OutputProperties) LogLevel() bool {
	return op.flagLevel
}

// AddNewLine permet d'indiquer à l'encodeur si il doit ajouté un "\n" à la fin du message.
func (op *OutputProperties) AddNewLine() bool {
	return op.flagNewLine
}

// WriterOutput permet de définir des sorties basées sur un "Writer".
type WriterOutput struct {
	io.Writer
	*OutputProperties
}

// NewWriterOutput permet de créer une instance du type "WriterOutput".
func NewWriterOutput(iow io.Writer, op *OutputProperties) *WriterOutput {
	return &WriterOutput{
		Writer:           iow,
		OutputProperties: op,
	}
}

// Log écrit le message résultant de l'encodage.
func (o *WriterOutput) Log(level Level, buf []byte) error {
	_, err := o.Write(buf)
	return err
}

/*
######################################################################################################## @(°_°)@ #######
*/
