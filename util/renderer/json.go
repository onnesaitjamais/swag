/*
#######
##        ____    _____ ____ _
##       (_-< |/|/ / _ `/ _ `/
##      /___/__,__/\_,_/\_, /
##                     /___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package renderer

import (
	"encoding/json"
	"net/http"
)

// JSON AFAIRE
func (r *Renderer) JSON(status int, data interface{}) error {
	r.Header().Set("Content-Type", "application/json; charset=utf-8")
	r.WriteHeader(status)

	err := json.NewEncoder(r).Encode(data)
	if err != nil {
		return err
	}

	return nil
}

// JSONOk AFAIRE
func (r *Renderer) JSONOk(data interface{}) error {
	return r.JSON(http.StatusOK, data)
}

/*
######################################################################################################## @(°_°)@ #######
*/
