package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wow-ruby/internal/dto"
	"wow-ruby/internal/service"

	"github.com/go-chi/chi"
)

type NpcHandler struct {
	npcService *service.NpcService
}

func NewNpcHandler(npcService *service.NpcService) *NpcHandler {
	return &NpcHandler{
		npcService: npcService,
	}
}

// @Summary Получить список НИП с фильтрацией и пагинацией
// @Description Возвращает список НИП с поддержкой пагинации, сортировки, фильтрации и поиска.
// @Tags НИП
// @Accept json
// @Produce json
// @Param request body dto.ListRequest true "Параметры запроса для фильтрации, сортировки и пагинации"
// @Param Authorization header string true "Access token в формате: Bearer {token}"
// @Success 200 {object} dto.NpcListResponse "Успешный ответ со списком НИП"
// @Failure 400 {object} map[string]string "Неверный формат запроса"
// @Failure 404 {object} map[string]string "НИП не найдены"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /api/npc/list [post]
func (n *NpcHandler) GetNpcList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var list_request dto.ListRequest

	if err := json.NewDecoder(r.Body).Decode(&list_request); err != nil {
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	npc_list, err := n.npcService.GetNpcList(ctx, &list_request)
	if err != nil {
		http.Error(w, "Не удалось найти НИП", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(npc_list)
}

// @Summary Получить НИП по ID
// @Description Возвращает НИП по уникальному ID
// @Tags НИП
// @Produce json
// @Param id path int true "ID нип"
// @Param Authorization header string true "Access token в формате: Bearer {token}"
// @Success 200 {object} dto.NpcDetailResponse "НИП"
// @Failure 400 {object} map[string]string "Неверный формат ID"
// @Failure 404 {object} map[string]string "НИП не найден"
// @Router /api/npc/{id} [get]
func (n *NpcHandler) GetById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")

	npc, err := n.npcService.GetNpcById(ctx, id)
	if err != nil {
		http.Error(w, "Не удалось найти НИП", http.StatusNotFound)
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(npc)
}
