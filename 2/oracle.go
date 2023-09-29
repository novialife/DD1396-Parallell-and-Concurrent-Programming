// Stefan Nilsson 2013-03-13

// This program implements an ELIZA-like oracle (en.wikipedia.org/wiki/ELIZA).
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	star   = "Pythia"
	venue  = "Delphi"
	prompt = "> "
)

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	oracle := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", star, line)
		oracle <- line // The channel doesn't block.
	}
}

// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.
func Oracle() chan<- string {
	questions := make(chan string)
	answers := make(chan string)
	go prophecy("", answers)

	go func() {
		for question := range questions {
			go ansQuestion(question, answers) // go routine that answers users questions
		}
	}()
	go printAnswers(answers) // go routine that prints out answers nicely

	// TODO: Answer questions.
	// TODO: Make prophecies.
	// TODO: Print answers.
	return questions
}

func ansQuestion(questions string, answers chan<- string) {
	randomanswers := []string{ // list of random answers
		"Lorem Ipsum",
		"Dolor sit amet",
		"U41",
		"U32",
	}

	time.Sleep(10 * time.Second)
	answers <- randomanswers[rand.Intn(len(randomanswers))]

}

func printAnswers(answers chan string) {
	for ans := range answers { // For each answer
		for _, c := range ans { // For each letter in each anwer, omit the index
			time.Sleep(50 * time.Millisecond) // Wait some time
			fmt.Printf(string(c))             // print the letter
		}
		fmt.Printf("\n")
	}
}

// This is the oracle's secret algorithm.
// It waits for a while and then sends a message on the answer channel.
// TODO: make it better.
func prophecy(question string, answer chan<- string) {
	for { // Add conditionless for to keep this running
		// Keep them waiting. Pythia, the original oracle at Delphi,
		// only gave prophecies on the seventh day of each month.

		time.Sleep(time.Duration(20+rand.Intn(10)) * time.Second)

		// Find the longest word.
		longestWord := ""
		words := strings.Fields(question) // Fields extracts the words into a slice.
		for _, w := range words {
			if len(w) > len(longestWord) {
				longestWord = w
			}
		}

		// Cook up some pointless nonsense.
		nonsense := []string{
			"The moon is dark.",
			"The sun is bright.",
		}
		answer <- longestWord + "... " + nonsense[rand.Intn(len(nonsense))]
	}
}

func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}
