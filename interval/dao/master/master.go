package master

type MasterServe struct{
	Email	*Dispatch
}

func NewMasterServe(status bool) *MasterServe {
	url := "http://wwwijetchorg/"
	s := &MasterServe{
		Email: CreateEmailDispatch(status, url),
	}
	return s
}

func (m *MasterServe) HandleNewIpRegistry(ip string) {
	m.Email.HandleNewIpRegistry(ip)
}

