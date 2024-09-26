package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	hellopb "mygrpc/pkg/grpc"
)

var (
	scanner *bufio.Scanner
	client  hellopb.GreetingServiceClient
)

func main() {
	fmt.Println("start gRPC client.")

	scanner = bufio.NewScanner(os.Stdin)

	address := "localhost:8080"
	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Connection failed: ", err)
		return
	}
	defer conn.Close()

	client = hellopb.NewGreetingServiceClient(conn)

	for {
		fmt.Println("1: send Request")
		fmt.Println("2: HelloServerStream")
		fmt.Println("3: HelloClientStream")
		fmt.Println("4: HelloBiStreams")
		fmt.Println("5: exit")
		fmt.Print("please enter the number: ")

		scanner.Scan()
		in := scanner.Text()

		switch in {
		case "1":
			Hello()
		case "2":
			HelloServerStream()
		case "3":
			HelloClientStream()
		case "4":
			HelloBiStreams()
		case "5":
			fmt.Println("bye.")
			goto M
		}
	}
M:
}

func HelloBiStreams() {
	stream, err := client.HelloBiStreams(context.Background())
	if err != nil {
		fmt.Println("failed to call HelloBiStreams: ", err)
		return
	}

	sendNum := 5
	fmt.Printf("please enter %d names.\n", sendNum)

	var sendEnd, recvEnd bool
	sendCount := 0
	for !(sendEnd && recvEnd) {
		if !sendEnd {
			scanner.Scan()
			name := scanner.Text()

			sendCount++
			if err := stream.Send(&hellopb.HelloRequest{
				Name: name,
			}); err != nil {
				fmt.Println("failed to send request: ", err)
				return
				sendEnd = true
			}

			if sendCount == sendNum {
				sendEnd = true
				if err := stream.CloseSend(); err != nil {
					fmt.Println("failed to close send: ", err)
				}
			}
		}

		if !recvEnd {
			if res, err := stream.Recv(); err != nil {
				if !errors.Is(err, io.EOF) {
					fmt.Println("failed to receive response: ", err)
				}
				recvEnd = true
			} else {
				fmt.Println(res.GetMessage())
			}
		}
	}
}

func HelloClientStream() {
	stream, err := client.HelloClientStream(context.Background())
	if err != nil {
		fmt.Println("failed to call HelloClientStream: ", err)
		return
	}

	sendCount := 5
	fmt.Printf("please enter %d names.\n", sendCount)
	for i := 0; i < sendCount; i++ {
		scanner.Scan()
		name := scanner.Text()

		if err := stream.Send(&hellopb.HelloRequest{
			Name: name,
		}); err != nil {
			fmt.Println("failed to send request: ", err)
			return
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Println("failed to close and receive: ", err)
	} else {
		fmt.Println("response: ", res.GetMessage())
	}
}

func HelloServerStream() {
	fmt.Print("Please enter your name: ")
	scanner.Scan()
	name := scanner.Text()

	req := &hellopb.HelloRequest{
		Name: name,
	}
	stream, err := client.HelloServerStream(context.Background(), req)
	if err != nil {
		fmt.Println("failed to call HelloServerStream: ", err)
		return
	}

	for {
		res, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("all the responses have already received.")
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
	}
}

func Hello() {
	fmt.Print("please enter your name: ")
	scanner.Scan()
	name := scanner.Text()

	req := &hellopb.HelloRequest{
		Name: name,
	}
	res, err := client.Hello(context.Background(), req)
	if err != nil {
		log.Fatal("failed to call Hello: ", err)
		return
	} else {
		log.Println("response: ", res.GetMessage())
	}
}
