package common

import "glog"

func CheckError(err error) {
	if err != nil {
		glog.Errorf("[ERROR] %s", err.Error())
	}
}

type PingTimeOutErr struct {
}

func (e *PingTimeOutErr) Error() string {
	return "ping time out error"
}

type NetConnectErr struct {
}

func (e *NetConnectErr) Error() string {
	return "net connect error"
}
