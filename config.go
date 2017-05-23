package lua

import (
	"os"
)

var CompatVarArg = true
var FieldsPerFlush = 50
var RegistrySize = 256 * 20
var CallStackSize = 256
var MaxTableGetLoop = 100
var MaxArrayIndex = 67108864

type LInt int64

const LIntBit = 64
const LIntScanFormat = "%d"

type LFloat float64

const LFloatBit = 64
const LFloatScanFormat = "%f"

var LuaPath = "LUA_PATH"
var LuaOS string

func init() {
	if os.PathSeparator == '/' { // unix-like
		LuaOS = "unix"
	} else { // windows
		LuaOS = "windows"
	}
}
