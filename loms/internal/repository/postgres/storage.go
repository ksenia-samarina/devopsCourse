package postgres

type Storage struct {
	QueryEngineProvider QueryEngineProvider
}

func New(queryEngineProvider QueryEngineProvider) *Storage {
	return &Storage{
		QueryEngineProvider: queryEngineProvider,
	}
}
