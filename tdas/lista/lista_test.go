package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const TAMAÑO_TEST_VOLUMEN = 10000

func TestCrearUnaLista(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() }, "La lista inicializada esta vacia por defecto y no tiene elementos")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() }, "La lista inicializada esta vacia por defecto y no tiene elementos")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() }, "La lista inicializada esta vacia por defecto y no tiene elementos")
}

func TestInsertarPrimeroUnElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(9)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 9, lista.VerPrimero(), "Al insertar en el principio un unico elemento en una lista vacia, el primer elemento deberia ser el mismo.")
	require.Equal(t, 9, lista.VerUltimo(), "Al insertar en el principio un unico elemento en una lista vacia, el ultimo elemento deberia ser el mismo.")
	require.Equal(t, 9, lista.BorrarPrimero())
}

func TestInsertarUltimoUnElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarUltimo(5)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 5, lista.VerUltimo(), "Al insertar a lo ultimo un unico elemento en una lista vacia, el ultimo elemento deberia ser el mismo.")
	require.Equal(t, 5, lista.VerPrimero(), "Al insertar a lo ultimo un unico elemento en una lista vacia, el primer elemento deberia ser el mismo.")
	require.Equal(t, 5, lista.BorrarPrimero())
}

func TestLargoLista(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	require.Equal(t, 0, lista.Largo(), "El largo de una lista recien creada debe ser cero")
	lista.InsertarPrimero(5)
	lista.InsertarUltimo(4)
	lista.InsertarPrimero(7)
	require.Equal(t, 3, lista.Largo(), "El largo de una lista con tres elementos debe ser tres")
	lista.BorrarPrimero()
	lista.BorrarPrimero()
	lista.BorrarPrimero()
	require.Equal(t, 0, lista.Largo(), "El largo de una lista cuyos elementos fueron todos borrados debe ser cero")

}

func TestVolumenInsertarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < TAMAÑO_TEST_VOLUMEN+1; i++ {
		lista.InsertarPrimero(i)
		require.Equal(t, i, lista.VerPrimero(), "El ultimo elemento ingresado queda al principio de la lista")
		require.Equal(t, 0, lista.VerUltimo(), "El primer elemento ingresado queda al final de la lista.")
	}
	require.False(t, lista.EstaVacia())
	for i := TAMAÑO_TEST_VOLUMEN; i > -1; i-- {
		require.Equal(t, i, lista.BorrarPrimero())
	}

	require.True(t, lista.EstaVacia())
}

func TestVolumenInsertarUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < TAMAÑO_TEST_VOLUMEN+1; i++ {
		lista.InsertarUltimo(i)
		require.Equal(t, 0, lista.VerPrimero(), "El primer elemento ingresado queda al principio de la lista")
		require.Equal(t, i, lista.VerUltimo(), "El ultimo elemento ingresado queda al final de la lista.")
	}
	require.False(t, lista.EstaVacia())
	for i := 0; i < TAMAÑO_TEST_VOLUMEN+1; i++ {
		require.Equal(t, i, lista.BorrarPrimero())
	}

	require.True(t, lista.EstaVacia())
}

func TestListaPostBorrar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 11; i++ {
		lista.InsertarPrimero(i)
	}
	for j := 0; j < 11; j++ {
		lista.BorrarPrimero()
	}
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })

	lista.InsertarUltimo(11)
	require.Equal(t, 11, lista.VerPrimero())
	require.Equal(t, 11, lista.VerUltimo())
	require.Equal(t, 11, lista.BorrarPrimero())
}

func TestInsertarPrimeroVariosElementosTipoString(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()

	lista.InsertarPrimero("t")
	require.False(t, lista.EstaVacia(), "Una lista con un elemento no esta vacia")
	lista.InsertarPrimero("K")
	lista.InsertarPrimero("!")
	lista.InsertarPrimero(" ")
	lista.InsertarPrimero("%")
	require.False(t, lista.EstaVacia(), "Una lista con  varios elementos no esta vacia")
	require.Equal(t, "%", lista.VerPrimero(), "Al insertar varios elementos al principio el ultimo insertado queda como el primer elemento")
	require.Equal(t, "t", lista.VerUltimo(), "Al insertar varios elementos al principio el primero queda como el ultimo")
	require.Equal(t, "%", lista.BorrarPrimero())
	require.Equal(t, " ", lista.BorrarPrimero())
	require.Equal(t, "!", lista.BorrarPrimero())
	require.Equal(t, "K", lista.BorrarPrimero())
	require.Equal(t, "t", lista.BorrarPrimero())

}

func TestInsertarPrimeroVariosElementosTipoArregloDeFloats(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[[]float32]()

	lista.InsertarPrimero([]float32{3.56, 6.45})
	require.False(t, lista.EstaVacia(), "Una lista con un elemento no esta vacia")
	lista.InsertarPrimero([]float32{0.56, 8.2})
	lista.InsertarPrimero([]float32{})
	lista.InsertarPrimero([]float32{1, 0.0})
	lista.InsertarPrimero([]float32{3.66, 9.99})
	require.False(t, lista.EstaVacia(), "Una lista con  varios elementos no esta vacia")
	require.Equal(t, []float32{3.66, 9.99}, lista.VerPrimero(), "Al insertar varios elementos al principio el ultimo insertado queda como el primer elemento")
	require.Equal(t, []float32{3.56, 6.45}, lista.VerUltimo(), "Al insertar varios elementos al principio el primero queda como el ultimo")
	require.Equal(t, []float32{3.66, 9.99}, lista.BorrarPrimero())
	require.Equal(t, []float32{1, 0.0}, lista.BorrarPrimero())
	require.Equal(t, []float32{}, lista.BorrarPrimero())
	require.Equal(t, []float32{0.56, 8.2}, lista.BorrarPrimero())
	require.Equal(t, []float32{3.56, 6.45}, lista.BorrarPrimero())

}

