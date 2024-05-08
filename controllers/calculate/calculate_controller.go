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

	response := Response{
		Expression: request.Query,
		Error:      false,
		Message:    "",
		Answer:     "55",
		Steps:      []string{"Hello", "How are you"},
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
