import "github.com/lyft/odie/client/mongo"

sess, _ := imgo.Dial("localhost")
mgr := mongo.New(sess)
defer mgr.Close()

c := mgr.Client()
defer c.Close()

var user User
err := client.Get().Key("ID", 1).Table("users").Into(&user).Select("Name").Response()

fmt.Println(user.Name) => "Other Chris"