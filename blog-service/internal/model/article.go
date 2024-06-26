package model

import "github.com/go-programming-tour-book/blog-service/pkg/app"

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (t Article) TableName() string {
	return "blog_article"
}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}
