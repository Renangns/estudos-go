package contas

import (
	"bank/clientes"
)

type ContaPoupanca struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Operacao      int
	saldo         float64
}

func (c *ContaPoupanca) Sacar(valor float64) string {
	if valor < 0 {
		return "Valor para saque deve ser maior que zero"
	}

	if c.saldo < valor {
		return "Saldo insuficiente para o saque"
	}

	c.saldo -= valor
	return "Saldo Ok"
}

func (c *ContaPoupanca) Depositar(valor float64) (string, float64) {
	if valor < 0 {
		return "Valor do deposito menor que zero", c.saldo
	}

	c.saldo += valor
	return "Deposito realizado com sucesso", c.saldo
}

func (c *ContaPoupanca) Transferencia(valor float64, contaDestino *ContaPoupanca) string {
	if valor < 0 {
		return "Valor é menor que zero"
	}
	if valor > c.saldo {
		return "Conta não possui saldo para realizar a transferencia"
	}

	contaDestino.Depositar(valor)
	c.saldo -= valor

	return "Transferencia realizada com sucesso"
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
