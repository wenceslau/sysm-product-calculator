package main

import (
	"database/sql"
	"fmt"

	"github.com/wenceslau/sysm-product-calculator/internal/order/infra/database"
	"github.com/wenceslau/sysm-product-calculator/internal/order/usecase"
)

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/orders")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	respository := database.NewOrderRepository(db)
	usecaseOrder := usecase.NewCalculateFinalPriceUseCase(respository)

	input := usecase.OrderInputDTO{
		ID:    "1234",
		Price: 10.0,
		Tax:   0.1,
	}

	output, err := usecaseOrder.Execute(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)

}
