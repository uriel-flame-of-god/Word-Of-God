package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ============================
// HOLY TOKEN DEFINITIONS
// ============================
type TokenType int

const (
	TOKEN_EOF TokenType = iota
	TOKEN_AND_GOD_SAID
	TOKEN_LET_THERE_BE
	TOKEN_THOU_SHALT
	TOKEN_AND
	TOKEN_YE_SHALL_NOT
	TOKEN_VERILY_VERILY
	TOKEN_BEHOLD
	TOKEN_WOE_UNTO
	TOKEN_AND_IT_CAME_TO_PASS
	TOKEN_GO_YE_UNTO
	TOKEN_IDENTIFIER
	TOKEN_NUMBER
	TOKEN_STRING
	TOKEN_COLON
	TOKEN_COMMA
)

func (t TokenType) String() string {
	names := map[TokenType]string{
		TOKEN_EOF:                 "EOF",
		TOKEN_AND_GOD_SAID:        "AND_GOD_SAID",
		TOKEN_LET_THERE_BE:        "LET_THERE_BE",
		TOKEN_THOU_SHALT:          "THOU_SHALT",
		TOKEN_AND:                 "AND",
		TOKEN_YE_SHALL_NOT:        "YE_SHALL_NOT",
		TOKEN_VERILY_VERILY:       "VERILY_VERILY",
		TOKEN_BEHOLD:              "BEHOLD",
		TOKEN_WOE_UNTO:            "WOE_UNTO",
		TOKEN_AND_IT_CAME_TO_PASS: "AND_IT_CAME_TO_PASS",
		TOKEN_GO_YE_UNTO:          "GO_YE_UNTO",
		TOKEN_IDENTIFIER:          "IDENTIFIER",
		TOKEN_NUMBER:              "NUMBER",
		TOKEN_STRING:              "STRING",
		TOKEN_COLON:               "COLON",
		TOKEN_COMMA:               "COMMA",
	}
	if name, ok := names[t]; ok {
		return name
	}
	return fmt.Sprintf("UNKNOWN(%d)", t)
}

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal interface{}
	Line    int
}

// ============================
// AST NODES
// ============================
type NodeType int

const (
	NODE_PROGRAM NodeType = iota
	NODE_DECLARATION
	NODE_ASSIGNMENT
	NODE_BINARY_OP
	NODE_UNARY_OP
	NODE_LITERAL
	NODE_VARIABLE
	NODE_PRINT
	NODE_JUMP
	NODE_CONDITIONAL_JUMP
	NODE_ERROR
	NODE_LABEL
	NODE_RETURN
)

type Node struct {
	Type     NodeType
	Value    interface{}
	Children []*Node
	Line     int
}

// ============================
// SYMBOL TABLE
// ============================
type SymbolType int

const (
	SYM_VOID SymbolType = iota
	SYM_BOOL
	SYM_INT
	SYM_STRING
	SYM_FLOAT
	SYM_DOUBLE
)

type Symbol struct {
	Name    string
	SymType SymbolType
	Lineage int
	IsLabel bool
	LabelPC int
}

// ============================
// LEXER
// ============================
type Lexer struct {
	source  string
	tokens  []Token
	start   int
	current int
	line    int
}

func NewLexer(source string) *Lexer {
	return &Lexer{
		source: source,
		line:   1,
	}
}

func (l *Lexer) ScanTokens() []Token {
	for !l.isAtEnd() {
		l.start = l.current
		l.scanToken()
	}

	l.tokens = append(l.tokens, Token{TOKEN_EOF, "", nil, l.line})
	return l.tokens
}

func (l *Lexer) scanToken() {
	c := l.advance()

	switch c {
	case ':':
		l.addToken(TOKEN_COLON, nil)
	case ',':
		l.addToken(TOKEN_COMMA, nil)
	case ' ', '\r', '\t':
		// Ignore whitespace
	case '\n':
		l.line++
	case '"':
		l.string()
	case '/':
		if l.match('/') {
			for l.peek() != '\n' && !l.isAtEnd() {
				l.advance()
			}
		}
	default:
		if l.isDigit(c) {
			l.number()
		} else if l.isAlpha(c) {
			l.identifier()
		}
	}
}

