package pila_test

import (
	"github.com/stretchr/testify/require"
	TDAPila "tdas/pila"
	"testing"
)

func TestPilaVacia(t *testing.T) {
	pilaInt := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pilaInt.EstaVacia())

	pilaString := TDAPila.CrearPilaDinamica[string]()
	require.True(t, pilaString.EstaVacia())
}

func TestPilaLIFO(t *testing.T) {
	pilaInt := TDAPila.CrearPilaDinamica[int]()
	pilaInt.Apilar(1)
	require.Equal(t, 1, pilaInt.VerTope())
	pilaInt.Apilar(2)
	require.Equal(t, 2, pilaInt.VerTope())
	pilaInt.Apilar(3)
	require.Equal(t, 3, pilaInt.VerTope())

	require.Equal(t, 3, pilaInt.Desapilar())
	require.Equal(t, 2, pilaInt.VerTope())
	require.Equal(t, 2, pilaInt.Desapilar())
	require.Equal(t, 1, pilaInt.VerTope())
	require.Equal(t, 1, pilaInt.Desapilar())
	require.True(t, pilaInt.EstaVacia())

	pilaString := TDAPila.CrearPilaDinamica[string]()
	pilaString.Apilar("uno")
	require.Equal(t, "uno", pilaString.VerTope())
	pilaString.Apilar("dos")
	require.Equal(t, "dos", pilaString.VerTope())
	pilaString.Apilar("tres")
	require.Equal(t, "tres", pilaString.VerTope())

	require.Equal(t, "tres", pilaString.Desapilar())
	require.Equal(t, "dos", pilaString.VerTope())
	require.Equal(t, "dos", pilaString.Desapilar())
	require.Equal(t, "uno", pilaString.VerTope())
	require.Equal(t, "uno", pilaString.Desapilar())
	require.True(t, pilaString.EstaVacia())
}

func TestPilaVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	const n = 10000

	for i := 0; i < n; i++ {
		require.NotPanics(t, func() { pila.Apilar(i) })
		require.Equal(t, i, pila.VerTope())
		require.Equal(t, false, pila.EstaVacia())
	}

	for i := n - 1; i >= 0; i-- {
		require.False(t, pila.EstaVacia())
		require.Equal(t, i, pila.VerTope())
		require.Equal(t, i, pila.Desapilar())
	}

	require.True(t, pila.EstaVacia())
	require.Panics(t, func() { pila.Desapilar() })
	require.Panics(t, func() { pila.VerTope() })
}

func TestCondicionesDeBorde(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())

	// Al crear la pila
	require.Panics(t, func() { pila.Desapilar() })
	require.Panics(t, func() { pila.VerTope() })

	pila.Apilar(3)
	require.Equal(t, 3, pila.VerTope())
	pila.Apilar(2)
	require.Equal(t, 2, pila.VerTope())

	require.Equal(t, 2, pila.Desapilar())
	require.Equal(t, 3, pila.VerTope())
	require.Equal(t, 3, pila.Desapilar())
	require.True(t, pila.EstaVacia())

	// Luego de usarla
	require.Panics(t, func() { pila.Desapilar() })
	require.Panics(t, func() { pila.VerTope() })
}

func TestPilaTiposGenericos(t *testing.T) {

	pilaInt := TDAPila.CrearPilaDinamica[int]()
	pilaInt.Apilar(42)
	require.Equal(t, 42, pilaInt.VerTope())

	pilaString := TDAPila.CrearPilaDinamica[string]()
	pilaString.Apilar("Go")
	require.Equal(t, "Go", pilaString.VerTope())

	pilaBool := TDAPila.CrearPilaDinamica[bool]()
	pilaBool.Apilar(true)
	require.Equal(t, true, pilaBool.VerTope())
}
