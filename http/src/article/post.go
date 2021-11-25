package article

type Post struct {
	UserId string `json:"userId"`
	Id     string `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
