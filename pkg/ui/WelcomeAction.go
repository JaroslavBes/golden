package ui

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"path/filepath"
	"html/template"
)

type WelcomeAction struct {
	Action
}

func (self *WelcomeAction) ServeHTTP(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	lp := filepath.Join("views", "layout.tmpl")
	fp := filepath.Join("views", "welcome.tmpl")
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		panic(err)
	}
	outParams := make(map[string]interface{})
	outParams["Areas"] = self.Site.app.AreaList.Areas
	tmpl.ExecuteTemplate(w, "layout", outParams)
}
