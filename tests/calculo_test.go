/*
	PARA EXECUTAR O TESTE BASTA ACESSAR O DITETÓRIO DO ARQUIVO E EXECUTAR O COMANDO
	GO TEST

	COMANDO GO TEST -COVER EXIBE A PORCENTAGEM DE COBERTURA DOS TESTES
	COMANDO GO TEST -COVERPROFILE=RESULTADO.OUT EXIBE A PORCENTAGEM DE COBERTURA DOS TESTES
	COMANDO GO TOOL COVER  -FUNC=RESULT.OUT EXIBE OS DADOS DO ARQUIVO CRIADO
	COMANDO GO TOOL COVER  -HTML=RESULT.OUT CRIA UM ARQUIVO COM OS DADOS DE SAÍDA NO FORMATO HTML

*/

package tests

import (
	"testing"
)

const errorMessage string = "O valor esperado é %v mas o resultado encontrado foi %v"

func TestDeveRetornar10(t *testing.T) {
	expectedValue := 10.0

	value := Average(10.0, 10.0)

	if value != expectedValue {
		t.Errorf(errorMessage, expectedValue, value)
	}
}

func TestDeveRetornar15(t *testing.T) {
	expectedValue := 15.0

	value := Average(15.0, 15.0)

	if value != expectedValue {
		t.Errorf(errorMessage, expectedValue, value)
	}
}
