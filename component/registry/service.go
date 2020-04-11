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

import "math/rand"

// Service AFAIRE
type Service struct {
	Name       string `json:"name"`
	ID         string `json:"id"`
	Version    string `json:"version"`
	BuiltAt    int64  `json:"built_at"`
	StartedAt  int64  `json:"started_at"`
	FQDN       string `json:"fqdn"`
	Port       int    `json:"port"`
	SdInstance string `json:"sd_instance"`
	Status     string `json:"status"`
	Heartbeat  int64  `json:"heartbeat"`
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
