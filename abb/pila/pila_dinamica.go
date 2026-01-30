package pila

const (
	_TAM_INICIAL                = 10
	_FACTOR_REDIMENSION_AMPLIAR = 2
	_FACTOR_REDIMENSION_REDUCIR = 4
)

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func (p *pilaDinamica[T]) redimensionar(nuevaCapacidad int) {
	nuevosDatos := make([]T, nuevaCapacidad)
	copy(nuevosDatos, p.datos)
	p.datos = nuevosDatos
}

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{
		datos:    make([]T, _TAM_INICIAL),
		cantidad: 0,
	}
}

func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p *pilaDinamica[T]) Apilar(valor T) {
	if p.cantidad == cap(p.datos) {
		p.redimensionar(_FACTOR_REDIMENSION_AMPLIAR * cap(p.datos))
	}
	p.datos[p.cantidad] = valor
	p.cantidad++
}

func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	if p.cantidad*_FACTOR_REDIMENSION_REDUCIR <= cap(p.datos) && cap(p.datos)/_FACTOR_REDIMENSION_AMPLIAR >= _TAM_INICIAL {
		p.redimensionar(cap(p.datos) / _FACTOR_REDIMENSION_AMPLIAR)
	}
	p.cantidad--
	return p.datos[p.cantidad]
}

func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	return p.datos[p.cantidad-1]
}
