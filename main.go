package main

import (
	"os"
	"bufio"
	"fmt"
	"net/http"
	"log"
	"strings"
	"encoding/json"
)

type ResJson struct {
	Code       int32  `json:"code"`
	IsExist    bool   `json:"isExist"`
	OriginText string `json:"originText"`
	Reason     string `json:"reason"`
}

var rootTrie *TrieNode

func init() {
	root := TrieNode{nil, false}
	inFile, _ := os.Open("/Users/admin/nickname_uniq.txt")
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	lineNum := 0
	for scanner.Scan() {
		lineStr := scanner.Text()
		if len(lineStr) > 6 {
			lineNum++
			root.AddWord(lineStr)
		}
		if lineNum%100000 == 0 {
			fmt.Println(lineStr, lineNum)
			PrintMem()
		}
	}

	rootTrie = &root
	fmt.Println("######### finish init #########")
	PrintMem()
}

func main() {
	http.HandleFunc("/", search)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func search(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	text := strings.Join(r.Form["text"], "")

	isExist, existStr := rootTrie.IsExist(text)

	resJson := ResJson{
		0,
		isExist,
		text,
		existStr,
	}

	resStr, err := json.Marshal(resJson)
	if err != nil {
		fmt.Println("json marshal err:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {

		w.Header().Set("Content-Type", "application/json")
		w.Write(resStr)
	}
}
