package slave

import (
	"io/ioutil"
	"log"
	"net/http"
	"spider/interval/conf"
	"spider/interval/dao/utils"

	pb "spider/interval/serve/grpc"
)


func RegisterIp(times int) {
	if (times > conf.RETRY_REGISTER_TIMES) {
		log.Fatal("register many time exit")
	}
	resp, err := http.Get(conf.MASTER_HOST + "/register")
	if err != nil {
		times ++
		RegisterIp(times)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		times ++
		RegisterIp(times)
		return
	}
	if (body[0] == 0) {
		log.Println("register success")
	} else {
		log.Fatal("register res unhandle %v", body)
	}
}

/*
 * 选择任务类型
 * 1000 发送 email
 * 1001 爬取页面 email
 */
func HandleReq(req *pb.HandleTaskReq) (*pb.HandleTaskResp) {
	resp := &pb.HandleTaskResp{
		SpiderInfo: &pb.SpiderInfo{},
		Code: conf.SUCCESS_TASK,
	}
	switch req.TaskCode {
	case conf.SEND_EMAIL:
		err := SendMail(req.EmailInfo.Ac,
			req.EmailInfo.Ps,
			req.EmailInfo.Host,
			req.EmailInfo.Receive,
			"无主题",
			conf.EmailModalList[req.EmailInfo.ModalIndex],
			"html")
		if err != nil {
			resp.Code = conf.ERROR_EMAIL_TASK
			resp.ErrorMsg = err.Error()
		} else {
			resp.Code = conf.SUCCESS_TASK
		}

		break
	case conf.SPIDER_EMAIL:
		log.Println(resp)
		err, emails, urls := SpiderEmail(req.SpiderUrl, 0)
		if err != nil {
			resp.Code = conf.ERROR_SPIDER_TASK
			resp.ErrorMsg = err.Error()
		} else {
			resp.Code = conf.SUCCESS_TASK
			resp.SpiderInfo.Emails = emails
			resp.SpiderInfo.Urls = urls
		}
		break
	case conf.SYNC_DATA:
		utils.Log.Info("sync get data", req)
		switch req.SyncData.SyncType {
			case conf.SYNC_ALL:
				break;
			case conf.SYNC_RECORD:
				log.Println("req.SyncData.SpiderSyncData 长度", len(req.SyncData.SpiderSyncData.SpiderRecordData))
				for _, v := range req.SyncData.SpiderSyncData.SpiderRecordData {
					log.Println("v: ", v)
				}
				break;
			default:
				break;
		}
		break;
	default:
		resp.Code = conf.ERROR_UNAHDNLE_TASK
		resp.ErrorMsg = "unhandle task code"
		break
	}
	return resp
}