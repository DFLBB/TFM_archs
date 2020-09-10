package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	cc_util "github.com/chaincode/netcan/netcan/lib/netcan_util"
	cc_cfg "github.com/chaincode/netcan/perfiles/lib/perfiles_cfg"

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

	fechaHoraActual := time.Now()
	fechaHoraActualAsString := fechaHoraActual.String()

	nuevoPerfilPersona := &cc_cfg.PerfilesPersonas{cc_cfg.CFG_ObjectType, 0, 0, cc_cfg.CFG_PerfilAdministrador, fechaHoraActualAsString, ""}
	fmt.Println(nuevoPerfilPersona)

	nuevoPerfilPersonaAsBytes, err := json.Marshal(nuevoPerfilPersona)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(cc_cfg.CFG_ObjectType+"0", nuevoPerfilPersonaAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nuevoPerfilPersonaAsBytes)
}

func (tcc *ThisChainCode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))

	function, args := stub.GetFunctionAndParameters()

	fmt.Println(function)
	fmt.Println(args)

	if function == "registrarPerfilPersona" {
		return tcc.registrarPerfilPersona(stub, args)
	} else if function == "cancelarPerfilPersona" {
		return tcc.cancelarPerfilPersona(stub, args)

	} else if function == "esPerfilPersona" {
		return tcc.esPerfilPersona(stub, args[0], args[1])
	} else if function == "esAdministrador" {
		return tcc.esPerfilPersona(stub, cc_cfg.CFG_PerfilAdministrador, args[0])
	} else if function == "esMiembroFederacion" {
		return tcc.esPerfilPersona(stub, cc_cfg.CFG_PerfilFederacion, args[0])
	} else if function == "esVeterinario" {
		return tcc.esPerfilPersona(stub, cc_cfg.CFG_PerfilFederacion, args[0])

	} else if function == "cargarDatosIniciales" {
		return tcc.cargarDatosIniciales(stub, args)

	} else if function == "asignarEstado_OnlyAdmin" {
		return tcc.asignarEstado_OnlyAdmin(stub, args)
	} else if function == "borrarEstado_OnlyAdmin" {
		return tcc.borrarEstado_OnlyAdmin(stub, args)

	} else if function == "consultarEstado" {
		return tcc.consultarEstado(stub, args)
	} else if function == "consultarRangoEstados" {
		return tcc.consultarRangoEstados(stub, args)
	} else if function == "ejecutarConsulta" {
		return tcc.ejecutarConsulta(stub, args)
	} else {
		return shim.Error("(" + cc_cfg.CFG_ObjectType + ") Invalida un nombre de funcion no valida (" + function + ")")
	}

	return shim.Success(nil)
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////

