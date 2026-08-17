package bencode

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type DECODED struct {
	INTVAL    *int64
	STRINGVAL *string
}

var BASE string = "Hello World"

var CASE_INT_REGEX = regexp.MustCompile(`^i(-?\d+)e$`)
var CASE_STR_REGEX = regexp.MustCompile(`^(\d+):(.*)`)
var CASE_LIS_REGEX = regexp.MustCompile(`^l(.*)e$`)

func TOKENIZE(line string) []rune {
	r := []rune(line)
	return r
}

func EVALUATE(expr string) []DECODED {
	tokens := TOKENIZE(expr)
	tokens = append(tokens, ' ') //Shit Code number 1 , Fadi added this line to process the last token becuase he's lazy
	var full_expr = []DECODED{}
	var temp = []rune{}
	for _, token := range tokens {
		if token != ' ' {
			temp = append(temp, token)
		} else {
			full_token := string(temp)
			fmt.Printf("Full Token: %s\n", full_token)
			temp = []rune{}
			eval := DECODE_SINGLE(full_token)
			full_expr = append(full_expr, eval)
		}
	}
	return full_expr
}

func DECODE_INT(expr string) int64 {
	str_val := expr[1 : len(expr)-1]
	int_val, _ := strconv.ParseInt(str_val, 10, 64)
	return int_val
}

func DECODE_STR(expr string) string {
	index := strings.Index(expr, ":")
	str_val := expr[index : len(expr)-1]
	return str_val
}

func DECODE_LIS(expr string) []DECODED {

}

func DECODE_SINGLE(expr string) DECODED {
	switch {
	case CASE_INT_REGEX.MatchString(expr):
		int_val := DECODE_INT(expr)
		return DECODED{INTVAL: &int_val}
	case CASE_STR_REGEX.MatchString(expr):
		str_val := DECODE_STR(expr)
		return DECODED{STRINGVAL: &str_val}
	case CASE_LIS_REGEX.MatchString(expr):

	}
	return DECODED{}
}
