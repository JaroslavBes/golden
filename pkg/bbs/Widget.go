package bbs

type Widget interface {
	Render(cs *ConnState)
	ProcessEvent(cs *ConnState, event *TerminalEvent)
}
