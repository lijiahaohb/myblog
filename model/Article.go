package model

import (
	"ginblog/utils/errmsg"

	"github.com/jinzhu/gorm"
)

type Article struct {
	// 指定外键
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

// 添加文章
func CreateArticle(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询单个文章
func GetArticleInfo(id int) (Article, int) {
	var article  Article
	err := db.Preload("Category").Where("id = ?", id).First(&article).Error
	if err != nil {
		return article, errmsg.ERROR_ARTICLE_NOT_EXIST
	}
	return article, errmsg.SUCCESS
}

// 查询文章列表
func GetArticles(pageSize int, pageNum int) ([]Article, int, int) {
	var articles []Article
	var total int
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articles).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return articles, errmsg.SUCCESS, total
}

// 查询分类下所有文章
func GetCateArticles(id int, pageSize int, pageNum int) ([]Article, int, int) {
	var cateArticles []Article
	var total int
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Where("cid = ?", id).Find(&cateArticles).Count(&total).Error
	if err != nil {
		return nil, errmsg.ERROR_CATEGORY_NOT_EXIST, 0
	}
	return cateArticles, errmsg.SUCCESS, total
}

// 编辑文章
func EditArticle(id int, data *Article) int {
	var article Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	err := db.Model(&article).Where("id = ?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除文章
func DeleteArticle(id int) int {
	var article Article
	err := db.Where("id = ?", id).Delete(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
