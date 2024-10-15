package diccionario

import TDAPila "tdas/pila"

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      func(K, K) int
}

type iterAbb[K comparable, V any] struct {
	pila TDAPila.Pila[*nodoAbb[K, V]]
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{nil, 0, funcion_cmp}
}

func (ab *abb[K, V]) Guardar(clave K, dato V) {
	nuevaHoja := &nodoAbb[K, V]{nil, nil, clave, dato}
	if ab.raiz == nil {
		ab.raiz = nuevaHoja
	} else {
		puntero := buscarPuntero(ab.raiz, ab.cmp, clave)
		*puntero = nuevaHoja
	}
	ab.cantidad++
}

func (ab *abb[K, V]) Pertenece(clave K) bool {
	puntero := buscarPuntero(ab.raiz, ab.cmp, clave)
	return *puntero != nil
}

func (ab *abb[K, V]) Obtener(clave K) V {
	puntero := buscarPuntero(ab.raiz, ab.cmp, clave)
	if *puntero == nil {
		panic("La clave no pertenece al diccionario")
	}
	return (*puntero).dato
}

func (ab *abb[K, V]) Borrar(clave K) V {
	puntero := buscarPuntero(ab.raiz, ab.cmp, clave)
	if *puntero == nil {
		panic("La clave no pertenece al diccionario")
	}
	datoBorrado := (*puntero).dato
	if (*puntero).izquierdo == nil && (*puntero).derecho == nil {
		*puntero = nil
	} else if (*puntero).izquierdo == nil {
		*puntero = (*puntero).derecho
	} else if (*puntero).derecho == nil {
		*puntero = (*puntero).izquierdo
	} else {
		claveMin, valorMin := hallarMaximoRamaIzquierda(ab.raiz)
		(*puntero).clave = claveMin
		(*puntero).dato = valorMin
		ab.Borrar(claveMin)
	}
	ab.cantidad--
	return datoBorrado
}

func (ab *abb[K, V]) Cantidad() int {
	return ab.cantidad
}

func (ab *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	if ab.raiz == nil {
		return
	}

}

func (ab *abb[K, V]) Iterador() IterDiccionario[K, V] {

}

func (iter *iterAbb[K, V]) HaySiguiente() bool {

}

func (iter *iterAbb[K, V]) VerActual() (K, V) {

}

func (iter *iterAbb[K, V]) Siguiente() {

}

func (ab *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	if ab.raiz == nil {
		return
	}

}

func (ab *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {

}

func hallarMaximoRamaIzquierda[K comparable, V any](raiz *nodoAbb[K, V]) (K, V) {
	if raiz.izquierdo == nil && raiz.derecho == nil {
		return raiz.clave, raiz.dato
	}
	if raiz.izquierdo == nil {
		return hallarMaximoRamaIzquierda(raiz.derecho)
	}
	return hallarMaximoRamaIzquierda(raiz.izquierdo)
}

func buscarPuntero[K comparable, V any](raiz *nodoAbb[K, V], cmp func(K, K) int, clave K) **nodoAbb[K, V] {
	if cmp(clave, raiz.clave) < 0 {
		if raiz.izquierdo != nil {
			return buscarPuntero(raiz.izquierdo, cmp, clave)
		} else {
			return &raiz.izquierdo
		}
	} else if cmp(clave, raiz.clave) > 0 {
		if raiz.derecho != nil {
			return buscarPuntero(raiz.derecho, cmp, clave)
		} else {
			return &raiz.derecho
		}
	}
	return &raiz
}
