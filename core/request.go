package core

//请求,参考SOCKS5协议
type Request struct {
	//协议版本号
	Ver byte
	//指令
	//0x01表示CONNECT请求
	//0x02表示BIND请求
	//0x03表示UDP转发
	Cmd byte
	//目标地址类型
	// 0x01 IPv4地址，DST.ADDR部分4字节长度
	// 0x04 IPv6地址，16个字节长度
	Atyp byte
	//目标地址ip
	DstAddr *[]byte
	//目标地址端口
	DstPort int16
}
