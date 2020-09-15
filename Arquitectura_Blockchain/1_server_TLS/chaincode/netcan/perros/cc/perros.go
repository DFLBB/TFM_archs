package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"sort"
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

	} else if function == "registrarReconocimientoRaza" {
		return tcc.registrarReconocimientoRaza(stub, args)

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
	} else if function == "consultarDatosEjemplar" {
		return tcc.consultarDatosEjemplar(stub, args)
	} else if function == "obtenerCertificadoRegistro" {
		return tcc.obtenerCertificadoRegistro(stub, args)
	} else if function == "obtenerPedigri" {
		return tcc.obtenerPedigri(stub, args)
	} else if function == "registrarCesionTemporal " {
		return tcc.registrarCesionTemporal(stub, args)

	} else {
		return shim.Error("(" + cc_cfg.CFG_ObjectType + ") Invoca un nombre de funcion no valida (" + function + ")")
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
	var response pb.Response

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

	DatosPerroComoJson := []byte(args[0])

	var datosPerro cc_cfg.TipoRegistrarCamada
	err = json.Unmarshal(DatosPerroComoJson, &datosPerro)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(datosPerro)

	// ---------------------------------------------------------------------------------------------------

	IDPerroMadre := datosPerro.IDPerroMadre
	IDRazaMadre := 0

	if IDPerroMadre > 0 {

		queryString = "{\"selector\":{\"docType\":\"" + cc_cfg.CFG_ObjectType + "\",\"IDPerro\":" + strconv.Itoa(IDPerroMadre) + "}}"

		queryResults, err = getQueryResultForQueryString(stub, queryString)
		if err != nil {
			return shim.Error(err.Error())
		}

		if string(queryResults) == "[]" {
			return shim.Error("(Args[0]) IDPerroMadre: la perra " + strconv.Itoa(IDPerroMadre) + " no existe")
		}

		var queryPerros []cc_cfg.TipoQueryPerros

		err = json.Unmarshal(queryResults, &queryPerros)
		if err != nil {
			return shim.Error(err.Error())
		}

		// Comprobar IDMadre es hembra
		if queryPerros[0].Record.IDSexo != cc_cfg.HEMBRA {
			return shim.Error("(Args[0]) IDPerroMadre: el perro " + strconv.Itoa(IDPerroMadre) + " esta registrado como un ejemplar macho")
		}

		// Comprobar IDMadre no esta dada de baja
		if queryPerros[0].Record.FechaBaja != "" {
			return shim.Error("(Args[0]) IDPerroMadre: la perra " + strconv.Itoa(IDPerroMadre) + " esta dada de baja")
		}

		// Comprobar IDPerroMadre la edad necesaria ( 1 < edad < 10 años )
		fmt.Println("************** FALTA IMPLEMENTAR: Comprobacion de la edad ( 1 < edad < 10 años ) de IDPerroMadre")

		IDRazaMadre = queryPerros[0].Record.IDRaza

	} else {
		IDPerroMadre = 0
	}

	// ---------------------------------------------------------------------------------------------------

	IDPerroPadre := datosPerro.IDPerroPadre
	IDRazaPadre := 0

	if IDPerroPadre > 0 {

		queryString = "{\"selector\":{\"docType\":\"" + cc_cfg.CFG_ObjectType + "\",\"IDPerro\":" + strconv.Itoa(IDPerroPadre) + "}}"

		queryResults, err = getQueryResultForQueryString(stub, queryString)
		if err != nil {
			return shim.Error(err.Error())
		}

		if string(queryResults) == "[]" {
			return shim.Error("(Args[0]) IDPerroPadre: el perro " + strconv.Itoa(IDPerroPadre) + " no existe")
		}

		var queryPerros []cc_cfg.TipoQueryPerros

		err = json.Unmarshal(queryResults, &queryPerros)
		if err != nil {
			return shim.Error(err.Error())
		}

		// Comprobar IDPadre es hembra
		if queryPerros[0].Record.IDSexo != cc_cfg.MACHO {
			return shim.Error("(Args[0]) IDPerroPadre: el perro " + strconv.Itoa(IDPerroPadre) + " esta registrado como un ejemplar hembra")
		}

		// Comprobar IDPadre no esta dada de baja
		if queryPerros[0].Record.FechaBaja != "" {
			return shim.Error("(Args[0]) IDPerroPadre: la perra " + strconv.Itoa(IDPerroPadre) + " esta dada de baja")
		}

		// Comprobar IDPerroPadre la edad necesaria ( 1 < edad < 10 años )
		fmt.Println("************** FALTA IMPLEMENTAR: Comprobacion de la edad ( 1 < edad < 12 años ) de IDPerroPadre")

		IDRazaPadre = queryPerros[0].Record.IDRaza

	} else {
		IDPerroPadre = 0
	}

	IDRaza := 0

	if IDRazaMadre == IDRazaPadre {

		IDRaza = IDRazaMadre

		queryString = "{\"selector\":{\"docType\":\"" + cc_razas_cfg.CFG_ObjectType + "\",\"IDRaza\":" + strconv.Itoa(IDRaza) + ",\"FechaBaja\":\"\"}}"
		response := stub.InvokeChaincode(cc_razas_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("ejecutarConsulta", queryString), stub.GetChannelID())

		if response.Status != shim.OK {
			return shim.Error(string(response.Payload))
		}

		if string(response.Payload) == "[]" {
			return shim.Error("(Args[0]) IDRaza: [ " + strconv.Itoa(IDRaza) + " ] no existe o no es valido")
		}

	}

	// ---------------------------------------------------------------------------------------------------

	FechaNacimiento := datosPerro.FechaNacimiento
	if len(FechaNacimiento) <= 0 {
		return shim.Error("(Args[0]) FechaNacimiento: no tiene un Fecha definida")
	}
	fmt.Println("************** FALTA IMPLEMENTAR!!!: Comprobacion de la FechaNacimiento sea una Fecha valida")

	// ---------------------------------------------------------------------------------------------------

	PerrosCamada := datosPerro.Perros
	if len(PerrosCamada) <= 0 {
		return shim.Error("(Args[0]) PerrosCamada: debe tener un perro")
	}

	// ---------------------------------------------------------------------------------------------------

	propietariosArray := []int{}
	IDAfijo := 0

	if IDPerroMadre > 0 {

		if len(datosPerro.Propietarios) > 0 {
			return shim.Error("(Args[0]) Propietarios: no debe tener definido un propietario")
		}

		var queryPerrosPropietarios []cc_cfg.TipoQueryPerrosPropietarios

		queryString = "{\"selector\":{\"docType\":\"" + cc_cfg.CFG_ObjectType_Propietarios + "\",\"IDPerro\":" + strconv.Itoa(IDPerroMadre) + ",\"FechaBaja\":\"\"}}"
		queryResults, err = getQueryResultForQueryString(stub, queryString)

		if err != nil {
			return shim.Error(err.Error())
		}

		err = json.Unmarshal(queryResults, &queryPerrosPropietarios)
		if err != nil {
			fmt.Println(err)
		}

		afijosArray := []int{}

		for propietario := range queryPerrosPropietarios {

			propietariosArray = append(propietariosArray, queryPerrosPropietarios[propietario].Record.IDPersona)

			queryString = "{\"selector\":{\"docType\":\"" + cc_afijos_cfg.CFG_ObjectType_Propietarios + "\",\"IDPersona\":" + strconv.Itoa(queryPerrosPropietarios[propietario].Record.IDPersona) + ",\"FechaBaja\":\"\"}}"
			response := stub.InvokeChaincode(cc_afijos_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("ejecutarConsulta", queryString), stub.GetChannelID())

			if response.Status != shim.OK {
				return shim.Error(string(response.Payload))
			}

			fmt.Println(string(response.Payload))

			if string(response.Payload) != "[]" {

				var queryAfijos []cc_afijos_cfg.TipoQueryAfijos

				err = json.Unmarshal(response.Payload, &queryAfijos)
				if err != nil {
					fmt.Println(err)
				}
				afijosArray = append(afijosArray, queryAfijos[0].Record.IDAfijo)
			}
		}

		afijosArray = cc_util.ArrayIntSinDuplicados(afijosArray)

		if len(afijosArray) == 1 {

			propietariosAfijoArray := []int{}

			queryString = "{\"selector\":{\"docType\":\"" + cc_afijos_cfg.CFG_ObjectType_Propietarios + "\",\"IDAfijo\":" + strconv.Itoa(afijosArray[0]) + ",\"FechaBaja\":\"\"}}"
			response := stub.InvokeChaincode(cc_afijos_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("ejecutarConsulta", queryString), stub.GetChannelID())

			if response.Status != shim.OK {
				return shim.Error(string(response.Payload))
			}

			fmt.Println(string(response.Payload))

			if string(response.Payload) != "[]" {

				var queryAfijosPropietarios []cc_afijos_cfg.TipoQueryAfijosPropietarios

				err = json.Unmarshal(response.Payload, &queryAfijosPropietarios)
				if err != nil {
					fmt.Println(err)
				}

				for propietario := range queryAfijosPropietarios {
					propietariosAfijoArray = append(propietariosAfijoArray, queryAfijosPropietarios[propietario].Record.IDPersona)
				}

				propietariosArray = cc_util.ArrayIntSinDuplicados(propietariosArray)
				sort.Ints(propietariosArray)
				propietariosAfijoArray = cc_util.ArrayIntSinDuplicados(propietariosAfijoArray)
				sort.Ints(propietariosAfijoArray)

				if reflect.DeepEqual(propietariosArray, propietariosAfijoArray) {
					IDAfijo = afijosArray[0]
				}
			}
			fmt.Println("-------------------------------------------")
			fmt.Println(propietariosArray)
			fmt.Println(afijosArray)
			fmt.Println(propietariosAfijoArray)
			fmt.Println("-------------------------------------------")

		}

	} else {

		var Propietarios []cc_cfg.TipoRegistrarCamadaPropietario

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

			propietariosArray = append(propietariosArray, Propietarios[propietario].IDPersona)
		}

		propietariosArray = cc_util.ArrayIntSinDuplicados(propietariosArray)

	}

	for perroCamada := range PerrosCamada {

		// Comprobar IDSexo es 0 o 1
		IDSexo := PerrosCamada[perroCamada].IDSexo
		if IDSexo != cc_cfg.HEMBRA && IDSexo != cc_cfg.MACHO {
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

		for propietario := range propietariosArray {

			IDPerroPropietario += 1

			nuevoPerroPropietario := &cc_cfg.PerrosPropietarios{cc_cfg.CFG_ObjectType_Propietarios, IDPerroPropietario, IDPerro, propietariosArray[propietario], FechaHoraActualAsString, ""}
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

	// ---------------------------------------------------------------------------------------------------
	// VALIDAR Argumentos
	// ---------------------------------------------------------------------------------------------------

	if len(args) != 2 {
		return shim.Error("Incorrecto numero de argumentos. Esperando 2")
	}

	DatosPerroComoJson := []byte(args[0])

	var datosPerro cc_cfg.TipoRegistrarCambioPropietario
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

	var datosRegistroPropietarios []cc_cfg.TipoQueryPerrosPropietarios

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

	// ---------------------------------------------------------------------------------------------------
	// VALIDAR Argumentos
	// ---------------------------------------------------------------------------------------------------

	if len(args) != 2 {
		return shim.Error("Incorrecto numero de argumentos. Esperando 2")
	}

	DatosPerroComoJson := []byte(args[0])

	var datosPerro cc_cfg.TipoRegistrarDefuncionPerro
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

	var datosRegistro []cc_cfg.TipoQueryPerros

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

	var datosRegistroPropietarios []cc_cfg.TipoQueryPerrosPropietarios

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

func (tcc *ThisChainCode) registrarReconocimientoRaza(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))

	// ---------------------------------------------------------------------------------------------------
	// VALIDAR Argumentos
	// ---------------------------------------------------------------------------------------------------

	if len(args) != 2 {
		return shim.Error("Incorrecto numero de argumentos. Esperando 2")
	}

	DatosPerroComoJson := []byte(args[0])

	var datosPerro cc_cfg.PerroReconocimientoRaza
	err := json.Unmarshal(DatosPerroComoJson, &datosPerro)
	if err != nil {
		fmt.Println(err)
	}

	// ---------------------------------------------------------------------------------------------------

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

	var datosRegistro []cc_cfg.TipoQueryPerros

	err = json.Unmarshal(queryResults, &datosRegistro)
	if err != nil {
		fmt.Println(err)
	}

	// ---------------------------------------------------------------------------------------------------

	IDRaza := datosPerro.IDRaza
	if IDRaza <= 0 {
		return shim.Error("(Args[0]) IDRaza: debe tener un valor")
	}

	queryString = "{\"selector\":{\"docType\":\"" + cc_razas_cfg.CFG_ObjectType + "\",\"IDRaza\":" + strconv.Itoa(IDRaza) + ",\"FechaBaja\":\"\"}}"
	response := stub.InvokeChaincode(cc_razas_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("ejecutarConsulta", queryString), stub.GetChannelID())

	if response.Status != shim.OK {
		return shim.Error(string(response.Payload))
	}

	if string(response.Payload) == "[]" {
		return shim.Error("(Args[0) IDRaza: [ " + strconv.Itoa(IDPerro) + " ] no existe o no es valido")
	}

	// ---------------------------------------------------------------------------------------------------

	Justificacion := datosPerro.Justificacion
	if len(Justificacion) <= 0 {
		return shim.Error("(Args[0]) Justificacion: debe tener un valor")
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

	fechaHoraActual := time.Now()
	fechaHoraActualAsString := fechaHoraActual.String()

	// ---------------------------------------------------------------------------------------------------

	IDPerroReconocimientoRaza := 0

	infoChainCodeAsBytes, err := stub.GetState(cc_cfg.CFG_ObjectType_ReconocimientosRazas)

	if err != nil {
		return shim.Error(err.Error())
	}

	if len(infoChainCodeAsBytes) > 0 {

		var queryInfoChaincode cc_util.InfoChaincode
		err = json.Unmarshal(infoChainCodeAsBytes, &queryInfoChaincode)
		if err != nil {
			fmt.Println(err)
		}
		IDPerroReconocimientoRaza = queryInfoChaincode.IDMaximo
	}

	IDPerroReconocimientoRaza += 1

	// ---------------------------------------------------------------------------------------------------

	InfoChaincodeAsBytes, err := json.Marshal(&cc_util.InfoChaincode{cc_util.CC_INFO_CONTADOR, IDPerroReconocimientoRaza})
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println([]string{cc_cfg.CFG_ObjectType_ReconocimientosRazas, string(InfoChaincodeAsBytes)})

	err = stub.PutState(cc_cfg.CFG_ObjectType_ReconocimientosRazas, InfoChaincodeAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ---------------------------------------------------------------------------------------------------

	nuevoReconocimientoRaza := &cc_cfg.PerroReconocimientoRaza{cc_cfg.CFG_ObjectType_ReconocimientosRazas, IDPerroReconocimientoRaza, IDPerro, IDRaza, Justificacion, IDPersonaEjecuta, fechaHoraActualAsString}

	nuevoReconocimientoRazaAsBytes, err := json.Marshal(nuevoReconocimientoRaza)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(cc_cfg.CFG_ObjectType_ReconocimientosRazas+strconv.Itoa(IDPerroReconocimientoRaza), nuevoReconocimientoRazaAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println(nuevoReconocimientoRaza)

	// ---------------------------------------------------------------------------------------------------

	datosRegistro[0].Record.IDRaza = IDRaza

	datosRegistroAsBytes, err := json.Marshal(datosRegistro[0].Record)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(cc_cfg.CFG_ObjectType+strconv.Itoa(IDPerro), datosRegistroAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println(datosRegistro)

	// ---------------------------------------------------------------------------------------------------

	return shim.Success(datosRegistroAsBytes)

}
func (tcc *ThisChainCode) consultarDatosEjemplar(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))

	retorno := " ********** FUNCION " + cc_util.NombreFuncion() + " SIN IMPLEMENTAR **********"
	fmt.Println(retorno)

	return shim.Success([]byte(retorno))
}

func (tcc *ThisChainCode) obtenerCertificadoRegistro(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))

	retorno := " ********** FUNCION " + cc_util.NombreFuncion() + " SIN IMPLEMENTAR **********"
	fmt.Println(retorno)

	return shim.Success([]byte(retorno))
}

func (tcc *ThisChainCode) obtenerPedigri(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))

	retorno := " ********** FUNCION " + cc_util.NombreFuncion() + " SIN IMPLEMENTAR **********"
	fmt.Println(retorno)

	return shim.Success([]byte(retorno))
}

func (tcc *ThisChainCode) registrarCesionTemporal(stub shim.ChaincodeStubInterface, args []string) pb.Response {
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
