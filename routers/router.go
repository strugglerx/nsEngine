package routers

import (
	"github.com/astaxie/beego/context"
	"server/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin", &controllers.MainController{},"get:Login;post:Login")
	beego.Router("/logout", &controllers.MainController{},"get:Logout")
	//API路由
	apiNs :=beego.NewNamespace("api",
		beego.NSBefore(func(ctx *context.Context) {
			ctx.Output.Header("Content-Type", "application/json;charset=UTF-8")
		}),
		beego.NSRouter("/",&controllers.ApiController{},"get:ApiIndex"),
		beego.NSRouter("msg",&controllers.ApiController{},"get:Msg;post:MsgPost"),
		beego.NSRouter("point",&controllers.ApiController{},"get:PointIndex;post:PointPost"),
		beego.NSRouter("indexswiper",&controllers.ApiController{},"get:IndexSwiper"),
		beego.NSRouter("indexconfig",&controllers.ApiController{},"get:IndexConfig"),
		beego.NSRouter("jobsdetail",&controllers.ApiController{},"get:JobDetail"),
		beego.NSRouter("jobslist",&controllers.ApiController{},"get:JobList"),
		beego.NSRouter("stepUpdate",&controllers.ApiController{},"post:StepUpdate"),
		beego.NSRouter("steplist",&controllers.ApiController{},"get:StepList"),
		beego.NSRouter("sports",&controllers.EipController{},"post:SportEntry"),
		beego.NSRouter("eip", &controllers.EipController{},"post:Entry"),
		beego.NSRouter("artlist", &controllers.ApiController{}),
		beego.NSRouter("artdetail", &controllers.ApiController{},"get:ArtDetail"),
		beego.NSRouter("artuplike", &controllers.ApiController{},"post:ArtUplike"),
		)
	//后台管理结构
	managerNs :=beego.NewNamespace("manager",
		beego.NSCond(func(ctx *context.Context) bool {
			role:=ctx.Input.Session("role")
			if role=="admin"{
				return true
			}
			return false
		}),
		beego.NSBefore(func(ctx *context.Context) {
			ctx.Output.Header("Content-Type", "application/json;charset=UTF-8")
		}),
		beego.NSRouter("feedback", &controllers.ManagerController{},"get:FeedBackList"),
		beego.NSRouter("artinsert", &controllers.ManagerController{},"post:ArtInsert"),
		beego.NSRouter("delete", &controllers.ManagerController{},"post:DbDelete"),
	)
	//注册 namespace
	beego.AddNamespace(apiNs)
	beego.AddNamespace(managerNs)
}
