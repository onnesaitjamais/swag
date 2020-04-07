/*
#######
##        ____    _____ ____ _
##       (_-< |/|/ / _ `/ _ `/
##      /___/__,__/\_,_/\_, /
##                     /___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package value

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/dolmen-go/jsonptr"
	"gopkg.in/yaml.v3"

	"github.com/arnumina/swag/util/failure"
)

// ErrNotFound AFAIRE
var ErrNotFound = errors.New("value not found")

// Value AFAIRE
type Value struct {
	data interface{}
}

// FromJSON AFAIRE
func FromJSON(data []byte) (*Value, error) {
	value := &Value{}
	if err := json.Unmarshal(data, &value.data); err != nil {
		return nil, err
	}

	return value, nil
}

// FromYAML AFAIRE
func FromYAML(data []byte) (*Value, error) {
	value := &Value{}
	if err := yaml.Unmarshal(data, &value.data); err != nil {
		return nil, err
	}

	return value, nil
}

// Data AFAIRE
func (v *Value) Data() interface{} {
	return v.data
}

// Get AFAIRE
func (v *Value) Get(keys ...string) (string, *Value, error) {
	ptr := fmt.Sprintf("/%s", strings.Join(append([]string{}, keys...), "/"))

	if ptr == "/" {
		return ptr, v, nil
	}

	value, err := jsonptr.Get(v.data, ptr)
	if err != nil {
		if errors.Is(err, jsonptr.ErrProperty) {
			return ptr, nil,
				failure.New(ErrNotFound).
					Set("pointer", ptr).
					Msg("this value does not seem to exist") ///////////////////////////////////////////////////////////
		}

		return ptr, nil, err
	}

	return ptr, &Value{data: value}, nil
}

// MapString AFAIRE
func (v *Value) MapString(keys ...string) (map[string]*Value, error) {
	ptr, value, err := v.Get(keys...)
	if err != nil {
		return nil, err
	}

	ms, ok := value.data.(map[string]interface{})
	if !ok {
		return nil,
			failure.New(nil).
				Set("pointer", ptr).
				Msg("this pointer do not reference a value of type 'map'") /////////////////////////////////////////////
	}

	result := make(map[string]*Value)

	for k, v := range ms {
		result[k] = &Value{data: v}
	}

	return result, nil
}

// Slice AFAIRE
func (v *Value) Slice(keys ...string) ([]*Value, error) {
	ptr, value, err := v.Get(keys...)
	if err != nil {
		return nil, err
	}

	slice, ok := value.data.([]interface{})
	if !ok {
		return nil,
			failure.New(nil).
				Set("pointer", ptr).
				Msg("this pointer do not reference a value of type 'slice'") ///////////////////////////////////////////
	}

	result := []*Value{}

	for _, v := range slice {
		result = append(result, &Value{data: v})
	}

	return result, nil
}

// Bool AFAIRE
func (v *Value) Bool(keys ...string) (bool, error) {
	ptr, value, err := v.Get(keys...)
	if err != nil {
		return false, err
	}

	result, ok := value.data.(bool)
	if !ok {
		return false,
			failure.New(nil).
				Set("pointer", ptr).
				Msg("this pointer do not reference a value of type 'bool'") ////////////////////////////////////////////
	}

	return result, nil
}

// DBool AFAIRE
func (v *Value) DBool(d bool, keys ...string) (bool, error) {
	if v == nil {
		return d, nil
	}

	value, err := v.Bool(keys...)
	if errors.Is(err, ErrNotFound) {
		return d, nil
	}

	return value, err
}

// Int AFAIRE
func (v *Value) Int(keys ...string) (int, error) {
	ptr, value, err := v.Get(keys...)
	if err != nil {
		return 0, err
	}

	result, ok := value.data.(int)
	if !ok {
		return 0,
			failure.New(nil).
				Set("pointer", ptr).
				Msg("this pointer do not reference a value of type 'int'") /////////////////////////////////////////////
	}

	return result, nil
}

// DInt AFAIRE
func (v *Value) DInt(d int, keys ...string) (int, error) {
	if v == nil {
		return d, nil
	}

	value, err := v.Int(keys...)
	if errors.Is(err, ErrNotFound) {
		return d, nil
	}

	return value, err
}

// String AFAIRE
func (v *Value) String(keys ...string) (string, error) {
	ptr, value, err := v.Get(keys...)
	if err != nil {
		return "", err
	}

	result, ok := value.data.(string)
	if !ok {
		return "",
			failure.New(nil).
				Set("pointer", ptr).
				Msg("this pointer do not reference a value of type 'string'") //////////////////////////////////////////
	}

	return result, nil
}

// DString AFAIRE
func (v *Value) DString(d string, keys ...string) (string, error) {
	if v == nil {
		return d, nil
	}

	value, err := v.String(keys...)
	if errors.Is(err, ErrNotFound) {
		return d, nil
	}

	return value, err
}

/*
######################################################################################################## @(°_°)@ #######
*/
