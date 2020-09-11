package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	cc_util "github.com/chaincode/netcan/netcan/lib/netcan_util"
	cc_cfg "github.com/chaincode/netcan/solicitudes/lib/solicitudes_cfg"

	cc_afijos_cfg "github.com/chaincode/netcan/afijos/lib/afijos_cfg"
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

	if function == "solicitarRegistrarCamada" {
		return tcc.solicitarRegistrarPerro(stub, args, function)
	} else if function == "solicitarRegistrarPerro" {
		return tcc.solicitarRegistrarPerro(stub, args, function)
	} else if function == "solicitarRegistrarPerroConCertificado" {
		return tcc.solicitarRegistrarPerro(stub, args, function)
	} else if function == "solicitarRegistrarCambioPropietarioPerro" {
		return tcc.solicitarRegistrarCambioPropietarioPerro(stub, args, function)

	} else if function == "solicitarRegistrarAfijo" {
		return tcc.solicitarRegistrarAfijo(stub, args, function)
	} else if function == "solicitarRegistrarCambioPropietarioAfijo" {
		return tcc.solicitarRegistrarCambioPropietarioAfijo(stub, args, function)
	} else if function == "solicitarRegistrarCancelacionAfijo" {
		return tcc.solicitarRegistrarCancelacionAfijo(stub, args, function)

	} else if function == "cargarDatosIniciales" {
		return tcc.cargarDatosIniciales(stub, args)
	} else if function == "cargarDatosIniciales_Autorizaciones" {
		return tcc.cargarDatosIniciales_Autorizaciones(stub, args)

	} else if function == "querySolicitudes" {
		return tcc.querySolicitudes(stub, args)

	} else if function == "validarSolicitud" {
		return tcc.validarSolicitud(stub, args)

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

	} else {
		return shim.Error("(" + cc_cfg.CFG_ObjectType + ") Invoca un nombre de funcion no valida (" + function + ")")
	}
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (tcc *ThisChainCode) validarSolicitud(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))

	// ---------------------------------------------------------------------------------------------------
	// VALIDAR Argumentos
	// ---------------------------------------------------------------------------------------------------

	var queryString string
	var queryResults []byte
	var err error

	if len(args) != 2 {
		return shim.Error("Incorrecto numero de argumentos. Esperando 2")
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

	ValidadionSolicitudComoJson := []byte(args[0])

	var validacionSolicitud cc_cfg.TipoValidarSolicitud
	err = json.Unmarshal(ValidadionSolicitudComoJson, &validacionSolicitud)
	if err != nil {
		fmt.Println(err)
	}

	IDSolicitud := validacionSolicitud.IDSolicitud
	if IDSolicitud <= 0 {
		return shim.Error("(Args[0]) IDSolicitud: debe tener un valor")
	}

	SolicitudEstado := validacionSolicitud.EstadoSolicitud

	if SolicitudEstado != cc_cfg.SolicitudEstadoAprobado && SolicitudEstado != cc_cfg.SolicitudEstadoRechazado {
		return shim.Error(fmt.Sprintf("(Args[0]) SolicitudEstado: %s no valido", SolicitudEstado))
	}

	queryString = fmt.Sprintf(
		"{\"selector\":{ \"docType\":\"%s\", \"IDSolicitud\":%d, \"IDPersona\":%d}}",
		cc_cfg.CFG_ObjectType_Autorizaciones,
		IDSolicitud,
		IDPersonaEjecuta)

	queryResults, err = getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	if string(queryResults) == "[]" {
		return shim.Error(fmt.Sprintf("(Args[0]) La autorizacion de la solicitud %d para la persona %d no existe", IDSolicitud, IDPersonaEjecuta))
	}

	// ---------------------------------------------------------------------------------------------------

	var querySolicitudesAutorizaciones []cc_cfg.TipoQuerySolicitudesAutorizaciones
	err = json.Unmarshal(queryResults, &querySolicitudesAutorizaciones)
	if err != nil {
		fmt.Println(err)
	}

	if querySolicitudesAutorizaciones[0].Record.EstadoSolicitud != cc_cfg.SolicitudEstadoPendiente {
		return shim.Error(fmt.Sprintf("(Args[0]) La autorizacion de la solicitud %d para la persona %d no esta %s: ", IDSolicitud, IDPersonaEjecuta, cc_cfg.SolicitudEstadoPendiente))
	}

	fechaHoraActual := time.Now()
	fechaHoraActualAsString := fechaHoraActual.String()

	if fechaHoraActualAsString > querySolicitudesAutorizaciones[0].Record.FechaBaja {
		return shim.Error(fmt.Sprintf("(Args[0]) La autorizacion de la solicitud %d para la persona %d ha caducado", IDSolicitud, IDPersonaEjecuta))
	}

	// ---------------------------------------------------------------------------------------------------
	// GRABAR Registros
	// ---------------------------------------------------------------------------------------------------

	querySolicitudesAutorizaciones[0].Record.EstadoSolicitud = SolicitudEstado
	querySolicitudesAutorizaciones[0].Record.FechaEjecucion = fechaHoraActualAsString

	querySolicitudesAutorizacionesAsBytes, err := json.Marshal(querySolicitudesAutorizaciones[0].Record)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(
		cc_cfg.CFG_ObjectType_Autorizaciones+strconv.Itoa(IDSolicitud)+"_"+strconv.Itoa(IDPersonaEjecuta),
		querySolicitudesAutorizacionesAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ---------------------------------------------------------------------------------------------------
	// Ejecutar solicitud (si se puede)
	// ---------------------------------------------------------------------------------------------------

	if SolicitudEstado == cc_cfg.SolicitudEstadoRechazado {

		queryString = fmt.Sprintf(
			"{\"selector\":{ \"docType\":\"%s\", \"IDSolicitud\":%d}}",
			cc_cfg.CFG_ObjectType,
			IDSolicitud)

		queryResults, err = getQueryResultForQueryString(stub, queryString)
		if err != nil {
			return shim.Error(err.Error())
		}

		var querySolicitudes []cc_cfg.TipoQuerySolicitudes

		err = json.Unmarshal(queryResults, &querySolicitudes)
		if err != nil {
			fmt.Println(err)
		}

		querySolicitudes[0].Record.EstadoSolicitud = cc_cfg.SolicitudEstadoRechazado
		querySolicitudes[0].Record.FechaEjecucion = fechaHoraActualAsString

		cambiarSolicitudAsBytes, err := json.Marshal(querySolicitudes[0].Record)
		if err != nil {
			return shim.Error(err.Error())
		}

		err = stub.PutState(cc_cfg.CFG_ObjectType+strconv.Itoa(IDSolicitud), cambiarSolicitudAsBytes)
		if err != nil {
			return shim.Error(err.Error())
		}

	} else if SolicitudEstado == cc_cfg.SolicitudEstadoAprobado {

		queryString = fmt.Sprintf(
			"{\"selector\":{ \"docType\":\"%s\", \"IDSolicitud\":%d}}",
			cc_cfg.CFG_ObjectType_Autorizaciones,
			IDSolicitud)

		queryResults, err = getQueryResultForQueryString(stub, queryString)
		if err != nil {
			return shim.Error(err.Error())
		}

		err = json.Unmarshal(queryResults, &querySolicitudesAutorizaciones)
		if err != nil {
			fmt.Println(err)
		}

		ejecutarSolicitud := true

		for querySolicitudAutorizacion := range querySolicitudesAutorizaciones {

			fmt.Println("---------------------------------------------------------------")
			fmt.Println(querySolicitudesAutorizaciones[querySolicitudAutorizacion].Record)

			if querySolicitudesAutorizaciones[querySolicitudAutorizacion].Record.EstadoSolicitud != cc_cfg.SolicitudEstadoAprobado &&
				!(querySolicitudesAutorizaciones[querySolicitudAutorizacion].Record.IDPersona == IDPersonaEjecuta &&
					querySolicitudesAutorizaciones[querySolicitudAutorizacion].Record.EstadoSolicitud == cc_cfg.SolicitudEstadoPendiente) {
				ejecutarSolicitud = false
			}
		}

		if ejecutarSolicitud {

			queryString = fmt.Sprintf(
				"{\"selector\":{ \"docType\":\"%s\", \"IDSolicitud\":%d}}",
				cc_cfg.CFG_ObjectType,
				IDSolicitud)

			queryResults, err = getQueryResultForQueryString(stub, queryString)
			if err != nil {
				return shim.Error(err.Error())
			}

			var querySolicitudes []cc_cfg.TipoQuerySolicitudes
			var response pb.Response

			err = json.Unmarshal(queryResults, &querySolicitudes)
			if err != nil {
				fmt.Println(err)
			}

			querySolicitudes[0].Record.EstadoSolicitud = cc_cfg.SolicitudEstadoAprobado
			querySolicitudes[0].Record.FechaEjecucion = fechaHoraActualAsString

			cambiarSolicitudAsBytes, err := json.Marshal(querySolicitudes[0].Record)
			if err != nil {
				return shim.Error(err.Error())
			}

			err = stub.PutState(cc_cfg.CFG_ObjectType+strconv.Itoa(IDSolicitud), cambiarSolicitudAsBytes)
			if err != nil {
				return shim.Error(err.Error())
			}

			/*
				fmt.Println("****************************************")
				fmt.Println(strconv.Itoa(IDSolicitud))
				fmt.Println(string(cambiarSolicitudAsBytes))
				fmt.Println("****************************************")
			*/

			TipoSolicitud := querySolicitudes[0].Record.TipoSolicitud
			JSONSolicitud := querySolicitudes[0].Record.JSONSolicitud

			if TipoSolicitud == "solicitarRegistrarCamada" {
				response = stub.InvokeChaincode(cc_perros_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("registrarPerro", JSONSolicitud, args[1]), stub.GetChannelID())

			} else if TipoSolicitud == "solicitarRegistrarPerro" {
				response = stub.InvokeChaincode(cc_perros_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("registrarPerro", JSONSolicitud, args[1]), stub.GetChannelID())

			} else if TipoSolicitud == "solicitarRegistrarPerroConCertificado" {
				response = stub.InvokeChaincode(cc_perros_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("registrarPerro", JSONSolicitud, args[1]), stub.GetChannelID())

			} else if TipoSolicitud == "solicitarRegistrarCambioPropietarioPerro" {
				response = stub.InvokeChaincode(cc_perros_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("registrarCambioPropietario", JSONSolicitud, args[1]), stub.GetChannelID())

			} else if TipoSolicitud == "solicitarRegistrarAfijo" {
				response = stub.InvokeChaincode(cc_afijos_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("registrarAfijo", JSONSolicitud, args[1]), stub.GetChannelID())
			} else if TipoSolicitud == "solicitarRegistrarCambioPropietarioAfijo" {
				response = stub.InvokeChaincode(cc_afijos_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("registrarCambioPropietario", JSONSolicitud, args[1]), stub.GetChannelID())
			} else if TipoSolicitud == "solicitarRegistrarCancelacionAfijo" {
				response = stub.InvokeChaincode(cc_afijos_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("registrarCancelacionAfijo", JSONSolicitud, args[1]), stub.GetChannelID())

			} else {
				return shim.Error("Invalido nombre de tipo de registro (" + TipoSolicitud + ")")
			}

			if response.Status != shim.OK {
				return shim.Error(response.Message)
			} else {
				return shim.Success(response.Payload)
			}

		}
	}
	return shim.Success(querySolicitudesAutorizacionesAsBytes)
}

func (tcc *ThisChainCode) cargarDatosIniciales(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))

	nombreArchivo := args[0]

	nuevosRegistrosAsJson, err := ioutil.ReadFile(nombreArchivo)
	if err != nil {
		return shim.Error(err.Error())
	}

	var nuevosRegistros []cc_cfg.Solicitudes
	json.Unmarshal(nuevosRegistrosAsJson, &nuevosRegistros)

	for _, nuevoRegistro := range nuevosRegistros {

		fmt.Println(nuevoRegistro)

		nuevoRegistroAsBytes, err := json.Marshal(nuevoRegistro)
		if err != nil {
			return shim.Error(err.Error())
		}

		err = stub.PutState(cc_cfg.CFG_ObjectType+strconv.Itoa(nuevoRegistro.IDSolicitud), nuevoRegistroAsBytes)
		if err != nil {
			return shim.Error(err.Error())

		}
	}

	return shim.Success([]byte(strconv.Itoa(len(nuevosRegistros))))
}

