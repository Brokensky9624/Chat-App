package line

import "example/homework/chatapp/service/web"

func RegistryRouter() {
	webMngr := web.GetWebManager()
	router := webMngr.WebEngine()
	group := router.Group("/line")
	group.POST("/callback", apiCallBack)
}
