package lua

import (
	"math"
	"math/rand"
)

func OpenMath(L *LState) int {
	mod := L.RegisterModule(MathLibName, mathFuncs).(*LTable)
	mod.RawSetString("pi", LFloat(math.Pi))
	mod.RawSetString("huge", LFloat(math.MaxFloat64))
	L.Push(mod)
	return 1
}

var mathFuncs = map[string]LGFunction{
	"abs":        mathAbs,
	"acos":       mathAcos,
	"asin":       mathAsin,
	"atan":       mathAtan,
	"atan2":      mathAtan2,
	"ceil":       mathCeil,
	"cos":        mathCos,
	"cosh":       mathCosh,
	"deg":        mathDeg,
	"exp":        mathExp,
	"floor":      mathFloor,
	"fmod":       mathFmod,
	"frexp":      mathFrexp,
	"ldexp":      mathLdexp,
	"log":        mathLog,
	"log10":      mathLog10,
	"max":        mathMax,
	"min":        mathMin,
	"mod":        mathMod,
	"modf":       mathModf,
	"pow":        mathPow,
	"rad":        mathRad,
	"random":     mathRandom,
	"randomseed": mathRandomseed,
	"sin":        mathSin,
	"sinh":       mathSinh,
	"sqrt":       mathSqrt,
	"tan":        mathTan,
	"tanh":       mathTanh,
}

func mathAbs(L *LState) int {
	L.Push(LFloat(math.Abs(L.CheckFloat(1))))
	return 1
}

func mathAcos(L *LState) int {
	L.Push(LFloat(math.Acos(L.CheckFloat(1))))
	return 1
}

func mathAsin(L *LState) int {
	L.Push(LFloat(math.Asin(L.CheckFloat(1))))
	return 1
}

func mathAtan(L *LState) int {
	L.Push(LFloat(math.Atan(L.CheckFloat(1))))
	return 1
}

func mathAtan2(L *LState) int {
	L.Push(LFloat(math.Atan2(L.CheckFloat(1), L.CheckFloat(2))))
	return 1
}

func mathCeil(L *LState) int {
	L.Push(LFloat(math.Ceil(L.CheckFloat(1))))
	return 1
}

func mathCos(L *LState) int {
	L.Push(LFloat(math.Cos(L.CheckFloat(1))))
	return 1
}

func mathCosh(L *LState) int {
	L.Push(LFloat(math.Cosh(L.CheckFloat(1))))
	return 1
}

func mathDeg(L *LState) int {
	L.Push(LFloat(L.CheckFloat(1) * 180 / math.Pi))
	return 1
}

func mathExp(L *LState) int {
	L.Push(LFloat(math.Exp(L.CheckFloat(1))))
	return 1
}

func mathFloor(L *LState) int {
	L.Push(LFloat(math.Floor(L.CheckFloat(1))))
	return 1
}

func mathFmod(L *LState) int {
	L.Push(LFloat(math.Mod(L.CheckFloat(1), L.CheckFloat(2))))
	return 1
}

func mathFrexp(L *LState) int {
	v1, v2 := math.Frexp(L.CheckFloat(1))
	L.Push(LFloat(v1))
	L.Push(LFloat(v2))
	return 2
}

func mathLdexp(L *LState) int {
	L.Push(LFloat(math.Ldexp(L.CheckFloat(1), L.CheckInt(2))))
	return 1
}

func mathLog(L *LState) int {
	L.Push(LFloat(math.Log(L.CheckFloat(1))))
	return 1
}

func mathLog10(L *LState) int {
	L.Push(LFloat(math.Log10(L.CheckFloat(1))))
	return 1
}

func mathMax(L *LState) int {
	if L.GetTop() == 0 {
		L.RaiseError("wrong number of arguments")
	}
	max := L.CheckFloat(1)
	top := L.GetTop()
	for i := 2; i <= top; i++ {
		v := L.CheckFloat(i)
		if v > max {
			max = v
		}
	}
	L.Push(LFloat(max))
	return 1
}

func mathMin(L *LState) int {
	if L.GetTop() == 0 {
		L.RaiseError("wrong number of arguments")
	}
	min := L.CheckFloat(1)
	top := L.GetTop()
	for i := 2; i <= top; i++ {
		v := L.CheckFloat(i)
		if v < min {
			min = v
		}
	}
	L.Push(LFloat(min))
	return 1
}

func mathMod(L *LState) int {
	lhs := L.CheckFloat(1)
	rhs := L.CheckFloat(2)
	L.Push(luaModulo(LFloat(lhs), LFloat(rhs)))
	return 1
}

func mathModf(L *LState) int {
	v1, v2 := math.Modf(L.CheckFloat(1))
	L.Push(LFloat(v1))
	L.Push(LFloat(v2))
	return 2
}

func mathPow(L *LState) int {
	L.Push(LFloat(math.Pow(L.CheckFloat(1), L.CheckFloat(2))))
	return 1
}

func mathRad(L *LState) int {
	L.Push(LFloat(L.CheckFloat(1) * math.Pi / 180))
	return 1
}

func mathRandom(L *LState) int {
	switch L.GetTop() {
	case 0:
		L.Push(LFloat(rand.Float64()))
	case 1:
		n := L.CheckInt(1)
		L.Push(LInt(rand.Intn(n-1) + 1))
	default:
		min := L.CheckInt(1)
		max := L.CheckInt(2) + 1
		L.Push(LInt(rand.Intn(max-min) + min))
	}
	return 1
}

func mathRandomseed(L *LState) int {
	rand.Seed(L.CheckInt64(1))
	return 0
}

func mathSin(L *LState) int {
	L.Push(LFloat(math.Sin(L.CheckFloat(1))))
	return 1
}

func mathSinh(L *LState) int {
	L.Push(LFloat(math.Sinh(L.CheckFloat(1))))
	return 1
}

func mathSqrt(L *LState) int {
	L.Push(LFloat(math.Sqrt(L.CheckFloat(1))))
	return 1
}

func mathTan(L *LState) int {
	L.Push(LFloat(math.Tan(L.CheckFloat(1))))
	return 1
}

func mathTanh(L *LState) int {
	L.Push(LFloat(math.Tanh(L.CheckFloat(1))))
	return 1
}

//
