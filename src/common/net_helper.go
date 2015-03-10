package common

import (
	"glog"
	"net"
)

func ReadFromConn(conn net.Conn) ([]byte, error) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		glog.Errorf("[ERROR] %s", err.Error())
		return nil, err
	}

	return buf, nil
}
