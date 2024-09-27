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

type TransferActionSuite struct {
	suite.Suite
	action              *TransferAction
	transactionManager  *mocktransaction.MockManager
	accountRepository   *mockrepository.MockAccountRepository
	operationRepository *mockrepository.MockOperationRepository
}

func TestTransferActionSuite(t *testing.T) {
	suite.Run(t, new(TransferActionSuite))
}

func (suite *TransferActionSuite) SetupTest() {
	t := suite.T()
	suite.transactionManager = mocktransaction.NewMockManager(t)
	suite.accountRepository = mockrepository.NewMockAccountRepository(t)
	suite.operationRepository = mockrepository.NewMockOperationRepository(t)
	suite.action = NewTransferAction(
		suite.accountRepository,
		suite.operationRepository,
		suite.transactionManager,
	)
}

func (suite *TransferActionSuite) TestShouldTransferSuccessfully() {
	userId := uuid.New()

	input := &TransferInput{
		FromAccountID: uuid.New(),
		ToAccountID:   uuid.New(),
		RequesterID:   userId,
		Amount:        100,
	}

	expectedOperation := &model.Operation{
		FromAccountID: input.FromAccountID,
		ToAccountID:   input.ToAccountID,
		Amount:        input.Amount,
	}

	expectedAccount := &model.Account{
		ID:      input.ToAccountID,
		UserID:  userId,
		Balance: 10000,
	}

	suite.accountRepository.EXPECT().
		GetById(mock.Anything, input.FromAccountID).
		Return(expectedAccount, nil)

	suite.transactionManager.EXPECT().
		RunTransaction(mock.Anything, mock.Anything).
		RunAndReturn(func(ctx context.Context, fn func(context.Context) error) error {
			return fn(ctx)
		})

	suite.operationRepository.EXPECT().
		Create(mock.Anything, mock.Anything).
		Return(expectedOperation, nil)

	suite.operationRepository.EXPECT().
		UpdateStatus(mock.Anything, mock.Anything, mock.Anything).
		Return(expectedOperation, nil)

	suite.accountRepository.EXPECT().
		AddBalanceById(mock.Anything, mock.Anything, mock.Anything).
		Return(expectedAccount, nil)

	output, err := suite.action.Execute(context.TODO(), *input)

	suite.NoError(err)
	suite.NotNil(output)

	suite.Equal(input.FromAccountID, output.Operation.FromAccountID)
	suite.Equal(input.ToAccountID, output.Operation.ToAccountID)
	suite.Equal(input.Amount, output.Operation.Amount)
}

func (suite *TransferActionSuite) TestShouldFailTransferWithInvalidInputs() {
	accountID := uuid.New()
	inputs := []TransferInput{
		{FromAccountID: accountID, ToAccountID: accountID, Amount: 100},
		{FromAccountID: uuid.New(), ToAccountID: uuid.New(), Amount: 0},
		{FromAccountID: uuid.New(), ToAccountID: uuid.New(), Amount: -100},
	}

	for _, input := range inputs {
		output, err := suite.action.Execute(context.TODO(), input)
		suite.Error(err)
		suite.Nil(output)
	}
}

func (suite *TransferActionSuite) TestShouldFailOnOperationCreationError() {
	userId := uuid.New()

	input := TransferInput{
		FromAccountID: uuid.New(),
		ToAccountID:   uuid.New(),
		RequesterID:   userId,
		Amount:        100,
	}

	expectedAccount := &model.Account{
		ID:      input.FromAccountID,
		UserID:  userId,
		Balance: 10000,
	}

	suite.accountRepository.EXPECT().
		GetById(mock.Anything, input.FromAccountID).
		Return(expectedAccount, nil)

	suite.operationRepository.EXPECT().
		Create(mock.Anything, mock.Anything).
		Return(nil, fmt.Errorf("failed to create operation"))

	output, err := suite.action.Execute(context.TODO(), input)
	suite.Error(err)
	suite.Nil(output)
}

