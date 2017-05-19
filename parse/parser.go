//line parser.go.y:2
package parse

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"github.com/lunixbochs/luaish/ast"
)

//line parser.go.y:37
type yySymType struct {
	yys   int
	token ast.Token

	stmts []ast.Stmt
	stmt  ast.Stmt

	funcname *ast.FuncName
	funcexpr *ast.FunctionExpr

	exprlist []ast.Expr
	expr     ast.Expr

	fieldlist []*ast.Field
	field     *ast.Field
	fieldsep  string

	namelist []string
	parlist  *ast.ParList
}

const TAnd = 57346
const TBreak = 57347
const TDo = 57348
const TElse = 57349
const TElseIf = 57350
const TEnd = 57351
const TFalse = 57352
const TFor = 57353
const TFunction = 57354
const TIf = 57355
const TIn = 57356
const TLocal = 57357
const TNil = 57358
const TNot = 57359
const TOr = 57360
const TReturn = 57361
const TRepeat = 57362
const TThen = 57363
const TTrue = 57364
const TUntil = 57365
const TWhile = 57366
const TOn = 57367
const TEqeq = 57368
const TNeq = 57369
const TLte = 57370
const TGte = 57371
const T2Comma = 57372
const T3Comma = 57373
const TIdent = 57374
const TIdentSpace = 57375
const TNumber = 57376
const TString = 57377
const TShr = 57378
const TShl = 57379
const UNARY = 57380
const TPow = 57381

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"TAnd",
	"TBreak",
	"TDo",
	"TElse",
	"TElseIf",
	"TEnd",
	"TFalse",
	"TFor",
	"TFunction",
	"TIf",
	"TIn",
	"TLocal",
	"TNil",
	"TNot",
	"TOr",
	"TReturn",
	"TRepeat",
	"TThen",
	"TTrue",
	"TUntil",
	"TWhile",
	"TOn",
	"TEqeq",
	"TNeq",
	"TLte",
	"TGte",
	"T2Comma",
	"T3Comma",
	"TIdent",
	"TIdentSpace",
	"TNumber",
	"TString",
	"'{'",
	"'('",
	"TShr",
	"TShl",
	"'~'",
	"'^'",
	"'|'",
	"'&'",
	"'>'",
	"'<'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"UNARY",
	"TPow",
	"';'",
	"'='",
	"','",
	"':'",
	"'.'",
	"'['",
	"']'",
	"'#'",
	"')'",
	"'}'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.go.y:619

func TokenName(c int) string {
	if c >= TAnd && c-TAnd < len(yyToknames) {
		if yyToknames[c-TAnd] != "" {
			return yyToknames[c-TAnd]
		}
	}
	return string([]byte{byte(c)})
}

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 19,
	54, 37,
	55, 37,
	-2, 80,
	-1, 110,
	54, 38,
	55, 38,
	-2, 80,
}

const yyPrivate = 57344

const yyLast = 839