func (tcc *ThisChainCode) registrarPerfilPersona(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))

	// ---------------------------------------------------------------------------------------------------
	// VALIDAR Argumentos
	// ---------------------------------------------------------------------------------------------------

	if len(args) != 2 {
		return shim.Error("Incorrecto numero de argumentos. Esperando 2")
	}

	var queryString string
	var queryResults []byte
	var err error
	var response pb.Response

	// ---------------------------------------------------------------------------------------------------

	DatosSeguridadComoJson := []byte(args[1])

	var datosSeguridad cc_util.TipoSeguridad

	err = json.Unmarshal(DatosSeguridadComoJson, &datosSeguridad)
	if err != nil {
		fmt.Println(err)
	}

	// ---------------------------------------------------------------------------------------------------

	IDPersonaEjecuta := datosSeguridad.IDPersona

	fechaHoraActual := time.Now()
	fechaHoraActualAsString := fechaHoraActual.String()

	queryString = "{\"selector\":{\"docType\":\"" + cc_cfg.CFG_ObjectType + "\",\"CODPerfil\":\"" + cc_cfg.CFG_PerfilAdministrador + "\",\"IDPersona\":" + strconv.Itoa(IDPersonaEjecuta) + ",\"FechaAlta\":{\"$lt\":     \"" + fechaHoraActualAsString + "\"},\"FechaBaja\":\"\"}}"
	queryResults, err = getQueryResultForQueryString(stub, queryString)

	if err != nil {
		return shim.Error(err.Error())
	}

	if string(queryResults) == "[]" {
		return shim.Error("(Args[0]) IDPersonaEjecuta:  [ " + strconv.Itoa(IDPersonaEjecuta) + " ] no tiene un perfil de administrador activo ")
	}

	// ---------------------------------------------------------------------------------------------------

	DatosRegistroComoJson := []byte(args[0])
	var DatosRegistro cc_cfg.PerfilesPersonas
	err = json.Unmarshal(DatosRegistroComoJson, &DatosRegistro)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(DatosRegistro)
	fmt.Println(args[0])

	// ---------------------------------------------------------------------------------------------------

	IDPersona := DatosRegistro.IDPersona

	if IDPersona < 0 {
		return shim.Error("(Args[0]) IDPersona: debe tener un persona")
	}

	queryString = "{\"selector\":{\"docType\":\"" + cc_personas_cfg.CFG_ObjectType + "\",\"IDPersona\":" + strconv.Itoa(IDPersona) + "}}"
	response = stub.InvokeChaincode(cc_personas_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("ejecutarConsulta", queryString), stub.GetChannelID())

	if response.Status != shim.OK {
		return shim.Error(string(response.Payload))
	}

	if string(response.Payload) == "[]" {
		return shim.Error("(Args[0]) IDPersona: [ " + strconv.Itoa(IDPersona) + " ] no existe o no es valido")
	}

	// ---------------------------------------------------------------------------------------------------

	CODPerfil := strings.ToUpper(DatosRegistro.CODPerfil)

	if len(CODPerfil) <= 0 {
		return shim.Error("(Args[0]) CODPerfil: debe contener un valor")
	}

	if CODPerfil != cc_cfg.CFG_PerfilAdministrador && CODPerfil != cc_cfg.CFG_PerfilVeterinario && CODPerfil != cc_cfg.CFG_PerfilFederacion {
		return shim.Error("(Args[0]) CODPerfil: [ " + CODPerfil + " ] no es un perfil valido")
	}

	queryString = "{\"selector\":{\"docType\":\"" + cc_cfg.CFG_ObjectType + "\",\"CODPerfil\":\"" + CODPerfil + "\",\"IDPersona\":" + strconv.Itoa(IDPersona) + ",\"FechaBaja\":\"\"}}"
	queryResults, err = getQueryResultForQueryString(stub, queryString)

	if err != nil {
		return shim.Error(err.Error())
	}

	if string(queryResults) != "[]" {
		return shim.Error("(Args[0]) CODPerfil: [ " + CODPerfil + " ] IDPersona: [ " + strconv.Itoa(IDPersona) + " ] ya esta registrado " + string(queryResults))
	}

	// ---------------------------------------------------------------------------------------------------
	// GRABAR Registros
	// ---------------------------------------------------------------------------------------------------

	IDPerfilPersona := 0

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
		IDPerfilPersona = queryInfoChaincode.IDMaximo
	}

	// ---------------------------------------------------------------------------------------------------

	IDPerfilPersona += 1

	InfoChaincodeAsBytes, err := json.Marshal(&cc_util.InfoChaincode{cc_util.CC_INFO_CONTADOR, IDPerfilPersona})
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println([]string{cc_cfg.CFG_ObjectType, string(InfoChaincodeAsBytes)})

	err = stub.PutState(cc_cfg.CFG_ObjectType, InfoChaincodeAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ---------------------------------------------------------------------------------------------------

	nuevoPerfilPersona := &cc_cfg.PerfilesPersonas{cc_cfg.CFG_ObjectType, IDPerfilPersona, IDPersona, CODPerfil, fechaHoraActualAsString, ""}
	fmt.Println(nuevoPerfilPersona)

	nuevoPerfilPersonaAsBytes, err := json.Marshal(nuevoPerfilPersona)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(cc_cfg.CFG_ObjectType+strconv.Itoa(IDPerfilPersona), nuevoPerfilPersonaAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ---------------------------------------------------------------------------------------------------

	return shim.Success(nuevoPerfilPersonaAsBytes)
}

func (tcc *ThisChainCode) cancelarPerfilPersona(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))

	type tipocancelarPerfilPersona struct {
		CODPerfil string `json:"CODPerfil"`
		IDPersona int    `json:"IDPersona"`
	}

	type tipoQueryPerfilesPersonas struct {
		Key    string                  `json:"Key"`
		Record cc_cfg.PerfilesPersonas `json:"Record"`
	}

	// ---------------------------------------------------------------------------------------------------
	// VALIDAR Argumentos
	// ---------------------------------------------------------------------------------------------------

	if len(args) != 2 {
		return shim.Error("Incorrecto numero de argumentos. Esperando 2")
	}

	var queryString string
	var queryResults []byte
	var err error

	// ---------------------------------------------------------------------------------------------------

	DatosSeguridadComoJson := []byte(args[1])

	var datosSeguridad cc_util.TipoSeguridad

	err = json.Unmarshal(DatosSeguridadComoJson, &datosSeguridad)
	if err != nil {
		fmt.Println(err)
	}

	// ---------------------------------------------------------------------------------------------------

	IDPersonaEjecuta := datosSeguridad.IDPersona

	fechaHoraActual := time.Now()
	fechaHoraActualAsString := fechaHoraActual.String()

	queryString = "{\"selector\":{\"docType\":\"" + cc_cfg.CFG_ObjectType + "\",\"CODPerfil\":\"" + cc_cfg.CFG_PerfilAdministrador + "\",\"IDPersona\":" + strconv.Itoa(IDPersonaEjecuta) + ",\"FechaAlta\":{\"$lt\":     \"" + fechaHoraActualAsString + "\"},\"FechaBaja\":\"\"}}"
	queryResults, err = getQueryResultForQueryString(stub, queryString)

	if err != nil {
		return shim.Error(err.Error())
	}

	if string(queryResults) == "[]" {
		return shim.Error("(Args[0]) IDPersonaEjecuta:  [ " + strconv.Itoa(IDPersonaEjecuta) + " ] no tiene un perfil de administrador activo ")
	}

	// ---------------------------------------------------------------------------------------------------

	DatosPerfilPersonaComoJson := []byte(args[0])

	var DatosPerfilPersona tipocancelarPerfilPersona
	err = json.Unmarshal(DatosPerfilPersonaComoJson, &DatosPerfilPersona)
	if err != nil {
		fmt.Println(err)
	}

	// ---------------------------------------------------------------------------------------------------

	if DatosPerfilPersona.IDPersona < 0 {
		return shim.Error("(Args[0]) IDPersona: debe tener un valor")
	}
	IDPersona := DatosPerfilPersona.IDPersona

	// ---------------------------------------------------------------------------------------------------

	if len(DatosPerfilPersona.CODPerfil) <= 0 {
		return shim.Error("(Args[0]) Nombre: debe tener un valor")
	}
	CODPerfil := strings.ToUpper(DatosPerfilPersona.CODPerfil)

	// ---------------------------------------------------------------------------------------------------

	queryString = "{\"selector\":{\"docType\":\"" + cc_cfg.CFG_ObjectType + "\",\"CODPerfil\":\"" + CODPerfil + "\",\"IDPersona\":" + strconv.Itoa(IDPersona) + ",\"FechaBaja\":\"\"}}"
	queryResults, err = getQueryResultForQueryString(stub, queryString)

	if err != nil {
		return shim.Error(err.Error())
	}

	if string(queryResults) == "[]" {
		return shim.Error("(Args[0]) CODPerfil: [ " + CODPerfil + " ] IDPersona: [ " + strconv.Itoa(IDPersona) + " ] no existe o no es valido ")
	}

	// ---------------------------------------------------------------------------------------------------
	// GRABAR Registros
	// ---------------------------------------------------------------------------------------------------

	var retorno bytes.Buffer

	// ---------------------------------------------------------------------------------------------------

	var datosRegistro []tipoQueryPerfilesPersonas

	err = json.Unmarshal(queryResults, &datosRegistro)
	if err != nil {
		fmt.Println(err)
	}

	datosRegistro[0].Record.FechaBaja = fechaHoraActualAsString

	datosRegistroAsBytes, err := json.Marshal(datosRegistro[0].Record)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(cc_cfg.CFG_ObjectType+strconv.Itoa(datosRegistro[0].Record.IDPerfilPersona), datosRegistroAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println(datosRegistro[0].Record)
	retorno.WriteString(string(datosRegistroAsBytes))

	return shim.Success(retorno.Bytes())
}

