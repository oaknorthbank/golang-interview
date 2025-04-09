package creditscore

import (
	"testing"
)

func TestGetCreditScore(t *testing.T) {
	tests := []struct {
		name              string
		creditUtilisation float64
		expectedScore     int
		expectedCategory  CreditScoreCategory
	}{
		{
			name:              "returns 560 for credit utilisation > 90%",
			creditUtilisation: 0.95,
			expectedScore:     560,
			expectedCategory:  VeryPoor,
		},
		{
			name:              "returns 720 for credit utilisation between 70-90%",
			creditUtilisation: 0.8,
			expectedScore:     720,
			expectedCategory:  Poor,
		},
		{
			name:              "returns 880 for credit utilisation between 50-70%",
			creditUtilisation: 0.6,
			expectedScore:     880,
			expectedCategory:  Fair,
		},
		{
			name:              "returns 960 for credit utilisation between 30-50%",
			creditUtilisation: 0.4,
			expectedScore:     960,
			expectedCategory:  Good,
		},
		{
			name:              "returns 999 for credit utilisation < 30%",
			creditUtilisation: 0.2,
			expectedScore:     999,
			expectedCategory:  Excellent,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			creditReport := &CreditReport{
				PaymentHistory:              []*Invoice{},
				CreditUtilisationPercentage: tt.creditUtilisation,
			}

			got := GetCreditScore(creditReport)

			// do we want to keep using vanilla go testing?
			// most often we use a testing library like testify
			// https://pkg.go.dev/github.com/stretchr/testify/assert
			if got.Value != tt.expectedScore {
				t.Errorf("expected score %d, got %d", tt.expectedScore, got.Value)
			}
			if got.Category != tt.expectedCategory {
				t.Errorf("expected category %v, got %v", tt.expectedCategory, got.Category)
			}
		})
	}
}
