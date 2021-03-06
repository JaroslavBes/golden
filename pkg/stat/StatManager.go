package stat

import (
	"database/sql"
	"fmt"
	"github.com/vit1251/golden/pkg/registry"
	"github.com/vit1251/golden/pkg/storage"
	"time"
)

type StatManager struct {
	registry       *registry.Container
}

type Stat struct {
	TicReceived      int
	TicSent          int
	EchomailReceived int
	EchomailSent     int
	NetmailReceived  int
	NetmailSent      int

	PacketReceived   int
	PacketSent       int

	MessageReceived  int
	MessageSent      int

	SessionIn        int
	SessionOut       int
}

func NewStatManager(r *registry.Container) *StatManager {
	statm := new(StatManager)
	statm.registry = r
	statm.createStat()
	return statm
}

func (self *StatManager) RegisterInFile(filename string) error {

	storageManager := self.restoreStorageManager()

	self.createStat()

	query := "UPDATE `stat` SET `statFileRXcount` = `statFileRXcount` + 1 WHERE `statDate` = ?"
	statDate := self.makeToday()
	var params []interface{}
	params = append(params, statDate)
	err1 := storageManager.Exec(query, params, func(result sql.Result, err error) error {
		return err
	})

	return err1
}

func (self *StatManager) RegisterOutFile(filename string) error {
	self.createStat()
	return nil
}

type SummaryRow struct {
	Date string
	Value int
}

func (self *StatManager) GetStatRow(statDate string) (*Stat, error) {

	storageManager := self.restoreStorageManager()

	var result *Stat

	query1 := "SELECT `statMessageRXcount`, `statMessageTXcount`, `statSessionIn`, `statSessionOut`, `statFileRXcount`, `statFileTXcount`, `statPacketIn`, `statPacketOut` FROM `stat` WHERE `statDate` = $1"
	var params []interface{}
	params = append(params, statDate)

	err1 := storageManager.Query(query1, params, func(rows *sql.Rows) error {

		var statMessageInCount int64
		var statMessageOutCount int64
		var statSessionInCount int64
		var statSessionOutCount int64
		var statFileInCount int64
		var statFileOutCount int64
		var statPacketInCount int64
		var statPacketOutCount int64

		err2 := rows.Scan(&statMessageInCount, &statMessageOutCount, &statSessionInCount, &statSessionOutCount, &statFileInCount, &statFileOutCount, &statPacketInCount, &statPacketOutCount)
		if err2 != nil {
			return err2
		}

		result = new(Stat)
		result.MessageReceived = int(statMessageInCount)
		result.MessageSent = int(statMessageOutCount)
		result.SessionIn = int(statSessionInCount)
		result.SessionOut = int(statSessionOutCount)
		result.TicReceived = int(statFileInCount)
		result.TicSent = int(statFileOutCount)
		result.PacketReceived = int(statPacketInCount)
		result.PacketSent = int(statPacketOutCount)

		return nil
	})

	return result, err1
}

func (self *StatManager) GetStat() (*Stat, error) {
	statDate := self.makeToday()
	stat, err := self.GetStatRow(statDate)
	return stat, err
}

func (self *StatManager) RegisterInPacket() error {

	storageManager := self.restoreStorageManager()

	self.createStat()

	query := "UPDATE `stat` SET `statPacketIn` = `statPacketIn` + 1 WHERE `statDate` = ?"
	statDate := self.makeToday()
	var params []interface{}
	params = append(params, statDate)
	err1 := storageManager.Exec(query, params, func(result sql.Result, err error) error {
		return err
	})

	return err1
}

func (self *StatManager) RegisterOutPacket() error {

	storageManager := self.restoreStorageManager()

	self.createStat()
	//
	query := "UPDATE `stat` SET `statPacketOut` = `statPacketOut` + 1 WHERE `statDate` = ?"
	statDate := self.makeToday()
	var params []interface{}
	params = append(params, statDate)
	err1 := storageManager.Exec(query, params, func(result sql.Result, err error) error {
		return err
	})

	return err1
}

func (self *StatManager) RegisterDupe() error {
	self.createStat()
	return nil
}

func (self *StatManager) RegisterInMessage() error {

	storageManager := self.restoreStorageManager()

	self.createStat()

	query := "UPDATE `stat` SET `statMessageRXcount` = `statMessageRXcount` + 1 WHERE `statDate` = ?"
	statDate := self.makeToday()
	var params []interface{}
	params = append(params, statDate)
	err1 := storageManager.Exec(query, params, func(result sql.Result, err error) error {
		return err
	})

	return err1
}

func (self *StatManager) RegisterOutMessage() error {

	storageManager := self.restoreStorageManager()

	self.createStat()

	query := "UPDATE `stat` SET `statMessageTXcount` = `statMessageTXcount` + 1 WHERE `statDate` = ?"
	statDate := self.makeToday()
	var params []interface{}
	params = append(params, statDate)
	err1 := storageManager.Exec(query, params, func(result sql.Result, err error) error {
		return err
	})

	return err1
}

func (self *StatManager) makeToday() string {
	currentTime := time.Now()
	//result := currentTime.Format("2006-01-02")
	result := fmt.Sprintf("%04d-%02d-%02d", currentTime.Year(), currentTime.Month(), currentTime.Day())
	return result
}

func (self *StatManager) createStat() error {
	statDate := self.makeToday()
	stat, err1 := self.GetStatRow(statDate)
	if err1 != nil {
		return err1
	}
	if stat == nil {
		self.createStat2(statDate)
	}
	return nil
}

func (self *StatManager) createStat2(statDate string) error {

	storageManager := self.restoreStorageManager()

	query := "INSERT INTO `stat` (`statDate`) VALUES ( ? )"

	var params []interface{}
	params = append(params, statDate)

	err1 := storageManager.Exec(query, params, func(result sql.Result, err error) error {
		return err
	})

	return err1
}

func (self *StatManager) RegisterInSession() error {

	storageManager := self.restoreStorageManager()

	/* Initialize statistic record */
	self.createStat()

	/* Update statistic */
	query := "UPDATE `stat` SET `statSessionIn` = `statSessionIn` + 1 WHERE `statDate` = ?"
	statDate := self.makeToday()

	var params []interface{}
	params = append(params, statDate)

	err1 := storageManager.Exec(query, params, func(result sql.Result, err error) error {
		return err
	})

	return err1
}

func (self *StatManager) RegisterOutSession() error {

	storageManager := self.restoreStorageManager()

	/* Initialize statistic record */
	self.createStat()

	/* Update statistic */
	query := "UPDATE `stat` SET `statSessionOut` = `statSessionOut` + 1 WHERE `statDate` = ?"

	statDate := self.makeToday()

	var params []interface{}
	params = append(params, statDate)

	err1 := storageManager.Exec(query, params, func(result sql.Result, err error) error {
		return err
	})

	return err1
}

func (self *StatManager) restoreStorageManager() *storage.StorageManager {

	storageManagerPtr := self.registry.Get("StorageManager")
	if storageManager, ok := storageManagerPtr.(*storage.StorageManager); ok {
		return storageManager
	} else {
		panic("no storage manager")
	}

}
