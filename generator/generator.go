package generator

import (
	"math/rand"
	"strings"
	"time"

	"github.com/GoodT9/password-generator-go-v1/alphabet"
	"github.com/GoodT9/password-generator-go-v1/password"
)

// Generator handles password generation and related operations
type Generator struct {
	alphabet *alphabet.Alphabet
}

// New creates a new Generator with the specified character sets
func New(includeUpper, includeLower, includeNum, includeSym bool) *Generator {
	return &Generator{
		alphabet: alphabet.New(includeUpper, includeLower, includeNum, includeSym),
	}
}

// GeneratePassword creates a random password of the specified length
func (g *Generator) GeneratePassword(length int) *password.Password {
	if length < 1 {
		panic("Password length must be at least 1")
	}

	var pass strings.Builder
	alphabetStr := g.alphabet.Get()
	alphabetLength := len(alphabetStr)

	// Initialize random source
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < length; i++ {
		index := r.Intn(alphabetLength)
		pass.WriteByte(alphabetStr[index])
	}

	return password.New(pass.String())
}

// PrintUsefulInfo displays password security tips
func (g *Generator) PrintUsefulInfo() string {
	tips := []string{
		"\n=== Password Security Tips ===",
		"1. Use a minimum password length of 8 or more characters if permitted",
		"2. Include lowercase and uppercase alphabetic characters, numbers and symbols if permitted",
		"3. Generate passwords randomly where feasible",
		"4. Avoid using the same password twice (e.g., across multiple user accounts and/or software systems)",
		"5. Avoid character repetition, keyboard patterns, dictionary words, letter or number sequences,\n   usernames, relative or pet names, romantic links (current or past) and biographical information (e.g., ID numbers, ancestors' names or dates).",
		"6. Avoid using information that the user's colleagues and/or acquaintances might know to be associated with the user",
		"7. Do not use passwords which consist wholly of any simple combination of the aforementioned weak components",
		"8. Use a unique password for each of your important accounts",
		"9. Use a password manager to generate and store complex passwords securely",
		"10. Enable two-factor authentication (2FA) whenever possible for additional security",
		"11. Regularly update your passwords, especially if you suspect they might have been compromised",
		"12. Avoid sharing your passwords with others, even if they claim to be from IT support",
		"13. Do not use the same password for multiple accounts if the accounts are not related to each other",
		"14. Consider using passphrases: long sequences of random words that are easy to remember but hard to crack",
		"15. Set up security questions with answers that are not easily guessable or found on social media",
		"16. If you suspect your password has been compromised, contact your IT support team immediately",
		"==============================",
	}

	return strings.Join(tips, "\n")
}
