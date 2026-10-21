package transaction

// Tx represents the transaction shared by repository operations.
type Tx interface {
	Commit() error
	Rollback() error
}
