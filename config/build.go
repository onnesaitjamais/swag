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
)

const (
	_defaultPort    = -1
	_defaultPortMax = 65530
	_defaultPortMin = 65000
)

func (c *config) setPort(opts options.Options, runner *runner.Runner) error {
	const option = "port"

	d, err := runner.CfgValue().DInt(_defaultPort, "services", runner.Name(), option)
	if err != nil {
		return err
	}

	opts.SetOption(
		option,
		"CONFIG_PORT",
		runner.Name(),
		d,
	)

	port, err := opts.Int(option)
	if err != nil {
		return err
	}

	c.port = port

	return nil
}

func (c *config) setPortMax(opts options.Options, runner *runner.Runner) error {
	const option = "port_max"

	d, err := runner.CfgValue().DInt(_defaultPortMax, "components", "config", option)
	if err != nil {
		return err
	}

	opts.SetOption(
		option,
		"CONFIG_PORT_MAX",
		runner.Name(),
		d,
	)

	portMax, err := opts.Int(option)
	if err != nil {
		return err
	}

	c.portMax = portMax

	return nil
}

func (c *config) setPortMin(opts options.Options, runner *runner.Runner) error {
	const option = "port_min"

	d, err := runner.CfgValue().DInt(_defaultPortMin, "components", "config", option)
	if err != nil {
		return err
	}

	opts.SetOption(
		option,
		"CONFIG_PORT_MIN",
		runner.Name(),
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
	config := &config{}

	if err := config.setPort(opts, runner); err != nil {
		return nil, err
	}

	if err := config.setPortMax(opts, runner); err != nil {
		return nil, err
	}

	if err := config.setPortMin(opts, runner); err != nil {
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
