package main

import (
	"fmt"
	"hael/evaluator"
	"hael/lexer"
	"hael/parser"
	"io"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Erro: Nenhum comando fornecido.")
		fmt.Println("Uso: hael <comando> <caminho_do_arquivo>")
		os.Exit(1)
	}

	// Obter o comando (primeiro argumento após o nome do programa)
	command := os.Args[1]

	// Verificar se há um argumento para o caminho do arquivo
	if len(os.Args) < 3 {
		fmt.Println("Erro: O caminho do arquivo é obrigatório.")
		fmt.Println("Uso: hael <comando> <caminho_do_arquivo>")
		os.Exit(1)
	}

	// Obter o caminho do arquivo (argumento após o comando)
	filePath := os.Args[2]

	// Resolvendo o caminho relativo para o caminho absoluto
	// filepath.Abs lida automaticamente com ./test ou .\test no Windows
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		fmt.Printf("Erro ao resolver o caminho do arquivo: %v\n", err)
		os.Exit(1)
	}

	// Verificando se o arquivo existe
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		fmt.Printf("Erro: O arquivo %s não existe.\n", absPath)
		os.Exit(1)
	}

	switch command {
	case "run":
		// fmt.Printf("\n      __             __\n   .-'.'     .-.     '.'-.\n .'.((      ( ^ `>     )).'. \n/`'- \\'._____\\ (_____.'/ -'`\\\n|-''`.'------' '------'.`''-|\n|.-'`.'.'.`/ | | \\`.'.'.`'-.|\n \\ .' . /  | | | |  \\ . '. /\n  '._. :  _|_| |_|_  : ._.'\n     ````` /T\"Y\"T\\ `````\n          / | | | \\\n         `'`'`'`'`'`\n\n")
		evalAndRun(filePath)
	default:
		fmt.Printf("Comando desconhecido: %s\n", command)
		fmt.Println("Comandos disponíveis: run")
		os.Exit(1)
	}
}

func evalAndRun(path string) {
	content, err := os.ReadFile(path)
	out := os.Stdout

	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	input := string(content)

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
	}

	evaluated := evaluator.Eval(program)

	if evaluated != nil {
		io.WriteString(out, ">>> ")
		io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")
	}

}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
