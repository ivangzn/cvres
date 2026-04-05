package resume

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestNewResume(t *testing.T) {
	t.Parallel()

	style := func(io.Writer, *Data) error { return nil }
	data := Data{Person: Person{Name: "Alice", Role: "Dev"}}
	r := NewResume(style, data)

	if r == nil {
		t.Fatal("expected Resume, got nil")
	}
	if r.Data.Person.Name != "Alice" {
		t.Fatalf("got Name %q, want %q", r.Data.Person.Name, "Alice")
	}
}

type writeToTest struct {
	name    string
	style   StyleFunc
	want    string
	wantErr bool
}

func TestWriteTo(t *testing.T) {
	t.Parallel()

	cases := []writeToTest{
		{
			name: "Style writes HTML",
			style: func(w io.Writer, d *Data) error {
				fmt.Fprintf(w, "<h1>%s</h1>", d.Person.Name)
				return nil
			},
			want: "<h1>Alice</h1>",
		},
		{
			name: "Style returns error",
			style: func(w io.Writer, d *Data) error {
				return errors.New("render failed")
			},
			wantErr: true,
		},
		{
			name: "Style writes nothing",
			style: func(w io.Writer, d *Data) error {
				return nil
			},
			want: "",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			data := Data{Person: Person{Name: "Alice", Role: "Dev"}}
			r := NewResume(tt.style, data)
			var buf bytes.Buffer

			_, err := r.WriteTo(&buf)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got := buf.String(); got != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
		})
	}
}

type newDecoderTest struct {
	name    string
	input   string
	wantErr bool
	isJSON  bool
}

func TestNewDecoder(t *testing.T) {
	t.Parallel()

	cases := []newDecoderTest{
		{
			name:    "JSON input",
			input:   `{"person":{"name":"Alice","role":"Dev"}}`,
			wantErr: false,
		},
		{
			name:    "YAML input",
			input:   "person:\n  name: Alice\n  role: Dev\n",
			wantErr: false,
		},
		{
			name:    "Unknown format",
			input:   "!!!",
			wantErr: true,
		},
		{
			name:    "Empty input",
			input:   "",
			wantErr: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := strings.NewReader(tt.input)
			_, err := NewDecoder(r)
			if tt.wantErr && err == nil {
				t.Fatalf("expected error, but got nil")
			}
		})
	}
}
