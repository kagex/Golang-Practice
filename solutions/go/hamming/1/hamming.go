package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {return 0, errors.New("can`t compare this DNA`s")}

    counter := 0
    for i:= 0; i < len(a); i++ {
        if a[i] != b[i] {counter++}
    }
    return counter, nil
}
