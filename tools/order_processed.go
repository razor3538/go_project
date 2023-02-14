package tools

import (
	"encoding/json"
	"example.com/m/config"
	"math"
	"net/http"
)

type Order struct {
	Order   string
	Status  string
	Accrual float64
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

	defer req.Body.Close()
	println(order.Status)
	println(Round(order.Accrual, 4))
	return order.Status, order.Accrual, nil
}

func Round(x float64, prec int) float64 {
	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)
	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow
}
