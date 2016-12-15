package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
)

func Increment(num []byte) []byte {
	num[len(num)-1] += 1
	for i := len(num) - 1; i > 0; i-- {
		if num[i] > '9' {
			num[i] = '0'
			num[i-1] += 1
		}
	}
	if num[0] > '9' {
		num[0] = '0'
		return append([]byte{'1'}, num...)
	} else {
		return num
	}
}

func MustWrite(w io.Writer, buf []byte) {
	n, err := w.Write(buf)
	if err != nil {
		log.Fatalf("Failed must write: %s", err)
	}

	if n != len(buf) {
		log.Fatalf("Failed must write: n=%d != len(buf)=%d", n, len(buf))
	}
}

func main() {
	log.SetFlags(0)
	h := md5.New()

	password := ""
	key := []byte(Input)
	num := []byte{'0'}

	for i := 0; len(password) < 8; i++ {
		MustWrite(h, key)
		MustWrite(h, num)
		sum := h.Sum(nil)
		h.Reset()

		if sum[0] == 0 && sum[1] == 0 && sum[2] < 0x10 {
			password = fmt.Sprintf("%s%x", password, sum[2])
			log.Print(password)
		}

		if i%1000000 == 0 {
			log.Print(string(num))
			log.Print(string(key) + string(num))
		}

		num = Increment(num)
	}

	log.Printf("password=%s", string(password))
}

var Input = "cxdnnyjw"
