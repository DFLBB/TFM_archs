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
	cc_cfg "github.com/chaincode/netcan/vacunas/lib/vacunas_cfg"

	cc_perfiles_cfg "github.com/chaincode/netcan/perfiles/lib/perfiles_cfg"
	cc_perros_cfg "github.com/chaincode/netcan/perros/lib/perros_cfg"
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

	if function == "registrarVacunaPerro" {
		return tcc.registrarVacunaPerro(stub, args)

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
	} else if function == "cargarDatosIniciales_VacunasPerrosProteccion" {
		return tcc.cargarDatosIniciales_VacunasPerrosProteccion(stub, args)
	} else if function == "cargarDatosIniciales_VacunasProteccion" {
		return tcc.cargarDatosIniciales_VacunasProteccion(stub, args)
	} else if function == "obtenerCertificadoVacunaciones " {
		return tcc.obtenerCertificadoVacunaciones(stub, args)
	} else if function == "consultarVacunaciones" {
		return tcc.consultarVacunaciones(stub, args)

	} else {
		return shim.Error("(" + cc_cfg.CFG_ObjectType + ") Invoca un nombre de funcion no valida (" + function + ")")
	}

	return shim.Success(nil)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (tcc *ThisChainCode) cargarDatosIniciales(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))

	NombreArchivo := args[0]

	nuevosRegistrosAsJson, err := ioutil.ReadFile(NombreArchivo)
	if err != nil {
		return shim.Error(err.Error())
	}

	var nuevosRegistros []cc_cfg.VacunasPerros
	json.Unmarshal(nuevosRegistrosAsJson, &nuevosRegistros)

	for _, nuevoRegistro := range nuevosRegistros {

		fmt.Println(nuevoRegistro)

		nuevoRegistroAsBytes, err := json.Marshal(nuevoRegistro)
		if err != nil {
			return shim.Error(err.Error())
		}

		err = stub.PutState(cc_cfg.CFG_ObjectType+strconv.Itoa(nuevoRegistro.IDVacunaPerro), nuevoRegistroAsBytes)
		if err != nil {
			return shim.Error(err.Error())

		}
	}

	return shim.Success([]byte(strconv.Itoa(len(nuevosRegistros))))
}

func (tcc *ThisChainCode) cargarDatosIniciales_VacunasPerrosProteccion(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))

	NombreArchivo := args[0]

	nuevosRegistrosAsJson, err := ioutil.ReadFile(NombreArchivo)
	if err != nil {
		return shim.Error(err.Error())
	}

	var nuevosRegistros []cc_cfg.VacunasPerrosProteccion
	json.Unmarshal(nuevosRegistrosAsJson, &nuevosRegistros)

	for _, nuevoRegistro := range nuevosRegistros {

		fmt.Println(nuevoRegistro)

		nuevoRegistroAsBytes, err := json.Marshal(nuevoRegistro)
		if err != nil {
			return shim.Error(err.Error())
		}

		err = stub.PutState(cc_cfg.CFG_ObjectType_VacunasPerrosProteccion+strconv.Itoa(nuevoRegistro.IDVacunaPerroProteccion), nuevoRegistroAsBytes)
		if err != nil {
			return shim.Error(err.Error())

		}
	}

	return shim.Success([]byte(strconv.Itoa(len(nuevosRegistros))))
}

func (tcc *ThisChainCode) cargarDatosIniciales_VacunasProteccion(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))

	NombreArchivo := args[0]

	nuevosRegistrosAsJson, err := ioutil.ReadFile(NombreArchivo)
	if err != nil {
		return shim.Error(err.Error())
	}

	var nuevosRegistros []cc_cfg.VacunasProteccion
	json.Unmarshal(nuevosRegistrosAsJson, &nuevosRegistros)

	for _, nuevoRegistro := range nuevosRegistros {

		fmt.Println(nuevoRegistro)

		nuevoRegistroAsBytes, err := json.Marshal(nuevoRegistro)
		if err != nil {
			return shim.Error(err.Error())
		}

		err = stub.PutState(cc_cfg.CFG_ObjectType_VacunasProteccion+strconv.Itoa(nuevoRegistro.IDVacunaProteccion), nuevoRegistroAsBytes)
		if err != nil {
			return shim.Error(err.Error())

		}
	}

	return shim.Success([]byte(strconv.Itoa(len(nuevosRegistros))))
}

