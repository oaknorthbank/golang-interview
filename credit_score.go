package creditscore

import "time"

type InvoiceStatus string

const (
	Paid   InvoiceStatus = "PAID"
	Unpaid InvoiceStatus = "UNPAID"
)

type Invoice struct {
	DueDate time.Time
	Status  InvoiceStatus
}

type CreditReport struct {
	PaymentHistory              []Invoice
	CreditUtilisationPercentage float64
}

type CreditScoreCategory string

const (
	Fair      CreditScoreCategory = "fair"
	Good      CreditScoreCategory = "good"
	Excellent CreditScoreCategory = "excellent"
	Poor      CreditScoreCategory = "poor"
	VeryPoor  CreditScoreCategory = "very poor"
)

type CreditScore struct {
	Value    int
	Category CreditScoreCategory
}

func GetCreditScore(c CreditReport) CreditScore {
	x := c.CreditUtilisationPercentage

	switch {
	case x > 0.9:
		return CreditScore{
			Value:    560,
			Category: VeryPoor,
		}
	case x > 0.7:
		return CreditScore{
			Value:    720,
			Category: Poor,
		}
	case x > 0.5:
		return CreditScore{
			Value:    880,
			Category: Fair,
		}
	case x > 0.3:
		return CreditScore{
			Value:    960,
			Category: Good,
		}
	default:
		return CreditScore{
			Value:    999,
			Category: Excellent,
		}
	}
}

func CalculatePercentage(value, total float64) float64 {
	return float64(int((value/total)*100)) / 100
}
