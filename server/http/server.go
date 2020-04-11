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
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/arnumina/swag/util/failure"
)

type server struct {
	tls       bool
	local     bool
	healthURI string
	certFile  string
	keyFile   string
	server    *http.Server
	client    *http.Client
}

// Start AFAIRE
func (s *server) Start(port int) error {
	if s.local {
		s.server.Addr = fmt.Sprintf("localhost:%d", port)
	} else {
		s.server.Addr = fmt.Sprintf(":%d", port)
	}

	if s.tls {
		cert, err := tls.LoadX509KeyPair(s.certFile, s.keyFile)
		if err != nil {
			return err
		}

		caCert, err := ioutil.ReadFile(s.certFile)
		if err != nil {
			return err
		}

		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		s.client = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					RootCAs:      caCertPool,
					Certificates: []tls.Certificate{cert},
				},
			},
		}

		s.server.TLSConfig = &tls.Config{
			ClientCAs:  caCertPool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		}

		if err := s.server.ListenAndServeTLS(s.certFile, s.keyFile); err != http.ErrServerClosed {
			return err
		}
	} else if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}

	return nil
}

// CheckHealth AFAIRE
func (s *server) CheckHealth(fqdn string, port int) error {
	var (
		err error
		res *http.Response
	)

	if s.tls {
		res, err = s.client.Get(fmt.Sprintf("https://%s:%d%s", fqdn, port, s.healthURI))
		if err != nil {
			return err
		}

		res.Body.Close()
	} else {
		res, err = http.Get(fmt.Sprintf("http://localhost:%d%s", port, s.healthURI))
		if err != nil {
			return err
		}

		res.Body.Close()
	}

	if res.StatusCode != http.StatusNoContent {
		return failure.New(nil).
			Set("status", res.StatusCode).
			Msg("the status of the response is not as expected") ///////////////////////////////////////////////////////
	}

	return nil
}

// Stop AFAIRE
func (s *server) Stop() error {
	if err := s.server.Shutdown(context.Background()); err != nil {
		return err
	}

	return nil
}

// Close AFAIRE
func (s *server) Close() error {
	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
