package usecase

import (
	"context"
	"fmt"
	"hacktiv/model"
	"time"

	"hacktiv/repository"

	"github.com/sirupsen/logrus"
	"go.elastic.co/apm/module/apmlogrus"
	"go.elastic.co/apm/v2"
	"go.elastic.co/ecslogrus"
)

type userUsecase struct {
	userRepo repository.IUserRepository
}

type IUserUsecase interface {
	CreateUser(ctx context.Context, user model.User) error
	GetAllUsers(ctx context.Context) ([]model.User, error)
}

func NewUserUsecase(userRepo repository.IUserRepository) IUserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) CreateUser(ctx context.Context, user model.User) error {
	// harusnya ada logic validaito disini
	// ...
	err := u.userRepo.CreateUser(ctx, user)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (u *userUsecase) GetAllUsers(ctx context.Context) ([]model.User, error) {
	span, _ := apm.StartSpan(ctx, "usecase.GetAllUsers", "userUsecase")
	defer span.End()

	time.Sleep(100 * time.Millisecond)

	log := logrus.New()
	log.SetFormatter(&ecslogrus.Formatter{})
	log.ReportCaller = true

	log.AddHook(&apmlogrus.Hook{})

	log.Info("Ini log")

	log.WithContext(apm.ContextWithSpan(ctx, span)).Info("Mencoba menggunakan with context apmlogrus")

	tx := apm.TransactionFromContext(ctx)

	traceID := ""
	if tx != nil {
		traceID = tx.TraceContext().Trace.String()
	}

	log.WithFields(logrus.Fields{
		"trace.id": traceID,
		"span.id":  span.TraceContext().Span.String(),
	}).Info("Log message with trace info")

	// berisi logic business (validation, etc)
	// ....
	var users []model.User
	users, err := u.userRepo.GetAllUsers(ctx)
	if err != nil {
		fmt.Println(err)
		return users, err
	}

	span2, _ := apm.StartSpan(ctx, "usecase.DoSomething", "userUsecase")
	// disini ceritanya do something
	time.Sleep(10 * time.Second)
	span2.End()

	return users, nil
}
