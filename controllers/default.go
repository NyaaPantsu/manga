package controllers

// MainController operations for main page
type MainController struct {
	BaseController
}

// Get ...
// @Title Get
// @Description get Main Page
// @Success 200 {object} models.ChapterGroup
// @Failure 403
// @router / [get
func (c *MainController) Get() {
	c.TplName = "home.html"
	c.Render()
}