func (tcc *ThisChainCode) esPerfilPersona(stub shim.ChaincodeStubInterface, CODPerfil string, stringIDPersona string) pb.Response {

	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))

	type tipoQueryPerfilesPersonas struct {
		Key    string                  `json:"Key"`
		Record cc_cfg.PerfilesPersonas `json:"Record"`
	}

	// ---------------------------------------------------------------------------------------------------
	// VALIDAR Argumentos
	// ---------------------------------------------------------------------------------------------------

	if len(CODPerfil) <= 0 {
		return shim.Error("CODPerfil: debe tener un valor")
	}

	_, err := strconv.Atoi(stringIDPersona)
	if err != nil {
		return shim.Error("IDPersona: [ " + stringIDPersona + " ] debe ser un entero")
	}

	// ---------------------------------------------------------------------------------------------------
	// REALIZAR consulta
	// ---------------------------------------------------------------------------------------------------

	var queryString string
	var queryResults []byte

	queryString = "{\"selector\":{\"docType\":\"" + cc_cfg.CFG_ObjectType + "\",\"CODPerfil\":\"" + CODPerfil + "\",\"IDPersona\":" + stringIDPersona + ",\"FechaBaja\":\"\"}}"
	queryResults, err = getQueryResultForQueryString(stub, queryString)

	if err != nil {
		return shim.Error(err.Error())
	}

	var retorno []byte
	if string(queryResults) == "[]" {
		retorno = []byte("0")
	} else {
		retorno = []byte("1")
	}

	return shim.Success(retorno)
}

