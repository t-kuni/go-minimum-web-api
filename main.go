package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Response struct {
	RequestURI     string              `json:"request_uri"`
	RequestMethod  string              `json:"request_method"`
	Headers        map[string][]string `json:"headers"`
	Body           string              `json:"body"`
	EnvironmentVar map[string]string   `json:"environment_variables"`
}

func main() {
	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	fmt.Printf("Server started on port %s\n", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Server stopped with error: %v\n", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	// リクエストヘッダーの取得
	headers := r.Header

	// リクエストボディの取得
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read the request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	body := string(bodyBytes)

	// 環境変数の取得
	environmentVars := make(map[string]string)
	for _, e := range os.Environ() {
		pair := splitEnv(e)
		environmentVars[pair[0]] = pair[1]
	}

	// レスポンスの構築
	response := &Response{
		RequestURI:     r.RequestURI,
		RequestMethod:  r.Method,
		Headers:        headers,
		Body:           body,
		EnvironmentVar: environmentVars,
	}

	// JSONレスポンスの内容を変数として取得
	responseData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to marshal the response", http.StatusInternalServerError)
		return
	}

	// JSONレスポンスの内容を標準出力に出力
	fmt.Println(string(responseData))

	// クライアントへのレスポンスの書き出し
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseData)
}

func splitEnv(envVar string) []string {
	for i := 0; i < len(envVar); i++ {
		if envVar[i] == '=' {
			return []string{envVar[:i], envVar[i+1:]}
		}
	}
	return []string{envVar, ""}
}
