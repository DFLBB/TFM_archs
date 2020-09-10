package personas_cfg

const (
	CFG_ChainCodeName string = "personas"
	CFG_ObjectType    string = "PERSONAS"
)

type Personas struct {
	ObjectType             string `json:"docType"`
	IDPersona              int    `json:"IDPersona"`
	Nombre                 string `json:"Nombre"`
	Apellido1              string `json:"Apellido1"`
	Apellido2              string `json:"Apellido2"`
	TipoDocumento          string `json:"TipoDocumento"`
	IdentificadorDocumento string `json:"IdentificadorDocumento"`
	PaisEmisor             string `json:"PaisEmisor"`
	Certificado            string `json:"Certificado"`
	FechaAlta              string `json:"FechaAlta"`
	FechaBaja              string `json:"FechaBaja"`
}
