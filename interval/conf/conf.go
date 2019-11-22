package conf

type SendInfo struct {
	Ac		string
	Ps 		string
	Host	string
}

const WAIT_SEND_EMAIL_TIME = 60 * 60 * 3 // 60 * 50
const WAIT_SYNC_DATA = 60 
const WAIT_SPIDER_TIME = 60 * 2
const SPIDER_TIMEOUT = 15 
const HTTP_TRY_REQUEST_TIMES = 2
const RETRY_REGISTER_TIMES = 10

var MASTER_IP = [2]string{
	"144.202.19.110",
	"47.103.12.134",
}

const DB_URL = "http://wwwijetchorg/"
const SPIDER_URL = "http://dpi-proceedings.com/index.php/dtem/article/download/31137/29718"

const (
	RegisterCodeSuccess = 0
	RegisterMsgSuccess = " register Success "
	RegisterCodeError = -1
	RegisterMsgErrorRepeat = " this ip Repeat registered "
)

const (
	Retry_Spider_Times = 20
	Retry_Send_Email_Times = 10
)

const (
	SEND_EMAIL = 1000
	SPIDER_EMAIL = 1001
	SYNC_DATA = 1002
)
