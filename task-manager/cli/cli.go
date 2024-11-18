package main

import (
	"fmt"
	"os"
	"example.com/tasks"
)

func main() {
	err := tasks.LoadTasksFromFile()
	if err != nil {
		fmt.Println("Erro ao carregar tarefas:", err)
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Uso: task-manager [add|list|complete] [argumentos]")
		return
	}

	command := os.Args[1]

	switch command {
		case "add":
			if len(os.Args) < 3 {
				fmt.Println("Por favor, forneça a descrição da tarefa.")
				return
			}
			description := os.Args[2]
			tasks.AddTask(description)
		case "list":
			tasks.ListTasks()
		case "complete":
			if len(os.Args) < 3 {
				fmt.Println("Por favor, forneça o ID da tarefa para completar.")
				return
			}
			taskID := os.Args[2]
			tasks.CompleteTask(taskID)
		default:
			fmt.Println("Comando não reconhecido.")
	}
}
