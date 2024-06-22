package pow

type Transaction struct {
	Sender    string
	Receiver  string
	Amount    float64
	Signature string
}

func (t *Transaction) IsValid() bool {
	// Check for non-empty sender and receiver, positive amount, and potentially signature validation
	return t.Sender != "" && t.Receiver != "" && t.Amount > 0
}
