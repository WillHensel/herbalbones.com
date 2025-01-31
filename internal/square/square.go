package square

type Square struct {
	accessToken string
	apiUri      string
	version     string
}

func NewSquare(accessToken string) Square {
	return Square{
		accessToken: accessToken,
		apiUri:      "https://connect.squareup.com/v2",
		version:     "2025-01-23",
	}
}
