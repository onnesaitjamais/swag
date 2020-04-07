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

// SvcList AFAIRE
type SvcList []*Service

// Len AFAIRE
func (s SvcList) Len() int {
	return len(s)
}

// Filter AFAIRE
func (s SvcList) Filter(fn func(*Service) bool) SvcList {
	var sl SvcList

	for _, svc := range s {
		if fn(svc) {
			sl = append(sl, svc)
		}
	}

	return sl
}

// Shuffle AFAIRE
func (s SvcList) Shuffle() {
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
