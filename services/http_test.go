package services

import (
	"ncrypt-cli/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpService_SendRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusCreated)
		_, _ = rw.Write([]byte(
			`{"Code":201,"Message":"Note stored.","Data":{"id":"209c64c9-0e01-42e3-ad2a-d5bda2713262","url":""},"Error":null,"Meta":null}`,
		),
		)
	}))
	defer server.Close()

	hs := HttpService{
		Client: server.Client(),
		Url:    server.URL,
	}

	requestModel := models.SecureMessageRequest{
		Note:                 "an encrypted message",
		SelfDestruct:         0,
		DestructAfterOpening: false,
	}

	_, err := hs.SendRequest(requestModel)
	if err != nil {
		t.Fatal(err)
	}
}

func TestHttpService_SendRequest_failure(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusInternalServerError)
		_, _ = rw.Write([]byte(
			`{"Code":201,"Message":"Note stored.","Data":{"id":"209c64c9-0e01-42e3-ad2a-d5bda2713262","url":""},"Error":null,"Meta":null}`,
		),
		)
	}))
	defer server.Close()

	hs := HttpService{
		Client: server.Client(),
		Url:    server.URL,
	}

	requestModel := models.SecureMessageRequest{
		Note:                 "an encrypted message",
		SelfDestruct:         0,
		DestructAfterOpening: false,
	}

	_, err := hs.SendRequest(requestModel)
	if err == nil {
		t.Fail()
	}
}
