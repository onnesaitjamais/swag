/*
#######
##        ____    _____ ____ _
##       (_-< |/|/ / _ `/ _ `/
##      /___/__,__/\_,_/\_, /
##                     /___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package swag

import (
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/arnumina/swag/application"
	"github.com/arnumina/swag/application/component/factory"
	"github.com/arnumina/swag/service"
	"github.com/arnumina/swag/util/options"
)

func initializeComponent(name, builder string, opts options.Options) func() error {
	return func() error {
		return factory.Initialize(name, builder, opts)
	}
}

// Broker AFAIRE
func Broker(builder string, opts options.Options) func() error {
	return initializeComponent("broker", builder, opts)
}

// Config AFAIRE
func Config(builder string, opts options.Options) func() error {
	return initializeComponent("config", builder, opts)
}

// Logger AFAIRE
func Logger(builder string, opts options.Options) func() error {
	return initializeComponent("logger", builder, opts)
}

// Registry AFAIRE
func Registry(builder string, opts options.Options) func() error {
	return initializeComponent("registry", builder, opts)
}

// Server AFAIRE
func Server(builder string, opts options.Options) func() error {
	return initializeComponent("server", builder, opts)
}

func initialize(name, version, builtAt string, cmpts ...func() error) (*service.Service, error) {
	rand.Seed(time.Now().UnixNano())

	application, err := application.New(name, version, builtAt)
	if err != nil {
		return nil, err
	}

	if err := application.Initialize(cmpts...); err != nil {
		return nil, err
	}

	service := service.New(application)

	return service, nil
}

// NewService AFAIRE
func NewService(name, version, builtAt string, cmpts ...func() error) (*service.Service, error) {
	service, err := initialize(name, version, builtAt, cmpts...)
	if err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return nil, err
		}

		fmt.Fprintf( ///////////////////////////////////////////////////////////////////////////////////////////////////
			os.Stderr,
			"Error when initializing this service: service=%s version=%s >>> %s\n",
			name,
			version,
			err,
		)

		return nil, err
	}

	return service, nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
