package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AskAdminDbTestSuite struct {
	suite.Suite
}

func (suite *AskAdminDbTestSuite) SetupTest() {
}

func (suite *AskAdminDbTestSuite) Test_EmptyOnInit() {
	db := InitDB()
	var askAdmins []AskAdmin
	db.Find(&askAdmins)
	assert.Equal(suite.T(), 0, len(askAdmins))
}

func (suite *AskAdminDbTestSuite) Test_AddAskAdmin() {
	db := InitDB()
	askAdmin := AskAdmin{
		Action: "action",
		Name: "name",
		Surname: "surname",
	}
	db.Create(&askAdmin)
	var askAdmins []AskAdmin
	db.Find(&askAdmins)
	assert.Equal(suite.T(), 1, len(askAdmins))

	askAdmin = askAdmins[0]
	assert.Equal(suite.T(), "action", askAdmin.Action)
	assert.Equal(suite.T(), "name", askAdmin.Name)
	assert.Equal(suite.T(), "surname", askAdmin.Surname)

	db.Delete(&askAdmin)
}

func (suite *AskAdminDbTestSuite) Test_DeleteAskAdmin() {
	db := InitDB()
	askAdmin := AskAdmin{
		Action: "action",
		Name: "name",
		Surname: "surname",
	}
	db.Create(&askAdmin)
	var askAdmins []AskAdmin
	db.Find(&askAdmins)
	assert.Equal(suite.T(), 1, len(askAdmins))

	askAdmin = askAdmins[0]
	db.Delete(&askAdmin)

	askAdmins = []AskAdmin{}
	db.Find(&askAdmins)
	assert.Equal(suite.T(), 0, len(askAdmins))
}

func (suite *AskAdminDbTestSuite) Test_UpdateAskAdmin() {
	db := InitDB()
	askAdmin := AskAdmin{
		Action: "action",
		Name: "name",
		Surname: "surname",
	}
	db.Create(&askAdmin)
	var askAdmins []AskAdmin
	db.Find(&askAdmins)
	assert.Equal(suite.T(), 1, len(askAdmins))

	askAdmin = askAdmins[0]
	askAdmin.Action = "action2"
	askAdmin.Name = "name2"
	askAdmin.Surname = "surname2"
	db.Save(&askAdmin)

	askAdmins = []AskAdmin{}
	db.Find(&askAdmins)
	assert.Equal(suite.T(), 1, len(askAdmins))
	askAdmin = askAdmins[0]
	assert.Equal(suite.T(), "action2", askAdmin.Action)
	assert.Equal(suite.T(), "name2", askAdmin.Name)
	assert.Equal(suite.T(), "surname2", askAdmin.Surname)

	db.Delete(&askAdmin)
}

func TestAskAdminDbTestSuite(t *testing.T) {
	suite.Run(t, new(AskAdminDbTestSuite))
}
