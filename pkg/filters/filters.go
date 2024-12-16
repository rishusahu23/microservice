package filters

import "go.mongodb.org/mongo-driver/bson"

type FilterOption interface {
	// Apply applies the filter to the given bson.M (MongoDB query filter).
	Apply(filter bson.M) bson.M
}

type funcMongoFilterOption struct {
	fn func(filter bson.M) bson.M
}

func NewFuncMongoFilterOption(fn func(filter bson.M) bson.M) *funcMongoFilterOption {
	return &funcMongoFilterOption{fn: fn}
}

func (f *funcMongoFilterOption) Apply(filter bson.M) bson.M {
	return f.fn(filter)
}
