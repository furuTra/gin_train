package user

import (
    "github.com/gin-gonic/gin"
	"github.com/furuTra/gin_train/db"
	"github.com/furuTra/gin_train/entity"
)

type Service struct{}

type User entity.User

// ユーザー全取得
func (s Service) GetAll() ([]User, error) {
	db := db.GetDB()
	var u []User
    if err := db.Find(&u).Error; err != nil {
        return nil, err
    }
	return u, nil
}

// ユーザー取得(Paginator)
func (s Service) PaginateUser(p entity.Pagination) ([]User, error) {
    db := db.GetDB()
    var u []User
    if err := db.Limit(p.Limit).Offset(p.Offset).Order("id " + p.Sort).Find(&u).Error; err != nil {
        return nil, err
    }
    return u, nil
}

// idでユーザー取得
func (s Service) GetByID(id string) (User, error) {
    db := db.GetDB()
    var u User

    if err := db.Where("id = ?", id).First(&u).Error; err != nil {
        return u, err
    }

    return u, nil
}

// ユーザーの作成
func (s Service) CreateUser(c *gin.Context) (User, error) {
    db := db.GetDB()
    var u User

	// user構造体をjsonとして受け取る
    if err := c.BindJSON(&u); err != nil {
        return u, err
    }

    if err := db.Create(&u).Error; err != nil {
        return u, err
    }

    return u, nil
}

// ユーザー更新
func (s Service) UpdateById(id string, c *gin.Context) (User, error) {
    db := db.GetDB()
    var u User

    // 該当するユーザーの確認
    if err := db.Where("id = ?", id).First(&u).Error; err != nil {
        return u, err
    }

    if err := c.BindJSON(&u); err != nil {
        return u, err
    }

    db.Save(&u)

    return u, nil
}

// ユーザー削除
func (s Service) DeleteById(id string) error {
    db := db.GetDB()
    var u User

    // db.Delete(&u)はデフォルトで`where id = ?`を削除する
    // しかし、今回はidのみを受け取る実装としているので、一度検索してからdeleteを呼び出している
    if err := db.Where("id = ?", id).Delete(&u).Error; err != nil {
        return err
    }

    return nil
}