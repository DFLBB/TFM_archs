package vacunas_cfg

const (
	CFG_ChainCodeName                      string = "vacunas"
	CFG_ObjectType                         string = "VACUNAS_PERROS"
	CFG_ObjectType_VacunasPerrosProteccion string = "VACUNAS_PERROS_PROTECCION"
	CFG_ObjectType_VacunasProteccion       string = "VACUNAS_PROTECCION"
)

type VacunasPerros struct {
	ObjectType           string `json:"docType"`
	IDVacunaPerro        int    `json:"IDVacunaPerro"`
	IDPerro              int    `json:"IDPerro"`
	IDPersonaVeterinario int    `json:"IDPersonaVeterinario"`
	CODVacuna            string `json:"CODVacuna"`
	FechaAlta            string `json:"FechaAlta"`
	FechaBaja            string `json:"FechaBaja"`
}

type VacunasPerrosProteccion struct {
	ObjectType              string `json:"docType"`
	IDVacunaPerroProteccion int    `json:"IDVacunaPerroProteccion"`
	IDVacunaPerro           int    `json:"IDVacunaPerro"`
	IDVacunaProteccion      int    `json:"IDVacunaProteccion"`
	FechaAlta               string `json:"FechaAlta"`
	FechaBaja               string `json:"FechaBaja"`
}

type VacunasProteccion struct {
	ObjectType         string `json:"docType"`
	IDVacunaProteccion int    `json:"IDVacunaProteccion"`
	Nombre             string `json:"Nombre"`
	FechaAlta          string `json:"FechaAlta"`
	FechaBaja          string `json:"FechaBaja"`
}
