import "github.com/lyft/odie/client/dynamo"

db := dynamodb.New(session)
c := client.New(db, config.Timeout(25*time.Millisecond))

var user User
err := client.Get().Key("ID", 1).Table("users").Into(&user).Select("Name").Response()

fmt.Println(user.Name) => "Chris"