package datacon

import (
	"context"

	g "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //dialect for mysql
)

type gorm struct {
	ctx              context.Context
	dialect          string
	connectionString string
	db               *g.DB
}

//NewGorm : Gorm instance for Session
func NewGorm(ctx context.Context, dialect string, connectionString string) Database {
	return &gorm{ctx: ctx, dialect: dialect, connectionString: connectionString}
}

func (gr *gorm) createConnection() (*g.DB, error) {
	return g.Open(gr.dialect, gr.connectionString)
}

func (gr *gorm) GetConnection() (interface{}, error) {
	var err error
  if gr.db != nil {
		gr.db, err = gr.createConnection()
		return gr.db, err
	}
	return gr.db, nil
}
