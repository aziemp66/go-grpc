package main

import (
	"context"
	"log"
	"net"

	"github.com/aziemp66/go-grpc/invoicer"
	"google.golang.org/grpc"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s myInvoicerServer) Create(context.Context, *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf:  []byte("Pdf File"),
		Docx: []byte("Docx File"),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("Cannot create listener : %s", err.Error())
	}

	serverRegistrar := grpc.NewServer()
	service := &myInvoicerServer{}

	invoicer.RegisterInvoicerServer(serverRegistrar, service)

	log.Println("Service is running")
	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to server : %s", err.Error())
	}
}
