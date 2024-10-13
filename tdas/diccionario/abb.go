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
		puntero := buscarPuntero(ab, clave)
		*puntero = nuevaHoja
	}
	ab.cantidad++
}

func (ab *abb[K, V]) Pertenece(clave K) bool {
	puntero := buscarPuntero(ab, clave)
	return *puntero != nil
}

func (ab *abb[K, V]) Obtener(clave K) V {
	puntero := buscarPuntero(ab, clave)
	if *puntero == nil {
		panic("La clave no pertenece al diccionario")
	}
	return (*puntero).dato
}

func (ab *abb[K, V]) Borrar(clave K) V {
	puntero := buscarPuntero(ab, clave)
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
		claveMin, valorMin := hallarMinimo(ab.raiz)
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

func hallarMinimo[K comparable, V any](raiz *nodoAbb[K, V]) (K, V) {
	if raiz.izquierdo == nil && raiz.derecho == nil {
		return raiz.clave, raiz.dato
	}
	if raiz.izquierdo == nil {
		return hallarMinimo(raiz.derecho)
	}
	return hallarMinimo(raiz.izquierdo)
}

func buscarPuntero[K comparable, V any](ab *abb[K, V], clave K) **nodoAbb[K, V] {
	if ab.cmp(clave, ab.raiz.clave) < 0 {
		if ab.raiz.izquierdo != nil {
			return buscarPuntero(&abb[K, V]{ab.raiz.izquierdo, ab.cantidad, ab.cmp}, clave)
		} else {
			return &ab.raiz.izquierdo
		}
	} else if ab.cmp(clave, ab.raiz.clave) > 0 {
		if ab.raiz.derecho != nil {
			return buscarPuntero(&abb[K, V]{ab.raiz.derecho, ab.cantidad, ab.cmp}, clave)
		} else {
			return &ab.raiz.derecho
		}
	}
	return &ab.raiz
}
