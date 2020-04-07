/*
#######
##        ____    _____ ____ _
##       (_-< |/|/ / _ `/ _ `/
##      /___/__,__/\_,_/\_, /
##                     /___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package main

import (
	"os"

	"github.com/arnumina/swag"
	"github.com/arnumina/swag/util/options"
)

var version, builtAt string

func run() error {
	service, err := swag.NewService(
		"hamm",
		version,
		builtAt,
		swag.Logger(
			"default",
			options.Options{
				"level":  "trace",
				"output": "stdout",
			},
		),
	)
	if err != nil {
		return err
	}

	defer service.Close()

	if err := service.Run(); err != nil {
		return err
	}

	return nil
}

func main() {
	if run() != nil {
		os.Exit(-1)
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
