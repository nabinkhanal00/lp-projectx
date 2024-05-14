package main

import (
	"encoding/json"
	"syscall/js"
)

type Response struct {
	Expression string   `json:"expression"`
	Error      bool     `json:"error"`
	Message    string   `json:"message,omitempty"`
	Answer     string   `json:"answer,omitempty"`
	Steps      []string `json:"steps,omitempty"`
}

func CalculateWASM(this js.Value, args []js.Value) any {
	if len(args) == 0 {
		return nil
	}
	query := args[0].String()
	answer, steps, err := Calculate(query)
	var response Response
	if err != nil {
		response = Response{
			Expression: query,
			Error:      true,
			Message:    err.Error(),
		}
	} else {
		response = Response{
			Expression: query,
			Error:      false,
			Answer:     answer,
			Steps:      steps,
		}
	}
	content, _ := json.Marshal(response)
	return string(content)
}

func main() {
	js.Global().Set("Calculate", js.FuncOf(CalculateWASM))
	<-make(chan bool)
}
