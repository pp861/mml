package main

import (

	"mml/parser"
	"os"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// 自定义一个错误监听器，用于捕获语法错误
type ErrorListener struct {
	*antlr.DefaultErrorListener
	HasSyntaxError bool
	line int
	column int
	msg string
}

func (el *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	el.HasSyntaxError = true
	el.line = line 
	el.column = column
	el.msg = msg
}

func main() {

	csvFile := os.Args[1]
	is, err := antlr.NewFileStream(csvFile)
	if err != nil {
		fmt.Printf("new file stream error: %s\n", err)
		return
	}

	// is := antlr.NewInputStream("DEL SIPLBS_LOCAL_CONFIG ne= 10753 , Index = 0 ;")

	lexer := parser.NewMmlLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewMmlParser(stream)

	listener := &ErrorListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(listener)

	p.Prog()

	if listener.HasSyntaxError {
		fmt.Printf("mml has error, line %d:%d %s", listener.line, listener.column, listener.msg);
		return 
	}

	fmt.Printf("mml is ok");
}