package controllers

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/gin-gonic/gin"
	"github.com/shadab/avasapi/forms"
	"github.com/shadab/avasapi/model"
)

var userModel = new(model.UserModel)

// UserController Struct
type UserController struct{}

//Create a user
func (user *UserController) Create(c *gin.Context) {
	var data forms.CreateUserCommand
	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{
			"message": "Invalid form",
			"form":    data,
		})
		c.Abort()
		return
	}

	err := userModel.Create(data)
	if err != nil {
		c.JSON(406, gin.H{
			"message": "user could not be created",
			"error":   err,
		})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

//AddConf a user
func (user *UserController) AddConf(c *gin.Context) {
	var body forms.GroupConf
	if c.BindJSON(&body) != nil {
		c.JSON(406, gin.H{
			"message": "Invalid form",
			"form":    body,
		})
		c.Abort()
		return
	}
	err := userModel.Add(body)
	if err != nil {
		c.JSON(406, gin.H{
			"message": "Group Conf Could not be add",
			"error":   err,
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"message": "Group Conf added success",
	})
}

//UpdateGroup user
func (user *UserController) UpdateGroup(c *gin.Context) {
	id := c.Param("id")
	data := forms.GroupConf{}

	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{
			"message": "Invalid Parameters",
		})
		c.Abort()
		return
	}

	err := userModel.UpdateConf(id, data)
	if err != nil {
		c.JSON(406, gin.H{
			"message": "Update Group Failed",
		})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"message": "Update Group Success",
	})
}

//DeleteGroup user
func (user *UserController) DeleteGroup(c *gin.Context) {
	id := c.Param("id")
	data := forms.DeleteConf{}

	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{
			"message": "Invalid Parameters",
		})
		c.Abort()
		return
	}

	err := userModel.DeleteConf(id, data)
	if err != nil {
		c.JSON(406, gin.H{
			"message": "Delete Group Failed",
		})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"message": "Delete Group Success",
	})

}

//GetAll users
func (user *UserController) GetAll(c *gin.Context) {
	list, err := userModel.GetAll()
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Find Error",
			"error":   err.Error(),
		})
		c.Abort()
	} else {
		c.JSON(200, gin.H{
			"data": list,
		})
	}
}

//Forget API
// func (user *UserController) Forget(c *gin.Context) {
// 	id := c.Param("id")
// 	data := forms.ForgetAPI{}
// 	if c.BindJSON(&data) != nil {
// 		c.JSON(406, gin.H{
// 			"message": "Invalid form",
// 			"form":    data,
// 		})
// 		c.Abort()
// 		return
// 	}
// 	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
// 	result, err := userModel.Forget(data.Email, seededRand)
// }

//GetMD5Hash func
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

//Login user
func (user *UserController) Login(c *gin.Context) {
	data := forms.LoginCommand{}
	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{
			"message": "Invalid form",
			"form":    data,
		})
		c.Abort()
		return
	}
	result, err := userModel.Login(data.Email)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Find User Error",
			"error":   err.Error(),
		})
		c.Abort()
	} else {
		if result.Password == GetMD5Hash(data.Password) {
			c.JSON(200, gin.H{
				"200":  "Login Successful",
				"data": result,
			})
		} else {
			c.JSON(404, gin.H{
				"404": "Login Failed",
			})
		}
	}
}

//GetOne user
func (user *UserController) GetOne(c *gin.Context) {
	id := c.Param("id")
	result, err := userModel.GetOne(id)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Find User Error",
			"error":   err.Error(),
		})
		c.Abort()
	} else {
		c.JSON(200, gin.H{
			"data": result,
		})
	}
}

//UpdateOne user
func (user *UserController) UpdateOne(c *gin.Context) {
	id := c.Param("id")
	data := forms.UpdateUserCommand{}

	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{
			"message": "Invalid Parameters",
		})
		c.Abort()
		return
	}

	err := userModel.UpdateOne(id, data)
	if err != nil {
		c.JSON(406, gin.H{
			"message": "Update Failed",
		})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"message": "Update Success",
	})

}

//DeleteOne user
func (user *UserController) DeleteOne(c *gin.Context) {
	id := c.Param("id")
	err := userModel.DeleteOne(id)
	if err != nil {
		c.JSON(406, gin.H{
			"message": "Delete Failed",
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"message": "Delete Success",
	})

}
