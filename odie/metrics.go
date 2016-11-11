import (
	"github.com/lyft/golyft/stats"
	golyft_stats "github.com/lyft/golyft/stats/odie"
	odie_stats "github.com/lyft/odie/client/decorators/stats"
	client "github.com/lyft/odie/client/mongo"
)

c := client.New(sess)
scope := stats.NewDefaultStore().Scope("odie_test")
c.Decorate(
	odie_stats.Decorator(golyft_stats.New(scope)),
)