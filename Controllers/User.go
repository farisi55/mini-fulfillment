package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"theapp/Models"
)

// Get Users godoc
// @Summary Show an account
// @Description get string by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} Models.User
// @Router /userapi/user [get]
func GetUsers(c *gin.Context)  {
	var subs[]Models.Subs
	err := Models.GetAllUser(&subs)
	if err != nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		format := c.DefaultQuery("format", "json")
		if format == "json"{
			c.JSON(http.StatusOK, subs)
		}else{
			c.XML(http.StatusOK, subs)
		}
	}
}

// Create Users godoc
// @Summary Create an account
// @Description get string by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Success 200 {object} Models.User
// @Router /admin/user [POST]
func BuyPacket(c *gin.Context){
	var subs Models.Subs
	c.BindJSON(&subs)
	err := Models.PurchasePacket(&subs)
	if err != nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK, subs)
	}
}

func GetPacket(c *gin.Context)  {
	var subs Models.Subs
	id := c.Params.ByName("id")
	err := Models.GetSubsInfo(&subs, id)
	if err != nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK, subs)
	}
}

func UpdateUser(c *gin.Context)  {
	var subs Models.Subs
	id := c.Params.ByName("subsid")
	err := Models.GetSubsInfo(&subs, id)
	if err != nil{
		c.JSON(http.StatusNotFound, subs)
	}

	c.BindJSON(&subs)
	err = Models.UpdateUser(&subs, id)
	if err != nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK, subs)
	}
}

func RemovePacket(c *gin.Context)  {
	var subs Models.Subs
	id := c.Params.ByName("subsid")
	serviceid := c.Params.ByName("serviceid")
	fmt.Println("======", id, serviceid)
	err := Models.UnregPacket(&subs, id, serviceid)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"serviceid" + serviceid: "is deleted"})
	}
}