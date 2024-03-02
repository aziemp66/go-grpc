package main

import (
	"context"
	"log"
	"time"

	invoicer "github.com/aziemp66/go-grpc/client/invoicer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8089", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	c := invoicer.NewInvoicerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Create(ctx, &invoicer.CreateRequest{
		Amount: &invoicer.Amount{
			Amount:   200000,
			Currency: "IDR",
		},
		From: "azie",
		To:   "fadhil",
	})
	if err != nil {
		panic(err)
	}

	log.Println(string(r.Docx))

	log.Println(string(r.Pdf))
}
