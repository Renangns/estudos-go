package main

import (
	"bank/contas"
	"fmt"
)

func PagarBoleto(conta verificaConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificaConta interface {
	Sacar(valor float64) string
}

func main() {
	conta := contas.ContaCorrente{}
	conta.Titular.Nome = "Titular 1"
	conta.Depositar(100)

	conta2 := contas.ContaPoupanca{}
	conta2.Titular.Nome = "Titular 2"
	conta2.Depositar(200)

	PagarBoleto(&conta, 45)
	PagarBoleto(&conta2, 190)
	fmt.Println(conta)
	fmt.Println(conta2)

}
