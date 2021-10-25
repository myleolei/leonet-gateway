package core

//应答,参考SOCKS5协议
type Response struct {

	//协议版本号
	Ver byte

	//响应码
	// 0x00表示成功
	// 0x01普通SOCKS服务器连接失败
	// 0x02现有规则不允许连接
	// 0x03网络不可达
	// 0x04主机不可达
	// 0x05连接被拒
	// 0x06 TTL超时
	// 0x07不支持的命令
	// 0x08不支持的地址类型
	// 0x09 - 0xFF未定义
	Rep byte

	//目标地址类型
	// 0x01 IPv4地址，DST.ADDR部分4字节长度
	// 0x04 IPv6地址，16个字节长度
	Atyp byte

	//目标地址ip
	DstAddr *[]byte

	//目标地址端口
	DstPort int16
}
