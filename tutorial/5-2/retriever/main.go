package main

import (
	"fmt"
	"golearn/tutorial/5-2/retriever/mock"
	"golearn/tutorial/5-2/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

const url = "http://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name":   "ccmouse",
			"course": "golang",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url,
		map[string]string{
			"contents": "another faked imooc.com",
		})
	return s.Get(url)
}

func main() {

	//io.ReadWriteCloser()

	var r Retriever

	//r = mock.Retriever{"this is a fake imooc.com"}
	retriever := mock.Retriever{"this is a fake imooc.com"}
	r = &retriever
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)

	// Type assertion，类型断言
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)

	//mockRetriever := r.(mock.Retriever)
	//fmt.Println(mockRetriever.Contents)
	// panic: interface conversion: main.Retriever is *real.Retriever, not mock.Retriever

	// 未匹配到类型断言，进行错误处理
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	fmt.Println("Try a session")
	fmt.Println(session(&retriever))

	//fmt.Println(download(r))
	/*fmt.Println(download(
	mock.Retriever{
		"this is a fake imooc.com"}))*/
}

// 检查实现者类型的方法
func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	fmt.Println("Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
