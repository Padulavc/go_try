package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func mayhem(e error){
	if e != nil {
		panic(e)
	}
}
func getFiles(){
	/*
			CONFIGURE REQUEST by year and page
	*/
	for year := 2010; year <= 2018; year++{
		for page := 1; page<= 60; page++ {
			url := fmt.Sprintf("https://gatewayapi.prodam.sp.gov.br:443/financas/orcamento/sof/v2.1.0/consultaEmpenhos?anoEmpenho=%d&mesEmpenho=12&codOrgao=16&numPagina=%d", year, page)
			req, _ := http.NewRequest("GET", url, nil)
			req.Header.Add("Authorization", "REPLACE_ME")
			req.Header.Add("Cache-Control", "no-cache")
			fmt.Println(fmt.Sprintf("REQUESTING FILE year %d pag %d", year, page))
			res, _ := http.DefaultClient.Do(req)
			body, _ := ioutil.ReadAll(res.Body)
			fmt.Println(string(body))

			//OPEN AND WRITE FILE
			fmt.Println("writing file")
			file , err := os.Create("masterfile.json")
			mayhem(err)

			//weird defer func
			//defer func() {
			//	if err := file.Close(); err != nil {
			//		panic(err)
			//	}
			//}()

			//write the chunk, i guess?
			fmt.Println("writing data")
			file.Write(body)
			fmt.Println("Data written")
		}
	}
}

func main(){
	go getFiles()
}

