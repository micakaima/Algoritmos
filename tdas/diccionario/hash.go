package diccionario

import (
	"fmt"
	"hash/fnv"
)

type estado int

const (
	VACIO estado = iota
	OCUPADO
	BORRADO
	TAM_INICIAL                 = 10
	FACTOR_REDIMENSION_AGRANDAR = 0.7
	FACTOR_REDIMENSION_ACHICAR  = 0.3
	PROPORCION_REDIMENSION      = 2
)

type celdaHash[K comparable, V any] struct {
	clave  K
	dato   V
	estado estado
}

type hashCerrado[K comparable, V any] struct {
	tabla     []celdaHash[K, V]
	tam       int
	guardados int
	borrados  int
}

type iterHashCerrado[K comparable, V any] struct {
	posActual int
	hash      *hashCerrado[K, V]
}

func crearTabla[K comparable, V any](tamanio int) []celdaHash[K, V] {
	return make([]celdaHash[K, V], tamanio)
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return &hashCerrado[K, V]{tabla: crearTabla[K, V](TAM_INICIAL), tam: TAM_INICIAL}
}

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {
	if float64(hash.guardados+hash.borrados)/float64(hash.tam) >= FACTOR_REDIMENSION_AGRANDAR {
		redimensionarTabla(hash, hash.tam*PROPORCION_REDIMENSION)
	}

	pos := buscarPos(hash, clave)
	if hash.tabla[pos].estado == VACIO {
		hash.tabla[pos] = celdaHash[K, V]{clave, dato, OCUPADO}
		hash.guardados++
	} else {
		hash.tabla[pos].dato = dato
	}
}

func (hash *hashCerrado[K, V]) Pertenece(clave K) bool {
	return hash.tabla[buscarPos(hash, clave)].estado == OCUPADO
}

func (hash *hashCerrado[K, V]) Obtener(clave K) V {
	posicion := buscarPos(hash, clave)
	if hash.tabla[posicion].estado == OCUPADO {
		return hash.tabla[posicion].dato
	} else {
		panic("La clave no pertenece al diccionario")
	}
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V {
	pos := buscarPos(hash, clave)
	if hash.tabla[pos].estado != OCUPADO {
		panic("La clave no pertenece al diccionario")
	}
	datoBorrado := hash.tabla[pos].dato
	hash.tabla[pos].estado = BORRADO
	hash.guardados--
	hash.borrados++

	if float64(hash.guardados+hash.borrados)/float64(hash.tam) <= FACTOR_REDIMENSION_ACHICAR && hash.tam > TAM_INICIAL {
		redimensionarTabla(hash, hash.tam/PROPORCION_REDIMENSION)
	}
	return datoBorrado
}

func (hash *hashCerrado[K, V]) Cantidad() int {
	return hash.guardados
}

func (hash *hashCerrado[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for _, celda := range hash.tabla {
		if celda.estado != OCUPADO || visitar(celda.clave, celda.dato) {
			continue
		} else {
			return
		}
	}
}

func (hash *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	posActual := 0
	for posActual < hash.tam && hash.tabla[posActual].estado != OCUPADO {
		posActual++
	}
	return &iterHashCerrado[K, V]{posActual, hash}
}

func (iter *iterHashCerrado[K, V]) HaySiguiente() bool {
	return iter.posActual < iter.hash.tam
}

func (iter *iterHashCerrado[K, V]) VerActual() (K, V) {
	validarIteradorFinalizado(iter)
	celda := iter.hash.tabla[iter.posActual]
	return celda.clave, celda.dato
}

func (iter *iterHashCerrado[K, V]) Siguiente() {
	validarIteradorFinalizado(iter)
	iter.posActual++
	for iter.HaySiguiente() && iter.hash.tabla[iter.posActual].estado != OCUPADO {
		iter.posActual++
	}
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

// Funcion de Hashing importada desde el paquete "fnv" de Go
func fnvHashing[K comparable](clave K) uint32 {
	h := fnv.New32()
	h.Write(convertirABytes(clave))
	return h.Sum32()
}

// Me da la pos donde debo actualizar [o] la primera posicion VACIA que encuentre
func buscarPos[K comparable, V any](hash *hashCerrado[K, V], clave K) int {
	pos := fnvHashing(clave) % uint32(hash.tam)
	for hash.tabla[pos].estado != VACIO && !(hash.tabla[pos].clave == clave && hash.tabla[pos].estado != BORRADO) {
		pos++
		if int(pos) >= hash.tam {
			pos = 0
		}
	}
	return int(pos)
}

func redimensionarTabla[K comparable, V any](hash *hashCerrado[K, V], nuevaCapacidad int) {
	nuevaTabla := crearTabla[K, V](nuevaCapacidad)
	hashAux := &hashCerrado[K, V]{tabla: nuevaTabla, tam: nuevaCapacidad}
	for _, celda := range hash.tabla {
		if celda.estado == OCUPADO {
			nuevaPos := buscarPos(hashAux, celda.clave)
			nuevaTabla[nuevaPos] = celda
		}
	}
	hash.tabla = nuevaTabla
	hash.tam = nuevaCapacidad
}

func validarIteradorFinalizado[K comparable, V any](iter *iterHashCerrado[K, V]) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}
