package main

import (
	"fmt"
	"net/http"

	"github.com/jeffprestes/cursodego/banco_sql/manipulador"
	"github.com/jeffprestes/cursodego/banco_sql/repo"
)

func init() {
	//fmt.Println("Vamos começar a subir o servidor...")
}

func main() {
	//repo.TestOracle()
	err := repo.AbreConexaoComBancoDeDadosSQL()
	if err != nil {
		fmt.Println("Parando a carga do servidor. Erro ao abrir o banco de dados: ", err.Error())
		return
	}
	db, err := repo.GetDB()
	if err != nil {
		fmt.Println("Deu ruim")
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Deu ruim")
		return
	}
	fmt.Println("Conectou")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá Mundo!")
	})
	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", manipulador.Ola)
	http.HandleFunc("/local/", manipulador.Local)
	fmt.Println("Estou escutando, comandante...")
	http.ListenAndServe(":8181", nil)
}