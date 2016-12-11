package models

import "github.com/jinzhu/gorm"

func FindPage(db *gorm.DB, slug string) Page {
	var page Page
	db.Where("slug = ?", slug).First(&page)
	return page
}

func (p Page) CreatePage(db *gorm.DB, page Page) (int, string) {
	page.Body = processMarkdown(page.Body)
	if page.Slug == "posts" {
		return 0, "Page slug cannot be 'posts'"
	}
	db.NewRecord(page)
	db.Create(&page)
	return 1, "Page created successfully"
}

func (p Page) UpdatePage(db *gorm.DB, page Page) (int, string) {
	page.Body = processMarkdown(page.Body)
	if page.Slug == "posts" {
		return 0, "Page slug cannot be 'posts'"
	}
	db.Model(&page).Updates(page)
	return 1, "Page updated successfully"
}

func (p Page) DeletePage(db *gorm.DB) {
	db.Where("ID = ?", p.ID).Delete(&p)
}

func (p Page) UnDeletePage(db *gorm.DB) {
	db.Unscoped().Where("ID = ?", p.ID).Find(&p).Updates(map[string]interface{}{"deleted_at": nil})
}
