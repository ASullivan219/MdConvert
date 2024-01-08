// Package lexer tokenizes markdown files 
//
// Provide a markdown file as input and receive a stream of tokens
package lexer

import (
	"fmt"
	"token"
)

// A Lexer holds information about the current position in
// the input file
type Lexer struct {
	input			string
	position		int
	readPosition	int
	ch				byte
	prevToken		*token.Token
}


// Create a new lexer from the file input
func New( input string ) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l;
}


func (l *Lexer ) readChar(){

	if l.readPosition >= len(l.input){
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}


// Retrieve the next token from the lexer
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	switch l.ch {
	case '\n':
		tok = newTokenStringLiteral(token.NEW_LINE, " ")
	case '#':
		if l.prevToken == nil || l.prevToken.Type == token.NEW_LINE{
			headerLevel, literal := l.readHeader()
			tok = createHeaderToken(headerLevel, literal)
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
		return tok
	case '-':
		if l.prevToken == nil || l.prevToken.Type == token.NEW_LINE{
			literal := l.readListItem()
			tok = newTokenStringLiteral(token.UNORDERED_LIST, literal)
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
		return tok
	case '*':
		return l.readModifier()
	default:
		return l.readText()
	}

	l.prevToken = &tok
	l.readChar()
	return tok
}

func (l * Lexer) readHeader() (int,string){
	headerLevel := 0
	for l.ch == '#' && l.ch != 0{
		headerLevel += 1
		l.readChar()
	}

	position := l.position
	for !isNewLine(l.ch) && l.ch != 0{
		l.readChar()
	}
	return headerLevel, l.input[position : l.position]
}

func (l *Lexer) readListItem() string{
	for l.ch == '-' || l.ch == ' '{
		l.readChar()
	}
	position := l.position
	for !l.isModifier() && l.ch != 0  && !isNewLine(l.ch){
		l.readChar()
	}
	return l.input[position : l.position]
}

func (l *Lexer) readModifier() token.Token{
	l.readChar()

	switch l.ch{
	case ' ':
		return l.readText()
	case '*':
		l.readChar()
		return l.parseBoldText()

	default:
		return l.parseItalic()

	}
}

func (l *Lexer) parseBoldText() token.Token{
	fmt.Println("trying to parse bold text")
	position := l.position
	endPosition := l.position
	for !isNewLine(l.ch) && l.ch != 0{
		if ( l.ch == '*' ){
			endPosition = l.position
			l.readChar()
			if (l.ch == '*'){
				break
			}
		}
		l.readChar()
	}
	if l.ch != '*' {
		return newTokenStringLiteral(token.ILLEGAL, "did not receive matching closing modifier")
	}
	l.readChar()
	return newTokenStringLiteral(token.BOLD, l.input[position : endPosition])
}

func (l *Lexer) parseItalic() token.Token{
	position := l.position
	for !l.isModifier() &&  !l.isEof() && !l.isNewline(){
		l.readChar()
	}
	return newTokenStringLiteral(token.ITALIC, l.input[position : l.position])
}

func (l *Lexer) readText() token.Token{
	position := l.position
	for  !l.isEof() && !l.isNewline() && !l.isModifier(){
		l.readChar()
	}

	return newTokenStringLiteral(token.TEXT, l.input[position: l.position])
}


func (l *Lexer) isModifier() bool {
	return l.ch == '*'
}

func (l *Lexer) isNewline() bool {
	return l.ch == '\n'
}

func (l *Lexer) isEof() bool {
	return l.ch == 0
}

func newTokenStringLiteral(tokenType token.TokenType, literal string) token.Token{
	return token.Token{Type: tokenType, Literal: literal}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func createHeaderToken(headerLevel int, literal string) token.Token{
	var tok token.Token
	switch headerLevel {
	case 1:
		tok = newTokenStringLiteral(token.H1, literal)
	case 2:
		tok = newTokenStringLiteral(token.H2, literal)
	case 3:
		tok = newTokenStringLiteral(token.H3, literal)
	default:
		tok = newTokenStringLiteral(token.ILLEGAL, "Too many headerlevels")
	}
	return tok
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isNewLine(ch byte) bool{
	return ch == '\n'
}

func isReservedChar(ch byte) bool {
	return (
		ch == '*' || 
		ch == '_' || 
		ch == '-' || 
		ch == '`' || 
		ch == '#' ||
		ch == '\n' || 
		ch == 0)
}






