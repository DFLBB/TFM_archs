package perros_cfg

const (
	CFG_ChainCodeName           string = "perros"
	CFG_ObjectType              string = "PERROS"
	CFG_ObjectType_Propietarios string = "PERROS_PROPIETARIOS"
)

type Perros struct {
	ObjectType      string `json:"docType"`
	IDPerro         int    `json:"IDPerro"`
	Nombre          string `json:"Nombre"`
	IDAfijo         int    `json:"IDAfijo"`
	IDSexo          int    `json:"IDSexo"`
	IDPerroMadre    int    `json:"IDPerroMadre"`
	IDPerroPadre    int    `json:"IDPerroPadre"`
	IDRaza          int    `json:"IDRaza"`
	FechaNacimiento string `json:"FechaNacimiento"`
	FechaDefuncion  string `json:"FechaDefuncion"`
	FechaAlta       string `json:"FechaAlta"`
	FechaBaja       string `json:"FechaBaja"`
}

type PerrosPropietarios struct {
	ObjectType         string `json:"docType"`
	IDPerroPropietario int    `json:"IDPerroPropietario"`
	IDPerro            int    `json:"IDPerro"`
	IDPersona          int    `json:"IDPersona"`
	FechaAlta          string `json:"FechaAlta"`
	FechaBaja          string `json:"FechaBaja"`
}
