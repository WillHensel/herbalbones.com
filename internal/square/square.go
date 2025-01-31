package square

type Square struct {
	accessToken string
}

func NewSquare(accessToken string) Square {
	return Square{accessToken: accessToken}
}
