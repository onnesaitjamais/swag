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
	"github.com/arnumina/swag/util/cfgstring"
	"github.com/arnumina/swag/util/failure"
	"github.com/arnumina/swag/util/options"
	"github.com/arnumina/swag/util/value"
)

// Loader AFAIRE
type Loader func(string, options.Options) (*value.Value, error)

var _loaders = make(map[string]Loader)

// AddLoader AFAIRE
func AddLoader(name string, fn Loader) {
	_loaders[name] = fn
}

func defaultLoader(t string, opts options.Options) (*value.Value, error) {
	switch t {
	case "none":
		return value.FromJSON([]byte("{}"))
	case "json":
		return loadJSONFile(opts)
	case "yaml":
		return loadYAMLFile(opts)
	default:
		return nil,
			failure.New(nil).
				Set("type", t).
				Msg("there is no configuration loader for this type") //////////////////////////////////////////////////
	}
}

// Load AFAIRE
func Load(cfgString string) (*value.Value, error) {
	t, opts, err := cfgstring.Parse(cfgString)
	if err != nil {
		return nil, err
	}

	loader, ok := _loaders[t]
	if !ok {
		return nil,
			failure.New(nil).
				Set("type", t).
				Msg("there is no configuration loader for this type") //////////////////////////////////////////////////
	}

	return loader(t, opts)
}

/*
######################################################################################################## @(°_°)@ #######
*/
