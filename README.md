<p align="center">
  <img width="871" alt="iPhone 14 - 1 (1)" src="https://user-images.githubusercontent.com/79533745/214399490-7347a24a-9c4e-4eb8-a2a2-5322c8845a67.png">
</p>

<h1 align="center">Snow Language</h1>
<p align="center">A programming language made in Go for fun</p>

<div align="center">
  
  <a href="https://github.com/snowlanguage/go-snow/releases">![GitHub release (latest by date)](https://img.shields.io/github/v/release/snowlanguage/go-snow?label=Stable%20release)</a>
  <a href="">![GitHub commit activity](https://img.shields.io/github/commit-activity/m/snowlanguage/go-snow?label=Commit%20activity)</a>
  
</div>

- [Command line tool](#command-line-tool)
- [How to use](#how-to-use)
  - [Expressions](#expressions)
  - [Comments](#comments)
    - [Line comments](#line-comments)
    - [Inline comments](#inline-comments)
  - [Variables](#variables)
    - [Declaration](#declaration)
    - [Setting](#setting)
    - [Getting](#getting)
  - [If statements](#if-statements)
  - [While statement](#while-statement)
    - [Break statement](#break-statement)
    - [Continue statement](#continue-statement)
  - [Functions](#functions)
    - [Declaration](#declaration-1)
    - [Calling](#calling)
    - [Returning a value](#returning-a-value)
    - [Arguments](#arguments)


## Command line tool

How to use the command line tool

```bash
./snow <path?: string>
```

* `path?: string`: a path to the file of code you want to run. If not specified you'll run the repl instead.

## How to use

How to use Snow Language

A list of all of Snow Language grammar rules can be found **[here](https://github.com/snowlanguage/snowlang-grammar)**

### Expressions

Order of expressions, with highest priority last

| Operation  | Description                                                                             |
|------------|-----------------------------------------------------------------------------------------|
| Assignment | Variable assignments                                                                    |
| Or         | The or operator                                                                         |
| And        | The and operator                                                                        |
| Comparison | The equality and nonequality operators                                                  |
| Comparison | The greater and less than operators                                                     |
| Term       | The addition and subtraction operators                                                  |
| Factor     | The multiplication and division operators                                               |
| Unary      | The invert and negative operators                                                       |
| Call       | A function call or attribute get                                                        |
| Primary    | Numbers, booleans, strings, null, identifiers, grouped expressions and super expression |

### Comments

Comments are used for parts of your code you don't want to be run.

#### Line comments

Makes the rest of the line a comment

```snow
1 + 1 # Comments are declared with the # symbol
```

#### Inline comments

Starts with `#/` and ends with `/#` and can span multiple lines

```snow
42 #/ The meaning of life /# + 1 # Results in 43
```

### Variables

Used to store data
#### Declaration

Muteable

```snow
var varName = "expression" # Can be changed
```

Constant

```snow
const varName = "expression" # Cannot be changed
```

#### Setting

```snow
varName = "expression" # Only works for muteable variables
```

#### Getting

```snow
varName # Just the variable name
```

### If statements

If statements can have one `if`, infinite `elif` and one `else` block.

```snow
if expression == 2 {
  # Do something
} elif expression > 100 {
  # Only runs if all above expressions where false
} else {
  # Only runs if all above expressions where false
}
```

### While statement

Currently the only loop in Snow Language

```snow
while expression {
  # Do something here
}
```

#### Break statement

Breaks out of the loop if used

```snow
while expression { # This loop will only loop once
  if true {
    break
  }
}
```

#### Continue statement

Directly jumps to the beginning of the loop, if the expression of the while loop is `true`

```snow
while expression { # This loop will loop forever
  if true {        # while the expression is true
    continue
  }
}
```

### Functions

You know what a function is

#### Declaration

Functions are always `constant`/`imuteable`

```snow
function funcName() {
  # Do something
}
```

#### Calling

You also need to be able to call functions, don't you?

```snow
function funcName() {
  # Do something
}

funcName() # Don't forget the parenthesis to call the function
```

#### Returning a value

Functions can also return values with the `return` keyword.

```snow
function funcName() {
  return 42
}

funcName() # Results in a value of 42
```

#### Arguments

Function can have arguments which will be passed to the function when its called. Arguments are separated by a comma. 

```snow
function add(x, y) {
  return x + y
}

add(1, add(2, 4 * 3)) # Results in a value of 15
```