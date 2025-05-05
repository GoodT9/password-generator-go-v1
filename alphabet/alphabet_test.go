package alphabet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAlphabet(t *testing.T) {
	tests := []struct {
		name                    string
		uppercaseIncluded       bool
		lowercaseIncluded       bool
		numbersIncluded         bool
		specialCharsIncluded    bool
		expectedContainsUpper   bool
		expectedContainsLower   bool
		expectedContainsNumbers bool
		expectedContainsSpecial bool
	}{
		{
			name:                    "All character types included",
			uppercaseIncluded:       true,
			lowercaseIncluded:       true,
			numbersIncluded:         true,
			specialCharsIncluded:    true,
			expectedContainsUpper:   true,
			expectedContainsLower:   true,
			expectedContainsNumbers: true,
			expectedContainsSpecial: true,
		},
		{
			name:                    "Only uppercase included",
			uppercaseIncluded:       true,
			lowercaseIncluded:       false,
			numbersIncluded:         false,
			specialCharsIncluded:    false,
			expectedContainsUpper:   true,
			expectedContainsLower:   false,
			expectedContainsNumbers: false,
			expectedContainsSpecial: false,
		},
		{
			name:                    "No characters included",
			uppercaseIncluded:       false,
			lowercaseIncluded:       false,
			numbersIncluded:         false,
			specialCharsIncluded:    false,
			expectedContainsUpper:   false,
			expectedContainsLower:   false,
			expectedContainsNumbers: false,
			expectedContainsSpecial: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := New(tt.uppercaseIncluded, tt.lowercaseIncluded, tt.numbersIncluded, tt.specialCharsIncluded)
			
			// Check if the alphabet contains the expected character types
			containsUpper := containsAny(a.Get(), UppercaseLetters)
			containsLower := containsAny(a.Get(), LowercaseLetters)
			containsNumbers := containsAny(a.Get(), Numbers)
			containsSpecial := containsAny(a.Get(), Symbols)
			
			assert.Equal(t, tt.expectedContainsUpper, containsUpper)
			assert.Equal(t, tt.expectedContainsLower, containsLower)
			assert.Equal(t, tt.expectedContainsNumbers, containsNumbers)
			assert.Equal(t, tt.expectedContainsSpecial, containsSpecial)
		})
	}
}

func containsAny(s, chars string) bool {
	for _, c := range chars {
		for _, sc := range s {
			if c == sc {
				return true
			}
		}
	}
	return false
}