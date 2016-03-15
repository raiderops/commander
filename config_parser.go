package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"sync"
)

type Playbook []map[string]map[string]string

func main() {
	filename, _ := filepath.Abs("./test.yml")
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var pbook Playbook

	err = yaml.Unmarshal(yamlFile, &pbook)
	if err != nil {
		panic(err)
	}

	wg := new(sync.WaitGroup)
	defer wg.Wait()

	//fmt.Printf("Vaule: %s\n", pbook)
	for k := range pbook {
		if _, found := pbook[k]["file"]; found {
			wg.Add(1)
			go handleFileTask(pbook[k], wg)
		} else if _, found := pbook[k]["user"]; found {
			wg.Add(1)
			go handleUserTask(pbook[k], wg)
		}
	}

}

func handleFileTask(ftask map[string]map[string]string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Creating new File: " + ftask["file"]["path"] + "\n")
}

func handleUserTask(utask map[string]map[string]string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Creating new User: " + utask["user"]["name"])
}
