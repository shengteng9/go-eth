package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/rpc"
)

type ETHRPCClient struct {
	NodeUrl string // 代表节点的url链接
	client *rpc.Client // 代表rpc客户端句柄实例
}

// NewETHRPCClient 代表的是新建一个“RPC”客户端
// 入参nodeUrl就是节点的链接，返回的是带有*的ETHRPCClient对象指针
func NewETHRPCClient(nodeUrl string) *ETHRPCClient {
	// & 符号代表的是取指针
	client := &ETHRPCClient {
		NodeUrl: nodeUrl,
	}
	client.initRpc() // 进行初始化 rpc哭护短句柄实例
	return client
}


// 初始化rpc客户端句柄实例
func (erc *ETHRPCClient) initRpc()  {
	// 使用go-ethereum 库中的rpc库来初始化
	// DialHTTP的意思是使用http版本的rpc实现方式
	rpcClient, err := rpc.DialHTTP(erc.NodeUrl)
	if err != nil {
		// 初始化失败， 终结程序，并将错误信息显示到控制台中
		errInfo := fmt.Errorf("初始化 rpc client 失败 %s", err.Error()).Error()
		panic(errInfo)
	}
	// 初始化成功，将新实例化的rpc句柄赋值给ETHRPCClient 结构体里面的client
	erc.client = rpcClient
}



// go中，大写字母开头的变量或者函数（方法）才能够被外部引用
// 小写字母的变量或函数（方法）只能内部调用
// GetRpc 函数（方法）是为了方便外部能够获取client *rpc.Client, 以便进行访问
func (erc *ETHRPCClient) GetRpc() *rpc.Client {
	if erc.client == nil {
		erc.initRpc()
	}
	return erc.client
}