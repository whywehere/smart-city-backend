// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/whywehere/smart-city-backend/web/dal/model"
)

func newCityAirQuality(db *gorm.DB, opts ...gen.DOOption) cityAirQuality {
	_cityAirQuality := cityAirQuality{}

	_cityAirQuality.cityAirQualityDo.UseDB(db, opts...)
	_cityAirQuality.cityAirQualityDo.UseModel(&model.CityAirQuality{})

	tableName := _cityAirQuality.cityAirQualityDo.TableName()
	_cityAirQuality.ALL = field.NewAsterisk(tableName)
	_cityAirQuality.ID = field.NewInt32(tableName, "id")
	_cityAirQuality.PmNum = field.NewInt32(tableName, "pmNum")

	_cityAirQuality.fillFieldMap()

	return _cityAirQuality
}

type cityAirQuality struct {
	cityAirQualityDo cityAirQualityDo

	ALL   field.Asterisk
	ID    field.Int32
	PmNum field.Int32

	fieldMap map[string]field.Expr
}

func (c cityAirQuality) Table(newTableName string) *cityAirQuality {
	c.cityAirQualityDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cityAirQuality) As(alias string) *cityAirQuality {
	c.cityAirQualityDo.DO = *(c.cityAirQualityDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cityAirQuality) updateTableName(table string) *cityAirQuality {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt32(table, "id")
	c.PmNum = field.NewInt32(table, "pmNum")

	c.fillFieldMap()

	return c
}

func (c *cityAirQuality) WithContext(ctx context.Context) ICityAirQualityDo {
	return c.cityAirQualityDo.WithContext(ctx)
}

func (c cityAirQuality) TableName() string { return c.cityAirQualityDo.TableName() }

func (c cityAirQuality) Alias() string { return c.cityAirQualityDo.Alias() }

func (c cityAirQuality) Columns(cols ...field.Expr) gen.Columns {
	return c.cityAirQualityDo.Columns(cols...)
}

func (c *cityAirQuality) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cityAirQuality) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 2)
	c.fieldMap["id"] = c.ID
	c.fieldMap["pmNum"] = c.PmNum
}

func (c cityAirQuality) clone(db *gorm.DB) cityAirQuality {
	c.cityAirQualityDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c cityAirQuality) replaceDB(db *gorm.DB) cityAirQuality {
	c.cityAirQualityDo.ReplaceDB(db)
	return c
}

type cityAirQualityDo struct{ gen.DO }

type ICityAirQualityDo interface {
	gen.SubQuery
	Debug() ICityAirQualityDo
	WithContext(ctx context.Context) ICityAirQualityDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICityAirQualityDo
	WriteDB() ICityAirQualityDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICityAirQualityDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICityAirQualityDo
	Not(conds ...gen.Condition) ICityAirQualityDo
	Or(conds ...gen.Condition) ICityAirQualityDo
	Select(conds ...field.Expr) ICityAirQualityDo
	Where(conds ...gen.Condition) ICityAirQualityDo
	Order(conds ...field.Expr) ICityAirQualityDo
	Distinct(cols ...field.Expr) ICityAirQualityDo
	Omit(cols ...field.Expr) ICityAirQualityDo
	Join(table schema.Tabler, on ...field.Expr) ICityAirQualityDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICityAirQualityDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICityAirQualityDo
	Group(cols ...field.Expr) ICityAirQualityDo
	Having(conds ...gen.Condition) ICityAirQualityDo
	Limit(limit int) ICityAirQualityDo
	Offset(offset int) ICityAirQualityDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICityAirQualityDo
	Unscoped() ICityAirQualityDo
	Create(values ...*model.CityAirQuality) error
	CreateInBatches(values []*model.CityAirQuality, batchSize int) error
	Save(values ...*model.CityAirQuality) error
	First() (*model.CityAirQuality, error)
	Take() (*model.CityAirQuality, error)
	Last() (*model.CityAirQuality, error)
	Find() ([]*model.CityAirQuality, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CityAirQuality, err error)
	FindInBatches(result *[]*model.CityAirQuality, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.CityAirQuality) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICityAirQualityDo
	Assign(attrs ...field.AssignExpr) ICityAirQualityDo
	Joins(fields ...field.RelationField) ICityAirQualityDo
	Preload(fields ...field.RelationField) ICityAirQualityDo
	FirstOrInit() (*model.CityAirQuality, error)
	FirstOrCreate() (*model.CityAirQuality, error)
	FindByPage(offset int, limit int) (result []*model.CityAirQuality, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICityAirQualityDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c cityAirQualityDo) Debug() ICityAirQualityDo {
	return c.withDO(c.DO.Debug())
}

func (c cityAirQualityDo) WithContext(ctx context.Context) ICityAirQualityDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c cityAirQualityDo) ReadDB() ICityAirQualityDo {
	return c.Clauses(dbresolver.Read)
}

