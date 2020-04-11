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
	"os"
	"os/signal"
	"syscall"

	"github.com/arnumina/swag/application"
	"github.com/arnumina/swag/util/systemd"
	"github.com/heptio/workgroup"
)

// Service AFAIRE
type Service struct {
	*application.Application
	group      workgroup.Group
	registered bool
}

// New AFAIRE
func New(application *application.Application) *Service {
	return &Service{
		Application: application,
	}
}

// AddGroupFn AFAIRE
func (s *Service) AddGroupFn(fn func(<-chan struct{}) error) {
	s.group.Add(fn)
}

func (s *Service) runGroup() error {
	logger := s.Logger()

	s.AddGroupFn(
		func(stop <-chan struct{}) error {
			sigEnd := make(chan os.Signal, 1)
			defer close(sigEnd)
			signal.Notify(sigEnd, syscall.SIGABRT, syscall.SIGTERM)

			systemd.NotifyReady()

			logger.Info("READY") //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

			select {
			case <-stop:
				signal.Stop(sigEnd)
			case <-sigEnd:
			}

			systemd.NotifyStopping()

			logger.Info("Stopping...") //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

			return nil
		},
	)

	return s.group.Run()
}

func (s *Service) run() error {
	if err := s.preregister(); err != nil {
		return err
	}

	defer s.deregister()

	if err := s.maybeSetupServer(); err != nil {
		return err
	}

	s.register()

	if err := s.runGroup(); err != nil {
		return err
	}

	return nil
}

// Run AFAIRE
func (s *Service) Run() error {
	if err := s.run(); err != nil {
		s.Logger().Critical( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
			"Runtime error",
			"reason", err.Error(),
		)

		return err
	}

	return nil
}

// Close AFAIRE
func (s *Service) Close() {
	s.Application.Close()
}

/*
######################################################################################################## @(°_°)@ #######
*/
