package lista

type Lista[T any] interface {

	// EstaVacia devuelve verdadero si la lista no tiene elementos insertados, false en caso contrario.
	EstaVacia() bool

	// InsertarPrimero agrega un nuevo elemento al principio de la lista.
	InsertarPrimero(T)

	// InsertarUltimo agrega un nuevo elemento al final de la lista.
	InsertarUltimo(T)

	// BorrarPrimero elimina el primer elemento de la lista. Si la lista tiene elementos, se quita el primero de la misma y se devuelve ese valor.
	// Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
	BorrarPrimero() T

	// VerPrimero obtiene el valor el primer valor de la lista. Si la lista tiene elementos se devuelve el primer valor de la misma.
	// Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
	VerPrimero() T

	// VerUltimo obtiene el valor el ultimo valor de la lista. Si la lista tiene elementos se devuelve el ultimo valor de la misma.
	// Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
	VerUltimo() T

	// Largo devuelve la cantidad de elementos que tiene la lista.
	Largo() int

	// Iterar recorre los elementos de la lista permitiendo aplicar una función de visita en todos ellos hasta que se termine de recorrer la lista o hasta que la función de visita devuelva false.
	Iterar(visitar func(T) bool)

	// Iterador devuelve una instancia de un iterador del tipo IteradorLista, que permite recorrer los elementos de la lista de manera controlada.
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	// VerActual obtiene el valor del elemento actual en el iterador.
	// Si ya se han iterado todos los elementos o la lista está vacía, entra en pánico con un mensaje "El iterador termino de iterar".
	VerActual() T

	// HaySiguiente devuelve verdadero si el elemento actual no es nulo, false en caso contrario.
	HaySiguiente() bool

	// Siguiente avanza el iterador al próximo elemento de la lista.
	// Si ya se han iterado todos los elementos o la lista está vacía, entra en pánico con un mensaje "El iterador termino de iterar".
	Siguiente()

	// Insertar agrega un elemento en la posición actual en la lista. Luego de la inserción, la iteración sigue desde la posición de inserción misma.
	Insertar(T)

	// Borrar remueve el elemento de la posición actual de la lista y se devuelve ese valor. Luego de borrarlo, la iteración sigue desde el elemento siguiente.
	// Si ya se han iterado todos los elementos o la lista está vacía, entra en pánico con un mensaje "El iterador termino de iterar".
	Borrar() T
}
