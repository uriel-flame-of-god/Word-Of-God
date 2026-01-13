# Word of God

A biblical esoteric programming language where code reads like scripture.

## Overview

Word of God is an esoteric programming language designed to make executable code that reads like biblical text. The language treats divine authority as mutable state and judgment as control flow interruption. Programs are structured as sacred proclamations, with keywords drawn from biblical cadence and authority.

This repository contains Gabriel, the official compiler and interpreter for Word of God, along with syntax highlighting support for Visual Studio Code.

## Features

- Imperative, linear execution model
- Label-based control flow with goto semantics
- Biblical lineage-tracked symbol declarations
- Type-annotated variable declarations with no runtime enforcement
- String and numeric literal support
- Conditional and unconditional jumps
- Error handling via divine judgment
- Case-insensitive keyword matching
- Full AST-based interpretation

## Installation

### Prerequisites

- Go 1.19 or later
- Git (optional, for cloning)

### Building from Source

Clone or download this repository, then choose your build method:

#### Using Make (Linux/macOS)

```bash
make build
make install
```

#### Using build.sh (Linux/macOS)

```bash
chmod +x build.sh
./build.sh install
```

#### Using build.bat (Windows)

```batch
build.bat install
```

#### Manual Build

```bash
go build -o gabriel main.go
```

### VS Code Syntax Highlighting

Install the Word of God syntax extension for color-coded highlighting:

```bash
make setup-vscode
```

Or manually copy the extension:

```bash
cp -r .vscode/extensions/word-of-god-syntax ~/.vscode/extensions/
```

## Quick Start

### 1. Create a Program

Create a file named `hello.wog`:

```
And God said

Behold "Hello, World"

And it came to pass
```

### 2. Run the Program

```bash
gabriel hello.wog
```

### 3. Output

```
GABRIEL interpreting hello.wog...

Tokens received:
   Line 1: AND_GOD_SAID -> 'And God said'
   Line 3: BEHOLD -> 'Behold'
   Line 3: STRING -> '"Hello, World"'
   Line 5: AND_IT_CAME_TO_PASS -> 'And it came to pass'

AST generated with 3 nodes

Executing divine commands...

Hello, World

Execution complete.
```

## Language Guide

### Program Structure

All Word of God programs have this structure:

```
And God said
    [declarations and statements]
And it came to pass
```

The program begins with "And God said" and must end with "And it came to pass". Everything outside these markers is ignored.

### Keywords

Word of God has nine keyword phrases:

| Keyword | Purpose | Example |
|---------|---------|---------|
| **And God said** | Program entry point | And God said |
| **Let there be** | Variable declaration | Let there be x 2:1 |
| **Thou shalt** | Assignment | Thou shalt x And 42 |
| **And** | Addition | Thou shalt sum And x And y |
| **Ye shall not** | Negation / conditional jump | Ye shall not valid, go ye unto error |
| **Verily, verily** | Boolean true | Thou shalt flag Verily, verily |
| **Behold** | Output / print | Behold x |
| **Woe unto** | Error / termination | Woe unto "Fatal error" |
| **Go ye unto** | Unconditional jump | Go ye unto loop_start |
| **And it came to pass** | Program end | And it came to pass |

### Variables and Declaration

Variables are declared with type and lineage annotations:

```
Let there be variable_name type:lineage
```

**Type numbers** (0-5):

- 0 = void
- 1 = boolean
- 2 = integer
- 3 = string
- 4 = float
- 5 = double

Types are syntactic only and do not affect runtime behavior.

**Lineage** represents variable scope depth. If a variable is redeclared, its lineage must be higher than the previous declaration.

Example:

```
Let there be count 2:1
Let there be flag 1:1
Let there be message 3:1
```

### Assignment

Assign values to variables using "Thou shalt":

```
Thou shalt variable_name expression
```

Examples:

```
Thou shalt x Verily, verily
Thou shalt count And 42
Thou shalt sum And x And y And 10
```

### Arithmetic

All arithmetic uses addition via the "And" keyword:

```
Thou shalt result And a And b And c
```

This evaluates as: a + b + c

To increment a variable:

```
Thou shalt counter And counter And Verily, verily
```

This evaluates as: counter + 1

### Output

Print values using "Behold":

```
Behold variable_name
Behold 42
Behold "String literal"
Behold And x And y
```

### Conditionals and Jumps

#### Unconditional Jump

```
Go ye unto label_name
```

#### Conditional Jump with Negation

```
Ye shall not expression, go ye unto label_name
```

The jump occurs if the expression evaluates to 0 (false).

Example:

```
Ye shall not valid, go ye unto error_handler
```

### Labels

Labels mark jump targets and use biblical lineage notation:

```
label_name type:lineage
```

Example:

```
loop 0:1
Behold x
Thou shalt x And x And Verily, verily
Ye shall not x, go ye unto loop

error_handler 0:1
Woe unto "Error occurred"
```

Lineage numbers must increment if the same label name is used multiple times.

### Error Handling

Terminate execution with a judgment:

```
Woe unto "error message"
```

This immediately halts the program and reports the error.

### Comments

Comments start with double slashes and extend to the end of the line:

```
// This is a comment
Thou shalt x And 42  // Assign 42 to x
```

### Strings

String literals are enclosed in double quotes:

```
Behold "Hello, World"
Woe unto "Program failed"
```

### Numbers

Integers and floating-point numbers are supported:

```
Thou shalt a And 42
Thou shalt pi And 3.14159
```

## Example Programs

### Example 1: Simple Counter

```
And God said

Let there be i 2:1
Let there be limit 2:1

Thou shalt i And 0
Thou shalt limit And 10

loop 0:1

Behold i
Thou shalt i And i And Verily, verily

Ye shall not i, go ye unto done

Go ye unto loop

done 0:2
Behold "Counting complete"

And it came to pass
```

