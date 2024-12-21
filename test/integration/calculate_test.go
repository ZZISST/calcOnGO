package integration

import (
	"calcOnGO/internal/service"
	"testing"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		want       float64
		expectErr  bool
	}{
		{"Simple addition", "2+2", 4, false},
		{"Multiplication and addition", "2+2*2", 6, false},
		{"Division by zero", "10/0", 0, true},
		{"Invalid characters", "2+2a", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.Calc(tt.expression)
			if (err != nil) != tt.expectErr {
				t.Errorf("Calc() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if got != tt.want && !tt.expectErr {
				t.Errorf("Calc() got = %v, want %v", got, tt.want)
			}
		})
	}
}
