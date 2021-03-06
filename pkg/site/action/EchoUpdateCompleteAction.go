package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type EchoUpdateCompleteAction struct {
	Action
}

func NewEchoUpdateCompleteAction() *EchoUpdateCompleteAction {
	euc := new(EchoUpdateCompleteAction)
	return euc
}

func (self *EchoUpdateCompleteAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	areaManager := self.restoreAreaManager()

	/* Parse POST parameters */
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	/* ... */
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	/* ... */
	area, err1 := areaManager.GetAreaByName(echoTag)
	if err1 != nil {
		panic(err1)
	}
	log.Printf("area = %+v", area)

	/* Update summary */
	area.Summary = r.PostForm.Get("summary")

	/* Update charset */
	newCharset := r.PostForm.Get("charset")
	if newCharset == "CP866" || newCharset == "UTF-8" {
		area.Charset = newCharset
	}

	/* Update area property */
	err2 := areaManager.Update(area)
	if err2 != nil {
		panic(err2)
	}

	/* Render */
	newLocation := fmt.Sprintf("/echo/%s", echoTag)
	http.Redirect(w, r, newLocation, 303)

}
