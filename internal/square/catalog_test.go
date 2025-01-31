package square

import (
	"os"
	"testing"
)

func TestCatalogList(t *testing.T) {
	token := os.Getenv("SQUARE_ACCESS_TOKEN")
	sq := NewSquare(token)
	sq.CatalogList()
}
