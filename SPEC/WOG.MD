# Word of God Language Specification

## Overview

Word of God is an esoteric programming language designed to read like biblical scripture while maintaining computational completeness within its feature set. This document formally specifies the language syntax, semantics, and execution model.

## Scope

This specification defines:

- Lexical structure of Word of God
- Formal syntax and grammar
- Type system and annotations
- Runtime semantics
- Program structure requirements
- Execution guarantees

## Version

Word of God 1.0 (Standard)

---

## 1. Language Philosophy

Word of God operates on three theological principles:

### 1.1 Authority Mutates State

The keyword "Thou shalt" establishes declarations that persist throughout program execution. Once proclaimed, a variable's state persists until explicitly reassigned or program termination.

### 1.2 Judgment Halts Execution

The keyword "Woe unto" represents divine judgment. Upon judgment, execution halts immediately, regardless of remaining statements. No recovery mechanism exists.

### 1.3 Lineage Tracks Descent

Variables and labels maintain genealogy through lineage numbers. If a symbol is declared multiple times, each declaration must have strictly greater lineage than previous declarations, preventing ambiguous shadowing.

---

## 2. Lexical Structure

### 2.1 Source Encoding

Word of God source files use UTF-8 encoding with `.wog` file extension.

### 2.2 Character Set

Valid characters include:

