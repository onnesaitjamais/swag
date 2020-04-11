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

import "log/syslog"

// GetFacilityFromString TODO
func GetFacilityFromString(facility string) syslog.Priority {
	switch facility {
	case "local0":
		return syslog.LOG_LOCAL0
	case "local1":
		return syslog.LOG_LOCAL1
	case "local2":
		return syslog.LOG_LOCAL2
	case "local3":
		return syslog.LOG_LOCAL3
	case "local4":
		return syslog.LOG_LOCAL4
	case "local5":
		return syslog.LOG_LOCAL5
	case "local6":
		return syslog.LOG_LOCAL6
	default:
		return syslog.LOG_LOCAL7
	}
}

// SyslogOutput AFAIRE
type SyslogOutput struct {
	*syslog.Writer
	*OutputProperties
}

// NewSyslogOutput AFAIRE
func NewSyslogOutput(facility string) (*SyslogOutput, error) {
	writer, err := syslog.New(GetFacilityFromString(facility), "quag")
	if err != nil {
		return nil, err
	}

	output := &SyslogOutput{
		Writer:           writer,
		OutputProperties: NewOutputProperties(false, false, false),
	}

	return output, nil
}

// Log AFAIRE
func (o *SyslogOutput) Log(level Level, buf []byte) error {
	switch level {
	case ILevel:
		return o.Info(string(buf))
	case NLevel:
		return o.Notice(string(buf))
	case WLevel:
		return o.Warning(string(buf))
	case ELevel:
		return o.Err(string(buf))
	case CLevel:
		return o.Crit(string(buf))
	default:
		return o.Debug(string(buf)) // TRACE & DEBUG
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
