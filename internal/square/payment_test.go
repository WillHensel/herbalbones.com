package square

import (
	"os"
	"testing"
)

func TestPaymentLink(t *testing.T) {
	token := os.Getenv("SQUARE_ACCESS_TOKEN")
	sq := NewSquare(token)
	t.Logf(sq.GetSingleItemPaymentLink("CE2ZLR6QK2B2EJGHUW3A54FR"))
}
