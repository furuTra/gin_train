package user

import (
	"fmt"
    "github.com/gin-gonic/gin"
	"github.com/furuTra/gin_train/service"
)

type Controller struct{}

// Index action: GET /users
func (pc Controller) Index(c *gin.Context) {
    var s user.Service
    page := s.GeneratePaginationFromRequest(c)
    p, err := s.PaginateUser(page)
    count := s.GetUserCount()

    if err != nil {
        c.AbortWithStatus(404)
        fmt.Println(err)
    } else {
        c.JSON(200, gin.H{
            "User": p,
            "Count": count,
        })
    }
}

// Show action: GET /users/:id
func (pc Controller) Show(c *gin.Context) {
    id := c.Params.ByName("id")
    var s user.Service
    p, err := s.GetByID(id)

    if err != nil {
        c.AbortWithStatus(404)
        fmt.Println(err)
    } else {
        c.JSON(200, p)
    }
}

// Create action: POST /users
func (pc Controller) Create(c *gin.Context) {
    var s user.Service
    p, err := s.CreateUser(c)

    if err != nil {
        c.AbortWithStatus(400)
        fmt.Println(err)
    } else {
        c.JSON(201, p)
    }
}

// Update action: PUT /users
func (pc Controller) Update(c *gin.Context) {
    id := c.Params.ByName("id")
    var s user.Service
    p, err := s.UpdateById(id, c)

    if err != nil {
        c.AbortWithStatus(400)
        fmt.Println(err)
    } else {
        c.JSON(200, p)
    }
}

// Delete action: DELETE /users/:id
func (pc Controller) Delete(c *gin.Context) {
    id := c.Params.ByName("id")
    var s user.Service

    if err := s.DeleteById(id); err != nil {
        c.AbortWithStatus(403)
        fmt.Println(err)
    } else {
        c.JSON(204, gin.H{"id #" + id: "deleted"})
    }
}