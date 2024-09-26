package lista

type Lista[T any] interface {

	// EstaVacia devuelve verdadero si la lista no tiene elementos insertados, false en caso contrario.
	EstaVacia() bool

	// InsertarPrimero agrega un nuevo elemento al principio de la lista.
	InsertarPrimero(T)

	// InsertarUltimo agrega un nuevo elemento al final de la lista.
	InsertarUltimo(T)

	// BorrarPrimero elimina el primer elemento de la lista. Si la lista tiene elementos, se quita el primero de la misma y se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
	BorrarPrimero() T

	// VerPrimero obtiene el valor el primer valor de la lista. Si la lista tiene elementos se devuelve el primer valor de la misma. Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
	VerPrimero() T

	// VerUltimo obtiene el valor el ultimo valor de la lista. Si la lista tiene elementos se devuelve el ultimo valor de la misma. Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
	VerUltimo() T

	// Largo devuelve la cantidad de elementos que tiene la lista.
	Largo() int

	// Iterar
	Iterar(visitar func(T) bool)

	// Iterador
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	// VerActual
	VerActual() T

	// HaySiguiente
	HaySiguiente() bool

	// Siguiente
	Siguiente()

	// Insertar
	Insertar(T)

	// Borrar
	Borrar() T
}
