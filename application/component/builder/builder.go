/*
#######
##        ____    _____ ____ _
##       (_-< |/|/ / _ `/ _ `/
##      /___/__,__/\_,_/\_, /
##                     /___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package builder

import (
	"github.com/arnumina/swag/runner"
	"github.com/arnumina/swag/util/failure"
	"github.com/arnumina/swag/util/options"
)

// Builder AFAIRE
type Builder func(opts options.Options, runner *runner.Runner) (interface{}, error)

// ComponentBuilder AFAIRE
type ComponentBuilder struct {
	name        string
	builders    map[string]Builder
	builder     string
	opts        options.Options
	initialized bool
}

// NewComponentBuilder AFAIRE
func NewComponentBuilder(name string, builders map[string]Builder) *ComponentBuilder {
	return &ComponentBuilder{
		name:     name,
		builders: builders,
	}
}

// Name AFAIRE
func (cb *ComponentBuilder) Name() string {
	return cb.name
}

// Initialize AFAIRE
func (cb *ComponentBuilder) Initialize(builder string, opts options.Options) {
	cb.builder = builder
	cb.opts = opts
	cb.initialized = true
}

// IsInitialized AFAIRE
func (cb *ComponentBuilder) IsInitialized() bool {
	return cb.initialized
}

// Builder AFAIRE
func (cb *ComponentBuilder) Builder() string {
	return cb.builder
}

// Build AFAIRE
func (cb *ComponentBuilder) Build(runner *runner.Runner) (interface{}, error) {
	builder, ok := cb.builders[cb.builder]
	if !ok {
		return nil,
			failure.New(nil).
				Set("component", cb.name).
				Set("builder", cb.builder).
				Msg("this builder does not exist for this component") //////////////////////////////////////////////////
	}

	return builder(cb.opts, runner)
}

// AddBuilder AFAIRE
func (cb *ComponentBuilder) AddBuilder(name string, builder Builder) {
	cb.builders[name] = builder
}

/*
######################################################################################################## @(°_°)@ #######
*/
