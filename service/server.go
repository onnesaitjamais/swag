/*
#######
##        ____    _____ ____ _
##       (_-< |/|/ / _ `/ _ `/
##      /___/__,__/\_,_/\_, /
##                     /___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package service

import (
	"time"

	"github.com/arnumina/swag/util/systemd"
)

func (s *Service) maybeSetupServer() error {
	if s.Server() == nil {
		return nil
	}

	s.AddGroupFn(func(stop <-chan struct{}) error {
		err := s.Server().Start(s.Config().Port())

		s.Logger().Info("Server stopped") //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

		return err
	})

	delay, err := systemd.WatchdogDelay()
	if err != nil {
		return err
	}

	s.AddGroupFn(func(stop <-chan struct{}) error {
		server := s.Server()

		if delay == 0 {
			<-stop
		} else {
			fqdn := s.FQDN()
			port := s.Config().Port()

		loop:
			for {
				select {
				case <-stop:
					break loop
				case <-time.After(delay):
					if err := server.CheckHealth(fqdn, port); err != nil {
						s.Logger().Warning( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
							"The server seems to have a problem",
							"reason", err.Error(),
						)
					} else {
						systemd.NotifyWatchdog()
					}
				}
			}
		}

		return server.Stop()
	})

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
