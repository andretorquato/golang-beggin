package controllers

import "net/http"

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreatePost"))
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetAllPosts"))
}

func GetAnPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetAnPost"))
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdatePost"))
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeletePost"))
}
