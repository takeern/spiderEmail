package master

import (
	"context"
	"spider/interval/modal"
)

type IDomainState {
	index				int
	ch 					chan pb.HandleTaskReq
	ctx					context.Context
}

type IControl struct {
	connClients				map[string]pb.TaskClient
	domainState				map[string]IDomainState
	emailIpList				*modal.Queue
	spiderIpList			*modal.Queue
}

// create new send control
func newControl() *IControl {

}