func (tcc *ThisChainCode) cargarDatosIniciales_Autorizaciones(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))

	nombreArchivo := args[0]

	nuevosRegistrosAsJson, err := ioutil.ReadFile(nombreArchivo)
	if err != nil {
		return shim.Error(err.Error())
	}

	var nuevosRegistros []cc_cfg.SolicitudesAutorizaciones
	json.Unmarshal(nuevosRegistrosAsJson, &nuevosRegistros)

	for _, nuevoRegistro := range nuevosRegistros {

		fmt.Println(nuevoRegistro)

		nuevoRegistroAsBytes, err := json.Marshal(nuevoRegistro)
		if err != nil {
			return shim.Error(err.Error())
		}

		err = stub.PutState(cc_cfg.CFG_ObjectType_Autorizaciones+strconv.Itoa(nuevoRegistro.IDSolicitud)+"_"+strconv.Itoa(nuevoRegistro.IDPersona), nuevoRegistroAsBytes)
		if err != nil {
			return shim.Error(err.Error())

		}
	}

	return shim.Success([]byte(strconv.Itoa(len(nuevosRegistros))))
}

func (tcc *ThisChainCode) solicitarRegistrarPerro(stub shim.ChaincodeStubInterface, args []string, tipoRegistro string) pb.Response {

	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))

	// ---------------------------------------------------------------------------------------------------
	// VALIDAR Argumentos
	// ---------------------------------------------------------------------------------------------------

	if len(args) != 2 {
		return shim.Error("Incorrecto numero de argumentos. Esperando 2")
	}

	if len(tipoRegistro) <= 0 {
		return shim.Error("Incorrecto numero de argumentos. Esperando valor de tipoRegistro")
	}

	// ---------------------------------------------------------------------------------------------------

	DatosSeguridadComoJson := []byte(args[1])

	var datosSeguridad cc_util.TipoSeguridad

	err := json.Unmarshal(DatosSeguridadComoJson, &datosSeguridad)
	if err != nil {
		fmt.Println(err)
	}

	IDPersonaEjecuta := datosSeguridad.IDPersona

	queryString := "{\"selector\":{\"docType\":\"" + cc_personas_cfg.CFG_ObjectType + "\",\"IDPersona\":" + strconv.Itoa(IDPersonaEjecuta) + "}}"
	response := stub.InvokeChaincode(cc_personas_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("ejecutarConsulta", queryString), stub.GetChannelID())

	if response.Status != shim.OK {
		return shim.Error(response.Message)
	}

	if string(response.Payload) == "[]" {
		return shim.Error("(Args[1]) IDPersona: [ " + strconv.Itoa(IDPersonaEjecuta) + " ] no existe o no es valido")
	}

	// ---------------------------------------------------------------------------------------------------

	DatosSolicitudComoJson := []byte(args[0])
	var datosSolicitud cc_perros_cfg.TipoRegistrarCamada
	err = json.Unmarshal(DatosSolicitudComoJson, &datosSolicitud)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(datosSolicitud)
	fmt.Println(args[0])

	PerrosCamada := datosSolicitud.Perros
	if len(PerrosCamada) <= 0 {
		return shim.Error("(Args[0]) PerrosCamada: debe tener un perro")
	}

	IDPerroMadre := datosSolicitud.IDPerroMadre

	var Propietarios []cc_perros_cfg.TipoRegistrarCamadaPropietario
	propietariosArray := []int{}

	if IDPerroMadre == 0 {
		Propietarios = datosSolicitud.Propietarios
		if len(Propietarios) <= 0 {
			return shim.Error("(Args[0]) Propietarios: debe tener un propietario")
		}
		for propietario := range Propietarios {

			IDPropietarioNuevo := Propietarios[propietario].IDPersona

			queryString := "{\"selector\":{\"docType\":\"" + cc_personas_cfg.CFG_ObjectType + "\",\"IDPersona\":" + strconv.Itoa(IDPropietarioNuevo) + "}}"
			response := stub.InvokeChaincode(cc_personas_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("ejecutarConsulta", queryString), stub.GetChannelID())

			if response.Status != shim.OK {
				return shim.Error(response.Message)
			}

			if string(response.Payload) == "[]" {
				return shim.Error("(Args[0]) IDPersona: [ " + strconv.Itoa(IDPropietarioNuevo) + " ] no existe o no es valido")
			}

		}
	} else {

		if len(Propietarios) > 0 {
			return shim.Error("(Args[0]) Propietarios: no debe tener definido un propietario")
		}

		queryString = "{\"selector\":{\"docType\":\"" + cc_perros_cfg.CFG_ObjectType_Propietarios + "\",\"IDPerro\":" + strconv.Itoa(IDPerroMadre) + ",\"FechaBaja\":\"\"}}"
		response = stub.InvokeChaincode(cc_perros_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("ejecutarConsulta", queryString), stub.GetChannelID())

		if response.Status != shim.OK {
			return shim.Error(response.Message)
		}

		if string(response.Payload) == "[]" {
			return shim.Error("(Args[0]) IDPerroMadre: [ " + strconv.Itoa(IDPerroMadre) + " ] no tiene propietarios")
		}

		var queryPerrosPropietarios []cc_perros_cfg.TipoQueryPerrosPropietarios

		err = json.Unmarshal(response.Payload, &queryPerrosPropietarios)
		if err != nil {
			return shim.Error(err.Error())
		}

		for propietario := range queryPerrosPropietarios {

			if queryPerrosPropietarios[propietario].Record.IDPersona > 0 {
				propietariosArray = append(propietariosArray, queryPerrosPropietarios[propietario].Record.IDPersona)
			}
		}

	}

	IDPerroPadre := datosSolicitud.IDPerroPadre

	if IDPerroPadre > 0 {

		queryString = "{\"selector\":{\"docType\":\"" + cc_perros_cfg.CFG_ObjectType_Propietarios + "\",\"IDPerro\":" + strconv.Itoa(IDPerroPadre) + "}}"
		response := stub.InvokeChaincode(cc_perros_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("ejecutarConsulta", queryString), stub.GetChannelID())

		if response.Status != shim.OK {
			return shim.Error(response.Message)
		}

		if string(response.Payload) == "[]" {
			return shim.Error("(Args[0]) IDPerroPadre: [ " + strconv.Itoa(IDPerroPadre) + " ] no tiene propietarios")
		}

		var queryPerrosPropietarios []cc_perros_cfg.TipoQueryPerrosPropietarios

		err = json.Unmarshal(response.Payload, &queryPerrosPropietarios)
		if err != nil {
			return shim.Error(err.Error())
		}

		for propietario := range queryPerrosPropietarios {

			if queryPerrosPropietarios[propietario].Record.IDPersona > 0 {
				propietariosArray = append(propietariosArray, queryPerrosPropietarios[propietario].Record.IDPersona)
			}
		}

	}

	// ---------------------------------------------------------------------------------------------------
	// GRABAR Registros
	// ---------------------------------------------------------------------------------------------------

	var retorno bytes.Buffer

	fechaHoraActual := time.Now()
	fechaHoraActualAsString := fechaHoraActual.String()
	fechaLimite := fechaHoraActual.Add(3 * 24 * time.Hour)
	fechaLimiteAsString := fechaLimite.String()

	// ---------------------------------------------------------------------------------------------------

	IDSolicitud := 0

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
		IDSolicitud = queryInfoChaincode.IDMaximo
	}

	IDSolicitud += 1

	// ---------------------------------------------------------------------------------------------------

	InfoChaincodeAsBytes, err := json.Marshal(&cc_util.InfoChaincode{cc_util.CC_INFO_CONTADOR, IDSolicitud})
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println([]string{cc_cfg.CFG_ObjectType, string(InfoChaincodeAsBytes)})

	err = stub.PutState(cc_cfg.CFG_ObjectType, InfoChaincodeAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ---------------------------------------------------------------------------------------------------

	var SolicitudEstado string
	var FechaHoraEjecucionAsString string
	var ejecutarSolicitud bool

	propietariosArray = cc_util.ArrayIntSinDuplicados(propietariosArray)

	if len(propietariosArray) > 0 {
		for propietario := range propietariosArray {
			if IDPersonaEjecuta == propietariosArray[propietario] && len(propietariosArray) == 1 {
				ejecutarSolicitud = true
			}
		}
	} else {
		ejecutarSolicitud = true
	}

	if ejecutarSolicitud {
		SolicitudEstado = cc_cfg.SolicitudEstadoAprobado
		FechaHoraEjecucionAsString = fechaHoraActualAsString
	} else {
		SolicitudEstado = cc_cfg.SolicitudEstadoPendiente
		FechaHoraEjecucionAsString = ""
	}

	nuevoSolicitud := &cc_cfg.Solicitudes{cc_cfg.CFG_ObjectType, IDSolicitud, tipoRegistro, args[0], IDPersonaEjecuta, SolicitudEstado, FechaHoraEjecucionAsString, fechaHoraActualAsString, fechaLimiteAsString}
	fmt.Println(nuevoSolicitud)

	nuevoSolicitudAsBytes, err := json.Marshal(nuevoSolicitud)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(cc_cfg.CFG_ObjectType+strconv.Itoa(IDSolicitud), nuevoSolicitudAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	retorno.WriteString(string(nuevoSolicitudAsBytes))

	// ---------------------------------------------------------------------------------------------------

	for propietario := range propietariosArray {

		if IDPersonaEjecuta == propietariosArray[propietario] {
			SolicitudEstado = cc_cfg.SolicitudEstadoAprobado
			FechaHoraEjecucionAsString = fechaHoraActualAsString
		} else {
			SolicitudEstado = cc_cfg.SolicitudEstadoPendiente
			FechaHoraEjecucionAsString = ""
		}

		nuevaSolicitudAutorizacion := &cc_cfg.SolicitudesAutorizaciones{
			cc_cfg.CFG_ObjectType_Autorizaciones,
			IDSolicitud,
			propietariosArray[propietario],
			SolicitudEstado,
			FechaHoraEjecucionAsString,
			fechaHoraActualAsString,
			fechaLimiteAsString}
		// fmt.Println(nuevaSolicitudAutorizacion)

		nuevaSolicitudAutorizacionAsBytes, err := json.Marshal(nuevaSolicitudAutorizacion)
		if err != nil {
			return shim.Error(err.Error())
		}

		err = stub.PutState(cc_cfg.CFG_ObjectType_Autorizaciones+strconv.Itoa(IDSolicitud)+"_"+strconv.Itoa(propietariosArray[propietario]), nuevaSolicitudAutorizacionAsBytes)
		if err != nil {
			return shim.Error(err.Error())
		}

		retorno.WriteString(string(nuevaSolicitudAutorizacionAsBytes))
	}

	if ejecutarSolicitud {

		fmt.Println(ejecutarSolicitud)
		response := stub.InvokeChaincode(cc_perros_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("registrarPerro", args[0], args[1]), stub.GetChannelID())
		if response.Status != shim.OK {
			return shim.Error(response.Message)
		} else {
			return shim.Success(response.Payload)
		}

	}

	return shim.Success(retorno.Bytes())
}

func (tcc *ThisChainCode) solicitarRegistrarPerroConCertificado(stub shim.ChaincodeStubInterface, args []string, tipoRegistro string) pb.Response {
	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))

	retorno := " ********** FUNCION " + cc_util.NombreFuncion() + " SIN IMPLEMENTAR **********"
	fmt.Println(retorno)

	return shim.Success([]byte(retorno))
}

