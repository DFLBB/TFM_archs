package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	cc_cfg "github.com/chaincode/netcan/afijos/lib/afijos_cfg"
	cc_util "github.com/chaincode/netcan/netcan/lib/netcan_util"

	cc_personas_cfg "github.com/chaincode/netcan/personas/lib/personas_cfg"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type ThisChainCode struct {
}

func main() {

	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))

	err := shim.Start(new(ThisChainCode))
	if err != nil {
		fmt.Printf("Error starting %s: %s", cc_cfg.CFG_ChainCodeName, err)
	}

}

func (tcc *ThisChainCode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))
	return shim.Success(nil)
}

func (tcc *ThisChainCode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))

	function, args := stub.GetFunctionAndParameters()

	fmt.Println(function)
	fmt.Println(args)

	if function == "registrarAfijo" {
		return tcc.registrarAfijo(stub, args)
	} else if function == "registrarCambioPropietario" {
		return tcc.registrarCambioPropietario(stub, args)
	} else if function == "registrarCancelacionAfijo" {
		return tcc.registrarCancelacionAfijo(stub, args)

	} else if function == "asignarEstado" {
		return tcc.asignarEstado(stub, args)
	} else if function == "consultarEstado" {
		return tcc.consultarEstado(stub, args)
	} else if function == "consultarRangoEstados" {
		return tcc.consultarRangoEstados(stub, args)
	} else if function == "borrarEstado" {
		return tcc.borrarEstado(stub, args)
	} else if function == "ejecutarConsulta" {
		return tcc.ejecutarConsulta(stub, args)

	} else if function == "cargarDatosIniciales" {
		return tcc.cargarDatosIniciales(stub, args)
	} else if function == "cargarDatosIniciales_Propietarios" {
		return tcc.cargarDatosIniciales_Propietarios(stub, args)
	} else if function == "consultarDatosAfijo" {
		return tcc.consultarDatosAfijo(stub, args)
	} else if function == "obtenerCertificadoAfijo" {
		return tcc.obtenerCertificadoAfijo(stub, args)

	} else {
		return shim.Error("(" + cc_cfg.CFG_ObjectType + ") Invoca un nombre de funcion no valida (" + function + ")")
	}
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (tcc *ThisChainCode) registrarAfijo(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))

	type TipoRegistrarAfijoPropietario struct {
		IDPersona int `json:"IDPersona"`
	}

	type TipoRegistrarAfijo struct {
		Nombre       string                          `json:"Nombre"`
		Propietarios []TipoRegistrarAfijoPropietario `json:"Propietarios"`
	}

	// ---------------------------------------------------------------------------------------------------
	// VALIDAR Argumentos
	// ---------------------------------------------------------------------------------------------------

	if len(args) != 2 {
		return shim.Error("Incorrecto numero de argumentos. Esperando 2")
	}

	DatosAfijoComoJson := []byte(args[0])

	var datosAfijo TipoRegistrarAfijo
	err := json.Unmarshal(DatosAfijoComoJson, &datosAfijo)
	if err != nil {
		fmt.Println(err)
	}

	NombreAfijo := strings.ToUpper(datosAfijo.Nombre)
	if len(NombreAfijo) <= 0 {
		return shim.Error("(Args[0]) Nombre: debe tener un valor")
	}

	queryString := "{\"selector\":{\"docType\":\"" + cc_cfg.CFG_ObjectType + "\",\"Nombre\":\"" + NombreAfijo + "\"}}"
	// fmt.Println(queryString)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	if string(queryResults) != "[]" {
		return shim.Error("(Args[0]) Nombre: ya existe el Afijo " + string(queryResults))
	}

	// ---------------------------------------------------------------------------------------------------

	Propietarios := datosAfijo.Propietarios
	if len(Propietarios) <= 0 {
		return shim.Error("(Args[0]) Propierarios: debe tener un valor")
	}

	for propietario := range Propietarios {

		IDPersonaPropietario := Propietarios[propietario].IDPersona

		// ---------------------------------------------------------------------------------------------------

		queryString = "{\"selector\":{\"docType\":\"" + cc_personas_cfg.CFG_ObjectType + "\",\"IDPersona\":" + strconv.Itoa(IDPersonaPropietario) + "}}"
		// fmt.Println(queryString)
		response := stub.InvokeChaincode(cc_personas_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("ejecutarConsulta", queryString), stub.GetChannelID())

		if response.Status != shim.OK {
			return shim.Error(response.Message)
		}

		if string(response.Payload) == "[]" {
			return shim.Error("(Args[0, " + strconv.Itoa(propietario) + "]) IDPersona: [ " + strconv.Itoa(IDPersonaPropietario) + " ] no existe o no es valido")
		}

		// ---------------------------------------------------------------------------------------------------

		queryString = "{\"selector\":{\"docType\":\"" + cc_cfg.CFG_ObjectType_Propietarios + "\",\"IDPersona\":" + strconv.Itoa(IDPersonaPropietario) + ",\"FechaBaja\":\"\"}}"
		queryResults, err = getQueryResultForQueryString(stub, queryString)

		if err != nil {
			return shim.Error(err.Error())
		}

		if string(queryResults) != "[]" {
			return shim.Error("(Args[0, " + strconv.Itoa(propietario) + "]) IDPersona: [ " + strconv.Itoa(IDPersonaPropietario) + " ] tiene asignado un afijo activo " + string(queryResults))
		}

		// ---------------------------------------------------------------------------------------------------

	}

	// ---------------------------------------------------------------------------------------------------

	DatosSeguridadComoJson := []byte(args[1])

	var datosSeguridad cc_util.TipoSeguridad

	err = json.Unmarshal(DatosSeguridadComoJson, &datosSeguridad)
	if err != nil {
		fmt.Println(err)
	}

	IDPersonaEjecuta := datosSeguridad.IDPersona

	queryString = "{\"selector\":{\"docType\":\"" + cc_personas_cfg.CFG_ObjectType + "\",\"IDPersona\":" + strconv.Itoa(IDPersonaEjecuta) + "}}"
	fmt.Println(queryString)
	response := stub.InvokeChaincode(cc_personas_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("ejecutarConsulta", queryString), stub.GetChannelID())

	if response.Status != shim.OK {
		return shim.Error(response.Message)
	}

	if string(response.Payload) == "[]" {
		return shim.Error("(Args[1]) IDPersona: [ " + strconv.Itoa(IDPersonaEjecuta) + " ] no existe o no es valido")
	}

	// ---------------------------------------------------------------------------------------------------
	// GRABAR Registros
	// ---------------------------------------------------------------------------------------------------

	var retorno bytes.Buffer

	fechaHoraActual := time.Now()
	fechaHoraActualAsString := fechaHoraActual.String()

	// ---------------------------------------------------------------------------------------------------

	IDAfijo := 0

	infoChainCodeAsBytes, err := stub.GetState(cc_cfg.CFG_ObjectType)

	if err != nil {
		return shim.Error(err.Error())
	}

	if len(infoChainCodeAsBytes) > 0 {

		var queryInfoChaincode cc_util.InfoChaincode
		err = json.Unmarshal(infoChainCodeAsBytes, &queryInfoChaincode)
		if err != nil {
			fmt.Println(err)
		}
		IDAfijo = queryInfoChaincode.IDMaximo
	}

	IDAfijo += 1

	// ---------------------------------------------------------------------------------------------------

	InfoChaincodeAsBytes, err := json.Marshal(&cc_util.InfoChaincode{cc_util.CC_INFO_CONTADOR, IDAfijo})
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println([]string{cc_cfg.CFG_ObjectType, string(InfoChaincodeAsBytes)})

	err = stub.PutState(cc_cfg.CFG_ObjectType, InfoChaincodeAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ---------------------------------------------------------------------------------------------------

	nuevoAfijo := &cc_cfg.Afijos{cc_cfg.CFG_ObjectType, IDAfijo, NombreAfijo, fechaHoraActualAsString, ""}
	fmt.Println(nuevoAfijo)

	nuevoAfijoAsBytes, err := json.Marshal(nuevoAfijo)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(cc_cfg.CFG_ObjectType+strconv.Itoa(IDAfijo), nuevoAfijoAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println(string(nuevoAfijoAsBytes))
	retorno.WriteString(string(nuevoAfijoAsBytes))

	// ---------------------------------------------------------------------------------------------------

	IDAfijoPropietario := 0

	infoChainCodeAsBytes, err = stub.GetState(cc_cfg.CFG_ObjectType)

	if err != nil {
		return shim.Error(err.Error())
	}

	if len(infoChainCodeAsBytes) > 0 {

		var queryInfoChaincode cc_util.InfoChaincode
		err = json.Unmarshal(infoChainCodeAsBytes, &queryInfoChaincode)
		if err != nil {
			fmt.Println(err)
		}
		IDAfijoPropietario = queryInfoChaincode.IDMaximo
	}

	// ---------------------------------------------------------------------------------------------------

	for propietario := range Propietarios {

		IDAfijoPropietario += 1

		nuevoAfijoPropietario := &cc_cfg.AfijosPropietarios{cc_cfg.CFG_ObjectType_Propietarios, IDAfijoPropietario, IDAfijo, Propietarios[propietario].IDPersona, fechaHoraActualAsString, ""}
		fmt.Println(nuevoAfijoPropietario)

		nuevoAfijoPropietarioAsBytes, err := json.Marshal(nuevoAfijoPropietario)
		if err != nil {
			return shim.Error(err.Error())
		}

		err = stub.PutState(cc_cfg.CFG_ObjectType_Propietarios+strconv.Itoa(IDAfijoPropietario), nuevoAfijoPropietarioAsBytes)
		if err != nil {
			return shim.Error(err.Error())
		}

		fmt.Println(string(nuevoAfijoPropietarioAsBytes))
		retorno.WriteString(string(nuevoAfijoPropietarioAsBytes))

	}

	// ---------------------------------------------------------------------------------------------------

	InfoChaincodeAsBytes, err = json.Marshal(&cc_util.InfoChaincode{cc_util.CC_INFO_CONTADOR, IDAfijoPropietario})
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println([]string{cc_cfg.CFG_ObjectType_Propietarios, string(InfoChaincodeAsBytes)})

	err = stub.PutState(cc_cfg.CFG_ObjectType_Propietarios, InfoChaincodeAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ---------------------------------------------------------------------------------------------------

	return shim.Success(retorno.Bytes())
}

func (tcc *ThisChainCode) registrarCambioPropietario(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))

	type TipoRegistrarCambioPropietarioPropietario struct {
		IDPersona int `json:"IDPersona"`
	}

	type TipoRegistrarCambioPropietario struct {
		IDAfijo      int                                         `json:"IDAfijo"`
		Propietarios []TipoRegistrarCambioPropietarioPropietario `json:"Propietarios"`
	}

	type tipoQueryAfijosPropietarios struct {
		Key    string                    `json:"Key"`
		Record cc_cfg.AfijosPropietarios `json:"Record"`
	}

	// ---------------------------------------------------------------------------------------------------
	// VALIDAR Argumentos
	// ---------------------------------------------------------------------------------------------------

	if len(args) != 2 {
		return shim.Error("Incorrecto numero de argumentos. Esperando 2")
	}

	DatosAfijoComoJson := []byte(args[0])

	var datosAfijo TipoRegistrarCambioPropietario
	err := json.Unmarshal(DatosAfijoComoJson, &datosAfijo)
	if err != nil {
		fmt.Println(err)
	}

	IDAfijo := datosAfijo.IDAfijo
	if IDAfijo <= 0 {
		return shim.Error("(Args[0]) Nombre: debe tener un valor")
	}

	queryString := "{\"selector\":{\"docType\":\"" + cc_cfg.CFG_ObjectType + "\",\"IDAfijo\":" + strconv.Itoa(IDAfijo) + ",\"FechaBaja\":\"\"}}"
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	// fmt.Println(string(queryResults))

	if string(queryResults) == "[]" {
		return shim.Error("(Args[0) IDAfijo: [ " + strconv.Itoa(IDAfijo) + " ] no existe o no es valido")
	}

	// ---------------------------------------------------------------------------------------------------

	Propietarios := datosAfijo.Propietarios
	if len(Propietarios) <= 0 {
		return shim.Error("(Args[0]) Propierarios: debe tener un valor")
	}

	for propietario := range Propietarios {

		IDPersonaPropietario := Propietarios[propietario].IDPersona

		// ---------------------------------------------------------------------------------------------------

		queryString = "{\"selector\":{\"docType\":\"" + cc_personas_cfg.CFG_ObjectType + "\",\"IDPersona\":" + strconv.Itoa(IDPersonaPropietario) + "}}"
		// fmt.Println(queryString)
		response := stub.InvokeChaincode(cc_personas_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("ejecutarConsulta", queryString), stub.GetChannelID())

		if response.Status != shim.OK {
			return shim.Error(response.Message)
		}

		if string(response.Payload) == "[]" {
			return shim.Error("(Args[0, " + strconv.Itoa(propietario) + "]) IDPersona: [ " + strconv.Itoa(IDPersonaPropietario) + " ] no existe o no es valido")
		}

		// ---------------------------------------------------------------------------------------------------

		queryString = "{\"selector\":{\"docType\":\"" + cc_cfg.CFG_ObjectType_Propietarios + "\",\"IDPersona\":" + strconv.Itoa(IDPersonaPropietario) + ",\"FechaBaja\":\"\",\"IDAfijo\":{\"$ne\":" + strconv.Itoa(IDAfijo) + "}}}"
		queryResults, err = getQueryResultForQueryString(stub, queryString)

		if err != nil {
			return shim.Error(err.Error())
		}

		if string(queryResults) != "[]" {
			return shim.Error("(Args[0, " + strconv.Itoa(propietario) + "]) IDPersona: [ " + strconv.Itoa(IDPersonaPropietario) + " ] tiene asignado un afijo activo " + string(queryResults))
		}

		// ---------------------------------------------------------------------------------------------------

	}

	// ---------------------------------------------------------------------------------------------------

	DatosSeguridadComoJson := []byte(args[1])

	var datosSeguridad cc_util.TipoSeguridad

	err = json.Unmarshal(DatosSeguridadComoJson, &datosSeguridad)
	if err != nil {
		fmt.Println(err)
	}

	IDPersonaEjecuta := datosSeguridad.IDPersona

	queryString = "{\"selector\":{\"docType\":\"" + cc_personas_cfg.CFG_ObjectType + "\",\"IDPersona\":" + strconv.Itoa(IDPersonaEjecuta) + "}}"
	response := stub.InvokeChaincode(cc_personas_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("ejecutarConsulta", queryString), stub.GetChannelID())

	if response.Status != shim.OK {
		return shim.Error(response.Message)
	}

	if string(response.Payload) == "[]" {
		return shim.Error("(Args[1]) IDPersona: [ " + strconv.Itoa(IDPersonaEjecuta) + " ] no existe o no es valido")
	}

	// ---------------------------------------------------------------------------------------------------
	// GRABAR Registros
	// ---------------------------------------------------------------------------------------------------

	var retorno bytes.Buffer

	fechaHoraActual := time.Now()
	fechaHoraActualAsString := fechaHoraActual.String()

	// ---------------------------------------------------------------------------------------------------

	queryString = "{\"selector\":{\"docType\":\"" + cc_cfg.CFG_ObjectType_Propietarios + "\",\"IDAfijo\":" + strconv.Itoa(IDAfijo) + ",\"FechaBaja\":\"\"}}"
	queryResults, err = getQueryResultForQueryString(stub, queryString)

	if err != nil {
		return shim.Error(err.Error())
	}

	var datosRegistroPropietarios []tipoQueryAfijosPropietarios

	err = json.Unmarshal(queryResults, &datosRegistroPropietarios)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("***BAJAS***")

	for propietario := range datosRegistroPropietarios {

		datosRegistroPropietarios[propietario].Record.FechaBaja = fechaHoraActualAsString

		datosRegistroPropietariosAsBytes, err := json.Marshal(datosRegistroPropietarios[propietario].Record)
		if err != nil {
			return shim.Error(err.Error())
		}

		err = stub.PutState(cc_cfg.CFG_ObjectType_Propietarios+strconv.Itoa(datosRegistroPropietarios[propietario].Record.IDAfijoPropietario), datosRegistroPropietariosAsBytes)
		if err != nil {
			return shim.Error(err.Error())
		}

		fmt.Println(datosRegistroPropietarios[propietario].Record)
		retorno.WriteString(string(datosRegistroPropietariosAsBytes))
	}

	// ---------------------------------------------------------------------------------------------------

	IDAfijoPropietario := 0

	infoChainCodeAsBytes, err := stub.GetState(cc_cfg.CFG_ObjectType_Propietarios)

	if err != nil {
		return shim.Error(err.Error())
	}

	if len(infoChainCodeAsBytes) > 0 {

		var queryInfoChaincode cc_util.InfoChaincode
		err = json.Unmarshal(infoChainCodeAsBytes, &queryInfoChaincode)
		if err != nil {
			fmt.Println(err)
		}
		IDAfijoPropietario = queryInfoChaincode.IDMaximo
	}

	// ---------------------------------------------------------------------------------------------------

	fmt.Println("***ALTAS***")

	for propietario := range Propietarios {

		IDAfijoPropietario += 1

		nuevoAfijoPropietario := &cc_cfg.AfijosPropietarios{cc_cfg.CFG_ObjectType_Propietarios, IDAfijoPropietario, IDAfijo, Propietarios[propietario].IDPersona, fechaHoraActualAsString, ""}
		// fmt.Println(nuevoAfijoPropietario)

		nuevoAfijoPropietarioAsBytes, err := json.Marshal(nuevoAfijoPropietario)
		if err != nil {
			return shim.Error(err.Error())
		}

		err = stub.PutState(cc_cfg.CFG_ObjectType_Propietarios+strconv.Itoa(IDAfijoPropietario), nuevoAfijoPropietarioAsBytes)
		if err != nil {
			return shim.Error(err.Error())
		}

		fmt.Println(string(nuevoAfijoPropietarioAsBytes))
		retorno.WriteString(string(nuevoAfijoPropietarioAsBytes))

	}

	// ---------------------------------------------------------------------------------------------------

	InfoChaincodeAsBytes, err := json.Marshal(&cc_util.InfoChaincode{cc_util.CC_INFO_CONTADOR, IDAfijoPropietario})
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println([]string{cc_cfg.CFG_ObjectType_Propietarios, string(InfoChaincodeAsBytes)})

	err = stub.PutState(cc_cfg.CFG_ObjectType_Propietarios, InfoChaincodeAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ---------------------------------------------------------------------------------------------------

	return shim.Success(retorno.Bytes())
}

func (tcc *ThisChainCode) registrarCancelacionAfijo(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))

	// ---------------------------------------------------------------------------------------------------
	// VALIDAR Argumentos
	// ---------------------------------------------------------------------------------------------------

	if len(args) != 2 {
		return shim.Error("Incorrecto numero de argumentos. Esperando 2")
	}

	DatosAfijoComoJson := []byte(args[0])

	var datosAfijo cc_cfg.TipoRegistrarCancelacionAfijo
	err := json.Unmarshal(DatosAfijoComoJson, &datosAfijo)
	if err != nil {
		fmt.Println(err)
	}

	IDAfijo := datosAfijo.IDAfijo
	if IDAfijo <= 0 {
		return shim.Error("(Args[0]) Nombre: debe tener un valor")
	}

	queryString := "{\"selector\":{\"docType\":\"" + cc_cfg.CFG_ObjectType + "\",\"IDAfijo\":" + strconv.Itoa(IDAfijo) + ",\"FechaBaja\":\"\"}}"
	// fmt.Println(queryString)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	// fmt.Println(string(queryResults))

	if string(queryResults) == "[]" {
		return shim.Error("(Args[0) IDAfijo: [ " + strconv.Itoa(IDAfijo) + " ] no existe o no es valido")
	}

	// ---------------------------------------------------------------------------------------------------

	DatosSeguridadComoJson := []byte(args[1])

	var datosSeguridad cc_util.TipoSeguridad

	err = json.Unmarshal(DatosSeguridadComoJson, &datosSeguridad)
	if err != nil {
		fmt.Println(err)
	}

	IDPersonaEjecuta := datosSeguridad.IDPersona

	queryString = "{\"selector\":{\"docType\":\"" + cc_personas_cfg.CFG_ObjectType + "\",\"IDPersona\":" + strconv.Itoa(IDPersonaEjecuta) + "}}"
	response := stub.InvokeChaincode(cc_personas_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("ejecutarConsulta", queryString), stub.GetChannelID())

	if response.Status != shim.OK {
		return shim.Error(response.Message)
	}

	if string(response.Payload) == "[]" {
		return shim.Error("(Args[1]) IDPersona: [ " + strconv.Itoa(IDPersonaEjecuta) + " ] no existe o no es valido")
	}

	// ---------------------------------------------------------------------------------------------------
	// GRABAR Registros
	// ---------------------------------------------------------------------------------------------------

	var retorno bytes.Buffer

	fechaHoraActual := time.Now()
	fechaHoraActualAsString := fechaHoraActual.String()

	// ---------------------------------------------------------------------------------------------------

	var datosRegistro []cc_cfg.TipoQueryAfijos

	err = json.Unmarshal(queryResults, &datosRegistro)
	if err != nil {
		fmt.Println(err)
	}

	datosRegistro[0].Record.FechaBaja = fechaHoraActualAsString

	datosRegistroAsBytes, err := json.Marshal(datosRegistro[0].Record)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(cc_cfg.CFG_ObjectType+strconv.Itoa(IDAfijo), datosRegistroAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println(datosRegistro[0].Record)
	retorno.WriteString(string(datosRegistroAsBytes))

	// ---------------------------------------------------------------------------------------------------

	queryString = "{\"selector\":{\"docType\":\"" + cc_cfg.CFG_ObjectType_Propietarios + "\",\"IDAfijo\":" + strconv.Itoa(IDAfijo) + ",\"FechaBaja\":\"\"}}"
	queryResults, err = getQueryResultForQueryString(stub, queryString)

	if err != nil {
		return shim.Error(err.Error())
	}

	var datosRegistroPropietarios []cc_cfg.TipoQueryAfijosPropietarios

	err = json.Unmarshal(queryResults, &datosRegistroPropietarios)
	if err != nil {
		fmt.Println(err)
	}

	for propietario := range datosRegistroPropietarios {

		datosRegistroPropietarios[propietario].Record.FechaBaja = fechaHoraActualAsString

		datosRegistroPropietariosAsBytes, err := json.Marshal(datosRegistroPropietarios[propietario].Record)
		if err != nil {
			return shim.Error(err.Error())
		}

		err = stub.PutState(cc_cfg.CFG_ObjectType_Propietarios+strconv.Itoa(datosRegistroPropietarios[propietario].Record.IDAfijoPropietario), datosRegistroPropietariosAsBytes)
		if err != nil {
			return shim.Error(err.Error())
		}

		fmt.Println(datosRegistroPropietarios[propietario].Record)
		retorno.WriteString(string(datosRegistroPropietariosAsBytes))
	}

	return shim.Success(retorno.Bytes())

}

