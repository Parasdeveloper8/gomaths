package mathematics

import "fmt"

func Mean(Sfixi float64, Sfi float64) float64 {
	x := Sfixi / Sfi
	fmt.Println(x)
	return x
}
func Mode(f1, f0, f2, l, h float64) float64 {
	//l + (f1-f0)/(2(f1)-f0-f2)*h
	formula := l + ((f1-f0)/(2*f1-f0-f2))*h
	return formula
}
func Median(cf, l, n, f, h float64) float64 {
	//l + (n/2 - cf/f) * h
	formula := l + (n/2-cf/f)*h
	return formula
}
func FindMedianData(num []float64) any {
	length := len(num)
	//If even
	if length%2 == 0 {
		length := len(num)        //n
		oTerm := (length / 2)     // n/2
		tTerm := (length / 2) + 1 // n/2 + 1
		oTerm_1 := oTerm - 1
		//index in array starts from 0 .
		//If we try to use oTerm and tTerm without -1 ,result will be wrong
		tTerm_1 := tTerm - 1
		var final float64 = (num[oTerm_1] + num[tTerm_1]) / 2
		return final
	}
	//if odd
	n := len(num)
	form := ((n + 1) / 2) - 1
	final := num[form]
	return final
}
