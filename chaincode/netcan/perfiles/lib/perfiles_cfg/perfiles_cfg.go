package perfiles_cfg

const (
	CFG_ChainCodeName string = "perfiles"
	CFG_ObjectType    string = "PERFILES_PERSONAS"
)

type PerfilesPersonas struct {
	ObjectType      string `json:"docType"`
	IDPerfilPersona int    `json:"IDPerfilPersona"`
	IDPersona       int    `json:"IDPersona"`
	CODPerfil       string `json:"CODPerfil"`
	FechaAlta       string `json:"FechaAlta"`
	FechaBaja       string `json:"FechaBaja"`
}

const (
	CFG_PerfilVeterinario   string = "VETERINARIO"
	CFG_PerfilAdministrador string = "ADMINISTRADOR"
	CFG_PerfilFederacion    string = "FEDERACION"
)
