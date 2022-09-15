package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	pb "github.com/Kendovvul/Ejemplo/Proto"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

var n_merc int

func main() {
	qName := "Emergencias" //Nombre de la cola
	hostQ := "dist145"     //Host de RabbitMQ 172.17.0.1
	//hostS := "localhost"   //Host de un Laboratorio
	n_merc = 2

	//Conexion con RabbitMQ
	connQ, err := amqp.Dial("amqp://test:test@" + hostQ + ":5672")
	if err != nil {
		log.Fatal(err)
	}
	defer connQ.Close()

	ch, err := connQ.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	//Se crea la cola en RabbitMQ
	q, err := ch.QueueDeclare(qName, false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	ch.QueuePurge(qName, true)
	fmt.Println(q)

	// Sucede la magia
	fmt.Println("Esperando Emergencias\n")
	chDelivery, err := ch.Consume(qName, "", true, false, false, false, nil) //obtiene la cola de RabbitMQ
	if err != nil {
		log.Fatal(err)
	}

	for delivery := range chDelivery {

		if n_merc > 0 {
			// Obtiene el primer mensaje de la cola con el nombre del lab
			labName := string(delivery.Body)
			labPort := ""
			labHost := ""

			switch labName {
			case "Renca":
				labPort = ":50051"
				labHost = "dist145"
			case "Pohang":
				labPort = ":50052"
				labHost = "dist146"
			case "Kampala":
				labPort = ":50053"
				labHost = "dist147"
			case "Pripiat":
				labPort = ":50054"
				labHost = "dist148"
			default:
				// Raise error
				fmt.Println("Error al asignar los puertos y el host!")
			}

			fmt.Println("------------------------\nHost: " + labHost + "\nPort: " + labPort)

			// BORRAR ESTA COSAAAA
			//labHost = "localhost"

			conn, err := grpc.Dial(labHost+labPort, grpc.WithInsecure()) //crea la conexion sincrona con el laboratorio
			if err != nil {
				panic("No se pudo conectar con el servidor" + err.Error())
			}
			defer conn.Close()

			serviceCliente := pb.NewMessageServiceClient(conn)

			squadSend := strconv.Itoa(n_merc)
			fmt.Println("mandando el squad " + squadSend)

			n_merc--
			consultas := 0

			for {
				// TODO: informar del escuadron mandado

				time.Sleep(5 * time.Second)
				// Envio del mensaje al lab
				res, err := serviceCliente.Intercambio(context.Background(),
					&pb.Message{
						Body: "Estallido resuelto?",
					})
				if err != nil {
					panic("No se puede crear el mensaje " + err.Error())
				}
				consultas++

				fmt.Println("\tEstallido resuelto?: " + res.Body)

				if res.Body == "SI" { // Se resolvio el estallido
					// TODO: Informar que squad volvio
					n_merc++

					// TODO: Cerrar la conexión
					fmt.Println("Cerrando conexion con " + labName)
					_, err := serviceCliente.Intercambio(context.Background(),
						&pb.Message{
							Body: "STOP",
						})
					if conn.GetState().String() != "IDLE" {
						panic("No se puede cerrar la conexion " + err.Error())
					}

					// TODO: Escribir en el archivo "SOLICITUDES.txt"
					time.Sleep(5 * time.Second)
					break
				}
				//time.Sleep(5 * time.Second)
			}

			fmt.Println("Finalizó el trabajo del escuadrón " + squadSend)

		}
	}
	fmt.Println("fin?")

}
