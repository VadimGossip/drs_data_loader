package oracle

import (
	"context"

	db "github.com/VadimGossip/platform_common/pkg/db/oracle"
	"github.com/sirupsen/logrus"

	def "github.com/VadimGossip/drs_data_loader/internal/repository"
)

var _ def.SrcGatewayRepository = (*repository)(nil)

type repository struct {
	db db.Client
}

func NewRepository(db db.Client) *repository {
	return &repository{db: db}
}

func (r *repository) GetSupGwgrIds(ctx context.Context) ([]int64, error) {
	rows, err := r.db.DB().QueryContext(ctx, sqlSUPGWQuery)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = rows.Close(); err != nil {
			logrus.WithFields(logrus.Fields{
				"handler": "GetSupGwgrIds",
				"problem": "rows close",
			}).Error(err)
		}
	}()

	var gwgrId int64
	result := make([]int64, 0)
	for rows.Next() {
		if err = rows.Scan(&gwgrId); err != nil {
			return nil, err
		}
		result = append(result, gwgrId)
	}

	return result, nil
}
