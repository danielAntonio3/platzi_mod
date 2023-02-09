package main

// Estamos usando un modulo externo creado
import (
	"github.com/danielAntonio3/calculadora-go"
	"strings"
)

func main() {
	operacion := calculadora_go.LeerEntrada()

	// Obtener el operador
	operator := calculadora_go.GetTypeOperation(operacion)

	// Divide la cadena de texto, con un valor como patron de separación
	valores := strings.Split(operacion, operator)

	// Creamos un objeto de nuestro struct
	calculadora := calculadora_go.Calc{}
	calculadora.Value1 = calculadora_go.GetValue(valores[0])
	calculadora.Value2 = calculadora_go.GetValue(valores[1])
	calculadora.Operator = operator

	// Realizamos la operación
	calculadora.GetOperation()
}
