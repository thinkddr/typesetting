// Fork de github.com/go-text/typesetting para Sytue One.
//
// La ruta del módulo cambia porque un `replace` de Go EXIGE que el módulo de
// destino declare la ruta por la que se le requiere. Los ficheros siguen
// importando `github.com/go-text/typesetting/...` a propósito: el `replace`
// del módulo principal los resuelve aquí, y así el diff contra upstream es
// exactamente el parche y nada más — que es lo que permite rebasar.
//
// Único cambio funcional: `shaping/shaping.go` conforma a ppem FRACCIONARIO en
// vez de a `ceil`. El porqué, en ADR 0048 de sytue/one.
module github.com/thinkddr/typesetting

go 1.19

require (
	github.com/go-text/typesetting-utils v0.0.0-20260223113751-2d88ac90dae3
	golang.org/x/image v0.23.0
	golang.org/x/text v0.21.0
)
