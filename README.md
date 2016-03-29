# qu

A command `qufmt` and associated packages to format Qu source code into Go code. The syntax of Qu code is a derivative of Go's syntax, using Kanji for keywords, special identifiers, and common package names, with the aim of eliminating the need for all whitespace in source code.

### License

Copyright © 2016 Gavin "Groovy" Grover

Distributed under the same BSD-style license as Go that can be found in the LICENSE file.

### Status

Version 0.2.

All the functionality described below is implemented, though of course it's raw and there's bugs.

### Installation

Run `go get github.com/gavingroovygrover/qu` to get the command and packages.

Run `go install github.com/gavingroovygrover/qu/cmd/qufmt` to compile and install the `qufmt` command from the downloaded source.

Run `go get github.com/gavingroovygrover/qutests` to get some sample qu code.

Run `qufmt -o src/github.com/gavingroovygrover/qutests/sample.go src/github.com/gavingroovygrover/qutests/sample.qu` to format one of the supplied qu code samples, which can then be run using the standard `go run src/github.com/gavingroovygrover/qutests/sample.go`.


## Intro: Spaceless Programming

By making semicolons optional at line ends, Go departs from other C-derivative languages by allowing newlines to determine semantics in some places in the code. It doesn't require the code to have any newlines, though, even though its style guide and `gofmt` utility recommend their use. For the other whitespace, however, Go does require spaces (or tabs) in the syntax to determine semantics.

Qu extends the optionality of newlines to the other whitespace, enabling a Go program to be written without any whitespace. It does so by introducing one small prohibition to Go's syntax: prohibiting the use of Kanji in identifier names. We then use the Kanji as aliases for Go's keywords, special identifiers, and package names in a modified syntax. Because all of the approx 80,000 Kanji available in Unicode have an implied space both before and after it in the Qu grammar, this allows Qu code to be written without any spaces or tabs, as well as no newlines.

As an added bonus, when a program is written using Kanji in all the places it's permitted in the code, any of the 25 Go keywords can be used as identifier or label names, so a dedicated Qu programmer doesn't need to know any Go-specific exceptions to code. And the Kanji, unlike other non-Ascii characters like `÷`, `≥`, or `←`, are easily enterable via the many IME's (input method editors) available for Chinese and Japanese that ship for free on OS's such as Linux and Windows.

## Usage

### Kanji-based syntax

The 25 keywords of Go can be substituted by any of their respective Kanji below:

* 包 package
* 入 import
* 变 var
* 久 const
* 种 type
* 功 func
* 构 struct
* 图 map
* 面 interface
* 通 chan
* 如 if
* 否 else
* 择 switch
* 事 case
* 别 default
* 掉 fallthrough
* 选 select
* 为 for
* 围 range
* 去 go
* 终 defer
* 回 return
* 破 break
* 继 continue
* 跳 goto

Any of the 25 Go keywords will be treated as an identifier when used within the lexical scope of any Kanji keyword. (This is implemented in Qu 0.1 for all 25 keywords.) To enable this, assignments can be begun with 让, a Kanji best verbalized as "let", and this style should be the prefered style for Qu programmers:

```go
	//added to demo 让:
    让range:="abc" //when used with 让, Go keywords like "range" can be used as identifiers
    让range="abcdefg"
	形Printf("range: %v\n",range)
```

Qi provides `做`, best verbalized as "do", which can be put at the beginning of any stand-alone block. The only effect is to treat the lexical scope of the block as referencing the protected identifier names during parsing.

The 39 special identifiers in Go can also be substituted by their associated Kanji:

* 真 true
* 假 false
* 空 nil
* 毫 iota
* 双 bool
* 节 byte
* 整 int
* 整8 int8
* 整16 int16
* 整32 int32
* 整64 int64
* 绝 uint
* 绝8 uint8
* 绝16 uint16
* 绝32 uint32
* 绝64 uint64
* 漂32 float32
* 漂64 float64
* 复 complex
* 复64 complex64
* 复128 complex128
* 字 rune
* 串 string
* 错 error
* 镇 uintptr
* 造 make
* 新 new
* 关 close
* 删 delete
* 能 cap
* 度 len
* 实 real
* 虚 imag
* 加 append
* 副 copy
* 丢 panic
* 抓 recover
* 写 print
* 线 println

As well as Go aliases `byte` (`节`) and `rune` (`字`), Qu adds alias `任` for `interface{}`, best verbalized as "any".

We also enable Kanji aliases for package names, and they aren't followed by a dot when used. Only 6 are implemented for now:

* 形 fmt
* 网 net
* 序 sort
* 数 math
* 大 math/big
* 时 time


### The seamless mixing of both syntaxes

By remembering a few simple rules, we can easily mix and mingle Go and Qu source code in the same file.

In Go, the 3 categories of identifier are global (the 25 keywords and 39 special identifiers), public (uppercase-initial identifiers), and private (all identifiers beginning with an underscore or lowercase letter, except the 25 keywords).

In Qu, our 3 categories of identifier match more intuitively to their lexical class. They are:

* Kanji. All single-token Kanji which are aliases for keywords, special identifiers, and package names. Every one of the approx 80,000 Kanji available in Unicode has an implied space both before and after it in the Qu grammar.

