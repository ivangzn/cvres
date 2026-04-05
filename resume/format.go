package resume

import "unicode"

type fileKind int

const (
	fileJSON fileKind = iota
	fileYAML
	fileUnknown
)

// guessFormat identifies if the given data is JSON or YAML
// by looking at the first non-whitespace character.
func guessFormat(data []byte) fileKind {
	for _, b := range data {
		char := rune(b)
		if unicode.IsSpace(char) {
			continue
		}

		if char == '{' || char == '[' {
			return fileJSON
		}

		if unicode.IsLetter(char) {
			return fileYAML
		}

		return fileUnknown
	}
	return fileUnknown
}
