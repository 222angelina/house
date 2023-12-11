package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	// Создаем команду для выполнения команды docker build
	cmd := exec.Command("docker", "build", "-t", "myimage", ".")

	// Устанавливаем рабочую директорию, где находится Dockerfile
	cmd.Dir = "C:\\Users\\gelya\\GolandProjects\\awesomeProject\\хаус\\Dockerfile"

	// Устанавливаем вывод команды на стандартный вывод
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Выполняем команду
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Не удалось выполнить команду docker build: %v", err)
	}
}
