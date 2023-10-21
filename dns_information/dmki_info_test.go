package dnsinformation

import "testing"

func TestHasDKIMRecord(t *testing.T) {
	// Test cases for hasDKIMRecord function
	t.Run("Valid DKIM record", func(t *testing.T) {
		record := "v=DKIM1alidABCDEFG"
		result := hasDKIMRecord(record)
		if !result {
			t.Errorf("Expected true for a valid DKIM record, got false")
		}
	})

	t.Run("Short DKIM record", func(t *testing.T) {
		record := "vDKIM1dsdds"
		result := hasDKIMRecord(record)
		if !result {
			t.Errorf("Expected true for a short DKIM record, got false")
		}
	})

	t.Run("Invalid DKIM record", func(t *testing.T) {
		record := "InvalidDKIM"
		result := hasDKIMRecord(record)
		if result {
			t.Errorf("Expected false for an invalid DKIM record, got true")
		}
	})

	t.Run("Empty string", func(t *testing.T) {
		record := ""
		result := hasDKIMRecord(record)
		if result {
			t.Errorf("Expected false for an empty string, got true")
		}
	})
}

func TestIsValidDKIM(t *testing.T) {
	// Test cases for isValidDKIM function
	t.Run("Valid DKIM record", func(t *testing.T) {
		record := "v=DKIM1ValidABCDEFG"
		result := isValidDKIM(record)
		if !result {
			t.Errorf("Expected true for a valid DKIM record, got false")
		}
	})

	t.Run("Short DKIM record", func(t *testing.T) {
		record := "DKIMVal"
		result := isValidDKIM(record)
		if result {
			t.Errorf("Expected true for a short DKIM record, got false")
		}
	})

	t.Run("Invalid DKIM record", func(t *testing.T) {
		record := "InvalidDKIM"
		result := isValidDKIM(record)
		if result {
			t.Errorf("Expected false for an invalid DKIM record, got true")
		}
	})

	t.Run("Empty string", func(t *testing.T) {
		record := ""
		result := isValidDKIM(record)
		if result {
			t.Errorf("Expected false for an empty string, got true")
		}
	})
}

func TestGetDKIM(t *testing.T) {
	t.Run("Valid DKIM Record", func(t *testing.T) {
		domain := "example.com"
		selector := "valid"
		_, err := getDKIM(domain, selector)
		if err != nil {
			t.Errorf("Expected no error, but got an error: %v", err)
		}
	})

	t.Run("No DKIM Record", func(t *testing.T) {
		domain := "example.com"
		selector := "nodkim"
		record, err := getDKIM(domain, selector)
		if err != nil {
			t.Errorf("Expected no error, but got an error: %v", err)
		}
		if record != "" {
			t.Error("Expected an empty string, but got a DKIM record")
		}
	})
}
