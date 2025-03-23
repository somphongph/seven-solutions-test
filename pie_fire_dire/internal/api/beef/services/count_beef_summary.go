package services

import (
	"strings"
)

type BeefCount struct {
	Beef interface{} `json:"beef"`
}

func (s *service) CountBeefSummary(str string) BeefCount {
	replaceComma := strings.Replace(str, ",", " ", -1)
	replacePoint := strings.Replace(replaceComma, ".", " ", -1)
	lower := strings.ToLower(replacePoint)
	word := strings.Fields(lower)

	ch := make(chan map[string]int)
	go beefCount(word, ch)

	b := BeefCount{
		Beef: <-ch,
	}

	return b
}

func beefCount(content []string, ch chan map[string]int) {
	var count = make(map[string]int)
	for _, ct := range content {
		count[ct]++
	}

	ch <- count

	close(ch)
}
