package handlers

import (
	"encoding/json"
	"net/http"
	"wow-ruby/internal/domain"
	"wow-ruby/internal/dto"
	"wow-ruby/internal/service"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userSerivice *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userSerivice,
	}
}

// @Summary Получить пользователя по ID
// @Description Возвращает пользователя по уникальному ID
// @Tags users
// @Accept json
// @Produce json
// @Param input body dto.UserDetailRequest true "Данные пользователя"
// @Success 200 {object} domain.UserDomain
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/user/details [post]
func (h *UserHandler) GetById(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	var data dto.UserDetailRequest

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		return
	}

	user, err := h.userService.GetUserById(ctx, data.ID)
	if err != nil {
		http.Error(w, "Не удалось найти пользователя", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// @Summary Создать пользователя
// @Description Создаёт нового пользователя в системе
// @Tags users
// @Accept json
// @Produce json
// @Param input body dto.UserCreateRequest true "Данные пользователя"
// @Success 201 {object} dto.Response "Пользователь успешно создан"
// @Failure 400 {object} map[string]string "Неверный формат данных или пользователь уже существует"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /api/user/register [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.UserCreateRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	data := domain.UserDomain{
		ID:       "",
		Login:    req.Login,
		Password: req.Password,
	}

	userId, err := h.userService.CreateUser(ctx, &data)
	if err != nil {
		http.Error(w, userId, http.StatusBadRequest)
		return
	}

	response := dto.Response{
		Status:  201,
		Message: "Пользователь успешно добавлен",
		Content: map[string]interface{}{
			"userId": userId,
		},
	}
	json.NewEncoder(w).Encode(response)
}

// @Summary Авторизация пользователя
// @Description Авторизирует пользователя в системе
// @Tags users
// @Accept json
// @Produce json
// @Param input body dto.UserLoginRequest true "Данные пользователя"
// @Success 200 {object} dto.UserLoginResponse "Пользователь успешно авторизован"
// @Failure 400 {object} map[string]string "Неверный формат данных"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /api/users [post]
func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.UserLoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	data := domain.UserDomain{
		ID:       "",
		Login:    req.Login,
		Password: req.Password,
	}

	tokens, err := h.userService.LoginUser(ctx, &data)
	if err != nil {
		http.Error(w, "Пользователь не существует", http.StatusBadRequest)
		return
	}

	response := &dto.UserLoginResponse{
		AcessToken:   tokens["acess_token"],
		RefreshToken: tokens["refresh_token"],
	}

	json.NewEncoder(w).Encode(response)
}
