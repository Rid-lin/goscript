package main

import (
	"database/sql"

	_ "github.com/lib/pq" //..
)

type Store struct {
	db *sql.DB
}

type Article struct {
	ID           int
	Magazines    string
	ArticleTypes string
	Author       string
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func NewDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func (a *App) GetAllArticles() ([]*Article, error) {
	var ars []*Article

	rows, err := a.store.db.Query("SELECT public.articles.id, public.magazines.name AS magazines, public.article_types.type AS article_types, public.author.author FROM magazines, article_types, author, articles WHERE articles.magazines_id = magazines.Id AND articles.article_type_id = article_types.id AND articles.author_id = author.id;")
	if err != nil {
		return ars, err
	}

	defer rows.Close()
	for rows.Next() {
		ar := Article{}
		err = rows.Scan(&ar.ID, &ar.Magazines, &ar.ArticleTypes, &ar.Author)
		if err != nil {
			continue
		}

		ars = append(ars, &ar)
	}

	return ars, nil
}
