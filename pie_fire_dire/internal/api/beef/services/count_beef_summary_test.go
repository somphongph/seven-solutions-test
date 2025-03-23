package services

import (
	"testing"
)

func TestBeefCount(t *testing.T) {
	t.Run("Beef Count", func(t *testing.T) {
		ch := make(chan map[string]int)

		content := []string{"fatback", "t-bone", "t-bone", "pastrami", "t-bone", "pork", "t-bone"}
		go beefCount(content, ch)

		result := <-ch

		expected := map[string]int{"t-bone": 4, "fatback": 1, "pork": 1, "pastrami": 1}
		for key, value := range expected {
			if result[key] != value {
				t.Errorf("Expected %s to be %d, got %d", key, value, result[key])
			}
		}
	})
}
