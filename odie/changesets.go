import (
	"github.com/lyft/golyft/middleware/odie/changes" // HL
	client "github.com/lyft/odie/client/mongo"
)

c := client.New(sess)
client.Decorate(
	changes.DefaultDecorator("user_property_change"),
)