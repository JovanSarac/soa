package service

import (
	"blogs_service/model"
	"blogs_service/repository"
	"time"
)

type BlogService struct {
	BlogRepository *repository.BlogRepository
}

func (service *BlogService) Create(blog *model.BlogPage) error {
	blog.CreationDate = time.Now()
	err := service.BlogRepository.CreateBlog(blog)
	if err != nil {
		return err
	}
	return nil
}

func (service *BlogService) GetAllBlogs() ([]model.BlogPage, error) {
	return service.BlogRepository.GetAll()
}

func (s *BlogService) FindByID(id int) (*model.BlogPage, error) {
	return s.BlogRepository.FindByID(id)
}

func (service *BlogService) UpdateOneBlog(blog *model.BlogPage) error {
	err := service.BlogRepository.UpdateOneBlog(blog)
	if err != nil {
		return err
	}
	return nil
}