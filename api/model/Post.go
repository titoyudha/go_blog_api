package model

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Post struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"_id"`
	Title     string    `gorm:"size:255;not null;unique" json:"title"`
	Content   string    `gorm:"size:255;not null;" json:"content"`
	Author    User      `json:"author"`
	AuthorID  uint32    `gorm:"not null" json:"author_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Post) Prepare() {
	p.ID = 0
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Content = html.EscapeString(strings.TrimSpace(p.Content))
	p.Author = User{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Post) Validate() error {

	if p.Title == "" {
		return errors.New("Title Required")
	}

	if p.Content == "" {
		return errors.New("Content Required")
	}

	if p.AuthorID < 1 {
		return errors.New("Author Required")
	}
	return nil
}

func (p *Post) SavePost(db *gorm.DB) (*Post, error) {
	var err error
	err = db.Debug().Model(&Post{}).Create(&p).Error
	if err != nil {
		return &Post{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
		if err != nil {
			return &Post{}, err
		}
	}
	return p, nil
}

func (p *Post) FindAllPost(db *gorm.DB) (*[]Post, error) {
	var err error
	posts := []Post{}

	err = db.Debug().Model(&Post{}).Limit(10).Find(&posts).Error
	if err != nil {
		return &[]Post{}, err
	}

	if len(posts) > 0 {
		for i := range posts {
			err := db.Debug().Model(&User{}).Where("id = ?", posts[i].AuthorID).Take(&posts[i]).Error
			if err != nil {
				return &[]Post{}, err
			}
		}
	}
	return &posts, nil
}

func (p *Post) FindPostByID(db *gorm.DB, pid uint64) (*Post, error) {
	var err error
	err = db.Debug().Model(&Post{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Post{}, err
	}

	if p.ID != 0 {
		err = db.Debug().Model(&Post{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
		if err != nil {
			return &Post{}, err
		}
	}
	return p, nil
}

func (p *Post) UpdatePost(db *gorm.DB) (*Post, error) {
	var err error

	err = db.Debug().Model(&Post{}).Where("id = ?", p.ID).Update(Post{
		Title:     p.Title,
		Content:   p.Content,
		UpdatedAt: time.Now(),
	}).Error

	if err != nil {
		return &Post{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&Post{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
		if err != nil {
			return &Post{}, err
		}
	}
	return p, nil
}

func (p *Post) DeletePost(db *gorm.DB, pid uint64, uid uint32) (int64, error) {
	db = db.Debug().Model(&Post{}).Where("id = ? and author_id = ?", pid, uid).Take(&Post{}).Delete(&Post{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Post not Found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