func (l *Lexer) identifier() {
	start := l.current - 1
	remaining := l.source[start:]
	remainingLower := strings.ToLower(remaining)

	// Check longest keywords first to avoid partial matches
	switch {
	case strings.HasPrefix(remainingLower, "and it came to pass"):
		l.current = start + len("and it came to pass")
		l.addToken(TOKEN_AND_IT_CAME_TO_PASS, nil)
	case strings.HasPrefix(remainingLower, "and god said"):
		l.current = start + len("and god said")
		l.addToken(TOKEN_AND_GOD_SAID, nil)
	case strings.HasPrefix(remainingLower, "let there be"):
		l.current = start + len("let there be")
		l.addToken(TOKEN_LET_THERE_BE, nil)
	case strings.HasPrefix(remainingLower, "thou shalt"):
		l.current = start + len("thou shalt")
		l.addToken(TOKEN_THOU_SHALT, nil)
	case strings.HasPrefix(remainingLower, "ye shall not"):
		l.current = start + len("ye shall not")
		l.addToken(TOKEN_YE_SHALL_NOT, nil)
	case strings.HasPrefix(remainingLower, "verily, verily"):
		l.current = start + len("verily, verily")
		l.addToken(TOKEN_VERILY_VERILY, nil)
	case strings.HasPrefix(remainingLower, "go ye unto"):
		l.current = start + len("go ye unto")
		l.addToken(TOKEN_GO_YE_UNTO, nil)
	case strings.HasPrefix(remainingLower, "woe unto"):
		l.current = start + len("woe unto")
		l.addToken(TOKEN_WOE_UNTO, nil)
	case strings.HasPrefix(remainingLower, "behold"):
		l.current = start + len("behold")
		l.addToken(TOKEN_BEHOLD, nil)
	case strings.HasPrefix(remainingLower, "and"):
		if start+3 >= len(l.source) || !l.isAlpha(l.source[start+3]) {
			l.current = start + 3
			l.addToken(TOKEN_AND, nil)
		} else {
			// Part of a longer identifier
			for l.isAlphaNumeric(l.peek()) {
				l.advance()
			}
			l.addToken(TOKEN_IDENTIFIER, l.source[start:l.current])
		}
	default:
		for l.isAlphaNumeric(l.peek()) {
			l.advance()
		}
		l.addToken(TOKEN_IDENTIFIER, l.source[start:l.current])
	}
}

func (l *Lexer) number() {
	for l.isDigit(l.peek()) {
		l.advance()
	}

	if l.peek() == '.' && l.isDigit(l.peekNext()) {
		l.advance()
		for l.isDigit(l.peek()) {
			l.advance()
		}
	}

	value, _ := strconv.ParseFloat(l.source[l.start:l.current], 64)
	l.addToken(TOKEN_NUMBER, value)
}

func (l *Lexer) string() {
	for l.peek() != '"' && !l.isAtEnd() {
		if l.peek() == '\n' {
			l.line++
		}
		l.advance()
	}

	if l.isAtEnd() {
		fmt.Printf("Gabriel proclaims: Line %d: Unterminated string.\n", l.line)
		return
	}

	l.advance() // Closing "
	value := l.source[l.start+1 : l.current-1]
	l.addToken(TOKEN_STRING, value)
}

func (l *Lexer) isAtEnd() bool {
	return l.current >= len(l.source)
}

func (l *Lexer) advance() byte {
	c := l.source[l.current]
	l.current++
	return c
}

func (l *Lexer) peek() byte {
	if l.isAtEnd() {
		return 0
	}
	return l.source[l.current]
}

func (l *Lexer) peekNext() byte {
	if l.current+1 >= len(l.source) {
		return 0
	}
	return l.source[l.current+1]
}

