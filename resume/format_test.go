package resume

import "testing"

type guessTest struct {
	name string
	text string
	want fileKind
}

func TestGuessFormat(t *testing.T) {
	t.Parallel()

	cases := []guessTest{
		{
			name: "Empty content",
			text: "",
			want: fileUnknown,
		},
		{
			name: "JSON object",
			text: `{"name": "John", "age": 30}`,
			want: fileJSON,
		},
		{
			name: "JSON array",
			text: `[1, 2, 3]`,
			want: fileJSON,
		},
		{
			name: "JSON with leading whitespace",
			text: "   \n\t{\"key\": \"value\"}",
			want: fileJSON,
		},
		{
			name: "YAML content",
			text: "name: John\nage: 30\n",
			want: fileYAML,
		},
		{
			name: "YAML with leading whitespace",
			text: "  \n  name: value",
			want: fileYAML,
		},
		{
			name: "Whitespace only",
			text: "   \t\n  ",
			want: fileUnknown,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := guessFormat([]byte(tt.text))
			if got != tt.want {
				t.Fatalf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
