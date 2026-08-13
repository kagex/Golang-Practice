package main

import (
	"fmt"
	"sort"
)

type movies map[string]map[string]float64

func main() {
	listOfMovies := movies{
		"Экшен": {
			"Фильм1": 9.52,
			"Фильм2": 6.0,
		},
		"Драма": {
			"Фильм3": 7.524,
			"Фильм4": 7.527,
			"Фильм5": 5.54,
		},
	}
	printRecommendations(listOfMovies)
}

func printRecommendations(list movies) {
	genres := []string{}

	for genre := range list {
		genres = append(genres, genre)
	}

	sort.Strings(genres)

	copiedList := make(movies, len(list))

	for genre, movies := range list {
		for movie, score := range movies {
			if score >= 7 {
				if copiedList[genre] == nil {
					copiedList[genre] = map[string]float64{}
				}

				copiedList[genre][movie] = score
			}
		}
	}

	for _, genre := range genres {
		if len(copiedList[genre]) == 0 {
			continue
		}

		movies := copiedList[genre]

		names := []string{}

		for movie := range movies {
			names = append(names, movie)
		}

		sort.Slice(names, func(i, j int) bool {
			leftScore := movies[names[i]]
			rightScore := movies[names[j]]

			if leftScore != rightScore {
				return leftScore > rightScore
			}

			return names[i] < names[j]
		})

		fmt.Printf("%s: ", genre)

		for i, movie := range names {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%s (%.1f)", movie, movies[movie])
		}
		fmt.Println(".")
	}
}
