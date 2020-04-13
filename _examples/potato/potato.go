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
	"time"

	"github.com/arnumina/swag"
	"github.com/arnumina/swag/component/broker"
	"github.com/arnumina/swag/service"
	"github.com/arnumina/swag/util/options"
)

var version, builtAt string

func initialize(s *service.Service) {
	s.Broker().Subscribe(
		s.Name(),
		func(m *broker.Message) bool {
			s.Logger().Trace("Message 1", "event", m.Event)
			return true
		},
	)

	s.Broker().Subscribe(
		s.Name(),
		func(m *broker.Message) bool {
			s.Logger().Trace("Message 2", "event", m.Event)
			return true
		},
	)

	s.AddGroupFn(
		func(stop <-chan struct{}) error {
		loop:
			for {
				select {
				case <-stop:
					break loop
				case <-time.After(100 * time.Millisecond):
					s.Broker().Publish("event_Hello", nil)
				}
			}

			return nil
		},
	)
}

func run() error {
	service, err := swag.NewService(
		"potato",
		version,
		builtAt,
		swag.Broker(
			"mongodb",
			options.Options{},
		),
	)
	if err != nil {
		return err
	}

	defer service.Close()

	initialize(service)

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
