package models

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Question struct {
	ID      int      `json:"id"`
	Text    string   `json:"text"`
	Answers []Answer `json:"answers"`
}

//var questions map[int]Question

func GetQuestions() []Question {
	var questions = make([]Question, 0)
	loadJSON(&questions)
	return questions
}

func loadJSON(questions *[]Question) {
	jsonFile, _ := os.Open("questions.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &questions)
	defer jsonFile.Close()
}
