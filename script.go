package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq"
)

type Dados struct {
	Host     string `json:"host"`
	Port     uint   `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

func main() {

	content, err := ioutil.ReadFile("/etc/default/config.api.json")
	if err != nil {
		log.Fatal("Erro para abrir o arquivo: ", err)
	}

	var payload Dados
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Erro during Unmarshal(): ", err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		payload.Host, payload.Port, payload.User, payload.Password, payload.Dbname)
	//	newFunction(payload)

	db, err := sql.Open("postgres", psqlInfo)
	CheckError(err)

	defer db.Close()

	// insert
	insertStmt := `INSERT INTO metricas (status, data_criacao) VALUES(true, now())`
	_, e := db.Exec(insertStmt)
	CheckError(e)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
