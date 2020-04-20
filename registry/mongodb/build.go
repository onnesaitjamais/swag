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
	"time"

	"github.com/arnumina/swag/runner"
	"github.com/arnumina/swag/util/failure"
	"github.com/arnumina/swag/util/options"
)

const (
	_defaultInterval = 30
)

func (r *registry) setInterval(opts options.Options, runner *runner.Runner) error {
	const option = "interval"

	d, err := runner.CfgValue().DInt(_defaultInterval, "components", "registry", option)
	if err != nil {
		return err
	}

	opts.SetOption(
		option,
		"REGISTRY_INTERVAL",
		runner.Name(),
		d,
	)

	interval, err := opts.Int(option)
	if err != nil {
		return err
	}

	r.interval = time.Duration(interval) * time.Second

	return nil
}

func (r *registry) setURI(opts options.Options, runner *runner.Runner) error {
	const option = "URI"

	d, err := runner.CfgValue().DString("", "components", "registry", option)
	if err != nil {
		return err
	}

	opts.SetOption(
		option,
		"REGISTRY_URI",
		runner.Name(),
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
	registry := &registry{
		runner: runner,
	}

	if err := registry.setInterval(opts, runner); err != nil {
		return nil, err
	}

	if err := registry.setURI(opts, runner); err != nil {
		return nil, err
	}

	return registry.build()
}

/*
######################################################################################################## @(°_°)@ #######
*/