- ASCII letters (a-z, A-Z)
- Digits (0-9)
- Punctuation (: , " / spaces)
- Whitespace (space, tab, newline, carriage return)

### 2.3 Tokens

A program is composed of tokens separated by whitespace or delimiters.

#### 2.3.1 Keywords (Case-Insensitive)

```
And God said         // Entry point
Let there be         // Declaration
Thou shalt           // Assignment
And                  // Addition
Ye shall not         // Negation / conditional
Verily, verily       // Boolean true
Behold               // Output
Woe unto             // Error
Go ye unto           // Jump
And it came to pass  // Program termination
```

#### 2.3.2 Identifiers (Case-Sensitive)

Identifiers name variables and labels:

```
pattern: [a-zA-Z_][a-zA-Z0-9_]*
```

Examples: `x`, `counter`, `_private`, `my_var_123`

#### 2.3.3 Literals

**Numbers** (integers and floats):

```
42
3.14159
0
-5 (via And operator)
```

**Strings** (double-quoted):

```
"Hello, World"
"multi line string"
"with \"escaped\" quotes"
```

#### 2.3.4 Delimiters

```
:   (colon, separates type and lineage)
,   (comma, separates conditional condition and jump)
```

#### 2.3.5 Comments

Single-line comments beginning with `//`:

```
// This entire line is a comment
Thou shalt x And 42  // Inline comment
```

---

## 3. Syntax and Grammar

### 3.1 Program Structure

```
Program = EntryPoint Statements Terminator

EntryPoint = "And God said"
Statements = Statement*
Terminator = "And it came to pass"

Statement = Declaration
          | Assignment
          | Label
          | PrintStatement
          | JumpStatement
          | ConditionalJump
          | ErrorStatement
```

### 3.2 Declarations

```
Declaration = "Let there be" Identifier Type ":" Lineage

Type      = Integer (0-5, modulo 6)
Lineage   = Integer (positive)

Type Map:
  0 -> void
  1 -> boolean
  2 -> integer
  3 -> string
  4 -> float
  5 -> double
```

**Semantics**: Declares a variable with optional type annotation. Type annotations are **syntactic only** and do not enforce runtime type checking.

**Lineage Rule**: If the identifier was previously declared, its new lineage must be strictly greater:

```
Lineage[new] > Lineage[previous]  must be true
```

**Example**:

```
Let there be x 2:1        // x: int, lineage 1
Let there be flag 1:1     // flag: bool, lineage 1
Let there be x 2:2        // x: int, lineage 2 (valid redeclaration)
Let there be x 2:1        // ERROR: lineage not greater
```

### 3.3 Assignments

```
Assignment = "Thou shalt" Identifier [Expression]

Expression is optional; if omitted, evaluates to 0.
```

**Semantics**: Assigns the value of the expression to the variable.

**Examples**:

```
Thou shalt x And 42
Thou shalt y Verily, verily
Thou shalt z                      // z = 0
Thou shalt result And a And b
```

### 3.4 Labels

```
Label = Identifier Type ":" Lineage

Type      = Integer (0-5)
Lineage   = Integer (positive)
```

**Semantics**: Defines a jump target. Labels are statements that mark positions for branching.

**Lineage Rule**: Same as declarations. Repeated label names require strictly increasing lineage.

**Example**:

```
loop 0:1
    Behold x
    Thou shalt x And x And Verily, verily
    Ye shall not x, go ye unto loop

error_handler 0:1
Woe unto "Error"

success 0:2
Behold "Done"
```

### 3.5 Print Statements

```
PrintStatement = "Behold" Expression
```

**Semantics**: Evaluates the expression and outputs its value.

- If the value is a string reference (negative integer), the referenced string is printed.
- Otherwise, the numeric value is printed.

**Examples**:

```
Behold x
Behold 42
Behold "Hello"
Behold And x And y
```

### 3.6 Jump Statements

```
JumpStatement = "Go ye unto" Identifier
```

**Semantics**: Unconditional jump to the label named by the identifier.

**Error**: If label does not exist, runtime error.

**Example**:

```
Go ye unto loop_start
```

### 3.7 Conditional Jumps

```
ConditionalJump = "Ye shall not" Expression "," "Go ye unto" Identifier
```

**Semantics**: If the expression evaluates to 0 (false), jump to the label. Otherwise, continue to next statement.

**Error**: If label does not exist, runtime error.

**Example**:

```
Ye shall not valid, go ye unto error_handler
Ye shall not i, go ye unto done
```

### 3.8 Error Statements

```
ErrorStatement = "Woe unto" String
```

**Semantics**: Immediately terminates program execution and reports the error message.

**Example**:

```
Woe unto "Fatal error: invalid input"
```

### 3.9 Expressions

```
Expression = ["And"] Primary ("And" Primary)*

Primary = Number
        | String
        | "Verily, verily"
        | Identifier
```

**Semantics**:

- Numbers evaluate to their numeric value
- Strings create string references (negative integers)
- "Verily, verily" evaluates to 1 (boolean true)
- Identifiers evaluate to their current variable value
- "And" performs addition between operands

**Examples**:

```
42                           // Evaluates to 42
Verily, verily               // Evaluates to 1
And 5 And 3                  // Evaluates to 0 + 5 + 3 = 8
x                            // Evaluates to value of x
And x And 10                 // Evaluates to 0 + x + 10
And a And b And c And d      // Evaluates to a + b + c + d
```

### 3.10 Negation in Conditionals

```
"Ye shall not" Expression, "Go ye unto" Label
```

The expression is implicitly negated within a conditional jump context. If expression is 0, the condition is true.

**Semantic equivalence**:

```
Ye shall not expr, go ye unto label
// Equivalent to:
if NOT(expr):
    goto label
```

---

## 4. Type System

### 4.1 Types

Word of God has a syntactic type system with six types:

| Type ID | Name | Purpose |
|---------|------|---------|
| 0 | void | Untyped / no value |
| 1 | boolean | Truth values (0 or 1) |
| 2 | integer | Whole numbers |
| 3 | string | Text |
| 4 | float | Decimal numbers |
| 5 | double | High-precision decimal |

### 4.2 Type Annotations

Types are declared in variable declarations and labels:

```
Let there be x 2:1        // x is declared as integer (type 2)
Thou shalt flag 1:1       // flag is declared as boolean (type 1)
message 3:1               // label message has type string (3)
```

### 4.3 Runtime Type Semantics

**Type annotations are syntactic only.** At runtime, no type checking is performed.

All values are stored as 64-bit signed integers in memory, except strings which are stored separately with negative reference integers.

### 4.4 Implicit Type Coercion

No implicit type coercion occurs. All operations treat operands as numeric.

---

## 5. Memory Model

### 5.1 Variable Storage

Variables store 64-bit signed integers. Uninitialized variables have default value 0.

### 5.2 String Storage

Strings are stored separately in a string pool with negative integer references.

**String reference assignment**:

```
ref = -(pool_index + 1)
```

Example:

```
strings[-1] = "First"
strings[-2] = "Second"
strings[-3] = "Third"
```

### 5.3 Variable Scope

All variables are global within a program. No block scoping or lexical scope.

Redeclaring a variable with higher lineage does not create a new scope; it refers to the same variable.

### 5.4 Memory Lifetime

Variables persist for the entire program duration unless explicitly reassigned.

---

## 6. Execution Model

### 6.1 Execution Order

Statements execute sequentially from top to bottom unless interrupted by:

- Unconditional jump ("Go ye unto")
- Conditional jump ("Ye shall not ... go ye unto")
- Error termination ("Woe unto")
- Program termination ("And it came to pass")

### 6.2 Program Termination

A program terminates in one of these ways:

1. **Normal termination**: Execution reaches "And it came to pass"
2. **Jump termination**: Execution reaches end of statement list via jumps
3. **Error termination**: "Woe unto" error statement encountered

### 6.3 Control Flow Semantics

#### 6.3.1 Unconditional Jump

```
Go ye unto label
```

Sets the instruction pointer to the statement at `label` without condition.

#### 6.3.2 Conditional Jump

```
Ye shall not expression, go ye unto label
```

Evaluates the expression. If result is 0, jump to label. Otherwise, continue.

#### 6.3.3 Error Termination

```
Woe unto "message"
```

Immediately halts execution. No recovery possible. Message is reported to stderr or output stream.

---

## 7. Semantics

### 7.1 Arithmetic

All arithmetic operations perform **addition only**.

```
a And b        // a + b
a And b And c  // a + b + c
```

No subtraction, multiplication, division, or modulo operations exist.

Negative numbers are created through assignment:

```
Thou shalt neg And 0 And -5  // neg = 0 + (-5) = -5
```

### 7.2 Boolean Logic

Booleans are represented as integers:

```
0              // False
1 (Verily)     // True
non-zero       // True (in conditional context)
```

Negation via "Ye shall not" inverts the truth value:

```
Ye shall not 0         // True (negation of false)
Ye shall not 1         // False (negation of true)
Ye shall not x         // True if x is 0
```

### 7.3 String Semantics

Strings are immutable. Each string literal creates a reference.

Multiple identical string literals create separate references:

```
Behold "Hello"
Behold "Hello"  // Creates two separate string entries
```

String references (negative integers) are never directly comparable to positive integers in semantic terms, but comparisons work numerically.

### 7.4 Operator Precedence

Word of God has no operator precedence beyond left-to-right evaluation within expressions.

```
And a And b And c  // Evaluated left-to-right: ((0 + a) + b) + c
```

### 7.5 Undefined Behavior

The following conditions are undefined:

- Arithmetic overflow (system-dependent)
- Reading uninitialized variables (default to 0, but not guaranteed)
- Circular jump loops without termination condition (infinite loop)

---

## 8. Well-Formedness

A valid Word of God program must satisfy:

1. **Entry requirement**: Must begin with "And God said"
2. **Termination requirement**: Must end with "And it came to pass"
3. **Label validity**: All jump targets must be declared labels
4. **Unique labels**: Label names must be unique per lineage level
5. **Lineage monotonicity**: Redeclared symbols must have strictly increasing lineage

A program violating these requirements is rejected at parse time.

---

## 9. Examples

### 9.1 Hello World

```
And God said

Behold "Hello, World"

And it came to pass
```

### 9.2 Counter Loop

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

Behold "Complete"

And it came to pass
```

### 9.3 Conditional Branching

```
And God said

Let there be value 2:1

Thou shalt value And 42

Ye shall not value, go ye unto zero_value

Behold "Value is non-zero"
Go ye unto end

zero_value 0:1
Behold "Value is zero"

end 0:2
Behold "Program done"

And it came to pass
```

### 9.4 Error Handling

```
And God said

Let there be input 2:1

Thou shalt input And 100

Ye shall not input, go ye unto error

Behold "Input valid"
Go ye unto success

error 0:1
Woe unto "Input is zero"

success 0:2
Behold "Exiting"

And it came to pass
```

---

## 10. Grammar (Formal EBNF)

```ebnf
Program         ::= "And God said" Statement* "And it came to pass"

Statement       ::= Declaration
                 |  Assignment
                 |  Label
                 |  PrintStatement
                 |  JumpStatement
                 |  ConditionalJump
                 |  ErrorStatement

Declaration     ::= "Let there be" Identifier Number ":" Number

Assignment      ::= "Thou shalt" Identifier (Expression)?

Label           ::= Identifier Number ":" Number

PrintStatement  ::= "Behold" Expression

JumpStatement   ::= "Go ye unto" Identifier

ConditionalJump ::= "Ye shall not" Expression "," "Go ye unto" Identifier

ErrorStatement  ::= "Woe unto" String

Expression      ::= ("And")? Primary ("And" Primary)*

Primary         ::= Number
                 |  String
                 |  "Verily, verily"
                 |  Identifier

Identifier      ::= [a-zA-Z_][a-zA-Z0-9_]*

Number          ::= [0-9]+ ("." [0-9]+)?

String          ::= '"' [^"]* '"'
```

---

## 11. Computational Properties

### 11.1 Turing Completeness

Word of God is **Turing-complete within its feature set**:

- Arbitrary variables store state
- Conditional branching (via "Ye shall not")
- Unconditional jumps ("Go ye unto")
- Sequential composition

These features enable:

- Arbitrary computation
- Loop constructs (via labels and conditionals)
- Decision trees (via branching)

### 11.2 Limitations

Word of God **lacks**:

- Function definitions and calls
- Arrays or data structures
- Dynamic memory allocation
- Recursive calls
- Higher-order functions

These limitations prevent some algorithms from being expressed elegantly but do not affect Turing completeness.

---

## 12. Conformance

A language implementation conforms to this specification if:

1. It accepts all syntactically valid Word of God programs as defined
2. It rejects all syntactically invalid programs
3. It executes valid programs according to semantics defined in sections 5-7
4. It correctly reports errors for non-conformant programs

---

## 13. Extension Points

### 13.1 Additional Operators

Future versions could add:

- Multiplication ("Multiplied by")
- Subtraction ("Minus")
- Modulo ("Remainder")
- Bitwise operations

### 13.2 Advanced Control Flow

- Function definitions
- Local scoping
- Recursion
- Exception handling beyond "Woe unto"

### 13.3 Data Structures

- Arrays
- Structures / records
- Hash maps

---

## Appendix A: Keyword Reference

| Keyword | Arity | Purpose |
|---------|-------|---------|
| And God said | 0 | Entry point |
| Let there be | 2 (type:lineage) | Declaration |
| Thou shalt | 1+ (var + expr) | Assignment |
| And | Binary | Addition |
| Ye shall not | 1 (expr) | Negation/conditional |
| Verily, verily | 0 | Boolean true |
| Behold | 1 | Output |
| Woe unto | 1 (string) | Error |
| Go ye unto | 1 (label) | Jump |
| And it came to pass | 0 | Termination |

---

## Appendix B: Lineage Examples

```
Let there be x 2:1          // x declared with lineage 1
Thou shalt x And 10         // x assigned

Let there be x 2:2          // x redeclared with lineage 2 (VALID)
Thou shalt x And 20         // x now holds 20

Let there be x 2:2          // ERROR: lineage 2 not > 2
```

---

## Appendix C: Symbol Table Structure

```
Symbol = {
    Name: string        // Variable or label name
    Type: int (0-5)     // Type annotation
    Lineage: int        // Declaration order
    IsLabel: bool       // True if label, false if variable
}
```

---
