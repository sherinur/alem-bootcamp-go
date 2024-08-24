package firststruct

// Password complexity checks
func (u *User) PasswordReliability() string {
	if u.password == "" {
		return "undefined"
	}

	var (
		hasSpecial  bool
		lengthValid bool
		hasUpper    bool
		hasLower    bool
		hasDigit    bool
		count       int
	)

	// Length of at least 8 characters
	if len(u.password) >= 8 {
		lengthValid = true
	}

	// character checks
	for _, char := range u.password {
		switch {
		case char >= 'A' && char <= 'Z':
			hasUpper = true
		case char >= 'a' && char <= 'z':
			hasLower = true
		case char >= '0' && char <= '9':
			hasDigit = true
		default:
			hasSpecial = true
		}
	}

	if lengthValid {
		count++
	}

	if hasSpecial {
		count++
	}

	if hasDigit {
		count++
	}

	if hasUpper {
		count++
	}

	if hasLower {
		count++
	}

	switch {
	case count == 5:
		return "strong"
	case count >= 3:
		return "medium"
	case count >= 1:
		return "easy"
	default:
		return "undefined"
	}
}
