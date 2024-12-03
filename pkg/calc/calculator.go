package calc

import (
	"errors"
	"strconv"
)

var (
	ErrMissingOperand        = errors.New("missing operand in expression")
	ErrMismatchedParentheses = errors.New("missing parentheses in expression")
	ErrDivisionByZero        = errors.New("division by zero")
	ErrUnknownOperator       = errors.New("unknown operator in expression")
	ErrUnknown               = errors.New("this error does not exist :)")
)

func Calc(expression string) (float64, error) {
	if len(expression) == 0 {
		return 0.0, ErrMissingOperand
	}

	if len(expression) == 1 {
		return strconv.ParseFloat(expression, 64)
	}

	// Removing outer parentheses in the expression
	if expression[0] == '(' && expression[len(expression)-1] == ')' {
		counter := 0
		outer := true
		for i := 0; i < len(expression); i++ {
			switch expression[i] {
			case '(':
				counter++
			case ')':
				counter--
			}

			if counter == 0 && i != len(expression)-1 {
				outer = false
				break
			}
		}

		if outer {
			return Calc(expression[1 : len(expression)-1])
		}
	}

	curOpIndex := 0
	parenthesisCounter := 0
	// Choose current opertion by priority
	for i := 0; i < len(expression); i++ {
		switch expression[i] {
		case '(':
			parenthesisCounter++
		case ')':
			parenthesisCounter--
		}

		// If we meet only ')', then the expression is already incorrect
		if parenthesisCounter < 0 {
			return 0.0, ErrMismatchedParentheses
		}

		// Skip the expression in parenthesis
		if parenthesisCounter != 0 {
			continue
		}

		if isOperator(expression[i]) {
			if curOpIndex == -1 || opPriority(expression[curOpIndex]) >= opPriority(expression[i]) {
				curOpIndex = i
			}
		} else if (expression[i] < '0' || expression[i] > '9') && expression[i] != '(' && expression[i] != ')' {
			return 0.0, ErrUnknownOperator
		}

	}

	// If != 0, then there is a parenthesis without a pair
	if parenthesisCounter != 0 {
		return 0.0, ErrMismatchedParentheses
	}

	value1, err1 := Calc(expression[:curOpIndex])
	value2, err2 := Calc(expression[curOpIndex+1:])
	switch {
	case err1 != nil:
		return 0.0, err1
	case err2 != nil:
		return 0.0, err2
	}

	switch expression[curOpIndex] {
	case '-':
		return value1 - value2, nil
	case '+':
		return value1 + value2, nil
	case '/':
		if value2 == 0 {
			return 0.0, ErrDivisionByZero
		}
		return value1 / value2, nil
	case '*':
		return value1 * value2, nil
	default:
		return 0.0, ErrUnknown
	}
}

func opPriority(b byte) int {
	switch {
	case b == '+' || b == '-':
		return 0
	case b == '*' || b == '/':
		return 1
	default:
		return 2
	}
}

func isOperator(b byte) bool {
	return b == '+' || b == '-' || b == '*' || b == '/'
}