### Example 2: Boolean Validation

```
And God said

Let there be input 2:1
Let there be valid 1:1

Thou shalt input And 42
Thou shalt valid Verily, verily

Ye shall not valid, go ye unto error_handler

Behold "Validation passed!"
Behold input
Go ye unto success_end

error_handler 0:1
Woe unto "Validation failed"

success_end 0:2
Behold "Program complete"

And it came to pass
```

### Example 3: Fibonacci Sequence

```
And God said

Let there be a 2:1
Let there be b 2:1
Let there be c 2:1
Let there be counter 2:1

Thou shalt a Verily, verily
Thou shalt b Verily, verily
Thou shalt counter And 0

loop 0:1

Verily, verily, go ye unto end

Behold a

Thou shalt c And a And b
Thou shalt a And b And 0
Thou shalt b And c And 0
Thou shalt counter And counter And Verily, verily

Go ye unto loop

end 0:2
Behold "Fibonacci complete"

And it came to pass
```

## Build Instructions

### Using Make (Linux/macOS)

```bash
make help           # Show all targets
make build          # Build Gabriel
make run FILE=prog.wog
make install        # Install to system
make clean          # Remove artifacts
```

### Using build.sh (Linux/macOS)

```bash
./build.sh              # Build for current OS
./build.sh all          # Build for all platforms
./build.sh windows      # Build for Windows
./build.sh linux        # Build for Linux
./build.sh macos        # Build for macOS
./build.sh install      # Install to system
./build.sh clean        # Clean artifacts
```

### Using build.bat (Windows)

```batch
build.bat               # Build for Windows
build.bat all           # Build for all platforms
build.bat linux         # Build for Linux
build.bat macos         # Build for macOS
build.bat install       # Install to system
build.bat clean         # Clean artifacts
```

### Manual Build

```bash
go build -o gabriel main.go
```

## Project Structure

```
Word-of-God/
├── main.go                      # Gabriel compiler source
├── Makefile                     # Build automation (Make)
├── build.sh                     # Build script (Linux/macOS)
├── build.bat                    # Build script (Windows)
├── LICENSE                      # MIT License
├── README.md                    # This file
├── GABRIEL.md                   # Gabriel compiler documentation
├── .vscode/                     # VS Code configuration
│   ├── settings.json            # Workspace settings
│   ├── extensions.json          # Extension recommendations
│   └── extensions/
│       └── word-of-god-syntax/  # Syntax highlighting extension
│           ├── package.json
│           ├── language-configuration.json
│           ├── syntaxes/
│           │   └── wog.json
│           └── README.md
├── build/                       # Build output directory (generated)
│   └── gabriel                  # Compiled binary
└── examples/                    # Example programs
    ├── hello.wog
    ├── counter.wog
    ├── fibonacci.wog
    └── validation.wog
```

## Interpreter Architecture

Gabriel operates in three stages:

### 1. Lexical Analysis (Lexer)

The source code is tokenized into meaningful units. The lexer recognizes:

- Multi-word keywords (case-insensitive)
- Identifiers
- Numbers (integers and floats)
- Strings (with escape sequences)
- Delimiters (colons, commas)
- Comments

### 2. Parsing (Parser)

Tokens are assembled into an Abstract Syntax Tree (AST). The parser:

- Validates program structure
- Resolves symbols and labels
- Enforces lineage constraints
- Tracks type annotations
- Generates semantic nodes

### 3. Interpretation (Interpreter)

The AST is executed sequentially. The interpreter:

- Manages variable memory
- Handles string references
- Resolves label jumps
- Evaluates expressions
- Reports runtime errors

## Error Handling

Gabriel provides detailed error messages with line numbers:

```
Line 14: Unknown label 'nonexistent'
Line 8: undefined variable: 'x'
Line 5: Expected expression after 'And'
```

All runtime errors are prefixed with context information to aid debugging.

## Limitations

- Word of God is a goto-based language with no function declarations or function calls
- Only addition is supported for arithmetic
- Type annotations are for documentation only and do not enforce runtime type checking
- No recursive data structures or arrays
- All variables are 64-bit signed integers (except strings)
- Strings are stored separately with negative integer references

## Contributing

Contributions are welcome. Please submit issues and pull requests to improve Gabriel or the language specification.

## License

Word of God and Gabriel are licensed under the MIT License. See LICENSE file for details.

## Semantic Philosophy

In Word of God:

- **Authority mutates state**: The "Thou shalt" keyword establishes declarations that persist until program termination.
- **Judgment interrupts control**: The "Woe unto" keyword represents divine judgment, immediately halting execution.
- **Lineage tracks descent**: Variables and labels track their declaration order through lineage numbers, maintaining a genealogy of symbols.
- **Veracity is true**: "Verily, verily" represents absolute truth, evaluating to 1.
- **Negation inverts**: "Ye shall not" inverts boolean values and enables conditional branching.

## Changelog

### Version 1.0.0

- Initial release of Gabriel compiler
- Full language specification implementation
- VS Code syntax highlighting extension
- Cross-platform build support (Windows, Linux, macOS)
- Comprehensive documentation

## Future Enhancements

Potential improvements for future versions:

- Multiplication and subtraction operators
- Array and structure support
- Function declarations and calls
- Stack-based memory management
- Improved error messages with suggestions
- Debugger support
- Performance optimizations

## Acknowledgments

Word of God is an esoteric language created as a artistic exploration of biblical cadence in code. It demonstrates how programming language design can be used to create unique and expressive syntactic experiences.

---

For more information about Gabriel, see SPEC/GABRIEL.md.

For VS Code extension details, see .vscode/extensions/word-of-god-syntax/README.md.