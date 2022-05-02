package main

import (
	"bank/contas"
	"fmt"
)

func PagarBoleto(conta contas.Conta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificaConta interface {
	Sacar(valor float64) string
}

func main() {
	var conta1 contas.Conta
	var conta2 contas.Conta

	contaCorrente := contas.ContaCorrente{}
	contaCorrente.Titular.Nome = "Titular 1"
	conta1 = &contaCorrente

	conta1.Depositar(100)

	contaPoupanca := contas.ContaPoupanca{}
	contaPoupanca.Titular.Nome = "Titular 2"
	conta2 = &contaPoupanca

	conta2.Depositar(200)

	PagarBoleto(conta1, 45)
	PagarBoleto(conta2, 190)
	conta1.Transferencia(50, conta2)
	fmt.Println(conta1)
	fmt.Println(conta2)

}
