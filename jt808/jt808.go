package jt808

import "go-jt808/common"

const (
	T0001 common.Proto = 0x0001 // 终端通用应答
	T0002 common.Proto = 0x0002 // 终端心跳
	T0003 common.Proto = 0x0003 // 终端注销
	T0004 common.Proto = 0x0004 // 查询服务器时间 2019 new
	T0005 common.Proto = 0x0005 // 终端补传分包请求 2019 new
	T0100 common.Proto = 0x0100 // 终端注册
	T0102 common.Proto = 0x0102 // 终端鉴权 2019 modify
	T0104 common.Proto = 0x0104 // 查询终端参数应答
	T0107 common.Proto = 0x0107 // 查询终端属性应答
	T0108 common.Proto = 0x0108 // 终端升级结果通知
	T0200 common.Proto = 0x0200 // 位置信息汇报
	T0201 common.Proto = 0x0201 // 位置信息查询应答
	T0301 common.Proto = 0x0301 // 事件报告 2019 del
	T0302 common.Proto = 0x0302 // 提问应答 2019 del
	T0303 common.Proto = 0x0303 // 信息点播_取消 2019 del
	T0500 common.Proto = 0x0500 // 车辆控制应答
	T0608 common.Proto = 0x0608 // 查询区域或线路数据应答 2019 new
	T0700 common.Proto = 0x0700 // 行驶记录数据上传
	T0701 common.Proto = 0x0701 // 电子运单上报
	T0702 common.Proto = 0x0702 // 驾驶员身份信息采集上报 2019 modify
	T0704 common.Proto = 0x0704 // 定位数据批量上传
	T0705 common.Proto = 0x0705 // CAN总线数据上传
	T0800 common.Proto = 0x0800 // 多媒体事件信息上传
	T0801 common.Proto = 0x0801 // 多媒体数据上传
	T0802 common.Proto = 0x0802 // 存储多媒体数据检索应答
	T0805 common.Proto = 0x0805 // 摄像头立即拍摄命令应答
	T0900 common.Proto = 0x0900 // 数据上行透传
	T0901 common.Proto = 0x0901 // 数据压缩上报
	T0A00 common.Proto = 0x0A00 // 终端RSA公钥

	//0x0F00 - 0x0FFF // 终端上行消息保留

	T8001 common.Proto = 0x8001 // 平台通用应答
	T8003 common.Proto = 0x8003 // 服务器补传分包请求
	T8004 common.Proto = 0x8004 // 查询服务器时间应答 2019 new
	T8100 common.Proto = 0x8100 // 终端注册应答

	T8103 common.Proto = 0x8103 // 设置终端参数
	T8104 common.Proto = 0x8104 // 查询终端参数
	T8105 common.Proto = 0x8105 // 终端控制
	T8106 common.Proto = 0x8106 // 查询指定终端参数
	T8107 common.Proto = 0x8107 // 查询终端属性
	T8108 common.Proto = 0x8108 // 下发终端升级包
	T8201 common.Proto = 0x8201 // 位置信息查询
	T8202 common.Proto = 0x8202 // 临时位置跟踪控制
	T8203 common.Proto = 0x8203 // 人工确认报警消息
	T8204 common.Proto = 0x8204 // 服务器向终端发起链路检测请求 2019 new
	T8300 common.Proto = 0x8300 // 文本信息下发 2019 modify
	T8301 common.Proto = 0x8301 // 事件设置 2019 del
	T8302 common.Proto = 0x8302 // 提问下发 2019 del
	T8303 common.Proto = 0x8303 // 信息点播菜单设置 2019 del
	T8304 common.Proto = 0x8304 // 信息服务 2019 del
	T8400 common.Proto = 0x8400 // 电话回拨
	T8401 common.Proto = 0x8401 // 设置电话本
	T8500 common.Proto = 0x8500 // 车辆控制 2019 modify
	T8600 common.Proto = 0x8600 // 设置圆形区域 2019 modify
	T8601 common.Proto = 0x8601 // 删除圆形区域
	T8602 common.Proto = 0x8602 // 设置矩形区域 2019 modify
	T8603 common.Proto = 0x8603 // 删除矩形区域
	T8604 common.Proto = 0x8604 // 设置多边形区域 2019 modify
	T8605 common.Proto = 0x8605 // 删除多边形区域
	T8606 common.Proto = 0x8606 // 设置路线
	T8607 common.Proto = 0x8607 // 删除路线
	T8608 common.Proto = 0x8608 // 查询区域或线路数据 2019 new
	T8700 common.Proto = 0x8700 // 行驶记录仪数据采集命令
	T8701 common.Proto = 0x8701 // 行驶记录仪参数下传命令
	T8702 common.Proto = 0x8702 // 上报驾驶员身份信息请求

	T8800 common.Proto = 0x8800 // 多媒体数据上传应答

	T8801 common.Proto = 0x8801 // 摄像头立即拍摄命令
	T8802 common.Proto = 0x8802 // 存储多媒体数据检索
	T8803 common.Proto = 0x8803 // 存储多媒体数据上传
	T8804 common.Proto = 0x8804 // 录音开始命令
	T8805 common.Proto = 0x8805 // 单条存储多媒体数据检索上传命令
	T8900 common.Proto = 0x8900 // 数据下行透传
	T8A00 common.Proto = 0x8A00 // 平台RSA公钥

	//0x8F00 - 0x8FFF // 平台下行消息保留
	//0xE000 - 0xEFFF // 厂商自定义上行消息 2019 new
	//0xF000 - 0xFFFF // 厂商自定义下行消息 2019 new
)
