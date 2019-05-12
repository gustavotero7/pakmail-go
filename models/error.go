package models

import "fmt"

// Error represents pakmail error
type Error struct {
	Errors interface{} `json:"errors"`
}

func (er *Error) Error() string {
	return fmt.Sprintf("%v", er.Errors)
}
