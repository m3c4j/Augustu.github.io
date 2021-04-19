package main

import (
	"fmt"
	"strings"
)

type Tokener interface {
	String() string
}

type Token struct {
	tag Tag
}

func (t Token) String() string {
	return string(rune(t.tag))
}

func newToken(r rune) Token {
	return Token{Tag(r)}
}

type Tag int

const (
	NUM   Tag = 256
	ID    Tag = 257
	TRUE  Tag = 258
	FALSE Tag = 259
)

type Num struct {
	Token
	value int
}

func (n Num) String() string {
	return fmt.Sprintf("%d", n.value)
}

func newNum(v int) Num {
	n := Num{}
	n.tag = NUM
	n.value = v
	return n
}

type Word struct {
	Token
	lexeme string
}

func (w Word) String() string {
	return w.lexeme
}

func newWord(tag Tag, lexeme string) Word {
	w := Word{}
	w.tag = tag
	w.lexeme = lexeme
	return w
}

type Lexer struct {
	line  int
	peek  rune
	words map[string]Word
	s     string
	len   int
	cur   int
	fin   bool
}

func (l *Lexer) reserve(t Word) {
	l.words[t.lexeme] = t
}

func (l *Lexer) init() {
	l.reserve(newWord(TRUE, "true"))
	l.reserve(newWord(FALSE, "false"))
}

func (l *Lexer) next() bool {
	if l.cur < l.len-1 {
		l.cur++
		l.peek = rune(l.s[l.cur])
		return true
	}

	l.peek = ' '
	l.fin = true
	return false
}

func (l *Lexer) scan() (Tokener, bool) {
	for l.next() {
		if l.peek == ' ' || l.peek == '\t' {
			continue
		} else if l.peek == '\n' {
			l.line++
		} else {
			break
		}
	}

	if isDigit(l.peek) {
		var v int = 0
		for isDigit(l.peek) {
			v = 10*v + digit(l.peek)
			l.next()
		}
		return newNum(v), false
	}

	if isLetter(l.peek) {
		var b strings.Builder
		for isLetter(l.peek) {
			b.WriteString(string(l.peek))
			l.next()
		}

		s := b.String()
		w, ok := l.words[s]
		if ok {
			return w, false
		}

		w = newWord(ID, s)
		l.words[s] = w
		return w, false
	}

	t := newToken(l.peek)
	l.peek = ' '
	return t, l.fin
}

func NewLexer(s string) *Lexer {
	l := &Lexer{
		line:  1,
		peek:  ' ',
		words: make(map[string]Word),
		s:     s,
		len:   len(s),
		cur:   -1,
	}
	l.init()
	return l
}
