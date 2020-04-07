/*
#######
##        ____    _____ ____ _
##       (_-< |/|/ / _ `/ _ `/
##      /___/__,__/\_,_/\_, /
##                     /___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package factory

import (
	"github.com/arnumina/swag/application/component/builder"
	"github.com/arnumina/swag/runner"
	"github.com/arnumina/swag/util/failure"
	"github.com/arnumina/swag/util/options"
)

// ComponentBuilder AFAIRE
type ComponentBuilder interface {
	// Name AFAIRE
	Name() string
	// Initialize AFAIRE
	Initialize(builder string, opts options.Options)
	// IsInitialized AFAIRE
	IsInitialized() bool
	// Builder AFAIRE
	Builder() string
	// Build AFAIRE
	Build(runner *runner.Runner) (interface{}, error)
	// AddBuilder AFAIRE
	AddBuilder(name string, builder builder.Builder)
}

var _factory = make(map[string]ComponentBuilder)

func addComponentBuilder(cb ComponentBuilder) {
	_factory[cb.Name()] = cb
}

func get(component string) (ComponentBuilder, error) {
	if cb, ok := _factory[component]; ok {
		return cb, nil
	}

	return nil,
		failure.New(nil).
			Set("name", component).
			Msg("this component does not exist") ///////////////////////////////////////////////////////////////////////
}

// Initialize AFAIRE
func Initialize(component, builder string, opts options.Options) error {
	cb, err := get(component)
	if err != nil {
		return err
	}

	cb.Initialize(builder, opts)

	return nil
}

// Build AFAIRE
func Build(component string, runner *runner.Runner) (interface{}, error) {
	cb, err := get(component)
	if err != nil {
		return nil, err
	}

	if !cb.IsInitialized() {
		return nil, nil
	}

	logger := runner.Logger()

	if logger != nil {
		logger.Info( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
			"Component building",
			"name", component,
			"builder", cb.Builder(),
		)
	}

	instance, err := cb.Build(runner)
	if err != nil {
		return nil,
			failure.New(err).
				Set("name", component).
				Msg("this component cannot be built") //////////////////////////////////////////////////////////////////
	}

	return instance, nil
}

// AddComponentBuilder AFAIRE
func AddComponentBuilder(component, name string, builder builder.Builder) error {
	cb, err := get(component)
	if err != nil {
		return err
	}

	cb.AddBuilder(name, builder)

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
