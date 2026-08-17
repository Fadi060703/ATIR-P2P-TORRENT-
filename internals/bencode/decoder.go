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
	LISTVAL   []DECODED
}

var BASE string = "Hello World"

var CASE_INT_REGEX = regexp.MustCompile(`^i(-?\d+)e$`)
var CASE_STR_REGEX = regexp.MustCompile(`^(\d+):(.*)`)
var CASE_LIS_REGEX = regexp.MustCompile(`^l(.*)e$`)
var CASE_LIS_ITEM = regexp.MustCompile(`i-?\d+e|\d+:[a-zA-Z0-9_]*`)

func TOKENIZE(line string) []rune {
	r := []rune(line)
	return r
}

func TOKENIZE_LIST(line string) []string {
	var full_tokens = []string{}
	var cursor = 0

	for cursor < len(line) {
		switch line[cursor] {
		case 'i':
			start := cursor
			for cursor < len(line) && line[cursor] != 'e' {
				cursor++
			}
			cursor++
			full_tokens = append(full_tokens, line[start:cursor])

		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			start := cursor
			for cursor < len(line) && line[cursor] != ':' {
				cursor++
			}
			len_str := line[start:cursor]
			len_int, _ := strconv.ParseInt(len_str, 10, 64)
			cursor++
			cursor += int(len_int)
			full_tokens = append(full_tokens, line[start:cursor])
		case 'l':
			start := cursor
			nest := 0
			for cursor < len(line) {
				if line[cursor] == 'l' {
					nest++
					cursor++
				} else if line[cursor] == 'e' {
					nest--
					cursor++
					if nest == 0 {
						break
					}
				} else if line[cursor] == 'i' {
					for cursor < len(line) && line[cursor] != 'e' {
						cursor++
					}
					cursor++
				} else if line[cursor] >= '0' && line[cursor] <= '9' {
					strStart := cursor
					for cursor < len(line) && line[cursor] != ':' {
						cursor++
					}
					lenStr := line[strStart:cursor]
					length, _ := strconv.ParseInt(lenStr, 10, 64)
					cursor++
					cursor += int(length)
				} else {
					cursor++
				}
			}
			full_tokens = append(full_tokens, line[start:cursor])
		}
	}
	return full_tokens
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
	var decoded_items = []DECODED{}
	items := TOKENIZE_LIST(expr)
	for _, item := range items {
		decoded_items = append(decoded_items, DECODE_SINGLE(item))
	}
	return decoded_items
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
		lis_val := DECODE_LIS(expr[1 : len(expr)-1])
		return DECODED{LISTVAL: lis_val}
	}
	return DECODED{}
}
