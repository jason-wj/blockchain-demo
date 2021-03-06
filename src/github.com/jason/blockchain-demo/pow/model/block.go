package model

import "sync"

type Block struct {
	Index int  //是区块链中数据记录的位置
	Timestamp string  //是自动确定的，并且是写入数据的时间
	BPM int  //或每分钟跳动，是你的脉率,
	Hash string  //是代表这个数据记录的SHA256标识符
	PrevHash string  //是链中上一条记录的SHA256标识符
	Diffculty int
	Nonce string
}

var Blockchain []Block

type Message struct {
	BPM int
}

var Mutex = &sync.Mutex{}