package lista

type nodoLista[T any] struct {
	dato    T
	proximo *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type iterListaEnlazada[T any] struct {
	actual *nodoLista[T]
}

func crearNodo[T any]() *nodoLista[T] {
	return &nodoLista[T]{}
}

func CrearListaEnlazada[T any]() Lista[T] {
	nodo := crearNodo[T]()
	return &listaEnlazada[T]{nodo, nodo, 0}
}

func (l *listaEnlazada[T]) EstaVacia() bool {
	return l.largo == 0
}

func (l *listaEnlazada[T]) InsertarPrimero(elem T) {
	if l.EstaVacia() {
		l.primero.dato = elem
		l.largo++
	}
	nuevoNodo := crearNodo[T]()
	nuevoNodo.dato = elem
	nuevoNodo.proximo = l.primero
	l.primero = nuevoNodo
	l.largo++
}

func (l *listaEnlazada[T]) InsertarUltimo(elem T) {
	if l.EstaVacia() {
		l.ultimo.dato = elem
		l.largo++
	}
	l.ultimo.proximo = crearNodo[T]()
	l.ultimo = l.ultimo.proximo
	l.ultimo.dato = elem
	l.largo++
}

func (l *listaEnlazada[T]) BorrarPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	primero := l.primero.dato
	l.primero = l.primero.proximo
	return primero
}

func (l *listaEnlazada[T]) VerPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.primero.dato
}

func (l *listaEnlazada[T]) VerUltimo() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.ultimo.dato
}

func (l *listaEnlazada[T]) Largo() int {
	return l.largo
}

func (l *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := l.primero
	for actual != nil {
		if visitar(actual.dato) {
			actual = actual.proximo
		} else {
			return
		}
	}

}

func (l *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterListaEnlazada[T]{}
}

func (i *iterListaEnlazada[T]) VerActual() T {
	return i.actual.dato
}

func (i *iterListaEnlazada[T]) HaySiguiente() bool {
	return i.actual != nil
}

func (i *iterListaEnlazada[T]) Siguiente() {
	i.actual = i.actual.proximo
}

func (i *iterListaEnlazada[T]) Insertar(T) {

}

func (i *iterListaEnlazada[T]) Borrar() T {

}
