package password

import (
    "fmt"
    "math"
    "regexp"
    "strings"
)

// Password represents a password with its value and properties
type Password struct {
    Value  string
    Length int
}

// New creates a new Password instance
func New(s string) *Password {
    return &Password{
        Value:  s,
        Length: len(s),
    }
}

// CalculateEntropy calculates the entropy of the password in bits
func (p *Password) CalculateEntropy() float64 {
    var poolSize int

    if regexp.MustCompile(`[A-Z]`).MatchString(p.Value) {
        poolSize += 26
    }
    if regexp.MustCompile(`[a-z]`).MatchString(p.Value) {
        poolSize += 26
    }
    if regexp.MustCompile(`\d`).MatchString(p.Value) {
        poolSize += 10
    }
    if regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\|,.<>\/?]`).MatchString(p.Value) {
        poolSize += 32
    }

    entropy := float64(p.Length) * (math.Log(float64(poolSize)) / math.Log(2))
    return entropy
}

// CharType determines the type of a character
// Returns: 1 for uppercase, 2 for lowercase, 3 for digit, 4 for symbol
func (p *Password) CharType(c rune) int {
    if c >= 'A' && c <= 'Z' {
        return 1
    } else if c >= 'a' && c <= 'z' {
        return 2
    } else if c >= '0' && c <= '9' {
        return 3
    } else {
        return 4
    }
}

// PasswordStrength calculates a strength score for the password
func (p *Password) PasswordStrength() int {
    usedUpper := false
    usedLower := false
    usedNum := false
    usedSym := false
    score := 0

    for _, c := range p.Value {
        charType := p.CharType(c)

        if charType == 1 {
            usedUpper = true
        }
        if charType == 2 {
            usedLower = true
        }
        if charType == 3 {
            usedNum = true
        }
        if charType == 4 {
            usedSym = true
        }
    }

    if usedUpper {
        score++
    }
    if usedLower {
        score++
    }
    if usedNum {
        score++
    }
    if usedSym {
        score++
    }

    if p.Length >= 8 {
        score++
    }
    if p.Length >= 16 {
        score++
    }

    return score
}

// CalculateScore returns a string representation of the password strength
func (p *Password) CalculateScore() string {
    entropy := p.CalculateEntropy()
    visualization := p.VisualizePasswordStrength()

    entropyInfo := fmt.Sprintf("\nPassword Entropy: %.2f bits", entropy)

    if entropy >= 80 {
        return visualization + entropyInfo + "\nThis is a very strong password with high entropy. Great job!"
    } else if entropy >= 60 {
        return visualization + entropyInfo + "\nThis is a strong password with good entropy. You're on the right track!"
    } else if entropy >= 40 {
        return visualization + entropyInfo + "\nThis password has moderate entropy. Consider making it stronger."
    } else {
        return visualization + entropyInfo + "\nThis password has low entropy. It's recommended to choose a stronger password."
    }
}

// VisualizePasswordStrength creates a visual representation of password strength
func (p *Password) VisualizePasswordStrength() string {
    entropy := p.CalculateEntropy()
    var visualization strings.Builder

    visualization.WriteString("[")

    filledBars := int(math.Min(entropy/20, 6)) // Scale entropy to 0-6 range
    for i := 0; i < 6; i++ {
        if i < filledBars {
            visualization.WriteString("█")
        } else {
            visualization.WriteString("░")
        }
    }

    visualization.WriteString("] ")

    if entropy >= 80 {
        visualization.WriteString("Ironclad")
    } else if entropy >= 60 {
        visualization.WriteString("Strong")
    } else if entropy >= 40 {
        visualization.WriteString("Medium")
    } else {
        visualization.WriteString("Weak")
    }

    return visualization.String()
}

// String returns the password value
func (p *Password) String() string {
    return p.Value
}