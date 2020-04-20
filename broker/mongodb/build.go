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
	"regexp"

	"github.com/arnumina/swag/runner"
	"github.com/arnumina/swag/util/failure"
	"github.com/arnumina/swag/util/options"
)

func (b *broker) setURI(opts options.Options, runner *runner.Runner) error {
	const option = "URI"

	d, err := runner.CfgValue().DString("", "components", "broker", option)
	if err != nil {
		return err
	}

	opts.SetOption(
		option,
		"BROKER_URI",
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

	b.uri = uri

	return nil
}

func (b *broker) setBindings(runner *runner.Runner) error {
	bindings := make(map[string][]*regexp.Regexp)

	vm, err := runner.CfgValue().MapString("components", "broker", "bindings")
	if err != nil {
		return err
	}

	for queue, v := range vm {
		vs, err := v.Slice()
		if err != nil {
			return err
		}

		var events []*regexp.Regexp

		for _, ve := range vs {
			event, err := ve.String()
			if err != nil {
				return err
			}

			re, err := regexp.Compile(event)
			if err != nil {
				return err
			}

			events = append(events, re)
		}

		bindings[queue] = events
	}

	b.bindings = bindings

	return nil
}

// Build AFAIRE
func Build(opts options.Options, runner *runner.Runner) (interface{}, error) {
	broker := &broker{
		runner: runner,
	}

	if err := broker.setURI(opts, runner); err != nil {
		return nil, err
	}

	if err := broker.setBindings(runner); err != nil {
		return nil, err
	}

	return broker.build()
}

/*
######################################################################################################## @(°_°)@ #######
*/
