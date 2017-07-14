package a1

import (
	"fmt"
	"os"
	"bufio"
	"errors"
	"reflect"
	"math"
)

/* Assignment: A1
 * Course CMPT 383
 * Date: May 23, 2017
 * Author: Joshua Campbell
 * StudentNo: 301266191
 */

func countPrimes(n int) int {
	var numPrimes int = 0
	if n <= 1 {
		return 0
	}
	for i := 2; i <= n; i++ {
		var isPrime bool = true
		for j := 2; j < i; j++ {
			var number float64 = float64(i)
			var answer float64 = number / float64(j)

			var answerTest float64 = float64(int32(answer))

			if answer - answerTest <= 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			numPrimes++
		}
	}
	return numPrimes
}

// Go file reading procedure by Stefan Arentz on StackOverflow
// Found at: http://stackoverflow.com/questions/8757389/reading-file-line-by-line-in-go
//
// Go word by word reading example found at https://www.dotnetperls.com/file-go
func countStrings(filename string) map[string]int {
	var stringCount map[string]int = make(map[string]int)

	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		token := scanner.Text()
		var count int = stringCount[token]
		count++
		stringCount[token] = count
	}
	file.Close()

	return stringCount
}

// 0 <= hour < 24
// 0 <= minute < 60
// 0 <= second < 60
type Time24 struct {
    hour, minute, second uint8
}

func equalsTime24(a Time24, b Time24) bool {
	if a.hour == b.hour && a.minute == b.minute && a.second == b.second {
		return true
	} else {
		return false
	}
}

func lessThanTime24(a Time24, b Time24) bool {
	if a.hour < b.hour {
		return true
	} else if a.hour == b.hour && a.minute < b.minute {
		return true
	} else if a.hour == b.hour && a.minute == b.minute && a.second < b.second {
		return true
	} else {
		return false
	}
}

func (t Time24) String() string {
	var hourString string
	var minuteString string
	var secondString string
	var timeString string

	if t.hour < 10 {
		hourString = fmt.Sprintf("0%v", t.hour)
	} else {
		hourString = fmt.Sprintf("%v", t.hour)
	}

	if t.minute < 10 {
		minuteString = fmt.Sprintf("0%v", t.minute)
	} else {
		minuteString = fmt.Sprintf("%v", t.minute)
	}

	if t.second < 10 {
		secondString = fmt.Sprintf("0%v", t.second)
	} else {
		secondString = fmt.Sprintf("%v", t.second)
	}

	timeString = fmt.Sprintf("%v:%v:%v", hourString, minuteString, secondString)
	return timeString
}

func (t Time24) validTime24() bool {
	if t.hour >= 24 || t.hour < 0 {
		return false
	}
	if t.minute >= 60 || t.minute < 0 {
		return false
	}
	if t.second >= 60 || t.second < 0 {
		return false
	}
	return true
}

func minTime24(times []Time24) (Time24, error) {
	if len(times) == 0 {
		return Time24{0, 0, 0}, errors.New("Error: the times slice was empty")
	}
	minTime := times[0]
	for i := 1; i < len(times); i++ {
		if lessThanTime24(times[i], minTime) {
			minTime = times[i]
		}
	}
	return minTime, nil
}

func linearSearch(x interface{}, lst interface{}) int {
	switch z := lst.(type) {
	case []int:
		xType := reflect.TypeOf(x)
		lstType := reflect.TypeOf(z[0])
		if xType == lstType {
			for i := 0; i < len(z); i++ {
				if z[i] == x {
					return i
				}
			}
		} else {
			panic("lst element types do not match the type of x")
		}
	case []string:
		xType := reflect.TypeOf(x)
		lstType := reflect.TypeOf(z[0])
		if xType == lstType {
			for i := 0; i < len(z); i++ {
				if z[i] == x {
					return i
				}
			}
		} else {
			panic("lst element types do not match the type of x")
		}
	default:
		panic("lst is not a slice")
	}
	
	return -1
}

func allBitSeqs(n int) [][]int {
	var result [][]int
	if n <= 0 {
		return result
	}
	numSequences := int(math.Pow(2, float64(n)))
	result = make([][]int, numSequences)
	for i := 0; i < numSequences; i++ {
		result[i] = make([]int, n)
	}
	for seq := 1; seq < numSequences; seq++ {
		for i := n - 1; i >= 0; i-- {
			result[seq][i] = result[seq - 1][i]
		}
		for i := n - 1; i >= 0; i-- {
			if result[seq][i] == 0 {
				result[seq][i] = 1
				break
			} else {
				result[seq][i] = 0
			}
		}
	}
	return result
}
