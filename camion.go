package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
	"strconv"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := courier.NewCourierServiceClient(conn)

	}
}

func cargaRetail1() { // Función para camión 1
	message := courier.Camion{ // Para identificar el camion que envía
		Id: "01",
	}
	ret1, err := c.PedirRetail(context.Background(), &message) // Busca un retail en cola retail
	if err != nil {
		log.Fatalf("Error when calling PedirRetail: %s", err)
	}
	if ret1.Id != nil { // Caso en que si encontró retail
		ret2, err := c.PedirRetail(context.Background(), &message) // Ahora evalúa si hay un segundo retail en cola retail
		if err != nil {
				log.Fatalf("Error when calling PedirRetail: %s", err)
		}
		if ret2.Id != nil { // Encontró un segundo retail.

			//TIRAR FUNCION YA QUE HAY 2 PAQUETES RT RT

		} else if ret2.Id == nil{ // No encontró un segundo retail.
			prio1, err := c.PedirPrioritario(context.Background(), &message) // Prueba entonces con prioritario.
			if err != nil {
				log.Fatalf("Error when calling PedirPrioritario: %s", err)
			}
		if prio1.Id != nil { // Si encuentra prioritario.
			//TIRAR FUNCION YA QUE HAY 2 PAQUETES RT PR
		} else if prio1.Id == nil { //No encuentra prioritario.
			//TIRAR FUNCION YA QUE HAY 1 PAQUETE RT
		}

		}
	} else if ret1.Id == nil { // No encontró un primer retail.
		prio1, err := c.PedirPrioritario(context.Background(), &message) // Prueba entonces buscar prioritario en cola prioritario.
		if err != nil {
			log.Fatalf("Error when calling PedirPrioritario: %s", err)
		}
		if prio1.Id != nil { // Encuentra prioritario.
			prio2, err := c.PedirPrioritario(context.Background(), &message) // Busca el segundo prioritario.
			if err != nil {
				log.Fatalf("Error when calling PedirPrioritario: %s", err)
			}
			if prio2.Id != nil { // Encuentra segundo prioritario.
				//TIRAR FUNCION YA QUE HAY 2 PAQUETES PR PR 
			} else if prio2.Id == nil { // No encuentra un segundo prioritario.
				//TIRAR FUNCION YA QUE HAY 1 PAQUETE PR
			}			
		}
	}
}

func cargaRetail2() { // Función para camión 2 (similar a la función del camión 1)
	message := courier.Camion{
		Id: "02",
	}
	ret1, err := c.PedirRetail(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling PedirRetail: %s", err)
	}
	if ret1.Id != nil {
		ret2, err := c.PedirRetail(context.Background(), &message)
		if err != nil {
				log.Fatalf("Error when calling PedirRetail: %s", err)
		}
		if ret2.Id != nil {

			//TIRAR FUNCION YA QUE HAY 2 PAQUETES RT RT (recordar hacer comparacion de valores antes)

		} else if ret2.Id == nil{
			prio1, err := c.PedirPrioritario(context.Background(), &message)
			if err != nil {
				log.Fatalf("Error when calling PedirPrioritario: %s", err)
			}
		if prio1.Id != nil {
			//TIRAR FUNCION YA QUE HAY 2 PAQUETES RT PR
		} else if prio1.Id == nil {
			//TIRAR FUNCION YA QUE HAY 1 PAQUETE RT
		}

		}
	} else if ret1.Id == nil {
		prio1, err := c.PedirPrioritario(context.Background(), &message)
		if err != nil {
			log.Fatalf("Error when calling PedirPrioritario: %s", err)
		}
		if prio1.Id != nil {
			prio2, err := c.PedirPrioritario(context.Background(), &message)
			if err != nil {
				log.Fatalf("Error when calling PedirPrioritario: %s", err)
			}
			if prio2.Id != nil {
				//TIRAR FUNCION YA QUE HAY 2 PAQUETES PR PR 
			} else if prio2.Id == nil {
				//TIRAR FUNCION YA QUE HAY 1 PAQUETE PR
			}			
		}
	}
}

func cargaNormal() { // Función para camión 3 (similar a las anteriores, sólo que en vez de retail -> prioritario y prioritario -> normal)
	message := courier.Camion{
		Id: "03",
	}
	prio1, err := c.PedirPrioritario(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling PedirRetail: %s", err)
	}
	if prio1.Id != nil {
		prio2, err := c.PedirPrioritario(context.Background(), &message)
		if err != nil {
				log.Fatalf("Error when calling PedirPrioritario: %s", err)
		}
		if prio2.Id != nil {

			//TIRAR FUNCION YA QUE HAY 2 PAQUETES PR PR

		} else if prio2.Id == nil{
			nor1, err := c.PedirNormal(context.Background(), &message)
			if err != nil {
				log.Fatalf("Error when calling PedirNormal: %s", err)
			}
		if nor1.Id != nil {
			//TIRAR FUNCION YA QUE HAY 2 PAQUETES PR NR
		} else if nor1.Id == nil {
			//TIRAR FUNCION YA QUE HAY 1 PAQUETE PR
		}

		}
	} else if prio1.Id == nil {
		nor1, err := c.PedirNormal(context.Background(), &message)
		if err != nil {
			log.Fatalf("Error when calling PedirNormal: %s", err)
		}
		if nor1.Id != nil {
			nor2, err := c.PedirNormal(context.Background(), &message)
			if err != nil {
				log.Fatalf("Error when calling PedirNormal: %s", err)
			}

			if nor2.Id != nil {
				//TIRAR FUNCION YA QUE HAY 2 PAQUETES NR NR
			} else if nor2.Id == nil {
				//TIRAR FUNCION YA QUE HAY 1 PAQUETE NR
			}			
		}
	}
}

func dospaquetes(){ //Falta definir función de envío para cuando hay dos paquetes. (está lista la idea)

}

func unpaquete(){ //Falta definir función de envío para cuando hay un paquete. (está lista la idea)

}