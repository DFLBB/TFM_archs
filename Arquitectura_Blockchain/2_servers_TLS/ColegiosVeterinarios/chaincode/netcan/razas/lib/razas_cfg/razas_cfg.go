package razas_cfg

const (
	CFG_ChainCodeName     string = "razas"
	CFG_ObjectType        string = "RAZAS"
	CFG_ObjectType_Grupos string = "GRUPOS"
)

type Razas struct {
	ObjectType string `json:"docType"`
	IDRaza     int    `json:"IDRaza"`
	Nombre     string `json:"Nombre"`
	IDGrupo    int    `json:"IDGrupo"`
	FechaAlta  string `json:"FechaAlta"`
	FechaBaja  string `json:"FechaBaja"`
}

type Grupos struct {
	ObjectType string `json:"docType"`
	IDGrupo    int    `json:"IDGrupo"`
	Nombre     string `json:"Nombre"`
	FechaAlta  string `json:"FechaAlta"`
	FechaBaja  string `json:"FechaBaja"`
}
