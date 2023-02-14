package tools

import (
	"encoding/json"
	"example.com/m/config"
	"net/http"
)

type Order struct {
	Order   string
	Status  string
	Accrual float32
}

func OrderProcessed(number string) (string, float32, error) {
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

	defer req.Body.Close()
	println(order.Status)
	println(order.Accrual)
	return order.Status, order.Accrual, nil
}
