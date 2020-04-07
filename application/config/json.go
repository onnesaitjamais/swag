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
	"io/ioutil"

	"github.com/arnumina/swag/util/options"
	"github.com/arnumina/swag/util/value"
)

func loadJSONFile(opts options.Options) (*value.Value, error) {
	filename, err := opts.String("file")
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return value.FromJSON(data)
}

/*
######################################################################################################## @(°_°)@ #######
*/
