package ui

import (
	"fmt"
	"github.com/vit1251/golden/pkg/common"
	"log"
	"net/http"
)

type NetmailComposeCompleteAction struct {
	Action
}

func NewNetmailComposeCompleteAction() (*NetmailComposeCompleteAction) {
	nm := new(NetmailComposeCompleteAction)
	return nm
}

func (self *NetmailComposeCompleteAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	master := common.GetMaster()

	/* Parse */
	err1 := r.ParseForm()
	if err1 != nil {
		panic(err1)
	}

	/* Create netmail message */
	to := r.PostForm.Get("to")
	subj := r.PostForm.Get("subject")
	body := r.PostForm.Get("body")
	log.Printf("Compose netmail: to = %s subj = %s body = %s", to, subj, body)

	//
	nm := master.TosserManager.NewNetmailMessage()
	nm.Subject = subj
	nm.Body = body
	nm.To = to

	/* Delivery message */
	err2 := master.TosserManager.WriteNetmailMessage(nm)
	if err2 != nil {
		panic(err2)
	}

	/* Redirect */
	newLocation := fmt.Sprintf("/netmail")
	http.Redirect(w, r, newLocation, 303)
}