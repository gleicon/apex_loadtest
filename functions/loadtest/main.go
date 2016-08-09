package main

import (
	"encoding/json"
	"time"

	"github.com/apex/go-apex"
	"github.com/tsenart/vegeta/lib"
)

type message struct {
	Value string `json:"value"`
}

type ReturnMessage struct {
	Value interface{} `json:"value"`
}

func loadtest(url string) (*vegeta.Metrics, error) {
	rate := uint64(100) // per second
	duration := 5 * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    url,
	})
	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration) {
		metrics.Add(res)
	}
	metrics.Close()

	return &metrics, nil
}

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		var r ReturnMessage
		var m message
		//	var j []byte
		var err error

		r = ReturnMessage{}

		if err = json.Unmarshal(event, &m); err != nil {
			return nil, err
		}

		if r.Value, err = loadtest(m.Value); err != nil {
			return nil, err
		}

		//		if j, err = json.Marshal(r); err != nil {
		//			return nil, err
		//		}

		return r, nil
	})
}
