# Exercícios do capítulo 20

## 1)

- Alem da goroutine principal, crie duas outras goroutines.
- Cada goroutine adicional devem fazer um print separado.
- Utilize waitgroups para fazer com que suas goroutines finalizem antes de o programa terminar.


## 2)

- Esse exercício vai reforçar seus conhecimentos sobre conjuntos de métodos.
- Crie um tipo para um struct chamado "pessoa"
- Crie um método "falar" para este tipo que tenha um receiver ponteiro (*pessoa)
- Crie uma interface, "humanos", que seja implementada por tipos com o método "falar"
- Crie uma função "dizerAlgumaCoisa" cujo parâmetro seja do tipo "humanos" e
  que chame o método "falar"
- Demonstre no seu código:
    - Que você pode utilizar um valor do tipo "*pessoa" na função "dizerAlgumaCoisa"
    - Que você não pode utilizar um valor do tipo "pessoa" na função "dizerAlgumaCoisa"


## 3)

- Utilizando goroutines, crie um programa incrementador:
    - Tenha uma variável com o valor da contagem
    - Crie um monte de goroutines, onde cada uma deve:
        - Ler o valor do contador
        - Salvar este valor
        - Fazer yield da thread com runtime.Gosched()
        - Incrementar o valor salvo
        - Copiar o novo valor para a variável inicial
    - Utilize WaitGroups para que seu programa não finalize antes de suas goroutines.
    - Demonstre que há uma condição de corrida utilizando a flag -race

## 4)

- Utilize mutex para consertar a condição de corrida do exercício anterior

## 5)
