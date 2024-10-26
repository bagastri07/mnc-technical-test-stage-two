package user

import (
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/contract"
	"github.com/go-redsync/redsync/v4"
	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

type UseCase struct {
	db                            *gorm.DB
	asynqClient                   *asynq.Client
	redSync                       *redsync.Redsync
	userRepository                contract.UserRepository
	userBalanceRepository         contract.UserBalanceRepository
	userTransactionRepository     contract.UserTransactionRepository
	userTransferHistoryRepository contract.UserTransferHistoryRepository
}

func New() *UseCase {
	return new(UseCase)
}

func (u *UseCase) WithPostgresqlDB(db *gorm.DB) *UseCase {
	u.db = db
	return u
}

func (u *UseCase) WithAsynqClient(client *asynq.Client) *UseCase {
	u.asynqClient = client
	return u
}

func (u *UseCase) WithRedSync(rs *redsync.Redsync) *UseCase {
	u.redSync = rs
	return u
}

func (u *UseCase) WithUserRepository(repo contract.UserRepository) *UseCase {
	u.userRepository = repo
	return u
}

func (u *UseCase) WithUserBalanceRepository(repo contract.UserBalanceRepository) *UseCase {
	u.userBalanceRepository = repo
	return u
}

func (u *UseCase) WithUserTransactionRepository(repo contract.UserTransactionRepository) *UseCase {
	u.userTransactionRepository = repo
	return u
}

func (u *UseCase) WithUserTransferHistoryRepository(repo contract.UserTransferHistoryRepository) *UseCase {
	u.userTransferHistoryRepository = repo
	return u
}
