package main

import (
	"fmt"
	"log"
)

func (cli *CLI) pack(addr string) {
	if !ValidateAddress(addr) {
		log.Panic("ERROR: 发送地址非法")
	}

	bc := NewBlockchain() //打开数据库，读取区块链并构建区块链实例
	defer bc.Db.Close()   //转账完毕，关闭数据库
	//txs := conn_recv(addr)

	//开启一个主进程处理交易数据
	txs := recv_tx(addr)
	bc.MineBlock(txs)

	//开启一个进程处理评论数据
	//go recvReview()

	fmt.Println("成功打包交易")
}
