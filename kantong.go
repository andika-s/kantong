package kantong

import (
	"reflect"
)

type (
	Kantong interface {
		Result(any, error) *Handler
	}

	Handler struct {
		Version string `json:"version"`
		Id      string `json:"id"`
		Data    *data  `json:"data,omitempty"`
		Error   string `json:"error,omitempty"`
	}

	data struct {
		Item  any `json:"item,omitempty"`
		Items any `json:"items,omitempty"`
	}
)

func (h *Handler) Result(result any, err error) *Handler {
	if err != nil {
		return h.failed(err)
	}
	return h.success(result)
}

func (h *Handler) failed(err error) *Handler {
	return &Handler{
		Version: h.Version,
		Id:      h.Id,
		Error:   err.Error(),
	}
}

func isList(result any) bool {
	return reflect.TypeOf(result).Kind() == reflect.Array
}

func (h *Handler) success(result any) *Handler {
	if isList(result) {
		length := len(result.(map[any]any))
		return h.items(result, length)
	}
	return h.item(result)
}

func (h *Handler) items(result any, length int) *Handler {
	return &Handler{
		Version: h.Version,
		Id:      h.Id,
		Data: &data{
			Items: append(make([]any, length), result),
		},
	}
}

func (h *Handler) item(result any) *Handler {
	return &Handler{
		Version: h.Version,
		Id:      h.Id,
		Data: &data{
			Item: result,
		},
	}
}
