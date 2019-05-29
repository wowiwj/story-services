package repository

import (
	pb "github.com/wowiwj/story-services/user-service/proto/user"
	"github.com/jinzhu/gorm"
)


type Repository interface {
	GetAll() ([]*pb.User, error)
	Get(id string) (*pb.User, error)
	Create(user *pb.User) error
	GetByEmailAndPassword(user *pb.User) (*pb.User, error)
}

type UserRepository struct {
	db *gorm.DB
}