* public. Uppercase-initial identifiers are visible outside a package in Qu, just like in Go. They have the same format as in Go.

* protected. Identifiers that begin with an underscore followed by lowercase. They are accessible by both Qu and Go code within a single file. When defined or used within Qu code (i.e. within the static scope of a Kanji), the initial underscore is omitted, and so all possible lowercase-initial identifiers, including Go keywords, are permitted there.

Go's private identifiers are inaccessible within Qu code, which generally isn't a problem because, well, they're usually used as parameters and local variables. If a private variable needs to be accessed by both Go and Qu code, put an underscore in front of it when declaring it in Go context.

One example of an identifier that can't have an underscore in front of it is `main`, so we provide the Kanji `正`.


## Examples

This code from "Go by Example":

```go
package main

import (
    "fmt"
    "math/rand"
    "sync/atomic"
    "time"
)

type readOp struct {
    key  int
    resp chan int
}
type writeOp struct {
    key  int
    val  int
    resp chan bool
}

func main() {
    var ops int64 = 0

    reads := make(chan *readOp)
    writes := make(chan *writeOp)

    go func() {
        var state = make(map[int]int)
        for {
            select {
            case read := <-reads:
                read.resp <- state[read.key]
            case write := <-writes:
                state[write.key] = write.val
                write.resp <- true
            }
        }
    }()

    for r := 0; r < 100; r++ {
        go func() {
            for {
                read := &readOp{
                    key:  rand.Intn(5),
                    resp: make(chan int)}
                reads <- read
                <-read.resp
                atomic.AddInt64(&ops, 1)
            }
        }()
    }

    for w := 0; w < 10; w++ {
        go func() {
            for {
                write := &writeOp{
                    key:  rand.Intn(5),
                    val:  rand.Intn(100),
                    resp: make(chan bool)}
                writes <- write
                <-write.resp
                atomic.AddInt64(&ops, 1)
            }
        }()
    }

    time.Sleep(time.Second)

    opsFinal := atomic.LoadInt64(&ops)
    fmt.Println("ops:", opsFinal)
}
```

can be replaced by the terser:

```go
包正;入("math/rand";"sync/atomic")
种readOp构{key整;resp通整} //each line is an example of spaceless programming
种writeOp构{key整;val整;resp通双}  
功正(){
    变ops整64=0
    reads:=造(通*readOp)
    writes:=造(通*writeOp)
    去功(){
        变state=造(图[整]整)
        为{选{事read:=<-reads:read.resp<-state[read.key]
              事write:=<-writes:state[write.key]=write.val;write.resp<-真
             }}}()
    为r:=0;r<100;r++{
        去功(){
            为{read:=&readOp{key:rand.Intn(5),resp:造(通整)}
               reads<-read
               <-read.resp
               atomic.AddInt64(&ops,1)
              }}()}
    为w:=0;w<10;w++{
        去功(){
            为{write:=&writeOp{key:rand.Intn(5),val:rand.Intn(100),resp:造(通双)}
               writes<-write
               <-write.resp
               atomic.AddInt64(&ops,1)
              }}()}
    时Sleep(时Second) //no dot when Kanji used as package name
    opsFinal:=atomic.LoadInt64(&ops)
    形Println("ops:",opsFinal)}
```

If we shorten the local identifier names, and use semicolon (`;`) to join lines together, we can achieve much greater tersity.

### Protecting special identifiers

Go only has 25 reserved words which can't be used as variables, but all the special identifiers such as `false` can be. The Kanji versions of them can't be redefined:

```go
package main
func main() {
	a:= true
	b:= 真
	nil:= true // Qu still allows special identifiers to be used on the left-hand side, ...
	iota:= 真
	//假:= true // ... but when the Kanji version is used, e.g. 假 for false,
	           // generates a parse error "expected non-kanji special identifier on left hand side"
	形Printf("a: %v, b: %v, nil: %v, iota: %v\n", a, b, nil, iota)
}
```

## Rationale

Qu has the purpose of generating discussion on which Kanji should map to which keyword, special identifier, and package in Go, in both China and Japan. If one standard repertoire of Kanji is eventually adopted under the control of some other party/s, then the author encourages others to clone, modify, and promote their own edition of Qu.

The name "qu" is the pinyin of the Mandarin translation of "to go". The Qu syntax is a modification of Go's, parsed by a modified edition of the new recursive descent parser shipped in Go 1.6. The `parser`, `scanner`, and `cmd/gofmt` packages were copied and modified, and the `golang.org/x/tools/go/ast/astutil` package copied.

Qu is a side-path on my goal to create a Unicode-based scripting language for the Go platform with the same inspiration that Groovy's creator had for the JVM platform, with the hope it inspires someone else more appropriate to standardize the meanings of Kanji characters in spaceless programming in the Golang ecosystem. My choice of Kanji are only interim because native Chinese and Japanese speakers will make the final choices. I'm now resuming work on the `ritch` dynamically-typed grammar, to join the `utf88` unicode encoding, `kern` parser combinators, and `thomp` dynamic operators as I rebuild `gro`, the Grover edition of the Groovy language. Hopefully those choices of Kanji for Go will be standardized by someone by the time I need them finalized for `gro`.

