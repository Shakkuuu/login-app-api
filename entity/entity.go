package entity

// User
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Ticket   int    `json:"ticket"`
	Coin     int    `json:"coin"`
}

// Memo
type Memo struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Title string `json:"title"`
	Text  string `json:"text"`
}
