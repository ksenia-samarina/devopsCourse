package order

type Domain struct {
	storage            Storage
	transactionManager TransactionManager
}

func New(storage Storage, transactionManager TransactionManager) *Domain {
	return &Domain{
		storage:            storage,
		transactionManager: transactionManager,
	}
}
