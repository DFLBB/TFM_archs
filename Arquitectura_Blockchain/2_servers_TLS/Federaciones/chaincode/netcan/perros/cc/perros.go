package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	cc_afijos_cfg "github.com/chaincode/netcan/afijos/lib/afijos_cfg"
	cc_util "github.com/chaincode/netcan/netcan/lib/netcan_util"

	cc_cfg "github.com/chaincode/netcan/perros/lib/perros_cfg"
	cc_personas_cfg "github.com/chaincode/netcan/personas/lib/personas_cfg"
	cc_razas_cfg "github.com/chaincode/netcan/razas/lib/razas_cfg"

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

	if function == "registrarPerro" {
		return tcc.registrarPerro(stub, args)

	} else if function == "registrarCambioPropietario" {
		return tcc.registrarCambioPropietario(stub, args)
	} else if function == "registrarDefuncionPerro" {
		return tcc.registrarDefuncionPerro(stub, args)

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

	} else {
		return shim.Error("(" + cc_cfg.CFG_ObjectType + ") Invalida un nombre de funcion no valida (" + function + ")")
	}
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (tcc *ThisChainCode) cargarDatosIniciales(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))

	NombreArchivo := args[0]

	nuevosRegistrosAsJson, err := ioutil.ReadFile(NombreArchivo)
	if err != nil {
		return shim.Error(err.Error())
	}

	var nuevosRegistros []cc_cfg.Perros
	json.Unmarshal(nuevosRegistrosAsJson, &nuevosRegistros)

	for _, nuevoRegistro := range nuevosRegistros {

		fmt.Println(nuevoRegistro)

		nuevoRegistroAsBytes, err := json.Marshal(nuevoRegistro)
		if err != nil {
			return shim.Error(err.Error())
		}

		err = stub.PutState(cc_cfg.CFG_ObjectType+strconv.Itoa(nuevoRegistro.IDPerro), nuevoRegistroAsBytes)
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

	var nuevosRegistros []cc_cfg.PerrosPropietarios
	json.Unmarshal(nuevosRegistrosAsJson, &nuevosRegistros)

	for _, nuevoRegistro := range nuevosRegistros {

		fmt.Println(nuevoRegistro)

		nuevoRegistroAsBytes, err := json.Marshal(nuevoRegistro)
		if err != nil {
			return shim.Error(err.Error())
		}

		err = stub.PutState(cc_cfg.CFG_ObjectType_Propietarios+strconv.Itoa(nuevoRegistro.IDPerroPropietario), nuevoRegistroAsBytes)
		if err != nil {
			return shim.Error(err.Error())

		}
	}

	return shim.Success([]byte(strconv.Itoa(len(nuevosRegistros))))
}

