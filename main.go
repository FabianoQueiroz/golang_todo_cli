package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"slices"
)

func criando_csv(path_csv string) {
	// Se nao existe, cria o arquivo com o cabecalho
	fmt.Println("INFO: Arquivo de dados nao existe")
	fmt.Println("INFO: Criando arquivo...")

	file, err := os.Create(path_csv)

	if err != nil {
		fmt.Println("ERROR: Erro ao criar o arquivo ->", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	atividades := [][]string{
		{"ID", "ATIVIDADE", "CRIACAO", "STATUS"},
	}

	for _, record := range atividades {
		if err := writer.Write(record); err != nil {
			fmt.Println("ERROR: Erro ao escrever o cabecalho no arquivo ->", err)
			return
		}
	}

	fmt.Println("INFO: Arquivo criado com sucesso")
}

func add_atividade(complemento string, file *os.File) {
	fmt.Println("TODO: Implementar a funcao add")
	return
}

func del_atividade(complemento string, file *os.File) {
	fmt.Println("TODO: Implementar a funcao del")
	return
}
func list_atividade(file *os.File) {
	fmt.Println("TODO: Implementar a funcao list")
	return
}
func done_atividade(file *os.File) {
	fmt.Println("TODO: Implementar a funcao done")
	return
}

func main() {

	path_csv := "todo_csv.csv"

	// Checa se o 'csv' que guarda as informacoes existe
	if _, err := os.Stat(path_csv); os.IsNotExist(err) {
		criando_csv(path_csv)
	} else {

		file, err := os.Open(path_csv)
		if err != nil {
			fmt.Print("ERROR: Arquivo de dados nao pode ser aberto ->", err)
			return
		}
		defer file.Close()

		// Identifica qual o tipo de acao o comando deve realizar (add, del, list, done)
		args_com_compl := []string{"add", "del", "done"}

		if qtde_args := len(os.Args); qtde_args > 3 {
			fmt.Println("ERROR: Muitos argumentos")
		} else if qtde_args == 1 {
			fmt.Println("              TODO-CLI in GO", "\n",
				"------------------------------------------------------", "\n",
				"main add  <atividade> : Adiciona uma atividade a lista", "\n",
				"main del  <atividade> : Deleta uma atividade a lista", "\n",
				"main done <atividade> : Marca uma atividade como feita", "\n",
				"main list             : Lista todas as atividades")
		} else if tipo_arg := os.Args[1]; qtde_args == 2 && slices.Contains(args_com_compl, tipo_arg) {
			fmt.Println("ERROR: Funcao", tipo_arg, "precisa do nome da atividade e nenhum foi fornecido")
			return
		} else {

			complemento := os.Args[len(os.Args)-1]

			switch tipo_arg {
			case "add":
				add_atividade(complemento, file)
			case "del":
				del_atividade(complemento, file)
			case "list":
				list_atividade(file)
			case "done":
				done_atividade(file)
			default:
				fmt.Println("ERROR: Comando nao reconhecido ->", tipo_arg)
				fmt.Println("ERROR: Utilize add, del, list ou done")
				return
			}
		}
	}
}
