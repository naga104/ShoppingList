package routes

import "github.com/gin-gonic/gin"

//UserGroup has signup,login,logout
func UserGroup(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("signup", signupController.AddUser)
		userGroup.POST("login", signupController.Login)
		userGroup.POST("logout", signupController.Logout)
	}
}
