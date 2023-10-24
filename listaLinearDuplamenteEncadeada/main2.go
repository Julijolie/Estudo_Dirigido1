package main

import "fmt"

// Definição da estrutura do nó
type Node struct {
	Value    int
	Previous *Node
	Next     *Node
}

// Definição da estrutura da lista
type DoublyLinkedList struct {
	Head *Node
}

// Função para inserir um nó no início da lista
func (list *DoublyLinkedList) Inserir(valor int) {
	novoNo := &Node{Value: valor, Next: list.Head}

	if list.Head != nil {
		list.Head.Previous = novoNo
	}

	list.Head = novoNo
	novoNo.Previous = nil
}

// Função para remover um nó da lista
func (list *DoublyLinkedList) Remover(no *Node) {
	if no.Previous != nil {
		no.Previous.Next = no.Next
	} else {
		list.Head = no.Next
	}

	if no.Next != nil {
		no.Next.Previous = no.Previous
	}
}

// Função para buscar um nó na lista
func (list *DoublyLinkedList) Buscar(valor int) *Node {
	aux := list.Head
	for aux != nil {
		if aux.Value == valor {
			return aux // Nó encontrado
		}
		aux = aux.Next
	}
	return nil // Nó não encontrado
}

// Função para exibir os nós da lista
func (list *DoublyLinkedList) Imprimir() {
	aux := list.Head
	fmt.Print("Lista: ")
	for aux != nil {
		fmt.Printf("%d <-> ", aux.Value)
		aux = aux.Next
	}
	fmt.Println("nil")
}

func main() {
	// Inicializa uma lista duplamente encadeada
	lista := DoublyLinkedList{}

	// Inserindo alguns nós no início da lista
	lista.Inserir(30)
	lista.Inserir(20)
	lista.Inserir(10)

	// Exibindo a lista
	lista.Imprimir()

	// Buscando um nó na lista
	valorParaBuscar := 20
	resultadoBusca := lista.Buscar(valorParaBuscar)
	if resultadoBusca != nil {
		fmt.Printf("Nó com valor %d encontrado.\n", valorParaBuscar)
	} else {
		fmt.Printf("Nó com valor %d não encontrado.\n", valorParaBuscar)
	}

	// Removendo um nó da lista
	noParaRemover := lista.Buscar(20)
	if noParaRemover != nil {
		lista.Remover(noParaRemover)
	}

	// Exibindo a lista após a remoção
	lista.Imprimir()
}
