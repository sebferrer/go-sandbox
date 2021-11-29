package article

import "fmt"

type Article struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func (article Article) String() string {
	return fmt.Sprintf("%v\n%v\n%v\n%v\n",
		article.Id, article.Title, article.Desc, article.Content)
}
