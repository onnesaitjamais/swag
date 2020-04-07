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
	Register(id, name string, service *registry.Service) error
	// Deregister AFAIRE
	Deregister(id, name string) error
	// Get AFAIRE
	Get(name string) (registry.SvcList, error)
	// List AFAIRE
	List() (registry.SvcList, error)
	// Close AFAIRE
	Close() error
}

/*
######################################################################################################## @(°_°)@ #######
*/