func (tcc *ThisChainCode) registrarPerro(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))
	fmt.Println(args[0])
	fmt.Println(args[1])

	type tipoRegistrarCamadaPerro struct {
		Nombre string `json:"Nombre"`
		IDSexo int    `json:"IDSexo"`
	}

	type tipoRegistrarCamadaPropietario struct {
		IDPersona int `json:"IDPersona"`
	}

	type tipoRegistrarCamada struct {
		Perros          []tipoRegistrarCamadaPerro       `json:"Perros"`
		IDPerroMadre    int                              `json:"IDPerroMadre"`
		IDPerroPadre    int                              `json:"IDPerroPadre"`
		IDAfijo         int                              `json:"IDAfijo"`
		IDRaza          int                              `json:"IDRaza"`
		FechaNacimiento string                           `json:"FechaNacimiento"`
		Propietarios    []tipoRegistrarCamadaPropietario `json:"Propietarios"`
	}

	type tipoQueryAfijos struct {
		Key    string               `json:"Key"`
		Record cc_afijos_cfg.Afijos `json:"Record"`
	}

	type tipoQueryPerros struct {
		Key    string        `json:"Key"`
		Record cc_cfg.Perros `json:"Record"`
	}

	// ---------------------------------------------------------------------------------------------------
	// VALIDAR Argumentos
	// ---------------------------------------------------------------------------------------------------

	if len(args) != 2 {
		return shim.Error("Incorrecto numero de argumentos. Esperando 2")
	}

	// ---------------------------------------------------------------------------------------------------

	DatosSeguridadComoJson := []byte(args[1])

	var queryString string
	var queryResults []byte
	var err error

	var datosSeguridad cc_util.TipoSeguridad

	err = json.Unmarshal(DatosSeguridadComoJson, &datosSeguridad)
	if err != nil {
		fmt.Println(err)
	}

	IDPersonaEjecuta := datosSeguridad.IDPersona

	queryString = "{\"selector\":{\"docType\":\"" + cc_personas_cfg.CFG_ObjectType + "\",\"IDPersona\":" + strconv.Itoa(IDPersonaEjecuta) + "}}"
	//fmt.Println(queryString)
	response := stub.InvokeChaincode(cc_personas_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("ejecutarConsulta", queryString), stub.GetChannelID())

	if response.Status != shim.OK {
		return shim.Error(string(response.Payload))
	}

	if string(response.Payload) == "[]" {
		return shim.Error("(Args[1]) IDPersona: [ " + strconv.Itoa(IDPersonaEjecuta) + " ] no existe o no es valido")
	}

	// ---------------------------------------------------------------------------------------------------

	DatosPerroComoJson := []byte(args[0])

	var datosPerro tipoRegistrarCamada
	err = json.Unmarshal(DatosPerroComoJson, &datosPerro)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(datosPerro)

	IDAfijo := datosPerro.IDAfijo

	// Comprobar IDAfijo existe y esta activo

	if IDAfijo > 0 {

		queryString = "{\"selector\":{\"docType\":\"" + cc_afijos_cfg.CFG_ObjectType + "\",\"IDAfijo\":" + strconv.Itoa(IDAfijo) + ",\"FechaBaja\":\"\"}}"
		//fmt.Println(queryString)
		response := stub.InvokeChaincode(cc_afijos_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("ejecutarConsulta", queryString), stub.GetChannelID())

		if response.Status != shim.OK {
			return shim.Error(string(response.Payload))
		}

		if string(response.Payload) == "[]" {
			return shim.Error("(Args[0]) IDAfijo: [ " + strconv.Itoa(IDAfijo) + " ] no existe o no es valido")
		}

	} else {
		IDAfijo = 0
	}

	IDPerroMadre := datosPerro.IDPerroMadre

	if IDPerroMadre > 0 {

		queryString = "{\"selector\":{\"docType\":\"" + cc_cfg.CFG_ObjectType + "\",\"IDPerro\":" + strconv.Itoa(IDPerroMadre) + "}}"

		queryResults, err = getQueryResultForQueryString(stub, queryString)
		if err != nil {
			return shim.Error(err.Error())
		}

		if string(queryResults) == "[]" {
			return shim.Error("(Args[0]) IDPerroMadre: la perra " + strconv.Itoa(IDPerroMadre) + " no existe")
		}

		var queryPerros []tipoQueryPerros

		err = json.Unmarshal(queryResults, &queryPerros)
		if err != nil {
			return shim.Error(err.Error())
		}

		// Comprobar IDMadre es hembra
		if queryPerros[0].Record.IDSexo != 0 {
			return shim.Error("(Args[0]) IDPerroMadre: el perro " + strconv.Itoa(IDPerroMadre) + " esta registrado como un ejemplar macho")
		}

		// Comprobar IDMadre no esta dada de baja
		if queryPerros[0].Record.FechaBaja != "" {
			return shim.Error("(Args[0]) IDPerroMadre: la perra " + strconv.Itoa(IDPerroMadre) + " esta dada de baja")
		}

		// Comprobar IDPerroMadre la edad necesaria ( 1 < edad < 10 años )
		fmt.Println("************** FALTA IMPLEMENTAR: Comprobacion de la edad ( 1 < edad < 10 años ) de IDPerroMadre")

	} else {
		IDPerroMadre = 0
	}

	IDPerroPadre := datosPerro.IDPerroPadre

	if IDPerroPadre > 0 {

		queryString = "{\"selector\":{\"docType\":\"" + cc_cfg.CFG_ObjectType + "\",\"IDPerro\":" + strconv.Itoa(IDPerroPadre) + "}}"

		queryResults, err = getQueryResultForQueryString(stub, queryString)
		if err != nil {
			return shim.Error(err.Error())
		}

		if string(queryResults) == "[]" {
			return shim.Error("(Args[0]) IDPerroPadre: el perro " + strconv.Itoa(IDPerroPadre) + " no existe")
		}

		var queryPerros []tipoQueryPerros

		err = json.Unmarshal(queryResults, &queryPerros)
		if err != nil {
			return shim.Error(err.Error())
		}

		// Comprobar IDPadre es hembra
		if queryPerros[0].Record.IDSexo != 1 {
			return shim.Error("(Args[0]) IDPerroPadre: el perro " + strconv.Itoa(IDPerroPadre) + " esta registrado como un ejemplar hembra")
		}

		// Comprobar IDPadre no esta dada de baja
		if queryPerros[0].Record.FechaBaja != "" {
			return shim.Error("(Args[0]) IDPerroPadre: la perra " + strconv.Itoa(IDPerroPadre) + " esta dada de baja")
		}

		// Comprobar IDPerroPadre la edad necesaria ( 1 < edad < 10 años )
		fmt.Println("************** FALTA IMPLEMENTAR: Comprobacion de la edad ( 1 < edad < 12 años ) de IDPerroPadre")

	} else {
		IDPerroPadre = 0
	}

	IDRaza := datosPerro.IDRaza

	fmt.Println(IDRaza)

	if IDRaza > 0 {

		queryString = "{\"selector\":{\"docType\":\"" + cc_razas_cfg.CFG_ObjectType + "\",\"IDRaza\":" + strconv.Itoa(IDRaza) + "}}"
		fmt.Println(queryString)
		response := stub.InvokeChaincode(cc_razas_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("ejecutarConsulta", queryString), stub.GetChannelID())

		if response.Status != shim.OK {
			return shim.Error(string(response.Payload))
		}

		if string(response.Payload) == "[]" {
			return shim.Error("(Args[0]) IDRaza: [ " + strconv.Itoa(IDRaza) + " ] no existe o no es valido")
		}

	} else {
		IDRaza = 0
	}

	FechaNacimiento := datosPerro.FechaNacimiento
	if len(FechaNacimiento) <= 0 {
		return shim.Error("(Args[0]) FechaNacimiento: no tiene un Fecha definida")
	}
	fmt.Println("************** FALTA IMPLEMENTAR!!!: Comprobacion de la FechaNacimiento sea una Fecha valida")

	PerrosCamada := datosPerro.Perros
	if len(PerrosCamada) <= 0 {
		return shim.Error("(Args[0]) PerrosCamada: debe tener un perro")
	}

	for perroCamada := range PerrosCamada {

		// Comprobar IDSexo es 0 o 1
		IDSexo := PerrosCamada[perroCamada].IDSexo
		if IDSexo != 0 && IDSexo != 1 {
			return shim.Error("(Args[0, " + strconv.Itoa(perroCamada) + "]) IDSexo: el valor " + strconv.Itoa(IDSexo) + " no es valido (0=HEMBRA, 1=MACHO")
		}

		NombrePerro := strings.ToUpper(PerrosCamada[perroCamada].Nombre)
		fmt.Println(NombrePerro)

		// Comprobar que el Nombre esta relleno
		if len(NombrePerro) <= 0 {
			return shim.Error("(Args[0, " + strconv.Itoa(perroCamada) + "])Nombre: debe tener un valor")
		}

		if IDAfijo != 0 {

			queryString = "{\"selector\":{\"docType\":\"" + cc_cfg.CFG_ObjectType + "\",\"IDAfijo\":" + strconv.Itoa(IDAfijo) + ",\"Nombre\":\"" + NombrePerro + "\"}}"
			queryResults, err = getQueryResultForQueryString(stub, queryString)
			if err != nil {
				return shim.Error(err.Error())
			}

			// Comprobar que no existe un perro que se llame igual para ese afijo
			if string(queryResults) != "[]" {
				return shim.Error("(Args[0, " + strconv.Itoa(perroCamada) + "]) Nombre: ya un perro registrado para el afijo " + strconv.Itoa(IDAfijo) + " con el Nombre " + NombrePerro + " " + string(queryResults))
			}
		}

	}

	var Propietarios []tipoRegistrarCamadaPropietario

	if IDPerroMadre == 0 {

		Propietarios = datosPerro.Propietarios
		if len(Propietarios) <= 0 {
			return shim.Error("(Args[0]) Propietarios: debe tener un propietario")
		}

		for propietario := range Propietarios {

			queryString = "{\"selector\":{\"docType\":\"" + cc_personas_cfg.CFG_ObjectType + "\",\"IDPersona\":" + strconv.Itoa(Propietarios[propietario].IDPersona) + "}}"
			fmt.Println(queryString)
			response := stub.InvokeChaincode(cc_personas_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("ejecutarConsulta", queryString), stub.GetChannelID())

			if response.Status != shim.OK {
				return shim.Error(string(response.Payload))
			}

			if string(response.Payload) == "[]" {
				return shim.Error("(Args[0, " + strconv.Itoa(propietario) + "]) IDPersona: [ " + strconv.Itoa(Propietarios[propietario].IDPersona) + " ] no existe o no es valido")
			}

			// ---------------------------------------------------------------------------------------------------

		}

	} else {

		queryString = "{\"selector\":{\"docType\":\"" + cc_cfg.CFG_ObjectType_Propietarios + "\",\"IDPerro\":" + strconv.Itoa(IDPerroMadre) + ",\"FechaBaja\":\"\"}}"
		queryResults, err = getQueryResultForQueryString(stub, queryString)

		if err != nil {
			return shim.Error(err.Error())
		}

		fmt.Println("···································")
		fmt.Println(queryResults)

		err = json.Unmarshal(queryResults, &Propietarios)
		if err != nil {
			fmt.Println(err)
		}

	}

	// ---------------------------------------------------------------------------------------------------
	// GRABAR Registros
	// ---------------------------------------------------------------------------------------------------

	fmt.Println("GRABAR Registros")

	var retorno bytes.Buffer

	FechaHoraActual := time.Now()
	FechaHoraActualAsString := FechaHoraActual.String()

	// ---------------------------------------------------------------------------------------------------

	IDPerro := 0

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
		IDPerro = queryInfoChaincode.IDMaximo
	}

	// ---------------------------------------------------------------------------------------------------

	IDPerroPropietario := 0

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
		IDPerroPropietario = queryInfoChaincode.IDMaximo
	}

	// ---------------------------------------------------------------------------------------------------

	for perroCamada := range PerrosCamada {

		IDPerro += 1

		NombrePerro := strings.ToUpper(PerrosCamada[perroCamada].Nombre)
		IDSexo := PerrosCamada[perroCamada].IDSexo
		nuevoPerro := &cc_cfg.Perros{cc_cfg.CFG_ObjectType, IDPerro, NombrePerro, IDAfijo, IDSexo, IDPerroMadre, IDPerroPadre, IDRaza, FechaNacimiento, "", FechaHoraActualAsString, ""}

		nuevoPerroAsBytes, err := json.Marshal(nuevoPerro)
		if err != nil {
			return shim.Error(err.Error())
		}

		err = stub.PutState(cc_cfg.CFG_ObjectType+strconv.Itoa(IDPerro), nuevoPerroAsBytes)
		if err != nil {
			return shim.Error(err.Error())
		}

		fmt.Println(nuevoPerro)
		retorno.WriteString(string(nuevoPerroAsBytes))

		// ---------------------------------------------------------------------------------------------------

		for propietario := range Propietarios {

			IDPerroPropietario += 1

			nuevoPerroPropietario := &cc_cfg.PerrosPropietarios{cc_cfg.CFG_ObjectType_Propietarios, IDPerroPropietario, IDPerro, Propietarios[propietario].IDPersona, FechaHoraActualAsString, ""}
			nuevoPerroPropietarioAsBytes, err := json.Marshal(nuevoPerroPropietario)
			if err != nil {
				return shim.Error(err.Error())
			}

			err = stub.PutState(cc_cfg.CFG_ObjectType_Propietarios+strconv.Itoa(IDPerroPropietario), nuevoPerroPropietarioAsBytes)
			if err != nil {
				return shim.Error(err.Error())
			}

			fmt.Println(nuevoPerroPropietario)
			retorno.WriteString(string(nuevoPerroPropietarioAsBytes))

		}
	}

	// ---------------------------------------------------------------------------------------------------

	InfoChaincodeAsBytes, err := json.Marshal(&cc_util.InfoChaincode{cc_util.CC_INFO_CONTADOR, IDPerro})
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println([]string{cc_cfg.CFG_ObjectType, string(InfoChaincodeAsBytes)})

	err = stub.PutState(cc_cfg.CFG_ObjectType, InfoChaincodeAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ---------------------------------------------------------------------------------------------------

	InfoChaincodeAsBytes, err = json.Marshal(&cc_util.InfoChaincode{cc_util.CC_INFO_CONTADOR, IDPerroPropietario})
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

	type tiporegistrarCambioPropietarioPropietario struct {
		IDPersona int `json:"IDPersona"`
	}

	type tiporegistrarCambioPropietario struct {
		IDPerro      int                                         `json:"IDPerro"`
		Propietarios []tiporegistrarCambioPropietarioPropietario `json:"Propietarios"`
	}

	type tipoQueryPerrosPropietarios struct {
		Key    string                    `json:"Key"`
		Record cc_cfg.PerrosPropietarios `json:"Record"`
	}

	// ---------------------------------------------------------------------------------------------------
	// VALIDAR Argumentos
	// ---------------------------------------------------------------------------------------------------

	if len(args) != 2 {
		return shim.Error("Incorrecto numero de argumentos. Esperando 2")
	}

	DatosPerroComoJson := []byte(args[0])

	var datosPerro tiporegistrarCambioPropietario
	err := json.Unmarshal(DatosPerroComoJson, &datosPerro)
	if err != nil {
		fmt.Println(err)
	}

	IDPerro := datosPerro.IDPerro
	if IDPerro <= 0 {
		return shim.Error("(Args[0]) Nombre: debe tener un valor")
	}

	queryString := "{\"selector\":{\"docType\":\"" + cc_cfg.CFG_ObjectType + "\",\"IDPerro\":" + strconv.Itoa(IDPerro) + ",\"FechaBaja\":\"\"}}"
	// fmt.Println(queryString)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	// fmt.Println(string(queryResults))

	if string(queryResults) == "[]" {
		return shim.Error("(Args[0) IDPerro: [ " + strconv.Itoa(IDPerro) + " ] no existe o no es valido")
	}

	// ---------------------------------------------------------------------------------------------------

	Propietarios := datosPerro.Propietarios
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
			return shim.Error(string(response.Payload))
		}

		if string(response.Payload) == "[]" {
			return shim.Error("(Args[0, " + strconv.Itoa(propietario) + "]) IDPersona: [ " + strconv.Itoa(IDPersonaPropietario) + " ] no existe o no es valido")
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
	// fmt.Println(queryString)
	response := stub.InvokeChaincode(cc_personas_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("ejecutarConsulta", queryString), stub.GetChannelID())

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

	fechaHoraActual := time.Now()
	fechaHoraActualAsString := fechaHoraActual.String()

	// ---------------------------------------------------------------------------------------------------

	queryString = "{\"selector\":{\"docType\":\"" + cc_cfg.CFG_ObjectType_Propietarios + "\",\"IDPerro\":" + strconv.Itoa(IDPerro) + ",\"FechaBaja\":\"\"}}"
	queryResults, err = getQueryResultForQueryString(stub, queryString)

	if err != nil {
		return shim.Error(err.Error())
	}

	var datosRegistroPropietarios []tipoQueryPerrosPropietarios

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

		err = stub.PutState(cc_cfg.CFG_ObjectType_Propietarios+strconv.Itoa(datosRegistroPropietarios[propietario].Record.IDPerroPropietario), datosRegistroPropietariosAsBytes)
		if err != nil {
			return shim.Error(err.Error())
		}

		fmt.Println(datosRegistroPropietarios[propietario].Record)
		retorno.WriteString(string(datosRegistroPropietariosAsBytes))
	}

	// ---------------------------------------------------------------------------------------------------

	IDPerroPropietario := 0

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
		IDPerroPropietario = queryInfoChaincode.IDMaximo
	}

	// ---------------------------------------------------------------------------------------------------

	fmt.Println("***ALTAS***")

	for propietario := range Propietarios {

		IDPerroPropietario += 1

		nuevoPerroPropietario := &cc_cfg.PerrosPropietarios{cc_cfg.CFG_ObjectType_Propietarios, IDPerroPropietario, IDPerro, Propietarios[propietario].IDPersona, fechaHoraActualAsString, ""}

		nuevoPerroPropietarioAsBytes, err := json.Marshal(nuevoPerroPropietario)
		if err != nil {
			return shim.Error(err.Error())
		}

		err = stub.PutState(cc_cfg.CFG_ObjectType_Propietarios+strconv.Itoa(IDPerroPropietario), nuevoPerroPropietarioAsBytes)
		if err != nil {
			return shim.Error(err.Error())
		}

		fmt.Println(nuevoPerroPropietario)
		retorno.WriteString(string(nuevoPerroPropietarioAsBytes))

	}

	// ---------------------------------------------------------------------------------------------------

	InfoChaincodeAsBytes, err := json.Marshal(&cc_util.InfoChaincode{cc_util.CC_INFO_CONTADOR, IDPerroPropietario})
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

