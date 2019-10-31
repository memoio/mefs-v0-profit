package controllers
import(
	"github.com/kataras/iris"

	rm "longchain.com/memoriae/profit/web/resultModel"
	"longchain.com/memoriae/profit/web/service"
)

func New(app *iris.Application) {
	app.Get("/profit/{role}/{addr}", func(ctx iris.Context) {
		role := ctx.Params().Get("role")
		addr := ctx.Params().Get("addr")
		ctx.JSON(rm.Ok(service.Profit(role,addr)))
	})
}