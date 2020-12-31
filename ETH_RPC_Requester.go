package main


type ETHRPCRequester struct {
	// nonceManager *NonceManager // nonce 管理者实例
	client *ETHRPCClient // 小写字母开头， 私有
}

func NewETHRPCRequester(nodeUrl string) *ETHRPCRequester  {
	requester := &ETHRPCRequester{}
	// 实例化 nonce管理器
	// requester.nonceManager = NewNonceManager()
	// 实例化 rpc客户端
	requester.client = NewETHRPCClient(nodeUrl)
	return requester
}

func (r *ETHRPCRequester) GetTransactionByHash(txHash string) (model.Transaction, error) {
	methodName := "eth_getTransactionByHash"
	result := model.Transaction{}
	// 下面call 函数的result参数传入的是 model.Transaction 结构体的引用，
	// 这样内部所设置的值在函数执行完之后才能有效果
	err := r.client.GetRpc().Call(*result, methodName, txHash)
	return result, err
}