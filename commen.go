package main

import (
	"encoding/json"
	"strconv"

	"github.com/gorilla/websocket"
)

func GetMsgType(msg []byte) (uint64, string, map[string]interface{}) {
	var my_msg interface{}
	json.Unmarshal([]byte(msg), &my_msg)
	msgs := my_msg.(map[string]interface{})
	msg_type := uint64(msgs["msg_type"].(float64))
	msg_uuid := msgs["uuid"].(string)
	msg_data := msgs["msg_data"].(map[string]interface{})
	//log.Println(msg_data)
	return msg_type, msg_uuid, msg_data
}

//单发送
func SetMsgType(conn *websocket.Conn, msg_type int, msg_data string) {
	msg_str := []byte(`{"msg_type":` + strconv.Itoa(msg_type) + `,"msg_data":"` + msg_data + `"}`)
	var f interface{}
	json.Unmarshal(msg_str, &f)
	conn.WriteJSON(f)
}

//添加波次
func SetMyWellen(wellen string, uuid string) {
	for r, user_d := range users {
		if uuid == user_d.uuid {
			users[r].wellen = wellen
		}
	}
}

//判断波次是否重复
func WellenIsRepeat(wellen string) bool {
	var is_repeat = false
	for _, user_d := range users {
		if wellen == user_d.wellen {
			is_repeat = true
		}
	}
	return is_repeat
}

//清除多余用户
func UnsetUser(uuid string) {
	var marker int
	for r, user_d := range users {
		if uuid == user_d.uuid {
			marker = r
		}
	}
	users = append(users[:marker], users[marker+1:]...)
}
