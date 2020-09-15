package perros_cfg

const (
	CFG_ChainCodeName                   string = "perros"
	CFG_ObjectType                      string = "PERROS"
	CFG_ObjectType_Propietarios         string = "PERROS_PROPIETARIOS"
	CFG_ObjectType_ReconocimientosRazas string = "PERROS_RECONOCIMIENTO_RAZA"
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

type PerroReconocimientoRaza struct {
	ObjectType                string `json:"docType"`
	IDPerroReconocimientoRaza int    `json:"IDPerroReconocimientoRaza"`
	IDPerro                   int    `json:"IDPerro"`
	IDRaza                    int    `json:"IDRaza"`
	Justificacion             string `json:"Justificacion"`
	IDPersona                 int    `json:"IDPersona"`
	FechaEjecucion            string `json:"FechaEjecucion"`
}

const (
	HEMBRA int = 0
	MACHO  int = 1
)

type TipoRegistrarCamadaPerro struct {
	Nombre string `json:"Nombre"`
	IDSexo int    `json:"IDSexo"`
}

type TipoRegistrarCamadaPropietario struct {
	IDPersona int `json:"IDPersona"`
}

type TipoRegistrarCamada struct {
	Perros          []TipoRegistrarCamadaPerro       `json:"Perros"`
	IDPerroMadre    int                              `json:"IDPerroMadre"`
	IDPerroPadre    int                              `json:"IDPerroPadre"`
	IDAfijo         int                              `json:"IDAfijo"`
	IDRaza          int                              `json:"IDRaza"`
	FechaNacimiento string                           `json:"FechaNacimiento"`
	Propietarios    []TipoRegistrarCamadaPropietario `json:"Propietarios"`
}

type TipoRegistrarCambioPropietarioPropietario struct {
	IDPersona int `json:"IDPersona"`
}

type TipoRegistrarCambioPropietario struct {
	IDPerro      int                                         `json:"IDPerro"`
	Propietarios []TipoRegistrarCambioPropietarioPropietario `json:"Propietarios"`
}

type TipoRegistrarDefuncionPerro struct {
	IDPerro        int    `json:"IDPerro"`
	FechaDefuncion string `json:"FechaDefuncion"`
}

type TipoQueryPerros struct {
	Key    string `json:"Key"`
	Record Perros `json:"Record"`
}

type TipoQueryPerrosPropietarios struct {
	Key    string             `json:"Key"`
	Record PerrosPropietarios `json:"Record"`
}
