package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/solluzumo/wow-ruby/gateway/internal/domain"
	"github.com/solluzumo/wow-ruby/gateway/internal/dto"
	"github.com/solluzumo/wow-ruby/gateway/internal/service"
	"go.uber.org/zap"
)

type UserHandler struct {
	authService *service.AuthService
	logger      *zap.Logger
}

func NewUserHandler(authService *service.AuthService, logger *zap.Logger) *UserHandler {
	return &UserHandler{
		authService: authService,
		logger:      logger,
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
func (h *UserHandler) GetByEmail(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	var data dto.UserDetailRequest

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		return
	}

	user, err := h.authService.GetByEmail(ctx, data.Email)
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
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	ctx := r.Context()

	var req dto.UserCreateRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	data := domain.UserDomain{
		ID:       "",
		Email:    req.Email,
		Password: req.Password,
	}

	tokens, err := h.authService.Register(ctx, &data)
	if err != nil {
		http.Error(w, fmt.Sprintf("ошибка при регистрации: %v", err), http.StatusBadRequest)
		return
	}
	wholeDuration := time.Since(start)
	h.logger.Info(
		"http razniza",
		zap.Duration("duration", wholeDuration-tokens.GrpcDuration),
	)
	log.Printf("duration: %s\n", wholeDuration-tokens.GrpcDuration)
	response := dto.Response{
		Status:  201,
		Message: "Пользователь успешно добавлен",
		Content: tokens,
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
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.UserLoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	data := domain.UserDomain{
		ID:       "",
		Email:    req.Email,
		Password: req.Password,
	}

	tokens, err := h.authService.Login(ctx, &data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка аутенфикации: %v", err), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(tokens)
}
