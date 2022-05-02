package contas

type Conta interface {
	Sacar(valor float64) string
	Depositar(valor float64) (string, float64)
	Transferencia(valor float64, contaDestino Conta) string
	ObterSaldo() float64
}
