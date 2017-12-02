package main

import (
	"bytes"
	//"errors"
	"fmt"
	"strings"
	"time"
)

//StrApend xxx
func StrApend(buf bytes.Buffer, strs ...string) string {
	for _, str := range strs {
		buf.WriteString(str)
	}
	return buf.String()
}

func benchmarkStringFunction(n int, index int) (d time.Duration) {
	v := "ni shuo wo shi bu shi tai wu liao le a?"
	var s string
	var buf bytes.Buffer

	t0 := time.Now()
	for i := 0; i < n; i++ {
		switch index {
		case 0: // fmt.Sprintf
			s = fmt.Sprintf("%s[%s]", s, v)
		case 1: // string +
			s = s + "[" + v + "]"
		case 2: // strings.Join
			s = strings.Join([]string{s, "[", v, "]"}, "")
		case 3: // temporary bytes.Buffer
			b := bytes.Buffer{}
			s = StrApend(b, "[", v, "]")

			// b.WriteString("[")
			// b.WriteString(v)
			// b.WriteString("]")
			// s = b.String()
		case 4: // stable bytes.Buffer
			s = StrApend(buf, "[", v, "]")
			// buf.WriteString("[")
			// buf.WriteString(v)
			// buf.WriteString("]")
		case 5:
			buf.WriteString("[")
			buf.WriteString(v)
			buf.WriteString("]")
		}

		if i == n-1 {
			if index == 5 { // for stable bytes.Buffer
				s = buf.String()
			}
			fmt.Printf("length of way(%d) is : %d\n", index, len(s)) // consume s to avoid compiler optimization
		}
	}
	t1 := time.Now()
	d = t1.Sub(t0)
	fmt.Printf("time of way(%d)=%v\n", index, d)
	return d
}

func main() {
	k := 6
	d := [6]time.Duration{}
	for i := 0; i < k; i++ {
		d[i] = benchmarkStringFunction(10000, i)
	}

	for i := 0; i < k-1; i++ {
		fmt.Printf("way %d is %6.1f times of way %d\n", i, float32(d[i])/float32(d[k-1]), k-1)
	}
}
