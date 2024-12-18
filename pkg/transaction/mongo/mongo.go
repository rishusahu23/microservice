package mongo

import (
	"context"
	"github.com/google/wire"
	txn "github.com/rishu/microservice/pkg/transaction"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoTransactionManager implements TransactionManager for MongoDB
type MongoTransactionManager struct {
	client *mongo.Client
}

// NewMongoTransactionManager creates a new MongoTransactionManager
func NewMongoTransactionManager(client *mongo.Client) *MongoTransactionManager {
	return &MongoTransactionManager{client: client}
}

func (tm *MongoTransactionManager) RunInTxn(ctx context.Context, fn txn.InTxn) error {
	// Start a new session
	session, err := tm.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	// Run the transaction
	_, err = session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		return nil, fn(sessCtx)
	})
	return err
}

var _ txn.TransactionManager = &MongoTransactionManager{}

var (
	DefaultTxnManagerWireSet = wire.NewSet(NewMongoTransactionManager, wire.Bind(new(txn.TransactionManager), new(*MongoTransactionManager)))
)
