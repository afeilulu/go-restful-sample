package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"afeilulu.com/example/database"
	"afeilulu.com/example/model"
	"afeilulu.com/example/pagination"
)

// CreateGroup godoc
// @Summary Create a new group
// @Description
// @Tags groups
// @Accept  json
// @Produce  json
// @Param group body model.Group true "Create group"
// @Success 200 {object} model.Group
// @Router /api/groups [post]
func CreateGroup(c *fiber.Ctx) error {
	db := database.DB
	g := new(model.Group)
	if err := c.BodyParser(g); err != nil {
		return err
	}
	g.ID = uuid.New()
	db.Create(&g)
	return c.JSON(g)
}

// GetGroups godoc
// @Summary Get paged of all
// @Description
// @Tags groups
// @Accept  json
// @Produce  json
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Param sort query string false "sort"
// @Success 200 {array} model.Group
// @Router /api/groups-paged [get]
func GetPagedGroups(c *fiber.Ctx) error {
	db := database.DB
	var pagiable pagination.Pagination
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	sort := c.Query("sort", "Id desc")

	pagiable.Limit = limit
	pagiable.Page = page
	pagiable.Sort = sort

	list := new([]model.Group)
	db.Preload("Users")

	db.Scopes(pagination.Paginate(list, &pagiable, db)).Find(&list)
	pagiable.Rows = list

	return c.JSON(pagiable)
}

// GetGroups godoc
// @Summary Get list of all
// @Description
// @Tags groups
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Group
// @Router /api/groups-list [get]
func GetGroups(c *fiber.Ctx) error {
	db := database.DB
	list := new([]model.Group)
	db.Preload("Users").Find(&list)
	return c.JSON(list)
}

// GetGroup godoc
// @Summary Get by id
// @Description
// @Tags groups
// @Accept  json
// @Produce  json
// @Param id path string true "ID of the group"
// @Success 200 {object} model.Group
// @Router /api/groups/{id} [get]
func GetGroup(c *fiber.Ctx) error {
	db := database.DB
	input := c.Params("id")
	var g model.Group
	db.Preload("Users").Find(&g, "id = ?", input)

	// If no such note present return an error
	if g.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No present", "data": nil})
	}

	return c.JSON(g)
}

// UpdateGroup godoc
// @Summary Update
// @Description Update the group with specified field
// @Tags groups
// @Accept  json
// @Produce  json
// @Param id path string true "ID of the group to be updated"
// @Param group body model.Group true "group"
// @Success 200 {object} model.Group
// @Router /api/groups/{id} [post]
func UpdateGroup(c *fiber.Ctx) error {
	db := database.DB
	// 原记录
	var o model.Group

	// Read the param noteId
	id := c.Params("id")

	// Find the old record
	db.Find(&o, "id = ?", id)

	// If no such record present return an error
	if o.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No present", "data": nil})
	}

	// 更新记录
	g := new(model.Group)
	if err := c.BodyParser(g); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// 仅更新自有属性
	if len(g.Name) > 0 {
		o.Name = g.Name
	}
	if len(g.Memo) > 0 {
		o.Memo = g.Memo
	}
	if len(g.Users) > 0 {
		o.Users = g.Users
	}

	db.Save(&o)
	return c.JSON(o)
}

// DeleteGroup godoc
// @Summary Delete
// @Description Delete the group and assioated field User
// @Tags groups
// @Accept  json
// @Produce  json
// @Param id path string true "ID of the group to be deleted"
// @Success 204 "No Content"
// @Router /api/groups/{id} [delete]
func DeleteGroup(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")

	// delete assioated user
	var u model.User
	err := db.Delete(&u, "group_id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete", "data": nil})
	}

	// delete group
	var g model.Group
	err = db.Delete(&g, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete", "data": nil})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// CreateUser godoc
// @Summary Create
// @Description
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body model.User true "Create User"
// @Success 200 {object} model.User
// @Router /api/users [post]
func CreateUser(c *fiber.Ctx) error {
	db := database.DB
	n := new(model.User)
	if err := c.BodyParser(n); err != nil {
		return err
	}
	n.ID = uuid.New()

	// find group
	var g model.Group
	db.Find(&g, "id = ?", n.GroupID)
	if g.ID == uuid.Nil {
		// group not found
		db.Omit("GroupID").Create(&n)
	} else {
		// group found
		db.Create(&n)
	}
	return c.JSON(n)
}

// GetUsers godoc
// @Summary Get list of all
// @Description
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} model.User
// @Router /api/users [get]
func GetUsers(c *fiber.Ctx) error {
	db := database.DB
	list := new([]model.User)
	db.Preload("Users").Find(&list)
	return c.JSON(list)
}
