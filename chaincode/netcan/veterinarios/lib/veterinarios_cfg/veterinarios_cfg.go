package veterinarios_cfg

const (
	CFG_ChainCodeName string = "veterinarios"
	CFG_ObjectType    string = "COLEGIATURAS_PERSONAS"
)

type ColegiaturasPersonas struct {
	ObjectType           string `json:"docType"`
	IDColegiaturaPersona int    `json:"IDColegiaturaPersona"`
	IDPersona            int    `json:"IDPersona"`
	CODColegiatura       string `json:"CODColegiatura"`
	FechaAlta            string `json:"FechaAlta"`
	FechaBaja            string `json:"FechaBaja"`
}
