# The Go programming language

Solutions to the exercises of [The Go Programming Book](http://www.gopl.io/).

**Table of Contents**

- [The Go programming language](#the-go-programming-language)
- [Getting the book's examples](#getting-the-books-examples)
  - [Windows](#windows)
- [License](#license)

# Getting the book's examples

Using `go get` to get the examples as described in the book no longer 
works, use the following instead:

## Windows

```shell
cd %GOPATH%
go install gopl.io/ch1/helloworld@latest
%GOPATH%/bin/helloworld.exe
```

# License

These solutions are licensed under a 
[Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License](LICENSE.md)
as is the original [example code](https://github.com/adonovan/gopl.io/).