func (l *Lexer) match(expected byte) bool {
	if l.isAtEnd() || l.source[l.current] != expected {
		return false
	}
	l.current++
	return true
}

func (l *Lexer) addToken(tokenType TokenType, literal interface{}) {
	text := l.source[l.start:l.current]
	l.tokens = append(l.tokens, Token{tokenType, text, literal, l.line})
}

func (l *Lexer) isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func (l *Lexer) isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func (l *Lexer) isAlphaNumeric(c byte) bool {
	return l.isAlpha(c) || l.isDigit(c)
}

// ============================
// PARSER
// ============================
type Parser struct {
	tokens       []Token
	current      int
	symbols      map[string]*Symbol
	errors       []string
	seenEntry    bool
	labelLineage map[string]int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{
		tokens:       tokens,
		symbols:      make(map[string]*Symbol),
		labelLineage: make(map[string]int),
	}
}

func (p *Parser) Parse() *Node {
	program := &Node{
		Type:     NODE_PROGRAM,
		Children: make([]*Node, 0),
	}

	if !p.match(TOKEN_AND_GOD_SAID) {
		p.error("Program must begin with 'And God said'")
		return program
	}
	p.seenEntry = true

	for !p.isAtEnd() && !p.check(TOKEN_AND_IT_CAME_TO_PASS) {
		stmt := p.statement()
		if stmt != nil {
			program.Children = append(program.Children, stmt)
		}
	}

	if !p.match(TOKEN_AND_IT_CAME_TO_PASS) {
		p.error("Program must end with 'And it came to pass'")
	}

	return program
}

func (p *Parser) statement() *Node {
	if p.match(TOKEN_LET_THERE_BE) {
		return p.varDeclaration()
	}

	if p.match(TOKEN_BEHOLD) {
		return p.printStatement()
	}

	if p.match(TOKEN_THOU_SHALT) {
		return p.assignment()
	}

	if p.match(TOKEN_GO_YE_UNTO) {
		return p.jumpStatement()
	}

	if p.match(TOKEN_WOE_UNTO) {
		return p.errorStatement()
	}

	if p.match(TOKEN_YE_SHALL_NOT) {
		return p.negationOrCondJump()
	}

	// Check for labels (identifier followed by number:number)
	if p.check(TOKEN_IDENTIFIER) {
		checkpoint := p.current
		name := p.advance().Lexeme

		if p.match(TOKEN_NUMBER) {
			typeNum := int(p.previous().Literal.(float64))
			if p.match(TOKEN_COLON) && p.match(TOKEN_NUMBER) {
				lineage := int(p.previous().Literal.(float64))

				// Check lineage constraint
				if prev, exists := p.labelLineage[name]; exists {
					if lineage <= prev {
						p.error(fmt.Sprintf("Lineage violation: %s must increment from %d, got %d",
							name, prev, lineage))
					}
				}

				p.labelLineage[name] = lineage
				p.symbols[name] = &Symbol{
					Name:    name,
					SymType: SymbolType(typeNum % 6),
					Lineage: lineage,
					IsLabel: true,
				}

				return &Node{
					Type:  NODE_LABEL,
					Value: name,
					Line:  p.previous().Line,
				}
			}
		}

		// Not a label, backtrack
		p.current = checkpoint
	}

	return nil
}

func (p *Parser) varDeclaration() *Node {
	if !p.check(TOKEN_IDENTIFIER) {
		p.error("Expected variable name after 'Let there be'")
		return nil
	}

	name := p.advance().Lexeme

	if !p.match(TOKEN_NUMBER) {
		p.error("Expected type number")
		return nil
	}
	typeNum := int(p.previous().Literal.(float64))

	if !p.match(TOKEN_COLON) {
		p.error("Expected ':'")
		return nil
	}

	if !p.match(TOKEN_NUMBER) {
		p.error("Expected lineage number")
		return nil
	}
	lineage := int(p.previous().Literal.(float64))

	p.symbols[name] = &Symbol{
		Name:    name,
		SymType: SymbolType(typeNum % 6),
		Lineage: lineage,
		IsLabel: false,
	}

	return &Node{
		Type: NODE_DECLARATION,
		Value: map[string]interface{}{
			"name":    name,
			"type":    typeNum,
			"lineage": lineage,
		},
		Line: p.previous().Line,
	}
}

