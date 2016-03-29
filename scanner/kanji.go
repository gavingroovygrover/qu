// Copyright 2016 Gavin "Groovy" Grover. All rights reserved.
// Use of this source code is governed by the same BSD-style
// license as Go that can be found in the LICENSE file.

package scanner

import(
)

const (
	KeywordKanji = iota + 1
	IdentifierKanji
	PackageKanji
	TentativeKanji
)

type KanjiVal struct {
	Kind           int
	Word           string
	IsSuffixable   bool
	IsScoped       bool
	IsGoReserved   bool
}

// We put all the Kanjis into a single map so we'll always know
// every Kanji has a single unique meaning.
type KanjiMap map[rune]KanjiVal

var SuffixedIdents = map[string]string{
	"整": "int",
	"整8": "int8",
	"整16": "int16",
	"整32": "int32",
	"整64": "int64",

	"绝": "uint",
	"绝8": "uint8",
	"绝16": "uint16",
	"绝32": "uint32",
	"绝64": "uint64",

	"漂32": "float32",
	"漂64": "float64",

	"复": "complex",
	"复64": "complex64",
	"复128": "complex128",
}

var Kanjis = KanjiMap {
	'包': {Kind:KeywordKanji, Word:"package", IsGoReserved:true},
	'入': {Kind:KeywordKanji, Word:"import", IsGoReserved:true},
	'久': {Kind:KeywordKanji, Word:"const", IsGoReserved:true},
	'变': {Kind:KeywordKanji, Word:"var", IsGoReserved:true},
	'种': {Kind:KeywordKanji, Word:"type", IsGoReserved:true},
	'功': {Kind:KeywordKanji, Word:"func", IsGoReserved:true},
	'构': {Kind:KeywordKanji, Word:"struct", IsGoReserved:true},
	'图': {Kind:KeywordKanji, Word:"map", IsGoReserved:true},
	'面': {Kind:KeywordKanji, Word:"interface", IsGoReserved:true},
	'通': {Kind:KeywordKanji, Word:"chan", IsGoReserved:true},

	'如': {Kind:KeywordKanji, Word:"if", IsGoReserved:true},
	'否': {Kind:KeywordKanji, Word:"else", IsGoReserved:true},
	'择': {Kind:KeywordKanji, Word:"switch", IsGoReserved:true},
	'事': {Kind:KeywordKanji, Word:"case", IsGoReserved:true},
	'别': {Kind:KeywordKanji, Word:"default", IsGoReserved:true},
	'掉': {Kind:KeywordKanji, Word:"fallthrough", IsGoReserved:true},
	'选': {Kind:KeywordKanji, Word:"select", IsGoReserved:true},
	'为': {Kind:KeywordKanji, Word:"for", IsGoReserved:true},
	'围': {Kind:KeywordKanji, Word:"range", IsGoReserved:true},
	'终': {Kind:KeywordKanji, Word:"defer", IsGoReserved:true},
	'去': {Kind:KeywordKanji, Word:"go", IsGoReserved:true},
	'回': {Kind:KeywordKanji, Word:"return", IsGoReserved:true},
	'破': {Kind:KeywordKanji, Word:"break", IsGoReserved:true},
	'继': {Kind:KeywordKanji, Word:"continue", IsGoReserved:true},
	'跳': {Kind:KeywordKanji, Word:"goto", IsGoReserved:true},

	'真': {Kind:IdentifierKanji, Word:"true", IsGoReserved:true},
	'假': {Kind:IdentifierKanji, Word:"false", IsGoReserved:true},
	'空': {Kind:IdentifierKanji, Word:"nil", IsGoReserved:true},
	'毫': {Kind:IdentifierKanji, Word:"iota", IsGoReserved:true},

	'能': {Kind:IdentifierKanji, Word:"cap", IsGoReserved:true, IsScoped:true},
	'度': {Kind:IdentifierKanji, Word:"len", IsGoReserved:true, IsScoped:true},
	'实': {Kind:IdentifierKanji, Word:"real", IsGoReserved:true, IsScoped:true},
	'虚': {Kind:IdentifierKanji, Word:"imag", IsGoReserved:true, IsScoped:true},
	'造': {Kind:IdentifierKanji, Word:"make", IsGoReserved:true, IsScoped:true},
	'新': {Kind:IdentifierKanji, Word:"new", IsGoReserved:true, IsScoped:true},
	'关': {Kind:IdentifierKanji, Word:"close", IsGoReserved:true, IsScoped:true},
	'加': {Kind:IdentifierKanji, Word:"append", IsGoReserved:true, IsScoped:true},
	'副': {Kind:IdentifierKanji, Word:"copy", IsGoReserved:true, IsScoped:true},
	'删': {Kind:IdentifierKanji, Word:"delete", IsGoReserved:true, IsScoped:true},
	'丢': {Kind:IdentifierKanji, Word:"panic", IsGoReserved:true, IsScoped:true},
	'抓': {Kind:IdentifierKanji, Word:"recover", IsGoReserved:true, IsScoped:true},
	'写': {Kind:IdentifierKanji, Word:"print", IsGoReserved:true, IsScoped:true},
	'线': {Kind:IdentifierKanji, Word:"println", IsGoReserved:true, IsScoped:true},

	'节': {Kind:IdentifierKanji, Word:"byte", IsGoReserved:true, IsScoped:true},
	'字': {Kind:IdentifierKanji, Word:"rune", IsGoReserved:true, IsScoped:true},
	'串': {Kind:IdentifierKanji, Word:"string", IsGoReserved:true, IsScoped:true},
	'双': {Kind:IdentifierKanji, Word:"bool", IsGoReserved:true, IsScoped:true},
	'错': {Kind:IdentifierKanji, Word:"error", IsGoReserved:true, IsScoped:true},
	'镇': {Kind:IdentifierKanji, Word:"uintptr", IsGoReserved:true, IsScoped:true},

	//suffixable identifiers...
	'整': {Kind:IdentifierKanji, Word:"int", IsGoReserved:true, IsScoped:true, IsSuffixable:true}, //int, int8, int16, int32, int64
	'绝': {Kind:IdentifierKanji, Word:"uint", IsGoReserved:true, IsScoped:true, IsSuffixable:true}, //uint, uint8, uint16, uint32, uint64
	'漂': {Kind:IdentifierKanji, Word:"float", IsGoReserved:true, IsScoped:true, IsSuffixable:true}, //float32, float64
	'复': {Kind:IdentifierKanji, Word:"complex", IsGoReserved:true, IsScoped:true, IsSuffixable:true}, //complex, complex64, complex128

	'让': {Kind:IdentifierKanji}, //verbalize as "let"
	'做': {Kind:IdentifierKanji}, //verbalize as "do"
	'正': {Kind:IdentifierKanji, Word:"main"},
	'任': {Kind:IdentifierKanji}, //"interface{}" returned; verbalize as "any"

	//packages...
	'形': {Kind:PackageKanji, Word:"fmt"},
	'网': {Kind:PackageKanji, Word:"net"},
	'序': {Kind:PackageKanji, Word:"sort"},
	'数': {Kind:PackageKanji, Word:"math"},
	'大': {Kind:PackageKanji, Word:"math/big"},
	'时': {Kind:PackageKanji, Word:"time"},

	//tentative kanji...
	'这': {Kind:TentativeKanji, Word:"this"},
	'特': {Kind:TentativeKanji, Word:"special"},
	'愿': {Kind:TentativeKanji, Word:"source"},
	'试': {Kind:TentativeKanji, Word:"try"},
	'具': {Kind:TentativeKanji, Word:"util"},
	'动': {Kind:TentativeKanji, Word:"dyn"},
	'指': {Kind:TentativeKanji, Word:"spec"},
	'羔': {Kind:TentativeKanji, Word:"lamb"},
	'程': {Kind:TentativeKanji, Word:"proc"},
	'对': {Kind:TentativeKanji, Word:"assert"},
	'用': {Kind:TentativeKanji, Word:"use"},
	'准': {Kind:TentativeKanji, Word:"prepare"},
	'执': {Kind:TentativeKanji, Word:"execute"},
	'冲': {Kind:TentativeKanji, Word:"flush"},
	'建': {Kind:TentativeKanji, Word:"build"},
	'跑': {Kind:TentativeKanji, Word:"run"},
	'考': {Kind:TentativeKanji, Word:"test"},
	'洗': {Kind:TentativeKanji, Word:"clean"},
	'出': {Kind:TentativeKanji, Word:"exit"},
	'显': {Kind:TentativeKanji, Word:"vars"},
	'后': {Kind:TentativeKanji, Word:"next"},
	'前': {Kind:TentativeKanji, Word:"prev"},
	'学': {Kind:TentativeKanji, Word:"learn"},
	'解': {Kind:TentativeKanji, Word:"parse"},
	'类': {Kind:TentativeKanji, Word:"class"},
	'叫': {Kind:TentativeKanji, Word:"call"},
	'是': {Kind:TentativeKanji, Word:"is"},
	'侯': {Kind:TentativeKanji, Word:"while"},
	'它': {Kind:TentativeKanji, Word:"it"},
	'自': {Kind:TentativeKanji, Word:"self"},
	'滤': {Kind:TentativeKanji, Word:"filter"},
	'减': {Kind:TentativeKanji, Word:"reduce"},
	'组': {Kind:TentativeKanji, Word:"groupby"},
	'颠': {Kind:TentativeKanji, Word:"reverse"},
	'长': {Kind:TentativeKanji, Word:"long"},
	'除': {Kind:TentativeKanji, Word:"exception"},
	'摸': {Kind:TentativeKanji, Word:"pattern"},
}

