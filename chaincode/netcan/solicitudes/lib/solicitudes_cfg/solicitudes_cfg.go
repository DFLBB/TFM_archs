package solicitudes_cfg

const (
	CFG_ChainCodeName             string = "solicitudes"
	CFG_ObjectType                string = "SOLICITUDES"
	CFG_ObjectType_Autorizaciones string = "SOLICITUDES_AUTORIZACIONES"
)

const (
	SolicitudEstadoPendiente string = "PENDIENTE"
	SolicitudEstadoAprobado  string = "APROBADO"
	SolicitudEstadoRechazado string = "RECHAZADO"
)

type Solicitudes struct {
	ObjectType           string `json:"docType"`
	IDSolicitud          int    `json:"IDSolicitud"`
	TipoSolicitud        string `json:"TipoSolicitud"`
	JSONSolicitud        string `json:"JSONSolicitud"`
	IDPersonaSolicitante int    `json:"IDPersonaSolicitante"`
	EstadoSolicitud      string `json:"EstadoSolicitud"`
	FechaEjecucion       string `json:"FechaEjecucion"`
	FechaAlta            string `json:"FechaAlta"`
	FechaBaja            string `json:"FechaBaja"`
}

type SolicitudesAutorizaciones struct {
	ObjectType      string `json:"docType"`
	IDSolicitud     int    `json:"IDSolicitud"`
	IDPersona       int    `json:"IDPersona"`
	EstadoSolicitud string `json:"EstadoSolicitud"`
	FechaEjecucion  string `json:"FechaEjecucion"`
	FechaAlta       string `json:"FechaAlta"`
	FechaBaja       string `json:"FechaBaja"`
}
