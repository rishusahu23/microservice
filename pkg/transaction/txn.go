package txn

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type InTxn func(sessCtx mongo.SessionContext) error

// TransactionManager defines the interface for transaction execution
type TransactionManager interface {
	RunInTxn(ctx context.Context, fn InTxn) error
}
