package routers

import (
	"server/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	//错误页处理404/405
	beego.ErrorController(&controllers.ErrorController{})
	//外部展示路由
	beego.Router("/", &controllers.MainController{}, "get:Index")
	beego.Router("/admin", &controllers.MainController{}, "get:Login;post:LoginPost")
	beego.Router("/logout", &controllers.MainController{}, "get:Logout")
	//API路由
	apiNs := beego.NewNamespace("api",
		beego.NSBefore(func(ctx *context.Context) {
			ctx.Output.Header("Content-Type", "application/json;charset=UTF-8")
		}),
		beego.NSRouter("/", &controllers.ApiController{}, "get:ApiIndex"),
		beego.NSRouter("msg", &controllers.ApiController{}, "get:Msg;post:MsgPost"),
		beego.NSRouter("point", &controllers.ApiController{}, "get:PointIndex;post:PointPost"),
		beego.NSRouter("indexswiper", &controllers.ApiController{}, "get:IndexSwiper"),
		beego.NSRouter("indexconfig", &controllers.ApiController{}, "get:IndexConfig"),
		beego.NSRouter("jobsdetail", &controllers.ApiController{}, "get:JobDetail"),
		beego.NSRouter("jobslist", &controllers.ApiController{}, "get:JobList"),
		beego.NSRouter("stepupdate", &controllers.ApiController{}, "post:StepUpdate"),
		beego.NSRouter("steplist", &controllers.ApiController{}, "get:StepList"),
		beego.NSRouter("library", &controllers.EipController{}, "get:Library"),
		beego.NSRouter("librarydetail", &controllers.EipController{}, "get:LibraryDetail"),
		beego.NSRouter("sports", &controllers.EipController{}, "post:SportEntry"),
		beego.NSRouter("eip", &controllers.EipController{}, "post:Entry"),
		beego.NSRouter("artlist", &controllers.ApiController{}, "get:ArtList"),
		beego.NSRouter("artdetail", &controllers.ApiController{}, "get:ArtDetail"),
		beego.NSRouter("artuplike", &controllers.ApiController{}, "post:ArtUplike"),
		beego.NSRouter("adlist", &controllers.ApiController{}, "get:AdList"),
		beego.NSRouter("collectFormId", &controllers.ApiController{}, "post:CollectFormId"),
		beego.NSRouter("classswitch", &controllers.ApiController{}, "post:Classswitch"),
	)
	//后台管理结构
	managerNs := beego.NewNamespace("manager",
		beego.NSCond(func(ctx *context.Context) bool {
			role := ctx.Input.Session("role")
			if role == 1 {
				return true
			}
			//生产环境
			return false
			//开发环境
			// return true
		}),
		beego.NSBefore(func(ctx *context.Context) {
			ctx.Output.Header("Content-Type", "application/json;charset=UTF-8")
		}),
		beego.NSRouter("/", &controllers.ManagerController{}, "get:ManagerIndex"),
		beego.NSRouter("info", &controllers.ManagerController{}, "get:ManagerInfo"),
		beego.NSRouter("feedback", &controllers.ManagerController{}, "get:FeedBackList;post:FeedBackAction"),
		beego.NSRouter("artinsert", &controllers.ManagerController{}, "post:ArtInsert"),
		beego.NSRouter("delete", &controllers.ManagerController{}, "post:DbDelete"),
		beego.NSRouter("update", &controllers.ManagerController{}, "post:Update"),
		beego.NSRouter("option", &controllers.ManagerController{}, "post:Option"),
		beego.NSRouter("changePwd", &controllers.ManagerController{}, "post:ChangePwd"),
		beego.NSRouter("adinsert", &controllers.ManagerController{}, "post:AdInsert"),
		beego.NSRouter("adlist", &controllers.ManagerController{}, "get:AdList"),
		beego.NSRouter("keywordlist", &controllers.ManagerController{}, "get:KeywordList"),
		beego.NSRouter("keywordinsert", &controllers.ManagerController{}, "post:KeywordInsert"),
		beego.NSRouter("formIdList", &controllers.ManagerController{}, "post:FormId"),
		beego.NSRouter("pushToast", &controllers.ManagerController{}, "post:PushToast"),
	)
	//注册 namespace
	beego.AddNamespace(apiNs)
	beego.AddNamespace(managerNs)
}
