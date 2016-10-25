package controllers

type UserController struct {
	BaseController
}

func (this *UserController) Get() {
	action := this.GetString(":action")
	switch action {
	case "list":
		this.List()
	default:
		this.List()

	}
}
func (this *UserController) List() {
	this.Layout = "base/base.html"
	this.TplName = "user/user_list.html"
}
