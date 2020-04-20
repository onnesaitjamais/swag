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
	"sort"
	"time"
)

// Service AFAIRE
type Service struct {
	ID         string
	Name       string
	Version    string
	BuiltAt    time.Time
	StartedAt  time.Time
	FQDN       string
	Port       int
	SdInstance string
	Status     string
	Heartbeat  time.Time
	Interval   time.Duration
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

type lessFunc func(si, sj *Service) bool

type multiSorter struct {
	data Services
	less []lessFunc
}

func (ms *multiSorter) Len() int {
	return len(ms.data)
}

func (ms *multiSorter) Swap(i, j int) {
	ms.data[i], ms.data[j] = ms.data[j], ms.data[i]
}

func (ms *multiSorter) Less(i, j int) bool {
	di, dj := ms.data[i], ms.data[j]

	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]

		switch {
		case less(di, dj):
			return true
		case less(dj, di):
			return false
		}
	}

	return ms.less[k](di, dj)
}

// Sort AFAIRE
func (s Services) Sort(fields ...string) {
	less := []lessFunc{}

	for _, field := range fields {
		switch field {
		case "ID":
			less = append(less, func(si, sj *Service) bool { return si.ID < sj.ID })
		case "Name":
			less = append(less, func(si, sj *Service) bool { return si.Name < sj.Name })
		case "FQDN":
			less = append(less, func(si, sj *Service) bool { return si.FQDN < sj.FQDN })
		case "Port":
			less = append(less, func(si, sj *Service) bool { return si.Port < sj.Port })
		}
	}

	ms := &multiSorter{
		data: s,
		less: less,
	}

	sort.Sort(ms)
}

/*
######################################################################################################## @(°_°)@ #######
*/
