//line parser.go.y:2
package parse

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"fmt"
	"github.com/lunixbochs/luaish/ast"
)

//line parser.go.y:38
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
const TShr = 57377
const TShl = 57378
const UNARY = 57379
const TPow = 57380

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

//line parser.go.y:581

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
	53, 34,
	54, 34,
	-2, 77,
	-1, 106,
	53, 35,
	54, 35,
	-2, 77,
}

const yyPrivate = 57344

const yyLast = 766

var yyAct = [...]int{

	27, 23, 55, 26, 101, 97, 172, 126, 61, 50,
	181, 156, 57, 155, 59, 58, 60, 64, 67, 36,
	153, 72, 68, 71, 35, 66, 118, 161, 72, 53,
	151, 24, 45, 174, 54, 121, 122, 185, 93, 94,
	95, 96, 43, 44, 52, 104, 102, 124, 119, 108,
	105, 109, 157, 110, 53, 117, 112, 25, 72, 54,
	150, 98, 115, 51, 49, 48, 119, 123, 120, 46,
	47, 127, 206, 128, 129, 130, 131, 132, 133, 134,
	135, 136, 137, 138, 139, 140, 141, 142, 143, 144,
	145, 146, 147, 148, 43, 44, 52, 203, 89, 88,
	67, 90, 91, 92, 158, 154, 65, 152, 84, 85,
	86, 198, 87, 184, 24, 45, 163, 162, 165, 164,
	160, 167, 166, 168, 169, 24, 45, 53, 171, 170,
	53, 197, 54, 89, 88, 54, 90, 91, 92, 20,
	34, 100, 29, 9, 41, 24, 45, 87, 28, 39,
	21, 24, 45, 173, 30, 104, 102, 42, 176, 175,
	18, 69, 32, 24, 45, 31, 43, 44, 21, 191,
	183, 37, 187, 188, 186, 182, 167, 178, 38, 113,
	149, 189, 56, 1, 190, 33, 192, 74, 107, 194,
	193, 40, 111, 70, 17, 19, 8, 201, 200, 63,
	62, 73, 202, 3, 179, 106, 4, 205, 79, 80,
	78, 77, 81, 2, 0, 0, 0, 0, 0, 0,
	89, 88, 0, 90, 91, 92, 75, 76, 82, 83,
	84, 85, 86, 74, 87, 0, 0, 0, 0, 0,
	0, 0, 0, 125, 0, 0, 0, 73, 0, 0,
	0, 0, 0, 0, 79, 80, 78, 77, 81, 0,
	0, 0, 0, 0, 0, 0, 89, 88, 74, 90,
	91, 92, 75, 76, 82, 83, 84, 85, 86, 0,
	87, 0, 73, 0, 0, 0, 0, 177, 0, 79,
	80, 78, 77, 81, 0, 0, 0, 0, 0, 0,
	0, 89, 88, 0, 90, 91, 92, 75, 76, 82,
	83, 84, 85, 86, 29, 87, 41, 0, 0, 0,
	28, 39, 159, 0, 0, 0, 30, 0, 0, 0,
	0, 0, 0, 0, 32, 24, 45, 31, 43, 44,
	21, 0, 0, 37, 0, 74, 0, 195, 0, 0,
	38, 0, 0, 0, 0, 0, 0, 0, 0, 73,
	0, 103, 0, 40, 0, 99, 79, 80, 78, 77,
	81, 0, 0, 0, 0, 0, 0, 74, 89, 88,
	0, 90, 91, 92, 75, 76, 82, 83, 84, 85,
	86, 73, 87, 0, 0, 196, 0, 0, 79, 80,
	78, 77, 81, 0, 0, 0, 0, 0, 0, 0,
	89, 88, 0, 90, 91, 92, 75, 76, 82, 83,
	84, 85, 86, 29, 87, 41, 0, 180, 0, 28,
	39, 0, 0, 0, 0, 30, 0, 0, 0, 0,
	0, 0, 0, 32, 24, 45, 31, 43, 44, 21,
	0, 0, 37, 0, 29, 0, 41, 0, 0, 38,
	28, 39, 0, 0, 0, 0, 30, 0, 0, 0,
	103, 74, 40, 204, 32, 24, 45, 31, 43, 44,
	21, 0, 0, 37, 0, 73, 0, 0, 0, 0,
	38, 0, 79, 80, 78, 77, 81, 0, 0, 0,
	0, 0, 0, 40, 89, 88, 74, 90, 91, 92,
	75, 76, 82, 83, 84, 85, 86, 0, 87, 0,
	73, 0, 0, 199, 0, 0, 0, 79, 80, 78,
	77, 81, 0, 0, 0, 0, 0, 0, 74, 89,
	88, 0, 90, 91, 92, 75, 76, 82, 83, 84,
	85, 86, 73, 87, 0, 116, 0, 0, 0, 79,
	80, 78, 77, 81, 0, 0, 0, 0, 0, 0,
	0, 89, 88, 0, 90, 91, 92, 75, 76, 82,
	83, 84, 85, 86, 74, 87, 114, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 73, 0,
	0, 0, 0, 0, 0, 79, 80, 78, 77, 81,
	0, 0, 0, 0, 0, 0, 74, 89, 88, 0,
	90, 91, 92, 75, 76, 82, 83, 84, 85, 86,
	73, 87, 0, 0, 0, 0, 0, 79, 80, 78,
	77, 81, 0, 74, 0, 0, 0, 0, 0, 89,
	88, 0, 90, 91, 92, 75, 76, 82, 83, 84,
	85, 86, 0, 87, 79, 80, 78, 77, 81, 0,
	0, 0, 0, 0, 0, 0, 89, 88, 0, 90,
	91, 92, 75, 76, 82, 83, 84, 85, 86, 0,
	87, 79, 80, 78, 77, 81, 0, 0, 0, 0,
	0, 0, 0, 89, 88, 0, 90, 91, 92, 75,
	76, 82, 83, 84, 85, 86, 0, 87, 7, 10,
	0, 0, 0, 0, 14, 15, 13, 0, 16, 0,
	0, 0, 6, 12, 0, 0, 0, 11, 0, 81,
	0, 0, 0, 0, 24, 22, 0, 89, 88, 21,
	90, 91, 92, 0, 0, 82, 83, 84, 85, 86,
	0, 87, 0, 0, 0, 5,
}
var yyPact = [...]int{

	-1000, -1000, 713, 5, -1000, -1000, 444, -1000, 16, 8,
	-1000, 444, -1000, 444, 120, 120, 94, -1000, -1000, -1000,
	-1000, 444, 444, -1000, -1000, -1000, -26, 612, -1000, -1000,
	-1000, -1000, -1000, -1000, 8, -1000, -1000, 444, 444, 444,
	444, 25, -1000, -1000, 304, -1000, 444, 114, 444, 120,
	-1000, 120, 132, -1000, -1000, 170, -1000, 580, 39, 534,
	2, 12, 25, -20, -1000, 120, -6, -1000, 183, -53,
	444, 612, 444, 444, 444, 444, 444, 444, 444, 444,
	444, 444, 444, 444, 444, 444, 444, 444, 444, 444,
	444, 444, 444, 96, 96, 96, 96, -1000, 0, -1000,
	-41, -1000, -1, 444, 612, -26, -1000, 8, 264, -1000,
	60, -1000, -33, -1000, -1000, 444, -1000, 444, 444, 120,
	-1000, 120, 120, 25, 444, -1000, -1000, 612, 612, 639,
	666, 710, 710, 710, 710, 710, 710, 710, 61, 61,
	96, 96, 96, 96, -1000, -1000, -1000, -1000, -1000, -54,
	-1000, -1000, -21, -1000, 413, -1000, -1000, 444, 229, -1000,
	-1000, -1000, 168, 612, -1000, 373, 4, -1000, -1000, -1000,
	-1000, -26, -1000, 161, 83, -1000, 612, -16, -1000, 165,
	444, -1000, 160, -1000, -1000, 444, -1000, -1000, 444, 341,
	122, -1000, 612, 102, 502, -1000, 444, -1000, -1000, -1000,
	88, 467, -1000, -1000, -1000, 63, -1000,
}
var yyPgo = [...]int{

	0, 182, 213, 2, 206, 204, 203, 200, 199, 196,
	157, 8, 3, 0, 24, 140, 139, 195, 9, 194,
	193, 185, 5, 180, 19, 141, 4, 105, 1,
}
var yyR1 = [...]int{

	0, 1, 1, 1, 2, 2, 2, 3, 28, 28,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 5, 5, 6, 6, 6,
	7, 7, 8, 8, 9, 9, 10, 10, 10, 11,
	11, 12, 12, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 14, 15, 15, 15,
	15, 17, 16, 16, 19, 20, 20, 18, 18, 18,
	18, 21, 22, 22, 23, 23, 23, 24, 24, 25,
	25, 25, 26, 26, 26, 27, 27,
}
var yyR2 = [...]int{

	0, 1, 2, 3, 0, 2, 2, 1, 1, 1,
	3, 1, 3, 5, 4, 6, 8, 9, 11, 7,
	3, 4, 4, 2, 1, 0, 5, 1, 2, 1,
	1, 3, 1, 3, 1, 3, 1, 4, 3, 1,
	3, 1, 3, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 2, 2, 2, 2, 1, 1, 1, 1,
	3, 3, 2, 4, 2, 1, 2, 2, 3, 1,
	1, 2, 5, 4, 1, 1, 3, 2, 3, 1,
	3, 2, 3, 5, 1, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -2, -6, -4, 52, 19, 5, -9, -15,
	6, 24, 20, 13, 11, 12, 15, -19, -10, -17,
	-16, 36, 32, -28, 31, 52, -12, -13, 16, 10,
	22, 33, 30, -21, -15, -14, -24, 39, 46, 17,
	59, 12, -10, 34, 35, 32, 53, 54, 57, 56,
	-18, 55, 36, -24, -14, -3, -1, -13, -3, -13,
	-28, -11, -7, -8, -28, 12, -11, -28, -13, -16,
	-20, -13, 54, 18, 4, 43, 44, 28, 27, 25,
	26, 29, 45, 46, 47, 48, 49, 51, 38, 37,
	40, 41, 42, -13, -13, -13, -13, -22, 36, 61,
	-25, -26, -28, 57, -13, -12, -10, -15, -13, -28,
	-28, 60, -12, 9, 6, 23, 21, 53, 14, 54,
	-22, 55, 56, -28, 53, 60, 60, -13, -13, -13,
	-13, -13, -13, -13, -13, -13, -13, -13, -13, -13,
	-13, -13, -13, -13, -13, -13, -13, -13, -13, -23,
	60, 30, -11, 61, -27, 54, 52, 53, -13, 58,
	-18, 60, -3, -13, -3, -13, -12, -28, -28, -28,
	-22, -12, 60, -3, 54, -26, -13, 58, 9, -5,
	54, 6, -3, 9, 30, 53, 9, 7, 8, -13,
	-3, 9, -13, -3, -13, 6, 54, 9, 9, 21,
	-3, -13, -3, 9, 6, -3, 9,
}
var yyDef = [...]int{

	4, -2, 1, 2, 5, 6, 27, 29, 0, 11,
	4, 0, 4, 0, 0, 0, 0, 24, -2, 78,
	79, 0, 8, 36, 9, 3, 28, 41, 43, 44,
	45, 46, 47, 48, 49, 50, 51, 0, 0, 0,
	0, 0, 77, 76, 0, 8, 0, 0, 0, 0,
	82, 0, 0, 89, 90, 0, 7, 0, 0, 0,
	39, 0, 0, 30, 32, 0, 23, 39, 0, 79,
	84, 85, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 72, 73, 74, 75, 91, 0, 97,
	0, 99, 36, 0, 104, 10, -2, 0, 0, 38,
	0, 87, 0, 12, 4, 0, 4, 0, 0, 0,
	20, 0, 0, 0, 0, 80, 81, 86, 42, 52,
	53, 54, 55, 56, 57, 58, 59, 60, 61, 62,
	63, 64, 65, 66, 67, 68, 69, 70, 71, 0,
	4, 94, 95, 98, 101, 105, 106, 0, 0, 37,
	83, 88, 0, 14, 25, 0, 0, 40, 31, 33,
	21, 22, 4, 0, 0, 100, 102, 0, 13, 0,
	0, 4, 0, 93, 96, 0, 15, 4, 0, 0,
	0, 92, 103, 0, 0, 4, 0, 19, 16, 4,
	0, 0, 26, 17, 4, 0, 18,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 59, 3, 49, 42, 3,
	36, 60, 47, 45, 54, 46, 56, 48, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 55, 52,
	44, 53, 43, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 57, 3, 58, 40, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 35, 41, 61, 39,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 37, 38, 50, 51,
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
				yyVAL.stmt = &ast.HalfStmt{Expr: yyDollar[1].expr}
				yyVAL.stmt.SetLine(yyDollar[1].expr.Line())
			} else {
				yyVAL.stmt = &ast.FuncCallStmt{Expr: yyDollar[1].expr}
				yyVAL.stmt.SetLine(yyDollar[1].expr.Line())
			}
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:138
		{
			yyVAL.stmt = &ast.DoBlockStmt{Stmts: yyDollar[2].stmts}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[3].token.Pos.Line)
		}
	case 13:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:143
		{
			yyVAL.stmt = &ast.WhileStmt{Condition: yyDollar[2].expr, Stmts: yyDollar[4].stmts}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[5].token.Pos.Line)
		}
	case 14:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:148
		{
			yyVAL.stmt = &ast.RepeatStmt{Condition: yyDollar[4].expr, Stmts: yyDollar[2].stmts}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[4].expr.Line())
		}
	case 15:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:153
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
		//line parser.go.y:163
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
		//line parser.go.y:174
		{
			yyVAL.stmt = &ast.NumberForStmt{Name: yyDollar[2].token.Str, Init: yyDollar[4].expr, Limit: yyDollar[6].expr, Stmts: yyDollar[8].stmts}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[9].token.Pos.Line)
		}
	case 18:
		yyDollar = yyS[yypt-11 : yypt+1]
		//line parser.go.y:179
		{
			yyVAL.stmt = &ast.NumberForStmt{Name: yyDollar[2].token.Str, Init: yyDollar[4].expr, Limit: yyDollar[6].expr, Step: yyDollar[8].expr, Stmts: yyDollar[10].stmts}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[11].token.Pos.Line)
		}
	case 19:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:184
		{
			yyVAL.stmt = &ast.GenericForStmt{Names: yyDollar[2].namelist, Exprs: yyDollar[4].exprlist, Stmts: yyDollar[6].stmts}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[7].token.Pos.Line)
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:189
		{
			yyVAL.stmt = &ast.FuncDefStmt{Name: yyDollar[2].funcname, Func: yyDollar[3].funcexpr}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[3].funcexpr.LastLine())
		}
	case 21:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:194
		{
			yyVAL.stmt = &ast.LocalAssignStmt{Names: []string{yyDollar[3].token.Str}, Exprs: []ast.Expr{yyDollar[4].funcexpr}}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[4].funcexpr.LastLine())
		}
	case 22:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:199
		{
			yyVAL.stmt = &ast.LocalAssignStmt{Names: yyDollar[2].namelist, Exprs: yyDollar[4].exprlist}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:203
		{
			yyVAL.stmt = &ast.LocalAssignStmt{Names: yyDollar[2].namelist, Exprs: []ast.Expr{}}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:207
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
		//line parser.go.y:217
		{
			yyVAL.stmts = []ast.Stmt{}
		}
	case 26:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:220
		{
			yyVAL.stmts = append(yyDollar[1].stmts, &ast.IfStmt{Condition: yyDollar[3].expr, Then: yyDollar[5].stmts})
			yyVAL.stmts[len(yyVAL.stmts)-1].SetLine(yyDollar[2].token.Pos.Line)
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:226
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: nil}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:230
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyDollar[2].exprlist}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:234
		{
			yyVAL.stmt = &ast.BreakStmt{}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:240
		{
			yyVAL.funcname = yyDollar[1].funcname
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:243
		{
			yyVAL.funcname = &ast.FuncName{Func: nil, Receiver: yyDollar[1].funcname.Func, Method: yyDollar[3].token.Str}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:248
		{
			yyVAL.funcname = &ast.FuncName{Func: &ast.IdentExpr{Value: yyDollar[1].token.Str}}
			yyVAL.funcname.Func.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:252
		{
			key := &ast.StringExpr{Value: yyDollar[3].token.Str}
			key.SetLine(yyDollar[3].token.Pos.Line)
			fn := &ast.AttrGetExpr{Object: yyDollar[1].funcname.Func, Key: key}
			fn.SetLine(yyDollar[3].token.Pos.Line)
			yyVAL.funcname = &ast.FuncName{Func: fn}
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:261
		{
			yyVAL.exprlist = []ast.Expr{yyDollar[1].expr}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:264
		{
			yyVAL.exprlist = append(yyDollar[1].exprlist, yyDollar[3].expr)
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:269
		{
			yyVAL.expr = &ast.IdentExpr{Value: yyDollar[1].token.Str}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 37:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:273
		{
			yyVAL.expr = &ast.AttrGetExpr{Object: yyDollar[1].expr, Key: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:277
		{
			key := &ast.StringExpr{Value: yyDollar[3].token.Str}
			key.SetLine(yyDollar[3].token.Pos.Line)
			yyVAL.expr = &ast.AttrGetExpr{Object: yyDollar[1].expr, Key: key}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:285
		{
			yyVAL.namelist = []string{yyDollar[1].token.Str}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:288
		{
			yyVAL.namelist = append(yyDollar[1].namelist, yyDollar[3].token.Str)
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:293
		{
			yyVAL.exprlist = []ast.Expr{yyDollar[1].expr}
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:296
		{
			yyVAL.exprlist = append(yyDollar[1].exprlist, yyDollar[3].expr)
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:301
		{
			yyVAL.expr = &ast.NilExpr{}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:305
		{
			yyVAL.expr = &ast.FalseExpr{}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:309
		{
			yyVAL.expr = &ast.TrueExpr{}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:313
		{
			yyVAL.expr = &ast.NumberExpr{Value: yyDollar[1].token.Str}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:317
		{
			yyVAL.expr = &ast.Comma3Expr{}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:321
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:324
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:327
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:330
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:333
		{
			yyVAL.expr = &ast.LogicalOpExpr{Lhs: yyDollar[1].expr, Operator: "or", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:337
		{
			yyVAL.expr = &ast.LogicalOpExpr{Lhs: yyDollar[1].expr, Operator: "and", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:341
		{
			yyVAL.expr = &ast.RelationalOpExpr{Lhs: yyDollar[1].expr, Operator: ">", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:345
		{
			yyVAL.expr = &ast.RelationalOpExpr{Lhs: yyDollar[1].expr, Operator: "<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:349
		{
			yyVAL.expr = &ast.RelationalOpExpr{Lhs: yyDollar[1].expr, Operator: ">=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:353
		{
			yyVAL.expr = &ast.RelationalOpExpr{Lhs: yyDollar[1].expr, Operator: "<=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:357
		{
			yyVAL.expr = &ast.RelationalOpExpr{Lhs: yyDollar[1].expr, Operator: "==", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:361
		{
			yyVAL.expr = &ast.RelationalOpExpr{Lhs: yyDollar[1].expr, Operator: "!=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:365
		{
			yyVAL.expr = &ast.StringConcatOpExpr{Lhs: yyDollar[1].expr, Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:369
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "+", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:373
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "-", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:377
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "*", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:381
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "/", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:385
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "%", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:389
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "**", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:393
		{
			fmt.Println("TSHL")
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "<<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:398
		{
			fmt.Println("TSHR")
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: ">>", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:403
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "^", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:407
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "|", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:411
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 72:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:415
		{
			yyVAL.expr = &ast.UnaryNegateOpExpr{Expr: yyDollar[2].expr}
			yyVAL.expr.SetLine(yyDollar[2].expr.Line())
		}
	case 73:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:419
		{
			yyVAL.expr = &ast.UnaryMinusOpExpr{Expr: yyDollar[2].expr}
			yyVAL.expr.SetLine(yyDollar[2].expr.Line())
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:423
		{
			yyVAL.expr = &ast.UnaryNotOpExpr{Expr: yyDollar[2].expr}
			yyVAL.expr.SetLine(yyDollar[2].expr.Line())
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:427
		{
			yyVAL.expr = &ast.UnaryLenOpExpr{Expr: yyDollar[2].expr}
			yyVAL.expr.SetLine(yyDollar[2].expr.Line())
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:433
		{
			yyVAL.expr = &ast.StringExpr{Value: yyDollar[1].token.Str}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:439
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:442
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:445
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:448
		{
			yyVAL.expr = yyDollar[2].expr
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:454
		{
			yyDollar[2].expr.(*ast.FuncCallExpr).AdjustRet = true
			yyVAL.expr = yyDollar[2].expr
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:460
		{
			yyVAL.expr = &ast.FuncCallExpr{Func: yyDollar[1].expr, Args: yyDollar[2].exprlist}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 83:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:464
		{
			yyVAL.expr = &ast.FuncCallExpr{Method: yyDollar[3].token.Str, Receiver: yyDollar[1].expr, Args: yyDollar[4].exprlist}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:470
		{
			ident := &ast.IdentExpr{Value: yyDollar[1].token.Str}
			ident.SetLine(yyDollar[1].token.Pos.Line)

			yyVAL.expr = &ast.FuncCallExpr{Func: ident, Args: yyDollar[2].exprlist}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 85:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:479
		{
			yyVAL.exprlist = []ast.Expr{yyDollar[1].expr}
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:482
		{
			yyVAL.exprlist = append(yyDollar[1].exprlist, yyDollar[2].expr)
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:487
		{
			if yylex.(*Lexer).PNewLine {
				yylex.(*Lexer).TokenError(yyDollar[1].token, "ambiguous syntax (function call x new statement)")
			}
			yyVAL.exprlist = []ast.Expr{}
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:493
		{
			if yylex.(*Lexer).PNewLine {
				yylex.(*Lexer).TokenError(yyDollar[1].token, "ambiguous syntax (function call x new statement)")
			}
			yyVAL.exprlist = yyDollar[2].exprlist
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:499
		{
			yyVAL.exprlist = []ast.Expr{yyDollar[1].expr}
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:502
		{
			yyVAL.exprlist = []ast.Expr{yyDollar[1].expr}
		}
	case 91:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:507
		{
			yyVAL.expr = &ast.FunctionExpr{ParList: yyDollar[2].funcexpr.ParList, Stmts: yyDollar[2].funcexpr.Stmts}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.expr.SetLastLine(yyDollar[2].funcexpr.LastLine())
		}
	case 92:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:514
		{
			yyVAL.funcexpr = &ast.FunctionExpr{ParList: yyDollar[2].parlist, Stmts: yyDollar[4].stmts}
			yyVAL.funcexpr.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.funcexpr.SetLastLine(yyDollar[5].token.Pos.Line)
		}
	case 93:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:519
		{
			yyVAL.funcexpr = &ast.FunctionExpr{ParList: &ast.ParList{HasVargs: false, Names: []string{}}, Stmts: yyDollar[3].stmts}
			yyVAL.funcexpr.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.funcexpr.SetLastLine(yyDollar[4].token.Pos.Line)
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:526
		{
			yyVAL.parlist = &ast.ParList{HasVargs: true, Names: []string{}}
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:529
		{
			yyVAL.parlist = &ast.ParList{HasVargs: false, Names: []string{}}
			yyVAL.parlist.Names = append(yyVAL.parlist.Names, yyDollar[1].namelist...)
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:533
		{
			yyVAL.parlist = &ast.ParList{HasVargs: true, Names: []string{}}
			yyVAL.parlist.Names = append(yyVAL.parlist.Names, yyDollar[1].namelist...)
		}
	case 97:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:540
		{
			yyVAL.expr = &ast.TableExpr{Fields: []*ast.Field{}}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:544
		{
			yyVAL.expr = &ast.TableExpr{Fields: yyDollar[2].fieldlist}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:551
		{
			yyVAL.fieldlist = []*ast.Field{yyDollar[1].field}
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:554
		{
			yyVAL.fieldlist = append(yyDollar[1].fieldlist, yyDollar[3].field)
		}
	case 101:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:557
		{
			yyVAL.fieldlist = yyDollar[1].fieldlist
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:562
		{
			yyVAL.field = &ast.Field{Key: &ast.StringExpr{Value: yyDollar[1].token.Str}, Value: yyDollar[3].expr}
			yyVAL.field.Key.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 103:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:566
		{
			yyVAL.field = &ast.Field{Key: yyDollar[2].expr, Value: yyDollar[5].expr}
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:569
		{
			yyVAL.field = &ast.Field{Value: yyDollar[1].expr}
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:574
		{
			yyVAL.fieldsep = ","
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:577
		{
			yyVAL.fieldsep = ";"
		}
	}
	goto yystack /* stack new state and value */
}