func (tcc *ThisChainCode) cargarDatosIniciales(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))

	// ---------------------------------------------------------------------------------------------------
	// VALIDAR Argumentos
	// ---------------------------------------------------------------------------------------------------

	if len(args) != 2 {
		return shim.Error("Incorrecto numero de argumentos. Esperando 2")
	}

	var queryString string
	var queryResults []byte
	var err error

	// ---------------------------------------------------------------------------------------------------

	DatosSeguridadComoJson := []byte(args[1])

	var datosSeguridad cc_util.TipoSeguridad

	err = json.Unmarshal(DatosSeguridadComoJson, &datosSeguridad)
	if err != nil {
		fmt.Println(err)
	}

	// ---------------------------------------------------------------------------------------------------

	IDPersonaEjecuta := datosSeguridad.IDPersona

	fechaHoraActual := time.Now()
	fechaHoraActualAsString := fechaHoraActual.String()

	queryString = "{\"selector\":{\"docType\":\"" + cc_cfg.CFG_ObjectType + "\",\"CODPerfil\":\"" + cc_cfg.CFG_PerfilAdministrador + "\",\"IDPersona\":" + strconv.Itoa(IDPersonaEjecuta) + ",\"FechaAlta\":{\"$lt\":     \"" + fechaHoraActualAsString + "\"},\"FechaBaja\":\"\"}}"
	queryResults, err = getQueryResultForQueryString(stub, queryString)

	if err != nil {
		return shim.Error(err.Error())
	}

	if string(queryResults) == "[]" {
		return shim.Error("(Args[0]) IDPersonaEjecuta:  [ " + strconv.Itoa(IDPersonaEjecuta) + " ] no tiene un perfil de administrador activo ")
	}

	// ---------------------------------------------------------------------------------------------------

	nombreArchivo := args[0]

	nuevosRegistrosAsJson, err := ioutil.ReadFile(nombreArchivo)
	if err != nil {
		return shim.Error(err.Error())
	}

	var nuevosRegistros []cc_cfg.PerfilesPersonas
	json.Unmarshal(nuevosRegistrosAsJson, &nuevosRegistros)

	for _, nuevoRegistro := range nuevosRegistros {

		fmt.Println(nuevoRegistro)

		nuevoRegistroAsBytes, err := json.Marshal(nuevoRegistro)
		if err != nil {
			return shim.Error(err.Error())
		}

		err = stub.PutState(cc_cfg.CFG_ObjectType+strconv.Itoa(nuevoRegistro.IDPerfilPersona), nuevoRegistroAsBytes)
		if err != nil {
			return shim.Error(err.Error())

		}
	}

	return shim.Success([]byte(strconv.Itoa(len(nuevosRegistros))))
}

