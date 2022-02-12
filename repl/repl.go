package repl

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		lex := lexer.New(line)

		for tok := lex.NextToken(); tok.Type != token.EOF; tok = lex.NextToken() {
			var buf bytes.Buffer
			_, err := fmt.Fprintf(out, "%+v\n", tok)
			if err != nil {
				return
			}
			_, err = out.Write(buf.Bytes())
			if err != nil {
				return
			}
		}
	}
}