func (p *Parser) assignment() *Node {
	if !p.check(TOKEN_IDENTIFIER) {
		p.error("Expected variable name")
		return nil
	}

	name := p.advance().Lexeme

	var expr *Node
	if !p.isAtEnd() && !p.check(TOKEN_LET_THERE_BE) &&
		!p.check(TOKEN_BEHOLD) && !p.check(TOKEN_THOU_SHALT) &&
		!p.check(TOKEN_GO_YE_UNTO) && !p.check(TOKEN_WOE_UNTO) &&
		!p.check(TOKEN_YE_SHALL_NOT) && !p.check(TOKEN_AND_IT_CAME_TO_PASS) &&
		(p.check(TOKEN_IDENTIFIER) || p.check(TOKEN_NUMBER) ||
			p.check(TOKEN_STRING) || p.check(TOKEN_VERILY_VERILY) ||
			p.check(TOKEN_AND)) {
		expr = p.expression()
	} else {
		expr = &Node{Type: NODE_LITERAL, Value: int64(0)}
	}

	return &Node{
		Type:     NODE_ASSIGNMENT,
		Value:    name,
		Children: []*Node{expr},
		Line:     p.previous().Line,
	}
}

func (p *Parser) printStatement() *Node {
	expr := p.expression()
	return &Node{
		Type:     NODE_PRINT,
		Children: []*Node{expr},
		Line:     p.previous().Line,
	}
}

func (p *Parser) jumpStatement() *Node {
	if !p.check(TOKEN_IDENTIFIER) {
		p.error("Expected label name")
		return nil
	}

	label := p.advance().Lexeme
	return &Node{
		Type:  NODE_JUMP,
		Value: label,
		Line:  p.previous().Line,
	}
}

func (p *Parser) errorStatement() *Node {
	if !p.match(TOKEN_STRING) {
		p.error("Expected string after 'Woe unto'")
		return nil
	}

	return &Node{
		Type:  NODE_ERROR,
		Value: p.previous().Literal.(string),
		Line:  p.previous().Line,
	}
}

func (p *Parser) negationOrCondJump() *Node {
	expr := p.expression()

	if p.match(TOKEN_COMMA) {
		if !p.match(TOKEN_GO_YE_UNTO) {
			p.error("Expected 'Go ye unto' after comma")
			return nil
		}

		if !p.check(TOKEN_IDENTIFIER) {
			p.error("Expected label name")
			return nil
		}

		label := p.advance().Lexeme

		return &Node{
			Type:     NODE_CONDITIONAL_JUMP,
			Children: []*Node{expr},
			Value:    label,
			Line:     p.previous().Line,
		}
	}

	return &Node{
		Type:     NODE_UNARY_OP,
		Value:    "NOT",
		Children: []*Node{expr},
		Line:     p.previous().Line,
	}
}

