package main

import "fmt"

type Parser struct {
	lookahead rune
	s         string
	len       int
	cur       int
}

func (p *Parser) Parser(s string) {
	p.s = s
	p.len = len(s)
	p.cur = 0
	p.lookahead = rune(s[0])
}

func (p *Parser) next() {
	if p.cur < p.len-1 {
		p.cur++
	}
	p.lookahead = rune(p.s[p.cur])
}

func (p *Parser) expr() error {
	p.term()
	for {
		if p.lookahead == '+' {
			p.match('+')
			p.term()
			fmt.Print("+")
		} else if p.lookahead == '-' {
			p.match('-')
			p.term()
			fmt.Print("-")
		} else {
			return nil
		}
	}
}

func (p *Parser) term() error {
	if p.isDigit(p.lookahead) {
		fmt.Print(string(p.lookahead))
		p.match(p.lookahead)
	} else {
		panic("syntax error")
		// return fmt.Errorf("syntax error")
	}
	return nil
}

func (p *Parser) isDigit(i rune) bool {
	switch i {
	case '0':
		return true
	case rune('1'):
		return true
	case rune('2'):
		return true
	case rune('3'):
		return true
	case rune('4'):
		return true
	case rune('5'):
		return true
	case rune('6'):
		return true
	case rune('7'):
		return true
	case rune('8'):
		return true
	case rune('9'):
		return true
	default:
		return false
	}
}

func (p *Parser) match(t rune) error {
	if p.lookahead == t {
		p.next()
	} else {
		return fmt.Errorf("syntax error")
	}
	return nil
}
