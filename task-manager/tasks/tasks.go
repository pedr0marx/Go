package tasks

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID          int
	Description string
	Completed   bool
}

var taskList []Task
var nextID int = 1
var filename = "tasks.json"

func LoadTasksFromFile() error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil 
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("erro ao ler o arquivo: %v", err)
	}

	err = json.Unmarshal(data, &taskList)
	if err != nil {
		return fmt.Errorf("erro ao desserializar JSON: %v", err)
	}
	
	if len(taskList) > 0 {
		nextID = taskList[len(taskList)-1].ID + 1
	}
	return nil
}

func SaveTasksToFile() error {
	data, err := json.MarshalIndent(taskList, "", "  ")
	if err != nil {
		return fmt.Errorf("erro ao serializar tarefas: %v", err)
	}
	
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("erro ao salvar no arquivo: %v", err)
	}
	return nil
}

func AddTask(description string) {
	task := Task{
		ID:          nextID,
		Description: description,
		Completed:   false,
	}
	taskList = append(taskList, task)
	nextID++

	err := SaveTasksToFile()
	if err != nil {
		fmt.Println("Erro ao salvar tarefas:", err)
		return
	}

	fmt.Printf("Tarefa adicionada: %s (ID: %d)\n", description, task.ID)
}

func ListTasks() {
	if len(taskList) == 0 {
		fmt.Println("Nenhuma tarefa encontrada.")
		return
	}

	for _, task := range taskList {
		status := "não completada"
		if task.Completed {
			status = "completada"
		}
		fmt.Printf("ID: %d - %s [%s]\n", task.ID, task.Description, status)
	}
}

func CompleteTask(taskID string) {
	for i, task := range taskList {
		if fmt.Sprintf("%d", task.ID) == taskID {
			taskList[i].Completed = true

			err := SaveTasksToFile()
			if err != nil {
				fmt.Println("Erro ao salvar tarefas:", err)
				return
			}

			fmt.Printf("Tarefa ID %d marcada como completada.\n", task.ID)
			return
		}
	}
	fmt.Println("Tarefa não encontrada.")
}
