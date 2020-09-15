package netcan_util

import (
	"runtime"
)

type InfoChaincode struct {
	ObjectType string `json:"docType"`
	IDMaximo   int    `json:"IDMaximo"`
}

type TipoQueryInfoChaincode struct {
	Key    string        `json:"Key"`
	Record InfoChaincode `json:"Record"`
}

const (
	CC_INFO_CONTADOR string = "CONTADOR"
)

type TipoSeguridad struct {
	IDPersona int `json:"IDPersona"`
}

func NombreFuncion() string {
	pc, _, _, _ := runtime.Caller(1)
	nameFull := runtime.FuncForPC(pc).Name()
	// nameEnd := filepath.Ext(nameFull)
	// name := strings.TrimPrefix(nameFull , ".")
	// return name
	return nameFull
}

func ToChaincodeArgs(args ...string) [][]byte {
	bargs := make([][]byte, len(args))
	for i, arg := range args {
		bargs[i] = []byte(arg)
	}
	return bargs
}

/*
func existeEnArrayInt(inputArray []int, busqueda int) bool {
	for _, numero := range inputArray {
		if numero == busqueda {
			return true
		}
	}
	return false
}
*/
