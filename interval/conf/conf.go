package conf

type SendInfo struct {
	Ac		string
	Ps 		string
	Host	string
}

const WAIT_SPIDER_TIME = 30
const SPIDER_TIMEOUT = 15 
const HTTP_TRY_REQUEST_TIMES = 2
const RETRY_REGISTER_TIMES = 10

var MASTER_IP = [...]string{
	// "144.202.19.110",
	// "47.103.12.134",
	// "127.0.0.1",
}

const DB_URL = "http://wwwijetchorg/"
const SPIDER_URL = "http://www.ivypub.org/index.shtml"   // http://www.jissr.net/src/assets/pdf/2019-6-11_1.pdf  http://www.ivypub.org/ERF/download/54421.shtml

const (
	RegisterCodeSuccess = 0
	RegisterMsgSuccess = " register Success "
	CreateSpiderMsgSuccess = " create Spider Success "
	DeleteSpiderMsgSuccess = " delete Spider Success "
	RegisterCodeError = -1
	RegisterMsgErrorRepeat = " this ip Repeat registered "
	CreateSpiderURLError = " create spider task error "
	CreateSpiderURLRepeat = " create spider url Repeat "
	DeleteSpiderURLError = " delete spider task error no such url"
	GetSpiderInfoError = " can not find this url "
)

const (
	Retry_Spider_Times = 20
	Retry_Send_Email_Times = 10
	WAIT_SEND_EMAIL_TIME = 60 * 60 * 3 // 60 * 50
	WAIT_SYNC_DATA = 60 * 3 
)

// slave 任务
const (
	SEND_EMAIL = 1000
	SPIDER_EMAIL = 1001
)

const TASK_BOUNDARY = 2000

// master 任务
const (
	SYNC_DATA = 2001
)

const (
	SYNC_ALL = 10
	SYNC_RECORD = 11
)

const (
	TYPE_MASTER = "TYPE_MASTER"
	TYPE_SLAVE = "TYPE_SLAVE"
)

const (
	SUCCESS_TASK		= 10000
	ERROR_EMAIL_TASK	= 10001
	ERROR_SPIDER_TASK	= 10002
	ERROR_SYNCDATA_TASK	= 10003
	ERROR_UNAHDNLE_TASK = 10004
	ERROR_MASTER_TASK	= 10005
	ERROR_SLAVE_TASK	= 10006
)