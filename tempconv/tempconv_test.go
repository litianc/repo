package tempconv

import "testing"
import "fmt"

func TestTempconv(t *testing.T) { 
	var tests = []struct {
        	c Celsius
        	f Fahrenheit
    	}{
		{ 100, 212 },
		{ 5, 41 },
		//{ 2, 35.6 },
	}
	for _, test := range tests {
		if (CToF(test.c) != test.f) {
			fmt.Println(test.c, test.f)
			fmt.Println(CToF(test.c))
			t.Error(`CToF(test.c) = test.f`)
		}
		if (FToC(test.f) != test.c) {
			fmt.Println(test.c, test.f)
			fmt.Println(FToC(test.f))
			t.Error(`FToC(test.f) != test.c`)
		}
	}
}
