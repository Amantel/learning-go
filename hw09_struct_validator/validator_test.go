package hw09structvalidator

import (
	"encoding/json"
	"errors"
	"testing"
)

type UserRole string

type User struct {
	ID     string          `json:"id" validate:"len:36"`
	Name   string          `validate:"len:36"`
	Age    int             `validate:"min:18|max:50"`
	Email  string          `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
	Role   UserRole        `validate:"in:admin,stuff"`
	Phones []string        `validate:"len:11"`
	meta   json.RawMessage // ignored
}

type App struct {
	Version string `validate:"len:5"`
}

type Token struct {
	Header    []byte
	Payload   []byte
	Signature []byte
}

type Response struct {
	Code int    `validate:"in:200,404,500"`
	Body string `json:"omitempty"`
}

// ---------------- validation-error tests ----------------

func TestValidate(t *testing.T) {
	tests := []struct {
		name      string
		in        interface{}
		wantErr   bool
		wantCodes []error // sentinels expected inside ValidationErrors (order-independent)
	}{
		{
			name: "valid user",
			in: User{
				ID:     "123456789012345678901234567890123456",
				Name:   "validnamevalidnamevalidnamevalidnabb",
				Age:    25,
				Email:  "test@mail.com",
				Role:   "admin",
				Phones: []string{"12345678901"},
			},
			wantErr: false,
		},
		{
			name: "len violation (ID + Name)",
			in: User{
				ID:    "short",
				Name:  "short",
				Age:   25,
				Email: "test@mail.com",
				Role:  "admin",
			},
			wantErr:   true,
			wantCodes: []error{ErrLen, ErrLen},
		},
		{
			name: "min violation",
			in: User{
				ID:    "123456789012345678901234567890123456",
				Name:  "validnamevalidnamevalidnamevalidnabb",
				Age:   10,
				Email: "test@mail.com",
				Role:  "admin",
			},
			wantErr:   true,
			wantCodes: []error{ErrMin},
		},
		{
			name: "max violation",
			in: User{
				ID:    "123456789012345678901234567890123456",
				Name:  "validnamevalidnamevalidnamevalidnabb",
				Age:   60,
				Email: "test@mail.com",
				Role:  "admin",
			},
			wantErr:   true,
			wantCodes: []error{ErrMax},
		},
		{
			name: "regexp violation",
			in: User{
				ID:    "123456789012345678901234567890123456",
				Name:  "validnamevalidnamevalidnamevalidnabb",
				Age:   25,
				Email: "bad-email",
				Role:  "admin",
			},
			wantErr:   true,
			wantCodes: []error{ErrRegexp},
		},
		{
			name: "in violation",
			in: User{
				ID:    "123456789012345678901234567890123456",
				Name:  "validnamevalidnamevalidnamevalidnabb",
				Age:   25,
				Email: "test@mail.com",
				Role:  "unknown",
			},
			wantErr:   true,
			wantCodes: []error{ErrIn},
		},
		{
			name: "slice violation (Phones)",
			in: User{
				ID:     "123456789012345678901234567890123456",
				Name:   "validnamevalidnamevalidnamevalidnabb",
				Age:    25,
				Email:  "test@mail.com",
				Role:   "admin",
				Phones: []string{"short", "12345678901"},
			},
			wantErr:   true,
			wantCodes: []error{ErrLen},
		},
		{
			name:    "response valid",
			in:      Response{Code: 200, Body: "ok"},
			wantErr: false,
		},
		{
			name:      "response invalid in",
			in:        Response{Code: 999},
			wantErr:   true,
			wantCodes: []error{ErrIn},
		},
		{
			name:    "app valid",
			in:      App{Version: "12345"},
			wantErr: false,
		},
		{
			name:      "app invalid len",
			in:        App{Version: "123"},
			wantErr:   true,
			wantCodes: []error{ErrLen},
		},
		{
			name: "token no validation tags",
			in: Token{
				Header:    []byte("a"),
				Payload:   []byte("b"),
				Signature: []byte("c"),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Validate(tt.in)

			if !tt.wantErr {
				if err != nil {
					t.Fatalf("expected nil, got %v", err)
				}
				return
			}

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			var ve ValidationErrors
			if !errors.As(err, &ve) {
				t.Fatalf("expected ValidationErrors, got %T: %v", err, err)
			}

			if len(ve) != len(tt.wantCodes) {
				t.Fatalf("got %d errors, want %d: %v", len(ve), len(tt.wantCodes), ve)
			}

			for _, want := range tt.wantCodes {
				found := false
				for _, got := range ve {
					if errors.Is(got, want) {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("missing sentinel %v in %v", want, ve)
				}
			}
		})
	}
}

// ---------------- program-error tests ----------------

func TestValidateProgramErrors(t *testing.T) {
	type BadLen struct {
		S string `validate:"len:notanint"`
	}
	type BadRegexp struct {
		S string `validate:"regexp:[unclosed"`
	}
	type UnknownRule struct {
		S string `validate:"weird:1"`
	}
	type NoArg struct {
		S string `validate:"len"`
	}
	type Unsupported struct {
		F float64 `validate:"min:1"`
	}

	tests := []struct {
		name    string
		in      interface{}
		wantErr error
	}{
		{"not a struct (int)", 42, ErrNotStruct},
		{"not a struct (nil)", nil, ErrNotStruct},
		{"bad len arg", BadLen{S: "x"}, ErrInvalidRule},
		{"bad regexp", BadRegexp{S: "x"}, ErrInvalidRegexp},
		{"unknown rule", UnknownRule{S: "x"}, ErrInvalidRule},
		{"missing arg", NoArg{S: "x"}, ErrInvalidRule},
		{"unsupported kind", Unsupported{F: 0.5}, ErrUnsupportedFieldKind},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Validate(tt.in)
			if err == nil {
				t.Fatalf("expected program error, got nil")
			}

			var ve ValidationErrors
			if errors.As(err, &ve) {
				t.Fatalf("program error should not be ValidationErrors, got %v", ve)
			}

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected %v, got %v", tt.wantErr, err)
			}
		})
	}
}
