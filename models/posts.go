package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func AllPosts(db *gorm.DB) Post {
	var posts Post
	db.Find(&posts)
	fmt.Println(posts)
	return posts
}

func FindPost(db *gorm.DB, tag string) Post {
	var post Post
	db.Where("tag = ?", tag).First(&post)
	fmt.Println(post)
	return post
}

func (p Post) CreatePost(db *gorm.DB, post Post) {
	db.NewRecord(post)
	db.Create(&post)
}

func (p Post) UpdatePost(db *gorm.DB, post Post) {
	db.Model(&post).Updates(post)
}

func (p Post) DeletePost(db *gorm.DB) {
	db.Where("ID = ?", p.ID).Delete(&p)
}

func (p Post) UnDeletePost(db *gorm.DB) {
	db.Unscoped().Where("ID = ?", p.ID).Find(&p).Updates(map[string]interface{}{"deleted_at": nil})
}
