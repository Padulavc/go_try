package main

import (
	"io/ioutil"
	"github.com/isvaldo/converter/mock"
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"time"
	"os"
	"strconv"
)

type Job struct {
	Empenho mock.Empenho
	Year    int
	Pagina  int
}

func getEmpenhoData(year int, pagina int, cn chan Job) {

	if !CheckIfNotFileExist(year, pagina) {
		log.Println("ignore write file", year, pagina)
		return
	}
	url := fmt.Sprintf("https://gatewayapi.prodam.sp.gov.br/financas/orcamento/sof/v2.1.0/consultaEmpenhos?anoEmpenho=%d&mesEmpenho=12&codOrgao=16&numPagina=%d", year, pagina)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error on start request")
		return
	}

	req.Header.Add("Authorization", "REPLACE")
	req.Header.Add("Cache-Control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error on execute request", year, pagina, err)
		return
	}

	fmt.Println("Request executed", year, pagina, res.StatusCode)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Error on readall request", year, pagina, res.StatusCode)
		return
	}
	var empenho mock.Empenho
	if err := json.Unmarshal(body, &empenho); err != nil {
		log.Println("Error on unmarshal request", year, pagina, string(body), res.StatusCode)
		return
	}
	log.Println("end year,page", year, pagina)
	cn <- Job{
		Empenho: empenho,
		Year:    year,
		Pagina:  pagina,
	}
}

func GetGovSourceData(year int, cn chan Job) {

	url := fmt.Sprintf("https://gatewayapi.prodam.sp.gov.br/financas/orcamento/sof/v2.1.0/consultaEmpenhos?anoEmpenho=%d&mesEmpenho=12&codOrgao=16&numPagina=1", year)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error on start request")
		return
	}

	req.Header.Add("Authorization", "Bearer 4abf361f8c520e74728a0d51817bc4")
	req.Header.Add("Cache-Control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error on execute request", year)
		return
	}

	fmt.Println("Request executed", year, res.StatusCode)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Error on readall request", year, res.StatusCode)
		return
	}
	var empenho mock.Empenho
	if err := json.Unmarshal(body, &empenho); err != nil {
		log.Println("Error on unmarshal request", year, string(body), res.StatusCode)
		return
	}

	qt := empenho.Metadados.QtdPaginas

	cn <- Job{
		Empenho: empenho,
		Year:    year,
		Pagina:  1,
	}
	for i := 1; i < int(qt)+1; i++ {
		log.Println("start year,page", year, i)
		time.Sleep(time.Second * 25)
		go getEmpenhoData(year, i, cn)
	}
}

func CheckIfNotFileExist(year int, page int) bool {
	if _, err := os.Stat(fmt.Sprintf("/home/isvaldo/goworkspace/src/github.com/isvaldo/converter/result/page-%d-%d.csv", year, page)); os.IsNotExist(err) {
		return true
	}
	return false
}

func createNewFile(empenho mock.Empenho, year int, page int) {
	log.Println("start write file!")
	f, err := os.Create(fmt.Sprintf("/home/isvaldo/goworkspace/src/github.com/isvaldo/converter/result/page-%d-%d.csv", year, page))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	list := empenho.LstEmpenhos

	for _, value := range list {

		f.WriteString(fmt.Sprintf("%f,%f,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s\n",
			value.AnoEmpenho,
			value.CodCategoria,
			value.TxtCategoriaEconomica,
			value.CodElemento,
			strconv.FormatFloat(value.CodEmpenho, 'f', 6, 64),
			value.CodEmpresa,
			value.CodFonteRecurso,
			value.CodFuncao,
			strconv.FormatFloat(value.CodGrupo, 'f', 6, 64),
			value.TxtGrupoDespesa,
			value.CodItemDespesa,
			strconv.FormatFloat(value.CodModalidade, 'f', 6, 64),
			value.TxtModalidadeAplicacao,
			value.CodOrgao,
			strconv.FormatFloat(value.CodProcesso, 'f', 6, 64),
			value.CodPrograma,
			value.CodProjetoAtividade,
			value.CodSubElemento,
			value.CodSubFuncao,
			value.CodUnidade,
			value.DatEmpenho,
			strconv.FormatFloat(value.MesEmpenho, 'f', 6, 64),
			value.NomEmpresa,
			value.NumCpfCnpj,
			strconv.FormatFloat(value.NumReserva, 'f', 6, 64),
			value.TxtDescricaoOrgao,
			value.TxtDescricaoUnidade,
			value.TxtDescricaoElemento,
			value.TxtDescricaoFonteRecurso,
			value.TxtDescricaoFuncao,
			value.TxtDescricaoItemDespesa,
			value.TxtDescricaoPrograma,
			value.TxtDescricaoProjetoAtividade,
			value.TxtRazaoSocial,
			value.TxtDescricaoSubElemento,
			value.TxtDescricaoSubFuncao,
			strconv.FormatFloat(value.ValAnuladoEmpenho, 'f', 6, 64),
			strconv.FormatFloat(value.ValEmpenhadoLiquido, 'f', 6, 64),
			strconv.FormatFloat(value.ValLiquidado, 'f', 6, 64),
			strconv.FormatFloat(value.ValPagoExercicio, 'f', 6, 64),
			strconv.FormatFloat(value.ValPagoRestos, 'f', 6, 64),
			strconv.FormatFloat(value.ValTotalEmpenhado, 'f', 6, 64),
		),
		)
		f.Sync()
	}
	log.Println("finish write file")
}

func main() {

	cn := make(chan Job, 2)
	//go GetGovSourceData(2011, cn)
	//go GetGovSourceData(2012, cn)
	go GetGovSourceData(2013, cn)
	//go GetGovSourceData(2015, cn)
	//go GetGovSourceData(2016, cn)
	//go GetGovSourceData(2017, cn)
	//go GetGovSourceData(2018, cn)

	log.Println("wait events!")
	for event := range cn {
		log.Println("new event arrived", event.Year, event.Pagina)
		if !CheckIfNotFileExist(event.Year, event.Pagina) {
			log.Println("Ignore file", event.Year, event.Pagina)
			return
		}
		createNewFile(event.Empenho, event.Year, event.Pagina)
	}

}
