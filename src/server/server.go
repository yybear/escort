package server

import (
	. "common"
	"glog"
	"net"
	"os"
)

func Work() {
	go startServer()
	doCheck()
}

func doCheck() {
	ProbeCS()

	ProbeDB()

}

func startServer() {
	listener, err := net.Listen("tcp", "localhost:8260")
	if err != nil {
		glog.Errorf("[ERROR] %s", err.Error())
		os.Exit(1)
	}

	glog.Infoln("Escort server is up!")

	for {
		conn, err := listener.Accept()
		if err != nil {
			glog.Errorf("[ERROR] %s", err.Error())
			continue
		}

		go doService(conn)
	}
}

func doService(conn net.Conn) {
	glog.Infoln("Escort server is connected!")
	defer func() {
		conn.Close()
	}()

	for {
		buf, err := ReadFromConn(conn)
		if err != nil {
			glog.Errorf("[ERROR] %s", err.Error())
			break
		}

		p := new(Packet)
		p.Decode(buf)

		if p.Flags == FLAG_REQUEST && p.Length == 0 { // ping命令
			glog.Infoln("Escort server get client ping cmd!")

			pong := new(Packet)
			pong.Version = V1
			pong.Flags = FLAG_RESPONSE
			pong.Sequence = 0
			pong.Length = 0

			var data []byte
			var err error
			data, err = pong.Encode()
			if err != nil {
				glog.Errorf("[ERROR] %s", err.Error())
				continue
			}
			_, err = conn.Write(data)
			if err == nil {
				glog.Infoln("Escort server send pong successfully!")
			} else {
				glog.Errorf("[ERROR] %s", err.Error())
			}
		}
	}

}
