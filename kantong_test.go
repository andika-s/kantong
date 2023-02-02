package kantong

import (
	"encoding/json"
	"errors"
	"log"
	"testing"
)

func TestSuccessItems(t *testing.T) {
}

func TestSuccessItem(t *testing.T) {
}

func TestFailed(t *testing.T) {
	values := []struct {
		handler Handler
		input   error
		output  string
	}{
		{
			handler: Handler{
				Version: "1.0",
				Id:      "52413",
			},
			input:  errors.New("not found"),
			output: `{"version":"1.0","id":"52413","error":"not found"}`,
		},
	}

	t.Run("result-failed-test", func(t *testing.T) {
		for _, value := range values {
			response, err := json.Marshal(value.handler.Result(nil, value.input))
			if err != nil {
				t.Error(err)
			}
			log.Println(string(value.output))
			log.Println(string(response))
			if string(response) != value.output {
				t.Errorf("output not same as expected")
			}
		}
	})
}
