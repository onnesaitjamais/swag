/*
#######
##        ____    _____ ____ _
##       (_-< |/|/ / _ `/ _ `/
##      /___/__,__/\_,_/\_, /
##                     /___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package application

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/arnumina/swag/application/component/factory"
	"github.com/arnumina/swag/application/config"
	"github.com/arnumina/swag/component"
	"github.com/arnumina/swag/runner"
	"github.com/arnumina/swag/util"
	"github.com/arnumina/swag/util/failure"
	"github.com/arnumina/swag/util/options"
)

// Application AFAIRE
type Application struct {
	*runner.Runner
}

// New AFAIRE
func New(name, version, builtAt string) (*Application, error) {
	runner, err := runner.New(name, version, builtAt)
	if err != nil {
		return nil, err
	}

	application := &Application{
		Runner: runner,
	}

	return application, nil
}

func (a *Application) printVersion() {
	fmt.Printf(
		"swag.%s version %s built at %s by Archivage Numérique © INA %d\n\n",
		a.Name(),
		a.Version(),
		a.BuiltAt().String(),
		time.Now().Year(),
	)
}

func (a *Application) loadConfig(cfgString string) error {
	cfgValue, err := config.Load(cfgString)
	if err != nil {
		return err
	}

	a.SetCfgValue(cfgValue)

	return nil
}

func (a *Application) parseFlags() error {
	flagSet := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	flagSet.SetOutput(os.Stdout)
	flagSet.Usage = func() {
		fmt.Println(os.Args[0])
		fmt.Println("=================================================================================================")
		flagSet.PrintDefaults()
		fmt.Println("-------------------------------------------------------------------------------------------------")
	}

	version := flagSet.Bool(
		"version",
		false,
		"version number of this service",
	)

	sdInstance := flagSet.String(
		"instance",
		"",
		"systemd instance identifier (%i in unit file)",
	)

	cfgString := flagSet.String(
		"config",
		util.EnvValue(a.Name(), "CFG_STRING", "none").(string),
		"configuration string: <type>[:option1=value1,option2=value2,...]",
	)

	if err := flagSet.Parse(os.Args[1:]); err != nil {
		return err
	}

	if *version {
		a.printVersion()
		return flag.ErrHelp
	}

	a.SetSdInstance(*sdInstance)

	return a.loadConfig(*cfgString)
}

func (a *Application) initializeComponents(cmpts ...func() error) error {
	for _, fn := range append(
		[]func() error{
			func() error {
				return factory.Initialize("config", "default", options.Options{})
			},
			func() error {
				return factory.Initialize("logger", "default", options.Options{})
			},
			func() error {
				return factory.Initialize("registry", "mongodb", options.Options{})
			},
		},
		cmpts...,
	) {
		if err := fn(); err != nil {
			return err
		}
	}

	return nil
}

func (a *Application) buildComponents() error {
	// L'ordre de construction est important
	for _, name := range []string{
		"logger",
		"config",
		"broker",
		"registry",
		"server",
	} {
		instance, err := factory.Build(name, a.Runner)
		if err != nil {
			return err
		}

		// NO Instance
		errNOI := failure.New(nil).
			Set("name", name).
			Msg("this component has no instance") //////////////////////////////////////////////////////////////////////

		// Interface Not Implemented
		errINI := failure.New(nil).
			Set("name", name).
			Msg("this component does not implement the associated interface") //////////////////////////////////////////

		switch name {
		case "broker": //..............................Broker...
			if instance == nil {
				break
			}

			broker, ok := instance.(component.Broker)
			if !ok {
				return errINI
			}

			a.SetBroker(broker)
		case "config": //..............................Config...
			if instance == nil {
				return errNOI
			}

			config, ok := instance.(component.Config)
			if !ok {
				return errINI
			}

			a.SetConfig(config)
		case "logger": //..............................Logger...
			if instance == nil {
				return errNOI
			}

			logger, ok := instance.(component.Logger)
			if !ok {
				return errINI
			}

			a.SetLogger(logger)

			logger.Info( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
				"===BEGIN",
				"id", a.ID(),
				"name", a.Name(),
				"version", a.Version(),
				"builtAt", a.BuiltAt().String(),
				"instance", a.SdInstance(),
				"pid", os.Getpid(),
			)

		case "registry": //..........................Registry...
			if instance == nil {
				return errNOI
			}

			registry, ok := instance.(component.Registry)
			if !ok {
				return errINI
			}

			a.SetRegistry(registry)
		case "server": //..............................Server...
			if instance == nil {
				break
			}

			server, ok := instance.(component.Server)
			if !ok {
				return errINI
			}

			a.SetServer(server)
		}
	}

	return nil
}

func (a *Application) closeComponents() {
	logger := a.Logger()

	onClose := func(component string, err error) {
		if err != nil {
			logger.Error( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
				"Error when closing this component",
				"name", component,
				"reason", err.Error(),
			)
		}
	}

	if c := a.Broker(); c != nil {
		onClose("broker", c.Close())
	}

	if c := a.Config(); c != nil {
		onClose("config", c.Close())
	}

	if c := a.Registry(); c != nil {
		onClose("registry", c.Close())
	}

	if c := a.Server(); c != nil {
		onClose("server", c.Close())
	}

	if logger != nil {
		logger.Info( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
			"===END",
			"uptime", time.Since(a.StartedAt()).Round(time.Second).String(),
		)

		logger.Close()
	}
}

// Initialize AFAIRE
func (a *Application) Initialize(cmpts ...func() error) error {
	if err := a.parseFlags(); err != nil {
		return err
	}

	if err := a.initializeComponents(cmpts...); err != nil {
		return err
	}

	if err := a.buildComponents(); err != nil {
		a.closeComponents()
		return err
	}

	return nil
}

// Close AFAIRE
func (a *Application) Close() {
	a.closeComponents()
}

/*
######################################################################################################## @(°_°)@ #######
*/
