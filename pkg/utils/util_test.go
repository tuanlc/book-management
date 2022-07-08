package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Foo struct {
	Bar string
}

func TestUtil(t *testing.T) {
	t.Run("should return parsed body", func(t *testing.T) {
		body := new(bytes.Buffer)
		objectToSerializeToJSON := Foo{Bar: "test"}
		json.NewEncoder(body).Encode(objectToSerializeToJSON)

		r := httptest.NewRequest(http.MethodPost, "/", body)

		foo := &Foo{}

		ParseBody(r, foo)

		if foo.Bar != "test" {
			t.Errorf("got %q want %q", foo.Bar, "test")
		}
	})

	t.Run("should parsed body if the request body is nil", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodPost, "/", nil)

		foo := &Foo{}

		ParseBody(r, foo)

		if foo.Bar != "" {
			t.Errorf("got %q want %q", foo.Bar, "")
		}
	})
}
