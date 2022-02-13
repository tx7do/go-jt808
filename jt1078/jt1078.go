package jt1078

import "go-jt808/common"

const (
	T1003     common.Proto = 0x1003     //终端上传音视频属性
	T1005     common.Proto = 0x1005     //终端上传乘客流量
	T1205     common.Proto = 0x1205     //终端上传音视频资源列表
	T1206     common.Proto = 0x1206     //文件上传完成通知
	T9003     common.Proto = 0x9003     //查询终端音视频属性
	T9101     common.Proto = 0x9101     //实时音视频传输请求
	T9102     common.Proto = 0x9102     //音视频实时传输控制
	T9105     common.Proto = 0x9105     //实时音视频传输状态通知
	T9201     common.Proto = 0x9201     //平台下发远程录像回放请求
	T9202     common.Proto = 0x9202     //平台下发远程录像回放控制
	T9205     common.Proto = 0x9205     //查询资源列表
	T9206     common.Proto = 0x9206     //文件上传指令
	T9207     common.Proto = 0x9207     //文件上传控制
	T9301     common.Proto = 0x9301     //云台旋转
	T9302     common.Proto = 0x9302     //云台调整焦距控制
	T9303     common.Proto = 0x9303     //云台调整光圈控制
	T9304     common.Proto = 0x9304     //云台雨刷控制
	T9305     common.Proto = 0x9305     //红外补光控制
	T9306     common.Proto = 0x9306     //云台变倍控制
	T30316364 common.Proto = 0x30316364 //实时音视频流及透传数据传输
	Wakeup    string       = "WAKEUP"   //平台手工唤醒请求_短消息
)

func getid() common.Proto {
	return T30316364
}
