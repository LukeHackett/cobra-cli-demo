package service

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

type CrudService interface {
	FetchAll() string
	FetchById(id int) []byte
}

type UserService struct {
	BaseUrl string
}

func (service UserService) FetchAll() string {
	slog.Debug("Making a request to: " + service.BaseUrl + "/users")
	resp, err := http.Get(service.BaseUrl + "/users")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {

	}

	var out bytes.Buffer
	t := json.Indent(&out, body, "", "  ")
	if t != nil {

	}

	return out.String()
}

func (service UserService) FetchById(id int) []byte {
	return []byte{}
}

// func main() {
// 	svc := UserService{baseUrl: "http://example.com"}

// 	svc.FetchAll()

// }
