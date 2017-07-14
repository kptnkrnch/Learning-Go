package a1

import (
	"testing"
	"fmt"
)

/* Assignment: A1 test
 * Course CMPT 383
 * Date: May 23, 2017
 * Author: Joshua Campbell
 * StudentNo: 301266191
 */

func TestCountPrimes(t *testing.T) {
	var v int
	
	v = countPrimes(5)
	if v != 3 {
		t.Error("Expected 3, got ", v)
	}

	v = countPrimes(10000)
	if v != 1229 {
		t.Error("Expected 1229, got ", v)
	}

	v = countPrimes(-6)
	if v != 0 {
		t.Error("Expected 0, got ", v)
	}
}

func TestCountStrings(t *testing.T) {
	var result map[string]int

	result = countStrings("test1.txt")

	if result["dog"] != 1 || result["The"] != 1 || result["the"] != 1 || result["big"] != 3 || result["ate"] != 1 || result["apple"] != 1 {
		t.Error("Expected map[ate:1 the:1 apple:1 The:1 big:3 dog:1], got ", result)
	}
}

func TestEqualsTime24(t *testing.T) {
	var a Time24 = Time24{hour: 1, minute: 55, second: 8}
	var b Time24 = Time24{hour: 1, minute: 55, second: 8}
	var c Time24 = Time24{hour: 11, minute: 55, second: 8}
	var d Time24 = Time24{hour: 1, minute: 54, second: 8}
	var e Time24 = Time24{hour: 1, minute: 55, second: 7}
	var f Time24 = Time24{hour: 12, minute: 17, second: 11}

	result := equalsTime24(a, b)
	if result != true {
		t.Error("Expected time 01:55:08 equals time 01:55:08, got time 01:55:08 does not equal time 01:55:08")
	}

	result = equalsTime24(a, c)
	if result != false {
		t.Error("Expected time 01:55:08 does not equal time 11:55:08, got time 01:55:08 equals time 11:55:08")
	}

	result = equalsTime24(a, d)
	if result != false {
		t.Error("Expected time 01:55:08 does not equal time 01:54:08, got time 01:55:08 equals time 01:54:08")
	}

	result = equalsTime24(a, e)
	if result != false {
		t.Error("Expected time 01:55:08 does not equal time 01:55:07, got time 01:55:08 equals time 01:55:07")
	}

	result = equalsTime24(a, f)
	if result != false {
		t.Error("Expected time 01:55:08 does not equal time 12:17:11, got time 01:55:08 equals time 12:17:11")
	}
}

func TestLessThanTime24(t *testing.T) {
	var a Time24 = Time24{hour: 1, minute: 55, second: 8}
	var b Time24 = Time24{hour: 1, minute: 55, second: 8}
	var c Time24 = Time24{hour: 2, minute: 55, second: 8}
	var d Time24 = Time24{hour: 1, minute: 54, second: 8}
	var e Time24 = Time24{hour: 1, minute: 55, second: 9}
	var f Time24 = Time24{hour: 1, minute: 55, second: 7}

	result := equalsTime24(a, b)
	if result != true {
		t.Error("Expected time 01:55:08 not less than time 01:55:08, got time 01:55:08 less than time 01:55:08")
	}

	result = equalsTime24(a, c)
	if result != false {
		t.Error("Expected time 01:55:08 less than time 02:55:08, got time 01:55:08 not less than time 02:55:08")
	}

	result = equalsTime24(a, d)
	if result != false {
		t.Error("Expected time 01:55:08 not less than time 01:54:08, got time 01:55:08 less than time 01:54:08")
	}

	result = equalsTime24(a, e)
	if result != false {
		t.Error("Expected time 01:55:08 less than time 01:55:09, got time 01:55:08 not less than time 01:55:09")
	}

	result = equalsTime24(a, f)
	if result != false {
		t.Error("Expected time 01:55:08 not less than time 01:55:07, got time 01:55:08 less than time 01:55:07")
	}
}

func TestTime24String(t *testing.T) {
	var time Time24 = Time24{hour: 5, minute: 39, second: 8}

	if time.String() != "05:39:08" {
		t.Error("Expected 05:39:08, got ", time)
	}
}

