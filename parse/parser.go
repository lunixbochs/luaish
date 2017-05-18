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

//line parser.go.y:605

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
	54, 35,
	55, 35,
	-2, 78,
	-1, 108,
	54, 36,
	55, 36,
	-2, 78,
}

const yyPrivate = 57344

const yyLast = 800

var yyAct = [...]int{

	28, 24, 58, 99, 103, 176, 51, 27, 64, 130,
	122, 57, 159, 60, 158, 62, 61, 63, 67, 70,
	74, 156, 37, 71, 57, 178, 69, 36, 74, 154,
	25, 46, 54, 26, 164, 125, 126, 55, 190, 95,
	96, 97, 98, 44, 45, 53, 106, 104, 186, 160,
	110, 123, 111, 121, 112, 107, 100, 116, 54, 153,
	119, 114, 211, 55, 52, 50, 49, 128, 123, 124,
	127, 47, 48, 208, 116, 131, 132, 133, 134, 135,
	136, 137, 138, 139, 140, 141, 142, 143, 144, 145,
	146, 147, 148, 149, 150, 151, 35, 74, 203, 9,
	91, 90, 70, 92, 93, 94, 161, 25, 46, 155,
	86, 87, 88, 68, 89, 44, 45, 53, 165, 163,
	167, 166, 169, 168, 157, 171, 56, 172, 173, 76,
	170, 174, 54, 25, 46, 54, 175, 55, 43, 202,
	55, 19, 21, 75, 196, 109, 189, 25, 46, 188,
	73, 81, 82, 80, 79, 83, 177, 183, 106, 104,
	182, 180, 179, 91, 90, 72, 92, 93, 94, 77,
	78, 84, 85, 86, 87, 88, 117, 89, 102, 187,
	171, 192, 193, 191, 181, 152, 194, 108, 34, 195,
	18, 197, 76, 20, 199, 198, 91, 90, 8, 92,
	93, 94, 206, 205, 25, 46, 75, 207, 66, 22,
	89, 65, 210, 3, 81, 82, 80, 79, 83, 59,
	1, 184, 4, 2, 0, 0, 91, 90, 76, 92,
	93, 94, 77, 78, 84, 85, 86, 87, 88, 0,
	89, 0, 75, 0, 0, 0, 0, 0, 0, 129,
	81, 82, 80, 79, 83, 0, 0, 0, 0, 0,
	0, 0, 91, 90, 0, 92, 93, 94, 77, 78,
	84, 85, 86, 87, 88, 30, 89, 42, 0, 0,
	0, 29, 40, 162, 0, 0, 0, 31, 0, 0,
	0, 0, 0, 0, 0, 0, 33, 25, 46, 32,
	44, 45, 22, 0, 0, 38, 76, 0, 200, 0,
	0, 0, 39, 0, 0, 0, 0, 0, 0, 0,
	75, 0, 0, 105, 0, 41, 0, 101, 81, 82,
	80, 79, 83, 0, 0, 0, 0, 0, 0, 0,
	91, 90, 0, 92, 93, 94, 77, 78, 84, 85,
	86, 87, 88, 30, 89, 42, 0, 201, 0, 29,
	40, 0, 0, 0, 0, 31, 0, 0, 0, 0,
	0, 76, 0, 0, 33, 25, 46, 32, 44, 45,
	22, 0, 0, 38, 0, 75, 0, 0, 0, 0,
	39, 0, 0, 81, 82, 80, 79, 83, 0, 0,
	0, 0, 0, 41, 113, 91, 90, 0, 92, 93,
	94, 77, 78, 84, 85, 86, 87, 88, 30, 89,
	42, 0, 185, 0, 29, 40, 0, 0, 0, 0,
	31, 0, 0, 0, 0, 0, 0, 0, 0, 33,
	25, 46, 32, 44, 45, 22, 115, 0, 38, 0,
	30, 0, 42, 0, 0, 39, 29, 40, 0, 0,
	0, 0, 31, 0, 0, 0, 105, 0, 41, 0,
	0, 33, 25, 46, 32, 44, 45, 22, 0, 0,
	38, 0, 30, 0, 42, 0, 0, 39, 29, 40,
	0, 0, 0, 0, 31, 0, 0, 0, 0, 76,
	41, 209, 0, 33, 25, 46, 32, 44, 45, 22,
	0, 0, 38, 75, 0, 0, 0, 0, 0, 39,
	0, 81, 82, 80, 79, 83, 0, 0, 0, 0,
	0, 0, 41, 91, 90, 76, 92, 93, 94, 77,
	78, 84, 85, 86, 87, 88, 0, 89, 0, 75,
	0, 0, 204, 0, 0, 0, 0, 81, 82, 80,
	79, 83, 0, 0, 0, 0, 0, 0, 76, 91,
	90, 0, 92, 93, 94, 77, 78, 84, 85, 86,
	87, 88, 75, 89, 0, 120, 0, 0, 0, 0,
	81, 82, 80, 79, 83, 0, 0, 0, 0, 0,
	0, 0, 91, 90, 0, 92, 93, 94, 77, 78,
	84, 85, 86, 87, 88, 76, 89, 118, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 75,
	0, 0, 0, 0, 0, 0, 0, 81, 82, 80,
	79, 83, 0, 0, 0, 0, 0, 0, 76, 91,
	90, 0, 92, 93, 94, 77, 78, 84, 85, 86,
	87, 88, 75, 89, 0, 0, 0, 0, 0, 0,
	81, 82, 80, 79, 83, 76, 0, 0, 0, 0,
	0, 0, 91, 90, 0, 92, 93, 94, 77, 78,
	84, 85, 86, 87, 88, 0, 89, 81, 82, 80,
	79, 83, 0, 0, 0, 0, 0, 0, 0, 91,
	90, 0, 92, 93, 94, 77, 78, 84, 85, 86,
	87, 88, 0, 89, 81, 82, 80, 79, 83, 0,
	0, 0, 0, 0, 0, 0, 91, 90, 0, 92,
	93, 94, 77, 78, 84, 85, 86, 87, 88, 0,
	89, 7, 11, 0, 0, 0, 0, 15, 16, 14,
	0, 17, 0, 0, 0, 6, 13, 0, 0, 0,
	12, 10, 0, 83, 0, 0, 0, 0, 25, 23,
	0, 91, 90, 22, 92, 93, 94, 0, 0, 84,
	85, 86, 87, 88, 0, 89, 0, 0, 0, 5,
}
var yyPact = [...]int{

	-1000, -1000, 746, -20, -1000, -1000, 472, -1000, 17, 8,
	472, -1000, 472, -1000, 472, 75, 75, 101, -1000, -1000,
	-1000, -1000, 472, 472, -1000, -1000, -1000, -35, 644, -1000,
	-1000, -1000, -1000, -1000, -1000, 8, -1000, -1000, 472, 472,
	472, 472, 19, -1000, -1000, 265, -1000, 472, 172, 472,
	75, -1000, 75, 343, -1000, -1000, 440, 644, 167, -1000,
	611, 37, 564, -1, -4, 19, -21, -1000, 75, 13,
	-1000, 188, -52, 472, 472, 472, 472, 472, 472, 472,
	472, 472, 472, 472, 472, 472, 472, 472, 472, 472,
	472, 472, 472, 472, 472, 158, 158, 158, 158, -1000,
	-2, -1000, -41, -1000, -5, 472, 644, -35, -1000, 8,
	224, -1000, 80, -1000, -27, -1000, 644, -1000, -1000, 472,
	-1000, 472, 472, 75, -1000, 75, 75, 19, 472, -1000,
	-1000, 644, 671, 698, 743, 743, 743, 743, 743, 743,
	743, 62, 62, 158, 158, 158, 158, -1000, -1000, -1000,
	-1000, -1000, -56, -1000, -1000, -30, -1000, 408, -1000, -1000,
	472, 125, -1000, -1000, -1000, 151, 148, 644, -1000, 367,
	42, -1000, -1000, -1000, -1000, -35, -1000, 140, 115, -1000,
	644, -16, -1000, -1000, 174, 472, -1000, 135, -1000, -1000,
	472, -1000, -1000, 472, 302, 130, -1000, 644, 89, 531,
	-1000, 472, -1000, -1000, -1000, 64, 495, -1000, -1000, -1000,
	53, -1000,
}
var yyPgo = [...]int{

	0, 219, 223, 2, 222, 221, 213, 211, 208, 198,
	138, 8, 7, 0, 27, 96, 142, 193, 6, 190,
	126, 188, 3, 185, 22, 178, 4, 124, 1,
}
var yyR1 = [...]int{

	0, 1, 1, 1, 2, 2, 2, 3, 28, 28,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 5, 5, 6, 6,
	6, 7, 7, 8, 8, 9, 9, 10, 10, 10,
	11, 11, 12, 12, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 14, 15, 15,
	15, 15, 17, 16, 16, 19, 20, 20, 18, 18,
	18, 18, 21, 22, 22, 23, 23, 23, 24, 24,
	25, 25, 25, 26, 26, 26, 27, 27,
}
var yyR2 = [...]int{

	0, 1, 2, 3, 0, 2, 2, 1, 1, 1,
	3, 1, 5, 3, 5, 4, 6, 8, 9, 11,
	7, 3, 4, 4, 2, 1, 0, 5, 1, 2,
	1, 1, 3, 1, 3, 1, 3, 1, 4, 3,
	1, 3, 1, 3, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 2, 2, 2, 2, 1, 1, 1,
	1, 3, 3, 2, 4, 2, 1, 2, 2, 3,
	1, 1, 2, 5, 4, 1, 1, 3, 2, 3,
	1, 3, 2, 3, 5, 1, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -2, -6, -4, 53, 19, 5, -9, -15,
	25, 6, 24, 20, 13, 11, 12, 15, -19, -10,
	-17, -16, 37, 33, -28, 32, 53, -12, -13, 16,
	10, 22, 34, 31, -21, -15, -14, -24, 40, 47,
	17, 60, 12, -10, 35, 36, 33, 54, 55, 58,
	57, -18, 56, 37, -24, -14, -20, -13, -3, -1,
	-13, -3, -13, -28, -11, -7, -8, -28, 12, -11,
	-28, -13, -16, -20, 55, 18, 4, 44, 45, 29,
	28, 26, 27, 30, 46, 47, 48, 49, 50, 52,
	39, 38, 41, 42, 43, -13, -13, -13, -13, -22,
	37, 62, -25, -26, -28, 58, -13, -12, -10, -15,
	-13, -28, -28, 61, -12, 6, -13, 9, 6, 23,
	21, 54, 14, 55, -22, 56, 57, -28, 54, 61,
	61, -13, -13, -13, -13, -13, -13, -13, -13, -13,
	-13, -13, -13, -13, -13, -13, -13, -13, -13, -13,
	-13, -13, -23, 61, 31, -11, 62, -27, 55, 53,
	54, -13, 59, -18, 61, -3, -3, -13, -3, -13,
	-12, -28, -28, -28, -22, -12, 61, -3, 55, -26,
	-13, 59, 9, 9, -5, 55, 6, -3, 9, 31,
	54, 9, 7, 8, -13, -3, 9, -13, -3, -13,
	6, 55, 9, 9, 21, -3, -13, -3, 9, 6,
	-3, 9,
}
var yyDef = [...]int{

	4, -2, 1, 2, 5, 6, 28, 30, 0, 11,
	0, 4, 0, 4, 0, 0, 0, 0, 25, -2,
	79, 80, 0, 8, 37, 9, 3, 29, 42, 44,
	45, 46, 47, 48, 49, 50, 51, 52, 0, 0,
	0, 0, 0, 78, 77, 0, 8, 0, 0, 0,
	0, 83, 0, 0, 90, 91, 0, 86, 0, 7,
	0, 0, 0, 40, 0, 0, 31, 33, 0, 24,
	40, 0, 80, 85, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 73, 74, 75, 76, 92,
	0, 98, 0, 100, 37, 0, 105, 10, -2, 0,
	0, 39, 0, 88, 0, 4, 87, 13, 4, 0,
	4, 0, 0, 0, 21, 0, 0, 0, 0, 81,
	82, 43, 53, 54, 55, 56, 57, 58, 59, 60,
	61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
	71, 72, 0, 4, 95, 96, 99, 102, 106, 107,
	0, 0, 38, 84, 89, 0, 0, 15, 26, 0,
	0, 41, 32, 34, 22, 23, 4, 0, 0, 101,
	103, 0, 12, 14, 0, 0, 4, 0, 94, 97,
	0, 16, 4, 0, 0, 0, 93, 104, 0, 0,
	4, 0, 20, 17, 4, 0, 0, 27, 18, 4,
	0, 19,
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
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:127
		{
			if _, ok := yyDollar[1].expr.(*ast.FuncCallExpr); !ok {
				// if this is just a simple ident, it can be executed as a function (anywhere)
				// otherwise this expression should be evaluated and printed if in a repl
				// this isn't a return statement, because it needs to support bare function calls in the middle of a block
				yyVAL.stmt = &ast.FuncCallStmt{Expr: &ast.FuncCallExpr{
					Func: &ast.IdentExpr{Value: "_fallback"},
					Args: []ast.Expr{yyDollar[1].expr},
				}}
			} else {
				yyVAL.stmt = &ast.FuncCallStmt{Expr: yyDollar[1].expr}
				yyVAL.stmt.SetLine(yyDollar[1].expr.Line())
			}
		}
	case 12:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:141
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
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:168
		{
			yyVAL.stmt = &ast.DoBlockStmt{Stmts: yyDollar[2].stmts}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[3].token.Pos.Line)
		}
	case 14:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:173
		{
			yyVAL.stmt = &ast.WhileStmt{Condition: yyDollar[2].expr, Stmts: yyDollar[4].stmts}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[5].token.Pos.Line)
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:178
		{
			yyVAL.stmt = &ast.RepeatStmt{Condition: yyDollar[4].expr, Stmts: yyDollar[2].stmts}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[4].expr.Line())
		}
	case 16:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:183
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
	case 17:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:193
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
	case 18:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:204
		{
			yyVAL.stmt = &ast.NumberForStmt{Name: yyDollar[2].token.Str, Init: yyDollar[4].expr, Limit: yyDollar[6].expr, Stmts: yyDollar[8].stmts}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[9].token.Pos.Line)
		}
	case 19:
		yyDollar = yyS[yypt-11 : yypt+1]
		//line parser.go.y:209
		{
			yyVAL.stmt = &ast.NumberForStmt{Name: yyDollar[2].token.Str, Init: yyDollar[4].expr, Limit: yyDollar[6].expr, Step: yyDollar[8].expr, Stmts: yyDollar[10].stmts}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[11].token.Pos.Line)
		}
	case 20:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:214
		{
			yyVAL.stmt = &ast.GenericForStmt{Names: yyDollar[2].namelist, Exprs: yyDollar[4].exprlist, Stmts: yyDollar[6].stmts}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[7].token.Pos.Line)
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:219
		{
			yyVAL.stmt = &ast.FuncDefStmt{Name: yyDollar[2].funcname, Func: yyDollar[3].funcexpr}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[3].funcexpr.LastLine())
		}
	case 22:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:224
		{
			yyVAL.stmt = &ast.LocalAssignStmt{Names: []string{yyDollar[3].token.Str}, Exprs: []ast.Expr{yyDollar[4].funcexpr}}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[4].funcexpr.LastLine())
		}
	case 23:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:229
		{
			yyVAL.stmt = &ast.LocalAssignStmt{Names: yyDollar[2].namelist, Exprs: yyDollar[4].exprlist}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 24:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:233
		{
			yyVAL.stmt = &ast.LocalAssignStmt{Names: yyDollar[2].namelist, Exprs: []ast.Expr{}}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:237
		{
			yyVAL.stmt = &ast.FuncCallStmt{Expr: yyDollar[1].expr}
			yyVAL.stmt.SetLine(yyDollar[1].expr.Line())
		}
	case 26:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:243
		{
			yyVAL.stmts = []ast.Stmt{}
		}
	case 27:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:246
		{
			yyVAL.stmts = append(yyDollar[1].stmts, &ast.IfStmt{Condition: yyDollar[3].expr, Then: yyDollar[5].stmts})
			yyVAL.stmts[len(yyVAL.stmts)-1].SetLine(yyDollar[2].token.Pos.Line)
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:252
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: nil}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:256
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyDollar[2].exprlist}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:260
		{
			yyVAL.stmt = &ast.BreakStmt{}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:266
		{
			yyVAL.funcname = yyDollar[1].funcname
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:269
		{
			yyVAL.funcname = &ast.FuncName{Func: nil, Receiver: yyDollar[1].funcname.Func, Method: yyDollar[3].token.Str}
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:274
		{
			yyVAL.funcname = &ast.FuncName{Func: &ast.IdentExpr{Value: yyDollar[1].token.Str}}
			yyVAL.funcname.Func.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:278
		{
			key := &ast.StringExpr{Value: yyDollar[3].token.Str}
			key.SetLine(yyDollar[3].token.Pos.Line)
			fn := &ast.AttrGetExpr{Object: yyDollar[1].funcname.Func, Key: key}
			fn.SetLine(yyDollar[3].token.Pos.Line)
			yyVAL.funcname = &ast.FuncName{Func: fn}
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:287
		{
			yyVAL.exprlist = []ast.Expr{yyDollar[1].expr}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:290
		{
			yyVAL.exprlist = append(yyDollar[1].exprlist, yyDollar[3].expr)
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:295
		{
			yyVAL.expr = &ast.IdentExpr{Value: yyDollar[1].token.Str}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:299
		{
			yyVAL.expr = &ast.AttrGetExpr{Object: yyDollar[1].expr, Key: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:303
		{
			key := &ast.StringExpr{Value: yyDollar[3].token.Str}
			key.SetLine(yyDollar[3].token.Pos.Line)
			yyVAL.expr = &ast.AttrGetExpr{Object: yyDollar[1].expr, Key: key}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:311
		{
			yyVAL.namelist = []string{yyDollar[1].token.Str}
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:314
		{
			yyVAL.namelist = append(yyDollar[1].namelist, yyDollar[3].token.Str)
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:319
		{
			yyVAL.exprlist = []ast.Expr{yyDollar[1].expr}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:322
		{
			yyVAL.exprlist = append(yyDollar[1].exprlist, yyDollar[3].expr)
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:327
		{
			yyVAL.expr = &ast.NilExpr{}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:331
		{
			yyVAL.expr = &ast.FalseExpr{}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:335
		{
			yyVAL.expr = &ast.TrueExpr{}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:339
		{
			yyVAL.expr = &ast.NumberExpr{Value: yyDollar[1].token.Str}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:343
		{
			yyVAL.expr = &ast.Comma3Expr{}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:347
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:350
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:353
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:356
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:359
		{
			yyVAL.expr = &ast.LogicalOpExpr{Lhs: yyDollar[1].expr, Operator: "or", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:363
		{
			yyVAL.expr = &ast.LogicalOpExpr{Lhs: yyDollar[1].expr, Operator: "and", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:367
		{
			yyVAL.expr = &ast.RelationalOpExpr{Lhs: yyDollar[1].expr, Operator: ">", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:371
		{
			yyVAL.expr = &ast.RelationalOpExpr{Lhs: yyDollar[1].expr, Operator: "<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:375
		{
			yyVAL.expr = &ast.RelationalOpExpr{Lhs: yyDollar[1].expr, Operator: ">=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:379
		{
			yyVAL.expr = &ast.RelationalOpExpr{Lhs: yyDollar[1].expr, Operator: "<=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:383
		{
			yyVAL.expr = &ast.RelationalOpExpr{Lhs: yyDollar[1].expr, Operator: "==", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:387
		{
			yyVAL.expr = &ast.RelationalOpExpr{Lhs: yyDollar[1].expr, Operator: "!=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:391
		{
			yyVAL.expr = &ast.StringConcatOpExpr{Lhs: yyDollar[1].expr, Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:395
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "+", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:399
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "-", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:403
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "*", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:407
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "/", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:411
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "%", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:415
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "**", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:419
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "<<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:423
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: ">>", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:427
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "^", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:431
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "|", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:435
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 73:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:439
		{
			yyVAL.expr = &ast.UnaryNegateOpExpr{Expr: yyDollar[2].expr}
			yyVAL.expr.SetLine(yyDollar[2].expr.Line())
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:443
		{
			yyVAL.expr = &ast.UnaryMinusOpExpr{Expr: yyDollar[2].expr}
			yyVAL.expr.SetLine(yyDollar[2].expr.Line())
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:447
		{
			yyVAL.expr = &ast.UnaryNotOpExpr{Expr: yyDollar[2].expr}
			yyVAL.expr.SetLine(yyDollar[2].expr.Line())
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:451
		{
			yyVAL.expr = &ast.UnaryLenOpExpr{Expr: yyDollar[2].expr}
			yyVAL.expr.SetLine(yyDollar[2].expr.Line())
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:457
		{
			yyVAL.expr = &ast.StringExpr{Value: yyDollar[1].token.Str}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:463
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:466
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:469
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:472
		{
			yyVAL.expr = yyDollar[2].expr
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:478
		{
			yyDollar[2].expr.(*ast.FuncCallExpr).AdjustRet = true
			yyVAL.expr = yyDollar[2].expr
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:484
		{
			yyVAL.expr = &ast.FuncCallExpr{Func: yyDollar[1].expr, Args: yyDollar[2].exprlist}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 84:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:488
		{
			yyVAL.expr = &ast.FuncCallExpr{Method: yyDollar[3].token.Str, Receiver: yyDollar[1].expr, Args: yyDollar[4].exprlist}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:494
		{
			ident := &ast.IdentExpr{Value: yyDollar[1].token.Str}
			ident.SetLine(yyDollar[1].token.Pos.Line)

			yyVAL.expr = &ast.FuncCallExpr{Func: ident, Args: yyDollar[2].exprlist}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:503
		{
			yyVAL.exprlist = []ast.Expr{yyDollar[1].expr}
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:506
		{
			yyVAL.exprlist = append(yyDollar[1].exprlist, yyDollar[2].expr)
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:511
		{
			if yylex.(*Lexer).PNewLine {
				yylex.(*Lexer).TokenError(yyDollar[1].token, "ambiguous syntax (function call x new statement)")
			}
			yyVAL.exprlist = []ast.Expr{}
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:517
		{
			if yylex.(*Lexer).PNewLine {
				yylex.(*Lexer).TokenError(yyDollar[1].token, "ambiguous syntax (function call x new statement)")
			}
			yyVAL.exprlist = yyDollar[2].exprlist
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:523
		{
			yyVAL.exprlist = []ast.Expr{yyDollar[1].expr}
		}
	case 91:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:526
		{
			yyVAL.exprlist = []ast.Expr{yyDollar[1].expr}
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:531
		{
			yyVAL.expr = &ast.FunctionExpr{ParList: yyDollar[2].funcexpr.ParList, Stmts: yyDollar[2].funcexpr.Stmts}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.expr.SetLastLine(yyDollar[2].funcexpr.LastLine())
		}
	case 93:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:538
		{
			yyVAL.funcexpr = &ast.FunctionExpr{ParList: yyDollar[2].parlist, Stmts: yyDollar[4].stmts}
			yyVAL.funcexpr.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.funcexpr.SetLastLine(yyDollar[5].token.Pos.Line)
		}
	case 94:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:543
		{
			yyVAL.funcexpr = &ast.FunctionExpr{ParList: &ast.ParList{HasVargs: false, Names: []string{}}, Stmts: yyDollar[3].stmts}
			yyVAL.funcexpr.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.funcexpr.SetLastLine(yyDollar[4].token.Pos.Line)
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:550
		{
			yyVAL.parlist = &ast.ParList{HasVargs: true, Names: []string{}}
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:553
		{
			yyVAL.parlist = &ast.ParList{HasVargs: false, Names: []string{}}
			yyVAL.parlist.Names = append(yyVAL.parlist.Names, yyDollar[1].namelist...)
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:557
		{
			yyVAL.parlist = &ast.ParList{HasVargs: true, Names: []string{}}
			yyVAL.parlist.Names = append(yyVAL.parlist.Names, yyDollar[1].namelist...)
		}
	case 98:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:564
		{
			yyVAL.expr = &ast.TableExpr{Fields: []*ast.Field{}}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:568
		{
			yyVAL.expr = &ast.TableExpr{Fields: yyDollar[2].fieldlist}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:575
		{
			yyVAL.fieldlist = []*ast.Field{yyDollar[1].field}
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:578
		{
			yyVAL.fieldlist = append(yyDollar[1].fieldlist, yyDollar[3].field)
		}
	case 102:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:581
		{
			yyVAL.fieldlist = yyDollar[1].fieldlist
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:586
		{
			yyVAL.field = &ast.Field{Key: &ast.StringExpr{Value: yyDollar[1].token.Str}, Value: yyDollar[3].expr}
			yyVAL.field.Key.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 104:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:590
		{
			yyVAL.field = &ast.Field{Key: yyDollar[2].expr, Value: yyDollar[5].expr}
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:593
		{
			yyVAL.field = &ast.Field{Value: yyDollar[1].expr}
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:598
		{
			yyVAL.fieldsep = ","
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:601
		{
			yyVAL.fieldsep = ";"
		}
	}
	goto yystack /* stack new state and value */
}
