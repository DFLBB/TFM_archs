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

type TiporegistrarCambioPropietarioPropietario struct {
	IDPersona int `json:"IDPersona"`
}

type TiporegistrarCambioPropietario struct {
	IDAfijo      int                                         `json:"IDAfijo"`
	Propietarios []TiporegistrarCambioPropietarioPropietario `json:"Propietarios"`
}

type TiporegistrarCancelacionAfijo struct {
	IDAfijo int `json:"IDAfijo"`
}

type TipoQueryAfijos struct {
	Key    string `json:"Key"`
	Record Afijos `json:"Record"`
}

type TipoQueryAfijosPropietarios struct {
	Key    string             `json:"Key"`
	Record AfijosPropietarios `json:"Record"`
}
