package lexer

import (
	"fmt"
	"testing"
	"token"
)

type TestToken struct {
	expectedType	token.TokenType
	expectedLiteral	string
}

func runTests( tests []TestToken, l *Lexer , t *testing.T){
	for i, tt := range tests {
		tok := l.NextToken()
		fmt.Println(tok)
		if tok.Type != tt.expectedType{
			t.Fatalf("Tests[%d] - token type wrong expected = %q, got = %q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral{
			t.Fatalf("Tests[%d] - literal wrong expected = %q, got %q",
				1, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken (t *testing.T){
	input := "#hello this is a test\n##"

	tests := []TestToken{
		{token.H1, "hello this is a test"},
		{token.NEW_LINE, " "},
		{token.H2, ""},
	}
	l := New(input)
	runTests(tests, l, t)
}


func TestUnorderedList(t *testing.T){
	input := "#This is a heading 1\n- this is an unordered list\n- this is another UL"
	tests := []TestToken{
		{token.H1,				"This is a heading 1"},
		{token.NEW_LINE,		" "},
		{token.UNORDERED_LIST,	"this is an unordered list"},
		{token.NEW_LINE,		" "},
		{token.UNORDERED_LIST,	"this is another UL"},
	}
	l := New(input)
	runTests(tests, l, t)
}

func TestListWithModifier(t *testing.T){
	input := "- this is a list **with bold**\n"

	tests := []TestToken{
		{token.UNORDERED_LIST, "this is a list "},
		{token.BOLD, "with bold"},
		{token.NEW_LINE, " "}, 
	}
	l := New(input)
	runTests(tests, l, t)
}

func TestListWithItalic( t *testing.T){
	input := "- This is a list *with italic*\n"
	tests := []TestToken {
		{expectedType: token.UNORDERED_LIST, expectedLiteral: "This is a list "},
		{expectedType: token.ITALIC, expectedLiteral: "with italic"},
	}
	l := New(input)
	runTests(tests, l, t)
}

func TestPlainText( t *testing.T){
	input := "This should just be a plain text block\n"
	tests := []TestToken {
		{expectedType: token.TEXT, expectedLiteral: "This should just be a plain text block"},
		{expectedType: token.NEW_LINE, expectedLiteral: " "},
	}
	l := New(input)
	runTests(tests, l, t)
}

func TestPlainTextWithModifiers(t *testing.T){
	input := "This should be a plain * text block\n"
	tests := []TestToken {
		{expectedType: token.TEXT, expectedLiteral: "This should be a plain * text block"},
		{expectedType: token.NEW_LINE, expectedLiteral: " "},
	}	
	l := New(input)
	runTests(tests, l, t)
}

func TestBoldList( t *testing.T){
	input := "- **Bold List**"
	tests := []TestToken {
		{expectedType: token.UNORDERED_LIST, expectedLiteral: ""},
		{expectedType: token.BOLD, expectedLiteral: "Bold List"},
	}
	l:= New(input)
	runTests(tests, l, t)
}









