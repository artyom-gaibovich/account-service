package main

import (
	"account-service/pkg/account_v1"
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
)

func main() {
	//c := app.NewContext()
	//log.Fatal(c.Server().Start())

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	account_v1.RegisterAccountV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

const grpcPort = 50051

type server struct {
	account_v1.UnimplementedAccountV1Server
}

func (s *server) Get(ctx context.Context, req *account_v1.GetRequest) (*account_v1.GetResponse, error) {
	log.Printf("Account id: %d", req.GetId())
	return &account_v1.GetResponse{
		Id:        gofakeit.Int64(),
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Role:      1,
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}, nil
}

func (s *server) Create(ctx context.Context, req *account_v1.CreateRequest) (*account_v1.CreateResponse, error) {
	return &account_v1.CreateResponse{Id: gofakeit.Int64()}, nil
}

func (s *server) Update(context.Context, *account_v1.UpdateRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *server) Delete(context.Context, *account_v1.DeleteRequest) (*emptypb.Empty, error) {
	return nil, nil
}
