package operation

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/pauloRohling/txplorer/internal/domain/account"
	"github.com/pauloRohling/txplorer/internal/domain/operation"
	mockaccount "github.com/pauloRohling/txplorer/mocks/account"
	mockoperation "github.com/pauloRohling/txplorer/mocks/operation"
	mocktransaction "github.com/pauloRohling/txplorer/mocks/transaction"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type WithdrawActionSuite struct {
	suite.Suite
	action              *WithdrawAction
	transactionManager  *mocktransaction.MockManager
	accountRepository   *mockaccount.MockRepository
	operationRepository *mockoperation.MockRepository
}

func TestWithdrawActionSuite(t *testing.T) {
	suite.Run(t, new(WithdrawActionSuite))
}

func (suite *WithdrawActionSuite) SetupTest() {
	t := suite.T()
	suite.transactionManager = mocktransaction.NewMockManager(t)
	suite.accountRepository = mockaccount.NewMockRepository(t)
	suite.operationRepository = mockoperation.NewMockRepository(t)
	suite.action = NewWithdrawAction(
		suite.accountRepository,
		suite.operationRepository,
		suite.transactionManager,
	)
}

func (suite *WithdrawActionSuite) TestShouldWithdrawSuccessfully() {
	userId := uuid.New()

	input := &WithdrawInput{
		AccountID:   uuid.New(),
		RequesterID: userId,
		Amount:      100,
	}

	expectedAccount := &account.Account{
		ID:      input.AccountID,
		UserID:  userId,
		Balance: 10000,
	}

	expectedOperation := &operation.Operation{
		ID:            uuid.New(),
		FromAccountID: input.AccountID,
		ToAccountID:   input.AccountID,
		Amount:        input.Amount,
		Type:          operation.WithdrawType.String(),
		CreatedBy:     input.RequesterID,
		Status:        operation.PendingStatus,
	}

	suite.accountRepository.EXPECT().
		GetById(mock.Anything, input.AccountID).
		Return(expectedAccount, nil)

	suite.operationRepository.EXPECT().
		Create(mock.Anything, mock.Anything).
		RunAndReturn(func(ctx context.Context, operation *operation.Operation) (*operation.Operation, error) {
			return operation, nil
		})

	suite.transactionManager.EXPECT().
		RunTransaction(mock.Anything, mock.Anything).
		RunAndReturn(func(ctx context.Context, fn func(context.Context) error) error {
			return fn(ctx)
		})

	suite.accountRepository.EXPECT().
		AddBalanceById(mock.Anything, mock.Anything, mock.Anything).
		Return(expectedAccount, nil)

	suite.operationRepository.EXPECT().
		UpdateStatus(mock.Anything, mock.Anything, mock.Anything).
		Return(expectedOperation, nil)

	output, err := suite.action.Execute(context.TODO(), *input)

	suite.NoError(err)
	suite.NotNil(output)

	suite.Equal(input.AccountID, output.Operation.FromAccountID)
	suite.Equal(input.AccountID, output.Operation.ToAccountID)
	suite.Equal(input.Amount, output.Operation.Amount)
}

func (suite *WithdrawActionSuite) TestShouldFailWithdrawWithInvalidInputs() {
	inputs := []WithdrawInput{
		{AccountID: uuid.New(), RequesterID: uuid.New(), Amount: 0},
		{AccountID: uuid.New(), RequesterID: uuid.New(), Amount: -100},
	}

	for _, input := range inputs {
		output, err := suite.action.Execute(context.TODO(), input)
		suite.Error(err)
		suite.Nil(output)
	}
}

func (suite *WithdrawActionSuite) TestShouldFailOnOperationCreationError() {
	userId := uuid.New()

	input := WithdrawInput{
		AccountID:   uuid.New(),
		RequesterID: userId,
		Amount:      100,
	}

	suite.accountRepository.EXPECT().
		GetById(mock.Anything, input.AccountID).
		Return(nil, fmt.Errorf("failed to get account"))

	output, err := suite.action.Execute(context.TODO(), input)
	suite.Error(err)
	suite.Nil(output)
}
