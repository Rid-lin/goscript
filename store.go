package main

import (
	"os"
	"path"
	"text/template"
)

type displayData struct {
	Style  string
	Header string
	Table  []*Article
	Author string
	Email  string
}

func (a *App) SaveArticlesToFile(articles []*Article, pathToWeb string) error {
	styleCss, err := a.fs.ReadFile("assets/style.css")
	if err != nil {
		return err
	}
	DisplayData := displayData{
		Style:  string(styleCss),
		Header: "Articles",
		Table:  articles,
		Author: "Vladislav Vegner",
		Email:  "vlad@vegner.org",
	}
	ts, err := template.ParseFS(a.fs, "assets/index.html")
	// ts := template.Must(template.New("index.html").ParseFiles("assets/index.html"))
	if err != nil {
		return err
	}

	f, err := os.Create(path.Join(pathToWeb, "web.html"))
	if err != nil {
		return err
	}
	defer f.Close()

	// err = ts.Execute(f, DisplayData)
	err = ts.ExecuteTemplate(f, "index.html", DisplayData)
	if err != nil {
		return err
	}

	// err = ts.Execute(os.Stdout, DisplayData)
	// err = ts.ExecuteTemplate(os.Stdout, "index.html", DisplayData)
	// if err != nil {
	// 	return err
	// }

	return nil
}
