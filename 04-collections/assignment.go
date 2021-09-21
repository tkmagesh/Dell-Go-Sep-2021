package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Dolore duis velit magna eu sunt et excepteur cupidatat ullamco non ex Aliquip amet consequat enim enim occaecat quis fugiat officia quis aliquip pariatur Consequat cillum eiusmod proident culpa duis dolor incididunt id occaecat ex consectetur id dolor Mollit fugiat sit irure do cupidatat officia deserunt laboris ametElit ut dolore incididunt irure elit consectetur Quis exercitation reprehenderit nostrud occaecat nisi ipsum nulla cillum quis labore tempor minim magna ullamco Et nostrud nostrud enim sunt esse excepteur exercitation ad officia sint aliqua exercitation doCillum Lorem reprehenderit minim minim Voluptate magna ipsum incididunt voluptate cillum enim Aute cupidatat fugiat tempor sint labore dolore dolore consectetur id anim ea voluptate Occaecat dolore do aute nulla reprehenderit Minim magna aliquip magna commodo dolore culpa sint nostrud sit non fugiat Eu sunt incididunt deserunt enim adipisicing nulla quis nostrud culpa dolore Mollit pariatur id velit qui duis irureId ea minim commodo labore ullamco proident laborum ad adipisicing Elit ullamco officia culpa magna amet voluptate dolor Occaecat pariatur ea consectetur sunt anim et Lorem in tempor labore velit pariatur velit Minim qui aliquip ad aliquip id nulla labore dolore nulla culpa culpa Irure laborum mollit veniam dolore Lorem Commodo aliqua esse incididunt sit in aliquip incididunt deserunt"

	//find out the length of the words that occurs the most by length and print the length and the number of occurrences

	//use strings.Split() to split the string into words
	words := strings.Split(str, " ")
	wordsCountBySize := getWordsCountBySize(words)
	size, occurances := getMaxOccuranceBySize(wordsCountBySize)
	fmt.Println(size, occurances)
}

func getWordsCountBySize(words []string) map[int]int {
	counts := make(map[int]int)
	for _, word := range words {
		counts[len(word)] += 1
	}
	return counts
}

func getMaxOccuranceBySize(wordsCountBySize map[int]int) (size int, occurances int) {
	maxOccurance := 0
	wordSize := 0
	for size, occurances := range wordsCountBySize {
		if occurances > maxOccurance {
			maxOccurance = occurances
			wordSize = size
		}
	}
	return wordSize, maxOccurance
}
