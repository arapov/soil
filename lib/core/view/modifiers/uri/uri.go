// Package uri adds URI shortcuts to the view template.
package uri

import (
	"net/http"
	"path"

	"github.com/arapov/soil/lib/core/view"
)

// Modify sets BaseURI, CurrentURI, ParentURI, and GrandparentURI variables
// for use in templates.
func Modify(w http.ResponseWriter, r *http.Request, v *view.Info) {
	v.Vars["BaseURI"] = v.BaseURI
	v.Vars["CurrentURI"] = r.URL.Path
	v.Vars["ParentURI"] = path.Dir(r.URL.Path)
	v.Vars["GrandparentURI"] = path.Dir(path.Dir(r.URL.Path))
}
