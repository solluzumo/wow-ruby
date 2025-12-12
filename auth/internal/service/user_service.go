package service

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/solluzumo/wow-ruby/auth/internal/app"
	"github.com/solluzumo/wow-ruby/auth/internal/domain"
	"github.com/solluzumo/wow-ruby/auth/internal/dto"
	"github.com/solluzumo/wow-ruby/auth/internal/repository"
	"github.com/solluzumo/wow-ruby/pkg"
	"go.uber.org/zap"
)

type UserService struct {
	repo         repository.UserRepository
	argonParams  *pkg.Argon2Params
	userChannels *app.UserChannels
	logger       *zap.Logger
}

func NewUserService(repo repository.UserRepository, argonParams *pkg.Argon2Params, uc *app.UserChannels, logger *zap.Logger) *UserService {
	return &UserService{
		repo:         repo,
		argonParams:  argonParams,
		userChannels: uc,
		logger:       logger,
	}
}

func (us *UserService) GenerateTokens(email string) (map[string]string, error) {
	tokens, err := pkg.GenerateToken(email)
	if (tokens["acess_token"] == "") || (err != nil) {
		return nil, err
	}

	return tokens, nil
}

func (us *UserService) GetUserByEmail(ctx context.Context, email string) (*domain.UserDomain, error) {
	data, err := us.repo.GetByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	return &domain.UserDomain{
		ID:    data.ID,
		Email: data.Email,
		Hash:  data.Hash,
	}, nil
}

func (us *UserService) CreateUser(ctx context.Context, userData *domain.UserDomain) (*dto.UserLoginResponse, error) {

	resultChan := make(chan string, 1)

	hashTask := &domain.HashTaskDomain{
		Password: userData.Password,
		Result:   resultChan,
	}

	start := time.Now()

	select {
	case *us.userChannels.HashChannel <- *hashTask:
		fmt.Println("задача в обработке")
	case <-ctx.Done():
		return nil, errors.New("запрос отменён пользователем")
	default:
		return nil, errors.New("канал переполнен")
	}

	result := <-hashTask.Result

	duration := time.Since(start)

	us.logger.Info("Password hashed",
		zap.Duration("duration", duration),
	)

	close(hashTask.Result)

	userData.Hash = result
	userData.IsActive = false

	userExists, err := us.GetUserByEmail(ctx, userData.Email)
	if userExists != nil {
		return nil, fmt.Errorf("пользователь уже существует:%v", err)
	}

	_, err = us.repo.Create(ctx, userData)
	if err != nil {
		return nil, fmt.Errorf("невозможно создать запись в бд:%v", err)
	}

	tokens, err := us.GenerateTokens(userData.Email)

	if err != nil {
		return nil, fmt.Errorf("не удалось сгенерировать токены: %v", err)
	}

	return &dto.UserLoginResponse{
		AcessToken:   tokens["acess_token"],
		RefreshToken: tokens["refresh_token"],
	}, nil
}

func (us *UserService) UpdateStatus(ctx context.Context, email string) error {
	return us.repo.UpdateStatus(ctx, email)
}

func (us *UserService) LoginUser(ctx context.Context, userData *domain.UserDomain) (*dto.UserLoginResponse, error) {

	user, err := us.GetUserByEmail(ctx, userData.Email)
	if user == nil {
		return nil, &pkg.ErrorExisting{
			Text: fmt.Sprintf("ошибка при получении пользователя: %v", err),
		}
	}

	if user.IsActive {
		return nil, fmt.Errorf("пользователь %s уже авторизован", user.ID)
	}

	isPassword, err := pkg.Verify(userData.Password, user.Hash)
	if err != nil || !isPassword {
		return nil, &pkg.ErrorExisting{
			Text: fmt.Sprintf("ошибка при сравнении пароля: %v", isPassword),
		}
	}

	tokens, err := us.GenerateTokens(userData.Email)
	if err != nil {
		return nil, err
	}

	//обновляем пользователя
	if err := us.UpdateStatus(ctx, user.Email); err != nil {
		return nil, fmt.Errorf("не удалось обновить статус пользователя: %v", err)
	}

	return &dto.UserLoginResponse{
		AcessToken:   tokens["acess_token"],
		RefreshToken: tokens["refresh_token"],
	}, nil

}

func (us *UserService) ProcessHash(task domain.HashTaskDomain) error {
	result, err := pkg.HashPassword(task.Password, us.argonParams)

	task.Result <- result

	return err
}

func HashWorker(id int, hashTaskChan <-chan domain.HashTaskDomain, us *UserService, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range hashTaskChan {
		fmt.Printf("Worker %d is WOKING\n", id)
		if err := us.ProcessHash(task); err != nil {
			fmt.Printf("Worker ERROR: %v\n", err)
		}
		fmt.Printf("Worker %d FINISHED\n", id)
	}
	fmt.Printf("Worker %d STOPPED\n", id)
}
