package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type ResJson struct {
	Code       int32  `json:"code"`
	IsExist    bool   `json:"isExist"`
	OriginText string `json:"originText"`
	Reason     string `json:"reason"`
}

var rootTrie *TrieNode

func init() {
	defer Destroy()

	root := TrieNode{nil, false}
	inFile, err := os.Open(Conf.WordsPath)
	if err != nil {
		Log.Error("read words fail")
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	lineNum := 0
	for scanner.Scan() {
		lineStr := scanner.Text()

		root.AddWord(lineStr)

		lineNum++
		if lineNum%100000 == 0 {
			Log.Info(lineStr + " : " + string(lineNum))
			PrintMem()
		}
	}

	rootTrie = &root
	Log.Info("######### finish init #########")
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
		Log.Info(resStr)

		w.Header().Set("Content-Type", "application/json")
		w.Write(resStr)
	}
}
