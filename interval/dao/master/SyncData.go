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
				firstSend := true
				if ms.status.syncId != 0 {
					firstSend = false
				}
				req := ms.getSyncData(firstSend)

				for _, masterIp := range ms.syncList.Q {
					c, ok := ms.connClients[masterIp]
					if ok {
						resp, err := c.HandleTask(context.Background(), req)
						if err != nil {
							utils.Log.Error("grpc: sync data error ", err)
						} else if resp.Code != 10000 {
							utils.Log.Info("grpc: sync data get init data", err)
							initReq := ms.getSyncData(true)
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

// 获取 需要同步的数据
// true 同步所有数据
// false 同步变化数据
func (ms *MasterServer) getSyncData(status bool) *pb.HandleTaskReq {
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
		},
	}

	if status {
		all := ms.SpiderDispatch.GetAllData()
		req.SyncData.SyncType = conf.SYNC_ALL
		req.SyncData.SpiderSyncData = &pb.SpiderSyncData{
			SpiderAllData: &pb.SpiderAllData{
				IpList: all.Ip_list.Q,
				CloseIpList: all.Close_ip_list.Q,
				WaitSpiderQueue: all.Wait_spider_queue.Q,
				HadSpiderQueue: all.Had_spider_queue.Q,
				ErrorSpiderQueue: all.Error_spider_queue.Q,
				CacheEmail: all.Cache_email,
				HostUrl: all.Host_url,
			},
		}
	} else {
		record := ms.SpiderDispatch.GetSyncData()
		req.SyncData.SyncType = conf.SYNC_RECORD
		req.SyncData.SpiderSyncData = &pb.SpiderSyncData{
			SpiderRecordData: record,
		}
	}

	return req
}
