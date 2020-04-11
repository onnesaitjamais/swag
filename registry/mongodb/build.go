/*
#######
##        ____    _____ ____ _
##       (_-< |/|/ / _ `/ _ `/
##      /___/__,__/\_,_/\_, /
##                     /___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package mongodb

import (
	"github.com/arnumina/swag/runner"
	"github.com/arnumina/swag/util/failure"
	"github.com/arnumina/swag/util/options"
	"github.com/arnumina/swag/util/value"
)

const (
	_defaultInterval = 30
)

func (r *registry) setInterval(opts options.Options, runner string, cfg *value.Value) error {
	const option = "interval"

	d, err := cfg.DInt(_defaultInterval, option)
	if err != nil {
		return err
	}

	opts.SetOption(
		option,
		"REGISTRY_INTERVAL",
		runner,
		d,
	)

	interval, err := opts.Int(option)
	if err != nil {
		return err
	}

	r.interval = interval

	return nil
}

func (r *registry) setURI(opts options.Options, runner string, cfg *value.Value) error {
	const option = "URI"

	d, err := cfg.DString("", option)
	if err != nil {
		return err
	}

	opts.SetOption(
		option,
		"REGISTRY_URI",
		runner,
		d,
	)

	uri, err := opts.String(option)
	if err != nil {
		return err
	}

	if uri == "" {
		return failure.New(nil).
			Set("name", option).
			Msg("this option is required") /////////////////////////////////////////////////////////////////////////////
	}

	r.uri = uri

	return nil
}

// Build AFAIRE
func Build(opts options.Options, runner *runner.Runner) (interface{}, error) {
	cfg, err := runner.ComponentCfg("registry")
	if err != nil {
		return nil, err
	}

	registry := &registry{}

	if err := registry.setInterval(opts, runner.Name(), cfg); err != nil {
		return nil, err
	}

	if err := registry.setURI(opts, runner.Name(), cfg); err != nil {
		return nil, err
	}

	return registry.build()
}

/*
######################################################################################################## @(°_°)@ #######
*/
