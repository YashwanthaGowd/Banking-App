package handler

import (
	"net/http"

	"github.com/banking-app/account-service/src/model"

	accountpb "github.com/banking-app/protos/generated/account"

	"github.com/gin-gonic/gin"
)

//create handlers and make use of service to handle the business logic

// use gin framework

// example json request
// {
//   "first_name": "yash",
//   "last_name": "yash",
//   "email": "yash@gmail.com",
//   "type": "savings",
//   "password": "123456"
// }

func (h handler) CreateUser(c *gin.Context) {

	req := &accountpb.CreateUserRequest{}

	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.NewUserFromProto(req)
	err := h.BankingService.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

// GetUserbyEmail gets a user by email
func (h handler) GetUserbyEmail(c *gin.Context) {
	userEmail := c.Param("userEmail")
	user, err := h.BankingService.GetUserbyEmail(userEmail)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// updateUser updates an user
func (h handler) UpdateUser(c *gin.Context) {
	req := &accountpb.UpdateUserRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// find if user exists
	user, err := h.BankingService.GetUserbyEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}
	// update user
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Type = req.Type
	user.Password = req.Password

	err = h.BankingService.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// disableUser disables an user
func (h handler) DisableUser(c *gin.Context) {

	req := &accountpb.DisableUserRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// find if user exists
	user, err := h.BankingService.GetUserbyEmail(req.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}
	// update user
	if user.Status == "disabled" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user is already disabled"})
		return
	}
	user.Status = "disabled"

	err = h.BankingService.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User Disabled successfully"})
}

// activateUser activates an user
func (h handler) ActivateUser(c *gin.Context) {
	req := &accountpb.ActivateUserRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// find if user exists
	user, err := h.BankingService.GetUserbyEmail(req.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}
	// update user
	if user.Status == "active" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user is already active"})
		return
	}
	user.Status = "active"

	err = h.BankingService.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User Activated successfully"})
}
