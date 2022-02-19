package modelos

type Operacion struct {
	Numero1   int    `json:"numero1"`
	Numero2   int    `json:"numero2"`
	Operacion string `json:"operacion"`
	Resultado int    `json:"resultado"`
	Fecha     string `json:"fecha"`
}
