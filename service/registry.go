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
	"math/rand"
	"time"

	_registry "github.com/arnumina/swag/component/registry"
	"github.com/arnumina/swag/util"
	"github.com/arnumina/swag/util/failure"
)

func (s *Service) newRegistryService(port int, status string) *_registry.Service {
	return &_registry.Service{
		Name:       s.Name(),
		ID:         s.ID(),
		Version:    s.Version(),
		BuiltAt:    s.BuiltAt().Unix(),
		StartedAt:  s.StartedAt().Unix(),
		FQDN:       s.FQDN(),
		Port:       port,
		SdInstance: s.SdInstance(),
		Status:     status,
		Heartbeat:  time.Now().Unix(),
	}
}

func (s *Service) updateRegistryService(service *_registry.Service) {
	service.Heartbeat = time.Now().Unix()
}

func (s *Service) preregister() error {
	config := s.Config()

	if config.Port() != 0 {
		s.Logger().Info("Config", "port", s.Config().Port()) //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::
		return nil
	}

	assign := func(used []int) (*_registry.Service, error) {
		min := config.PortMin()
		max := config.PortMax()

		for i := min; i <= max; i++ {
			port := rand.Intn(max-min+1) + min

			for _, p := range used {
				if p == port {
					port = 0
					break
				}
			}

			if port != 0 {
				config.SetPort(port)
				s.Logger().Info("Assignment", "port", port) //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

				return s.newRegistryService(port, "starting"), nil
			}
		}

		return nil, failure.New(nil).
			Set("min", min).
			Set("max", max).
			Msg("impossible to retrieve a free TCP port") //////////////////////////////////////////////////////////////
	}

	if err := s.Registry().Preregister(s.ID(), s.Name(), assign); err != nil {
		return err
	}

	s.registered = true

	return nil
}

func (s *Service) register() {
	s.AddGroupFn(func(stop <-chan struct{}) error {
		registry := s.Registry()
		service := s.newRegistryService(s.Config().Port(), "running")

		if err := registry.Register(service); err != nil {
			return err
		}

		s.registered = true

		if registry.Interval() == 0 {
			<-stop
		} else {
			interval := time.Duration(registry.Interval()) * time.Second

			for {
				select {
				case <-stop:
					s.deregister()
					return nil
				case <-time.After(interval):
					s.updateRegistryService(service)
					if err := registry.Register(service); err != nil {
						s.Logger().Warning( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
							"Impossible to register the service",
							"id", s.ID(),
							"name", s.Name(),
							"reason", err.Error(),
						)
					}
				}
			}
		}

		return nil
	})
}

func (s *Service) deregister() {
	if s.registered {
		if err := s.Registry().Deregister(s.Runner.ID(), s.Runner.Name()); err != nil {
			f := failure.New(err).
				Set("id", s.ID()).
				Set("name", s.Name()).
				Msg("Cannot deregister this service") //////////////////////////////////////////////////////////////////

			s.Logger().Error(f.Error()) //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

			util.Alert(f)
		}

		s.registered = false
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
