package mailer

import (
	"bytes"
	"github.com/vit1251/golden/pkg/mailer/stream"
	"log"
)

type MailerStateAdditionalStep struct {
	MailerState
}

func NewMailerStateAdditionalStep() *MailerStateAdditionalStep {
	return new(MailerStateAdditionalStep)
}

func (self *MailerStateAdditionalStep) String() string {
	return "MailerStateAdditionalStep"
}

func (self *MailerStateAdditionalStep) processCommandFrame(mailer *Mailer, nextFrame stream.Frame) IMailerState {

	var streamCommandId = nextFrame.CommandFrame.CommandID

	/* Use modern secure authorization */
	if streamCommandId == stream.M_NUL {
		if key, value, err1 := mailer.parseInfoFrame(nextFrame.CommandFrame.Body); err1 == nil {
			log.Printf("Remote side option: name = %s value = %s", key, value)
			if bytes.Equal(key, []byte("OPT")) {
				mailer.parseInfoOptFrame(value)
			}
		} else {
			log.Printf("Parse error: err = %+v", err1)
		}
	}

	/* Use unsecure password authorization */
	if streamCommandId == stream.M_ADR {
		if mailer.respAuthorization != "" {
			return NewMailerStateSecureAuthRemoteAction()
		} else {
			return NewMailerStateAuthRemote()
		}
	}

	return self
}

func (self *MailerStateAdditionalStep) processFrame(mailer *Mailer, nextFrame stream.Frame) IMailerState {

	if nextFrame.IsCommandFrame() {
		return self.processCommandFrame(mailer, nextFrame)
	}

	return self

}

func (self *MailerStateAdditionalStep) Process(mailer *Mailer) IMailerState {

	select {
	case nextFrame := <-mailer.stream.InDataFrames:
		return self.processFrame(mailer, nextFrame)
	}

}