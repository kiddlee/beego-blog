package models

import (
	"github.com/astaxie/beego"
	"github.com/russross/blackfriday"
    "github.com/astaxie/beego/logs"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type Post struct {
	Title   string
	Date    string
	Summary string
	Body    string
	File    string
}

func GetPosts() []Post {
	posts := []Post{}
	files, _ := filepath.Glob(beego.AppConfig.String("content_path") + "posts/*")
    logs.Info(files)
	for _, f := range files {
        file_arr := strings.Split(f, "/")
        file_name := file_arr[len(file_arr) - 1]
		fileread, _ := ioutil.ReadFile(f)
		lines := strings.Split(string(fileread), "\n")
		title := string(lines[0])
		date := string(lines[1])
		summary := string(lines[2])
		body := strings.Join(lines[3:len(lines)], "\n")
		body = string(blackfriday.MarkdownCommon([]byte(body)))
		posts = append(posts, Post{title, date, summary, body, file_name})
	}
    logs.Info(posts)
	return posts
}

func GetPost(file_name string) Post {
	post := Post{}
    file_path := beego.AppConfig.String("content_path") + "posts/" + file_name
	fileread, _ := ioutil.ReadFile(file_path)
	lines := strings.Split(string(fileread), "\n")
	title := string(lines[0])
	date := string(lines[1])
	summary := string(lines[2])
	body := strings.Join(lines[3:len(lines)], "\n")
	body = string(blackfriday.MarkdownCommon([]byte(body)))
	post = Post{title, date, summary, body, file_name}
    return post
}

func GetAboutMe() Post {
	post := Post{}
    file_path := beego.AppConfig.String("content_path") + "aboutme/index.md"
	fileread, _ := ioutil.ReadFile(file_path)
	lines := strings.Split(string(fileread), "\n")
	body := strings.Join(lines, "\n")
	body = string(blackfriday.MarkdownCommon([]byte(body)))
	post = Post{"about me", "5/3/1982", "about me", body, "about me"}
    logs.Info(post)
    return post
}
