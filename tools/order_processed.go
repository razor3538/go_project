package tools

import (
	"encoding/json"
	"example.com/m/config"
	"fmt"
	"net/http"
	"strconv"
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

	tmp := fmt.Sprintf("%.2f", order.Accrual)

	accrualRes, _ := strconv.ParseFloat(tmp, 64)

	println("order.Accrual")
	println(accrualRes)
	println(accrualRes)
	println(accrualRes)
	println("order.Accrual")

	if err != nil {
		println(err.Error())
		return "", 0.0, err
	}

	defer req.Body.Close()
	return order.Status, accrualRes, nil
}
