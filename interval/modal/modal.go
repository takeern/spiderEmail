package modal

const HTTP_TRY_REQUEST_TIMES = 3
const SPIDER_WAIT_TIME = 30
const SPIDER_TIMEOUT = 10

const (
	GETBOOKDESC	= iota
	GETBOOKLIST
	GETBOOKDATA
	GETBOOKALLDATA
)

type Server struct {}
