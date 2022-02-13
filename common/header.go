package common

import "fmt"

type Proto int

type Header struct {
	Length       uint32  // 数据长度(包括头标识、数据头、数据体和尾标识)
	SerialNo     uint32  // 报文序列号
	Type         uint16  // 业务数据类型
	GNSSCenterID uint32  // 下级平台接入码，上级平台给下级平台分配的唯一标识号
	Version      [3]byte // 协议版本号标识， 3 字节 上下级平台之间采用的标准协议版本编号；长度为三个字节来表示： 0x01 0x02 x0F 表示的版本号是V1.2.15,依此类推
	Encrypt      byte    // 报文加密标识位：0表示报文不加密，1表示报文加密
	EncryptKey   uint32  // 数据加密的密钥，长度为四个字节
}

func (h Header) Len() int {
	return 22
}

func (h Header) PacketLen() int {
	return int(h.Length)
}

func (h Header) String() string {
	return ""
}

func (h Header) VersionString() string {
	return fmt.Sprintf("%d.%d.%d", h.Version[0], h.Version[1], h.Version[2])
}

func (h Header) Encode() []byte {
	var data []byte
	return data
}

func (h Header) Decode(data []byte) {

}
