// 通知相关
package main

import (
	"github.com/gen2brain/beeep"
)

const logoFile = "acfunlogo.ico"

var logoFileLocation string

// 添加订阅指定uid的直播提醒
func addNotify(uid uint) {
	isExist := false
	streamers.mu.Lock()
	for i, s := range streamers.current {
		if s.UID == uid {
			isExist = true
			if s.Notify {
				logger.Println("已经订阅过" + s.ID + "的开播提醒")
			} else {
				streamers.current[i].Notify = true
				logger.Println("成功订阅" + s.ID + "的开播提醒")
			}
		}
	}
	streamers.mu.Unlock()

	if !isExist {
		id := getID(uid)
		if id == "" {
			logger.Println("不存在这个用户")
			return
		}

		newStreamer := streamer{UID: uid, ID: id, Notify: true, Record: false}
		streamers.mu.Lock()
		streamers.current = append(streamers.current, newStreamer)
		streamers.mu.Unlock()
		logger.Println("成功订阅" + id + "的开播提醒")
	}

	saveConfig()
}

// 取消订阅指定uid的直播提醒
func delNotify(uid uint) {
	streamers.mu.Lock()
	for i, s := range streamers.current {
		if s.UID == uid {
			if s.Record {
				streamers.current[i].Notify = false
			} else {
				deleteStreamer(uid)
			}
			logger.Println("成功取消订阅" + s.ID + "的开播提醒")
		}
	}
	streamers.mu.Unlock()

	saveConfig()
}

// 桌面通知
func desktopNotify(notifyWords string) {
	beeep.Alert("AcFun直播通知", notifyWords, logoFileLocation)
}
