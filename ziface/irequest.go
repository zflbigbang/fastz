package ziface

/*
IRequest 接口：
实际上是把客户端请求的链接信息 和 请求的数据 包装到了 Request里
*/
type IRequest interface {
	GetConnection() IConnection //获取请求连接信息
	GetData() []byte
	// GetMsgID 得到当前请求消息 ID
	GetMsgID() uint32
}
