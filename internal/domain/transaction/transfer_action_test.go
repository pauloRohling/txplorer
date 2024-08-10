package transaction

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
	action                *TransferAction
	transactionManager    *mocktransaction.MockManager
	accountRepository     *mockrepository.MockAccountRepository
	transactionRepository *mockrepository.MockTransactionRepository
}

func TestTransferActionSuite(t *testing.T) {
	suite.Run(t, new(TransferActionSuite))
}

func (suite *TransferActionSuite) SetupTest() {
	t := suite.T()
	suite.transactionManager = mocktransaction.NewMockManager(t)
	suite.accountRepository = mockrepository.NewMockAccountRepository(t)
	suite.transactionRepository = mockrepository.NewMockTransactionRepository(t)
	suite.action = NewTransferAction(
		suite.transactionManager,
		suite.accountRepository,
		suite.transactionRepository,
	)
}

func (suite *TransferActionSuite) TestShouldTransferSuccessfully() {
	input := &TransferInput{
		FromAccountID: uuid.New(),
		ToAccountID:   uuid.New(),
		Amount:        100,
	}

	expectedTransaction := model.NewTransaction(input.FromAccountID, input.ToAccountID, input.Amount)
	expectedAccount := &model.Account{
		ID:      input.ToAccountID,
		Balance: 10000,
	}

	suite.transactionManager.EXPECT().
		RunTransaction(mock.Anything, mock.Anything).
		RunAndReturn(func(ctx context.Context, fn func(context.Context) error) error {
			return fn(ctx)
		})

	suite.transactionRepository.EXPECT().
		Create(mock.Anything, mock.Anything).
		Return(expectedTransaction, nil)

	suite.transactionRepository.EXPECT().
		UpdateStatus(mock.Anything, mock.Anything, mock.Anything).
		Return(expectedTransaction, nil)

	suite.accountRepository.EXPECT().
		AddBalanceById(mock.Anything, mock.Anything, mock.Anything).
		Return(expectedAccount, nil)

	output, err := suite.action.Execute(context.TODO(), *input)

	suite.NoError(err)
	suite.NotNil(output)

	suite.Equal(input.FromAccountID, output.Transaction.FromAccountID)
	suite.Equal(input.ToAccountID, output.Transaction.ToAccountID)
	suite.Equal(input.Amount, output.Transaction.Amount)
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

func (suite *TransferActionSuite) TestShouldFailOnTransactionCreationError() {
	input := TransferInput{
		FromAccountID: uuid.New(),
		ToAccountID:   uuid.New(),
		Amount:        100,
	}

	suite.transactionRepository.EXPECT().
		Create(mock.Anything, mock.Anything).
		Return(nil, fmt.Errorf("failed to create transaction"))

	output, err := suite.action.Execute(context.TODO(), input)
	suite.Error(err)
	suite.Nil(output)
}

func (suite *TransferActionSuite) TestShouldUpdateStatusToFailedOnTransactionError() {
	input := TransferInput{
		FromAccountID: uuid.New(),
		ToAccountID:   uuid.New(),
		Amount:        100,
	}

	expectedTransaction := model.NewTransaction(input.FromAccountID, input.ToAccountID, input.Amount)

	suite.transactionRepository.EXPECT().
		Create(mock.Anything, mock.Anything).
		Return(expectedTransaction, nil)

	suite.transactionManager.EXPECT().
		RunTransaction(mock.Anything, mock.Anything).
		Return(fmt.Errorf("failed to update balances"))

	suite.transactionRepository.EXPECT().
		UpdateStatus(mock.Anything, expectedTransaction.ID, model.TransactionStatusFailed).
		Return(expectedTransaction, nil)

	output, err := suite.action.Execute(context.TODO(), input)
	suite.Error(err)
	suite.Nil(output)
}

func (suite *TransferActionSuite) TestShouldFailOnUpdateStatusError() {
	input := TransferInput{
		FromAccountID: uuid.New(),
		ToAccountID:   uuid.New(),
		Amount:        100,
	}

	expectedTransaction := model.NewTransaction(input.FromAccountID, input.ToAccountID, input.Amount)

	suite.transactionRepository.EXPECT().
		Create(mock.Anything, mock.Anything).
		Return(expectedTransaction, nil)

	suite.transactionManager.EXPECT().
		RunTransaction(mock.Anything, mock.Anything).
		Return(fmt.Errorf("failed to update balances"))

	suite.transactionRepository.EXPECT().
		UpdateStatus(mock.Anything, expectedTransaction.ID, model.TransactionStatusFailed).
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

	expectedTransaction := model.NewTransaction(input.FromAccountID, input.ToAccountID, input.Amount)
	expectedAccount := &model.Account{
		ID:      input.ToAccountID,
		Balance: 10000,
	}

	suite.accountRepository.EXPECT().
		AddBalanceById(mock.Anything, mock.Anything, mock.Anything).
		Return(expectedAccount, nil).
		Times(2)

	suite.transactionRepository.EXPECT().
		UpdateStatus(mock.Anything, expectedTransaction.ID, model.TransactionStatusSuccess).
		Return(nil, fmt.Errorf("failed to update status"))

	output, err := suite.action.updateBalances(context.TODO(), input, expectedTransaction.ID)
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
