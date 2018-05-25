package conf

import (
	"os"
	"log"
	"encoding/json"
	"fmt"
)

type Configuration struct {
	WordsPath string
	LogPath   string
}

var Conf Configuration

func init() {
	log.SetFlags(log.Lshortfile|log.Ldate|log.Lmicroseconds)

	confFile, err := os.Open("./conf.json")
	if err != nil {
		log.Fatalln("init conf fail", err)
	}

	decoder := json.NewDecoder(confFile)

	if err := decoder.Decode(&Conf); err != nil {
		log.Fatalln("conf.json can not decode", err)
	}

	fmt.Println(Conf)
}
