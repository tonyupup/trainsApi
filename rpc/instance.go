package rpc1

import (
	"google.golang.org/grpc"
	"log"
	"time"
)

var TrainsClient S_TrainsClient

func init() {
	if client, err := grpc.Dial("10.1.1.103:30001", grpc.WithInsecure(),grpc.WithTimeout(time.Second*3)); err != nil {
		log.Printf("dial server error.%s", err.Error())
		return
	} else {
		TrainsClient = NewS_TrainsClient(client)
		log.Println("Connected rpc server .")
	}

}
