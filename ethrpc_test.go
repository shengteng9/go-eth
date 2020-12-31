package main

import (
	"testing"
	"fmt"
)

func TestNewETHRPCClient(t *testing.T)  {
	// 首先是一个格式正确的链接测试初始化
	client2 := NewETHRPCClient("www.nihao.com").GetRpc()	
	if client2 == nil {
		fmt.Println("初始化失败")
	}
	client := NewETHRPCClient("123://456").GetRpc()
	if client == nil {
		fmt.Println("初始化失败")
	}
}