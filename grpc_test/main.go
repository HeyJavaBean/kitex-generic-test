package main

import (
	"context"
	"fmt"
	"generic-kitex-test/grpc_test/kitex_gen/hello/cloudwego/team"
	"generic-kitex-test/grpc_test/kitex_gen/hello/cloudwego/team/greet"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/transport"
	"net"
	"time"
)

type GreetImpl struct {
}

func (g *GreetImpl) Hello(ctx context.Context, req *team.MyReq) (r *team.MyResp, err error) {
	fmt.Println("server : ", req)
	return &team.MyResp{Text: "server reply"}, nil
}

func RunServer() {
	addr, _ := net.ResolveTCPAddr("tcp", "localhost:8888")
	s := greet.NewServer(&GreetImpl{}, server.WithServiceAddr(addr))
	s.Run()
}

func main() {

	go RunServer()

	time.Sleep(1 * time.Second)

	fmt.Println("\nnormal")
	normalClientCall()
	fmt.Println("\nbinary")
	binaryGenericCall()
	fmt.Println("\njson")
	jsonGenericCall()
	fmt.Println("\nmap")
	mapGenericCall()
}

func clientOptionsSuite() []client.Option {
	return []client.Option{client.WithHostPorts("hostport:8888"), client.WithTransportProtocol(transport.GRPC)}
}

func jsonGenericCall() {

	p, err := generic.NewThriftFileProvider("./thrift_test/example.thrift")
	if err != nil {
		panic(err)
	}

	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		return
	}
	genericCli, err := genericclient.NewClient("a.b.c", g, clientOptionsSuite()...)

	methodName := "Hello"

	jsonData := `
	{
		"name": "json hello",
		"id":   "789",
	}
	`
	result, err := genericCli.GenericCall(context.Background(), methodName, jsonData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("resp :", result)
}

func mapGenericCall() {

	p, err := generic.NewThriftFileProvider("./thrift_test/example.thrift")
	if err != nil {
		panic(err)
	}

	g, err := generic.MapThriftGeneric(p)
	if err != nil {
		return
	}
	genericCli, err := genericclient.NewClient("a.b.c", g, clientOptionsSuite()...)

	methodName := "Hello"

	mapData := map[string]interface{}{
		"name": "map hello",
		"id":   "123",
	}

	result, err := genericCli.GenericCall(context.Background(), methodName, mapData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("resp :", result)
}

func binaryGenericCall() {
	// todo
	genericCli, err := genericclient.NewClient("a.b.c", generic.BinaryGrpcGeneric(), clientOptionsSuite()...)

	// 要用 method 封装的结构体
	args := &team.MyReq{
		Name: "hello",
		Id:   "0000",
	}
	var buf []byte
	// todo mock buffer

	if err != nil {
		return
	}
	result, err := genericCli.GenericCall(context.Background(), methodName, buf, callopt.WithHostPort("localhost:8888"))
	// 实际上 uint8
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("raw binary resp :", result)
	respStruct := &team.MyResp{}
	method, seq, err := codec.Decode(result.([]byte), respStruct)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("method :", method)
	fmt.Println("seq :", seq)
	fmt.Println("result :", respStruct)
}

func normalClientCall() {

	cli, _ := greet.NewClient("a.b.c", clientOptionsSuite()...)
	resp, err := cli.Hello(context.Background(), &team.MyReq{
		Name: "Lee",
		Id:   "123",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}
