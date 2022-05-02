package contas

import (
	"testing"
)

const SALDO float64 = 500.00

func TestNaoDeveriaSacarValorMenorQueZero(t *testing.T) {
	conta := ContaCorrente{}
	conta.saldo = SALDO
	res := conta.Sacar(-500)
	if res != "Valor para saque deve ser maior que zero" {
		t.Fatalf("Erro %s", res)
	}
}

func TestNaoDeveSacarValorMaiorQueOSaldo(t *testing.T) {
	conta := ContaCorrente{}
	conta.saldo = SALDO
	res := conta.Sacar(SALDO + 1.)
	if res != "Saldo insuficiente para o saque" {
		t.Fatalf("Erro %s", res)
	}
}

func TestSacar(t *testing.T) {
	conta := ContaCorrente{}
	conta.saldo = SALDO
	res := conta.Sacar(SALDO - 1.)
	if res != "Saldo Ok" && conta.saldo != SALDO-1. {
		t.Fatalf("Erro %s", res)
	}
}

func TestNaoDeveDepositarValorMenorQueZero(t *testing.T) {
	conta := ContaCorrente{}
	conta.saldo = SALDO
	res, valor := conta.Depositar(-SALDO)
	if res != "Valor do deposito menor que zero" && conta.saldo != valor {
		t.Fatalf("Erro %s, valor %f", res, valor)
	}
}

func TestObterSaldo(t *testing.T) {
	conta := ContaCorrente{}
	conta.saldo = SALDO
	saldoConta := conta.ObterSaldo()

	if conta.saldo != saldoConta {
		t.Fatalf(
			"Saldo incorreto o saldo esperado era de %f saldo obtido foi de %f", SALDO, saldoConta)
	}
}