func (p *Parser) expression() *Node {
	if p.check(TOKEN_AND) {
		p.advance()
		if !p.check(TOKEN_IDENTIFIER) && !p.check(TOKEN_NUMBER) &&
			!p.check(TOKEN_STRING) && !p.check(TOKEN_VERILY_VERILY) {
			p.error("Expected expression after 'And'")
			return &Node{Type: NODE_LITERAL, Value: int64(0)}
		}

		expr := &Node{Type: NODE_LITERAL, Value: int64(0)}
		right := p.primary()
		expr = &Node{
			Type:     NODE_BINARY_OP,
			Value:    "ADD",
			Children: []*Node{expr, right},
			Line:     p.previous().Line,
		}

		for p.match(TOKEN_AND) {
			if !p.check(TOKEN_IDENTIFIER) && !p.check(TOKEN_NUMBER) &&
				!p.check(TOKEN_STRING) && !p.check(TOKEN_VERILY_VERILY) {
				p.error("Expected expression after 'And'")
				break
			}
			right := p.primary()
			expr = &Node{
				Type:     NODE_BINARY_OP,
				Value:    "ADD",
				Children: []*Node{expr, right},
				Line:     p.previous().Line,
			}
		}
		return expr
	}

	expr := p.primary()

	for p.match(TOKEN_AND) {
		if !p.check(TOKEN_IDENTIFIER) && !p.check(TOKEN_NUMBER) &&
			!p.check(TOKEN_STRING) && !p.check(TOKEN_VERILY_VERILY) {
			p.error("Expected expression after 'And'")
			break
		}
		right := p.primary()
		expr = &Node{
			Type:     NODE_BINARY_OP,
			Value:    "ADD",
			Children: []*Node{expr, right},
			Line:     p.previous().Line,
		}
	}

	return expr
}

func (p *Parser) primary() *Node {
	if p.match(TOKEN_VERILY_VERILY) {
		return &Node{
			Type:  NODE_LITERAL,
			Value: int64(1),
			Line:  p.previous().Line,
		}
	}

	if p.match(TOKEN_NUMBER) {
		return &Node{
			Type:  NODE_LITERAL,
			Value: int64(p.previous().Literal.(float64)),
			Line:  p.previous().Line,
		}
	}

	if p.match(TOKEN_STRING) {
		return &Node{
			Type:  NODE_LITERAL,
			Value: p.previous().Literal.(string),
			Line:  p.previous().Line,
		}
	}

	if p.match(TOKEN_IDENTIFIER) {
		name := p.previous().Lexeme
		return &Node{
			Type:  NODE_VARIABLE,
			Value: name,
			Line:  p.previous().Line,
		}
	}

	if !p.isAtEnd() {
		p.error(fmt.Sprintf("Expected expression, got %s",
			p.peek().Type))
	} else {
		p.error("Expected expression")
	}
	return &Node{Type: NODE_LITERAL, Value: int64(0)}
}

func (p *Parser) match(types ...TokenType) bool {
	for _, t := range types {
		if p.check(t) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) check(t TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == t
}

func (p *Parser) advance() Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == TOKEN_EOF
}

func (p *Parser) peek() Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() Token {
	return p.tokens[p.current-1]
}

func (p *Parser) error(message string) {
	p.errors = append(p.errors, fmt.Sprintf("Line %d: %s", p.peek().Line, message))
}

// ============================
// INTERPRETER
// ============================
type Interpreter struct {
	ast     *Node
	symbols map[string]*Symbol
	memory  map[string]int64
	strings map[string]string
	labels  map[string]int
	pc      int
	stmts   []*Node
}

func NewInterpreter(ast *Node, symbols map[string]*Symbol) *Interpreter {
	return &Interpreter{
		ast:     ast,
		symbols: symbols,
		memory:  make(map[string]int64),
		strings: make(map[string]string),
		labels:  make(map[string]int),
	}
}

func (i *Interpreter) Execute() error {
	i.stmts = i.ast.Children

	// First pass: collect labels
	for idx, stmt := range i.stmts {
		if stmt.Type == NODE_LABEL {
			i.labels[stmt.Value.(string)] = idx
		}
	}

	i.pc = 0
	for i.pc < len(i.stmts) {
		if err := i.executeNode(i.stmts[i.pc]); err != nil {
			return err
		}
		i.pc++
	}

	return nil
}

