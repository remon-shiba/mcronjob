package service

import (
	"encoding/json"
	"fmt"
	"time"
)

type (
	SampleTime struct {
		TimeZero  time.Time `json:"timeZero,omitzero"`
		TimeEmpty time.Time `json:"timeEmpty,omitempty"`
	}
)

func CheckTime() {
	timeData, _ := json.Marshal(&SampleTime{})
	fmt.Printf("RESULT: %s", string(timeData))
}