func (tcc *ThisChainCode) consultarDatosAfijo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))

	retorno := " ********** FUNCION " + cc_util.NombreFuncion() + " SIN IMPLEMENTAR **********"
	fmt.Println(retorno)

	return shim.Success([]byte(retorno))
}

func (tcc *ThisChainCode) obtenerCertificadoAfijo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))

	retorno := " ********** FUNCION " + cc_util.NombreFuncion() + " SIN IMPLEMENTAR **********"
	fmt.Println(retorno)

	return shim.Success([]byte(retorno))
}

func (tcc *ThisChainCode) cargarDatosIniciales(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))

	NombreArchivo := args[0]

	nuevosRegistrosAsJson, err := ioutil.ReadFile(NombreArchivo)
	if err != nil {
		return shim.Error(err.Error())
	}

	var nuevosRegistros []cc_cfg.Afijos
	json.Unmarshal(nuevosRegistrosAsJson, &nuevosRegistros)

	for _, nuevoRegistro := range nuevosRegistros {

		fmt.Println(nuevoRegistro)

		nuevoRegistroAsBytes, err := json.Marshal(nuevoRegistro)
		if err != nil {
			return shim.Error(err.Error())
		}

		err = stub.PutState(cc_cfg.CFG_ObjectType+strconv.Itoa(nuevoRegistro.IDAfijo), nuevoRegistroAsBytes)
		if err != nil {
			return shim.Error(err.Error())

		}
	}

	return shim.Success([]byte(strconv.Itoa(len(nuevosRegistros))))
}

