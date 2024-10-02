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
	anterior *nodoLista[T]
	actual   *nodoLista[T]
	lista    *listaEnlazada[T]
}

func crearNodo[T any](dato T) *nodoLista[T] {
	return &nodoLista[T]{dato, nil}
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{nil, nil, 0}
}

func (l *listaEnlazada[T]) EstaVacia() bool {
	return l.largo == 0
}

func (l *listaEnlazada[T]) InsertarPrimero(elem T) {
	nuevoNodo := crearNodo(elem)
	if l.EstaVacia() {
		l.ultimo = nuevoNodo
	} else {
		nuevoNodo.proximo = l.primero
	}
	l.primero = nuevoNodo
	l.largo++
}

func (l *listaEnlazada[T]) InsertarUltimo(elem T) {
	nuevoNodo := crearNodo(elem)
	if l.EstaVacia() {
		l.primero = nuevoNodo
	} else {
		l.ultimo.proximo = nuevoNodo
	}
	l.ultimo = nuevoNodo
	l.largo++
}

func (l *listaEnlazada[T]) BorrarPrimero() T {
	validarListaVacia(l)
	primero := l.primero.dato
	l.primero = l.primero.proximo
	if l.primero == nil {
		l.ultimo = nil
	}
	l.largo--
	return primero
}

func (l *listaEnlazada[T]) VerPrimero() T {
	validarListaVacia(l)
	return l.primero.dato
}

func (l *listaEnlazada[T]) VerUltimo() T {
	validarListaVacia(l)
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
	return &iterListaEnlazada[T]{actual: l.primero, lista: l}
}

func (i *iterListaEnlazada[T]) VerActual() T {
	validarIteradorFinalizado(i)
	return i.actual.dato
}

func (i *iterListaEnlazada[T]) HaySiguiente() bool {
	return i.actual != nil
}

func (i *iterListaEnlazada[T]) Siguiente() {
	validarIteradorFinalizado(i)
	i.anterior = i.actual
	i.actual = i.actual.proximo
}

func (i *iterListaEnlazada[T]) Insertar(elem T) {
	nuevoNodo := crearNodo(elem)
	if i.anterior == nil {
		i.lista.primero = nuevoNodo
	} else {
		i.anterior.proximo = nuevoNodo
	}
	if i.actual == nil {
		i.lista.ultimo = nuevoNodo
	}
	nuevoNodo.proximo = i.actual
	i.actual = nuevoNodo
	i.lista.largo++
}

func (i *iterListaEnlazada[T]) Borrar() T {
	validarIteradorFinalizado(i)
	datoEliminado := i.actual.dato
	if i.anterior == nil {
		i.lista.primero = i.actual.proximo
	} else {
		i.anterior.proximo = i.actual.proximo
	}
	i.actual = i.actual.proximo
	if i.actual == nil {
		i.lista.ultimo = i.anterior
	}
	i.lista.largo--
	return datoEliminado
}

func validarListaVacia[T any](l *listaEnlazada[T]) {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
}

func validarIteradorFinalizado[T any](i *iterListaEnlazada[T]) {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}
