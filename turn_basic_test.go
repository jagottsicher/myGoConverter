package main

import (
	"os"
	"testing"
)

func TestMainFunc(t *testing.T) {

	type testData struct {
		inputArgStrings []string
		//ergebnis  int
	}

	tests := []testData{
		//testData{[]string{"-d", "23", "-v"}},
		testData{[]string{"-bin", "10111", "-v"}},
	}

	for _, v := range tests {
		os.Args = append(os.Args, v.inputArgStrings[0])
		os.Args = append(os.Args, v.inputArgStrings[1])
		os.Args = append(os.Args, v.inputArgStrings[2])

		main()
	}

	// output := main()

	// if output != `Hexadecimal: 0x17 17 (without preceding "0x")` {
	// 	t.Error("Expected:", `Hexadecimal: 0x17 17 (without preceding "0x")`, "but got:", output)
	// }

	// Test results here, and decide pass/fail.
}

// func TestDieSummeVonTableTest(t *testing.T) {

// 	type testData struct {
// 		inputInts []int
// 		ergebnis  int
// 	}

// 	tests := []testData{
// 		testData{[]int{1, 2, 3, 4, 5, 6}, 21},
// 		testData{[]int{1, 2, 4, 8, 16}, 31},
// 		testData{[]int{1, -2, 3, -4, 5}, 3},
// 		testData{[]int{-5, 1, 1, 1, 1, 1}, 0},
// 		testData{[]int{6, 5, 4, 3, 2, 1}, 21},
// 	}

// 	for _, v := range tests {
// 		v1 := DieSummeVon(v.inputInts...)
// 		if v1 != v.ergebnis {
// 			t.Error("Expected", v.ergebnis, "got", v1)
// 		}
// 	}
// }
