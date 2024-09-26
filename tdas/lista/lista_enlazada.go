package lista

type listaEnlazada[T any] struct {
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{}
}

func (l *listaEnlazada[T]) EstaVacia() bool {
	return true
}

func (l *listaEnlazada[T]) InsertarPrimero(T)

func (l *listaEnlazada[T]) InsertarUltimo(T)

func (l *listaEnlazada[T]) BorrarPrimero() T

func (l *listaEnlazada[T]) VerPrimero() T

func (l *listaEnlazada[T]) VerUltimo() T

func (l *listaEnlazada[T]) Largo() int

func (l *listaEnlazada[T]) Iterar(visitar func(T) bool)

func (l *listaEnlazada[T]) Iterador() IteradorLista[T]
