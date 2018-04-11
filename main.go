package main

import (
	"os"
	"bufio"
	"fmt"
	"net/http"
	"log"
	"strings"
)

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

	fmt.Println(text)

	isExist, existStr := rootTrie.IsExist(text)
	resStr := fmt.Sprintf("%s isExist %t %s", text, isExist, existStr)

	fmt.Println(resStr)
	fmt.Fprintf(w, resStr)
}
