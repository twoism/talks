import (
    "github.com/lyft/lyft-idl/generated/go/test_repo"
    "github.com/lyft/odie/client/mongo"

    ddb "github.com/lyft/lyft-idl/generated/go/test_repo/dynamo"
    mgo "github.com/lyft/lyft-idl/generated/go/test_repo/mongo"
)

dc := dynamo.New(dynamoSession, config.Timeout(25*time.Milliseconds))
sess, _ := imgo.Dial("localhost")
mc := mongo.New(sess).Client()
cfg := config.TableName("users")

dyn := ddb.New(dc, cfg)
mgo := mongo.New(mc, cfg)
ctx := context.Background()

var user test_repo.UserModel
err := mgo.Get(ctx, 1).Into(&user).Response()

err = dyn.Put(ctx, &user).Response()