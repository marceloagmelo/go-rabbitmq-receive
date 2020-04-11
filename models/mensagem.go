package models

import (
	"fmt"
	"strconv"

	"github.com/marceloagmelo/go-rabbitmq-receive/lib"
	"github.com/marceloagmelo/go-rabbitmq-receive/utils"
)

//Mensagem estrutura de mensagem
type Mensagem struct {
	ID     int    `db:"id" json:"id"`
	Titulo string `db:"titulo" json:"titulo"`
	Texto  string `db:"texto" json:"texto"`
	Status int    `db:"status" json:"status"`
}

//MensagemModel recebe a tabela do banco de dados
var MensagemModel = lib.Sess.Collection("mensagem")

//AtualizarMensagem enviada
func AtualizarMensagem(id string) string {
	mensagem := ""
	var mensagemID, _ = strconv.Atoi(id)

	var mensagemModel Mensagem

	resultado := MensagemModel.Find("id=?", mensagemID)
	if count, err := resultado.Count(); count < 1 {
		mensagem = utils.CheckErr(err)
		mensagem = "Não foi possivel encontrar o usuário!"
	}

	if err := resultado.One(&mensagemModel); err != nil {
		mensagem = utils.CheckErr(err)
		mensagem = "Não foi possivel encontrar o usuário!"
	}
	mensagemModel.Status = 2

	if err := resultado.Update(mensagemModel); err != nil {
		mensagem = utils.CheckErr(err)
		mensagem = "Erro ao tentar atualizar o usuário!"
	}

	if mensagem == "" {
		mensagem = fmt.Sprintf("Mensagem %s atualiazada", id)
	}
	return mensagem
}
