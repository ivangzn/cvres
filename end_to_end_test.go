package main

import (
	"bytes"
	"os"
	"reflect"
	"testing"

	"github.com/ivangzn/cvres/resume"
	"github.com/ivangzn/cvres/styles"
)

// extractStringFields recursively extracts all string values from a struct
func extractStringFields(v reflect.Value, fields *[]string) {
	if !v.IsValid() {
		return
	}

	switch v.Kind() {
	case reflect.String:
		if str := v.String(); str != "" {
			*fields = append(*fields, str)
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			extractStringFields(v.Field(i), fields)
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			extractStringFields(v.Index(i), fields)
		}
	case reflect.Ptr:
		if !v.IsNil() {
			extractStringFields(v.Elem(), fields)
		}
	case reflect.Interface:
		if !v.IsNil() {
			extractStringFields(v.Elem(), fields)
		}
	}
}

func TestGenerateResumeFromData(t *testing.T) {
	// Create custom resume data structure
	data := resume.Data{
		Person: resume.Person{
			Name: "Person Name",
			Role: "Person Role",
		},
		Contact: resume.Contact{
			Email:    "person@email.com",
			Location: "Location City",
			LinkedIn: "linkedin.com/in/person",
		},
		Sections: []resume.Section{
			{
				Title: "Section 1",
				Articles: []resume.Article{
					{
						Desc: "Article 1 description.",
					},
				},
			},
			{
				Title: "Section 2",
				Articles: []resume.Article{
					{
						What:  "Article 2",
						Where: "Article Place 2",
						When:  "Article Time 2",
						List: []string{
							"List Item 1",
							"List Item 2",
							"List Item 3",
						},
					},
					{
						What:  "Article 3",
						Where: "Article Place 3",
						When:  "Article Time 3",
						List: []string{
							"List Item 4",
							"List Item 5",
							"List Item 6",
						},
					},
					{
						What:  "Article 4",
						Where: "Article Place 4",
						When:  "Article Time 4",
						Desc:  "Article 4 description.",
					},
				},
			},
			{
				Title: "Section 3",
				Articles: []resume.Article{
					{
						What:  "Article 5",
						Where: "Article Place 5",
						When:  "Article Time 5",
					},
				},
			},
			{
				Title: "Section 4",
				Articles: []resume.Article{
					{
						What:  "Article 6",
						Where: "Article Place 6",
					},
					{
						What:  "Article 7",
						Where: "Article Place 7",
					},
					{
						What:  "Article 8",
						Where: "Article Place 8",
					},
				},
			},
			{
				Title: "Section 5",
				Articles: []resume.Article{
					{
						FullList: []string{
							"Skill 1",
							"Skill 2",
							"Skill 3",
							"Skill 4",
							"Skill 5",
							"Skill 6",
							"Skill 7",
							"Skill 8",
							"Skill 9",
						},
					},
				},
			},
		},
	}

	// Extract all string fields using reflection
	var expectedFields []string
	extractStringFields(reflect.ValueOf(data), &expectedFields)

	output := bytes.NewBuffer([]byte{})

	style, err := styles.NewStyle("plain")
	if err != nil {
		t.Fatal(err)
	}

	resume := resume.NewResume(style, data)

	_, err = resume.WriteTo(output)
	if err != nil {
		t.Fatal(err)
	}

	// Get the rendered HTML content
	renderedHTML := output.String()

	// Check that all expected fields are present
	for _, field := range expectedFields {
		if !bytes.Contains([]byte(renderedHTML), []byte(field)) {
			t.Errorf("Expected field '%s' not found in rendered resume", field)
		}
	}

	// Verify the output is not empty
	if len(renderedHTML) == 0 {
		t.Error("Rendered resume is empty")
	}

	// Verify basic HTML structure (should be minified)
	if !bytes.Contains([]byte(renderedHTML), []byte("Person Name")) {
		t.Error("Basic HTML structure verification failed - name not found")
	}
}

func TestGenerateResumeFromFile(t *testing.T) {
	file := "./example/example.json"

	data, err := os.Open(file)
	if err != nil {
		t.Fatal(err)
	}
	defer data.Close()

	output := bytes.NewBuffer([]byte{})

	decoder, err := resume.NewDecoder(data)
	if err != nil {
		t.Fatal(err)
	}

	style, err := styles.NewStyle("plain")
	if err != nil {
		t.Fatal(err)
	}

	resume, err := resume.NewResumeFromDecoder(style, decoder)
	if err != nil {
		t.Fatal(err)
	}

	_, err = resume.WriteTo(output)
	if err != nil {
		t.Fatal(err)
	}
}
