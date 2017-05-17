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
const TEqeq = 57367
const TNeq = 57368
const TLte = 57369
const TGte = 57370
const T2Comma = 57371
const T3Comma = 57372
const TIdent = 57373
const TIdentSpace = 57374
const TNumber = 57375
const TString = 57376
const UNARY = 57377
const TPow = 57378

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

//line parser.go.y:549

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
	-1, 18,
	47, 34,
	48, 34,
	-2, 71,
	-1, 99,
	47, 35,
	48, 35,
	-2, 71,
}

const yyPrivate = 57344

const yyLast = 615

var yyAct = [...]int{

	27, 23, 54, 26, 94, 90, 160, 119, 60, 49,
	71, 144, 56, 143, 58, 57, 59, 63, 66, 36,
	141, 71, 67, 70, 35, 65, 169, 149, 111, 52,
	114, 115, 117, 112, 53, 45, 46, 162, 87, 88,
	89, 139, 24, 44, 97, 95, 173, 145, 101, 98,
	102, 110, 103, 25, 52, 105, 83, 84, 85, 53,
	86, 34, 112, 86, 9, 138, 116, 113, 71, 91,
	120, 108, 121, 122, 123, 124, 125, 126, 127, 128,
	129, 130, 131, 132, 133, 134, 135, 136, 42, 43,
	51, 41, 194, 66, 18, 24, 44, 146, 191, 64,
	140, 24, 44, 50, 48, 47, 21, 20, 100, 151,
	150, 153, 152, 148, 155, 154, 156, 157, 24, 44,
	52, 159, 158, 52, 186, 53, 185, 179, 53, 68,
	29, 171, 40, 42, 43, 51, 28, 38, 99, 166,
	106, 161, 30, 97, 95, 142, 164, 163, 55, 1,
	32, 24, 44, 31, 42, 43, 21, 172, 24, 44,
	37, 93, 137, 170, 155, 175, 176, 174, 33, 177,
	69, 96, 178, 39, 180, 92, 17, 182, 181, 73,
	19, 8, 62, 61, 3, 189, 188, 167, 4, 2,
	190, 0, 0, 72, 0, 193, 0, 0, 0, 0,
	78, 79, 77, 76, 80, 0, 0, 0, 0, 0,
	0, 0, 74, 75, 81, 82, 83, 84, 85, 29,
	86, 40, 0, 0, 0, 28, 38, 0, 0, 118,
	0, 30, 73, 0, 0, 0, 0, 0, 0, 32,
	24, 44, 31, 42, 43, 21, 72, 0, 0, 37,
	0, 0, 0, 78, 79, 77, 76, 80, 0, 0,
	0, 0, 39, 104, 0, 74, 75, 81, 82, 83,
	84, 85, 29, 86, 40, 0, 0, 0, 28, 38,
	165, 0, 0, 0, 30, 73, 0, 0, 0, 0,
	0, 0, 32, 24, 44, 31, 42, 43, 21, 72,
	0, 0, 37, 0, 0, 0, 78, 79, 77, 76,
	80, 0, 0, 96, 0, 39, 0, 0, 74, 75,
	81, 82, 83, 84, 85, 29, 86, 40, 0, 0,
	0, 28, 38, 147, 0, 0, 0, 30, 73, 0,
	183, 0, 0, 0, 0, 32, 24, 44, 31, 42,
	43, 21, 72, 0, 0, 37, 0, 0, 0, 78,
	79, 77, 76, 80, 73, 0, 0, 0, 39, 0,
	0, 74, 75, 81, 82, 83, 84, 85, 72, 86,
	0, 0, 184, 0, 0, 78, 79, 77, 76, 80,
	73, 0, 192, 0, 0, 0, 0, 74, 75, 81,
	82, 83, 84, 85, 72, 86, 0, 0, 168, 0,
	0, 78, 79, 77, 76, 80, 73, 0, 0, 0,
	0, 0, 0, 74, 75, 81, 82, 83, 84, 85,
	72, 86, 0, 187, 0, 0, 0, 78, 79, 77,
	76, 80, 73, 0, 0, 0, 0, 0, 0, 74,
	75, 81, 82, 83, 84, 85, 72, 86, 0, 109,
	0, 0, 0, 78, 79, 77, 76, 80, 73, 0,
	107, 0, 0, 0, 0, 74, 75, 81, 82, 83,
	84, 85, 72, 86, 0, 0, 0, 0, 0, 78,
	79, 77, 76, 80, 73, 0, 0, 0, 0, 0,
	0, 74, 75, 81, 82, 83, 84, 85, 72, 86,
	0, 0, 0, 0, 0, 78, 79, 77, 76, 80,
	73, 0, 0, 0, 0, 0, 0, 74, 75, 81,
	82, 83, 84, 85, 0, 86, 80, 0, 0, 0,
	0, 78, 79, 77, 76, 80, 81, 82, 83, 84,
	85, 0, 86, 74, 75, 81, 82, 83, 84, 85,
	0, 86, 7, 10, 0, 0, 0, 0, 14, 15,
	13, 0, 16, 0, 0, 0, 6, 12, 0, 0,
	0, 11, 0, 0, 0, 0, 0, 0, 24, 22,
	0, 0, 0, 21, 78, 79, 77, 76, 80, 0,
	0, 0, 0, 5, 0, 0, 74, 75, 81, 82,
	83, 84, 85, 0, 86,
}
var yyPact = [...]int{

	-1000, -1000, 557, 7, -1000, -1000, 315, -1000, -12, 54,
	-1000, 315, -1000, 315, 64, 64, 87, -1000, -1000, -1000,
	-1000, 315, 315, -1000, -1000, -1000, -38, 490, -1000, -1000,
	-1000, -1000, -1000, -1000, 54, -1000, -1000, 315, 315, 315,
	33, -1000, -1000, 120, -1000, 315, 70, 315, 64, -1000,
	64, 209, -1000, -1000, 131, -1000, 464, 48, 438, 4,
	14, 33, -19, -1000, 64, -15, -1000, 175, -47, 315,
	490, 315, 315, 315, 315, 315, 315, 315, 315, 315,
	315, 315, 315, 315, 315, 315, 315, 18, 18, 18,
	-1000, 11, -1000, -35, -1000, 0, 315, 490, -38, -1000,
	54, 281, -1000, 99, -1000, -27, -1000, -1000, 315, -1000,
	315, 315, 64, -1000, 64, 64, 33, 315, -1000, -1000,
	490, 490, 516, 569, 507, 507, 507, 507, 507, 507,
	507, 15, 15, 18, 18, 18, 18, -48, -1000, -1000,
	-11, -1000, 262, -1000, -1000, 315, 228, -1000, -1000, -1000,
	130, 490, -1000, 360, 20, -1000, -1000, -1000, -1000, -38,
	-1000, 122, 127, -1000, 490, -1, -1000, 158, 315, -1000,
	118, -1000, -1000, 315, -1000, -1000, 315, 334, 117, -1000,
	490, 115, 412, -1000, 315, -1000, -1000, -1000, 89, 386,
	-1000, -1000, -1000, 83, -1000,
}
var yyPgo = [...]int{

	0, 148, 189, 2, 188, 187, 184, 183, 182, 181,
	91, 8, 3, 0, 24, 61, 107, 180, 9, 176,
	170, 168, 5, 162, 19, 161, 4, 145, 1,
}
var yyR1 = [...]int{

	0, 1, 1, 1, 2, 2, 2, 3, 28, 28,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 5, 5, 6, 6, 6,
	7, 7, 8, 8, 9, 9, 10, 10, 10, 11,
	11, 12, 12, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	14, 15, 15, 15, 15, 17, 16, 16, 19, 20,
	20, 18, 18, 18, 18, 21, 22, 22, 23, 23,
	23, 24, 24, 25, 25, 25, 26, 26, 26, 27,
	27,
}
var yyR2 = [...]int{

	0, 1, 2, 3, 0, 2, 2, 1, 1, 1,
	3, 1, 3, 5, 4, 6, 8, 9, 11, 7,
	3, 4, 4, 2, 1, 0, 5, 1, 2, 1,
	1, 3, 1, 3, 1, 3, 1, 4, 3, 1,
	3, 1, 3, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 2, 2, 2,
	1, 1, 1, 1, 3, 3, 2, 4, 2, 1,
	2, 2, 3, 1, 1, 2, 5, 4, 1, 1,
	3, 2, 3, 1, 3, 2, 3, 5, 1, 1,
	1,
}
var yyChk = [...]int{

	-1000, -1, -2, -6, -4, 46, 19, 5, -9, -15,
	6, 24, 20, 13, 11, 12, 15, -19, -10, -17,
	-16, 36, 32, -28, 31, 46, -12, -13, 16, 10,
	22, 33, 30, -21, -15, -14, -24, 40, 17, 53,
	12, -10, 34, 35, 32, 47, 48, 51, 50, -18,
	49, 36, -24, -14, -3, -1, -13, -3, -13, -28,
	-11, -7, -8, -28, 12, -11, -28, -13, -16, -20,
	-13, 48, 18, 4, 37, 38, 28, 27, 25, 26,
	29, 39, 40, 41, 42, 43, 45, -13, -13, -13,
	-22, 36, 55, -25, -26, -28, 51, -13, -12, -10,
	-15, -13, -28, -28, 54, -12, 9, 6, 23, 21,
	47, 14, 48, -22, 49, 50, -28, 47, 54, 54,
	-13, -13, -13, -13, -13, -13, -13, -13, -13, -13,
	-13, -13, -13, -13, -13, -13, -13, -23, 54, 30,
	-11, 55, -27, 48, 46, 47, -13, 52, -18, 54,
	-3, -13, -3, -13, -12, -28, -28, -28, -22, -12,
	54, -3, 48, -26, -13, 52, 9, -5, 48, 6,
	-3, 9, 30, 47, 9, 7, 8, -13, -3, 9,
	-13, -3, -13, 6, 48, 9, 9, 21, -3, -13,
	-3, 9, 6, -3, 9,
}
var yyDef = [...]int{

	4, -2, 1, 2, 5, 6, 27, 29, 0, 11,
	4, 0, 4, 0, 0, 0, 0, 24, -2, 72,
	73, 0, 8, 36, 9, 3, 28, 41, 43, 44,
	45, 46, 47, 48, 49, 50, 51, 0, 0, 0,
	0, 71, 70, 0, 8, 0, 0, 0, 0, 76,
	0, 0, 83, 84, 0, 7, 0, 0, 0, 39,
	0, 0, 30, 32, 0, 23, 39, 0, 73, 78,
	79, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 67, 68, 69,
	85, 0, 91, 0, 93, 36, 0, 98, 10, -2,
	0, 0, 38, 0, 81, 0, 12, 4, 0, 4,
	0, 0, 0, 20, 0, 0, 0, 0, 74, 75,
	80, 42, 52, 53, 54, 55, 56, 57, 58, 59,
	60, 61, 62, 63, 64, 65, 66, 0, 4, 88,
	89, 92, 95, 99, 100, 0, 0, 37, 77, 82,
	0, 14, 25, 0, 0, 40, 31, 33, 21, 22,
	4, 0, 0, 94, 96, 0, 13, 0, 0, 4,
	0, 87, 90, 0, 15, 4, 0, 0, 0, 86,
	97, 0, 0, 4, 0, 19, 16, 4, 0, 0,
	26, 17, 4, 0, 18,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 53, 3, 43, 3, 3,
	36, 54, 41, 39, 48, 40, 50, 42, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 49, 46,
	38, 47, 37, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 51, 3, 52, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 35, 3, 55,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 44, 45,
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
		//line parser.go.y:76
		{
			yyVAL.stmts = yyDollar[1].stmts
			if l, ok := yylex.(*Lexer); ok {
				l.Stmts = yyVAL.stmts
			}
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:82
		{
			yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[2].stmt)
			if l, ok := yylex.(*Lexer); ok {
				l.Stmts = yyVAL.stmts
			}
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:88
		{
			yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[2].stmt)
			if l, ok := yylex.(*Lexer); ok {
				l.Stmts = yyVAL.stmts
			}
		}
	case 4:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:96
		{
			yyVAL.stmts = []ast.Stmt{}
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:99
		{
			yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[2].stmt)
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:102
		{
			yyVAL.stmts = yyDollar[1].stmts
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:107
		{
			yyVAL.stmts = yyDollar[1].stmts
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:112
		{
			yyVAL.token = yyDollar[1].token
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:113
		{
			yyVAL.token = yyDollar[1].token
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:116
		{
			yyVAL.stmt = &ast.AssignStmt{Lhs: yyDollar[1].exprlist, Rhs: yyDollar[3].exprlist}
			yyVAL.stmt.SetLine(yyDollar[1].exprlist[0].Line())
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:121
		{
			if _, ok := yyDollar[1].expr.(*ast.FuncCallExpr); !ok {
				// if this is just a simple ident, it can be executed as a function (anywhere)
				// otherwise this expression should be evaluated and printed if in a repl
				yyVAL.stmt = &ast.HalfStmt{Expr: yyDollar[1].expr}
				yyVAL.stmt.SetLine(yyDollar[1].expr.Line())
			} else {
				yyVAL.stmt = &ast.FuncCallStmt{Expr: yyDollar[1].expr}
				yyVAL.stmt.SetLine(yyDollar[1].expr.Line())
			}
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:132
		{
			yyVAL.stmt = &ast.DoBlockStmt{Stmts: yyDollar[2].stmts}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[3].token.Pos.Line)
		}
	case 13:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:137
		{
			yyVAL.stmt = &ast.WhileStmt{Condition: yyDollar[2].expr, Stmts: yyDollar[4].stmts}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[5].token.Pos.Line)
		}
	case 14:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:142
		{
			yyVAL.stmt = &ast.RepeatStmt{Condition: yyDollar[4].expr, Stmts: yyDollar[2].stmts}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[4].expr.Line())
		}
	case 15:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:147
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
	case 16:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:157
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
	case 17:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:168
		{
			yyVAL.stmt = &ast.NumberForStmt{Name: yyDollar[2].token.Str, Init: yyDollar[4].expr, Limit: yyDollar[6].expr, Stmts: yyDollar[8].stmts}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[9].token.Pos.Line)
		}
	case 18:
		yyDollar = yyS[yypt-11 : yypt+1]
		//line parser.go.y:173
		{
			yyVAL.stmt = &ast.NumberForStmt{Name: yyDollar[2].token.Str, Init: yyDollar[4].expr, Limit: yyDollar[6].expr, Step: yyDollar[8].expr, Stmts: yyDollar[10].stmts}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[11].token.Pos.Line)
		}
	case 19:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:178
		{
			yyVAL.stmt = &ast.GenericForStmt{Names: yyDollar[2].namelist, Exprs: yyDollar[4].exprlist, Stmts: yyDollar[6].stmts}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[7].token.Pos.Line)
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:183
		{
			yyVAL.stmt = &ast.FuncDefStmt{Name: yyDollar[2].funcname, Func: yyDollar[3].funcexpr}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[3].funcexpr.LastLine())
		}
	case 21:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:188
		{
			yyVAL.stmt = &ast.LocalAssignStmt{Names: []string{yyDollar[3].token.Str}, Exprs: []ast.Expr{yyDollar[4].funcexpr}}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[4].funcexpr.LastLine())
		}
	case 22:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:193
		{
			yyVAL.stmt = &ast.LocalAssignStmt{Names: yyDollar[2].namelist, Exprs: yyDollar[4].exprlist}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:197
		{
			yyVAL.stmt = &ast.LocalAssignStmt{Names: yyDollar[2].namelist, Exprs: []ast.Expr{}}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:201
		{
			if _, ok := yyDollar[1].expr.(*ast.FuncCallExpr); !ok {
				yylex.(*Lexer).Error("parse error 2")
			} else {
				yyVAL.stmt = &ast.FuncCallStmt{Expr: yyDollar[1].expr}
				yyVAL.stmt.SetLine(yyDollar[1].expr.Line())
			}
		}
	case 25:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:211
		{
			yyVAL.stmts = []ast.Stmt{}
		}
	case 26:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:214
		{
			yyVAL.stmts = append(yyDollar[1].stmts, &ast.IfStmt{Condition: yyDollar[3].expr, Then: yyDollar[5].stmts})
			yyVAL.stmts[len(yyVAL.stmts)-1].SetLine(yyDollar[2].token.Pos.Line)
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:220
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: nil}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:224
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyDollar[2].exprlist}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:228
		{
			yyVAL.stmt = &ast.BreakStmt{}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:234
		{
			yyVAL.funcname = yyDollar[1].funcname
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:237
		{
			yyVAL.funcname = &ast.FuncName{Func: nil, Receiver: yyDollar[1].funcname.Func, Method: yyDollar[3].token.Str}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:242
		{
			yyVAL.funcname = &ast.FuncName{Func: &ast.IdentExpr{Value: yyDollar[1].token.Str}}
			yyVAL.funcname.Func.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:246
		{
			key := &ast.StringExpr{Value: yyDollar[3].token.Str}
			key.SetLine(yyDollar[3].token.Pos.Line)
			fn := &ast.AttrGetExpr{Object: yyDollar[1].funcname.Func, Key: key}
			fn.SetLine(yyDollar[3].token.Pos.Line)
			yyVAL.funcname = &ast.FuncName{Func: fn}
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:255
		{
			yyVAL.exprlist = []ast.Expr{yyDollar[1].expr}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:258
		{
			yyVAL.exprlist = append(yyDollar[1].exprlist, yyDollar[3].expr)
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:263
		{
			yyVAL.expr = &ast.IdentExpr{Value: yyDollar[1].token.Str}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 37:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:267
		{
			yyVAL.expr = &ast.AttrGetExpr{Object: yyDollar[1].expr, Key: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:271
		{
			key := &ast.StringExpr{Value: yyDollar[3].token.Str}
			key.SetLine(yyDollar[3].token.Pos.Line)
			yyVAL.expr = &ast.AttrGetExpr{Object: yyDollar[1].expr, Key: key}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:279
		{
			yyVAL.namelist = []string{yyDollar[1].token.Str}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:282
		{
			yyVAL.namelist = append(yyDollar[1].namelist, yyDollar[3].token.Str)
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:287
		{
			yyVAL.exprlist = []ast.Expr{yyDollar[1].expr}
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:290
		{
			yyVAL.exprlist = append(yyDollar[1].exprlist, yyDollar[3].expr)
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:295
		{
			yyVAL.expr = &ast.NilExpr{}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:299
		{
			yyVAL.expr = &ast.FalseExpr{}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:303
		{
			yyVAL.expr = &ast.TrueExpr{}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:307
		{
			yyVAL.expr = &ast.NumberExpr{Value: yyDollar[1].token.Str}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:311
		{
			yyVAL.expr = &ast.Comma3Expr{}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:315
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:318
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:321
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:324
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:327
		{
			yyVAL.expr = &ast.LogicalOpExpr{Lhs: yyDollar[1].expr, Operator: "or", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:331
		{
			yyVAL.expr = &ast.LogicalOpExpr{Lhs: yyDollar[1].expr, Operator: "and", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:335
		{
			yyVAL.expr = &ast.RelationalOpExpr{Lhs: yyDollar[1].expr, Operator: ">", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:339
		{
			yyVAL.expr = &ast.RelationalOpExpr{Lhs: yyDollar[1].expr, Operator: "<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:343
		{
			yyVAL.expr = &ast.RelationalOpExpr{Lhs: yyDollar[1].expr, Operator: ">=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:347
		{
			yyVAL.expr = &ast.RelationalOpExpr{Lhs: yyDollar[1].expr, Operator: "<=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:351
		{
			yyVAL.expr = &ast.RelationalOpExpr{Lhs: yyDollar[1].expr, Operator: "==", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:355
		{
			yyVAL.expr = &ast.RelationalOpExpr{Lhs: yyDollar[1].expr, Operator: "!=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:359
		{
			yyVAL.expr = &ast.StringConcatOpExpr{Lhs: yyDollar[1].expr, Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:363
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "+", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:367
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "-", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:371
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "*", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:375
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "/", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:379
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "%", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:383
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "^", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:387
		{
			yyVAL.expr = &ast.UnaryMinusOpExpr{Expr: yyDollar[2].expr}
			yyVAL.expr.SetLine(yyDollar[2].expr.Line())
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:391
		{
			yyVAL.expr = &ast.UnaryNotOpExpr{Expr: yyDollar[2].expr}
			yyVAL.expr.SetLine(yyDollar[2].expr.Line())
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:395
		{
			yyVAL.expr = &ast.UnaryLenOpExpr{Expr: yyDollar[2].expr}
			yyVAL.expr.SetLine(yyDollar[2].expr.Line())
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:401
		{
			yyVAL.expr = &ast.StringExpr{Value: yyDollar[1].token.Str}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:407
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:410
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:413
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:416
		{
			yyVAL.expr = yyDollar[2].expr
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:422
		{
			yyDollar[2].expr.(*ast.FuncCallExpr).AdjustRet = true
			yyVAL.expr = yyDollar[2].expr
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:428
		{
			yyVAL.expr = &ast.FuncCallExpr{Func: yyDollar[1].expr, Args: yyDollar[2].exprlist}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 77:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:432
		{
			yyVAL.expr = &ast.FuncCallExpr{Method: yyDollar[3].token.Str, Receiver: yyDollar[1].expr, Args: yyDollar[4].exprlist}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:438
		{
			ident := &ast.IdentExpr{Value: yyDollar[1].token.Str}
			ident.SetLine(yyDollar[1].token.Pos.Line)

			yyVAL.expr = &ast.FuncCallExpr{Func: ident, Args: yyDollar[2].exprlist}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:447
		{
			yyVAL.exprlist = []ast.Expr{yyDollar[1].expr}
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:450
		{
			yyVAL.exprlist = append(yyDollar[1].exprlist, yyDollar[2].expr)
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:455
		{
			if yylex.(*Lexer).PNewLine {
				yylex.(*Lexer).TokenError(yyDollar[1].token, "ambiguous syntax (function call x new statement)")
			}
			yyVAL.exprlist = []ast.Expr{}
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:461
		{
			if yylex.(*Lexer).PNewLine {
				yylex.(*Lexer).TokenError(yyDollar[1].token, "ambiguous syntax (function call x new statement)")
			}
			yyVAL.exprlist = yyDollar[2].exprlist
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:467
		{
			yyVAL.exprlist = []ast.Expr{yyDollar[1].expr}
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:470
		{
			yyVAL.exprlist = []ast.Expr{yyDollar[1].expr}
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:475
		{
			yyVAL.expr = &ast.FunctionExpr{ParList: yyDollar[2].funcexpr.ParList, Stmts: yyDollar[2].funcexpr.Stmts}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.expr.SetLastLine(yyDollar[2].funcexpr.LastLine())
		}
	case 86:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:482
		{
			yyVAL.funcexpr = &ast.FunctionExpr{ParList: yyDollar[2].parlist, Stmts: yyDollar[4].stmts}
			yyVAL.funcexpr.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.funcexpr.SetLastLine(yyDollar[5].token.Pos.Line)
		}
	case 87:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:487
		{
			yyVAL.funcexpr = &ast.FunctionExpr{ParList: &ast.ParList{HasVargs: false, Names: []string{}}, Stmts: yyDollar[3].stmts}
			yyVAL.funcexpr.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.funcexpr.SetLastLine(yyDollar[4].token.Pos.Line)
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:494
		{
			yyVAL.parlist = &ast.ParList{HasVargs: true, Names: []string{}}
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:497
		{
			yyVAL.parlist = &ast.ParList{HasVargs: false, Names: []string{}}
			yyVAL.parlist.Names = append(yyVAL.parlist.Names, yyDollar[1].namelist...)
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:501
		{
			yyVAL.parlist = &ast.ParList{HasVargs: true, Names: []string{}}
			yyVAL.parlist.Names = append(yyVAL.parlist.Names, yyDollar[1].namelist...)
		}
	case 91:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:508
		{
			yyVAL.expr = &ast.TableExpr{Fields: []*ast.Field{}}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:512
		{
			yyVAL.expr = &ast.TableExpr{Fields: yyDollar[2].fieldlist}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:519
		{
			yyVAL.fieldlist = []*ast.Field{yyDollar[1].field}
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:522
		{
			yyVAL.fieldlist = append(yyDollar[1].fieldlist, yyDollar[3].field)
		}
	case 95:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:525
		{
			yyVAL.fieldlist = yyDollar[1].fieldlist
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:530
		{
			yyVAL.field = &ast.Field{Key: &ast.StringExpr{Value: yyDollar[1].token.Str}, Value: yyDollar[3].expr}
			yyVAL.field.Key.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 97:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:534
		{
			yyVAL.field = &ast.Field{Key: yyDollar[2].expr, Value: yyDollar[5].expr}
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:537
		{
			yyVAL.field = &ast.Field{Value: yyDollar[1].expr}
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:542
		{
			yyVAL.fieldsep = ","
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:545
		{
			yyVAL.fieldsep = ";"
		}
	}
	goto yystack /* stack new state and value */
}
