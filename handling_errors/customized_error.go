package handling_errors

import "fmt"

type CommsError struct{}

func (m CommsError) Error() string {
	return "An error happened during data transfer."
}

type SyntaxError struct {
	Line int
	Col  int
}

func (err *SyntaxError) Error() string {
	return fmt.Sprintf("Error at line %d, column %d", err.Line, err.Col)
}
