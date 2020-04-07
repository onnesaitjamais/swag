/*
#######
##        ____    _____ ____ _
##       (_-< |/|/ / _ `/ _ `/
##      /___/__,__/\_,_/\_, /
##                     /___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package cfgstring

import (
	"strings"

	"github.com/arnumina/swag/util/failure"
	"github.com/arnumina/swag/util/options"
)

// Parse AFAIRE
// Format: "<type>[:option1=value1,option2=value2,...]"
func Parse(cfgString string) (string, options.Options, error) {
	if cfgString == "" {
		return "", nil,
			failure.New(nil).
				Msg("the configuration string is empty") ///////////////////////////////////////////////////////////////
	}

	opts := make(options.Options)

	ls := strings.Split(cfgString, ":")

	if len(ls) != 1 {
		if len(ls) != 2 {
			return "", nil,
				failure.New(nil).
					Set("string", cfgString).
					Msg("this configuration string is not valid") //////////////////////////////////////////////////////
		}

		for _, opt := range strings.Split(ls[1], ",") {
			kv := strings.Split(opt, "=")
			if len(kv) != 2 {
				return "", nil,
					failure.New(nil).
						Set("string", cfgString).
						Set("option", opt).
						Msg("this option of this configuration string is not valid") ///////////////////////////////////
			}

			opts[kv[0]] = kv[1]
		}
	}

	return ls[0], opts, nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
