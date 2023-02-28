package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	// /代表着传进来任何的URI都执行返回helloword,因为会将pattern进行最长前缀匹配,匹配成功后,执行对应的处理函数.
	http.ListenAndServe(":8080", nil)

}
