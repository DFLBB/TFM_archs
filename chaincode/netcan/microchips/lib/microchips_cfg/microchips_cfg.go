package microchips_cfg

const (
	CFG_ChainCodeName string = "microchips"
	CFG_ObjectType    string = "MICROCHIPS_PERROS"
)

type MicrochipsPerros struct {
	ObjectType           string `json:"docType"`
	IDMicrochipPerro     int    `json:"IDMicrochipPerro"`
	IDPerro              int    `json:"IDPerro"`
	IDPersonaVeterinario int    `json:"IDPersonaVeterinario"`
	CODMicrochip         string `json:"CODMicrochip"`
	FechaAlta            string `json:"FechaAlta"`
	FechaBaja            string `json:"FechaBaja"`
}
