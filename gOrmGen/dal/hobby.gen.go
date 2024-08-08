// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"gorm_gen_demo/model"
)

func newHobby(db *gorm.DB, opts ...gen.DOOption) hobby {
	_hobby := hobby{}

	_hobby.hobbyDo.UseDB(db, opts...)
	_hobby.hobbyDo.UseModel(&model.Hobby{})

	tableName := _hobby.hobbyDo.TableName()
	_hobby.ALL = field.NewAsterisk(tableName)
	_hobby.ID = field.NewInt64(tableName, "id")
	_hobby.Name = field.NewString(tableName, "name")
	_hobby.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_hobby.CreatedAt = field.NewInt64(tableName, "created_at")
	_hobby.DeletedAt = field.NewInt64(tableName, "deleted_at")

	_hobby.fillFieldMap()

	return _hobby
}

type hobby struct {
	hobbyDo hobbyDo

	ALL       field.Asterisk
	ID        field.Int64
	Name      field.String
	UpdatedAt field.Int64
	CreatedAt field.Int64
	DeletedAt field.Int64

	fieldMap map[string]field.Expr
}

func (h hobby) Table(newTableName string) *hobby {
	h.hobbyDo.UseTable(newTableName)
	return h.updateTableName(newTableName)
}

func (h hobby) As(alias string) *hobby {
	h.hobbyDo.DO = *(h.hobbyDo.As(alias).(*gen.DO))
	return h.updateTableName(alias)
}

func (h *hobby) updateTableName(table string) *hobby {
	h.ALL = field.NewAsterisk(table)
	h.ID = field.NewInt64(table, "id")
	h.Name = field.NewString(table, "name")
	h.UpdatedAt = field.NewInt64(table, "updated_at")
	h.CreatedAt = field.NewInt64(table, "created_at")
	h.DeletedAt = field.NewInt64(table, "deleted_at")

	h.fillFieldMap()

	return h
}

func (h *hobby) WithContext(ctx context.Context) IHobbyDo { return h.hobbyDo.WithContext(ctx) }

func (h hobby) TableName() string { return h.hobbyDo.TableName() }

func (h hobby) Alias() string { return h.hobbyDo.Alias() }

func (h hobby) Columns(cols ...field.Expr) gen.Columns { return h.hobbyDo.Columns(cols...) }

func (h *hobby) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := h.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (h *hobby) fillFieldMap() {
	h.fieldMap = make(map[string]field.Expr, 5)
	h.fieldMap["id"] = h.ID
	h.fieldMap["name"] = h.Name
	h.fieldMap["updated_at"] = h.UpdatedAt
	h.fieldMap["created_at"] = h.CreatedAt
	h.fieldMap["deleted_at"] = h.DeletedAt
}

func (h hobby) clone(db *gorm.DB) hobby {
	h.hobbyDo.ReplaceConnPool(db.Statement.ConnPool)
	return h
}

func (h hobby) replaceDB(db *gorm.DB) hobby {
	h.hobbyDo.ReplaceDB(db)
	return h
}

type hobbyDo struct{ gen.DO }

