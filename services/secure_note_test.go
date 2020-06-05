package services

import (
	"encoding/base64"
	"ncrypt-cli/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateNote(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusCreated)
		_, _ = rw.Write([]byte(
			`{"Code":201,"Message":"Note stored.","Data":{"id":"209c64c9-0e01-42e3-ad2a-d5bda2713262","url":"https://ncrypt.site/note/209c64c9-0e01-42e3-ad2a-d5bda2713262"},"Error":null,"Meta":null}`,
		),
		)
	}))
	defer server.Close()

	hs := HttpService{
		Client: server.Client(),
		Url:    server.URL,
	}

	note, key, err := CreateNote(hs, models.Note{
		Note:                 []byte("Denial is the most predictable of all human responses."),
		SelfDestruct:         1,
		DestructAfterOpening: true,
	})

	if err != nil {
		t.Fatal(err)
	}

	if note.Code != http.StatusCreated {
		t.Fail()
	}

	k, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		t.Fatal(err)
	}

	if len(k) != 32 {
		t.Fail()
	}
}

func TestCreateNoteWithServerFailure(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusInternalServerError)
		_, _ = rw.Write([]byte(
			`{"Code":500,"Message":"Internal error","Data":null,"Error":null,"Meta":null}`,
		),
		)
	}))
	defer server.Close()

	hs := HttpService{
		Client: server.Client(),
		Url:    server.URL,
	}

	_, _, err := CreateNote(hs, models.Note{
		Note: []byte("There’s no escaping reason, no denying purpose. " +
			"Because as we both know, without purpose, we would not exist."),
		SelfDestruct:         1,
		DestructAfterOpening: true,
	})

	if err == nil {
		t.Fail()
	}
}

func TestCreateNoteWithInvalidJson(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusCreated)
		_, _ = rw.Write([]byte(
			`{"Code":201,"Message":"Note stored.","Data":{"id":"209c64c9-0e01-42e3-ad2a-d5bda2713262","url":"https://ncrypt.site/note/209c64c9-0e01-42e3-ad2a-d5bda2713262"},"Error":null,"Meta"`,
		),
		)
	}))
	defer server.Close()

	hs := HttpService{
		Client: server.Client(),
		Url:    server.URL,
	}

	_, _, err := CreateNote(hs, models.Note{
		Note: []byte("You have a problem with authority, Mr. Anderson. " +
			"You believe you are special, that somehow the rules do not apply to you. Obviously, you are mistaken."),
		SelfDestruct:         1,
		DestructAfterOpening: true,
	})

	if err == nil || err.Error() != "unexpected end of JSON input" {
		t.Fail()
	}
}
