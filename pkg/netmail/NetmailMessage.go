package netmail

import (
	commonfunc "github.com/vit1251/golden/pkg/common"
	"time"
)

type NetmailMessage struct {
	ID          string
	MsgID       string
	Hash        string
	From        string
	To          string
	OrigAddr    string
	DestAddr    string
	Subject     string
	Content     string
	UnixTime    int64
	ViewCount   int
	DateWritten time.Time
	Packet      []byte
}

func NewNetmailMessage() *NetmailMessage {
	nm := new(NetmailMessage)
	return nm
}

func (self *NetmailMessage) SetMsgID(msgID string) {
	self.MsgID = msgID
}

func (self *NetmailMessage) SetSubject(subject string) {
	self.Subject = subject
}

func (self *NetmailMessage) SetID(id string) {
	self.ID = id
}

func (self *NetmailMessage) SetFrom(from string) {
	self.From = from
}

func (self *NetmailMessage) SetTo(to string) {
	self.To = to
}

func (self *NetmailMessage) SetViewCount(count int) *NetmailMessage {
	self.ViewCount = count
	return self
}

func (self *NetmailMessage) SetContent(body string) *NetmailMessage {
	self.Content = body
	return self
}

func (self *NetmailMessage) SetHash(hash string) *NetmailMessage {
	self.Hash = hash
	return self
}

func (self *NetmailMessage) GetContent() string {
	return self.Content
}

func (self *NetmailMessage) SetUnixTime(unixTime int64) {
	self.UnixTime = unixTime
	self.DateWritten = time.Unix(unixTime, 0)
}

func (self *NetmailMessage) SetTime(ptm time.Time) {
	self.DateWritten = ptm
	self.UnixTime = ptm.Unix()
}

func (self NetmailMessage) GetAge() string {
	result := commonfunc.MakeHumanTime(self.DateWritten)
	return result
}

func (self *NetmailMessage) SetPacket(packet []byte) {
	self.Packet = packet
}

func (self *NetmailMessage) SetOrigAddr(addr string) {
	self.OrigAddr = addr
}

func (self *NetmailMessage) SetDestAddr(addr string) {
	self.DestAddr = addr
}

