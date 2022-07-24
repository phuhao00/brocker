package mongobrocker

import (
	"bufio"
	"bytes"
	"context"

	jsoniter "github.com/json-iterator/go"

	"go.mongodb.org/mongo-driver/x/bsonx"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetCOll(client *mongo.Client, dbName, collName string) *mongo.Collection {
	collection := client.Database(dbName).Collection(collName)
	return collection
}

func InsertOne(ctx context.Context, coll *mongo.Collection, data interface{}) (*mongo.InsertOneResult, error) {
	res, err := coll.InsertOne(ctx, data)
	return res, err
}

func InsertMany(ctx context.Context, coll *mongo.Collection, data []interface{}) (*mongo.InsertManyResult, error) {

	res, err := coll.InsertMany(ctx, data)
	return res, err
}

func FindOne(ctx context.Context, coll *mongo.Collection, filter interface{}) *mongo.SingleResult {

	return coll.FindOne(ctx, filter)
}

func Find(ctx context.Context, coll *mongo.Collection, filter interface{}) (*mongo.Cursor, error) {

	return coll.Find(ctx, filter)
}

func FindWithOption(ctx context.Context, coll *mongo.Collection, filter interface{},
	findOptions *options.FindOptions) (*mongo.Cursor, error) {

	return coll.Find(ctx, filter, findOptions)
}

func Distinct(ctx context.Context, coll *mongo.Collection, fieldName string, filter interface{}) ([]interface{}, error) {

	return coll.Distinct(ctx, fieldName, filter)
}

func UpdateOne(ctx context.Context, coll *mongo.Collection, filter interface{}, data interface{}) (*mongo.UpdateResult, error) {

	return coll.UpdateOne(ctx, filter, data)
}

func UpdateMany(ctx context.Context, coll *mongo.Collection, filter interface{}, data interface{}) (*mongo.UpdateResult, error) {

	return coll.UpdateMany(ctx, filter, data)
}

func UpdateByID(ctx context.Context, coll *mongo.Collection, id interface{}, data interface{}) (*mongo.UpdateResult, error) {

	return coll.UpdateByID(ctx, id, data)
}

func UpdateOneWithSession(ctx context.Context, client *mongo.Client, coll *mongo.Collection, filter interface{}, data interface{}) error {

	var (
		session mongo.Session
		err     error
	)
	session, err = client.StartSession()
	if err != nil {
		return err
	}
	if err = session.StartTransaction(); err != nil {
		return err
	}
	f := func(sessionContext mongo.SessionContext) error {
		_, err = coll.UpdateOne(sessionContext, filter, data)
		if err != nil {
			return err
		}
		err = session.CommitTransaction(sessionContext)
		if err != nil {
			return err
		}
		return nil
	}
	err = mongo.WithSession(ctx, session, f)
	if err != nil {
		return err
	}
	session.EndSession(ctx)
	return nil
}

func UpdateManyWithSession(ctx context.Context, client *mongo.Client, coll *mongo.Collection, filter interface{}, data interface{}) error {

	var (
		session mongo.Session
		err     error
	)
	session, err = client.StartSession()
	if err != nil {
		return err
	}
	if err = session.StartTransaction(); err != nil {
		return err
	}
	f := func(sessionContext mongo.SessionContext) error {
		_, err = coll.UpdateMany(sessionContext, filter, data)
		if err != nil {
			return err
		}
		err = session.CommitTransaction(sessionContext)
		if err != nil {
			return err
		}
		return nil
	}
	err = mongo.WithSession(ctx, session, f)
	if err != nil {
		return err
	}
	session.EndSession(ctx)
	return nil
}

func UpdateByIDWithSession(ctx context.Context, client *mongo.Client, coll *mongo.Collection, id interface{}, data interface{}) error {

	var (
		session mongo.Session
		err     error
	)
	session, err = client.StartSession()
	if err != nil {
		return err
	}
	if err = session.StartTransaction(); err != nil {
		return err
	}
	f := func(sessionContext mongo.SessionContext) error {
		_, err = coll.UpdateByID(sessionContext, id, data)
		if err != nil {
			return err
		}
		err = session.CommitTransaction(sessionContext)
		if err != nil {
			return err
		}
		return nil
	}
	err = mongo.WithSession(ctx, session, f)
	if err != nil {
		return err
	}
	session.EndSession(ctx)
	return nil
}

func DeleteOne(ctx context.Context, coll *mongo.Collection, filter interface{}) (*mongo.DeleteResult, error) {

	return coll.DeleteOne(ctx, filter)
}

func DeleteMany(ctx context.Context, coll *mongo.Collection, filter interface{}) (*mongo.DeleteResult, error) {

	return coll.DeleteMany(ctx, filter)
}

func Count(ctx context.Context, coll *mongo.Collection, filter interface{}) (int64, error) {

	return coll.CountDocuments(ctx, filter)
}

func ChangeStreamClient(client *mongo.Client, coll *mongo.Collection) {
	//
}

func ChangeStreamCollection(client *mongo.Client, coll *mongo.Collection) {
	//
}

func ChangeStreamDB(client *mongo.Client, coll *mongo.Collection) {
	//
}

//UploadGridFS ...
func UploadGridFS(ctx context.Context, filename string, data interface{}, db *mongo.Database, bucketOptions *options.BucketOptions) error {
	bucket, err := gridfs.NewBucket(db, bucketOptions)
	if err != nil {
		return err
	}
	opts := options.GridFSUpload()
	opts.SetMetadata(bsonx.Doc{{Key: "content-type", Value: bsonx.String("application/json")}})
	var upLoadStream *gridfs.UploadStream
	if upLoadStream, err = bucket.OpenUploadStream(filename, opts); err != nil {
		return err
	}
	str, err := jsoniter.MarshalToString(data)
	if err != nil {
		return err
	}
	if _, err = upLoadStream.Write([]byte(str)); err != nil {
		return err
	}
	upLoadStream.Close()
	return nil
}

//DownLoadGridFS ...
func DownLoadGridFS(ctx context.Context, fileID interface{}, db *mongo.Database, bucketOptions *options.BucketOptions) (string, error) {
	bucket, err := gridfs.NewBucket(db, bucketOptions)
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	if _, err = bucket.DownloadToStream(fileID, w); err != nil {
		return "", err
	}
	return b.String(), err
}
