# Gabriel Compiler Specification

## Overview

Gabriel is the reference compiler and interpreter for the Word of God esoteric programming language. This document specifies Gabriel's architecture, compilation stages, and runtime semantics.

## Scope

This specification defines:

- Lexical analysis requirements
- Parser behavior and AST structure
- Symbol table management
- Interpretation semantics
- Error reporting standards
- Extension points for future versions

## Version

Gabriel 1.0.0 (Reference Implementation)

---

## 1. Architecture

Gabriel operates as a three-stage compiler-interpreter:

```
Source Code
    |
    v
[LEXER] -> Token Stream
    |
    v
[PARSER] -> Abstract Syntax Tree (AST)
    |
    v
[INTERPRETER] -> Execution / Output
```

### 1.1 Execution Model

Gabriel uses:

- **Linear execution**: Statements execute sequentially unless interrupted by jumps
- **Label-based control flow**: Goto-style jumps to labeled statements
- **Flat memory model**: Single-level variable scope with no function calls
- **Immediate interpretation**: No optimization or compilation to bytecode

---

## 2. Lexical Analysis (Lexer)

### 2.1 Token Types

Gabriel recognizes the following token types:

```go
TOKEN_EOF                  // End of file
TOKEN_AND_GOD_SAID         // Program entry point
TOKEN_LET_THERE_BE         // Variable declaration
TOKEN_THOU_SHALT           // Assignment
TOKEN_AND                  // Addition operator
TOKEN_YE_SHALL_NOT         // Negation / conditional
TOKEN_VERILY_VERILY        // Boolean true (1)
TOKEN_BEHOLD               // Output statement
TOKEN_WOE_UNTO             // Error / termination
TOKEN_AND_IT_CAME_TO_PASS  // Program termination
TOKEN_GO_YE_UNTO           // Unconditional jump
TOKEN_IDENTIFIER           // Variable / label name
TOKEN_NUMBER               // Numeric literal
TOKEN_STRING               // String literal
TOKEN_COLON                // Delimiter (:)
TOKEN_COMMA                // Delimiter (,)
```

### 2.2 Keyword Matching

Multi-word keywords are matched **before** shorter keywords to prevent partial matches.

**Keyword priority** (longest first):

1. "And it came to pass"
2. "And God said"
3. "Let there be"
4. "Thou shalt"
5. "Ye shall not"
6. "Verily, verily"
7. "Go ye unto"
8. "Woe unto"
9. "Behold"
10. "And"

Keyword matching is **case-insensitive**. All keywords may appear in any capitalization (e.g., "and god said", "AND GOD SAID", "And God Said").

### 2.3 Identifiers

Identifiers follow standard programming language conventions:

- Start with letter (a-z, A-Z) or underscore (_)
- Contain letters, digits (0-9), or underscores
- Case-sensitive
- Maximum length: implementation-defined (Gabriel: no limit)

### 2.4 Literals

#### Numbers

Integers and floating-point numbers are recognized:

```
42          // Integer
3.14        // Float
0           // Zero
-5          // Negative (via And operator)
```

Numbers are parsed as 64-bit floats internally, then converted to int64 for storage.

#### Strings

Strings are enclosed in double quotes:

```
"Hello, World"
"Error message"
"Multi\nline"
```

String contents preserve all characters including whitespace and escape sequences.

### 2.5 Comments

Single-line comments begin with `//` and extend to end of line:

```
// This is a comment
Thou shalt x And 42  // Inline comment
```

### 2.6 Whitespace

Whitespace (spaces, tabs, carriage returns) is ignored except for line counting.

Newlines are tracked for error reporting but do not affect parsing.

---

## 3. Parsing (Parser)

### 3.1 Grammar (EBNF)

```ebnf
program         = TOKEN_AND_GOD_SAID statement* TOKEN_AND_IT_CAME_TO_PASS

statement       = declaration
                | assignment
                | label
                | print_stmt
                | jump_stmt
                | cond_jump_stmt
                | error_stmt

declaration     = TOKEN_LET_THERE_BE IDENTIFIER TOKEN_NUMBER TOKEN_COLON TOKEN_NUMBER

assignment      = TOKEN_THOU_SHALT IDENTIFIER expression?

label           = IDENTIFIER TOKEN_NUMBER TOKEN_COLON TOKEN_NUMBER

print_stmt      = TOKEN_BEHOLD expression

jump_stmt       = TOKEN_GO_YE_UNTO IDENTIFIER

cond_jump_stmt  = TOKEN_YE_SHALL_NOT expression TOKEN_COMMA TOKEN_GO_YE_UNTO IDENTIFIER

error_stmt      = TOKEN_WOE_UNTO TOKEN_STRING

expression      = (TOKEN_AND)? primary (TOKEN_AND primary)*

primary         = TOKEN_NUMBER
                | TOKEN_STRING
                | TOKEN_VERILY_VERILY
                | IDENTIFIER
```

### 3.2 AST Node Types

