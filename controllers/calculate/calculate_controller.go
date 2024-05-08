package calculate

import (
	"encoding/json"
	"net/http"
)

type Request struct {
	Query string `json:"query"`
}

type Response struct {
	Expression string   `json:"expression"`
	Error      bool     `json:"error"`
	Message    string   `json:"message,omitempty"`
	Answer     string   `json:"answer,omitempty"`
	Steps      []string `json:"steps,omitempty"`
}

func CalculateController(w http.ResponseWriter, r *http.Request) {
	request := Request{}
	json.NewDecoder(r.Body).Decode(&request)

	var response Response

	answer, steps, err := calculate(request.Query)
	if err != nil {

		response = Response{
			Expression: request.Query,
			Error:      true,
			Message:    err.Error(),
		}
	} else {

		response = Response{
			Expression: request.Query,
			Error:      false,
			Answer:     answer,
			Steps:      steps,
		}
	}

	content, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(500)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		w.Write(content)
	}
}
