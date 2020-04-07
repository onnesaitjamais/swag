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
	"crypto/rand"
	"fmt"
	mathrand "math/rand"
	"os"
	"strings"

	"github.com/Showmax/go-fqdn"

	"github.com/arnumina/swag/util/failure"
)

// Alert AFAIRE
func Alert(err error) {
	fmt.Fprintf(os.Stderr, "[>>> swag <<<] HUMAN INTERVENTION REQUIRED: %s\n", err)
}

// EnvValue AFAIRE
func EnvValue(runner, suffix string, d interface{}) interface{} {
	env := []string{
		strings.Join([]string{"SWAG", strings.ToUpper(runner), suffix}, "_"),
		strings.Join([]string{"SWAG", suffix}, "_"),
	}

	for _, name := range env {
		if value, ok := os.LookupEnv(name); ok {
			return value
		}
	}

	return d
}

// FQDN AFAIRE
func FQDN() (string, error) {
	fqdn := fqdn.Get()
	if fqdn == "unknown" {
		return "",
			failure.New(nil).
				Msg("impossible to retrieve the FQDN") /////////////////////////////////////////////////////////////////
	}

	return fqdn, nil
}

// NewUUID génère un UUID V4.
func NewUUID() string {
	b := make([]byte, 16)

	_, err := rand.Read(b)
	if err != nil {
		Alert(err)
		mathrand.Read(b) //nolint:gosec
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//
	// https://tools.ietf.org/html/rfc4122
	//
	// The version 4 UUID is meant for generating UUIDs from truly-random or
	// pseudo-random numbers.
	//
	// typedef struct {
	//    unsigned32  time_low;
	//    unsigned16  time_mid;
	//    unsigned16  time_hi_and_version;
	//    unsigned8   clock_seq_hi_and_reserved;
	//    unsigned8   clock_seq_low;
	//    byte        node[6];
	// } uuid_t;
	//
	// The algorithm is as follows:
	//
	// - Set the two most significant bits (bits 6 and 7) of the clock_seq_hi_and_reserved
	//   to zero and one, respectively.
	//

	b[8] = (b[8] & 0x7F) | 0x40

	//
	// - Set the four most significant bits (bits 12 through 15) of the time_hi_and_version
	//   field to the 4-bit version number from Section 4.1.3.
	//

	b[6] = (b[6] & 0x0F) | 0x40

	//
	// Set all the other bits to randomly (or pseudo-randomly) chosen
	// values.
	//
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

/*
######################################################################################################## @(°_°)@ #######
*/
