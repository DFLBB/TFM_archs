package afijos_cfg

const (
	CFG_ChainCodeName           string = "afijos"
	CFG_ObjectType              string = "AFIJOS"
	CFG_ObjectType_Propietarios string = "AFIJOS_PROPIETARIOS"
)

type Afijos struct {
	ObjectType string `json:"docType"`
	IDAfijo    int    `json:"IDAfijo"`
	Nombre     string `json:"Nombre"`
	FechaAlta  string `json:"FechaAlta"`
	FechaBaja  string `json:"FechaBaja"`
}

type AfijosPropietarios struct {
	ObjectType         string `json:"docType"`
	IDAfijoPropietario int    `json:"IDAfijoPropietario"`
	IDAfijo            int    `json:"IDAfijo"`
	IDPersona          int    `json:"IDPersona"`
	FechaAlta          string `json:"FechaAlta"`
	FechaBaja          string `json:"FechaBaja"`
}