```go
NODE_PROGRAM          // Root node
NODE_DECLARATION      // Variable declaration
NODE_ASSIGNMENT       // Variable assignment
NODE_LABEL            // Jump target label
NODE_PRINT            // Output statement
NODE_JUMP             // Unconditional jump
NODE_CONDITIONAL_JUMP // Conditional jump
NODE_ERROR            // Error / termination
NODE_RETURN           // (Internal) end execution
NODE_LITERAL          // Value literal
NODE_VARIABLE         // Variable reference
NODE_BINARY_OP        // Binary operation (ADD)
NODE_UNARY_OP         // Unary operation (NOT)
```

### 3.3 Symbol Table

The parser maintains a symbol table tracking:

```go
Symbol {
    Name      string      // Variable/label name
    SymType   int         // Type annotation (0-5)
    Lineage   int         // Declaration order
    IsLabel   bool        // True if label
    LabelPC   int         // Program counter for label
}
```

#### 3.3.1 Lineage Validation

If a symbol is redeclared with the same name, its lineage must be strictly greater than the previous declaration:

```
Lineage[new] > Lineage[previous]  // Must be true
```

This prevents ambiguous symbol redeclaration.

#### 3.3.2 Type Annotations

Type values are modulo 6:

```
0 -> void
1 -> boolean
2 -> integer
3 -> string
4 -> float
5 -> double
```

Types are **syntactic only** and have no runtime effect.

### 3.4 Entry Point Enforcement

Every valid Word of God program **must** begin with `TOKEN_AND_GOD_SAID` and end with `TOKEN_AND_IT_CAME_TO_PASS`.

Absence of either is a parse-time error.

### 3.5 Error Recovery

The parser collects errors without stopping, attempting to continue parsing remaining statements. This allows reporting multiple errors in a single pass.

---

## 4. Interpretation (Interpreter)

### 4.1 Execution Model

The interpreter executes statements sequentially via a program counter (pc):

```
pc = 0
while pc < len(statements):
    execute(statements[pc])
    pc++
```

### 4.2 Memory Model

The interpreter maintains:

#### 4.2.1 Variable Memory

```go
memory map[string]int64
```

Variables map to 64-bit signed integers. Uninitialized variables default to 0.

#### 4.2.2 String Storage

```go
strings map[string]string
```

Strings are stored separately from numeric memory to prevent collision with arithmetic.

String references use **negative integers**:

```
Reference = -(index + 1)
```

Example:

```
strings["-1"] = "First string"
strings["-2"] = "Second string"
strings["-3"] = "Third string"
```

#### 4.2.3 Label Resolution

```go
labels map[string]int
```

Labels map to statement indices in the AST. Resolved during first pass before execution.

### 4.3 Statement Execution

#### 4.3.1 NODE_DECLARATION

Variables are registered in the symbol table. No memory allocation occurs; memory is lazy-allocated on first assignment.

#### 4.3.2 NODE_ASSIGNMENT

```
memory[variable] = evaluate(expression)
```

#### 4.3.3 NODE_PRINT

```
value = evaluate(expression)
if strings[value] exists:
    print(strings[value])
else:
    print(value)
```

#### 4.3.4 NODE_JUMP

```
pc = labels[target_label] - 1
```

(Subtract 1 because pc is incremented after statement execution)

If label does not exist, runtime error.

#### 4.3.5 NODE_CONDITIONAL_JUMP

```
value = evaluate(expression)
if value == 0:
    pc = labels[target_label] - 1
```

Jump only if expression evaluates to 0 (false).

If label does not exist, runtime error.

#### 4.3.6 NODE_ERROR

```
throw RuntimeError(error_message)
```

Immediately terminates execution with error message.

#### 4.3.7 NODE_LABEL

No operation; labels are markers only.

#### 4.3.8 NODE_RETURN

```
pc = len(statements)  // End execution
```

### 4.4 Expression Evaluation

#### 4.4.1 NODE_LITERAL

```go
case int64:    return value
case string:   ref = -(len(strings) + 1)
               strings[fmt.Sprint(ref)] = value
               return ref
```

#### 4.4.2 NODE_VARIABLE

```
return memory[variable_name]
```

If variable undefined, runtime error.

#### 4.4.3 NODE_BINARY_OP (ADD)

```
left = evaluate(left_operand)
right = evaluate(right_operand)
return left + right
```

#### 4.4.4 NODE_UNARY_OP (NOT)

```
value = evaluate(operand)
return (value == 0) ? 1 : 0
```

### 4.5 Error Handling

Runtime errors include:

- Undefined variable reference
- Unknown label in jump
- Arithmetic overflow (implementation-defined)
- Type mismatch (not enforced)

All errors are reported with line number and context.

---

## 5. Semantics

### 5.1 Boolean Values

Booleans are represented as integers:

```
0  = false
1  = true
any non-zero = true (for conditionals)
```

### 5.2 Arithmetic

Only addition is supported. All binary operations perform addition:

```
And x And y  =>  x + y
And a And b And c  =>  a + b + c
```

