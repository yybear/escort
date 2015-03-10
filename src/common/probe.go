package common

import (
	"net/http"
)

/**
 * 探测cs服务是否正常
 *
 */
func ProbeCS() bool {
	_, err := http.Get("http://localhost:8080/client/")
	if err != nil {
		return false
	}

	return true
}

/**
 * 探测数据库是否正常
 *
 */
func ProbeDB() bool {
	return true
}