func TestValidTime24(t *testing.T) {
	var a Time24 = Time24{hour: 0, minute: 0, second: 0}
	var b Time24 = Time24{hour: 24, minute: 60, second: 60}
	var c Time24 = Time24{hour: 23, minute: 59, second: 59}
	var d Time24 = Time24{hour: 25, minute: 7, second: 8}
	var e Time24 = Time24{hour: 1, minute: 70, second: 9}
	var f Time24 = Time24{hour: 1, minute: 55, second: 70}

	result := a.validTime24()
	if result != true {
		t.Error("Expected true, got ", result)
	}

	result = b.validTime24()
	if result != false {
		t.Error("Expected false, got ", result)
	}

	result = c.validTime24()
	if result != true {
		t.Error("Expected true, got ", result)
	}

	result = d.validTime24()
	if result != false {
		t.Error("Expected false, got ", result)
	}

	result = e.validTime24()
	if result != false {
		t.Error("Expected false, got ", result)
	}

	result = f.validTime24()
	if result != false {
		t.Error("Expected false, got ", result)
	}
}

func TestMinTime24(t *testing.T) {
	var a Time24 = Time24{hour: 0, minute: 1, second: 3}
	var b Time24 = Time24{hour: 0, minute: 0, second: 3}
	var c Time24 = Time24{hour: 0, minute: 0, second: 1}
	var d Time24 = Time24{hour: 0, minute: 0, second: 0}
	var e Time24 = Time24{hour: 0, minute: 1, second: 0}
	var f Time24 = Time24{hour: 1, minute: 0, second: 0}

	var lst []Time24 = []Time24{ a, b, c, d, e, f }

	result, _ := minTime24(lst)
	if result != d {
		t.Error("Expected 00:00:00, got ", result)
	}
}

func TestLinearSearch(t *testing.T) {
	var result int
	
	result = linearSearch(5, []int{4, 2, -1, 5, 0})
	if result != 3 {
		t.Error("Expected 3, got ", result)
	}

	result = linearSearch(3, []int{4, 2, -1, 5, 0})
	if result != -1 {
		t.Error("Expected -1, got ", result)
	}

	result = linearSearch("egg", []string{"cat", "nose", "egg"})
	if result != 2 {
		t.Error("Expected 2, got ", result)
	}

	result = linearSearch("up", []string{"cat", "nose", "egg"})
	if result != -1 {
		t.Error("Expected -1, got ", result)
	}
}

func TestAllBitsSeqs(t *testing.T) {
	var result [][]int

	result = allBitSeqs(-1)
	if fmt.Sprintf("%v", result) != "[]" {
		t.Error("Expected -1, got ", result)
	}
	
	result = allBitSeqs(0)
	if fmt.Sprintf("%v", result) != "[]" {
		t.Error("Expected -1, got ", result)
	}
	
	result = allBitSeqs(1)
	if fmt.Sprintf("%v", result) != "[[0] [1]]" {
		t.Error("Expected -1, got ", result)
	}
	
	result = allBitSeqs(2)
	if fmt.Sprintf("%v", result) != "[[0 0] [0 1] [1 0] [1 1]]" {
		t.Error("Expected -1, got ", result)
	}
	
	result = allBitSeqs(3)
	if fmt.Sprintf("%v", result) != "[[0 0 0] [0 0 1] [0 1 0] [0 1 1] [1 0 0] [1 0 1] [1 1 0] [1 1 1]]" {
		t.Error("Expected -1, got ", result)
	}
	
	result = allBitSeqs(4)
	if fmt.Sprintf("%v", result) != "[[0 0 0 0] [0 0 0 1] [0 0 1 0] [0 0 1 1] [0 1 0 0] [0 1 0 1] [0 1 1 0] [0 1 1 1] [1 0 0 0] [1 0 0 1] [1 0 1 0] [1 0 1 1] [1 1 0 0] [1 1 0 1] [1 1 1 0] [1 1 1 1]]" {
		t.Error("Expected -1, got ", result)
	}
}
