package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	separator             = "---"
	superSeparator        = "%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%"
	blankLine             = "\n"
	minTopicsPerParagraph = 2
	maxTopicsPerParagraph = 4
	minParagraphs         = 2
	maxParagraphs         = 5
)

func main() {
	songs := getSongs()

	var result []string
	result = append(result, superSeparator)

	for _, song := range songs {
		result = append(result, song)
		result = append(result, superSeparator)
		result = append(result, getTopics()...)
		result = append(result, superSeparator)
	}

	printResult(result)
}

func getAllTopics() []string {
	return []string{"Procedentes de Madrid", "A menudo comparados con Corrupted", "Tres discos", "Iñaki, Edu y Raúl",
		"Sludge", "Empezaron en 2002", "Se separaron en 2011", "Cientos de conciertos"}
}

func getSongs() []string {
	return []string{"El Tren", "Culebra", "Rompiendo Huesos", "Punk", "Gusano de Fuego", "Caronte (Escape)",
		"Semana Santa", "El Duelo", "Fístula", "180º", "Coche Fúnebre", "Lava", "Gargantor", "El Segador",
		"Chotacabra", "San Mamés", "Gargantor", "Demacronch", "Terror Ultramarino", "Torpedo", "Anciago",
		"Garage Champion"}
}

func getTopics() []string {
	numberOfParagraphs := getRandomIntBetween(minParagraphs, maxParagraphs)

	var result []string

	for i := 0; i < numberOfParagraphs; i++ {
		numberOfTopics := getRandomIntBetween(minTopicsPerParagraph, maxTopicsPerParagraph)
		topics := getAllTopics()
		shuffledTopics := shuffle(topics)

		for j := 0; j < numberOfTopics; j++ {
			result = append(result, shuffledTopics[j])
		}

		if i != numberOfParagraphs-1 {
			result = append(result, separator)
		} else {
			result = append(result, blankLine)
		}
	}

	return result
}

func printResult(result []string) {
	for _, line := range result {
		fmt.Println(line)
	}
}

func shuffle(slice []string) []string {
	rand.Seed(time.Now().UnixNano())

	for i := len(slice) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}

	return slice
}

func getRandomIntBetween(min int, max int) int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max-min+1) + min
}
