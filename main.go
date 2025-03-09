package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/hheedat/stopfilter/conf"
	slog "github.com/hheedat/stopfilter/log"
	"github.com/hheedat/stopfilter/util"
	"strings"
)

type ResJson struct {
	Code       int32  `json:"code"`
	IsExist    bool   `json:"isExist"`
	OriginText string `json:"originText"`
	Reason     string `json:"reason"`
}

var rootTrie *util.TrieNode

func buildStruct() {
	root := util.TrieNode{nil, false}
	inFile, err := os.Open(conf.Conf.WordsPath)
	if err != nil {
		slog.Log.Error("read words fail")
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	lineNum := 0
	for scanner.Scan() {
		lineStr := scanner.Text()

		root.AddWord(lineStr)

		lineNum++
		if lineNum%10000 == 0 {
			slog.Log.Info(lineStr + " : " + string(lineNum))
			util.PrintMem()
		}
	}
	slog.Log.Info("### finish load")
	util.PrintMem()

	rootTrie = &root
	slog.Log.Info("### finish init")
	util.PrintMem()
}

func main() {
	defer slog.Destroy()
	buildStruct()

	http.HandleFunc("/", search)

	err := http.ListenAndServe(":"+conf.Conf.ListenPort, nil)
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
		slog.Log.Info(string(resStr))
		util.PrintMem()

		w.Header().Set("Content-Type", "application/json")
		w.Write(resStr)
	}
}
