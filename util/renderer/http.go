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

import "net/http"

// Renderer AFAIRE
type Renderer struct {
	http.ResponseWriter
}

// New AFAIRE
func New(w http.ResponseWriter) *Renderer {
	return &Renderer{
		ResponseWriter: w,
	}
}

// NoContent AFAIRE
func (r *Renderer) NoContent() {
	r.WriteHeader(http.StatusNoContent)
}

// ErrorStr AFAIRE
func (r *Renderer) ErrorStr(status int, error string) {
	http.Error(r.ResponseWriter, error, status)
}

// Error AFAIRE
func (r *Renderer) Error(status int, err error) {
	r.ErrorStr(status, err.Error())
}

// ErrorStr500 AFAIRE
func (r *Renderer) ErrorStr500(error string) {
	r.ErrorStr(http.StatusInternalServerError, error)
}

// Error500 AFAIRE
func (r *Renderer) Error500(err error) {
	r.ErrorStr500(err.Error())
}

/*
######################################################################################################## @(°_°)@ #######
*/
