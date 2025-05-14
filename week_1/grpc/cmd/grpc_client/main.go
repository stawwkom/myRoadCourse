package main

import (
	"context"
	"github.com/brianvoe/gofakeit"
	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	desc "template_course/week_1/grpc/pkg/simpleNote_v1"
	"time"
)

const (
	address = "localhost:50051"
	noteID  = 12
)

func main() {
	// создаем gRPC соедение с сервером по указанному адресу
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect server: %v", address)
	}
	// закрываем соединение по окончанию нашей работы с программой
	defer conn.Close()
	// Создаем gRPC клиента
	client := desc.NewSimpleNoteV1Client(conn)
	// контекст с таймаутом в 1 сек, для ограничения выполнения gRPC запроса
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel() // для освобождения ресурсов
	// вызов метода получение у сервера
	responseGet, err := client.Get(ctx, &desc.GetRequest{Id: noteID})
	if err != nil {
		log.Fatalf("failed to get not by id: %v", err)
	}

	noteInfo := &desc.NoteInfo{
		Title:    gofakeit.Name(),
		Content:  gofakeit.Company(),
		Author:   gofakeit.IPv4Address(),
		IsPublic: gofakeit.Bool(),
	}

	responseCreate, err := client.Create(ctx, &desc.CreateRequest{Info: noteInfo})
	if err != nil {
		log.Fatalf("Failed to create note %v", err)
	}

	log.Printf(color.BlueString("Note info :\n"), color.GreenString("%+v", responseGet.GetNote()))
	log.Printf(color.BlueString("Created new note with ID: "), color.GreenString("%d", responseCreate.GetId()))

}
