package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input string
	fmt.Print("Enter encoded: ")
	fmt.Scanln(&input)

	s, err := decode(input)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)
}

func decode(str string) (string, error) {
	if str == "" {
		return "", errors.New("string is empty can't decode")
	}

	upperStr := strings.ToUpper(str)
	count := len(upperStr)
	limit := 10
	for range count {
		limit *= 10
	}

	sum := 0
	leastSum := limit
	leastStr := ""

	for i := range limit {
		num := strconv.Itoa(i)
		padStr := padLeft(num, "0", count+1)
		strDecode, err := decodeStr(padStr, upperStr, 0)
		if err == nil {
			sum = 0
			for _, s := range strDecode {
				sum += int(s - '0')
			}

			if leastSum > sum {
				leastSum = sum
				leastStr = strDecode
			}
		}
	}

	return leastStr, nil
}

func padLeft(str, pad string, length int) string {
	for len(str) < length {
		str = pad + str
	}
	return str
}

func decodeStr(padStr string, upperStr string, pos int) (string, error) {
	if pos >= len(upperStr) {
		return padStr, nil
	}

	sCode := string(upperStr[pos])
	sLeft := string(padStr[pos])
	sRight := string(padStr[pos+1])
	pos++

	if sCode == "=" && sLeft == sRight {
		return decodeStr(padStr, upperStr, pos)
	} else if sCode == "L" && sLeft > sRight {
		return decodeStr(padStr, upperStr, pos)
	} else if sCode == "R" && sLeft < sRight {
		return decodeStr(padStr, upperStr, pos)
	} else {
		return "", errors.New("no mach")
	}
}
