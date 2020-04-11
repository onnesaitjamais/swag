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

import (
	"fmt"

	"github.com/arnumina/swag/runner"
	"github.com/arnumina/swag/util/failure"
	"github.com/arnumina/swag/util/options"
	"github.com/arnumina/swag/util/value"
)

const (
	_defaultLevel                = "debug"
	_defaultEncoder              = "logfmt"
	_defaultOutputFileName       = "/var/log/swag/swag.log"
	_defaultOutputSyslogFacility = "local0"
	_defaultOutput               = "syslog"
)

// EncBuilder AFAIRE
type EncBuilder func(string, options.Options) (Encoder, error)

// OutBuilder AFAIRE
type OutBuilder func(string, options.Options) (Output, error)

var (
	_encBuilders = make(map[string]EncBuilder)
	_outBuilders = make(map[string]OutBuilder)
)

// AddEncBuilder AFAIRE
func AddEncBuilder(name string, builder EncBuilder) {
	_encBuilders[name] = builder
}

// AddOutBuilder AFAIRE
func AddOutBuilder(name string, builder OutBuilder) {
	_outBuilders[name] = builder
}

func defaultEncBuilder(t string, _ options.Options) (Encoder, error) {
	switch t {
	case "logfmt":
		return NewLogFmtEncoder(), nil
	default:
		return nil, failure.New(nil).
			Set("type", t).
			Msg("this type of encoder does not exist") /////////////////////////////////////////////////////////////////
	}
}

func buildOutputFile(opts options.Options) (Output, error) {
	ofn, err := opts.String("output_file_name")
	if err != nil {
		return nil, err
	}

	return NewFileOutput(ofn)
}

func buildOutputSyslog(opts options.Options) (Output, error) {
	osf, err := opts.String("output_syslog_facility")
	if err != nil {
		return nil, err
	}

	return NewSyslogOutput(osf)
}

func defaultOutBuilder(t string, opts options.Options) (Output, error) {
	switch t {
	case "file":
		return buildOutputFile(opts)
	case "stderr":
		return Stderr, nil
	case "stdout":
		return Stdout, nil
	case "syslog":
		return buildOutputSyslog(opts)
	default:
		return nil, failure.New(nil).
			Set("type", t).
			Msg("this type of output does not exist") //////////////////////////////////////////////////////////////////
	}
}

func (l *logger) setLevel(opts options.Options, runner string, cfg *value.Value) error {
	const option = "level"

	d, err := cfg.DString(_defaultLevel, option)
	if err != nil {
		return err
	}

	opts.SetOption(
		option,
		"LOGGER_LEVEL",
		runner,
		d,
	)

	level, err := opts.String(option)
	if err != nil {
		return err
	}

	l.lvl = GetLevelFromString(level)

	return nil
}

func (l *logger) setEncoder(opts options.Options, runner string, cfg *value.Value) error {
	const option = "encoder"

	d, err := cfg.DString(_defaultEncoder, option)
	if err != nil {
		return err
	}

	opts.SetOption(
		option,
		"LOGGER_ENCODER",
		runner,
		d,
	)

	t, err := opts.String(option)
	if err != nil {
		return err
	}

	builder, ok := _encBuilders[t]
	if !ok {
		return failure.New(nil).
			Set("type", t).
			Msg("there is no builder for this type of encoder") ////////////////////////////////////////////////////////
	}

	enc, err := builder(t, opts)
	if err != nil {
		return err
	}

	l.enc = enc

	return nil
}

func (l *logger) setOutputFileName(opts options.Options, runner string, cfg *value.Value) error {
	const option = "output_file_name"

	d, err := cfg.DString(_defaultOutputFileName, option)
	if err != nil {
		return err
	}

	opts.SetOption(
		option,
		"LOGGER_OUTPUT_FILE_NAME",
		runner,
		d,
	)

	return nil
}

func (l *logger) setOutputSyslogFacility(opts options.Options, runner string, cfg *value.Value) error {
	const option = "output_syslog_facility"

	d, err := cfg.DString(_defaultOutputSyslogFacility, option)
	if err != nil {
		return err
	}

	opts.SetOption(
		option,
		"LOGGER_OUTPUT_SYSLOG_FACILITY",
		runner,
		d,
	)

	return nil
}

func (l *logger) setOutput(opts options.Options, runner string, cfg *value.Value) error {
	const option = "output"

	d, err := cfg.DString(_defaultOutput, option)
	if err != nil {
		return err
	}

	opts.SetOption(
		option,
		"LOGGER_OUTPUT",
		runner,
		d,
	)

	t, err := opts.String(option)
	if err != nil {
		return err
	}

	builder, ok := _outBuilders[t]
	if !ok {
		return failure.New(nil).
			Set("type", t).
			Msg("there is no builder for this type of output") /////////////////////////////////////////////////////////
	}

	out, err := builder(t, opts)
	if err != nil {
		return err
	}

	l.out = out

	return nil
}

// Build AFAIRE
func Build(opts options.Options, runner *runner.Runner) (interface{}, error) {
	cfg, err := runner.ComponentCfg("logger")
	if err != nil {
		return nil, err
	}

	logger := &logger{
		runner: fmt.Sprintf("%s.%s", runner.Name(), runner.ID()[:8]),
	}

	//////////////// Il faudrait ajouter l'appel d'une callback pour les options des builders externes /////////////////

	if err := logger.setLevel(opts, runner.Name(), cfg); err != nil {
		return nil, err
	}

	if err := logger.setEncoder(opts, runner.Name(), cfg); err != nil {
		return nil, err
	}

	if err := logger.setOutputFileName(opts, runner.Name(), cfg); err != nil {
		return nil, err
	}

	if err := logger.setOutputSyslogFacility(opts, runner.Name(), cfg); err != nil {
		return nil, err
	}

	if err := logger.setOutput(opts, runner.Name(), cfg); err != nil {
		return nil, err
	}

	return logger, nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
