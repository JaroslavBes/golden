package action

import (
	"fmt"
	"github.com/vit1251/golden/pkg/site/widgets"
	"net/http"
)

type StatAction struct {
	Action
}

func NewStatAction() *StatAction {
	sa := new(StatAction)
	return sa
}

func (self *StatAction) createMetric(tw *widgets.TableWidget, name string, rx string, tx  string) {
	tw.AddRow(widgets.NewTableRowWidget().
		AddCell(widgets.NewTableCellWidget().SetClass("table-cell").SetWidget(widgets.NewTextWidgetWithText(name))).
		AddCell(widgets.NewTableCellWidget().SetClass("table-cell").SetWidget(widgets.NewTextWidgetWithText(rx))).
		AddCell(widgets.NewTableCellWidget().SetClass("table-cell").SetWidget(widgets.NewTextWidgetWithText(tx))))
}

func (self *StatAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	statManager := self.restoreStatManager()

	/* Get statistics */
	stat, err1 := statManager.GetStat()
	if err1 != nil {
		response := fmt.Sprintf("Fail GetStat on StatManager: err = %+v", err1)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Create statistics */

	/* Render */
	bw := widgets.NewBaseWidget()

	vBox := widgets.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := self.makeMenu()
	vBox.Add(mmw)

	container := widgets.NewDivWidget()
	container.SetClass("container")

	containerVBox := widgets.NewVBoxWidget()

	container.SetWidget(containerVBox)

	vBox.Add(container)

	statWidget := widgets.NewTableWidget().
		SetClass("stat-index-items")

	statWidget.AddRow(widgets.NewTableRowWidget().
		SetClass("stat-index-header").
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Metric"))).
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Received"))).
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Sent"))))

	self.createMetric(statWidget,
			"Total TIC",
			fmt.Sprintf("%d", stat.TicReceived),
			fmt.Sprintf("%d", stat.TicSent))

	//self.createMetric(statWidget,
	//		"Total Echomail",
	//		fmt.Sprintf("%d", stat.EchomailReceived),
	//		fmt.Sprintf("%d", stat.EchomailSent))
	//
	//self.createMetric(statWidget,
	//	"Total Netmail",
	//	fmt.Sprintf("%d", stat.NetmailReceived),
	//	fmt.Sprintf("%d", stat.NetmailSent))

	self.createMetric(statWidget,
			"Total Packet",
			fmt.Sprintf("%d", stat.PacketReceived),
			fmt.Sprintf("%d", stat.PacketSent))

	self.createMetric(statWidget,
			"Total Message",
		fmt.Sprintf("%d", stat.MessageReceived),
			fmt.Sprintf("%d", stat.MessageSent))

	self.createMetric(statWidget,
			"Total session count",
			fmt.Sprintf("%d", stat.SessionIn),
			fmt.Sprintf("%d", stat.SessionOut))

	containerVBox.Add(statWidget)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}
