// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parser_test

import (
	"github.com/gavingroovygrover/qu/parser"
	"go/token"
	"go/format"
	"go/types"
	"go/ast"
	"go/importer"
	"testing"
)

type StringWriter struct{ data string }
func (sw *StringWriter) Write(b []byte) (int, error) {
	sw.data += string(b)
	return len(b), nil
}

func TestKanji(t *testing.T) {
	for src, dst:= range kanjiTests {
		fset := token.NewFileSet() // positions are relative to fset
		f, err := parser.ParseFile(fset, "", src, 0)
		if err != nil {
			t.Errorf("parse error: %q", err)
		} else {
			var conf = types.Config{
				Importer: importer.Default(),
			}
			info := types.Info{
				Types: make(map[ast.Expr]types.TypeAndValue),
				Defs:  make(map[*ast.Ident]types.Object),
				Uses:  make(map[*ast.Ident]types.Object),
			}
			_, err = conf.Check("testing", fset, []*ast.File{f}, &info)
			if err != nil {
				t.Errorf("type check error: %q", err)
			}
			sw:= StringWriter{""}
			_= format.Node(&sw, fset, f)
			if sw.data != dst {
				t.Errorf("unexpected Go source: received source: %q; expected source: %q", sw.data, dst)
			}
		}
	}
}