type IHobbyDo interface {
	gen.SubQuery
	Debug() IHobbyDo
	WithContext(ctx context.Context) IHobbyDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IHobbyDo
	WriteDB() IHobbyDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IHobbyDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IHobbyDo
	Not(conds ...gen.Condition) IHobbyDo
	Or(conds ...gen.Condition) IHobbyDo
	Select(conds ...field.Expr) IHobbyDo
	Where(conds ...gen.Condition) IHobbyDo
	Order(conds ...field.Expr) IHobbyDo
	Distinct(cols ...field.Expr) IHobbyDo
	Omit(cols ...field.Expr) IHobbyDo
	Join(table schema.Tabler, on ...field.Expr) IHobbyDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IHobbyDo
	RightJoin(table schema.Tabler, on ...field.Expr) IHobbyDo
	Group(cols ...field.Expr) IHobbyDo
	Having(conds ...gen.Condition) IHobbyDo
	Limit(limit int) IHobbyDo
	Offset(offset int) IHobbyDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IHobbyDo
	Unscoped() IHobbyDo
	Create(values ...*model.Hobby) error
	CreateInBatches(values []*model.Hobby, batchSize int) error
	Save(values ...*model.Hobby) error
	First() (*model.Hobby, error)
	Take() (*model.Hobby, error)
	Last() (*model.Hobby, error)
	Find() ([]*model.Hobby, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Hobby, err error)
	FindInBatches(result *[]*model.Hobby, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Hobby) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IHobbyDo
	Assign(attrs ...field.AssignExpr) IHobbyDo
	Joins(fields ...field.RelationField) IHobbyDo
	Preload(fields ...field.RelationField) IHobbyDo
	FirstOrInit() (*model.Hobby, error)
	FirstOrCreate() (*model.Hobby, error)
	FindByPage(offset int, limit int) (result []*model.Hobby, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IHobbyDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	FilterWithId(id int) (result []model.Hobby, err error)
}

// SELECT * FROM @@table WHERE id = @id{{end}}
func (h hobbyDo) FilterWithId(id int) (result []model.Hobby, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM hobby WHERE id = ? ")

	var executeSQL *gorm.DB
	executeSQL = h.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (h hobbyDo) Debug() IHobbyDo {
	return h.withDO(h.DO.Debug())
}

func (h hobbyDo) WithContext(ctx context.Context) IHobbyDo {
	return h.withDO(h.DO.WithContext(ctx))
}

func (h hobbyDo) ReadDB() IHobbyDo {
	return h.Clauses(dbresolver.Read)
}

func (h hobbyDo) WriteDB() IHobbyDo {
	return h.Clauses(dbresolver.Write)
}

func (h hobbyDo) Session(config *gorm.Session) IHobbyDo {
	return h.withDO(h.DO.Session(config))
}

func (h hobbyDo) Clauses(conds ...clause.Expression) IHobbyDo {
	return h.withDO(h.DO.Clauses(conds...))
}

func (h hobbyDo) Returning(value interface{}, columns ...string) IHobbyDo {
	return h.withDO(h.DO.Returning(value, columns...))
}

func (h hobbyDo) Not(conds ...gen.Condition) IHobbyDo {
	return h.withDO(h.DO.Not(conds...))
}

func (h hobbyDo) Or(conds ...gen.Condition) IHobbyDo {
	return h.withDO(h.DO.Or(conds...))
}

func (h hobbyDo) Select(conds ...field.Expr) IHobbyDo {
	return h.withDO(h.DO.Select(conds...))
}

func (h hobbyDo) Where(conds ...gen.Condition) IHobbyDo {
	return h.withDO(h.DO.Where(conds...))
}

func (h hobbyDo) Order(conds ...field.Expr) IHobbyDo {
	return h.withDO(h.DO.Order(conds...))
}

func (h hobbyDo) Distinct(cols ...field.Expr) IHobbyDo {
	return h.withDO(h.DO.Distinct(cols...))
}

func (h hobbyDo) Omit(cols ...field.Expr) IHobbyDo {
	return h.withDO(h.DO.Omit(cols...))
}

func (h hobbyDo) Join(table schema.Tabler, on ...field.Expr) IHobbyDo {
	return h.withDO(h.DO.Join(table, on...))
}

func (h hobbyDo) LeftJoin(table schema.Tabler, on ...field.Expr) IHobbyDo {
	return h.withDO(h.DO.LeftJoin(table, on...))
}

func (h hobbyDo) RightJoin(table schema.Tabler, on ...field.Expr) IHobbyDo {
	return h.withDO(h.DO.RightJoin(table, on...))
}

func (h hobbyDo) Group(cols ...field.Expr) IHobbyDo {
	return h.withDO(h.DO.Group(cols...))
}

func (h hobbyDo) Having(conds ...gen.Condition) IHobbyDo {
	return h.withDO(h.DO.Having(conds...))
}

func (h hobbyDo) Limit(limit int) IHobbyDo {
	return h.withDO(h.DO.Limit(limit))
}

func (h hobbyDo) Offset(offset int) IHobbyDo {
	return h.withDO(h.DO.Offset(offset))
}

func (h hobbyDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IHobbyDo {
	return h.withDO(h.DO.Scopes(funcs...))
}

func (h hobbyDo) Unscoped() IHobbyDo {
	return h.withDO(h.DO.Unscoped())
}

func (h hobbyDo) Create(values ...*model.Hobby) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Create(values)
}

func (h hobbyDo) CreateInBatches(values []*model.Hobby, batchSize int) error {
	return h.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (h hobbyDo) Save(values ...*model.Hobby) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Save(values)
}

func (h hobbyDo) First() (*model.Hobby, error) {
	if result, err := h.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Hobby), nil
	}
}

func (h hobbyDo) Take() (*model.Hobby, error) {
	if result, err := h.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Hobby), nil
	}
}

func (h hobbyDo) Last() (*model.Hobby, error) {
	if result, err := h.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Hobby), nil
	}
}

func (h hobbyDo) Find() ([]*model.Hobby, error) {
	result, err := h.DO.Find()
	return result.([]*model.Hobby), err
}

func (h hobbyDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Hobby, err error) {
	buf := make([]*model.Hobby, 0, batchSize)
	err = h.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (h hobbyDo) FindInBatches(result *[]*model.Hobby, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return h.DO.FindInBatches(result, batchSize, fc)
}

func (h hobbyDo) Attrs(attrs ...field.AssignExpr) IHobbyDo {
	return h.withDO(h.DO.Attrs(attrs...))
}

func (h hobbyDo) Assign(attrs ...field.AssignExpr) IHobbyDo {
	return h.withDO(h.DO.Assign(attrs...))
}

func (h hobbyDo) Joins(fields ...field.RelationField) IHobbyDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Joins(_f))
	}
	return &h
}

func (h hobbyDo) Preload(fields ...field.RelationField) IHobbyDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Preload(_f))
	}
	return &h
}

func (h hobbyDo) FirstOrInit() (*model.Hobby, error) {
	if result, err := h.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Hobby), nil
	}
}

func (h hobbyDo) FirstOrCreate() (*model.Hobby, error) {
	if result, err := h.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Hobby), nil
	}
}

func (h hobbyDo) FindByPage(offset int, limit int) (result []*model.Hobby, count int64, err error) {
	result, err = h.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = h.Offset(-1).Limit(-1).Count()
	return
}

func (h hobbyDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = h.Count()
	if err != nil {
		return
	}

	err = h.Offset(offset).Limit(limit).Scan(result)
	return
}

func (h hobbyDo) Scan(result interface{}) (err error) {
	return h.DO.Scan(result)
}

func (h hobbyDo) Delete(models ...*model.Hobby) (result gen.ResultInfo, err error) {
	return h.DO.Delete(models)
}

func (h *hobbyDo) withDO(do gen.Dao) *hobbyDo {
	h.DO = *do.(*gen.DO)
	return h
}
