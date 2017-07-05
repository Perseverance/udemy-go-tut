package main

import (
	"bufio"
	"fmt"
	"net/http"
)

const cBuckets = 12

func main() {
	url := "https://gist.githubusercontent.com/StevenClontz/4445774/raw/1722a289b665d940495645a5eaaad4da8e3ad4c7/mobydick.txt"
	dict, err := createWordsHashMap(url, cBuckets)
	if err != nil {
		fmt.Println(err)
	}
	for key, val := range dict {
		fmt.Println("=======================", key, "=======================")
		fmt.Println(len(val))
	}
}

func createWordsHashMap(url string, buckets int) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error Getting The Response")
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)

	result := make([][]string, buckets)

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		ind := hashBucket(word, buckets)
		result[ind] = append(result[ind], word)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func hashBucket(word string, buckets int) int {
	var total int

	for _, val := range word {
		total += int(val)
	}

	return total % buckets
}
