/*
#######
##        ____    _____ ____ _
##       (_-< |/|/ / _ `/ _ `/
##      /___/__,__/\_,_/\_, /
##                     /___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package component

import "github.com/arnumina/swag/component/registry"

// Registry AFAIRE
type Registry interface {
	// Interval AFAIRE
	Interval() int
	// Preregister AFAIRE
	Preregister(id, name string, fn func([]int) (*registry.Service, error)) error
	// Register AFAIRE
	Register(service *registry.Service) error
	// Deregister AFAIRE
	Deregister(id, name string) error
	// Get AFAIRE
	Find(name string) (registry.Services, error)
	// List AFAIRE
	List() (registry.Services, error)
	// Close AFAIRE
	Close() error
}

/*
######################################################################################################## @(°_°)@ #######
*/
