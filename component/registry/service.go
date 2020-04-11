/*
#######
##        ____    _____ ____ _
##       (_-< |/|/ / _ `/ _ `/
##      /___/__,__/\_,_/\_, /
##                     /___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package registry

import (
	"math/rand"
)

// Service AFAIRE
type Service struct {
	Name       string
	ID         string
	Version    string
	BuiltAt    int64
	StartedAt  int64
	FQDN       string
	Port       int
	SdInstance string
	Status     string
	Heartbeat  int64
}

// Services AFAIRE
type Services []*Service

// Len AFAIRE
func (s Services) Len() int {
	return len(s)
}

// Filter AFAIRE
func (s Services) Filter(fn func(*Service) bool) Services {
	var result Services

	for _, service := range s {
		if fn(service) {
			result = append(result, service)
		}
	}

	return result
}

// Shuffle AFAIRE
func (s Services) Shuffle() {
	rand.Shuffle(
		s.Len(),
		func(i, j int) {
			s[i], s[j] = s[j], s[i]
		},
	)
}

/*
######################################################################################################## @(°_°)@ #######
*/
