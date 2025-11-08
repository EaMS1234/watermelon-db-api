package crud
type Corpo_Localizacao struct{
	ID_Corpo_d_agua int `json:"corpo"`
	ID_Localizacao int `json:"localizacao"`
	ID_corpo_localizacao int `json:"id"`
}

func (Corpo_Localizacao) TableName() string {
	return "Corpo_Localizacao"
}