func TestIteradorInternoIterarTodaLaLista(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(4)
	lista.InsertarPrimero(23)
	lista.InsertarPrimero(15)
	lista.InsertarPrimero(81)
	suma := 0
	lista.Iterar(func(v int) bool {
		suma += v
		return true
	})
	require.Equal(t, 123, suma)
}

func TestIteradorInternoIterarConCondicionDeCorte(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(4)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(8)
	lista.InsertarPrimero(5)
	suma := 0
	lista.Iterar(func(v int) bool {
		suma++
		return v != 1
	})
	require.Equal(t, 3, suma)
}

func TestIteradorInternoCompleto(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(4)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(8)
	lista.InsertarPrimero(5)
	suma, cont := 0, 0
	lista.Iterar(func(v int) bool {
		if v%2 == 0 {
			suma += v
			cont += 1
		}
		return cont != 3
	})
	require.Equal(t, 14, suma)
	require.Equal(t, 3, cont)

}

func TestIteradorExterno(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 10; i++ {
		lista.InsertarUltimo(i)
	}
	j := 0
	for i := lista.Iterador(); i.HaySiguiente(); i.Siguiente() {
		require.Equal(t, j, i.VerActual())
		j++
	}
}

func TestIteradorExternoEnListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.False(t, iter.HaySiguiente())
}

func TestIteradorExternoInsertarEnListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(1)
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())
}

func TestIteradorExternoInsertarPrimeraPos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	// Comenzando con una lista vacia, cada elemento insertado se agregaran siempre
	// del lado izquierdo, quedando al principio de la lista.
	for i := 0; i <= 10; i++ {
		iter.Insertar(i)
	}
	require.Equal(t, 10, lista.VerPrimero())
	require.Equal(t, 0, lista.VerUltimo())
	require.Equal(t, 11, lista.Largo())
	valor := 10
	for iter2 := lista.Iterador(); iter2.HaySiguiente(); iter2.Siguiente() {
		require.Equal(t, valor, lista.BorrarPrimero())
		valor--
	}
}

func TestIteradorExternoInsertarUltimaPos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	// Comenzando con una lista vacia, cada elemento insertado luego de iterar al
	// siguiente se agregaran siempre al final de la lista, del lado derecho.
	for i := 0; i <= 10; i++ {
		iter.Insertar(i)
		iter.Siguiente()
	}
	require.Equal(t, 0, lista.VerPrimero())
	require.Equal(t, 10, lista.VerUltimo())
	require.Equal(t, 11, lista.Largo())
	valor := 0
	for iter2 := lista.Iterador(); iter2.HaySiguiente(); iter2.Siguiente() {
		require.Equal(t, valor, lista.BorrarPrimero())
		valor++
	}
}

func TestIteradorExternoInsertarPosMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	for i := 0; i <= 10; i++ {
		iter.Insertar(i)
		iter.Siguiente()
	}
	numRequerido := 0
	for iter2 := lista.Iterador(); iter2.HaySiguiente(); iter2.Siguiente() {
		require.Equal(t, numRequerido, iter2.VerActual())
		numRequerido++
	}
	require.Equal(t, 11, lista.Largo())
}

func TestIteradorExternoBorrarPrimerElem(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i <= 10; i++ {
		lista.InsertarUltimo(i)
	}
	iter := lista.Iterador()
	require.Equal(t, 0, iter.Borrar())
	require.Equal(t, 1, iter.VerActual())
	require.Equal(t, 1, lista.VerPrimero())
}

func TestIteradorExternoBorrarUltimoElem(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i <= 10; i++ {
		lista.InsertarUltimo(i)
	}
	iter := lista.Iterador()
	for j := 0; j < 10; j++ {
		iter.Siguiente()
	}
	require.Equal(t, 10, iter.Borrar())
	require.Equal(t, 9, lista.VerUltimo())
}

func TestIteradorExternoBorrarElemMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i <= 10; i++ {
		lista.InsertarUltimo(i)
	}
	for iter := lista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		if iter.VerActual() == 5 {
			require.Equal(t, 5, iter.Borrar())
		}
	}
	for iter2 := lista.Iterador(); iter2.HaySiguiente(); iter2.Siguiente() {
		require.NotEqual(t, 5, iter2.VerActual())
	}
	require.Equal(t, 10, lista.Largo())
}

func TestIteradorExternoPostBorrarTodos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i <= 10; i++ {
		lista.InsertarUltimo(i)
	}
	iter := lista.Iterador()
	for j := 0; j <= 10; j++ {
		require.Equal(t, j, iter.Borrar())
	}
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })

	require.True(t, lista.EstaVacia(), "La lista cuyos elementos fueron todos borrados actua como una vacia")
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
}
