package tools

import (
	"encoding/json"
	"example.com/m/config"
	"fmt"
	"net/http"
)

type Order struct {
	Order   string
	Status  string
	Accrual float32
}

func OrderProcessed(number string) (string, string, error) {
	var order Order

	url := config.Env.RemoteAPI + "/api/orders/" + number
	//url := "http://localhost:8080/api/orders/" + number

	req, err := http.Get(url)
	if err != nil {
		println(err.Error())
		return "", "0", err
	}

	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()

	err = dec.Decode(&order)

	println("order.Accrual")
	println(order.Accrual)
	println(order.Accrual)
	println(order.Accrual)
	println("order.Accrual")

	if err != nil {
		println(err.Error())
		return "", "0", err
	}

	defer req.Body.Close()
	return order.Status, fmt.Sprintf("%.2f", order.Accrual), nil
}
