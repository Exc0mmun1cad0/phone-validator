package tests

import (
	"net/url"
	"phone-validator/internal/handlers"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

const (
	host = "localhost:7777"
)

func TestPhoneValidator_ValidNumbers(t *testing.T) {
	u := url.URL{
		Scheme: "http",
		Host:   host,
	}
	e := httpexpect.Default(t, u.String())

	tests := []struct {
		name           string
		input          string
		wantStatus     bool
		wantNormalized string
	}{
		{
			name:           "Test 1",
			input:          "8 (912) 583-8616",
			wantStatus:     true,
			wantNormalized: "+7-912-583-8616",
		},
		{
			name:           "Test 2",
			input:          "+7 982 583 8616",
			wantStatus:     true,
			wantNormalized: "+7-982-583-8616",
		},
		{
			name:           "Test 3",
			input:          "89865838616",
			wantStatus:     true,
			wantNormalized: "+7-986-583-8616",
		},
	}

	for _, test := range tests {
		e.POST("/validate").
			WithJSON(handlers.Request{PhoneNumber: test.input}).
			Expect().
			Status(200).
			JSON().Object().
			IsEqual(handlers.Response{
				Status:     test.wantStatus,
				Normalized: test.wantNormalized,
			})
	}
}

func TestPhoneValidator_InvalidNumbers(t *testing.T) {
	u := url.URL{
		Scheme: "http",
		Host:   host,
	}
	e := httpexpect.Default(t, u.String())

	tests := []struct {
		name           string
		input          string
		wantStatus     bool
		wantNormalized string
	}{
		{
			name:       "Wrong opcode",
			input:      "8 (933) 583-8616",
			wantStatus: false,
		},
		{
			name:       "Wrong format with +7",
			input:      "8 9345838616",
			wantStatus: false,
		},
		{
			name:       "Wrong format with 8",
			input:      "+7(934)5838616",
			wantStatus: false,
		},
	}

	for _, test := range tests {
		e.POST("/validate").
			WithJSON(handlers.Request{PhoneNumber: test.input}).
			Expect().
			Status(200).
			JSON().Object().
			IsEqual(handlers.Response{Status: test.wantStatus})
	}
}

func TestPhoneValidor_EmptyPhoneNumber(t *testing.T) {
	u := url.URL{
		Scheme: "http",
		Host:   host,
	}
	e := httpexpect.Default(t, u.String())

	e.POST("/validate").
        WithJSON(handlers.Request{PhoneNumber: ""}).
		Expect().
		Status(400)
}

func TestPhoneValidator_WrongMediaType(t *testing.T) {
	u := url.URL{
		Scheme: "http",
		Host:   host,
	}
	e := httpexpect.Default(t, u.String())

	e.POST("/validate").
		WithHeader("Content-Type", "text/plain").
		WithJSON(handlers.Request{PhoneNumber: "89865838616"}).
		Expect().
		Status(415)
}

func TestPhoneValidator_WrongMethod(t *testing.T) {
	u := url.URL{
		Scheme: "http",
		Host:   host,
	}
	e := httpexpect.Default(t, u.String())

	e.PUT("/validate").
		WithJSON(handlers.Request{PhoneNumber: "89865838616"}).
		Expect().
		Status(405)
}