func (suite *TransferActionSuite) TestShouldUpdateStatusToFailedOnTransactionError() {
	userId := uuid.New()

	input := TransferInput{
		FromAccountID: uuid.New(),
		ToAccountID:   uuid.New(),
		RequesterID:   userId,
		Amount:        100,
	}

	expectedAccount := &model.Account{
		ID:      input.FromAccountID,
		UserID:  userId,
		Balance: 10000,
	}

	expectedOperation := &model.Operation{
		FromAccountID: input.FromAccountID,
		ToAccountID:   input.ToAccountID,
		Amount:        input.Amount,
	}

	suite.accountRepository.EXPECT().
		GetById(mock.Anything, input.FromAccountID).
		Return(expectedAccount, nil)

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

func (suite *TransferActionSuite) TestShouldFailOnUpdateStatusError() {
	userId := uuid.New()

	input := TransferInput{
		FromAccountID: uuid.New(),
		ToAccountID:   uuid.New(),
		RequesterID:   userId,
		Amount:        100,
	}

	expectedAccount := &model.Account{
		ID:      input.FromAccountID,
		UserID:  userId,
		Balance: 10000,
	}

	expectedOperation := &model.Operation{
		FromAccountID: input.FromAccountID,
		ToAccountID:   input.ToAccountID,
		Amount:        input.Amount,
	}

	suite.accountRepository.EXPECT().
		GetById(mock.Anything, input.FromAccountID).
		Return(expectedAccount, nil)

	suite.operationRepository.EXPECT().
		Create(mock.Anything, mock.Anything).
		Return(expectedOperation, nil)

	suite.transactionManager.EXPECT().
		RunTransaction(mock.Anything, mock.Anything).
		Return(fmt.Errorf("failed to update balances"))

	suite.operationRepository.EXPECT().
		UpdateStatus(mock.Anything, mock.Anything, model.OperationStatusFailed).
		Return(nil, fmt.Errorf("failed to update status"))

	output, err := suite.action.Execute(context.TODO(), input)
	suite.Error(err)
	suite.Nil(output)
}

func (suite *TransferActionSuite) TestShouldFailOnAddBalanceOnFromAccountError() {
	input := TransferInput{
		FromAccountID: uuid.New(),
		ToAccountID:   uuid.New(),
		Amount:        100,
	}

	suite.accountRepository.EXPECT().
		AddBalanceById(mock.Anything, input.FromAccountID, input.Amount*-1).
		Return(nil, fmt.Errorf("failed to update account %s balance", input.FromAccountID))

	output, err := suite.action.updateBalances(context.TODO(), input, uuid.New())

	suite.Error(err)
	suite.Nil(output)
}

func (suite *TransferActionSuite) TestShouldFailOnAddBalanceOnToAccountError() {
	input := TransferInput{
		FromAccountID: uuid.New(),
		ToAccountID:   uuid.New(),
		Amount:        100,
	}

	expectedAccount := &model.Account{
		ID:      input.ToAccountID,
		Balance: 10000,
	}

	suite.accountRepository.EXPECT().
		AddBalanceById(mock.Anything, input.FromAccountID, input.Amount*-1).
		Return(expectedAccount, nil)

	suite.accountRepository.EXPECT().
		AddBalanceById(mock.Anything, input.ToAccountID, input.Amount).
		Return(nil, fmt.Errorf("failed to update account %s balance", input.ToAccountID))

	output, err := suite.action.updateBalances(context.TODO(), input, uuid.New())

	suite.Error(err)
	suite.Nil(output)
}

func (suite *TransferActionSuite) TestShouldFailOnUpdateStatusToSuccessError() {
	input := TransferInput{
		FromAccountID: uuid.New(),
		ToAccountID:   uuid.New(),
		Amount:        100,
	}

	expectedOperation := &model.Operation{
		FromAccountID: input.FromAccountID,
		ToAccountID:   input.ToAccountID,
		Amount:        input.Amount,
	}

	expectedAccount := &model.Account{
		ID:      input.ToAccountID,
		Balance: 10000,
	}

	suite.accountRepository.EXPECT().
		AddBalanceById(mock.Anything, mock.Anything, mock.Anything).
		Return(expectedAccount, nil).
		Times(2)

	suite.operationRepository.EXPECT().
		UpdateStatus(mock.Anything, mock.Anything, model.OperationStatusSuccess).
		Return(nil, fmt.Errorf("failed to update status"))

	output, err := suite.action.updateBalances(context.TODO(), input, expectedOperation.ID)
	suite.Error(err)
	suite.Nil(output)
}

func (suite *TransferActionSuite) TestShouldFailOnFromAccountNegativeBalanceError() {
	input := TransferInput{
		FromAccountID: uuid.New(),
		ToAccountID:   uuid.New(),
		Amount:        100,
	}

	expectedAccount := &model.Account{
		ID:      input.FromAccountID,
		Balance: -10000,
	}

	suite.accountRepository.EXPECT().
		AddBalanceById(mock.Anything, input.FromAccountID, input.Amount*-1).
		Return(expectedAccount, nil)

	output, err := suite.action.updateBalances(context.TODO(), input, uuid.New())
	suite.Error(err)
	suite.Nil(output)
}

func (suite *TransferActionSuite) TestShouldFailOnToAccountNegativeBalanceError() {
	input := TransferInput{
		FromAccountID: uuid.New(),
		ToAccountID:   uuid.New(),
		Amount:        100,
	}

	expectedFromAccount := &model.Account{
		ID:      input.FromAccountID,
		Balance: 10000,
	}

	expectedToAccount := &model.Account{
		ID:      input.FromAccountID,
		Balance: -10000,
	}

	suite.accountRepository.EXPECT().
		AddBalanceById(mock.Anything, input.FromAccountID, input.Amount*-1).
		Return(expectedFromAccount, nil)

	suite.accountRepository.EXPECT().
		AddBalanceById(mock.Anything, input.ToAccountID, input.Amount).
		Return(expectedToAccount, nil)

	output, err := suite.action.updateBalances(context.TODO(), input, uuid.New())
	suite.Error(err)
	suite.Nil(output)
}
