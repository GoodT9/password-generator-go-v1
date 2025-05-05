package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/GoodT9/password-generator-go/generator"
	"github.com/GoodT9/password-generator-go/password"
)

var (
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			Padding(0, 1).
			Bold(true)

	infoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#2D3748")).
			Padding(0, 1)

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#E53E3E")).
			Padding(0, 1)

	successStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#38A169")).
			Padding(0, 1)

	menuItemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#2D3748")).
			PaddingLeft(4)

	selectedItemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#7D56F4")).
			PaddingLeft(2).
			SetString("> ")
)

type state int

const (
	stateMainMenu state = iota
	stateGeneratePassword
	stateCheckPassword
	stateShowInfo
	stateQuit
	stateAskUppercase
	stateAskLowercase
	stateAskNumbers
	stateAskSymbols
	stateAskLength
	stateShowGeneratedPassword
	stateEnterPasswordToCheck
	stateShowPasswordStrength
)

type model struct {
	state            state
	menuItems        []string
	selectedMenuItem int
	textInput        textinput.Model
	result           string
	error            string
	includeUpper     bool
	includeLower     bool
	includeNum       bool
	includeSym       bool
	passwordLength   int
	generatedPwd     *password.Password
	checkedPwd       *password.Password
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Type here"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 50

	return model{
		state:            stateMainMenu,
		menuItems:        []string{"Generate a new password", "Check password strength", "Display password security information", "Quit"},
		selectedMenuItem: 0,
		textInput:        ti,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEnter:
			switch m.state {
			case stateMainMenu:
				switch m.selectedMenuItem {
				case 0:
					m.state = stateAskUppercase
					m.textInput.SetValue("")
					m.textInput.Placeholder = "Yes/No"
				case 1:
					m.state = stateEnterPasswordToCheck
					m.textInput.SetValue("")
					m.textInput.Placeholder = "Enter your password"
				case 2:
					m.state = stateShowInfo
				case 3:
					m.state = stateQuit
					return m, tea.Quit
				}
			case stateAskUppercase:
				input := strings.ToLower(m.textInput.Value())
				if input == "yes" || input == "y" {
					m.includeUpper = true
				} else if input == "no" || input == "n" {
					m.includeUpper = false
				} else {
					m.error = "Please enter 'yes' or 'no'"
					return m, nil
				}
				m.error = ""
				m.state = stateAskLowercase
				m.textInput.SetValue("")
			case stateAskLowercase:
				input := strings.ToLower(m.textInput.Value())
				if input == "yes" || input == "y" {
					m.includeLower = true
				} else if input == "no" || input == "n" {
					m.includeLower = false
				} else {
					m.error = "Please enter 'yes' or 'no'"
					return m, nil
				}
				m.error = ""
				m.state = stateAskNumbers
				m.textInput.SetValue("")
			case stateAskNumbers:
				input := strings.ToLower(m.textInput.Value())
				if input == "yes" || input == "y" {
					m.includeNum = true
				} else if input == "no" || input == "n" {
					m.includeNum = false
				} else {
					m.error = "Please enter 'yes' or 'no'"
					return m, nil
				}
				m.error = ""
				m.state = stateAskSymbols
				m.textInput.SetValue("")
			case stateAskSymbols:
				input := strings.ToLower(m.textInput.Value())
				if input == "yes" || input == "y" {
					m.includeSym = true
				} else if input == "no" || input == "n" {
					m.includeSym = false
				} else {
					m.error = "Please enter 'yes' or 'no'"
					return m, nil
				}
				
				// Check if at least one character type is selected
				if !m.includeUpper && !m.includeLower && !m.includeNum && !m.includeSym {
					m.error = "You must include at least one character type"
					m.state = stateAskUppercase
					m.textInput.SetValue("")
					return m, nil
				}
				
				m.error = ""
				m.state = stateAskLength
				m.textInput.SetValue("")
				m.textInput.Placeholder = "Enter password length"
			case stateAskLength:
				length, err := strconv.Atoi(m.textInput.Value())
				if err != nil || length < 1 {
					m.error = "Please enter a valid positive number"
					return m, nil
				}
				m.error = ""
				m.passwordLength = length
				
				// Generate the password
				gen := generator.New(m.includeUpper, m.includeLower, m.includeNum, m.includeSym)
				m.generatedPwd = gen.GeneratePassword(m.passwordLength)
				m.result = fmt.Sprintf("Generated password: %s", m.generatedPwd)
				m.state = stateShowGeneratedPassword
			case stateShowGeneratedPassword:
				m.state = stateMainMenu
			case stateEnterPasswordToCheck:
				if m.textInput.Value() == "" {
					m.error = "Password cannot be empty"
					return m, nil
				}
				m.error = ""
				m.checkedPwd = password.New(m.textInput.Value())
				m.result = m.checkedPwd.CalculateScore()
				m.state = stateShowPasswordStrength
			case stateShowPasswordStrength:
				m.state = stateMainMenu
			case stateShowInfo:
				m.state = stateMainMenu
			}
		case tea.KeyUp:
			if m.state == stateMainMenu {
				m.selectedMenuItem = max(0, m.selectedMenuItem-1)
			}
		case tea.KeyDown:
			if m.state == stateMainMenu {
				m.selectedMenuItem = min(len(m.menuItems)-1, m.selectedMenuItem+1)
			}
		case tea.KeyEsc:
			if m.state != stateMainMenu && m.state != stateQuit {
				m.state = stateMainMenu
				m.error = ""
				m.result = ""
			}
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	var s strings.Builder

	// Display banner
	s.WriteString(titleStyle.Render("\n  The Password Thunderdome v3.0.0  \n"))
	s.WriteString("\n")

	switch m.state {
	case stateMainMenu:
		s.WriteString("Welcome to the Thunderdome! How can we be of service today?\n\n")
		for i, item := range m.menuItems {
			if i == m.selectedMenuItem {
				s.WriteString(selectedItemStyle.String() + item + "\n")
			} else {
				s.WriteString(menuItemStyle.Render(item) + "\n")
			}
		}
		s.WriteString("\nUse ↑/↓ to navigate, Enter to select, Ctrl+C to quit\n")
	case stateAskUppercase:
		s.WriteString("Do you want Uppercase letters \"ABCD...\" to be used? (yes/no)\n\n")
		s.WriteString(m.textInput.View())
	case stateAskLowercase:
		s.WriteString("Do you want Lowercase letters \"abcd...\" to be used? (yes/no)\n\n")
		s.WriteString(m.textInput.View())
	case stateAskNumbers:
		s.WriteString("Do you want Numbers \"1234...\" to be used? (yes/no)\n\n")
		s.WriteString(m.textInput.View())
	case stateAskSymbols:
		s.WriteString("Do you want Symbols \"!@#$...\" to be used? (yes/no)\n\n")
		s.WriteString(m.textInput.View())
	case stateAskLength:
		s.WriteString("Great! Now enter the length of the password:\n\n")
		s.WriteString(m.textInput.View())
	case stateShowGeneratedPassword:
		s.WriteString(successStyle.Render("Password Generated Successfully") + "\n\n")
		s.WriteString(m.result + "\n\n")
		s.WriteString("Press Enter to return to the main menu")
	case stateEnterPasswordToCheck:
		s.WriteString("Enter your password to check its strength:\n\n")
		s.WriteString(m.textInput.View())
	case stateShowPasswordStrength:
		s.WriteString(infoStyle.Render("Password Strength Analysis") + "\n\n")
		s.WriteString(m.result + "\n\n")
		s.WriteString("Press Enter to return to the main menu")
	case stateShowInfo:
		s.WriteString(infoStyle.Render("Password Security Tips") + "\n\n")
		gen := generator.New(false, false, false, false)
		s.WriteString(gen.PrintUsefulInfo() + "\n\n")
		s.WriteString("Press Enter to return to the main menu")
	}

	// Display error if any
	if m.error != "" {
		s.WriteString("\n\n" + errorStyle.Render(m.error))
	}

	return s.String()
}

// Helper functions
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
		os.Exit(1)
	}
}