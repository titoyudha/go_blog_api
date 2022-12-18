package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/titoyudha/go_blog_api/api/model"
)

var users = []model.User{
	model.User{

		Username: "test username 1",
		Email:    "test1@mail.com",
		Password: "12345678",
	},
	model.User{
		Username: "test username 2",
		Email:    "test2@mail.com",
		Password: "12345678",
	},
}

var posts = []model.Post{
	model.Post{
		Title:   "Title 1",
		Content: "Content 1",
	},
	model.Post{
		Title:   "title 2",
		Content: "Contenr 2",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&model.Post{}, model.User{}).Error
	if err != nil {
		log.Fatalf("cant drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&model.User{}, model.Post{}).Error
	if err != nil {
		log.Fatalf("cant migrate table: %v", err)
	}

	err = db.Debug().Model(&model.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&model.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = uint32(users[i].ID)

		err = db.Debug().Model(&model.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed post table: %v", err)
		}
	}
}
