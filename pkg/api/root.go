package api

import (
	"net/http"

	"github.com/gleich/lumber/v2"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://github.com/gleich/gex_slack", http.StatusPermanentRedirect)
	lumber.Success("Handled root redirect")
}
