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
	group workgroup.Group
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

// RunGroup AFAIRE
func (s *Service) RunGroup() error {
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

// Run AFAIRE
func (s *Service) Run() error {
	return nil
}

// Close AFAIRE
func (s *Service) Close() {}

/*
######################################################################################################## @(°_°)@ #######
*/