func (tcc *ThisChainCode) registrarVacunaPerro(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))

	type tipoRegistrarVacunaPerroPropietario struct {
		IDVacunaProteccion int    `json:"IDVacunaProteccion"`
		FechaBaja          string `json:"FechaBaja"`
	}

	type tipoRegistrarVacunaPerro struct {
		IDPerro              int                                   `json:"IDPerro"`
		IDPersonaVeterinario int                                   `json:"IDPersonaVeterinario"`
		CODVacuna            string                                `json:"CODVacuna"`
		FechaAlta            string                                `json:"FechaAlta"`
		FechaBaja            string                                `json:"FechaBaja"`
		Protecciones         []tipoRegistrarVacunaPerroPropietario `json:"Protecciones"`
	}

	// ---------------------------------------------------------------------------------------------------
	// VALIDAR Argumentos
	// ---------------------------------------------------------------------------------------------------

	if len(args) != 2 {
		return shim.Error("Incorrecto numero de argumentos. Esperando 2")
	}

	DatosVacunaPerroComoJson := []byte(args[0])

	var datosVacunaPerro tipoRegistrarVacunaPerro
	err := json.Unmarshal(DatosVacunaPerroComoJson, &datosVacunaPerro)
	if err != nil {
		fmt.Println(err)
	}

	// ---------------------------------------------------------------------------------------------------

	IDPerro := datosVacunaPerro.IDPerro

	queryString := "{\"selector\":{\"docType\":\"" + cc_perros_cfg.CFG_ObjectType + "\",\"IDPerro\":" + strconv.Itoa(IDPerro) + "}}"
	response := stub.InvokeChaincode(cc_perros_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("ejecutarConsulta", queryString), stub.GetChannelID())

	if response.Status != shim.OK {
		return shim.Error(string(response.Payload))
	}

	if string(response.Payload) == "[]" {
		return shim.Error("(Args[0]) IDPerro: [ " + strconv.Itoa(IDPerro) + " ] no existe o no es valido")
	}

	// ---------------------------------------------------------------------------------------------------

	CODVacuna := strings.ToUpper(datosVacunaPerro.CODVacuna)
	if len(CODVacuna) <= 0 {
		return shim.Error("(Args[0]) CODVacuna: debe tener un valor")
	}

	// ---------------------------------------------------------------------------------------------------

	IDPersonaVeterinario := datosVacunaPerro.IDPersonaVeterinario

	fechaHoraActual := time.Now()
	fechaHoraActualAsString := fechaHoraActual.String()

	queryString = "{\"selector\":{\"docType\":\"" + cc_perfiles_cfg.CFG_ObjectType + "\",\"CODPerfil\":\"" + cc_perfiles_cfg.CFG_PerfilVeterinario + "\",\"IDPersona\":" + strconv.Itoa(IDPersonaVeterinario) + ",\"FechaAlta\":{\"$lt\":     \"" + fechaHoraActualAsString + "\"},\"FechaBaja\":\"\"}}"
	response = stub.InvokeChaincode(cc_perfiles_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("ejecutarConsulta", queryString), stub.GetChannelID())

	if response.Status != shim.OK {
		return shim.Error(string(response.Payload))
	}

	if string(response.Payload) == "[]" {
		return shim.Error("(Args[0]) IDPersonaVeterinario:  [ " + strconv.Itoa(IDPersonaVeterinario) + " ] no tiene un perfil de veterinario activo ")
	}

	// ---------------------------------------------------------------------------------------------------

	FechaBaja := datosVacunaPerro.FechaBaja

	// ---------------------------------------------------------------------------------------------------

	queryString = "{\"selector\":{\"docType\":\"" + cc_cfg.CFG_ObjectType + "\",\"CODVacuna\":\"" + CODVacuna + "\"}}"
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	if string(queryResults) != "[]" {
		return shim.Error("(Args[0]) CODVacuna: [ " + CODVacuna + " ] ya existe")
	}

	// ---------------------------------------------------------------------------------------------------

	Protecciones := datosVacunaPerro.Protecciones
	if len(Protecciones) <= 0 {
		return shim.Error("(Args[0]) Protecciones: debe tener un valor")
	}

	for proteccion := range Protecciones {

		IDVacunaProteccion := Protecciones[proteccion].IDVacunaProteccion

		// ---------------------------------------------------------------------------------------------------

		queryString = "{\"selector\":{\"docType\":\"" + cc_cfg.CFG_ObjectType_VacunasProteccion + "\",\"IDVacunaProteccion\":" + strconv.Itoa(IDVacunaProteccion) + ",\"FechaBaja\":\"\"}}"
		fmt.Println(queryString)
		queryResults, err = getQueryResultForQueryString(stub, queryString)
		if err != nil {
			return shim.Error(err.Error())
		}

		if string(queryResults) == "[]" {
			return shim.Error("(Args[0, " + strconv.Itoa(proteccion) + "]) IDVacunaProteccion: [ " + strconv.Itoa(IDVacunaProteccion) + " ] no existe o no es valido")
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
	response = stub.InvokeChaincode(cc_personas_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("ejecutarConsulta", queryString), stub.GetChannelID())

	if response.Status != shim.OK {
		return shim.Error(string(response.Payload))
	}

	if string(response.Payload) == "[]" {
		return shim.Error("(Args[1]) IDPersona: [ " + strconv.Itoa(IDPersonaEjecuta) + " ] no existe o no es valido")
	}

	// ---------------------------------------------------------------------------------------------------
	// GRABAR Registros
	// ---------------------------------------------------------------------------------------------------

	var retorno bytes.Buffer

	// ---------------------------------------------------------------------------------------------------

	IDVacunaPerro := 0

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
		IDVacunaPerro = queryInfoChaincode.IDMaximo
	}

	IDVacunaPerro += 1

	// ---------------------------------------------------------------------------------------------------

	InfoChaincodeAsBytes, err := json.Marshal(&cc_util.InfoChaincode{cc_util.CC_INFO_CONTADOR, IDVacunaPerro})
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println([]string{cc_cfg.CFG_ObjectType, string(InfoChaincodeAsBytes)})

	err = stub.PutState(cc_cfg.CFG_ObjectType, InfoChaincodeAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ---------------------------------------------------------------------------------------------------

	nuevoVacunaPerro := &cc_cfg.VacunasPerros{cc_cfg.CFG_ObjectType, IDVacunaPerro, IDPerro, IDPersonaVeterinario, CODVacuna, fechaHoraActualAsString, FechaBaja}
	fmt.Println(nuevoVacunaPerro)

	nuevoVacunaPerroAsBytes, err := json.Marshal(nuevoVacunaPerro)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(cc_cfg.CFG_ObjectType+strconv.Itoa(IDVacunaPerro), nuevoVacunaPerroAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println(string(nuevoVacunaPerroAsBytes))
	retorno.WriteString(string(nuevoVacunaPerroAsBytes))

	// ---------------------------------------------------------------------------------------------------

	IDVacunaPerroProteccion := 0

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
		IDVacunaPerroProteccion = queryInfoChaincode.IDMaximo
	}

	// ---------------------------------------------------------------------------------------------------

	for proteccion := range Protecciones {

		IDVacunaPerroProteccion += 1

		IDVacunaProteccion := Protecciones[proteccion].IDVacunaProteccion
		FechaBaja := Protecciones[proteccion].FechaBaja

		nuevoVacunaPerroProteccion := &cc_cfg.VacunasPerrosProteccion{cc_cfg.CFG_ObjectType_VacunasPerrosProteccion, IDVacunaPerroProteccion, IDVacunaPerro, IDVacunaProteccion, fechaHoraActualAsString, FechaBaja}
		fmt.Println(nuevoVacunaPerroProteccion)

		nuevoVacunaPerroProteccionAsBytes, err := json.Marshal(nuevoVacunaPerroProteccion)
		if err != nil {
			return shim.Error(err.Error())
		}

		err = stub.PutState(cc_cfg.CFG_ObjectType_VacunasProteccion+strconv.Itoa(IDVacunaPerroProteccion), nuevoVacunaPerroProteccionAsBytes)
		if err != nil {
			return shim.Error(err.Error())
		}

		fmt.Println(string(nuevoVacunaPerroProteccionAsBytes))
		retorno.WriteString(string(nuevoVacunaPerroProteccionAsBytes))

	}

	// ---------------------------------------------------------------------------------------------------

	InfoChaincodeAsBytes, err = json.Marshal(&cc_util.InfoChaincode{cc_util.CC_INFO_CONTADOR, IDVacunaPerroProteccion})
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println([]string{cc_cfg.CFG_ObjectType_VacunasPerrosProteccion, string(InfoChaincodeAsBytes)})

	err = stub.PutState(cc_cfg.CFG_ObjectType_VacunasPerrosProteccion, InfoChaincodeAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ---------------------------------------------------------------------------------------------------

	return shim.Success(retorno.Bytes())
}

func (tcc *ThisChainCode) obtenerCertificadoVacunaciones(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))

	retorno := " ********** FUNCION " + cc_util.NombreFuncion() + " SIN IMPLEMENTAR **********"
	fmt.Println(retorno)

	return shim.Success([]byte(retorno))
}

func (tcc *ThisChainCode) consultarVacunaciones(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))

	retorno := " ********** FUNCION " + cc_util.NombreFuncion() + " SIN IMPLEMENTAR **********"
	fmt.Println(retorno)

	return shim.Success([]byte(retorno))
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
