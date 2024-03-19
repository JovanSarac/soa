package repository

import (
	"blogs_service/model"

	"gorm.io/gorm"
)

type CommentRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *CommentRepository) CreateComment(comment *model.Comment) error {

	dbResult := repo.DatabaseConnection.Table(`blog."Comments"`).Create(comment)

	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *CommentRepository) UpdateComment(comment *model.Comment) error {
	dbResult := repo.DatabaseConnection.Table(`blog."Comments"`).Save(comment)

	if dbResult.Error != nil {
		return dbResult.Error
	}

	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *CommentRepository) DeleteComment(commentId int) error {

	dbResult := repo.DatabaseConnection.Table(`blog."Comments"`).Delete(&model.Comment{}, commentId)

	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}