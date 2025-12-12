package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/solluzumo/wow-ruby/gateway/internal/dto"
	"github.com/solluzumo/wow-ruby/gateway/internal/service"

	"github.com/go-chi/chi"
)

type QuestHandler struct {
	questService *service.QuestService
}

func NewQuestHandler(questService *service.QuestService) *QuestHandler {
	return &QuestHandler{
		questService: questService,
	}
}

// @Summary Получить список квестов с фильтрацией и пагинацией
// @Description Возвращает список квестов с поддержкой пагинации, сортировки, фильтрации и поиска.
// @Tags Квесты
// @Accept json
// @Produce json
// @Param request body dto.ListRequest true "Параметры запроса для фильтрации, сортировки и пагинации"
// @Param Authorization header string true "Access token в формате: Bearer {token}"
// @Success 200 {object} dto.QuestListResponse "Успешный ответ со списком квестов"
// @Failure 400 {object} map[string]string "Неверный формат запроса"
// @Failure 404 {object} map[string]string "Квест не найдены"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /api/quest/list [post]
func (q *QuestHandler) GetQuestList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var list_request dto.ListRequest

	if err := json.NewDecoder(r.Body).Decode(&list_request); err != nil {
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	quest_list, err := q.questService.GetQuestList(ctx, &list_request)
	if err != nil {
		http.Error(w, "Не удалось найти квест", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(quest_list)
}

// @Summary Получить квест по ID
// @Description Возвращает квест по уникальному ID
// @Tags Квесты
// @Produce json
// @Param id path int true "ID квест"
// @Param Authorization header string true "Access token в формате: Bearer {token}"
// @Success 200 {object} dto.QuestDetailResponse "квест"
// @Failure 400 {object} map[string]string "Неверный формат ID"
// @Failure 404 {object} map[string]string "квест не найден"
// @Router /api/quest/{id} [get]
func (q *QuestHandler) GetById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")

	quest, err := q.questService.GetQuestById(ctx, id)
	if err != nil {
		http.Error(w, "Не удалось найти предмет", http.StatusNotFound)
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(quest)
}
