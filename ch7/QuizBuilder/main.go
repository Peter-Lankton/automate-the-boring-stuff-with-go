package main

import (
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	capitals := map[string]string{
		"Alabama":        "Montgomery",
		"Alaska":         "Juneau",
		"Arizona":        "Phoenix",
		"Arkansas":       "Little Rock",
		"California":     "Sacramento",
		"Colorado":       "Denver",
		"Connecticut":    "Hartford",
		"Delaware":       "Dover",
		"Florida":        "Tallahassee",
		"Georgia":        "Atlanta",
		"Hawaii":         "Honolulu",
		"Idaho":          "Boise",
		"Illinois":       "Springfield",
		"Indiana":        "Indianapolis",
		"Iowa":           "Des Moines",
		"Kansas":         "Topeka",
		"Kentucky":       "Frankfort",
		"Louisiana":      "Baton Rouge",
		"Maine":          "Augusta",
		"Maryland":       "Annapolis",
		"Massachusetts":  "Boston",
		"Michigan":       "Lansing",
		"Minnesota":      "Saint Paul",
		"Mississippi":    "Jackson",
		"Missouri":       "Jefferson City",
		"Montana":        "Helena",
		"Nebraska":       "Lincoln",
		"Nevada":         "Carson City",
		"New Hampshire":  "Concord",
		"New Jersey":     "Trenton",
		"New Mexico":     "Santa Fe",
		"New York":       "Albany",
		"North Carolina": "Raleigh",
		"North Dakota":   "Bismarck",
		"Ohio":           "Columbus",
		"Oklahoma":       "Oklahoma City",
		"Oregon":         "Salem",
		"Pennsylvania":   "Harrisburg",
		"Rhode Island":   "Providence",
		"South Carolina": "Columbia",
		"South Dakota":   "Pierre",
		"Tennessee":      "Nashville",
		"Texas":          "Austin",
		"Utah":           "Salt Lake City",
		"Vermont":        "Montpelier",
		"Virginia":       "Richmond",
		"Washington":     "Olympia",
		"West Virginia":  "Charleston",
		"Wisconsin":      "Madison",
		"Wyoming":        "Cheyenne",
	}

	var states, capitalsItems []string
	for y, x := range capitals {
		capitalsItems = append(capitalsItems, x)
		states = append(states, y)
	}

	for i := 0; i < 35; i++ {
		// сreate the quiz text file
		str1 := "quiz_" + strconv.Itoa(i+1) + ".txt"
		quizFile, err := os.Create(str1)
		check(err)
		defer quizFile.Close()
		// create the answer key to the quiz
		str2 := "answer_key_" + strconv.Itoa(i+1) + ".txt"
		answerKeyFile, err := os.Create(str2)
		check(err)
		defer answerKeyFile.Close()

		// Create portion for students to fill out
		quizFile.WriteString("Student Number:\n\nName:\n\nDate:\n\n")
		str3 := "Quiz " + strconv.Itoa(i+1)
		quizFile.WriteString(strings.Repeat(" ", 20) + str3)
		quizFile.WriteString("\n\n")

		rand.Seed(time.Now().UnixNano())
		// mix of the States
		shuffle(states)

		// Iterate through and build the question out
		for j := 0; j < 50; j++ {
			correctAnswer := capitals[states[j]]
			wrongAnswers := make([]string, len(capitalsItems))
			copy(wrongAnswers, capitalsItems)

			// shuffle wrong answers
			answNoCorrect := make([]string, len(wrongAnswers)-1)
			for l := 0; l < len(wrongAnswers); l++ {
				if wrongAnswers[l] == correctAnswer {
					copy(answNoCorrect, removeAtIndex(wrongAnswers, l))
				}
			}

			// create answer options A-D
			var answerOptions []string
			for l := 0; l < 3; l++ {
				answerOptions = append(answerOptions, answNoCorrect[l])
			}
			answerOptions = append(answerOptions, correctAnswer)
			shuffle(answerOptions)

			// create question
			str3 := strconv.Itoa(j+1) + " What is the Capital of " + states[j] + "?" + "\n"
			quizFile.WriteString(str3)
			strAbcd := "ABCD"
			for l := 0; l < 4; l++ {
				strAnsw := string(strAbcd[l]) + ". " + answerOptions[l] + "\n"
				quizFile.WriteString(strAnsw)
			}
			// make quiz and save it
			quizFile.WriteString("\n")

			// make answer key and save it
			strAnswerOk := ""
			for l := 0; l < len(answerOptions); l++ {
				if answerOptions[l] == correctAnswer {
					strAnswerOk += string(strAbcd[l])
				}
			}
			strCorAnsw := strconv.Itoa(j+1) + ". " + strAnswerOk + "\n"
			answerKeyFile.WriteString(strCorAnsw)
		}
	}
}

// helper functions for making quiz building easier
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func shuffle(a []string) {
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

func removeAtIndex(source []string, index int) []string {
	lastIndex := len(source) - 1
	source[index], source[lastIndex] = source[lastIndex], source[index]
	return source[:lastIndex]
}