func (i *Interpreter) executeNode(node *Node) error {
	switch node.Type {
	case NODE_DECLARATION:
		return nil

	case NODE_ASSIGNMENT:
		val, err := i.evaluate(node.Children[0])
		if err != nil {
			return fmt.Errorf("Line %d: %v", node.Line, err)
		}
		i.memory[node.Value.(string)] = val

	case NODE_PRINT:
		val, err := i.evaluate(node.Children[0])
		if err != nil {
			return fmt.Errorf("Line %d: %v", node.Line, err)
		}
		if str, ok := i.strings[fmt.Sprintf("%d", val)]; ok {
			fmt.Println(str)
		} else {
			fmt.Println(val)
		}

	case NODE_JUMP:
		label := node.Value.(string)
		pc, ok := i.labels[label]
		if !ok {
			return fmt.Errorf("Line %d: Unknown label '%s'",
				node.Line, label)
		}
		i.pc = pc - 1

	case NODE_CONDITIONAL_JUMP:
		val, err := i.evaluate(node.Children[0])
		if err != nil {
			return fmt.Errorf("Line %d: %v", node.Line, err)
		}
		if val == 0 {
			label := node.Value.(string)
			pc, ok := i.labels[label]
			if !ok {
				return fmt.Errorf("Line %d: Unknown label '%s'",
					node.Line, label)
			}
			i.pc = pc - 1
		}

	case NODE_ERROR:
		return fmt.Errorf("Line %d: %s",
			node.Line, node.Value.(string))

	case NODE_LABEL:
		// Do nothing, labels are markers

	case NODE_RETURN:
		i.pc = len(i.stmts)
	}

	return nil
}

func (i *Interpreter) evaluate(node *Node) (int64, error) {
	switch node.Type {
	case NODE_LITERAL:
		switch v := node.Value.(type) {
		case int64:
			return v, nil
		case string:
			// Use negative refs for strings: -1, -2, -3, etc.
			ref := int64(-len(i.strings) - 1)
			i.strings[fmt.Sprintf("%d", ref)] = v
			return ref, nil
		default:
			return 0, nil
		}

	case NODE_VARIABLE:
		name := node.Value.(string)
		if val, ok := i.memory[name]; ok {
			return val, nil
		}
		return 0, fmt.Errorf("undefined variable: '%s'", name)

	case NODE_BINARY_OP:
		left, err := i.evaluate(node.Children[0])
		if err != nil {
			return 0, err
		}
		right, err := i.evaluate(node.Children[1])
		if err != nil {
			return 0, err
		}
		return left + right, nil

	case NODE_UNARY_OP:
		val, err := i.evaluate(node.Children[0])
		if err != nil {
			return 0, err
		}
		if val == 0 {
			return 1, nil
		}
		return 0, nil
	}

	return 0, nil
}

// ============================
// MAIN
// ============================
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gabriel <scripture.wog>")
		os.Exit(1)
	}

	source, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error: Cannot read scripture: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("GABRIEL interpreting %s...\n\n", os.Args[1])

	// Lexical Analysis
	lexer := NewLexer(string(source))
	tokens := lexer.ScanTokens()

	fmt.Println("Tokens received:")
	for _, token := range tokens {
		if token.Type != TOKEN_EOF {
			fmt.Printf("   Line %d: %s -> '%s'\n", token.Line, token.Type, token.Lexeme)
		}
	}

	// Parsing
	parser := NewParser(tokens)
	ast := parser.Parse()

	if len(parser.errors) > 0 {
		fmt.Println("\nParsing Errors:")
		for _, e := range parser.errors {
			fmt.Printf("   %s\n", e)
		}
		os.Exit(1)
	}

	fmt.Printf("\nAST generated with %d nodes\n", countNodes(ast))

	// Interpretation
	interpreter := NewInterpreter(ast, parser.symbols)

	fmt.Printf("\nExecuting divine commands...\n\n")

	if err := interpreter.Execute(); err != nil {
		fmt.Printf("\nRuntime error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\nExecution complete.")
}

func countNodes(node *Node) int {
	if node == nil {
		return 0
	}
	count := 1
	for _, child := range node.Children {
		count += countNodes(child)
	}
	return count
}