var yyAct = [...]int{

	29, 24, 59, 100, 104, 178, 28, 65, 156, 25,
	58, 57, 161, 61, 160, 63, 62, 64, 68, 71,
	51, 158, 75, 72, 57, 70, 38, 37, 166, 132,
	127, 128, 130, 125, 47, 48, 54, 55, 155, 75,
	96, 97, 98, 99, 45, 46, 53, 107, 105, 188,
	112, 180, 113, 192, 108, 162, 124, 118, 123, 26,
	116, 101, 114, 54, 55, 52, 50, 49, 121, 36,
	126, 129, 9, 25, 58, 118, 133, 134, 135, 136,
	137, 138, 139, 140, 141, 142, 143, 144, 145, 146,
	147, 148, 149, 150, 151, 152, 153, 125, 75, 213,
	210, 92, 91, 71, 93, 94, 95, 163, 56, 157,
	205, 87, 88, 89, 204, 90, 25, 58, 111, 18,
	167, 22, 169, 168, 171, 170, 27, 173, 21, 174,
	175, 172, 74, 176, 69, 165, 159, 177, 54, 55,
	44, 54, 55, 19, 45, 46, 53, 191, 25, 58,
	31, 73, 43, 198, 25, 58, 30, 41, 179, 190,
	107, 105, 32, 182, 181, 92, 91, 109, 93, 94,
	95, 34, 25, 58, 33, 45, 46, 22, 185, 90,
	39, 189, 173, 194, 195, 193, 184, 40, 196, 110,
	119, 197, 103, 199, 77, 154, 201, 200, 106, 35,
	42, 20, 102, 8, 208, 207, 60, 1, 76, 209,
	67, 66, 3, 186, 212, 4, 82, 83, 81, 80,
	84, 2, 0, 0, 0, 0, 0, 0, 92, 91,
	77, 93, 94, 95, 78, 79, 85, 86, 87, 88,
	89, 0, 90, 0, 76, 0, 0, 0, 0, 0,
	0, 131, 82, 83, 81, 80, 84, 0, 0, 0,
	0, 0, 0, 0, 92, 91, 77, 93, 94, 95,
	78, 79, 85, 86, 87, 88, 89, 0, 90, 0,
	76, 0, 0, 0, 0, 183, 0, 0, 82, 83,
	81, 80, 84, 0, 0, 0, 0, 0, 0, 0,
	92, 91, 0, 93, 94, 95, 78, 79, 85, 86,
	87, 88, 89, 77, 90, 202, 0, 0, 0, 0,
	0, 164, 0, 0, 0, 0, 0, 76, 0, 0,
	0, 0, 0, 0, 0, 82, 83, 81, 80, 84,
	0, 0, 0, 0, 0, 0, 0, 92, 91, 0,
	93, 94, 95, 78, 79, 85, 86, 87, 88, 89,
	31, 90, 43, 0, 203, 0, 30, 41, 0, 0,
	0, 0, 32, 0, 0, 0, 0, 0, 77, 0,
	0, 34, 25, 58, 33, 45, 46, 22, 0, 0,
	39, 0, 76, 0, 0, 0, 0, 40, 0, 0,
	82, 83, 81, 80, 84, 0, 0, 0, 0, 0,
	42, 115, 92, 91, 0, 93, 94, 95, 78, 79,
	85, 86, 87, 88, 89, 31, 90, 43, 0, 187,
	0, 30, 41, 0, 0, 0, 0, 32, 0, 0,
	0, 0, 0, 0, 0, 0, 34, 25, 58, 33,
	45, 46, 22, 117, 0, 39, 0, 31, 0, 43,
	0, 0, 40, 30, 41, 0, 0, 0, 0, 32,
	0, 0, 0, 106, 0, 42, 0, 0, 34, 25,
	58, 33, 45, 46, 22, 0, 0, 39, 0, 31,
	0, 43, 0, 0, 40, 30, 41, 0, 0, 0,
	0, 32, 0, 0, 0, 0, 0, 42, 0, 0,
	34, 25, 58, 33, 45, 46, 22, 0, 0, 39,
	0, 31, 0, 43, 0, 0, 40, 30, 41, 0,
	0, 0, 0, 32, 0, 0, 0, 0, 77, 42,
	211, 0, 34, 25, 23, 33, 45, 46, 22, 0,
	0, 39, 76, 0, 0, 0, 0, 0, 40, 0,
	82, 83, 81, 80, 84, 0, 0, 0, 0, 0,
	0, 42, 92, 91, 77, 93, 94, 95, 78, 79,
	85, 86, 87, 88, 89, 0, 90, 0, 76, 0,
	0, 206, 0, 0, 0, 0, 82, 83, 81, 80,
	84, 0, 0, 0, 0, 0, 0, 77, 92, 91,
	0, 93, 94, 95, 78, 79, 85, 86, 87, 88,
	89, 76, 90, 0, 122, 0, 0, 0, 0, 82,
	83, 81, 80, 84, 0, 0, 0, 0, 0, 0,
	0, 92, 91, 0, 93, 94, 95, 78, 79, 85,
	86, 87, 88, 89, 77, 90, 120, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 76, 0,
	0, 0, 0, 0, 0, 0, 82, 83, 81, 80,
	84, 0, 0, 0, 0, 0, 0, 77, 92, 91,
	0, 93, 94, 95, 78, 79, 85, 86, 87, 88,
	89, 76, 90, 0, 0, 0, 0, 0, 0, 82,
	83, 81, 80, 84, 77, 0, 0, 0, 0, 0,
	0, 92, 91, 0, 93, 94, 95, 78, 79, 85,
	86, 87, 88, 89, 0, 90, 82, 83, 81, 80,
	84, 0, 0, 0, 0, 0, 0, 0, 92, 91,
	0, 93, 94, 95, 78, 79, 85, 86, 87, 88,
	89, 0, 90, 82, 83, 81, 80, 84, 0, 0,
	0, 0, 0, 0, 0, 92, 91, 0, 93, 94,
	95, 78, 79, 85, 86, 87, 88, 89, 0, 90,
	7, 11, 0, 0, 0, 0, 15, 16, 14, 0,
	17, 0, 0, 0, 6, 13, 0, 0, 0, 12,
	10, 0, 84, 0, 0, 0, 0, 25, 23, 0,
	92, 91, 22, 93, 94, 95, 0, 0, 85, 86,
	87, 88, 89, 0, 90, 0, 0, 0, 5,
}
var yyPact = [...]int{

	-1000, -1000, 785, 6, -1000, -1000, 511, -1000, -20, 9,
	479, -1000, 479, -1000, 479, 41, 41, 122, -1000, -1000,
	-1000, -1000, 479, 479, -1000, -1000, -1000, -1000, -16, 683,
	-1000, -1000, -1000, -1000, -1000, -1000, 9, -1000, -1000, 479,
	479, 479, 479, 24, -1000, -1000, 140, 511, 84, 479,
	41, -1000, 30, 350, -1000, -1000, 447, 683, -1000, 181,
	-1000, 650, 45, 603, 4, 42, 24, -26, -1000, 41,
	-22, -1000, 190, -32, 479, 479, 479, 479, 479, 479,
	479, 479, 479, 479, 479, 479, 479, 479, 479, 479,
	479, 479, 479, 479, 479, 479, 127, 127, 127, 127,
	-1000, -23, -1000, -41, -1000, 1, 479, 683, -16, -1000,
	-1000, 9, 262, -1000, 109, -1000, -33, -1000, 683, -1000,
	-1000, 479, -1000, 479, 479, 41, -1000, 41, 41, 24,
	479, -1000, -1000, 683, 710, 737, 782, 782, 782, 782,
	782, 782, 782, 63, 63, 127, 127, 127, 127, -1000,
	-1000, -1000, -1000, -1000, -56, -1000, -1000, -4, -1000, 415,
	-1000, -1000, 479, 226, -1000, -1000, -1000, 177, 169, 683,
	-1000, 374, 43, -1000, -1000, -1000, -1000, -16, -1000, 150,
	116, -1000, 683, -1, -1000, -1000, 176, 479, -1000, 144,
	-1000, -1000, 479, -1000, -1000, 479, 309, 105, -1000, 683,
	101, 570, -1000, 479, -1000, -1000, -1000, 91, 534, -1000,
	-1000, -1000, 90, -1000,
}
var yyPgo = [...]int{

	0, 206, 221, 2, 215, 213, 212, 211, 210, 203,
	140, 7, 6, 0, 27, 69, 128, 201, 20, 119,
	108, 199, 3, 195, 26, 192, 4, 136, 1,
}
var yyR1 = [...]int{

	0, 1, 1, 1, 2, 2, 2, 3, 28, 28,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 5, 5, 6,
	6, 6, 6, 7, 7, 8, 8, 9, 9, 10,
	10, 10, 11, 11, 12, 12, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 14,
	15, 15, 15, 15, 17, 16, 16, 19, 20, 20,
	18, 18, 18, 18, 21, 22, 22, 23, 23, 23,
	24, 24, 25, 25, 25, 26, 26, 26, 27, 27,
}
var yyR2 = [...]int{

	0, 1, 2, 3, 0, 2, 2, 1, 1, 1,
	3, 3, 1, 5, 3, 5, 4, 6, 8, 9,
	11, 7, 3, 4, 4, 2, 1, 0, 5, 2,
	1, 2, 1, 1, 3, 1, 3, 1, 3, 1,
	4, 3, 1, 3, 1, 3, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 2, 2, 2, 2, 1,
	1, 1, 1, 3, 3, 2, 4, 2, 1, 2,
	2, 3, 1, 1, 2, 5, 4, 1, 1, 3,
	2, 3, 1, 3, 2, 3, 5, 1, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -2, -6, -4, 53, 19, 5, -9, -15,
	25, 6, 24, 20, 13, 11, 12, 15, -19, -10,
	-17, -16, 37, 33, -28, 32, 53, -19, -12, -13,
	16, 10, 22, 34, 31, -21, -15, -14, -24, 40,
	47, 17, 60, 12, -10, 35, 36, 54, 55, 58,
	57, -18, 56, 37, -24, -14, -20, -13, 33, -3,
	-1, -13, -3, -13, -28, -11, -7, -8, -28, 12,
	-11, -28, -13, -16, -20, 55, 18, 4, 44, 45,
	29, 28, 26, 27, 30, 46, 47, 48, 49, 50,
	52, 39, 38, 41, 42, 43, -13, -13, -13, -13,
	-22, 37, 62, -25, -26, -28, 58, -13, -12, -19,
	-10, -15, -13, -28, 32, 61, -12, 6, -13, 9,
	6, 23, 21, 54, 14, 55, -22, 56, 57, -28,
	54, 61, 61, -13, -13, -13, -13, -13, -13, -13,
	-13, -13, -13, -13, -13, -13, -13, -13, -13, -13,
	-13, -13, -13, -13, -23, 61, 31, -11, 62, -27,
	55, 53, 54, -13, 59, -18, 61, -3, -3, -13,
	-3, -13, -12, -28, -28, -28, -22, -12, 61, -3,
	55, -26, -13, 59, 9, 9, -5, 55, 6, -3,
	9, 31, 54, 9, 7, 8, -13, -3, 9, -13,
	-3, -13, 6, 55, 9, 9, 21, -3, -13, -3,
	9, 6, -3, 9,
}
var yyDef = [...]int{

	4, -2, 1, 2, 5, 6, 30, 32, 0, 12,
	0, 4, 0, 4, 0, 0, 0, 0, 26, -2,
	81, 82, 0, 8, 39, 9, 3, 29, 31, 44,
	46, 47, 48, 49, 50, 51, 52, 53, 54, 0,
	0, 0, 0, 0, 80, 79, 0, 0, 0, 0,
	0, 85, 0, 0, 92, 93, 0, 88, 8, 0,
	7, 0, 0, 0, 42, 0, 0, 33, 35, 0,
	25, 42, 0, 82, 87, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 75, 76, 77, 78,
	94, 0, 100, 0, 102, 39, 0, 107, 10, 11,
	-2, 0, 0, 41, 0, 90, 0, 4, 89, 14,
	4, 0, 4, 0, 0, 0, 22, 0, 0, 0,
	0, 83, 84, 45, 55, 56, 57, 58, 59, 60,
	61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
	71, 72, 73, 74, 0, 4, 97, 98, 101, 104,
	108, 109, 0, 0, 40, 86, 91, 0, 0, 16,
	27, 0, 0, 43, 34, 36, 23, 24, 4, 0,
	0, 103, 105, 0, 13, 15, 0, 0, 4, 0,
	96, 99, 0, 17, 4, 0, 0, 0, 95, 106,
	0, 0, 4, 0, 21, 18, 4, 0, 0, 28,
	19, 4, 0, 20,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 60, 3, 50, 43, 3,
	37, 61, 48, 46, 55, 47, 57, 49, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 56, 53,
	45, 54, 44, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 58, 3, 59, 41, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 36, 42, 62, 40,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 38, 39, 51, 52,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:82
		{
			yyVAL.stmts = yyDollar[1].stmts
			if l, ok := yylex.(*Lexer); ok {
				l.Stmts = yyVAL.stmts
			}
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:88
		{
			yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[2].stmt)
			if l, ok := yylex.(*Lexer); ok {
				l.Stmts = yyVAL.stmts
			}
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:94
		{
			yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[2].stmt)
			if l, ok := yylex.(*Lexer); ok {
				l.Stmts = yyVAL.stmts
			}
		}
	case 4:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:102
		{
			yyVAL.stmts = []ast.Stmt{}
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:105
		{
			yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[2].stmt)
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:108
		{
			yyVAL.stmts = yyDollar[1].stmts
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:113
		{
			yyVAL.stmts = yyDollar[1].stmts
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:118
		{
			yyVAL.token = yyDollar[1].token
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:119
		{
			yyVAL.token = yyDollar[1].token
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:122
		{
			yyVAL.stmt = &ast.AssignStmt{Lhs: yyDollar[1].exprlist, Rhs: yyDollar[3].exprlist}
			yyVAL.stmt.SetLine(yyDollar[1].exprlist[0].Line())
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:126
		{
			yyVAL.stmt = &ast.AssignStmt{Lhs: yyDollar[1].exprlist, Rhs: []ast.Expr{yyDollar[3].expr}}
			yyVAL.stmt.SetLine(yyDollar[1].exprlist[0].Line())
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:131
		{
			if _, ok := yyDollar[1].expr.(*ast.FuncCallExpr); !ok {
				// if this is just a simple ident, it can be executed as a function (anywhere)
				// otherwise this expression should be evaluated and printed if in a repl
				// this isn't a return statement, because it needs to support bare function calls in the middle of a block
				args := []ast.Expr{yyDollar[1].expr}
				if v, ok := yyDollar[1].expr.(*ast.IdentExpr); ok {
					args = []ast.Expr{
						&ast.StringExpr{Value: v.Value}, yyDollar[1].expr,
					}
				}
				yyVAL.stmt = &ast.FuncCallStmt{Expr: &ast.FuncCallExpr{
					Func: &ast.IdentExpr{Value: "_fallback"},
					Args: args,
				}}
			} else {
				yyVAL.stmt = &ast.FuncCallStmt{Expr: yyDollar[1].expr}
				yyVAL.stmt.SetLine(yyDollar[1].expr.Line())
			}
		}
	case 13:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:151
		{
			// used for `on step [start] [end] do end`-style hooks
			// will be converted into `_on_hook("<keyword>", func() ... end, args...)`
			var keyword ast.Expr
			kw := yyDollar[2].exprlist[0]
			switch v := kw.(type) {
			case *ast.IdentExpr:
				keyword = &ast.StringExpr{Value: v.Value}
			case *ast.NumberExpr, *ast.StringExpr:
				keyword = v
			default:
				yylex.(*Lexer).Error("parse error: first argument to `on` is number or identifier")
			}
			fn := &ast.FunctionExpr{
				ParList: &ast.ParList{Names: []string{}},
				Stmts:   yyDollar[4].stmts,
			}
			args := []ast.Expr{keyword, fn, &ast.NilExpr{}, &ast.NilExpr{}}
			copy(args[2:], yyDollar[2].exprlist[1:])

			yyVAL.stmt = &ast.FuncCallStmt{
				Expr: &ast.FuncCallExpr{
					Func: &ast.IdentExpr{Value: "_on_hook"},
					Args: args,
				},
			}
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:178
		{
			yyVAL.stmt = &ast.DoBlockStmt{Stmts: yyDollar[2].stmts}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[3].token.Pos.Line)
		}
	case 15:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:183
		{
			yyVAL.stmt = &ast.WhileStmt{Condition: yyDollar[2].expr, Stmts: yyDollar[4].stmts}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[5].token.Pos.Line)
		}
	case 16:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:188
		{
			yyVAL.stmt = &ast.RepeatStmt{Condition: yyDollar[4].expr, Stmts: yyDollar[2].stmts}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[4].expr.Line())
		}
	case 17:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:193
		{
			yyVAL.stmt = &ast.IfStmt{Condition: yyDollar[2].expr, Then: yyDollar[4].stmts}
			cur := yyVAL.stmt
			for _, elseif := range yyDollar[5].stmts {
				cur.(*ast.IfStmt).Else = []ast.Stmt{elseif}
				cur = elseif
			}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[6].token.Pos.Line)
		}
	case 18:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:203
		{
			yyVAL.stmt = &ast.IfStmt{Condition: yyDollar[2].expr, Then: yyDollar[4].stmts}
			cur := yyVAL.stmt
			for _, elseif := range yyDollar[5].stmts {
				cur.(*ast.IfStmt).Else = []ast.Stmt{elseif}
				cur = elseif
			}
			cur.(*ast.IfStmt).Else = yyDollar[7].stmts
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[8].token.Pos.Line)
		}
	case 19:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:214
		{
			yyVAL.stmt = &ast.NumberForStmt{Name: yyDollar[2].token.Str, Init: yyDollar[4].expr, Limit: yyDollar[6].expr, Stmts: yyDollar[8].stmts}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[9].token.Pos.Line)
		}
	case 20:
		yyDollar = yyS[yypt-11 : yypt+1]
		//line parser.go.y:219
		{
			yyVAL.stmt = &ast.NumberForStmt{Name: yyDollar[2].token.Str, Init: yyDollar[4].expr, Limit: yyDollar[6].expr, Step: yyDollar[8].expr, Stmts: yyDollar[10].stmts}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[11].token.Pos.Line)
		}
	case 21:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:224
		{
			yyVAL.stmt = &ast.GenericForStmt{Names: yyDollar[2].namelist, Exprs: yyDollar[4].exprlist, Stmts: yyDollar[6].stmts}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[7].token.Pos.Line)
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:229
		{
			yyVAL.stmt = &ast.FuncDefStmt{Name: yyDollar[2].funcname, Func: yyDollar[3].funcexpr}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[3].funcexpr.LastLine())
		}
	case 23:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:234
		{
			yyVAL.stmt = &ast.LocalAssignStmt{Names: []string{yyDollar[3].token.Str}, Exprs: []ast.Expr{yyDollar[4].funcexpr}}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[4].funcexpr.LastLine())
		}
	case 24:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:239
		{
			yyVAL.stmt = &ast.LocalAssignStmt{Names: yyDollar[2].namelist, Exprs: yyDollar[4].exprlist}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:243
		{
			yyVAL.stmt = &ast.LocalAssignStmt{Names: yyDollar[2].namelist, Exprs: []ast.Expr{}}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:247
		{
			yyVAL.stmt = &ast.FuncCallStmt{Expr: yyDollar[1].expr}
			yyVAL.stmt.SetLine(yyDollar[1].expr.Line())
		}
	case 27:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:253
		{
			yyVAL.stmts = []ast.Stmt{}
		}
	case 28:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:256
		{
			yyVAL.stmts = append(yyDollar[1].stmts, &ast.IfStmt{Condition: yyDollar[3].expr, Then: yyDollar[5].stmts})
			yyVAL.stmts[len(yyVAL.stmts)-1].SetLine(yyDollar[2].token.Pos.Line)
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:262
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: []ast.Expr{yyDollar[2].expr}}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:266
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: nil}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:270
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyDollar[2].exprlist}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:274
		{
			yyVAL.stmt = &ast.BreakStmt{}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:280
		{
			yyVAL.funcname = yyDollar[1].funcname
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:283
		{
			yyVAL.funcname = &ast.FuncName{Func: nil, Receiver: yyDollar[1].funcname.Func, Method: yyDollar[3].token.Str}
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:288
		{
			yyVAL.funcname = &ast.FuncName{Func: &ast.IdentExpr{Value: yyDollar[1].token.Str}}
			yyVAL.funcname.Func.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:292
		{
			key := &ast.StringExpr{Value: yyDollar[3].token.Str}
			key.SetLine(yyDollar[3].token.Pos.Line)
			fn := &ast.AttrGetExpr{Object: yyDollar[1].funcname.Func, Key: key}
			fn.SetLine(yyDollar[3].token.Pos.Line)
			yyVAL.funcname = &ast.FuncName{Func: fn}
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:301
		{
			yyVAL.exprlist = []ast.Expr{yyDollar[1].expr}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:304
		{
			yyVAL.exprlist = append(yyDollar[1].exprlist, yyDollar[3].expr)
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:309
		{
			yyVAL.expr = &ast.IdentExpr{Value: yyDollar[1].token.Str}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 40:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:313
		{
			yyVAL.expr = &ast.AttrGetExpr{Object: yyDollar[1].expr, Key: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:317
		{
			key := &ast.StringExpr{Value: yyDollar[3].token.Str}
			key.SetLine(yyDollar[3].token.Pos.Line)
			yyVAL.expr = &ast.AttrGetExpr{Object: yyDollar[1].expr, Key: key}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:325
		{
			yyVAL.namelist = []string{yyDollar[1].token.Str}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:328
		{
			yyVAL.namelist = append(yyDollar[1].namelist, yyDollar[3].token.Str)
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:333
		{
			yyVAL.exprlist = []ast.Expr{yyDollar[1].expr}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:336
		{
			yyVAL.exprlist = append(yyDollar[1].exprlist, yyDollar[3].expr)
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:341
		{
			yyVAL.expr = &ast.NilExpr{}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:345
		{
			yyVAL.expr = &ast.FalseExpr{}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:349
		{
			yyVAL.expr = &ast.TrueExpr{}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:353
		{
			yyVAL.expr = &ast.NumberExpr{Value: yyDollar[1].token.Str}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:357
		{
			yyVAL.expr = &ast.Comma3Expr{}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:361
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:364
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:367
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:370
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:373
		{
			yyVAL.expr = &ast.LogicalOpExpr{Lhs: yyDollar[1].expr, Operator: "or", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:377
		{
			yyVAL.expr = &ast.LogicalOpExpr{Lhs: yyDollar[1].expr, Operator: "and", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:381
		{
			yyVAL.expr = &ast.RelationalOpExpr{Lhs: yyDollar[1].expr, Operator: ">", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:385
		{
			yyVAL.expr = &ast.RelationalOpExpr{Lhs: yyDollar[1].expr, Operator: "<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:389
		{
			yyVAL.expr = &ast.RelationalOpExpr{Lhs: yyDollar[1].expr, Operator: ">=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:393
		{
			yyVAL.expr = &ast.RelationalOpExpr{Lhs: yyDollar[1].expr, Operator: "<=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:397
		{
			yyVAL.expr = &ast.RelationalOpExpr{Lhs: yyDollar[1].expr, Operator: "==", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:401
		{
			yyVAL.expr = &ast.RelationalOpExpr{Lhs: yyDollar[1].expr, Operator: "!=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:405
		{
			yyVAL.expr = &ast.StringConcatOpExpr{Lhs: yyDollar[1].expr, Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:409
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "+", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:413
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "-", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:417
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "*", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:421
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "/", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:425
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "%", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:429
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "**", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:433
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "<<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:437
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: ">>", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:441
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "^", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:445
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "|", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:449
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:453
		{
			yyVAL.expr = &ast.UnaryNegateOpExpr{Expr: yyDollar[2].expr}
			yyVAL.expr.SetLine(yyDollar[2].expr.Line())
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:457
		{
			yyVAL.expr = &ast.UnaryMinusOpExpr{Expr: yyDollar[2].expr}
			yyVAL.expr.SetLine(yyDollar[2].expr.Line())
		}
	case 77:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:461
		{
			yyVAL.expr = &ast.UnaryNotOpExpr{Expr: yyDollar[2].expr}
			yyVAL.expr.SetLine(yyDollar[2].expr.Line())
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:465
		{
			yyVAL.expr = &ast.UnaryLenOpExpr{Expr: yyDollar[2].expr}
			yyVAL.expr.SetLine(yyDollar[2].expr.Line())
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:471
		{
			yyVAL.expr = &ast.StringExpr{Value: yyDollar[1].token.Str}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:477
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:480
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:483
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:486
		{
			yyVAL.expr = yyDollar[2].expr
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:492
		{
			yyDollar[2].expr.(*ast.FuncCallExpr).AdjustRet = true
			yyVAL.expr = yyDollar[2].expr
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:498
		{
			yyVAL.expr = &ast.FuncCallExpr{Func: yyDollar[1].expr, Args: yyDollar[2].exprlist}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 86:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:502
		{
			yyVAL.expr = &ast.FuncCallExpr{Method: yyDollar[3].token.Str, Receiver: yyDollar[1].expr, Args: yyDollar[4].exprlist}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:508
		{
			ident := &ast.IdentExpr{Value: yyDollar[1].token.Str}
			ident.SetLine(yyDollar[1].token.Pos.Line)

			yyVAL.expr = &ast.FuncCallExpr{Func: ident, Args: yyDollar[2].exprlist}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:517
		{
			yyVAL.exprlist = []ast.Expr{yyDollar[1].expr}
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:520
		{
			yyVAL.exprlist = append(yyDollar[1].exprlist, yyDollar[2].expr)
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:525
		{
			if yylex.(*Lexer).PNewLine {
				yylex.(*Lexer).TokenError(yyDollar[1].token, "ambiguous syntax (function call x new statement)")
			}
			yyVAL.exprlist = []ast.Expr{}
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:531
		{
			if yylex.(*Lexer).PNewLine {
				yylex.(*Lexer).TokenError(yyDollar[1].token, "ambiguous syntax (function call x new statement)")
			}
			yyVAL.exprlist = yyDollar[2].exprlist
		}
	case 92:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:537
		{
			yyVAL.exprlist = []ast.Expr{yyDollar[1].expr}
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:540
		{
			yyVAL.exprlist = []ast.Expr{yyDollar[1].expr}
		}
	case 94:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:545
		{
			yyVAL.expr = &ast.FunctionExpr{ParList: yyDollar[2].funcexpr.ParList, Stmts: yyDollar[2].funcexpr.Stmts}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.expr.SetLastLine(yyDollar[2].funcexpr.LastLine())
		}
	case 95:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:552
		{
			yyVAL.funcexpr = &ast.FunctionExpr{ParList: yyDollar[2].parlist, Stmts: yyDollar[4].stmts}
			yyVAL.funcexpr.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.funcexpr.SetLastLine(yyDollar[5].token.Pos.Line)
		}
	case 96:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:557
		{
			yyVAL.funcexpr = &ast.FunctionExpr{ParList: &ast.ParList{HasVargs: false, Names: []string{}}, Stmts: yyDollar[3].stmts}
			yyVAL.funcexpr.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.funcexpr.SetLastLine(yyDollar[4].token.Pos.Line)
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:564
		{
			yyVAL.parlist = &ast.ParList{HasVargs: true, Names: []string{}}
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:567
		{
			yyVAL.parlist = &ast.ParList{HasVargs: false, Names: []string{}}
			yyVAL.parlist.Names = append(yyVAL.parlist.Names, yyDollar[1].namelist...)
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:571
		{
			yyVAL.parlist = &ast.ParList{HasVargs: true, Names: []string{}}
			yyVAL.parlist.Names = append(yyVAL.parlist.Names, yyDollar[1].namelist...)
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:578
		{
			yyVAL.expr = &ast.TableExpr{Fields: []*ast.Field{}}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:582
		{
			yyVAL.expr = &ast.TableExpr{Fields: yyDollar[2].fieldlist}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:589
		{
			yyVAL.fieldlist = []*ast.Field{yyDollar[1].field}
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:592
		{
			yyVAL.fieldlist = append(yyDollar[1].fieldlist, yyDollar[3].field)
		}
	case 104:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:595
		{
			yyVAL.fieldlist = yyDollar[1].fieldlist
		}
	case 105:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:600
		{
			yyVAL.field = &ast.Field{Key: &ast.StringExpr{Value: yyDollar[1].token.Str}, Value: yyDollar[3].expr}
			yyVAL.field.Key.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 106:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:604
		{
			yyVAL.field = &ast.Field{Key: yyDollar[2].expr, Value: yyDollar[5].expr}
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:607
		{
			yyVAL.field = &ast.Field{Value: yyDollar[1].expr}
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:612
		{
			yyVAL.fieldsep = ","
		}
	case 109:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:615
		{
			yyVAL.fieldsep = ";"
		}
	}
	goto yystack /* stack new state and value */
}
