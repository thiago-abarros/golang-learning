# Variaveis no Golang

## Declaração x Atribuição

O conceito de declaração e atribuição é importante no Golang, pois, variáveis que são declaradas mas não recebem uma atribuição causam erros de compilação, visto que existem variáveis que não estão sendo utilizadas no código.

### Declaração

Quando apenas declaramos a **existência** de uma variável, sem **atribuir** valor à mesma.

### Atribuição

Para a atribuição, precisamos de uma variável existente para podermos **atribuir um valor à ela**, ou seja, após declararmos, nós atribuimos algum valor a ela.

### Exemplo Prático

```go
// Isso aqui é uma declaração de variável
var a string

func main() {
 // Isso aqui é uma atribuição de variável
 a = "hello"
 println(a)

 // Aqui fazemos a declaração e a atribuição de uma variável
 b := "world"
 println(b)
}
```

## Declarando Funções

Para declarar funções em golang, colocamos os argumentos e o tipo de retorno da função, que neste caso é inteiro. Caso a função não tenha retorno, deixe vazio.

```go
// função com retorno
func Soma(a int, b int) int {
 return a + b
}

// função sem retorno
func HelloWorld() {
    print("Hello World")
}
```
