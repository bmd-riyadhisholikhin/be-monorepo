// Code generated by candi v1.5.32.

package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	shareddomain "monorepo/services/shark/pkg/shared/domain"

	"pkg.agungdp.dev/candi/candihelper"
	"pkg.agungdp.dev/candi/candishared"
	"pkg.agungdp.dev/candi/tracer"
)

type contactRepoMongo struct {
	readDB, writeDB *mongo.Database
	collection      string
}

// NewContactRepoMongo mongo repo constructor
func NewContactRepoMongo(readDB, writeDB *mongo.Database) ContactRepository {
	return &contactRepoMongo{
		readDB, writeDB, "contacts",
	}
}

func (r *contactRepoMongo) FetchAll(ctx context.Context, filter *candishared.Filter) (data []shareddomain.Contact, err error) {
	trace := tracer.StartTrace(ctx, "ContactRepoMongo:FetchAll")
	defer func() { trace.SetError(err); trace.Finish() }()
	ctx = trace.Context()

	where := bson.M{}
	trace.SetTag("query", where)

	findOptions := options.Find()
	if len(filter.OrderBy) > 0 {
		findOptions.SetSort(filter)
	}

	if !filter.ShowAll {
		findOptions.SetLimit(int64(filter.Limit))
		findOptions.SetSkip(int64(filter.Offset))
	}
	cur, err := r.readDB.Collection(r.collection).Find(ctx, where, findOptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	cur.All(ctx, &data)
	return
}

func (r *contactRepoMongo) Find(ctx context.Context, data *shareddomain.Contact) (err error) {
	trace := tracer.StartTrace(ctx, "ContactRepoMongo:Find")
	defer func() { trace.SetError(err); trace.Finish() }()
	ctx = trace.Context()

	bsonWhere := make(bson.M)
	if data.ID != "" {
		bsonWhere["_id"] = data.ID
	}
	trace.SetTag("query", bsonWhere)

	return r.readDB.Collection(r.collection).FindOne(ctx, bsonWhere).Decode(data)
}

func (r *contactRepoMongo) Count(ctx context.Context, filter *candishared.Filter) int {
	trace := tracer.StartTrace(ctx, "ContactRepoMongo:Count")
	defer trace.Finish()

	where := bson.M{}
	count, err := r.readDB.Collection(r.collection).CountDocuments(trace.Context(), where)
	trace.SetError(err)
	return int(count)
}

func (r *contactRepoMongo) Save(ctx context.Context, data *shareddomain.Contact) (err error) {
	trace := tracer.StartTrace(ctx, "ContactRepoMongo:Save")
	defer func() { trace.SetError(err); trace.Finish() }()
	ctx = trace.Context()
	tracer.Log(ctx, "data", data)

	data.ModifiedAt = time.Now()
	if data.ID == "" {
		data.ID = primitive.NewObjectID().Hex()
		data.CreatedAt = time.Now()
		_, err = r.writeDB.Collection(r.collection).InsertOne(ctx, data)
	} else {
		opt := options.UpdateOptions{
			Upsert: candihelper.ToBoolPtr(true),
		}
		_, err = r.writeDB.Collection(r.collection).UpdateOne(ctx,
			bson.M{
				"_id": data.ID,
			},
			bson.M{
				"$set": data,
			}, &opt)
	}

	return
}

func (r *contactRepoMongo) Delete(ctx context.Context, id string) (err error) {
	trace := tracer.StartTrace(ctx, "ContactRepoMongo:Save")
	defer func() { trace.SetError(err); trace.Finish() }()

	_, err = r.writeDB.Collection(r.collection).DeleteOne(ctx, bson.M{"_id": id})
	return
}