func (tcc *ThisChainCode) solicitarRegistrarCambioPropietarioPerro(stub shim.ChaincodeStubInterface, args []string, tipoRegistro string) pb.Response {

	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))

	// ---------------------------------------------------------------------------------------------------
	// VALIDAR Argumentos
	// ---------------------------------------------------------------------------------------------------

	if len(args) != 2 {
		return shim.Error("Incorrecto numero de argumentos. Esperando 2")
	}

	if len(tipoRegistro) <= 0 {
		return shim.Error("Incorrecto numero de argumentos. Esperando valor de tipoRegistro")
	}

	// ---------------------------------------------------------------------------------------------------

	DatosSeguridadComoJson := []byte(args[1])

	var datosSeguridad cc_util.TipoSeguridad

	err := json.Unmarshal(DatosSeguridadComoJson, &datosSeguridad)
	if err != nil {
		fmt.Println(err)
	}

	IDPersonaEjecuta := datosSeguridad.IDPersona

	queryString := "{\"selector\":{\"docType\":\"" + cc_personas_cfg.CFG_ObjectType + "\",\"IDPersona\":" + strconv.Itoa(IDPersonaEjecuta) + "}}"
	response := stub.InvokeChaincode(cc_personas_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("ejecutarConsulta", queryString), stub.GetChannelID())

	if response.Status != shim.OK {
		return shim.Error(response.Message)
	}

	if string(response.Payload) == "[]" {
		return shim.Error("(Args[1]) IDPersona: [ " + strconv.Itoa(IDPersonaEjecuta) + " ] no existe o no es valido")
	}

	// ---------------------------------------------------------------------------------------------------

	DatosSolicitudComoJson := []byte(args[0])
	var datosSolicitud cc_perros_cfg.TipoRegistrarCambioPropietario
	err = json.Unmarshal(DatosSolicitudComoJson, &datosSolicitud)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(datosSolicitud)
	fmt.Println(args[0])

	IDPerro := datosSolicitud.IDPerro
	if IDPerro <= 0 {
		return shim.Error("(Args[0]) IDPerro: debe tener un perro")
	}

	queryString = "{\"selector\":{\"docType\":\"" + cc_perros_cfg.CFG_ObjectType_Propietarios + "\",\"IDPerro\":" + strconv.Itoa(IDPerro) + ",\"FechaBaja\":\"\"}}"
	response = stub.InvokeChaincode(cc_perros_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("ejecutarConsulta", queryString), stub.GetChannelID())

	if response.Status != shim.OK {
		return shim.Error(response.Message)
	}

	if string(response.Payload) == "[]" {
		return shim.Error("(Args[0]) IDPerro: [ " + strconv.Itoa(IDPerro) + " ] no tiene propietarios")
	}

	var queryPerrosPropietarios []cc_perros_cfg.TipoQueryPerrosPropietarios

	propietariosArray := []int{}
	err = json.Unmarshal(response.Payload, &queryPerrosPropietarios)
	if err != nil {
		return shim.Error(err.Error())
	}

	for propietario := range queryPerrosPropietarios {

		if queryPerrosPropietarios[propietario].Record.IDPersona > 0 {
			propietariosArray = append(propietariosArray, queryPerrosPropietarios[propietario].Record.IDPersona)
		}
	}

	// ---------------------------------------------------------------------------------------------------

	var PropietariosNuevos []cc_perros_cfg.TipoRegistrarCambioPropietarioPropietario

	PropietariosNuevos = datosSolicitud.Propietarios
	if len(PropietariosNuevos) <= 0 {
		return shim.Error("(Args[0]) Propietarios: debe tener un propietario")
	}
	for propietario := range PropietariosNuevos {

		IDPropietarioNuevo := PropietariosNuevos[propietario].IDPersona

		queryString := "{\"selector\":{\"docType\":\"" + cc_personas_cfg.CFG_ObjectType + "\",\"IDPersona\":" + strconv.Itoa(IDPropietarioNuevo) + "}}"
		response := stub.InvokeChaincode(cc_personas_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("ejecutarConsulta", queryString), stub.GetChannelID())

		if response.Status != shim.OK {
			return shim.Error(response.Message)
		}

		if string(response.Payload) == "[]" {
			return shim.Error("(Args[1]) IDPersona: [ " + strconv.Itoa(IDPropietarioNuevo) + " ] no existe o no es valido")
		}

	}

	// ---------------------------------------------------------------------------------------------------
	// GRABAR Registros
	// ---------------------------------------------------------------------------------------------------

	var retorno bytes.Buffer

	fechaHoraActual := time.Now()
	fechaHoraActualAsString := fechaHoraActual.String()
	fechaLimite := fechaHoraActual.Add(3 * 24 * time.Hour)
	fechaLimiteAsString := fechaLimite.String()

	// ---------------------------------------------------------------------------------------------------

	IDSolicitud := 0

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
		IDSolicitud = queryInfoChaincode.IDMaximo
	}

	IDSolicitud += 1

	// ---------------------------------------------------------------------------------------------------

	InfoChaincodeAsBytes, err := json.Marshal(&cc_util.InfoChaincode{cc_util.CC_INFO_CONTADOR, IDSolicitud})
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println([]string{cc_cfg.CFG_ObjectType, string(InfoChaincodeAsBytes)})

	err = stub.PutState(cc_cfg.CFG_ObjectType, InfoChaincodeAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ---------------------------------------------------------------------------------------------------

	var SolicitudEstado string
	var FechaHoraEjecucionAsString string
	var ejecutarSolicitud bool

	propietariosArray = cc_util.ArrayIntSinDuplicados(propietariosArray)

	for propietario := range propietariosArray {

		if IDPersonaEjecuta == propietariosArray[propietario] {
			SolicitudEstado = cc_cfg.SolicitudEstadoAprobado
			FechaHoraEjecucionAsString = fechaHoraActualAsString
			if len(propietariosArray) == 1 {
				ejecutarSolicitud = true
			}

		} else {
			SolicitudEstado = cc_cfg.SolicitudEstadoPendiente
			FechaHoraEjecucionAsString = ""
		}

	}

	nuevoSolicitud := &cc_cfg.Solicitudes{cc_cfg.CFG_ObjectType, IDSolicitud, tipoRegistro, args[0], IDPersonaEjecuta, SolicitudEstado, FechaHoraEjecucionAsString, fechaHoraActualAsString, fechaLimiteAsString}
	fmt.Println(nuevoSolicitud)

	nuevoSolicitudAsBytes, err := json.Marshal(nuevoSolicitud)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(cc_cfg.CFG_ObjectType+strconv.Itoa(IDSolicitud), nuevoSolicitudAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	retorno.WriteString(string(nuevoSolicitudAsBytes))

	// ---------------------------------------------------------------------------------------------------

	propietariosArray = cc_util.ArrayIntSinDuplicados(propietariosArray)

	for propietario := range propietariosArray {

		if IDPersonaEjecuta == propietariosArray[propietario] {
			SolicitudEstado = cc_cfg.SolicitudEstadoAprobado
			FechaHoraEjecucionAsString = fechaHoraActualAsString
		} else {
			SolicitudEstado = cc_cfg.SolicitudEstadoPendiente
			FechaHoraEjecucionAsString = ""
		}

		nuevaSolicitudAutorizacion := &cc_cfg.SolicitudesAutorizaciones{
			cc_cfg.CFG_ObjectType_Autorizaciones,
			IDSolicitud,
			propietariosArray[propietario],
			SolicitudEstado,
			FechaHoraEjecucionAsString,
			fechaHoraActualAsString,
			fechaLimiteAsString}
		// fmt.Println(nuevaSolicitudAutorizacion)

		nuevaSolicitudAutorizacionAsBytes, err := json.Marshal(nuevaSolicitudAutorizacion)
		if err != nil {
			return shim.Error(err.Error())
		}

		err = stub.PutState(cc_cfg.CFG_ObjectType_Autorizaciones+strconv.Itoa(IDSolicitud)+"_"+strconv.Itoa(propietariosArray[propietario]), nuevaSolicitudAutorizacionAsBytes)
		if err != nil {
			return shim.Error(err.Error())
		}

		retorno.WriteString(string(nuevaSolicitudAutorizacionAsBytes))
	}

	if ejecutarSolicitud {

		fmt.Println(ejecutarSolicitud)
		response := stub.InvokeChaincode(cc_perros_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("registrarCambioPropietario", args[0], args[1]), stub.GetChannelID())
		if response.Status != shim.OK {
			return shim.Error(response.Message)
		} else {
			return shim.Success(response.Payload)
		}
	}

	return shim.Success(retorno.Bytes())
}

func (tcc *ThisChainCode) querySolicitudes(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	type tipoFiltroQuerySolicitudes struct {
		IDPersona       int    `json:"IDPersona"`
		EstadoSolicitud string `json:"EstadoSolicitud"`
	}

	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))

	// ---------------------------------------------------------------------------------------------------
	// VALIDAR Argumentos
	// ---------------------------------------------------------------------------------------------------

	var queryString string
	var queryResults []byte
	var err error

	if len(args) != 2 {
		return shim.Error("Incorrecto numero de argumentos. Esperando 2")
	}

	FiltroQuerySolicitudesComoJson := []byte(args[0])

	var FiltroQuerySolicitudes tipoFiltroQuerySolicitudes

	err = json.Unmarshal(FiltroQuerySolicitudesComoJson, &FiltroQuerySolicitudes)
	if err != nil {
		fmt.Println(err)
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

	// ---------------------------------------------------------------------------------------------------
	// REALIZAR consulta
	// ---------------------------------------------------------------------------------------------------

	fmt.Println(FiltroQuerySolicitudes.IDPersona)

	filtroSolicitudesAutorizacionesIDPersona := ""
	if FiltroQuerySolicitudes.IDPersona > 0 {
		filtroSolicitudesAutorizacionesIDPersona = fmt.Sprintf(",\"IDPersona\":%d", FiltroQuerySolicitudes.IDPersona)
	}

	filtroSolicitudesAutorizacionesEstadoSolicitud := ""
	if len(FiltroQuerySolicitudes.EstadoSolicitud) > 0 {
		filtroSolicitudesAutorizacionesEstadoSolicitud = fmt.Sprintf(",\"EstadoSolicitud\":\"%s\"", FiltroQuerySolicitudes.EstadoSolicitud)
	}

	queryString = fmt.Sprintf(
		"{\"selector\":{\"docType\":\"%s\"%s %s}}",
		cc_cfg.CFG_ObjectType_Autorizaciones,
		filtroSolicitudesAutorizacionesIDPersona,
		filtroSolicitudesAutorizacionesEstadoSolicitud)

	fmt.Println(queryString)

	queryResults, err = getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (tcc *ThisChainCode) solicitarRegistrarAfijo(stub shim.ChaincodeStubInterface, args []string, tipoRegistro string) pb.Response {
	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))

	retorno := " ********** FUNCION " + cc_util.NombreFuncion() + " SIN IMPLEMENTAR ********** usar por el momento registrarAfijo de AFIJOS"
	fmt.Println(retorno)

	return shim.Success([]byte(retorno))
}

