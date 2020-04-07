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

	"github.com/arnumina/swag/application/config"
	"github.com/arnumina/swag/runner"
	"github.com/arnumina/swag/util"
	"github.com/arnumina/swag/version"
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
		"swag.%s v%s built at %s [swag v%s] Archivage Numérique © INA %d\n\n",
		a.Name(),
		a.Version(),
		a.BuiltAt().String(),
		version.SWAG(),
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

// Initialize AFAIRE
func (a *Application) Initialize(cmpts ...func() error) error {
	if err := a.parseFlags(); err != nil {
		return err
	}

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
