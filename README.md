# Go CLI Calculator

A simple command-line calculator written in Go (Golang). It supports basic arithmetic operations, exponentiation, and resetting/re-entering values during execution.

## Project Structure

* `main.go` — Entry point of the application.
* `pkg/calc` — Package containing core mathematical operations (`Add`, `Sub`, `Mult`, `Dev`, `Pow`).
* `pkg/choose` — Package for handling user input and the interactive menu.

## Features

* `1` — Addition (`+`)
* `2` — Subtraction (`-`)
* `3` — Multiplication (`*`)
* `4` — Division (`/`)
* `5` — Exponentiation (`^`)
* `6` — Reset and enter new numbers
* `0` — Exit application

## Installation and Setup

Make sure you have [Go](https://go.dev/) installed on your system.

1. Clone the repository:
   ```bash
   git clone [https://github.com/bubkub10-lang/calculator.git](https://github.com/bubkub10-lang/calculator.git)
   cd calculator
   ```

## Run The Program
 ```bash
go run main.go
# or
go build main.go
./main
```

## Usage Example
```text
> Choose Action 
  > [0] Exit 
  > [1] Addition 
  > [2] Substraction 
  > [3] Multiplication 
  > [4] Division 
  > [5] Pow 
  > [6] New Numbers> 6
> Calculator 
> (need 2 numbers) 
> 1: 10
> 2: 5
> Choose Action 
...
> 1
15
```
