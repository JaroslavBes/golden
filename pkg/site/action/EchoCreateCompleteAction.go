package action

import (
	"fmt"
	"github.com/vit1251/golden/pkg/echomail"
	"net/http"
)

type EchoCreateComplete struct {
	Action
}

func NewEchoCreateCompleteAction() *EchoCreateComplete {
	return new(EchoCreateComplete)
}

func (self *EchoCreateComplete) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	areaManager := self.restoreAreaManager()

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	//
	echoTag := r.Form.Get("echoname")
	fmt.Printf("echoTag = %v", echoTag)

	a := echomail.NewArea()
	a.SetName(echoTag)
	areaManager.Register(a)

	//
	newLocation := fmt.Sprintf("/echo/%s", a.GetName())
	http.Redirect(w, r, newLocation, 303)

}
