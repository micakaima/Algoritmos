package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCrearUnaLista(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() }, "La lista inicializada esta vacia por defecto y no tiene elementos")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() }, "La lista inicializada esta vacia por defecto y no tiene elementos")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() }, "La lista inicializada esta vacia por defecto y no tiene elementos")
}

func TestInsertarPrincipioUnElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia())
	lista.InsertarPrimero(9)
	require.Equal(t, 9, lista.VerPrimero(), "Al insertar en el principio un unico elemento en una lista vacia, el primer elemento deberia ser el mismo.")
	require.Equal(t, 9, lista.VerUltimo(), "Al insertar en el principio un unico elemento en una lista vacia, el ultimo elemento deberia ser el mismo.")
	require.Equal(t, 9, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
}

func TestInsertarUltimoUnElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia())
	lista.InsertarUltimo(5)
	require.Equal(t, 5, lista.VerUltimo(), "Al insertar a lo ultimo un unico elemento en una lista vacia, el ultimo elemento deberia ser el mismo.")
	require.Equal(t, 5, lista.VerPrimero(), "Al insertar a lo ultimo un unico elemento en una lista vacia, el primer elemento deberia ser el mismo.")
	require.Equal(t, 5, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
}

func TestLargoLista(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	require.Equal(t, 0, lista.Largo(), "El largo de una lista recien creada debe ser cero")
	lista.InsertarPrimero(5)
	lista.InsertarUltimo(4)
	lista.InsertarPrimero(7)
	require.Equal(t, 3, lista.Largo(), "El largo de una lista con tres elementos debe ser tres")
	lista.InsertarUltimo(2)
	lista.InsertarPrimero(9)
	lista.InsertarUltimo(1)
	lista.InsertarPrimero(3)
	require.Equal(t, 7, lista.Largo(), "El largo de una lista con siete elementos debe ser siete")

}

func TestInsertarPrincioioVariosElementosTipoEntero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(9)
	require.False(t, lista.EstaVacia(), "Una lista con un elemento no esta vacia")
	lista.InsertarPrimero(6)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(7)
	lista.InsertarPrimero(2)
	require.False(t, lista.EstaVacia(), "Una lista con  varios elementos no esta vacia")
	require.Equal(t, 2, lista.VerPrimero(), "Al insertar varios elementos al principio el ultimo insertado queda como el primer elemento")
	require.Equal(t, 9, lista.VerUltimo(), "Al insertar varios elementos al principio el primero queda como el ultimo")
	require.Equal(t, 2, lista.BorrarPrimero())
	require.Equal(t, 7, lista.BorrarPrimero())
	require.Equal(t, 3, lista.BorrarPrimero())
	require.Equal(t, 6, lista.BorrarPrimero())
	require.Equal(t, 9, lista.BorrarPrimero())

}

func TestDeVolumenInsertarPrimero10000Elementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	volumen := 10000
	for i := 0; i < volumen+1; i++ {
		lista.InsertarPrimero(i)
	}
	require.False(t, lista.EstaVacia())
	require.Equal(t, 10000, lista.VerPrimero(), "El ultimo elemento ingresado queda al principio de la lista")
	require.Equal(t, 0, lista.VerUltimo(), "El primer elemento ingresado queda al final de la lista.")
	for i := volumen; i > -1; i-- {
		require.Equal(t, i, lista.BorrarPrimero())
	}

	require.True(t, lista.EstaVacia())
}

func TestDeVolumenInsertarUltimo10000Elementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	volumen := 10000
	for i := 0; i < volumen+1; i++ {
		lista.InsertarUltimo(i)
	}
	require.False(t, lista.EstaVacia())
	require.Equal(t, 0, lista.VerPrimero(), "El primer elemento ingresado queda al principio de la lista")
	require.Equal(t, 10000, lista.VerUltimo(), "El ultimo elemento ingresado queda al final de la lista.")
	for i := 0; i < volumen+1; i++ {
		require.Equal(t, i, lista.BorrarPrimero())
	}

	require.True(t, lista.EstaVacia())
}

func TestInsertarPrincioioVariosElementosTipoString(t *testing.T) {
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

func TestInsertarPrincioioVariosElementosTipoArregloDeFloats(t *testing.T) {
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
		if v == 1 {
			return false
		}
		return true
	})
	require.Equal(t, 3, suma)
}

func TestIteradorInterno(t *testing.T) {
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
		if cont == 3 {
			return false
		}
		return true
	})
	require.Equal(t, 14, suma)

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
