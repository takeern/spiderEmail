package master

import (
	"spider/interval/conf"
	"spider/interval/dao/utils"
	"time"
	"context"
	"strconv"
	"fmt"
	pb "spider/interval/serve/grpc"
	d "spider/interval/dao/dispatch"
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

	msIp, msUrls := ms.getSyncData(record)
	req := &pb.HandleTaskReq {
		TaskCode: conf.SYNC_DATA,
		SyncData: &pb.SyncData{
			SyncId: ms.status.syncId,
			SyncLastId: syncLastId,
			MasterSyncData: &pb.MasterSyncData {
				IpList: msIp,
				SpiderUrls: msUrls,
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

// 睡眠状态下 主节点同步数据
func (ms *MasterServer) HandleReq(req *pb.HandleTaskReq) *pb.HandleTaskResp{
	resp := &pb.HandleTaskResp {
		Code: conf.SUCCESS_TASK,
	}
	switch req.TaskCode {
	case conf.SYNC_DATA:
		if req.SyncData.SyncType == conf.SYNC_ALL {
			// 同步所有数据, 清理所有数据
			ms.InjectInitData(req.SyncData.MasterSyncData)
			ms.status.syncIdHistory = ms.status.syncIdHistory[0:0]
			ms.status.syncIdHistory = append(ms.status.syncIdHistory, req.SyncData.SyncId)
			for k, v := range req.SyncData.SpiderSyncData {
				if d, ok := ms.SpiderDispatchs[k]; ok {
					d.InjectInitData(v.SpiderAllData)
				}
			}
		} else if req.SyncData.SyncType == conf.SYNC_RECORD {
			// 同步 record 数据
			i := len(ms.status.syncIdHistory) - 1
			if i >= 0 && ms.status.syncIdHistory[i] == req.SyncData.SyncLastId {	// 保证 顺序一致性
				utils.Log.Info("sync record data")
				ms.InjectRecordData(req.SyncData.MasterSyncData)
				ms.status.syncIdHistory = append(ms.status.syncIdHistory, req.SyncData.SyncId)
				for k, v := range req.SyncData.SpiderSyncData {
					if d, ok := ms.SpiderDispatchs[k]; ok {
						d.InjectRecordData(v.SpiderRecordData)
					}
				}
			} else {
				resp.Code = conf.ERROR_SYNCDATA_TASK
			}
		}
		break;
	default:
		break;
	}
	return resp
}

func (ms *MasterServer) InjectInitData(data *pb.MasterSyncData) {
	utils.Log.Info("初始化主节点数据", data)
	ms.SpiderDispatchs = make(map[string]*d.Spider)
	ms.InjectRecordData(data)
}

func (ms *MasterServer) InjectRecordData(data *pb.MasterSyncData) {
	utils.Log.Info("同步主节点数据", data)
	fmt.Println(data)
	for _, url := range data.SpiderUrls {
		fmt.Println(url)
		ms.creatSpider(url)
	}
	for ip, v := range data.IpList {
		if v == conf.IP_SPIDER {
			ms.register(ip, strconv.Itoa(conf.SPIDER_EMAIL))
		}
	}
}