func (tcc *ThisChainCode) solicitarRegistrarCambioPropietarioAfijo(stub shim.ChaincodeStubInterface, args []string, tipoRegistro string) pb.Response {

	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))

	// ---------------------------------------------------------------------------------------------------
	// VALIDAR Argumentos
	// ---------------------------------------------------------------------------------------------------

	if len(args) != 2 {
		return shim.Error("Incorrecto numero de argumentos. Esperando 2")
	}

	if len(tipoRegistro) <= 0 {
		return shim.Error("Incorrecto numero de argumentos. Esperando valor de tipoRegistro")
	}

	// ---------------------------------------------------------------------------------------------------

	DatosSeguridadComoJson := []byte(args[1])

	var datosSeguridad cc_util.TipoSeguridad

	err := json.Unmarshal(DatosSeguridadComoJson, &datosSeguridad)
	if err != nil {
		fmt.Println(err)
	}

	IDPersonaEjecuta := datosSeguridad.IDPersona

	queryString := "{\"selector\":{\"docType\":\"" + cc_personas_cfg.CFG_ObjectType + "\",\"IDPersona\":" + strconv.Itoa(IDPersonaEjecuta) + "}}"
	response := stub.InvokeChaincode(cc_personas_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("ejecutarConsulta", queryString), stub.GetChannelID())

	if response.Status != shim.OK {
		return shim.Error(response.Message)
	}

	if string(response.Payload) == "[]" {
		return shim.Error("(Args[1]) IDPersona: [ " + strconv.Itoa(IDPersonaEjecuta) + " ] no existe o no es valido")
	}

	// ---------------------------------------------------------------------------------------------------

	DatosSolicitudComoJson := []byte(args[0])
	var datosSolicitud cc_afijos_cfg.TipoRegistrarCambioPropietario
	err = json.Unmarshal(DatosSolicitudComoJson, &datosSolicitud)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(datosSolicitud)
	fmt.Println(args[0])

	IDAfijo := datosSolicitud.IDAfijo
	if IDAfijo <= 0 {
		return shim.Error("(Args[0]) IDAfijo: debe tener un Afijo")
	}

	queryString = "{\"selector\":{\"docType\":\"" + cc_afijos_cfg.CFG_ObjectType_Propietarios + "\",\"IDAfijo\":" + strconv.Itoa(IDAfijo) + ",\"FechaBaja\":\"\"}}"
	response = stub.InvokeChaincode(cc_afijos_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("ejecutarConsulta", queryString), stub.GetChannelID())

	if response.Status != shim.OK {
		return shim.Error(response.Message)
	}

	if string(response.Payload) == "[]" {
		return shim.Error("(Args[0]) IDAfijo: [ " + strconv.Itoa(IDAfijo) + " ] no tiene propietarios")
	}

	var queryAfijosPropietarios []cc_afijos_cfg.TipoQueryAfijosPropietarios

	propietariosArray := []int{}
	err = json.Unmarshal(response.Payload, &queryAfijosPropietarios)
	if err != nil {
		return shim.Error(err.Error())
	}

	for propietario := range queryAfijosPropietarios {

		if queryAfijosPropietarios[propietario].Record.IDPersona > 0 {
			propietariosArray = append(propietariosArray, queryAfijosPropietarios[propietario].Record.IDPersona)
		}
	}

	// ---------------------------------------------------------------------------------------------------

	var PropietariosNuevos []cc_afijos_cfg.TipoRegistrarCambioPropietarioPropietario

	PropietariosNuevos = datosSolicitud.Propietarios
	if len(PropietariosNuevos) <= 0 {
		return shim.Error("(Args[0]) Propietarios: debe tener un propietario")
	}
	for propietario := range PropietariosNuevos {

		IDPropietarioNuevo := PropietariosNuevos[propietario].IDPersona

		queryString := "{\"selector\":{\"docType\":\"" + cc_personas_cfg.CFG_ObjectType + "\",\"IDPersona\":" + strconv.Itoa(IDPropietarioNuevo) + "}}"
		response := stub.InvokeChaincode(cc_personas_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("ejecutarConsulta", queryString), stub.GetChannelID())

		if response.Status != shim.OK {
			return shim.Error(response.Message)
		}

		if string(response.Payload) == "[]" {
			return shim.Error("(Args[1]) IDPersona: [ " + strconv.Itoa(IDPropietarioNuevo) + " ] no existe o no es valido")
		}

		// ---------------------------------------------------------------------------------------------------

		queryString = "{\"selector\":{\"docType\":\"" + cc_afijos_cfg.CFG_ObjectType_Propietarios + "\",\"IDPersona\":" + strconv.Itoa(IDPropietarioNuevo) + ",\"FechaBaja\":\"\"}}"
		response = stub.InvokeChaincode(cc_afijos_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("ejecutarConsulta", queryString), stub.GetChannelID())

		if response.Status != shim.OK {
			return shim.Error(response.Message)
		}

		if string(response.Payload) != "[]" {
			return shim.Error("(Args[0, " + strconv.Itoa(propietario) + "]) IDPersona: [ " + strconv.Itoa(IDPropietarioNuevo) + " ] tiene asignado un afijo activo " + string(response.Payload))
		}

	}

	// ---------------------------------------------------------------------------------------------------
	// GRABAR Registros
	// ---------------------------------------------------------------------------------------------------

	var retorno bytes.Buffer

	fechaHoraActual := time.Now()
	fechaHoraActualAsString := fechaHoraActual.String()
	fechaLimite := fechaHoraActual.Add(3 * 24 * time.Hour)
	fechaLimiteAsString := fechaLimite.String()

	// ---------------------------------------------------------------------------------------------------

	IDSolicitud := 0

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
		IDSolicitud = queryInfoChaincode.IDMaximo
	}

	IDSolicitud += 1

	// ---------------------------------------------------------------------------------------------------

	InfoChaincodeAsBytes, err := json.Marshal(&cc_util.InfoChaincode{cc_util.CC_INFO_CONTADOR, IDSolicitud})
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println([]string{cc_cfg.CFG_ObjectType, string(InfoChaincodeAsBytes)})

	err = stub.PutState(cc_cfg.CFG_ObjectType, InfoChaincodeAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ---------------------------------------------------------------------------------------------------

	var SolicitudEstado string
	var FechaHoraEjecucionAsString string
	var ejecutarSolicitud bool

	propietariosArray = cc_util.ArrayIntSinDuplicados(propietariosArray)

	for propietario := range propietariosArray {

		if IDPersonaEjecuta == propietariosArray[propietario] {
			SolicitudEstado = cc_cfg.SolicitudEstadoAprobado
			FechaHoraEjecucionAsString = fechaHoraActualAsString
			if len(propietariosArray) == 1 {
				ejecutarSolicitud = true
			}

		} else {
			SolicitudEstado = cc_cfg.SolicitudEstadoPendiente
			FechaHoraEjecucionAsString = ""
		}

	}

	nuevoSolicitud := &cc_cfg.Solicitudes{cc_cfg.CFG_ObjectType, IDSolicitud, tipoRegistro, args[0], IDPersonaEjecuta, SolicitudEstado, FechaHoraEjecucionAsString, fechaHoraActualAsString, fechaLimiteAsString}
	fmt.Println(nuevoSolicitud)

	nuevoSolicitudAsBytes, err := json.Marshal(nuevoSolicitud)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(cc_cfg.CFG_ObjectType+strconv.Itoa(IDSolicitud), nuevoSolicitudAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	retorno.WriteString(string(nuevoSolicitudAsBytes))

	// ---------------------------------------------------------------------------------------------------

	propietariosArray = cc_util.ArrayIntSinDuplicados(propietariosArray)

	for propietario := range propietariosArray {

		if IDPersonaEjecuta == propietariosArray[propietario] {
			SolicitudEstado = cc_cfg.SolicitudEstadoAprobado
			FechaHoraEjecucionAsString = fechaHoraActualAsString
		} else {
			SolicitudEstado = cc_cfg.SolicitudEstadoPendiente
			FechaHoraEjecucionAsString = ""
		}

		nuevaSolicitudAutorizacion := &cc_cfg.SolicitudesAutorizaciones{
			cc_cfg.CFG_ObjectType_Autorizaciones,
			IDSolicitud,
			propietariosArray[propietario],
			SolicitudEstado,
			FechaHoraEjecucionAsString,
			fechaHoraActualAsString,
			fechaLimiteAsString}
		// fmt.Println(nuevaSolicitudAutorizacion)

		nuevaSolicitudAutorizacionAsBytes, err := json.Marshal(nuevaSolicitudAutorizacion)
		if err != nil {
			return shim.Error(err.Error())
		}

		err = stub.PutState(cc_cfg.CFG_ObjectType_Autorizaciones+strconv.Itoa(IDSolicitud)+"_"+strconv.Itoa(propietariosArray[propietario]), nuevaSolicitudAutorizacionAsBytes)
		if err != nil {
			return shim.Error(err.Error())
		}

		retorno.WriteString(string(nuevaSolicitudAutorizacionAsBytes))
	}

	if ejecutarSolicitud {

		fmt.Println(ejecutarSolicitud)
		response := stub.InvokeChaincode(cc_afijos_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("registrarCambioPropietario", args[0], args[1]), stub.GetChannelID())
		if response.Status != shim.OK {
			return shim.Error(response.Message)
		} else {
			return shim.Success(response.Payload)
		}
	}

	return shim.Success(retorno.Bytes())
}

func (tcc *ThisChainCode) solicitarRegistrarCancelacionAfijo(stub shim.ChaincodeStubInterface, args []string, tipoRegistro string) pb.Response {

	fmt.Println(fmt.Sprintf(" - %s --- %s()", cc_cfg.CFG_ChainCodeName, cc_util.NombreFuncion()))

	// ---------------------------------------------------------------------------------------------------
	// VALIDAR Argumentos
	// ---------------------------------------------------------------------------------------------------

	if len(args) != 2 {
		return shim.Error("Incorrecto numero de argumentos. Esperando 2")
	}

	if len(tipoRegistro) <= 0 {
		return shim.Error("Incorrecto numero de argumentos. Esperando valor de tipoRegistro")
	}

	// ---------------------------------------------------------------------------------------------------

	DatosSeguridadComoJson := []byte(args[1])

	var datosSeguridad cc_util.TipoSeguridad

	err := json.Unmarshal(DatosSeguridadComoJson, &datosSeguridad)
	if err != nil {
		fmt.Println(err)
	}

	IDPersonaEjecuta := datosSeguridad.IDPersona

	queryString := "{\"selector\":{\"docType\":\"" + cc_personas_cfg.CFG_ObjectType + "\",\"IDPersona\":" + strconv.Itoa(IDPersonaEjecuta) + "}}"
	response := stub.InvokeChaincode(cc_personas_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("ejecutarConsulta", queryString), stub.GetChannelID())

	if response.Status != shim.OK {
		return shim.Error(response.Message)
	}

	if string(response.Payload) == "[]" {
		return shim.Error("(Args[1]) IDPersona: [ " + strconv.Itoa(IDPersonaEjecuta) + " ] no existe o no es valido")
	}

	// ---------------------------------------------------------------------------------------------------

	DatosSolicitudComoJson := []byte(args[0])
	var datosSolicitud cc_afijos_cfg.TipoRegistrarCancelacionAfijo
	err = json.Unmarshal(DatosSolicitudComoJson, &datosSolicitud)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(datosSolicitud)
	fmt.Println(args[0])

	IDAfijo := datosSolicitud.IDAfijo
	if IDAfijo <= 0 {
		return shim.Error("(Args[0]) IDAfijo: debe tener un Afijo")
	}

	queryString = "{\"selector\":{\"docType\":\"" + cc_afijos_cfg.CFG_ObjectType_Propietarios + "\",\"IDAfijo\":" + strconv.Itoa(IDAfijo) + ",\"FechaBaja\":\"\"}}"
	response = stub.InvokeChaincode(cc_afijos_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("ejecutarConsulta", queryString), stub.GetChannelID())

	if response.Status != shim.OK {
		return shim.Error(response.Message)
	}

	if string(response.Payload) == "[]" {
		return shim.Error("(Args[0]) IDAfijo: [ " + strconv.Itoa(IDAfijo) + " ] no tiene propietarios")
	}

	var queryAfijosPropietarios []cc_afijos_cfg.TipoQueryAfijosPropietarios

	propietariosArray := []int{}
	err = json.Unmarshal(response.Payload, &queryAfijosPropietarios)
	if err != nil {
		return shim.Error(err.Error())
	}

	for propietario := range queryAfijosPropietarios {

		if queryAfijosPropietarios[propietario].Record.IDPersona > 0 {
			propietariosArray = append(propietariosArray, queryAfijosPropietarios[propietario].Record.IDPersona)
		}
	}

	// ---------------------------------------------------------------------------------------------------
	// GRABAR Registros
	// ---------------------------------------------------------------------------------------------------

	var retorno bytes.Buffer

	fechaHoraActual := time.Now()
	fechaHoraActualAsString := fechaHoraActual.String()
	fechaLimite := fechaHoraActual.Add(3 * 24 * time.Hour)
	fechaLimiteAsString := fechaLimite.String()

	// ---------------------------------------------------------------------------------------------------

	IDSolicitud := 0

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
		IDSolicitud = queryInfoChaincode.IDMaximo
	}

	IDSolicitud += 1

	// ---------------------------------------------------------------------------------------------------

	InfoChaincodeAsBytes, err := json.Marshal(&cc_util.InfoChaincode{cc_util.CC_INFO_CONTADOR, IDSolicitud})
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println([]string{cc_cfg.CFG_ObjectType, string(InfoChaincodeAsBytes)})

	err = stub.PutState(cc_cfg.CFG_ObjectType, InfoChaincodeAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ---------------------------------------------------------------------------------------------------

	var SolicitudEstado string
	var FechaHoraEjecucionAsString string
	var ejecutarSolicitud bool

	propietariosArray = cc_util.ArrayIntSinDuplicados(propietariosArray)

	for propietario := range propietariosArray {

		if IDPersonaEjecuta == propietariosArray[propietario] {
			SolicitudEstado = cc_cfg.SolicitudEstadoAprobado
			FechaHoraEjecucionAsString = fechaHoraActualAsString
			if len(propietariosArray) == 1 {
				ejecutarSolicitud = true
			}

		} else {
			SolicitudEstado = cc_cfg.SolicitudEstadoPendiente
			FechaHoraEjecucionAsString = ""
		}

	}

	nuevoSolicitud := &cc_cfg.Solicitudes{cc_cfg.CFG_ObjectType, IDSolicitud, tipoRegistro, args[0], IDPersonaEjecuta, SolicitudEstado, FechaHoraEjecucionAsString, fechaHoraActualAsString, fechaLimiteAsString}
	fmt.Println(nuevoSolicitud)

	nuevoSolicitudAsBytes, err := json.Marshal(nuevoSolicitud)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(cc_cfg.CFG_ObjectType+strconv.Itoa(IDSolicitud), nuevoSolicitudAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	retorno.WriteString(string(nuevoSolicitudAsBytes))

	// ---------------------------------------------------------------------------------------------------

	propietariosArray = cc_util.ArrayIntSinDuplicados(propietariosArray)

	for propietario := range propietariosArray {

		if IDPersonaEjecuta == propietariosArray[propietario] {
			SolicitudEstado = cc_cfg.SolicitudEstadoAprobado
			FechaHoraEjecucionAsString = fechaHoraActualAsString
		} else {
			SolicitudEstado = cc_cfg.SolicitudEstadoPendiente
			FechaHoraEjecucionAsString = ""
		}

		nuevaSolicitudAutorizacion := &cc_cfg.SolicitudesAutorizaciones{
			cc_cfg.CFG_ObjectType_Autorizaciones,
			IDSolicitud,
			propietariosArray[propietario],
			SolicitudEstado,
			FechaHoraEjecucionAsString,
			fechaHoraActualAsString,
			fechaLimiteAsString}
		// fmt.Println(nuevaSolicitudAutorizacion)

		nuevaSolicitudAutorizacionAsBytes, err := json.Marshal(nuevaSolicitudAutorizacion)
		if err != nil {
			return shim.Error(err.Error())
		}

		err = stub.PutState(cc_cfg.CFG_ObjectType_Autorizaciones+strconv.Itoa(IDSolicitud)+"_"+strconv.Itoa(propietariosArray[propietario]), nuevaSolicitudAutorizacionAsBytes)
		if err != nil {
			return shim.Error(err.Error())
		}

		retorno.WriteString(string(nuevaSolicitudAutorizacionAsBytes))
	}

	if ejecutarSolicitud {

		fmt.Println(ejecutarSolicitud)
		response := stub.InvokeChaincode(cc_afijos_cfg.CFG_ChainCodeName, cc_util.ToChaincodeArgs("registrarCancelacionAfijo", args[0], args[1]), stub.GetChannelID())
		if response.Status != shim.OK {
			return shim.Error(response.Message)
		} else {
			return shim.Success(response.Payload)
		}
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