var kanjiTests = map[string]string {

// ========== ========== ========== ==========
//test keywords: 功
//test keyword scoping: 入
`
package main;入"fmt"
功main(){
  fmt.Printf("Hi!\n") // comment here
}`:

// ---------- ---------- ---------- ----------
`package main

import _fmt "fmt"

func _main() {
	_fmt.Printf("Hi!\n")
}
`,

// ========== ========== ========== ==========
//test keyword: 回
//test keyword scoping: 功
//test specids: 度,串,整,整64,漂32,漂64,复,复64,复128
`
package main;入"fmt";
功main(){
  fmt.Printf("Len: %d\n", 度(fs("abcdefg")))
}
功fs(a串)串{回a+"xyz"}
功ff(a漂32)漂64{回漂64(a)}
功fc(a复64)复128{回复128(a)+复(1,1)}
功fi(a整)整64{回整64(a)}
`:

// ---------- ---------- ---------- ----------
`package main

import _fmt "fmt"

func _main() {
	_fmt.Printf("Len: %d\n", len(_fs("abcdefg")))
}
func _fs(_a string) string        { return _a + "xyz" }
func _ff(_a float32) float64      { return float64(_a) }
func _fc(_a complex64) complex128 { return complex128(_a) + complex(1, 1) }
func _fi(_a int) int64            { return int64(_a) }
`,

// ========== ========== ========== ==========
//test keyword scoping: 变,如,否
//test specids: 真,假
`
package main;入"fmt"
var n = 50
变p=70
变string=170
功main(){
  如真{
    fmt.Printf("Len: %d\n", 度("abcdefg") + p)
  }
}
func deputy(){
  if真{
    fmt.Printf("Len: %d\n", 度("abcdefg") + n)
  }
  如假{
    fmt.Printf("Len: %d\n", 度("hijk") + p)
  }否{
    fmt.Printf("Len: %d\n", 度("hi") + p)
  }
  fmt.Printf("Len: %d\n", len("lmnop") + n)
}
`:

// ---------- ---------- ---------- ----------
`package main

import _fmt "fmt"

var n = 50
var _p = 70
var _string = 170

func _main() {
	if true {
		_fmt.Printf("Len: %d\n", len("abcdefg")+_p)
	}
}
func deputy() {
	if true {
		fmt.Printf("Len: %d\n", len("abcdefg")+n)
	}
	if false {
		_fmt.Printf("Len: %d\n", len("hijk")+_p)
	} else {
		_fmt.Printf("Len: %d\n", len("hi")+_p)
	}
	fmt.Printf("Len: %d\n", len("lmnop")+n)
}
`,

// ========== ========== ========== ==========
//test keyword: 构
//test keyword scoping: 种,久
//test specids: 整8,整16,整32
`
package main
种A struct{a string; b 整8}
种B struct{a string; b 整16}
种C构{a string; b 整32}
type D struct{a string; b 整32}
种E构{a串;b整32}
久a=3.1416
const b=2.72
`:

// ---------- ---------- ---------- ----------
`package main

type A struct {
	_a _string
	_b int8
}
type B struct {
	_a _string
	_b int16
}
type C struct {
	_a _string
	_b int32
}
type D struct {
	a string
	b int32
}
type E struct {
	_a string
	_b int32
}

const _a = 3.1416
const b = 2.72
`,

// ========== ========== ========== ==========
//test keywords: 围,为,继,破
//test keyword scoping: 入,图
//test specids: 节,字
`
package main
import "fmt"
var (
	a= 图[字]节{'a': 127, 'b': 0, '7':7}
	b = byte(7)
	c= 图[字]节{'a': b} //FIX: want b to become _b
)
func main(){
	zx:为i:=0;i<19;i++{
		if i==3 {继}; if i==6{破}
		如 i== 16{破zx}
		如 i== 17{继zx}
		fmt.Print(i," ")
	}
	fmt.Println("abc")
	为i:=围a{fmt.Print(i," ")}
	for i:= 0; i<28; i++ {
		if i==3 { continue }
		if i==6 { break }
		fmt.Print(i, " ")
	}
}
`:

// ---------- ---------- ---------- ----------
`package main

import "fmt"

var (
	a = map[rune]byte{'a': 127, 'b': 0, '7': 7}
	b = byte(7)
	c = map[rune]byte{'a': b}
)

func main() {
zx:
	for _i := 0; _i < 19; _i++ {
		if _i == 3 {
			continue
		}
		if _i == 6 {
			break
		}
		if _i == 16 {
			break _zx
		}
		if _i == 17 {
			continue _zx
		}
		_fmt.Print(_i, " ")
	}
	fmt.Println("abc")
	for _i := range _a {
		_fmt.Print(_i, " ")
	}
	for i := 0; i < 28; i++ {
		if i == 3 {
			continue
		}
		if i == 6 {
			break
		}
		fmt.Print(i, " ")
	}
}
`,

// ========== ========== ========== ==========
//test keywords: 掉
//test keyword scoping: 择,事,别,面
//test specids: 双,空,绝,绝8,绝16,绝32,绝64
`package main;入"fmt"
type A interface {
  aMeth()绝
}
种B面{
  bMeth()绝8
}
种C interface{
  cMeth(theC uint16)绝16
}
type D面{
  dMeth(theD绝32)绝64
}
func abc()*双{回空}
func main(){
	a:=2
	择a{
	事1:
		fmt.Print('a');
	事2:
		fmt.Print('b')
		掉
	事3:
		fmt.Print('c')
	别:
		fmt.Print('d')
	}
}
`:

// ---------- ---------- ---------- ----------
`package main

import _fmt "fmt"

type A interface {
	aMeth() uint
}
type B interface {
	_bMeth() uint8
}
type C interface {
	_cMeth(_theC _uint16) uint16
}
type D interface {
	_dMeth(_theD uint32) uint64
}

func abc() *bool { return nil }
func main() {
	a := 2
	switch _a {
	case 1:
		_fmt.Print('a')
	case 2:
		_fmt.Print('b')
		fallthrough
	case 3:
		_fmt.Print('c')
	default:
		_fmt.Print('d')
	}
}
`,

// ========== ========== ========== ==========
//test keyword scoping: 选,去,通
`package main
import (
    "fmt"
    "math/rand"
)
入("sync/atomic";"time")

type readOp struct {
    key  int
    resp 通int
}
type writeOp struct {
    key  int
    val  int
    resp chan bool
}
func main() {
    var ops int64 = 0

    reads := make(通*readOp)
    writes := make(chan *writeOp)

    go func() {
        var state = make(map[int]int)
        for {
            选{
            case read := <-reads:
                read.resp <- state[read.key]
            case write := <-writes:
                state[write.key] = write.val
                write.resp <- true
            }
        }
    }()

    for r := 0; r < 100; r++ {
        去func() {
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

    时Sleep(time.Second)

    opsFinal := atomic.LoadInt64(&ops)
    形Println("ops:", opsFinal)
    形Println("absolute value:", 数Abs(-7.89))
}
`:

// ---------- ---------- ---------- ----------
`package main

import (
	"fmt"
	math "math"
	"math/rand"
)
import (
	_atomic "sync/atomic"
	_time "time"
)

type readOp struct {
	key  int
	resp chan _int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	var ops int64 = 0

	reads := make(chan *_readOp)
	writes := make(chan *writeOp)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case _read := <-_reads:
				_read._resp <- _state[_read._key]
			case _write := <-_writes:
				_state[_write._key] = _write._val
				_write._resp <- _true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			for {
				_read := &_readOp{
					_key:  _rand.Intn(5),
					_resp: _make(chan _int)}
				_reads <- _read
				<-_read._resp
				_atomic.AddInt64(&_ops, 1)
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
	fmt.Println("absolute value:", math.Abs(-7.89))
}
`,

// ========== ========== ========== ==========
`
package main;入"fmt"
功main(){
  让b:=7
  让func:=8
  fmt.Printf("Hi!\n") // comment here
}`:

// ---------- ---------- ---------- ----------
`package main

import _fmt "fmt"

func _main() {
	_b := 7
	_func := 8
	_fmt.Printf("Hi!\n")
}
`,

// ========== ========== ========== ==========
//test keyword scoping: 包
`
包main;入"fmt"
功main(){
  让b:=7
  让func:=8
  fmt.Printf("Hi!\n") // different comment here
}`:

// ---------- ---------- ---------- ----------
`package main

import _fmt "fmt"

func _main() {
	_b := 7
	_func := 8
	_fmt.Printf("Hi!\n")
}
`,

// ========== ========== ========== ==========
`包正;入("math/rand";"sync/atomic")
种readOp构{key整;resp通整}
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
    时Sleep(time.Second)
    opsFinal:=atomic.LoadInt64(&ops)
    形Println("ops:",opsFinal)}
`:

// ---------- ---------- ---------- ----------
"package main\n\nimport (\n\tfmt \"fmt\"\n\t_rand \"math/rand\"\n\t_atomic \"sync/atomic\"\n\ttime \"time\"\n)\n\ntype _readOp struct {\n\t_key  int\n\t_resp chan int\n}\ntype _writeOp struct {\n\t_key  int\n\t_val  int\n\t_resp chan bool\n}\n\nfunc main() {\n\tvar _ops int64 = 0\n\t_reads := make(chan *_readOp)\n\t_writes := make(chan *_writeOp)\n\tgo func() {\n\t\tvar _state = make(map[int]int)\n\t\tfor {\n\t\t\tselect {\n\t\t\tcase _read := <-_reads:\n\t\t\t\t_read._resp <- _state[_read._key]\n\t\t\tcase _write := <-_writes:\n\t\t\t\t_state[_write._key] = _write._val\n\t\t\t\t_write._resp <- true\n\t\t\t}\n\t\t}\n\t}()\n\tfor _r := 0; _r < 100; _r++ {\n\t\tgo func() {\n\t\t\tfor {\n\t\t\t\t_read := &_readOp{_key: _rand.Intn(5), _resp: make(chan int)}\n\t\t\t\t_reads <- _read\n\t\t\t\t<-_read._resp\n\t\t\t\t_atomic.AddInt64(&_ops, 1)\n\t\t\t}\n\t\t}()\n\t}\n\tfor _w := 0; _w < 10; _w++ {\n\t\tgo func() {\n\t\t\tfor {\n\t\t\t\t_write := &_writeOp{_key: _rand.Intn(5), _val: _rand.Intn(100), _resp: make(chan bool)}\n\t\t\t\t_writes <- _write\n\t\t\t\t<-_write._resp\n\t\t\t\t_atomic.AddInt64(&_ops, 1)\n\t\t\t}\n\t\t}()\n\t}\n\ttime.Sleep(_time.Second)\n\t_opsFinal := _atomic.LoadInt64(&_ops)\n\tfmt.Println(\"ops:\", _opsFinal)\n}\n",

// ========== ========== ========== ==========
`package main

import (
	"fmt"
	"github.com/gavingroovygrover/qu/parser"
	"go/token"
	"go/format"
	//"go/ast"
	//"os"
)

func main() {
	形Printf("Hi!\n")

	fset := token.NewFileSet() // positions are relative to fset

	//f, err := parser.ParseFile(fset, "src/github.com/gavingroovygrover/qu/first.go", nil, parser.ImportsOnly)
	//f, err := parser.ParseFile(fset, "", src, parser.ImportsOnly)
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, s := range f.Imports {
		fmt.Println(s.Path.Value)
	}

	//ast.Print(fset, f)
	//_= format.Node(os.Stdout, fset, f)

	sw:= StringWriter{""}
	_= format.Node(&sw, fset, f)
	fmt.Println(sw.data)
	fmt.Println(sw.data == dst)
}

type StringWriter struct{ data string }
func (sw *StringWriter) Write(b []byte) (int, error) {
	sw.data += string(b)
	return len(b), nil
}

var src string = "package main"`:

// ---------- ---------- ---------- ----------
"",

// ========== ========== ========== ==========
}
/*
more tests: 让
more keyword scoping tests: 包
test more default packages: 数,大,网,序,形
to test keyword scoping: 为,终,回,破,继,跳,构
convert to specid scoping: all except 真,假,空,毫
to test specids: 能,实,虚,造,新,关,加,副,删,丢,抓,写,线,毫,镇,错
convert to keyword scoping: 围
fix keyword scoping: 图
aliases: 任
fix error where pre-existing imports aren't re-added with different name
----------
test keywords as ids and labels in kanji-context
test blank: _
pseudo-keywords: 这正
----------
plug into qu/tools
write doco
release
*/

