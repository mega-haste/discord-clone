package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"upchat.com/server/model"
)

func Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "here from index api",
	})
}

func GetGlobalMessages(ctx *gin.Context) {
	messages := make([]map[string]interface{}, 0)
	model.DB.
		Model(&model.MessageModel{}).
		Order("id desc").
		Limit(25).
		Select("id", "Content", "from_id", "replying_to_id", "CreatedAt", "updated_at", "deleted_at").
		Find(&messages)
	ctx.JSON(http.StatusOK, messages)
}

func PostLogin(ctx *gin.Context) {
	var loginInfo struct { UserName string; Password string }
	if err := ctx.BindJSON(&loginInfo); err != nil {
		log.Println(err)
	}
	var dist model.UserModel
	res := model.DB.
		Where("user_name = ? AND password = ?", loginInfo.UserName, loginInfo.Password).
		First(&dist)
	
	if res.RowsAffected > 0 {
		token, err := genrateToken()
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H {
				"status": http.StatusInternalServerError,
				"msg": "Something bad happened when trying to create sometokens and stuff",
			})
			return
		}
		dist.Login_token = &token
		model.DB.Save(&dist)
		ctx.SetCookie("haha", "hghggg", 5484884, "", "", false, false)
		ctx.JSON(http.StatusOK, gin.H {
			"status": http.StatusOK,
			"msg": "Login went succesfuly.",
			"token": token,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H {
			"status": http.StatusBadRequest,
			"msg": "login didn't acomplish.",
		})
	}
}

func GetLogin(ctx *gin.Context)  {
	login_token := ctx.Query("login_token")
	if login_token == "" {
		ctx.JSON(http.StatusBadRequest, gin.H {
			"status": http.StatusBadRequest,
			"msg": "there is no login in that device for now.",
		})
		return
	}
	var dist model.UserModel
	res := model.DB. 
		Model(&model.UserModel{}).
		Where("login_token = ?", login_token).
		First(&dist)
	if res.RowsAffected > 0 { 
		ctx.JSON(http.StatusOK, gin.H {
			"status": http.StatusOK,
			"msg": "logged in!!",
			"data": dist,
		})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H {
		"status": http.StatusBadRequest,
		"msg": "Login data is invalid",
	})
}


func genrateToken() (string, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

