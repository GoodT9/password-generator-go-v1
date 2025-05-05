package alphabet

import "strings"

// Constants for character sets
const (
	UppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	Numbers          = "1234567890"
	Symbols          = "!@#$%^&*()-_=+\/~?"
)

// Alphabet represents a character pool for password generation
type Alphabet struct {
	pool string
}

// New creates a new Alphabet with the specified character sets
func New(uppercaseIncluded, lowercaseIncluded, numbersIncluded, specialCharactersIncluded bool) *Alphabet {
	var builder strings.Builder

	if uppercaseIncluded {
		builder.WriteString(UppercaseLetters)
	}

	if lowercaseIncluded {
		builder.WriteString(LowercaseLetters)
	}

	if numbersIncluded {
		builder.WriteString(Numbers)
	}

	if specialCharactersIncluded {
		builder.WriteString(Symbols)
	}

	return &Alphabet{
		pool: builder.String(),
	}
}

// Get returns the alphabet's character pool
func (a *Alphabet) Get() string {
	return a.pool
}