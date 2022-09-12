package tokenize

import (
	"strconv"
	"unicode"
)

var userInput []rune

// Pos
var lno int
var lat int
var wat int

var singleOpSymbols []string
var compositeOpSymbols []string

func init() {
	singleOpSymbols = []string{
		"(", ")", "[", "]", "{", "}",
		".", ",", ":", ";",
		"+", "-", "*", "/", "%",
		">", "<",
		"=", "!",
	}
	compositeOpSymbols = []string{
		"==", "!=", ">=", "<=",
		"+=", "-=", "*=", "/=", "%=",
		"&&", "||", ":=",
	}
}

func startWith(q string) bool {
	qRunes := []rune(q)
	for i := 0; i < len(qRunes); i++ {
		if len(userInput) <= wat+i || userInput[wat+i] != qRunes[i] {
			return false
		}
	}
	return true
}

func isIdentRune(r rune) bool {
	return ('a' <= r && r <= 'z') ||
		('A' <= r && r <= 'Z') ||
		('0' <= r && r <= '9') ||
		('_' == r)
}

func isNotEof() bool {
	return wat < len(userInput)
}

func consumeComment() string {
	// skip "#"
	lat++
	wat++

	var s string
	for isNotEof() {
		if userInput[wat] == '\n' {
			break
		}
		s += string(userInput[wat])
		lat++
		wat++
	}
	return s
}

func consumeIdent() string {
	var s string
	for isNotEof() {
		if !isIdentRune(userInput[wat]) {
			break
		}
		s += string(userInput[wat])
		lat++
		wat++
	}
	return s
}

func consumeString() (string, bool) {
	var s string
	isRaw := false
	if userInput[wat] == '`' {
		isRaw = true
	}
	// " / `
	lat++
	wat++

	for isNotEof() {
		if isRaw && userInput[wat] == '`' {
			break
		} else if !isRaw && userInput[wat] == '"' {
			break
		}

		if userInput[wat] == '\\' && userInput[wat+1] == '"' {
			s += "\""
			lat += 2
			wat += 2
			continue
		}

		if userInput[wat] == '\\' && userInput[wat+1] == 'n' {
			s += "\n"
			lat += 2
			wat += 2
			continue
		}

		if userInput[wat] == '\\' && userInput[wat+1] == 't' {
			s += "\t"
			lat += 2
			wat += 2
			continue
		}

		if userInput[wat] == '\\' && userInput[wat+1] == '\'' {
			s += "'"
			lat += 2
			wat += 2
			continue
		}

		if userInput[wat] == '\\' && userInput[wat+1] == '\\' {
			s += "\\"
			lat += 2
			wat += 2
			continue
		}

		s += string(userInput[wat])
		lat++
		wat++
	}
	// " / `
	lat++
	wat++

	return s, isRaw
}

func consumeNumber() (string, bool) {
	isFloat := false
	var s string
	for isNotEof() {
		if unicode.IsDigit(userInput[wat]) {
			s += string(userInput[wat])
			lat++
			wat++
			continue
		} else if userInput[wat] == '.' {
			if len(userInput) <= wat+1 || !unicode.IsDigit(userInput[wat+1]) {
				break
			}
			s += string(userInput[wat])
			lat++
			wat++
			isFloat = true
			continue
		} else {
			break
		}
	}

	return s, isFloat
}

func consumeWhite() string {
	var s string
	for isNotEof() {
		if userInput[wat] == ' ' || userInput[wat] == '\t' {
			s += string(userInput[wat])
			lat++
			wat++
		} else {
			break
		}
	}
	return s
}

func Tokenize(input string) (*Token, error) {
	// init
	userInput = []rune(input)
	lno = 1
	lat = 0
	wat = 0
	var head Token
	cur := &head

inputLoop:
	for isNotEof() {
		// white
		if userInput[wat] == ' ' || userInput[wat] == '\t' {
			_ = NewPosition(lno, lat, wat)
			_ = consumeWhite()
			// cur = NewWhite(cur, pos, s)
			continue
		}
		// newline
		if userInput[wat] == '\n' {
			// cur = NewNL(cur, NewPosition(lno, lat, wat), "\n")
			lno++
			lat = 0
			wat++
			continue
		}
		// comment
		if userInput[wat] == '#' {
			pos := NewPosition(lno, lat, wat)
			s := consumeComment()
			cur = NewComment(cur, pos, s)
			continue
		}
		// op symbols
		for _, r := range append(compositeOpSymbols, singleOpSymbols...) {
			if startWith(r) {
				cur = NewOpSymbol(cur, NewPosition(lno, lat, wat), r)
				lat += len(r)
				wat += len(r)
				continue inputLoop
			}
		}
		// ident
		if isIdentRune(userInput[wat]) && !unicode.IsDigit(userInput[wat]) {
			pos := NewPosition(lno, lat, wat)
			id := consumeIdent()
			cur = NewIdent(cur, pos, id)
			continue
		}

		// string
		if userInput[wat] == '`' || userInput[wat] == '"' {
			pos := NewPosition(lno, lat, wat)
			str, isRaw := consumeString()
			cur = NewString(cur, pos, str, isRaw)
			continue
		}

		// number
		if unicode.IsDigit(userInput[wat]) {
			pos := NewPosition(lno, lat, wat)
			numStr, isFloat := consumeNumber()
			if isFloat {
				n, err := strconv.ParseFloat(numStr, 64)
				if err != nil {
					return nil, err
				}
				cur = NewFloat(cur, pos, n)
				continue
			} else {
				n, err := strconv.ParseInt(numStr, 10, 0)
				if err != nil {
					return nil, err
				}
				cur = NewInt(cur, pos, int(n))
				continue
			}
		}

		return nil, NewUnexpectedCharacterErr(userInput[wat], NewPosition(lno, lat, wat))
	}

	cur = NewEof(cur, NewPosition(lno, lat, wat))
	return head.Next, nil
}
