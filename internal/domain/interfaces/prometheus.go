package interfaces

type Prometheus interface {
	CounterRequestStatusCode(router, protocol string, statusCode int)
	ObserveInstructionDBDuration(database, table, method string, duration float64)
	ObserveRequestDuration(router, protocol string, statusCode int, duration float64)
}
