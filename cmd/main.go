package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/wenceslau/sysm-product-calculator/internal/order/infra/database"
	"github.com/wenceslau/sysm-product-calculator/internal/order/usecase"
	"github.com/wenceslau/sysm-product-calculator/pkg/rabbitmq"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	maxWorkers := 10000
	wg := sync.WaitGroup{}

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/orders")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	respository := database.NewOrderRepository(db)
	usecaseOrder := usecase.NewCalculateFinalPriceUseCase(respository)
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	out := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, out)

	wg.Add(maxWorkers)
	for i := 0; i < maxWorkers; i++ {
		defer wg.Done()
		go worker(out, usecaseOrder, i)
	}
	fmt.Println("Waiting for workes to finish")
	wg.Wait()

	// input := usecase.OrderInputDTO{
	// 	ID:    "1234",
	// 	Price: 10.0,
	// 	Tax:   0.1,
	// }

	// output, err := usecaseOrder.Execute(&input)
	// if err != nil {
	// 	panic(err)
	// }

	//fmt.Println(output)

}

func worker(deliveryMessage <-chan amqp.Delivery, uc *usecase.CalculateFinalPriceUseCase, workerId int) {
	for msg := range deliveryMessage {
		//fmt.Printf("Worker %d: %s\n", workerId, msg.Body)
		var input usecase.OrderInputDTO
		err := json.Unmarshal(msg.Body, &input)
		if err != nil {
			fmt.Printf("Error decoding JSON: %s\n", err)
		}
		input.Tax = 0.1
		_, err = uc.Execute(&input)
		if err != nil {
			fmt.Printf("Error executing use case: %s\n", err)
		}

		msg.Ack(false)
	}
}
