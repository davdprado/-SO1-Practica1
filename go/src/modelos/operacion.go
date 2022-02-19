package modelos

type Operacion struct {
	Numero1   float32 `json:"numero1"`
	Numero2   float32 `json:"numero2"`
	Operacion string  `json:"operacion"`
	Resultado float32 `json:"resultado"`
	Fecha     string  `json:"fecha"`
}
