package utils

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// StringArray is a custom type for handling array of strings in DB (e.g. JSON)
type StringArray []string

// Scan implements the Scanner interface for StringArray
func (s *StringArray) Scan(value interface{}) error {
	if value == nil {
		*s = StringArray{}
		return nil
	}

	// Convert the value to a byte slice if needed
	b, ok := value.([]byte)
	if !ok {
		return errors.New("failed to scan StringArray")
	}

	// Unmarshal the byte slice into a StringArray
	return json.Unmarshal(b, &s)
}

// Value implements the Valuer interface for StringArray
func (s StringArray) Value() (driver.Value, error) {
	// Marshal the StringArray into JSON and return the value
	return json.Marshal(s)
}
