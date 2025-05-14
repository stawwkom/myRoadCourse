package main

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
	desc "template_course/week_1/grpc/pkg/simpleNote_v1"
)

const grpcPort = 50051

type server struct { // тут мы в структуру сервера встариваем наш сгенерированный интерфейс grpc Server
	desc.UnimplementedSimpleNoteV1Server
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Create note: %+v", req.GetInfo())

	id := gofakeit.Number(1000, 9999)

	return &desc.CreateResponse{
		Id: int64(id),
	}, nil
}

// Реализуем метод Get

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("Note id: %d", req.GetId())

	return &desc.GetResponse{
		Note: &desc.Note{
			Id: req.GetId(),
			Info: &desc.NoteInfo{
				Title:    gofakeit.BeerName(),
				Content:  gofakeit.IPv4Address(),
				Author:   gofakeit.Name(),
				IsPublic: gofakeit.Bool(),
			},
			CreatedAt: timestamppb.New(gofakeit.Date()),
			UpdatedAt: timestamppb.New(gofakeit.Date()),
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("Failded to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterSimpleNoteV1Server(s, &server{})

	log.Printf("server listening at %d", grpcPort)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to server %v", err)

	}
}
