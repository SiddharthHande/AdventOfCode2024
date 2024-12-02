package problemDayOne

import (
	"adventofcode/helper"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func ProblemOne(data *[]string) {
	var a, b []int64
	ans := 0
	for _, v := range *data {
		temp := strings.Fields(v)
		if i, err := strconv.ParseInt(temp[0], 10, 64); err == nil {
			a = append(a, i)
		}
		if i, err := strconv.ParseInt(temp[1], 10, 64); err == nil {
			b = append(b, i)
		}
	}
	slices.Sort(a)
	slices.Sort(b)
	for i := range len(*data) {
		ans += int(math.Abs(float64(a[i] - b[i])))
	}
	fmt.Println(ans)
}

func ProblemTwo(data *[]string) {
	a := []string{}
	b := make(map[string]int)
	ans := 0
	for _, v := range *data {
		temp := strings.Fields(v)
		a = append(a, temp[0])
		b[temp[1]]++
	}

	for _, k := range a {
		if b[k] != 0 {
			ans += helper.QuickParseInt(k) * b[k]
		}
	}

	fmt.Println(ans)
}
