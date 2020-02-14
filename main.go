package main

import (
	"context"
	"fmt"
	"log"
	"os"

	cat "github.com/wakuwaku3/example.grpc.proto"
	"google.golang.org/grpc"
)

func main() {
	apiHost := os.Getenv("API_HOST")
	apiPort := os.Getenv("API_PORT")
	target := fmt.Sprintf("%s:%s", apiHost, apiPort)
	//sampleなのでwithInsecure
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := cat.NewCatClient(conn)
	message := &cat.GetMyCatMessage{TargetCat: "tama"}
	res, err := client.GetMyCat(context.TODO(), message)
	fmt.Printf("result:%#v \n", res)
	fmt.Printf("error::%#v \n", err)
}
