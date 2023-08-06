package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	Nome     string
	Ra       int
	Idade    int
	Cpf      int
	Telefone int
}

func main() {
	bdUrl := "postgres://postgres:postgre@localhost:5432/exercicio1_orm"
	bd, err := gorm.Open(postgres.Open(bdUrl))

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = bd.AutoMigrate(&Aluno{})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ALUNO := Aluno{}

	bd.First(&ALUNO)

	if ALUNO.ID > 0 {
		if ALUNO.Ra == 2022014019 {
			fmt.Println("Existe registro e vai apagar!!!!")

			//exclui o registro CPF
			bd.Model(&ALUNO).UpdateColumn("Ra", nil)
		} else {
			//Altera os dados
			bd.Model(&ALUNO).Update("Telefone", 46999877414)

		}


	} else {
		fmt.Println("Realiza inclusão")
		//Inserindo dados no banco
		ALUNO.Nome = "Gabriel"
		ALUNO.Ra = 2022014019
		ALUNO.Idade = 20
		ALUNO.Cpf = 19862753951
		ALUNO.Telefone = 46999029477

		bd.Create(&ALUNO)
	}
}
