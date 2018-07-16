package main

import (
	"github.com/joho/godotenv"
	"log"
	"time"
	"github.com/jason/blockchain-demo/pos/model"
	"github.com/jason/blockchain-demo/pos/service"
	"github.com/davecgh/go-spew/spew"
	"os"
	"net"
)

func main() {
	err := godotenv.Load()
	if err != nil{
		log.Fatal(err)
	}

	t := time.Now()
	genesisBlock := model.Block{}
	genesisBlock = model.Block{0,t.String(),0,service.CalculateBlockHash(genesisBlock),"",""}
	spew.Dump(genesisBlock)  //美观的将数据打印出，创始块
	model.Blockchain = append(model.Blockchain,genesisBlock)
	httpPort := os.Getenv("PORT")

	server, err := net.Listen("tcp", ":"+httpPort)
	if err!=nil{
		log.Fatal(err)
	}
	log.Println("HTTP Server Listening on port :",httpPort)
	defer server.Close()

	go func() {
		for candidate := range model.CandidateBlocks{
			model.Mutex.Lock()
			model.TempBlocks = append(model.TempBlocks,candidate)
			model.Mutex.Unlock()
		}
	}()

	go func() {
		for{
			service.PickWinner()
		}
	}()

	for{
		conn, err:= server.Accept()
		if err !=nil{
			log.Fatal(err)
		}
		go service.HandleConn(conn)
	}

}
