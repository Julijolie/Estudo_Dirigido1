package main

import "fmt"

type ListNode struct {
	Info int
	Prox *ListNode
}

func criaNo(numero int) *ListNode {
	return &ListNode{Info: numero, Prox: nil}
}

func insereInicio(lista **ListNode, numero int) {
	no := criaNo(numero)
	if *lista == nil {
		*lista = no
		no.Prox = *lista
	} else {
		no.Prox = (*lista).Prox
		(*lista).Prox = no
	}
}

func removeInicio(lista **ListNode, numero *int) bool {
	if *lista == nil {
		return false
	} else {
		aux := (*lista).Prox
		*numero = aux.Info
		if aux == *lista {
			*lista = nil
		} else {
			(*lista).Prox = aux.Prox
		}
		return true
	}
}

func exibirLista(lista *ListNode) {
	if lista == nil {
		fmt.Println("Lista vazia\n\n")
		return
	}

	aux := lista
	for {
		fmt.Printf("%d\t", aux.Info)
		aux = aux.Prox
		if aux == lista {
			break
		}
	}
	fmt.Println("\n\n")
}

func main() {
	var minhaLista *ListNode
	insereInicio(&minhaLista, 10)
	insereInicio(&minhaLista, 20)
	insereInicio(&minhaLista, 30)

	var numeroRemovido int
	if removeInicio(&minhaLista, &numeroRemovido) {
		fmt.Printf("Número removido: %d\n", numeroRemovido)
	}

	exibirLista(minhaLista)
}
