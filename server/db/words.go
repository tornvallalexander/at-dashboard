package db

import (
	"fmt"
	"math/rand"
)

type TopWords struct {
	Words []string `json:"words"`
}

func (rdb *Database) GetAllTopWords() ([]string, error) {
	words, err := rdb.Client.SMembers(ctx, "topWords").Result()
	if err != nil {
		return nil, err
	}

	return words, nil
}

func (rdb *Database) GetRandomTopWords() ([]string, error) {
	words, err := rdb.Client.SMembers(ctx, "topWords").Result()
	if err != nil {
		return nil, err
	}

	var randomWords []string

	for i := 0; i < 50; i++ {
		randomIndex := rand.Intn(len(words))
		randomWords = append(randomWords, words[randomIndex])
	}

	return randomWords, nil
}

// SetTopWords adds a slice of words to the redis db
func (rdb *Database) SetTopWords(words []string) {
	for _, word := range words {
		res, err := rdb.Client.SAdd(ctx, "topWords", word).Result()
		SAddError(res, err)
	}
}

// SAddError checks for errors when performing rdb.SAdd
func SAddError(res int64, err error) {
	if err != nil {
		fmt.Printf("Error adding word: %d\n", res)
		return
	}
	fmt.Println("Success")
}
