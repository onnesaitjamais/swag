/*
#######
##        ____    _____ ____ _
##       (_-< |/|/ / _ `/ _ `/
##      /___/__,__/\_,_/\_, /
##                     /___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package util

import (
	"regexp"
	"testing"
)

func TestNewUUID(t *testing.T) {
	uuid := NewUUID()

	re := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
	if !re.MatchString(uuid) {
		t.Errorf("NewV4() - uuid not valid: got=%s", uuid)
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
