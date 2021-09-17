package quotes

import (
	"fmt"
	"math/rand"
)

// Get a random gex quote
func RandomQuote() string {
	quoteLen := len(Quotes)
	return fmt.Sprintf("\"%v\" - gex", Quotes[rand.Intn(quoteLen)])
}