### 5.3 String Representation

Strings are immutable once created. Each string literal creates a new entry in the strings map.

Printing a string reference (negative integer) automatically looks up the string value.

### 5.4 Variable Scope

All variables are global within a program. No block scoping or local scope.

Redeclaration is tracked but not enforced; redeclaring a variable overwrites it.

### 5.5 Execution Termination

A program terminates normally when:

1. Execution reaches the end of the statement list
2. A NODE_RETURN statement is encountered
3. NODE_AND_IT_CAME_TO_PASS is executed (implicit return)

A program terminates with error when:

1. A NODE_ERROR (Woe unto) is encountered
2. An undefined label is jumped to
3. An undefined variable is referenced

---

## 6. Compilation Pipeline

### 6.1 Lexer

- Input: Source code string
- Output: Token slice
- Complexity: O(n) where n = source length
- Side effects: Line number tracking

### 6.2 Parser

- Input: Token slice
- Output: AST with symbol table
- Complexity: O(t) where t = token count
- Side effects: Symbol registration, error collection

### 6.3 Interpreter

- Input: AST + symbol table
- Output: Program output + exit status
- Complexity: O(s) where s = statement count (with jumps, potentially O(s*j) with jump iterations)
- Side effects: Memory modification, stdout writes

---

## 7. Extension Points

### 7.1 New Operators

Addition of new binary operators (multiplication, subtraction, etc.) requires:

1. New TOKEN_* types in lexer
2. Keyword matching in identifier()
3. New NODE_BINARY_OP values
4. Evaluation logic in interpreter.evaluate()

### 7.2 New Control Flow

New control structures require:

1. New token types
2. Parser statement rules
3. New AST node types
4. Interpreter execution logic

### 7.3 Functions

Function support would require:

1. Function definition tokens and parsing
2. Function call expressions
3. Return statements
4. Call stack management
5. Local scope implementation

---

## 8. Implementation Notes

### 8.1 Go Specifics

Gabriel is implemented in Go 1.19+:

- Token stream uses slice of Token structs
- AST uses recursive Node pointers
- Symbol table uses map[string]*Symbol
- Memory uses map[string]int64
- Error handling uses Go error interface

### 8.2 Cross-Platform Support

Gabriel builds successfully for:

- Windows (x86-64)
- Linux (x86-64)
- macOS (x86-64)

Target GOOS/GOARCH can be specified during build.

### 8.3 Performance Characteristics

- Lexing: Linear in source length
- Parsing: Linear in token count
- Interpretation: Linear in statement count plus jump overhead
- Memory: O(v + s) where v = variables, s = strings

---

## 9. Testing

Gabriel should be tested with:

1. **Unit tests**: Individual lexer, parser, interpreter components
2. **Integration tests**: Complete programs end-to-end
3. **Error tests**: Invalid input handling
4. **Edge cases**: Infinite loops, deep nesting, large numbers

---

## 10. Conformance

An implementation conforms to this specification if it:

1. Recognizes all 14 token types correctly
2. Enforces entry point requirement
3. Implements all AST node types
4. Maintains accurate symbol table with lineage
5. Executes statements in correct order
6. Handles jumps and conditionals correctly
7. Reports errors with line numbers
8. Produces correct output for valid programs

---

## Appendix A: Token Reference

| Token | Lexeme | Type | Purpose |
|-------|--------|------|---------|
| AND_GOD_SAID | "And God said" | Keyword | Entry point |
| AND_IT_CAME_TO_PASS | "And it came to pass" | Keyword | Program end |
| LET_THERE_BE | "Let there be" | Keyword | Declaration |
| THOU_SHALT | "Thou shalt" | Keyword | Assignment |
| AND | "And" | Keyword | Addition |
| YE_SHALL_NOT | "Ye shall not" | Keyword | Negation |
| VERILY_VERILY | "Verily, verily" | Keyword | True |
| BEHOLD | "Behold" | Keyword | Output |
| WOE_UNTO | "Woe unto" | Keyword | Error |
| GO_YE_UNTO | "Go ye unto" | Keyword | Jump |
| IDENTIFIER | [a-zA-Z_][a-zA-Z0-9_]* | Identifier | Names |
| NUMBER | [0-9]+(.[0-9]+)? | Literal | Numbers |
| STRING | "[^"]*" | Literal | Strings |
| COLON | : | Delimiter | Type/lineage separator |
| COMMA | , | Delimiter | Statement separator |
| EOF | (implicit) | Marker | End of input |

---

## Appendix B: Error Messages

Gabriel reports errors in format:

```
Gabriel proclaims: Line N: error message
```

Example error messages:

```
Gabriel proclaims: Line 5: Unknown label 'undefined_label'
Gabriel proclaims: Line 8: undefined variable: 'x'
Gabriel proclaims: Line 14: Expected expression after 'And'
Gabriel proclaims: Program must begin with 'And God said'
Gabriel proclaims: Program must end with 'And it came to pass'
```

---

End of Gabriel Specification