func (c cityAirQualityDo) WriteDB() ICityAirQualityDo {
	return c.Clauses(dbresolver.Write)
}

func (c cityAirQualityDo) Session(config *gorm.Session) ICityAirQualityDo {
	return c.withDO(c.DO.Session(config))
}

func (c cityAirQualityDo) Clauses(conds ...clause.Expression) ICityAirQualityDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c cityAirQualityDo) Returning(value interface{}, columns ...string) ICityAirQualityDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c cityAirQualityDo) Not(conds ...gen.Condition) ICityAirQualityDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c cityAirQualityDo) Or(conds ...gen.Condition) ICityAirQualityDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c cityAirQualityDo) Select(conds ...field.Expr) ICityAirQualityDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c cityAirQualityDo) Where(conds ...gen.Condition) ICityAirQualityDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c cityAirQualityDo) Order(conds ...field.Expr) ICityAirQualityDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c cityAirQualityDo) Distinct(cols ...field.Expr) ICityAirQualityDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c cityAirQualityDo) Omit(cols ...field.Expr) ICityAirQualityDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c cityAirQualityDo) Join(table schema.Tabler, on ...field.Expr) ICityAirQualityDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c cityAirQualityDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICityAirQualityDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c cityAirQualityDo) RightJoin(table schema.Tabler, on ...field.Expr) ICityAirQualityDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c cityAirQualityDo) Group(cols ...field.Expr) ICityAirQualityDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c cityAirQualityDo) Having(conds ...gen.Condition) ICityAirQualityDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c cityAirQualityDo) Limit(limit int) ICityAirQualityDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c cityAirQualityDo) Offset(offset int) ICityAirQualityDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c cityAirQualityDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICityAirQualityDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c cityAirQualityDo) Unscoped() ICityAirQualityDo {
	return c.withDO(c.DO.Unscoped())
}

func (c cityAirQualityDo) Create(values ...*model.CityAirQuality) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c cityAirQualityDo) CreateInBatches(values []*model.CityAirQuality, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c cityAirQualityDo) Save(values ...*model.CityAirQuality) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c cityAirQualityDo) First() (*model.CityAirQuality, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CityAirQuality), nil
	}
}

func (c cityAirQualityDo) Take() (*model.CityAirQuality, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CityAirQuality), nil
	}
}

func (c cityAirQualityDo) Last() (*model.CityAirQuality, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CityAirQuality), nil
	}
}

func (c cityAirQualityDo) Find() ([]*model.CityAirQuality, error) {
	result, err := c.DO.Find()
	return result.([]*model.CityAirQuality), err
}

func (c cityAirQualityDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CityAirQuality, err error) {
	buf := make([]*model.CityAirQuality, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c cityAirQualityDo) FindInBatches(result *[]*model.CityAirQuality, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c cityAirQualityDo) Attrs(attrs ...field.AssignExpr) ICityAirQualityDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c cityAirQualityDo) Assign(attrs ...field.AssignExpr) ICityAirQualityDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c cityAirQualityDo) Joins(fields ...field.RelationField) ICityAirQualityDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c cityAirQualityDo) Preload(fields ...field.RelationField) ICityAirQualityDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c cityAirQualityDo) FirstOrInit() (*model.CityAirQuality, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CityAirQuality), nil
	}
}

func (c cityAirQualityDo) FirstOrCreate() (*model.CityAirQuality, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CityAirQuality), nil
	}
}

func (c cityAirQualityDo) FindByPage(offset int, limit int) (result []*model.CityAirQuality, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c cityAirQualityDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c cityAirQualityDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c cityAirQualityDo) Delete(models ...*model.CityAirQuality) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *cityAirQualityDo) withDO(do gen.Dao) *cityAirQualityDo {
	c.DO = *do.(*gen.DO)
	return c
}
