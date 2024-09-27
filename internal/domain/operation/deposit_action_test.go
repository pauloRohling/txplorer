package operation

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/pauloRohling/txplorer/internal/model"
	mockrepository "github.com/pauloRohling/txplorer/mocks/repository"
	mocktransaction "github.com/pauloRohling/txplorer/mocks/transaction"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type DepositActionSuite struct {
	suite.Suite
	action              *DepositAction
	transactionManager  *mocktransaction.MockManager
	accountRepository   *mockrepository.MockAccountRepository
	operationRepository *mockrepository.MockOperationRepository
}

func TestDepositActionSuite(t *testing.T) {
	suite.Run(t, new(DepositActionSuite))
}

func (suite *DepositActionSuite) SetupTest() {
	t := suite.T()
	suite.transactionManager = mocktransaction.NewMockManager(t)
	suite.accountRepository = mockrepository.NewMockAccountRepository(t)
	suite.operationRepository = mockrepository.NewMockOperationRepository(t)
	suite.action = NewDepositAction(
		suite.accountRepository,
		suite.operationRepository,
		suite.transactionManager,
	)
}

func (suite *DepositActionSuite) TestShouldDepositSuccessfully() {
	input := &DepositInput{
		AccountID:   uuid.New(),
		RequesterID: uuid.New(),
		Amount:      100,
	}

	expectedAccount := &model.Account{
		ID:      input.AccountID,
		Balance: 10000,
	}

	expectedOperation := &model.Operation{
		ID:            uuid.New(),
		FromAccountID: input.AccountID,
		ToAccountID:   input.AccountID,
		Amount:        input.Amount,
		Type:          model.OperationTypeDeposit.String(),
		CreatedBy:     input.RequesterID,
		Status:        model.OperationStatusPending,
	}

	suite.operationRepository.EXPECT().
		Create(mock.Anything, mock.Anything).
		RunAndReturn(func(ctx context.Context, operation *model.Operation) (*model.Operation, error) {
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

func (suite *DepositActionSuite) TestShouldFailDepositWithInvalidInputs() {
	inputs := []DepositInput{
		{AccountID: uuid.New(), RequesterID: uuid.New(), Amount: 0},
		{AccountID: uuid.New(), RequesterID: uuid.New(), Amount: -100},
	}

	for _, input := range inputs {
		output, err := suite.action.Execute(context.TODO(), input)
		suite.Error(err)
		suite.Nil(output)
	}
}

func (suite *DepositActionSuite) TestShouldFailOnOperationCreationError() {
	input := DepositInput{
		AccountID:   uuid.New(),
		RequesterID: uuid.New(),
		Amount:      100,
	}

	suite.operationRepository.EXPECT().
		Create(mock.Anything, mock.Anything).
		Return(nil, fmt.Errorf("failed to create operation"))

	output, err := suite.action.Execute(context.TODO(), input)
	suite.Error(err)
	suite.Nil(output)
}

func (suite *DepositActionSuite) TestShouldUpdateStatusToFailedOnOperationError() {
	input := DepositInput{
		AccountID:   uuid.New(),
		RequesterID: uuid.New(),
		Amount:      100,
	}

	expectedOperation := &model.Operation{
		ID:            uuid.New(),
		FromAccountID: input.AccountID,
		ToAccountID:   input.AccountID,
		Amount:        input.Amount,
		Type:          model.OperationTypeDeposit.String(),
		CreatedBy:     input.RequesterID,
		Status:        model.OperationStatusPending,
	}

	suite.operationRepository.EXPECT().
		Create(mock.Anything, mock.Anything).
		Return(expectedOperation, nil)

	suite.transactionManager.EXPECT().
		RunTransaction(mock.Anything, mock.Anything).
		Return(fmt.Errorf("failed to update balances"))

	suite.operationRepository.EXPECT().
		UpdateStatus(mock.Anything, mock.Anything, model.OperationStatusFailed).
		Return(expectedOperation, nil)

	output, err := suite.action.Execute(context.TODO(), input)
	suite.Error(err)
	suite.Nil(output)
}
