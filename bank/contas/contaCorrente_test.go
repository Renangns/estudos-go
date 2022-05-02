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

func TestNaoDeveTransferirValorMenorQueZero(t *testing.T) {
	conta1 := ContaCorrente{}
	conta1.saldo = SALDO

	conta2 := ContaCorrente{}
	conta2.saldo = SALDO

	res := conta1.Transferencia(-SALDO, &conta2)

	if res != "Valor é menor que zero" && conta1.saldo != SALDO && conta2.saldo != SALDO {
		t.Fatalf("Erro %s ", res)
	}
}

func TestNaoDeveTransferirValorMaiorQueSaldoDisponivel(t *testing.T) {
	conta1 := ContaCorrente{}
	conta1.saldo = SALDO

	conta2 := ContaCorrente{}
	conta2.saldo = SALDO

	res := conta1.Transferencia(SALDO+1, &conta2)

	if res != "Conta não possui saldo para realizar a transferencia" && conta1.saldo != SALDO && conta2.saldo != SALDO {
		t.Fatalf("Erro %s ", res)
	}
}

func TestTransferirValor(t *testing.T) {
	valorTransferencia := 200.

	conta1 := ContaCorrente{}
	conta1.saldo = SALDO

	conta2 := ContaCorrente{}
	conta2.saldo = SALDO

	res := conta1.Transferencia(valorTransferencia, &conta2)

	if res != "Transferencia realizada com sucesso" && conta1.saldo != SALDO-valorTransferencia && conta2.saldo != SALDO+valorTransferencia {
		t.Fatalf("Erro %s ", res)
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