func (tcc *ThisChainCode) cargarDatosIniciales_Propietarios(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))

	NombreArchivo := args[0]

	nuevosRegistrosAsJson, err := ioutil.ReadFile(NombreArchivo)
	if err != nil {
		return shim.Error(err.Error())
	}

	var nuevosRegistros []cc_cfg.AfijosPropietarios
	json.Unmarshal(nuevosRegistrosAsJson, &nuevosRegistros)

	for _, nuevoRegistro := range nuevosRegistros {

		fmt.Println(nuevoRegistro)

		nuevoRegistroAsBytes, err := json.Marshal(nuevoRegistro)
		if err != nil {
			return shim.Error(err.Error())
		}

		err = stub.PutState(cc_cfg.CFG_ObjectType_Propietarios+strconv.Itoa(nuevoRegistro.IDAfijoPropietario), nuevoRegistroAsBytes)
		if err != nil {
			return shim.Error(err.Error())

		}
	}

	return shim.Success([]byte(strconv.Itoa(len(nuevosRegistros))))
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////
// FUNCIONES COMUNES
////////////////////////////////////////////////////////////////////////////////////////////////////////////

func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {
	//fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)
	resultsIterator, err := stub.GetQueryResult(queryString)
	defer resultsIterator.Close()
	if err != nil {
		return nil, err
	}
	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse,
			err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")
		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")
	//fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())
	return buffer.Bytes(), nil
}
func (tcc *ThisChainCode) ejecutarConsulta(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))
	fmt.Println(args)

	if len(args) < 1 {
		return shim.Error("Incorrecto numero de argumentos. Esperando 1")
	}

	queryString := args[0]

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (tcc *ThisChainCode) consultarEstado(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))

	if len(args) != 1 {
		return shim.Error("Incorrecto numero de argumentos. Esperando 1")
	}

	Estado := args[0]

	EstadoAsByte, err := stub.GetState(Estado)
	if err != nil {
		return shim.Error(err.Error())
	}

	var buffer bytes.Buffer
	buffer.WriteString("[")

	buffer.WriteString("{\"Key\":")
	buffer.WriteString("\"")
	buffer.WriteString(args[0])
	buffer.WriteString("\"")

	buffer.WriteString(", \"Record\":")
	buffer.WriteString(string(EstadoAsByte))
	buffer.WriteString("}")

	buffer.WriteString("]")

	fmt.Printf("   - consultarEstado:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (tcc *ThisChainCode) consultarRangoEstados(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println("   - SolicitudesChaincode consultarRangoEstados()")

	if len(args) != 3 {
		return shim.Error("Incorrecto numero de argumentos. Esperando 2")
	}

	tipoObjeto := args[0]
	identificadorInicial := tipoObjeto + args[1]
	identificadorFinal := tipoObjeto + args[2]

	resultsIterator, err := stub.GetStateByRange(identificadorInicial, identificadorFinal)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("   - consultarRangoEstados:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (tcc *ThisChainCode) asignarEstado(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	// args[0] --> Tipo Objeto (string)
	// args[1] --> Nuevo valor del objeto (string)

	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))
	fmt.Println(args)

	// ---------------------------------------------------------------------------------------------------

	if len(args) != 2 {
		return shim.Error("Incorrecto numero de argumentos. Esperando 2")
	}

	tipoObjeto := args[0]
	if len(tipoObjeto) <= 0 {
		return shim.Error("(args[0]) tipoObjeto: no tiene valor")
	}

	valorObjeto := args[1]
	if len(valorObjeto) <= 0 {
		return shim.Error("(args[1]) valorObjeto: no tiene valor")
	}

	nuevoValorAsBytes := []byte(args[1])
	err := stub.PutState(tipoObjeto, nuevoValorAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ---------------------------------------------------------------------------------------------------

	return shim.Success(nuevoValorAsBytes)
}

func (tcc *ThisChainCode) borrarEstado(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	// args[0] --> estado a borrar (string)

	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))
	fmt.Println(args)

	// ---------------------------------------------------------------------------------------------------

	if len(args) != 1 {
		return shim.Error("Incorrecto numero de argumentos. Esperando 2")
	}

	if len(args[0]) <= 0 {
		return shim.Error("(args[0]) tipoObjeto: no tiene valor")
	}

	err := stub.DelState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}

	// ---------------------------------------------------------------------------------------------------

	return shim.Success([]byte(args[0]))
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////
