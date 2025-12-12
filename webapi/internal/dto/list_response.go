package dto

import "github.com/solluzumo/wow-ruby/ruby-api/internal/models"

type ListResponse[T any] struct {
	Data       []T `json:"data"`
	Total      int `json:"total"`
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	TotalPages int `json:"total_pages"`
}

type ItemListResponse = ListResponse[models.Item]
type QuestListResponse = ListResponse[models.Quest]
type NpcListResponse = ListResponse[models.Npc]