func (tcc *ThisChainCode) asignarEstado_OnlyAdmin(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))
	fmt.Println(args)

	// ---------------------------------------------------------------------------------------------------
	// VALIDAR Argumentos
	// ---------------------------------------------------------------------------------------------------

	if len(args) != 3 {
		return shim.Error("Incorrecto numero de argumentos. Esperando 3")
	}

	var queryString string
	var queryResults []byte
	var err error

	// ---------------------------------------------------------------------------------------------------

	DatosSeguridadComoJson := []byte(args[1])

	var datosSeguridad cc_util.TipoSeguridad

	err = json.Unmarshal(DatosSeguridadComoJson, &datosSeguridad)
	if err != nil {
		fmt.Println(err)
	}

	// ---------------------------------------------------------------------------------------------------

	IDPersonaEjecuta := datosSeguridad.IDPersona

	fechaHoraActual := time.Now()
	fechaHoraActualAsString := fechaHoraActual.String()

	queryString = "{\"selector\":{\"docType\":\"" + cc_cfg.CFG_ObjectType + "\",\"CODPerfil\":\"" + cc_cfg.CFG_PerfilAdministrador + "\",\"IDPersona\":" + strconv.Itoa(IDPersonaEjecuta) + ",\"FechaAlta\":{\"$lt\":     \"" + fechaHoraActualAsString + "\"},\"FechaBaja\":\"\"}}"
	queryResults, err = getQueryResultForQueryString(stub, queryString)

	if err != nil {
		return shim.Error(err.Error())
	}

	if string(queryResults) == "[]" {
		return shim.Error("(Args[0]) IDPersonaEjecuta:  [ " + strconv.Itoa(IDPersonaEjecuta) + " ] no tiene un perfil de administrador activo ")
	}

	// ---------------------------------------------------------------------------------------------------

	tipoObjeto := args[0]
	if len(tipoObjeto) <= 0 {
		return shim.Error("(args[0]) tipoObjeto: no tiene valor")
	}

	// ---------------------------------------------------------------------------------------------------

	valorObjeto := args[1]
	if len(valorObjeto) <= 0 {
		return shim.Error("(args[1]) valorObjeto: no tiene valor")
	}

	// ---------------------------------------------------------------------------------------------------
	// GRABAR Registros
	// ---------------------------------------------------------------------------------------------------

	nuevoValorAsBytes := []byte(args[1])
	err = stub.PutState(tipoObjeto, nuevoValorAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ---------------------------------------------------------------------------------------------------

	return shim.Success(nuevoValorAsBytes)
}

func (tcc *ThisChainCode) borrarEstado_OnlyAdmin(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))
	fmt.Println(args)

	// ---------------------------------------------------------------------------------------------------
	// VALIDAR Argumentos
	// ---------------------------------------------------------------------------------------------------

	if len(args) != 2 {
		return shim.Error("Incorrecto numero de argumentos. Esperando 2")
	}

	var queryString string
	var queryResults []byte
	var err error

	// ---------------------------------------------------------------------------------------------------

	DatosSeguridadComoJson := []byte(args[1])

	var datosSeguridad cc_util.TipoSeguridad

	err = json.Unmarshal(DatosSeguridadComoJson, &datosSeguridad)
	if err != nil {
		fmt.Println(err)
	}

	// ---------------------------------------------------------------------------------------------------

	IDPersonaEjecuta := datosSeguridad.IDPersona

	fechaHoraActual := time.Now()
	fechaHoraActualAsString := fechaHoraActual.String()

	queryString = "{\"selector\":{\"docType\":\"" + cc_cfg.CFG_ObjectType + "\",\"CODPerfil\":\"" + cc_cfg.CFG_PerfilAdministrador + "\",\"IDPersona\":" + strconv.Itoa(IDPersonaEjecuta) + ",\"FechaAlta\":{\"$lt\":     \"" + fechaHoraActualAsString + "\"},\"FechaBaja\":\"\"}}"
	queryResults, err = getQueryResultForQueryString(stub, queryString)

	if err != nil {
		return shim.Error(err.Error())
	}

	if string(queryResults) == "[]" {
		return shim.Error("(Args[0]) IDPersonaEjecuta:  [ " + strconv.Itoa(IDPersonaEjecuta) + " ] no tiene un perfil de administrador activo ")
	}

	// ---------------------------------------------------------------------------------------------------

	if len(args[0]) <= 0 {
		return shim.Error("(args[0]) tipoObjeto: no tiene valor")
	}

	// ---------------------------------------------------------------------------------------------------
	// GRABAR Registros
	// ---------------------------------------------------------------------------------------------------

	err = stub.DelState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}

	// ---------------------------------------------------------------------------------------------------

	return shim.Success([]byte(args[0]))
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

////////////////////////////////////////////////////////////////////////////////////////////////////////////
