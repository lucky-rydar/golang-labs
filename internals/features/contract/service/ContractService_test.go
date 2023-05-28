package service

import (
	"testing"

	"github.com/it-02/dormitory/internals/db"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ContractMock struct {
	contracts []db.Contract
}

func NewContractMock() *ContractMock {
	return &ContractMock{contracts: []db.Contract{}}
}

func (cm *ContractMock) AddContract() db.Contract {
	contract := db.Contract{}
	cm.contracts = append(cm.contracts, contract)
	return contract
}

func (cm *ContractMock) GetContracts() []db.Contract {
	return cm.contracts
}

func (cm *ContractMock) GetContractById(id uint, contract *db.Contract) error {
	for _, c := range cm.contracts {
		if c.Id == id {
			*contract = c
			return nil
		}
	}
	return nil
}

func (cm *ContractMock) RemoveContractById(id uint) error {
	new_contracts := []db.Contract{}
	for _, c := range cm.contracts {
		if c.Id != id {
			new_contracts = append(new_contracts, c)
		}
	}
	cm.contracts = new_contracts
	return nil
}

type UserServiceMock struct {
}

func NewUserServiceMock() *UserServiceMock {
	return &UserServiceMock{}
}

func (usm *UserServiceMock) IsUserAdmin(uuid string) bool {
	return true
}

type ContractserviceTestSuite struct {
	suite.Suite
	contract_service ContractService
}

func (suite *ContractserviceTestSuite) SetupTest() {
	suite.contract_service = ContractService{
		contract_repository: NewContractMock(),
		user_service: NewUserServiceMock(),
	}
}

func (suite *ContractserviceTestSuite) Test_AddContract() {
	contracts := suite.contract_service.contract_repository.GetContracts()
	len1 := len(contracts)

	contract, err := suite.contract_service.AddContract("uuid")
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), contract)

	contracts = suite.contract_service.contract_repository.GetContracts()
	len2 := len(contracts)

	assert.Equal(suite.T(), len1+1, len2)
}

func (suite *ContractserviceTestSuite) Test_GetContracts() {
	contracts, _ := suite.contract_service.GetContracts("uuid")
	len1 := len(contracts)

	contracts = suite.contract_service.contract_repository.GetContracts()
	len2 := len(contracts)

	assert.Equal(suite.T(), len1, len2)
}

func TestContractservice(t *testing.T) {
	suite.Run(t, new(ContractserviceTestSuite))
}
