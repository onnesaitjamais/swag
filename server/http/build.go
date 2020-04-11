/*
#######
##        ____    _____ ____ _
##       (_-< |/|/ / _ `/ _ `/
##      /___/__,__/\_,_/\_, /
##                     /___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package http

import (
	"net/http"

	"github.com/arnumina/swag/runner"
	"github.com/arnumina/swag/util/failure"
	"github.com/arnumina/swag/util/options"
	"github.com/arnumina/swag/util/value"
)

const (
	_defaultLocal     = true
	_defaultHealthURI = "/api/health"
	_defaultTLS       = false
)

func (s *server) setHandler(opts options.Options) error {
	const option = "handler"

	ov, ok := opts[option]
	if !ok {
		return failure.New(nil).
			Set("name", option).
			Msg("this option is required") /////////////////////////////////////////////////////////////////////////////
	}

	handler, ok := ov.(http.Handler)
	if !ok {
		return failure.New(nil).
			Set("name", option).
			Msg("this option is not of type 'http.Handler'") ///////////////////////////////////////////////////////////
	}

	s.server.Handler = handler

	return nil
}

func (s *server) setLocal(opts options.Options, runner string, cfg *value.Value) error {
	const option = "local"

	d, err := cfg.DBool(_defaultLocal, option)
	if err != nil {
		return err
	}

	opts.SetOption(
		option,
		"SERVER_LOCAL",
		runner,
		d,
	)

	local, err := opts.Bool(option)
	if err != nil {
		return err
	}

	s.local = local

	return nil
}

func (s *server) setHealthURI(opts options.Options, runner string, cfg *value.Value) error {
	const option = "health_URI"

	d, err := cfg.DString(_defaultHealthURI, option)
	if err != nil {
		return err
	}

	opts.SetOption(
		option,
		"SERVER_HEALTH_URI",
		runner,
		d,
	)

	healthURI, err := opts.String(option)
	if err != nil {
		return err
	}

	s.healthURI = healthURI

	return nil
}

func (s *server) setTLS(opts options.Options, runner string, cfg *value.Value) error {
	const option = "TLS"

	d, err := cfg.DBool(_defaultTLS, option)
	if err != nil {
		return err
	}

	opts.SetOption(
		option,
		"SERVER_TLS",
		runner,
		d,
	)

	tls, err := opts.Bool(option)
	if err != nil {
		return err
	}

	s.tls = tls

	return nil
}

func (s *server) setCertFile(opts options.Options, runner string, cfg *value.Value) error {
	const option = "cert_file"

	d, err := cfg.DString("", option)
	if err != nil {
		return err
	}

	opts.SetOption(
		option,
		"SERVER_CERT_FILE",
		runner,
		d,
	)

	certFile, err := opts.String(option)
	if err != nil {
		return err
	}

	if certFile == "" {
		return failure.New(nil).
			Set("name", option).
			Msg("this option is required") /////////////////////////////////////////////////////////////////////////////
	}

	s.certFile = certFile

	return nil
}

func (s *server) setKeyFile(opts options.Options, runner string, cfg *value.Value) error {
	const option = "key_file"

	d, err := cfg.DString("", option)
	if err != nil {
		return err
	}

	opts.SetOption(
		option,
		"SERVER_KEY_FILE",
		runner,
		d,
	)

	keyFile, err := opts.String(option)
	if err != nil {
		return err
	}

	if keyFile == "" {
		return failure.New(nil).
			Set("name", option).
			Msg("this option is required") /////////////////////////////////////////////////////////////////////////////
	}

	s.keyFile = keyFile

	return nil
}

// Build AFAIRE
func Build(opts options.Options, runner *runner.Runner) (interface{}, error) {
	cfg, err := runner.ComponentCfg("server")
	if err != nil {
		return nil, err
	}

	server := &server{
		server: &http.Server{},
	}

	if err := server.setHandler(opts); err != nil {
		return nil, err
	}

	if err := server.setLocal(opts, runner.Name(), cfg); err != nil {
		return nil, err
	}

	if err := server.setHealthURI(opts, runner.Name(), cfg); err != nil {
		return nil, err
	}

	if err := server.setTLS(opts, runner.Name(), cfg); err != nil {
		return nil, err
	}

	if server.tls {
		if err := server.setCertFile(opts, runner.Name(), cfg); err != nil {
			return nil, err
		}

		if err := server.setKeyFile(opts, runner.Name(), cfg); err != nil {
			return nil, err
		}
	}

	return server, nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
