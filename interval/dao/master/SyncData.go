package master

import (
	"spider/interval/conf"
	"spider/interval/dao/utils"
	"time"
	"context"
	"fmt"
	pb "spider/interval/serve/grpc"
)

func (ms *MasterServer) StarSyncData() {
	fmt.Println("start sync")
	go func ()  {
		for {
			ms.mu.Lock()
			// 将 sync ip 同步数据
			if (ms.syncList.Len() != 0) {
				useRecord := false
				if ms.status.syncId != 0 {
					useRecord = true
				}
				req := ms.getAllSyncData(useRecord)
				utils.Log.Info("sync send Data")
				for _, masterIp := range ms.syncList.Q {
					c, ok := ms.connClients[masterIp]
					if ok {
						resp, err := c.HandleTask(context.Background(), req)
						if err != nil {
							utils.Log.Error("grpc: sync data error ", err)
						} else if resp.Code != 10000 {
							utils.Log.Info("grpc: sync data get init data", err)
							initReq := ms.getAllSyncData(true)
							c.HandleTask(context.Background(), initReq)
						}
					}
				}
			}
			ms.status.syncIdHistory = append(ms.status.syncIdHistory, ms.status.syncId)
			ms.status.syncId = time.Now().Unix()
			ms.mu.Unlock()
			time.Sleep(conf.WAIT_SYNC_DATA * time.Second)
		}
	}()
}

/*
 * 同步所有数据
 * true 状态下仅同步本次时间段内 变化
 * false 情况下 获取所有需要的数据
 */
func (ms *MasterServer) getAllSyncData(record bool) *pb.HandleTaskReq {
	var syncLastId int64
	if ms.status.syncId == 0 {
		syncLastId = 0
	} else {
		syncLastId = ms.status.syncIdHistory[len(ms.status.syncIdHistory) - 1]
	}

	req := &pb.HandleTaskReq {
		TaskCode: conf.SYNC_DATA,
		SyncData: &pb.SyncData{
			SyncId: ms.status.syncId,
			SyncLastId: syncLastId,
			MasterSyncData: &pb.MasterSyncData {
				IpList: ms.getSyncData(record),
			},
			SpiderSyncData: make(map[string]*pb.SpiderSyncData),
		},
	}
	for k, v := range ms.SpiderDispatchs {
		req.SyncData.SpiderSyncData[k] = v.GetSyncData(record)
	}
	if !record {
		req.SyncData.SyncType = conf.SYNC_ALL
	} else {
		req.SyncData.SyncType = conf.SYNC_RECORD
	}

	return req
}
