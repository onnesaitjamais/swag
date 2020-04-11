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

import (
	"github.com/arnumina/swag/runner"
	"github.com/arnumina/swag/util/failure"
	"github.com/arnumina/swag/util/options"
	"github.com/arnumina/swag/util/value"
)

const (
	_defaultPort    = -1
	_defaultPortMax = 65530
	_defaultPortMin = 65000
)

func (c *config) setPort(opts options.Options, runner string, cfg *value.Value) error {
	const option = "port"

	d, err := cfg.DInt(_defaultPort, option)
	if err != nil {
		return err
	}

	opts.SetOption(
		option,
		"CONFIG_PORT",
		runner,
		d,
	)

	port, err := opts.Int(option)
	if err != nil {
		return err
	}

	c.port = port

	return nil
}

func (c *config) setPortMax(opts options.Options, runner string, cfg *value.Value) error {
	const option = "port_max"

	d, err := cfg.DInt(_defaultPortMax, option)
	if err != nil {
		return err
	}

	opts.SetOption(
		option,
		"CONFIG_PORT_MAX",
		runner,
		d,
	)

	portMax, err := opts.Int(option)
	if err != nil {
		return err
	}

	c.portMax = portMax

	return nil
}

func (c *config) setPortMin(opts options.Options, runner string, cfg *value.Value) error {
	const option = "port_min"

	d, err := cfg.DInt(_defaultPortMin, option)
	if err != nil {
		return err
	}

	opts.SetOption(
		option,
		"CONFIG_PORT_MIN",
		runner,
		d,
	)

	portMin, err := opts.Int(option)
	if err != nil {
		return err
	}

	c.portMin = portMin

	return nil
}

// Build AFAIRE
func Build(opts options.Options, runner *runner.Runner) (interface{}, error) {
	cfg, err := runner.ComponentCfg("config")
	if err != nil {
		return nil, err
	}

	config := &config{}

	if err := config.setPort(opts, runner.Name(), cfg); err != nil {
		return nil, err
	}

	if err := config.setPortMax(opts, runner.Name(), cfg); err != nil {
		return nil, err
	}

	if err := config.setPortMin(opts, runner.Name(), cfg); err != nil {
		return nil, err
	}

	if config.portMax <= config.portMin {
		return nil,
			failure.New(nil).
				Set("port_min", config.portMin).
				Set("port_max", config.portMax).
				Msg("the values of the 'port_min' and 'port_max' options are not valid") ///////////////////////////////
	}

	return config, nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
