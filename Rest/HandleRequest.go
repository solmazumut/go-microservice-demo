package Rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type RestObject struct {
	deneme string
}

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Title", Desc: "Test Description", Content: "Test Content"},
	}
	fmt.Println("EndPoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}
func testPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	http.HandleFunc("/test", testPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (rest *RestObject) GetRequest(text string) {
	rest.deneme = text
	handleRequests()
}
