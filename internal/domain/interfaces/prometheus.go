package interfaces

type Prometheus interface {
	CounterRequestStatusCode(router, protocol string, statusCode int)
}
