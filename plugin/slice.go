package plugin

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func init() {
	Register("slice", &Slice{})
	Register("slicestr", &Slice{true})
}

type Slice struct {
	isString bool
}

func (s *Slice) Exe(cmd string, args []string) ([]string, error) {
	if len(args) < 2 {
		return nil, nil
	}

	from, to, err := parseIndexs(args[0])
	if err != nil {
		return nil, err
	}

	result := make([]string, 0, len(args)-1)
	if s.isString {
		for _, s := range args[1:] {
			from, to := adjustIndexs(from, to, len(s))
			result = append(result, s[from:to])
		}
	} else {
		for _, s := range args[1:] {
			d := toBytes(s)
			from, to := adjustIndexs(from, to, len(d))
			result = append(result, toHex(d[from:to]))
		}
	}

	return result, nil
}

func parseIndexs(s string) (int, int, error) {
	indexs := strings.Split(s, ":")
	if len(indexs) == 2 {
		from := 0
		to := math.MaxInt64
		if n, err := strconv.Atoi(indexs[0]); err == nil {
			from = n
		}

		if n, err := strconv.Atoi(indexs[1]); err == nil {
			to = n
		}

		return from, to, nil
	}

	return 0, 0, fmt.Errorf("slice, invalid index format, %s", s)
}

func adjustIndexs(from, to, max int) (int, int) {
	if from > to {
		from, to = to, from
	}

	adjust := func(a int) int {
		if a < 0 {
			return 0
		}

		if a > max {
			return max
		}

		return a
	}

	return adjust(from), adjust(to)
}
