package runtime_utils

import (
	"runtime"
)

func GetCurrentFunctionName() string {
	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}
