// lexer/lexer.go
package lexer

type Lexer struct {
	input string
	position int //current position in input
	readPosition int //current readng position in input
	ch byte //current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}