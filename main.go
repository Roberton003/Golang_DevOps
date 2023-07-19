package main

import (
    "fmt"
)

func main() {
    // Define as variáveis name e salario
    name, salario := "Roberto", 5000

    // Chama a função setName para imprimir uma mensagem de saudação
    setName(name)

    // Chama a função addSalary para adicionar um bônus ao salário atual
    // e armazena o novo salário e o valor do bônus em novas variáveis
    newSalary, bonus := addSalary(salario, 500)

    // Imprime o novo salário e o valor do bônus
    fmt.Println("Novo salario é", name, newSalary)
    fmt.Println("Bonus é", bonus)
}

// Define a função setName para imprimir uma mensagem de saudação
func setName(name string) {
    fmt.Println("Hello", name)
}

// Define a função addSalary para adicionar um bônus ao salário atual
// e retorna o novo salário e o valor do bônus
func addSalary(valorAtual int, bonus int) (int, int) {
    return valorAtual + bonus, bonus
}

// Define a função getName para retornar uma string
func getName() string {
    return "Roberto"
}