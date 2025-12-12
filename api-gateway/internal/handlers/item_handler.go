package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/solluzumo/wow-ruby/gateway/internal/dto"
	"github.com/solluzumo/wow-ruby/gateway/internal/service"
)

type ItemHandler struct {
	itemService *service.ItemService
}

func NewItemHandler(itemService *service.ItemService) *ItemHandler {
	return &ItemHandler{
		itemService: itemService,
	}
}

// @Summary Получить предмет по ID
// @Description Возвращает предмет по уникальному ID. Может возвращать базовый предмет (Item), оружие (Weapon) или броню (Armor) в зависимости от типа предмета
// @Tags Предметы
// @Produce json
// @Param request body dto.ItemDetailRequest true
// @Param Authorization header string true "Access token в формате: Bearer {token}"
// @Success 200 {object} dto.ItemDetailResponse "Базовый предмет"
// @Failure 400 {object} map[string]string "Неверный формат ID"
// @Failure 404 {object} map[string]string "Предмет не найден"
// @Router /api/item/details [get]
func (h *ItemHandler) GetById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var data dto.ItemDetailRequest

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		return
	}

	item, err := h.itemService.GetItemById(ctx, data.ID)
	if err != nil {
		http.Error(w, "Не удалось найти предмет", http.StatusNotFound)
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(item)
}

// @Summary Получить список предметов с фильтрацией и пагинацией
// @Description Возвращает список предметов с поддержкой пагинации, сортировки, фильтрации и поиска. Поддерживает различные типы предметов: базовые предметы (Item), оружие (Weapon) и броню (Armor)
// @Tags Предметы
// @Accept json
// @Produce json
// @Param request body dto.ListRequest true "Параметры запроса для фильтрации, сортировки и пагинации"
// @Param Authorization header string true "Access token в формате: Bearer {token}"
// @Success 200 {object} dto.ItemListResponse "Успешный ответ со списком предметов"
// @Failure 400 {object} map[string]string "Неверный формат запроса"
// @Failure 404 {object} map[string]string "Предметы не найдены"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /api/item/list [post]
func (h *ItemHandler) GetItemList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var list_request dto.ListRequest

	if err := json.NewDecoder(r.Body).Decode(&list_request); err != nil {
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	defer r.Body.Close()

	item_list, err := h.itemService.GetItemList(ctx, &list_request)
	if err != nil {
		http.Error(w, "Не удалось найти предмет", http.StatusNotFound)
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(item_list)
}
