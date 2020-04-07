/*
#######
##        ____    _____ ____ _
##       (_-< |/|/ / _ `/ _ `/
##      /___/__,__/\_,_/\_, /
##                     /___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package options

import (
	"errors"
	"strconv"

	"github.com/arnumina/swag/util"
	"github.com/arnumina/swag/util/failure"
)

// ErrNotFound AFAIRE
var ErrNotFound = errors.New("option not found")

// Options AFAIRE
type Options map[string]interface{}

// Get AFAIRE
func (o Options) Get(name string) (interface{}, error) {
	if ov, ok := o[name]; ok {
		return ov, nil
	}

	return nil,
		failure.New(ErrNotFound).
			Set("name", name).
			Msg("this option does not seem to exist") //////////////////////////////////////////////////////////////////
}

// Bool AFAIRE
func (o Options) Bool(name string) (bool, error) {
	ov, err := o.Get(name)
	if err != nil {
		return false, err
	}

	switch v := ov.(type) {
	case bool:
		return v, nil
	case string:
		return strconv.ParseBool(v)
	default:
		return false,
			failure.New(nil).
				Set("name", name).
				Msg("this option is not of type 'bool'") ///////////////////////////////////////////////////////////////
	}
}

// Int AFAIRE
func (o Options) Int(name string) (int, error) {
	ov, err := o.Get(name)
	if err != nil {
		return 0, err
	}

	switch v := ov.(type) {
	case int:
		return v, nil
	case string:
		return strconv.Atoi(v)
	default:
		return 0,
			failure.New(nil).
				Set("name", name).
				Msg("this option is not of type 'int'") ////////////////////////////////////////////////////////////////
	}
}

// String AFAIRE
func (o Options) String(name string) (string, error) {
	ov, err := o.Get(name)
	if err != nil {
		return "", err
	}

	if v, ok := ov.(string); ok {
		return v, nil
	}

	return "",
		failure.New(nil).
			Set("name", name).
			Msg("this option is not of type 'string'") /////////////////////////////////////////////////////////////////
}

// SetOption AFAIRE
func (o Options) SetOption(name, envSuffix, runner string, d interface{}) {
	v, ok := o[name]
	if ok {
		d = v
	}

	o[name] = util.EnvValue(runner, envSuffix, d)
}

/*
######################################################################################################## @(°_°)@ #######
*/
