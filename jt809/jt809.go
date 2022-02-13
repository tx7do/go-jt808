package jt809

import "go-jt808/common"

const (
	// 链路管理类
	T1001 common.Proto = 0x1001 // 主链路登录请求消息 (主链路)
	T1002 common.Proto = 0x1002 // 主链路登录应答消息 (主链路)
	T1003 common.Proto = 0x1003 // 主链路注销请求消息 (主链路)
	T1004 common.Proto = 0x1004 // 主链路注销应答消息 (主链路)
	T1005 common.Proto = 0x1005 // 主链路连接保持请求消息 (主链路)
	T1006 common.Proto = 0x1006 // 主链路连接保持应答消息 (主链路)
	T1007 common.Proto = 0x1007 // 主链路断开通知消息 (从链路)
	T1008 common.Proto = 0x1008 // 下级平台主动关闭链路通知消息 (从链路)
	T9001 common.Proto = 0x9001 // 从链路连接请求消息 (从链路)
	T9002 common.Proto = 0x9002 // 从链路连接应答消息 (从链路)
	T9003 common.Proto = 0x9003 // 从链路注销请求消息 (从链路)
	T9004 common.Proto = 0x9004 // 从链路注销应答消息 (从链路)
	T9005 common.Proto = 0x9005 // 从链路连接保持请求消息 (从链路)
	T9006 common.Proto = 0x9006 // 从链路连接保持应答消息 (从链路)
	T9007 common.Proto = 0x9007 // 从链路断开通知消息 (主链路)
	T9008 common.Proto = 0x9008 // 上级平台主动关闭链路通知消息 (主链路)

	// 信息统计类
	T9101 common.Proto = 0x9101 // 发送车辆定位信息数量通知消息 (主链路)
	T9102 common.Proto = 0x9102 // 平台链路连接情况与车辆定位消息传输情况上报请求消息 (从链路)
	T1102 common.Proto = 0x1102 // 平台链路连接情况与车辆定位消息传输情况上报应答消息 (主链路)
	T9103 common.Proto = 0x9103 // 下发平台间消息序列号通知消息 (从链路)
	T1103 common.Proto = 0x1103 // 上传平台间消息序列号通知消息 (主链路)

	// 车辆动态信息交换类
	T1200 common.Proto = 0x1200 //  主链路动态信息交换消息 (主链路)
	T9200 common.Proto = 0x9200 //  从链路动态信息交换消息 (从链路)

	// 平台间信息交互类
	T1300 common.Proto = 0x1300 //  主链路平台间信息交互消息 (主链路)
	T9300 common.Proto = 0x9300 //  从链路平台间信息交互消息 (从链路)

	// 车辆报警信息交互类
	T1400 common.Proto = 0x1400 //  主链路报警信息交互信息 (主链路)
	T9400 common.Proto = 0x9400 //  从链路报警信息交互信息 (从链路)

	// 车辆监管类
	T1500 common.Proto = 0x1500 //  主链路车辆监管消息 (主链路)
	T9500 common.Proto = 0x9500 //  从链路车辆监管消息 (从链路)

	// 车辆静态信息交换类
	T1600 common.Proto = 0x1600 //  主链路静态信息交换消息 (主链路)
	T9600 common.Proto = 0x9600 //  从链路静态信息交换消息 (从链路)
)
