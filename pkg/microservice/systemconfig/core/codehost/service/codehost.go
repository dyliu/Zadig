package service

import (
	"go.uber.org/zap"

	"github.com/koderover/zadig/pkg/microservice/systemconfig/core/codehost/repository/models"
	"github.com/koderover/zadig/pkg/microservice/systemconfig/core/codehost/repository/mongodb"
)

func CreateCodeHost(codehost *models.CodeHost, _ *zap.SugaredLogger) (*models.CodeHost, error) {
	if codehost.Type == "codehub" {
		codehost.IsReady = "2"
	}
	list, err := mongodb.NewCodehostColl().CodeHostList()
	if err != nil {
		return nil, err
	}
	codehost.ID = len(list) + 1
	return mongodb.NewCodehostColl().AddCodeHost(codehost)
}

func FindCodeHost(_ *zap.SugaredLogger) ([]*models.CodeHost, error) {
	return mongodb.NewCodehostColl().FindCodeHosts()
}

func DeleteCodeHost(id int, _ *zap.SugaredLogger) error {
	return mongodb.NewCodehostColl().DeleteCodeHostByID(id)
}

func UpdateCodeHost(host *models.CodeHost, _ *zap.SugaredLogger) (*models.CodeHost, error) {
	return mongodb.NewCodehostColl().UpdateCodeHost(host)
}

func GetCodeHost(id int, _ *zap.SugaredLogger) (*models.CodeHost, error) {
	return mongodb.NewCodehostColl().GetCodeHostByID(id)
}