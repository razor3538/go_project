package tools

import (
	"encoding/json"
	"example.com/m/config"
	"net/http"
)

type Order struct {
	Order   string
	Status  string
	Accrual string
}

func OrderProcessed(number string) (string, float64, error) {
	var order Order

	url := config.Env.RemoteAPI + "/api/orders/" + number
	//url := "http://localhost:8080/api/orders/" + number

	req, err := http.Get(url)
	if err != nil {
		println(err.Error())
		return "", 0.0, err
	}

	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()

	err = dec.Decode(&order)

	if err != nil {
		println(err.Error())

		return "", 0.0, err
	}

	println("order.Accrual")
	println(order.Accrual)
	println(order.Accrual)
	println(order.Accrual)
	println(order.Accrual)
	println(order.Accrual)
	println(order.Accrual)
	println(order.Accrual)
	println("order.Accrual")

	defer req.Body.Close()
	return order.Status, 0, nil
}
