package main

import (
	"fmt"
	"webreqverbs/src/request"
)

func main() {
	fmt.Println("-----Handling Web Verbs (Get, Post, PostForm)-----")

	// Web Verbs
	// 1. Get
	// request.Get()
	// 2. Post (Data in JSON Form)
	// request.PostJson()
	// 3. PostForm (Data in urlencoded form)
	request.PostForm()
}
