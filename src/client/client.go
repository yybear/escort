package client

import (
	. "common"
	"glog"
	"net"
	"os"
	"time"
)

var (
	pingAtCh chan int = make(chan int, 1)
	netErrCh chan int = make(chan int, 1)
)

func ping(conn net.Conn) error {
	select {
	case pingAtCh <- 1:
		glog.Infoln("Escort slave send ping!")
	default:
		// 没有收到上一次ping的pong，超时
		glog.Warningln("Escort slave doesn't get pong!")
		return new(PingTimeOutErr)
	}

	pingcmd := new(Packet)
	pingcmd.Version = V1
	pingcmd.Flags = FLAG_REQUEST
	pingcmd.Length = 0
	pingcmd.Sequence = 0

	data, err := pingcmd.Encode()
	_, err = conn.Write(data)
	if err != nil {
		glog.Errorf("[ERROR] %s", err.Error())
		return err
	}

	return nil
}

func pong() {
	glog.Infoln("Escort slave get pong!")
	<-pingAtCh // 收到pong则清理pingAtCh
}

/**
 * 处理从对端获取的响应或请求
 **/
func handler(conn net.Conn) {
	glog.Infoln("Escort slave handler")
	for {
		data, err := ReadFromConn(conn)
		if err != nil {
			glog.Errorf("[ERROR] %s", err.Error())
			break
		} else {
			packet := new(Packet)
			packet.Decode(data)

			if packet.Flags == FLAG_REQUEST {
				// master 的请求
			} else {
				// master 的响应
				if packet.Length == 0 {
					// pong
					pong()
				}
			}
		}
	}

}

// 切换为master
func doSwitch() {

}

func connectServer() {
	//接通
	conn, err := net.Dial("tcp", "localhost:8260")
	defer func() {
		conn.Close()
	}()

	if err != nil {
		glog.Errorln("Escort slave can't connect to master!")
		netErrCh <- 1
		//return new(NetConnectErr)
	} else {
		glog.Infoln("Escort slave connected successfully!")
		// 定时发送ping
		go func() {
			timer := time.NewTicker(10 * time.Second)
			for {
				select {
				case <-timer.C:
					if err := ping(conn); err != nil {
						glog.Errorf("[ERROR] Escort slave send ping err: %s", err.Error())
						// ping发生错误，切换
						doSwitch()
					}
				}
			}
		}()

		handler(conn)
	}
}

func Work() {
	connectServer()
}
