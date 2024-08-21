package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Pessoa struct {
	ID        string    `json: "id,omitempty"`
	Nome      string    `json: "nome,omitempty"`
	Sobrenome string    `json: "sobrenome,omitempty"`
	Endereco  *Endereco `json: "endereco,omitempty"`
}

type Endereco struct {
	Cidade string `json: "endereco,omitempty"`
	Estado string `json: "estado,omitempty"`
}

var pessoas []Pessoa

// Esta função mostra todos os contatos da variável pessoas
func GetPessoas(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(pessoas)
}

// Esta função mostra apenas um contato
func GetPessoa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range pessoas {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Pessoa{})
}

// Função para criar um novo contato
func CreatePessoa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var pessoa Pessoa
	_ = json.NewDecoder(r.Body).Decode(&pessoa)
	pessoa.ID = params["id"]
	pessoas = append(pessoas, pessoa)
	json.NewEncoder(w).Encode(pessoas)
}

// Função para deletar um contato
func DeletePessoa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range pessoas {
		if item.ID == params["id"] {
			pessoas = append(pessoas[:index], pessoas[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(pessoas)
	}
}

func main() {
	router := mux.NewRouter()
	pessoas := append(pessoas, Pessoa{ID: "1", Nome: "Reydner", Sobrenome: "Jorge",
		Endereco: &Endereco{Cidade: "Itumbiara", Estado: "Goias"}})

	pessoas = append(pessoas, Pessoa{ID: "2", Nome: "Alan", Sobrenome: "Silva",
		Endereco: &Endereco{Cidade: "Itumbiara", Estado: "Goias"}})

	pessoas = append(pessoas, Pessoa{ID: "3", Nome: "Junio", Sobrenome: "Oliveira",
		Endereco: &Endereco{Cidade: "Itumbiara", Estado: "Goias"}})

	router.HandleFunc("/contato", GetPessoas).Methods("GET")
	router.HandleFunc("/contato{id}", GetPessoa).Methods("GET")
	router.HandleFunc("/contato{id}", CreatePessoa).Methods("POST")
	router.HandleFunc("/contato{id}", DeletePessoa).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
