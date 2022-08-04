package test

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/withlin/canal-go/client"
	protocol "github.com/withlin/canal-go/protocol/entry"
	"testing"
	"time"
)

func InitCanal() {
	connector := client.NewSimpleCanalConnector("127.0.0.1", 11111,
		"admin", "123456",
		"example", 60000, 60*60*1000)
	err := connector.Connect()
	if err != nil {
		panic(err)
	}
	// 订阅mysql的表
	err = connector.Subscribe(".*\\..*")
	if err != nil {
		panic(err)
	}

	for {
		message, err := connector.Get(100, nil, nil)
		if err != nil {
			panic(err)
		}
		messageId := message.Id
		if messageId == -1 || len(message.Entries) <= 0 {
			time.Sleep(300 * time.Millisecond)
			continue
		}

		// 处理消息
		doMessage(message.Entries)
	}
}

func doMessage(entries []protocol.Entry) {
	for _, entry := range entries {
		if entry.GetEntryType() == protocol.EntryType_TRANSACTIONBEGIN || entry.GetEntryType() == protocol.EntryType_TRANSACTIONEND {
			continue
		}
		rowChange := new(protocol.RowChange)
		err := proto.Unmarshal(entry.GetStoreValue(), rowChange)
		if err != nil {
			panic(err)
		}
		if rowChange != nil {
			eventType := rowChange.GetEventType()
			header := entry.Header

			if header.ExecuteTime < time.Now().UnixNano()/1e6 {
				// 变动时间早于当前时间，直接抛弃当前数据
				continue
			}
			for _, rowData := range rowChange.RowDatas {
				if eventType == protocol.EventType_DELETE {
					//todo 收到数据库删除数据推送，执行对应处理，rowData.GetBeforeColumns() 可以获取到删除的数据
					//rowData.GetBeforeColumns()
					fmt.Println("删数据")
				} else if eventType == protocol.EventType_INSERT {
					//todo 收到数据库新增数据推送，执行对应处理，rowData.GetAfterColumns() 可以获取到新增的数据
					columns := rowData.GetAfterColumns()
					fmt.Printf("%#v\n", columns)

				} else {
					//todo 收到数据库更新数据推送，执行对应处理，rowData.GetAfterColumns() 可以获取到更新后的数据，rowData.GetBeforeColumns() 可以获取到更新前的数据
					//rowData.GetBeforeColumns()
					//rowData.GetAfterColumns()
					fmt.Println("更新数据")
				}
			}
		}
	}
}

func TestCanal(t *testing.T) {
	InitCanal()
}
