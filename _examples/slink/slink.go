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
	"fmt"
	"net/http"
	"os"

	"github.com/arnumina/swag"
	"github.com/arnumina/swag/util/options"
)

var version, builtAt string

func run() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		fmt.Fprintf(w, "Welcome to the home page!")
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	service, err := swag.NewService(
		"slink",
		version,
		builtAt,
		swag.Config(
			"default",
			options.Options{
				"port": 0,
			},
		),
		swag.Server(
			"http",
			options.Options{
				"handler":    mux,
				"health_URI": "/health",
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
