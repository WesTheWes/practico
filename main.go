package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type exercise struct {
	name      string
	startTime time.Time
	endTime   time.Time
	maxBPM    int
}

func main() {
	var exercises []exercise
	for {
		name, minutes, bpm := getExerciseInfo()
		startTime := time.Now()
		endTime := startTime.Add(time.Duration(minutes) * time.Minute)
		for {
			for {
				fmt.Printf("%v more seconds to go!\n", int(math.Round(endTime.Sub(time.Now()).Seconds())))
				fmt.Printf("Repeat this exercise at %v bpm\n", bpm)
				difficulty := getDifficulty()
				newBPM := getNextBPM(bpm, difficulty)
				if endTime.Before(time.Now()) {
					break
				}
				bpm = newBPM
			}
			continueExercise := continueExercise(minutes)
			if continueExercise {
				minutesInput := getUserInput("How long do you want to work on this (minutes)?", testPostitveInteger)
				moreMinutes, _ := strconv.Atoi(minutesInput)
				minutes += minutes + moreMinutes
				endTime = endTime.Add(time.Duration(moreMinutes) * time.Minute)
			} else {
				fmt.Printf("Great job practicing %v!\nYou practiced for %v minutes, and your most recent bpm was %v\n", name, minutes, bpm)
				break
			}
		}
		newExercise := exercise{name: name, startTime: startTime, endTime: endTime}
		exercises = append(exercises, newExercise)
		continuePractice := continuePractice()
		if !continuePractice {
			break
		}
	}
	fmt.Println(exercises)
}

func getUserInput(inputString string, isValidInput func(string) error) string {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println(inputString)
		fmt.Print(">")
		scanner.Scan()
		input := scanner.Text()
		if err := isValidInput(input); err != nil {
			fmt.Printf("%v\n", err)
		} else {
			return input
		}
	}
}

func testStringLength(input string) error {
	if len(input) < 1 {
		return errors.New("Input cannot be blank")
	}
	return nil
}

func testPostitveInteger(input string) error {
	if convertedInt, err := strconv.Atoi(input); err != nil || convertedInt < 1 {
		return errors.New("Input must be a positive integer")
	}
	return nil
}

func testYesOrNo(input string) error {
	if !strings.EqualFold(input, "n") && !strings.EqualFold(input, "y") {
		return errors.New("Please enter y or n")
	}
	return nil
}

func testInArray(inputArray []string) func(string) error {
	return func(input string) error {
		for _, v := range inputArray {
			if v == input {
				return nil
			}
		}
		return fmt.Errorf("Input needs to be one of the following: %v", inputArray)
	}
}

func getExerciseInfo() (name string, minutes int, bpm int) {
	name = getUserInput("What would you like to work on?", testStringLength)
	bpmInput := getUserInput("About what speed do you think you can play this at? (quarter notes)?", testPostitveInteger)
	bpm, _ = strconv.Atoi(bpmInput)
	minutesInput := getUserInput("How long do you want to work on this (minutes)?", testPostitveInteger)
	minutes, _ = strconv.Atoi(minutesInput)
	return
}

func getDifficulty() int {
	testFunction := testInArray([]string{"1", "2", "3", "4", "5"})
	question := "On a scale of 1 to 5, how did that feel?\n"
	question += "1 - Impossible\n"
	question += "2 - Lots of Mistakes\n"
	question += "3 - A few mistakes\n"
	question += "4 - No mistakes but challenging\n"
	question += "5 - No mistakes and too easy"
	difficultyString := getUserInput(question, testFunction)
	difficulty, _ := strconv.Atoi(difficultyString)
	return difficulty
}

func continueExercise(minutesElapsed int) bool {
	continueQuestion := fmt.Sprintf("%v minutes have passed. Would you like to continue? [y,n]", minutesElapsed)
	yesOrNo := getUserInput(continueQuestion, testYesOrNo)
	return yesOrNo == "y"
}

func continuePractice() bool {
	continuePractice := getUserInput("Would you like to continue practice with another exercise? [y,n]", testYesOrNo)
	return continuePractice == "y"
}

func getNextBPM(bpm int, difficulty int) (newBPM int) {
	switch difficulty {
	case 1:
		newBPM = int(math.Round(float64(bpm) / 3))
	case 2:
		newBPM = int(math.Round(float64(bpm) / 2))
	case 3:
		newBPM = bpm - 5
	case 4:
		newBPM = bpm
	case 5:
		newBPM = bpm + 2
	}
	return
}
