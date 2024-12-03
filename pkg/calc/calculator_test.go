package calc_test

import (
	"testing"

	"github.com/m3owmurrr/calc/pkg/calc"
)

func TestCalc(t *testing.T) {
	testsOk := []struct {
		name       string
		expression string
		want       float64
	}{
		{"simple", "2+2", 4.0},
		{"simple", "2*3", 6.0},
		{"simple", "1/2", 0.5},
		{"simple", "1+2+3+4", 10.0},
		{"priority", "2+2*2", 6.0},
		{"priority", "2*2+2", 6.0},
		{"priority", "2/2*2", 2.0},
		{"priority(bracets)", "(2+2)*2", 8.0},
		{"priority(bracets)", "2*(2+2)*2", 16.0},
		{"priority(bracets)", "2*((2/2)+(2-1))*3", 12.0},
	}

	for _, tt := range testsOk {
		t.Run(tt.name, func(t *testing.T) {
			gotAnsw, gotErr := calc.Calc(tt.expression)
			if gotErr != nil {
				t.Errorf("successful case %s returns error: %v", tt.expression, gotErr)
				t.FailNow()
			}
			if gotAnsw != tt.want {
				t.Errorf("Calc(%v) = %v, want %v", tt.expression, gotAnsw, tt.want)
			}
		})
	}

	testsFail := []struct {
		name       string
		expression string
		wantErr    error
	}{
		{"division by Zero", "2/0", calc.ErrDivisionByZero},
		{"division by Zero", "1+(2/(2-2))*10", calc.ErrDivisionByZero},
		{"missing operang", "2+2*", calc.ErrMissingOperand},
		{"missing parentheses", "2+2*4)", calc.ErrMismatchedParentheses},
		{"missing parentheses", "2+(2*4", calc.ErrMismatchedParentheses},
		{"missing parentheses", "2+((2+2+(2+2))", calc.ErrMismatchedParentheses},
		{"unknown operang", "2=2", calc.ErrUnknownOperator},
	}

	for _, tt := range testsFail {
		t.Run(tt.name, func(t *testing.T) {
			_, gotErr := calc.Calc(tt.expression)
			if gotErr == nil {
				t.Errorf("failed case \"%v\" does not return an error", tt.expression)
			}
			if gotErr != tt.wantErr {
				t.Errorf("Calc(%v) = error: %v, want %v", tt.expression, gotErr, tt.wantErr)
			}
		})
	}
}

func BenchmarkCalc(b *testing.B) {
	expression := "2*((2/2)+(2-1))*3"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = calc.Calc(expression)
	}
}

// go test -bench . -benchmem -cpuprofile="cpu.out" -memprofile="mem.out"
