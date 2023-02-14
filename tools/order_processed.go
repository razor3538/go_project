package tools

import (
	"encoding/json"
	"example.com/m/config"
	"net/http"
)

type Order struct {
	Order   string
	Status  string
	Accrual int
}

func OrderProcessed(number string) (string, error) {
	var order Order

	url := config.Env.RemoteAPI + "//" + "/api/orders/" + number

	req, err := http.Get(url)
	if err != nil {
		return "", err
	}

	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()

	err = dec.Decode(&order)

	if err != nil {
		return "", err
	}

	defer req.Body.Close()
	return order.Status, nil
}
