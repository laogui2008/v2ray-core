package vmess

import (
	"net/http"

	"github.com/v2fly/v2ray-core/v5/common/protocol"
)

var AccoutMap = make(map[string]protocol.Account)

func StartWeb() {
	http.HandleFunc("/getStatusInfo", getStatusInfo)
	http.HandleFunc("/updateStatus", updateStatus)
	http.ListenAndServe("127.0.0.1:22222", nil)
}

func getStatusInfo(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()      // 获取?后面的参数
	email := q.Get("email") // 获取q参数
	var msg string
	if account, ok := AccoutMap[email].(*MemoryAccount); ok {
		msg = account.Status
	} else {
		msg = "error"
	}
	w.Write([]byte(msg))
}

func updateStatus(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()        // 获取?后面的参数
	email := q.Get("email")   // 获取q参数
	status := q.Get("status") // 获取q参数
	if status != "1" && status != "0" {
		w.Write([]byte("status: invalid value"))
		return
	}

	msg := "OK"
	if account, ok := AccoutMap[email].(*MemoryAccount); ok {
		account.Status = status
	} else {
		msg = "error"
	}

	w.Write([]byte(msg))
}
