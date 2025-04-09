package creditscore

import (
	"math"
	"time"
)

type InvoiceStatus int

const (
	Paid InvoiceStatus = iota
	Unpaid
)

type Invoice struct {
	DueDate time.Time
	Status  InvoiceStatus
}

type CreditReport struct {
	PaymentHistory              []*Invoice
	CreditUtilisationPercentage float64
}

type CreditScoreCategory int

const (
	Fair CreditScoreCategory = iota
	Good
	Excellent
	Poor
	VeryPoor
)

type CreditScore struct {
	Value    int
	Category CreditScoreCategory
}

func GetCreditScore(c *CreditReport) *CreditScore {
	x := c.CreditUtilisationPercentage

	switch {
	case x > 0.9:
		return &CreditScore{
			Value:    560,
			Category: VeryPoor,
		}
	case x > 0.7:
		return &CreditScore{
			Value:    720,
			Category: Poor,
		}
	case x > 0.5:
		return &CreditScore{
			Value:    880,
			Category: Fair,
		}
	case x > 0.3:
		return &CreditScore{
			Value:    960,
			Category: Good,
		}
	default:
		return &CreditScore{
			Value:    999,
			Category: Excellent,
		}
	}
}

func CalculatePercentage(value, total float64) float64 {
	return math.Round(value/total*100) / 100
}