func (tcc *ThisChainCode) registrarDefuncionPerro(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))

	type tiporegistrarDefuncionPerro struct {
		IDPerro        int    `json:"IDPerro"`
		FechaDefuncion string `json:"FechaDefuncion"`
	}

	type tipoQueryPerros struct {
		Key    string        `json:"Key"`
		Record cc_cfg.Perros `json:"Record"`
	}

	type tipoQueryPerrosPropietarios struct {
		Key    string                    `json:"Key"`
		Record cc_cfg.PerrosPropietarios `json:"Record"`
	}

	// ---------------------------------------------------------------------------------------------------
	// VALIDAR Argumentos
	// ---------------------------------------------------------------------------------------------------

	if len(args) != 2 {
		return shim.Error("Incorrecto numero de argumentos. Esperando 2")
	}

	DatosPerroComoJson := []byte(args[0])

	var datosPerro tiporegistrarDefuncionPerro
	err := json.Unmarshal(DatosPerroComoJson, &datosPerro)
	if err != nil {
		fmt.Println(err)
	}

	IDPerro := datosPerro.IDPerro
	if IDPerro <= 0 {
		return shim.Error("(Args[0]) IDPerro: debe tener un valor")
	}

	queryString := "{\"selector\":{\"docType\":\"" + cc_cfg.CFG_ObjectType + "\",\"IDPerro\":" + strconv.Itoa(IDPerro) + ",\"FechaBaja\":\"\"}}"
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	if string(queryResults) == "[]" {
		return shim.Error("(Args[0) IDPerro: [ " + strconv.Itoa(IDPerro) + " ] no existe o no es valido")
	}

	// ---------------------------------------------------------------------------------------------------

	FechaDefuncion := datosPerro.FechaDefuncion
	if len(FechaDefuncion) <= 0 {
		return shim.Error("(Args[0]) FechaDefuncion: no tiene un Fecha definida")
	}
	fmt.Println("************** FALTA IMPLEMENTAR: Comprobacion de la FechaDefuncion sea una Fecha valida")

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
		return shim.Error(string(response.Payload))
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

	var datosRegistro []tipoQueryPerros

	err = json.Unmarshal(queryResults, &datosRegistro)
	if err != nil {
		fmt.Println(err)
	}

	datosRegistro[0].Record.FechaBaja = fechaHoraActualAsString
	datosRegistro[0].Record.FechaDefuncion = FechaDefuncion

	datosRegistroAsBytes, err := json.Marshal(datosRegistro[0].Record)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(cc_cfg.CFG_ObjectType+strconv.Itoa(IDPerro), datosRegistroAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println(datosRegistro[0].Record)
	retorno.WriteString(string(datosRegistroAsBytes))

	// ---------------------------------------------------------------------------------------------------

	queryString = "{\"selector\":{\"docType\":\"" + cc_cfg.CFG_ObjectType_Propietarios + "\",\"IDPerro\":" + strconv.Itoa(IDPerro) + ",\"FechaBaja\":\"\"}}"
	queryResults, err = getQueryResultForQueryString(stub, queryString)

	if err != nil {
		return shim.Error(err.Error())
	}

	var datosRegistroPropietarios []tipoQueryPerrosPropietarios

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

		err = stub.PutState(cc_cfg.CFG_ObjectType_Propietarios+strconv.Itoa(datosRegistroPropietarios[propietario].Record.IDPerroPropietario), datosRegistroPropietariosAsBytes)
		if err != nil {
			return shim.Error(err.Error())
		}

		fmt.Println(datosRegistroPropietarios[propietario].Record)
		retorno.WriteString(string(datosRegistroPropietariosAsBytes))
	}

	return shim.Success(retorno.Bytes())

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
