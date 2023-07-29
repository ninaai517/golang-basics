package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const qntMonitoramento = 3
const delay = 5

func main() {

	exibeIntroducao()
	registraLog("site-teste", false)

	for {
		exibeMenu()

		comando := lerComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Comando incorreto!")
			os.Exit(-1)
		}
	}

}

func exibeIntroducao() {
	nome := "Lola"
	fmt.Println("Olá ", nome)
}

func lerComando() int {
	var comando int
	fmt.Scanf("%d", &comando)
	return comando
}

func exibeMenu() {
	fmt.Println("1-Iniciar monitoramento")
	fmt.Println("2-Exibir logs ")
	fmt.Println("0-Sair do programa")
}

func iniciarMonitoramento() {
	fmt.Println("Iniciando monitoramento...")
	sites := leSitesDoArquivo()

	for i := 0; i < qntMonitoramento; i++ {
		for i, site := range sites {
			fmt.Println("Testando site:", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")

}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Erro detectado:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("O site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("O site:", site, "está com problemas! Status:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {

	var sites []string
	arquivo, err := os.Open("sites.txt")

	//arquivo, err := ioutil.ReadFile("sites.txt")

	//ioutil.Readfile: retorna um array de bites []

	if err != nil {
		fmt.Println("Erro detectado:", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		fmt.Println(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool){

	arquivo, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + "-" + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs(){

	arquivo, err := ioutil.ReadFile("logs.txt")

	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}
