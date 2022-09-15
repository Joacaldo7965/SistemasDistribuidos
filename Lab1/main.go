package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	pb "github.com/Kendovvul/Ejemplo/Proto"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

var serv *grpc.Server

type server struct {
	pb.UnimplementedMessageServiceServer
}

func (s *server) Intercambio(ctx context.Context, msg *pb.Message) (*pb.Message, error) {

	if msg.Body == "STOP" {
		serv.Stop()
		return &pb.Message{Body: ""}, nil
	}

	fmt.Print("\t" + msg.Body + ": ")

	// Probabilidad de contencion
	if rand.Float64() < 0.6 {
		fmt.Println("SI")
		fmt.Println("Devolviendo escuadron")
		return &pb.Message{Body: "SI"}, nil
	} else {
		fmt.Println("NO")
		return &pb.Message{Body: "NO"}, nil
	}
}

func main() {
	// Random seed
	rand.Seed(time.Now().UnixNano())

	labName := "Pripiat"                                           //nombre del laboratorio
	qName := "Emergencias"                                         //nombre de la cola
	hostQ := "dist145"                                             //ip del servidor de RabbitMQ 172.17.0.1
	connQ, err := amqp.Dial("amqp://test:test@" + hostQ + ":5672") //conexion con RabbitMQ

	if err != nil {
		log.Fatal(err)
	}
	defer connQ.Close()

	ch, err := connQ.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	fmt.Println("se prendio la wea")

	for {
		// Every 5 seconds
		time.Sleep(5 * time.Second)

		if rand.Float64() < 0.8 { // Sucede estallido
			fmt.Println("------------------------\nSucede estallido social")
			// Se genera y envia una solicitud con el nombre del lab
			err = ch.Publish("", qName, false, false,
				amqp.Publishing{
					Headers:     nil,
					ContentType: "text/plain",
					Body:        []byte(labName), //Contenido del mensaje
				})

			if err != nil {
				log.Fatal(err)
			}

			// Escuchando respuesta
			listener, err := net.Listen("tcp", ":50054") //conexion sincrona
			if err != nil {
				panic("La conexion no se pudo crear" + err.Error())
			}

			// Se espera y establece conexion con server
			serv = grpc.NewServer()
			pb.RegisterMessageServiceServer(serv, &server{})

			if err = serv.Serve(listener); err != nil {
				panic("El server no se pudo iniciar" + err.Error())
			}
		}
		fmt.Println("------------------------\nNo pasa nada")

	}

}
