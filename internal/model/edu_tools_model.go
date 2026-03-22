package model

import "time"

type EduToolsResponse struct {
	ID        int       `json:"id"`
	Publisher string    `json:"publisher"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateEduToolsRequest struct {
	Publisher string `json:"publisher" binding:"required"`
	Title     string `json:"title" binding:"required"`
	Content   string `json:"content" binding:"required"`
}

type UpdateEduToolsRequest struct {
	ID        int    `json:"id" binding:"required"`
	Publisher string `json:"publisher"`
	Title     string `json:"title"`
	Content   string `json:"content"`
}