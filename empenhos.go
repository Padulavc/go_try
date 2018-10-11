package mock

type Empenho struct {
	Metadados struct {
		TxtStatus       string      `json:"txtStatus"`
		TxtMensagemErro interface{} `json:"txtMensagemErro"`
		QtdPaginas      float64         `json:"qtdPaginas"`
	} `json:"metadados"`
	LstEmpenhos []struct {
		AnoEmpenho                   float64    `json:"anoEmpenho"`
		CodCategoria                 float64    `json:"codCategoria"`
		TxtCategoriaEconomica        string  `json:"txtCategoriaEconomica"`
		CodElemento                  string  `json:"codElemento"`
		CodEmpenho                   float64    `json:"codEmpenho"`
		CodEmpresa                   string  `json:"codEmpresa"`
		CodFonteRecurso              string  `json:"codFonteRecurso"`
		CodFuncao                    string  `json:"codFuncao"`
		CodGrupo                     float64 `json:"codGrupo"`
		TxtGrupoDespesa              string  `json:"txtGrupoDespesa"`
		CodItemDespesa               string  `json:"codItemDespesa"`
		CodModalidade                float64   `json:"codModalidade"`
		TxtModalidadeAplicacao       string  `json:"txtModalidadeAplicacao"`
		CodOrgao                     string  `json:"codOrgao"`
		CodProcesso                  float64 `json:"codProcesso"`
		CodPrograma                  string  `json:"codPrograma"`
		CodProjetoAtividade          string  `json:"codProjetoAtividade"`
		CodSubElemento               string  `json:"codSubElemento"`
		CodSubFuncao                 string  `json:"codSubFuncao"`
		CodUnidade                   string  `json:"codUnidade"`
		DatEmpenho                   string  `json:"datEmpenho"`
		MesEmpenho                   float64     `json:"mesEmpenho"`
		NomEmpresa                   string  `json:"nomEmpresa"`
		NumCpfCnpj                   string  `json:"numCpfCnpj"`
		NumReserva                   float64    `json:"numReserva"`
		TxtDescricaoOrgao            string  `json:"txtDescricaoOrgao"`
		TxtDescricaoUnidade          string  `json:"txtDescricaoUnidade"`
		TxtDescricaoElemento         string  `json:"txtDescricaoElemento"`
		TxtDescricaoFonteRecurso     string  `json:"txtDescricaoFonteRecurso"`
		TxtDescricaoFuncao           string  `json:"txtDescricaoFuncao"`
		TxtDescricaoItemDespesa      string  `json:"txtDescricaoItemDespesa"`
		TxtDescricaoPrograma         string  `json:"txtDescricaoPrograma"`
		TxtDescricaoProjetoAtividade string  `json:"txtDescricaoProjetoAtividade"`
		TxtRazaoSocial               string  `json:"txtRazaoSocial"`
		TxtDescricaoSubElemento      string  `json:"txtDescricaoSubElemento"`
		TxtDescricaoSubFuncao        string  `json:"txtDescricaoSubFuncao"`
		ValAnuladoEmpenho            float64     `json:"valAnuladoEmpenho"`
		ValEmpenhadoLiquido          float64 `json:"valEmpenhadoLiquido"`
		ValLiquidado                 float64 `json:"valLiquidado"`
		ValPagoExercicio             float64 `json:"valPagoExercicio"`
		ValPagoRestos                float64 `json:"valPagoRestos"`
		ValTotalEmpenhado            float64 `json:"valTotalEmpenhado"`
	} `json:"lstEmpenhos"`
}
