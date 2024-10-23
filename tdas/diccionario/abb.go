package diccionario

import (
	TDAPila "tdas/pila"
)

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
	arbol *abb[K, V]
	pila  TDAPila.Pila[*nodoAbb[K, V]]
	desde *K
	hasta *K
}

func crearNodo[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{nil, nil, clave, dato}
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{nil, 0, funcion_cmp}
}

func (ab *abb[K, V]) Guardar(clave K, dato V) {
	if ab.raiz == nil {
		ab.raiz = crearNodo(clave, dato)
	} else {
		puntero := buscarPuntero(ab.raiz, ab.cmp, clave)
		if *puntero == nil {
			*puntero = crearNodo(clave, dato)
		} else {
			(*puntero).dato = dato
		}

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
	if (*puntero).izquierdo == nil {
		*puntero = (*puntero).derecho
	} else if (*puntero).derecho == nil {
		*puntero = (*puntero).izquierdo
	} else {
		claveMax, valorMax := hallarReemplazo((*puntero).izquierdo)
		(*puntero).clave = claveMax
		(*puntero).dato = valorMax
		return ab.Borrar(claveMax)
	}
	ab.cantidad--
	return datoBorrado

	/*
		(*puntero).izquierdo == nil && (*puntero).derecho == nil {
			*puntero = nil
		} else if
	*/
}

func (ab *abb[K, V]) Cantidad() int {
	return ab.cantidad
}

func (ab *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	iterarRango(ab.raiz, nil, nil, ab.cmp, visitar)
}

func (ab *abb[K, V]) Iterador() IterDiccionario[K, V] {
	pila := TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	apilarClaves(ab.raiz, nil, nil, ab.cmp, pila)
	return &iterAbb[K, V]{ab, pila, nil, nil}
}

func (iter *iterAbb[K, V]) HaySiguiente() bool {
	return !iter.pila.EstaVacia()
}

func (iter *iterAbb[K, V]) VerActual() (K, V) {
	return iter.pila.VerTope().clave, iter.pila.VerTope().dato
}

func (iter *iterAbb[K, V]) Siguiente() {
	nodo := iter.pila.Desapilar()
	if nodo.derecho != nil {
		apilarClaves(nodo.derecho, iter.desde, iter.hasta, iter.arbol.cmp, iter.pila)
	}
}

func (ab *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	iterarRango(ab.raiz, desde, hasta, ab.cmp, visitar)
}

func (ab *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	pila := TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	apilarClaves(ab.raiz, desde, hasta, ab.cmp, pila)
	return &iterAbb[K, V]{ab, pila, desde, hasta}
}

func hallarReemplazo[K comparable, V any](nodo *nodoAbb[K, V]) (K, V) {
	if nodo.izquierdo == nil && nodo.derecho == nil {
		return nodo.clave, nodo.dato
	}
	if nodo.derecho != nil {
		return hallarReemplazo(nodo.derecho)
	}
	return hallarReemplazo(nodo.izquierdo)

	/*
		if raiz.izquierdo == nil && raiz.derecho == nil {
			return raiz.clave, raiz.dato
		}
		if raiz.izquierdo == nil {
			return hallarMaximoRamaIzquierda(raiz.derecho)
		}
		return hallarMaximoRamaIzquierda(raiz.izquierdo)
	*/
}

func buscarPuntero[K comparable, V any](nodo *nodoAbb[K, V], cmp func(K, K) int, clave K) **nodoAbb[K, V] {
	if nodo == nil || cmp(nodo.clave, clave) == 0 {
		return &nodo
	}
	if cmp(nodo.clave, clave) > 0 {
		return buscarPuntero(nodo.izquierdo, cmp, clave)
	} else {
		return buscarPuntero(nodo.derecho, cmp, clave)
	}
	/*if nodo == nil || cmp(nodo.clave, clave) == 0 {
		return &nodo
	}
	if cmp(nodo.clave, clave) > 0 {
		if nodo.izquierdo != nil {
			return buscarPuntero(nodo.izquierdo, cmp, clave)
		} else {
			return &nodo.izquierdo
		}
	} else {
		if nodo.derecho != nil {
			return buscarPuntero(nodo.derecho, cmp, clave)
		} else {
			return &nodo.derecho
		}
	}
	*/
}

func iterarRango[K comparable, V any](nodo *nodoAbb[K, V], desde *K, hasta *K, cmp func(K, K) int, visitar func(clave K, dato V) bool) bool {
	if nodo == nil {
		return true
	}
	if desde == nil || cmp(nodo.clave, *desde) > 0 {
		if !iterarRango(nodo.izquierdo, desde, hasta, cmp, visitar) {
			return false
		}
	}
	if (desde == nil || cmp(nodo.clave, *desde) >= 0) && (hasta == nil || cmp(nodo.clave, *hasta) <= 0) {
		if !visitar(nodo.clave, nodo.dato) {
			return false
		}
	}
	if hasta == nil || cmp(nodo.clave, *hasta) < 0 {
		if !iterarRango(nodo.derecho, desde, hasta, cmp, visitar) {
			return false
		}
	}
	return true
}

func apilarClaves[K comparable, V any](nodo *nodoAbb[K, V], desde *K, hasta *K, cmp func(K, K) int, pila TDAPila.Pila[*nodoAbb[K, V]]) {
	if nodo == nil {
		return
	}
	if (desde == nil || cmp(nodo.clave, *desde) >= 0) && (hasta == nil || cmp(nodo.clave, *hasta) <= 0) {
		pila.Apilar(nodo)
	}
	if desde == nil || cmp(nodo.clave, *desde) > 0 {
		apilarClaves(nodo.izquierdo, desde, hasta, cmp, pila)
	}

}
