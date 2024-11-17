package tasks

import (
	"encoding/json"
	"fmt"
	"os"
)

// Estrutura para armazenar tarefas
type Task struct {
	ID          int
	Description string
	Completed   bool
}

var taskList []Task
var nextID int = 1
var filename = "tasks.json"

// Função que carrega as tarefas do arquivo JSON
func LoadTasksFromFile() error {
	// Verifica se o arquivo tasks.json existe
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil // Se o arquivo não existir, retorna sem erro
	}

	// Lê o conteúdo do arquivo JSON
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("erro ao ler o arquivo: %v", err)
	}

	// Desserializa o JSON para a lista de tarefas
	err = json.Unmarshal(data, &taskList)
	if err != nil {
		return fmt.Errorf("erro ao desserializar JSON: %v", err)
	}

	// Atualiza o próximo ID, baseado no maior ID existente
	if len(taskList) > 0 {
		nextID = taskList[len(taskList)-1].ID + 1
	}
	return nil
}

// Função que salva as tarefas no arquivo JSON
func SaveTasksToFile() error {
	// Serializa a lista de tarefas para JSON
	data, err := json.MarshalIndent(taskList, "", "  ")
	if err != nil {
		return fmt.Errorf("erro ao serializar tarefas: %v", err)
	}

	// Escreve os dados no arquivo tasks.json
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("erro ao salvar no arquivo: %v", err)
	}
	return nil
}

// Adiciona uma nova tarefa à lista
func AddTask(description string) {
	task := Task{
		ID:          nextID,
		Description: description,
		Completed:   false,
	}
	taskList = append(taskList, task)
	nextID++

	// Salva as tarefas no arquivo
	err := SaveTasksToFile()
	if err != nil {
		fmt.Println("Erro ao salvar tarefas:", err)
		return
	}

	fmt.Printf("Tarefa adicionada: %s (ID: %d)\n", description, task.ID)
}

// Lista todas as tarefas
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

// Marca uma tarefa como completa
func CompleteTask(taskID string) {
	for i, task := range taskList {
		if fmt.Sprintf("%d", task.ID) == taskID {
			taskList[i].Completed = true

			// Salva as tarefas no arquivo após a alteração
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
