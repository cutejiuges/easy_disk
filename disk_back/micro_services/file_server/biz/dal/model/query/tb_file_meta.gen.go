// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"github.com/cutejiuges/disk_back/micro_services/file_server/biz/dal/model/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newFileMeta(db *gorm.DB, opts ...gen.DOOption) fileMeta {
	_fileMeta := fileMeta{}

	_fileMeta.fileMetaDo.UseDB(db, opts...)
	_fileMeta.fileMetaDo.UseModel(&model.FileMeta{})

	tableName := _fileMeta.fileMetaDo.TableName()
	_fileMeta.ALL = field.NewAsterisk(tableName)
	_fileMeta.ID = field.NewInt64(tableName, "id")
	_fileMeta.FileKey = field.NewString(tableName, "file_key")
	_fileMeta.FileName = field.NewString(tableName, "file_name")
	_fileMeta.FileSize = field.NewInt64(tableName, "file_size")
	_fileMeta.FileAddr = field.NewString(tableName, "file_addr")
	_fileMeta.CreateAt = field.NewTime(tableName, "create_at")
	_fileMeta.UpdateAt = field.NewTime(tableName, "update_at")
	_fileMeta.Status = field.NewInt8(tableName, "status")

	_fileMeta.fillFieldMap()

	return _fileMeta
}

// fileMeta 存储上传的文件元信息
type fileMeta struct {
	fileMetaDo fileMetaDo

	ALL      field.Asterisk
	ID       field.Int64  // 文件的唯一id
	FileKey  field.String // 文件hash
	FileName field.String // 文件名
	FileSize field.Int64  // 文件大小，单位字节
	FileAddr field.String // 文件地址
	CreateAt field.Time   // 创建时间
	UpdateAt field.Time   // 更新时间
	Status   field.Int8   // 文件状态 0-未知 1-生效中 2-已删除

	fieldMap map[string]field.Expr
}

func (f fileMeta) Table(newTableName string) *fileMeta {
	f.fileMetaDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f fileMeta) As(alias string) *fileMeta {
	f.fileMetaDo.DO = *(f.fileMetaDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *fileMeta) updateTableName(table string) *fileMeta {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewInt64(table, "id")
	f.FileKey = field.NewString(table, "file_key")
	f.FileName = field.NewString(table, "file_name")
	f.FileSize = field.NewInt64(table, "file_size")
	f.FileAddr = field.NewString(table, "file_addr")
	f.CreateAt = field.NewTime(table, "create_at")
	f.UpdateAt = field.NewTime(table, "update_at")
	f.Status = field.NewInt8(table, "status")

	f.fillFieldMap()

	return f
}

func (f *fileMeta) WithContext(ctx context.Context) IFileMetaDo { return f.fileMetaDo.WithContext(ctx) }

func (f fileMeta) TableName() string { return f.fileMetaDo.TableName() }

func (f fileMeta) Alias() string { return f.fileMetaDo.Alias() }

func (f fileMeta) Columns(cols ...field.Expr) gen.Columns { return f.fileMetaDo.Columns(cols...) }

func (f *fileMeta) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *fileMeta) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 8)
	f.fieldMap["id"] = f.ID
	f.fieldMap["file_key"] = f.FileKey
	f.fieldMap["file_name"] = f.FileName
	f.fieldMap["file_size"] = f.FileSize
	f.fieldMap["file_addr"] = f.FileAddr
	f.fieldMap["create_at"] = f.CreateAt
	f.fieldMap["update_at"] = f.UpdateAt
	f.fieldMap["status"] = f.Status
}

func (f fileMeta) clone(db *gorm.DB) fileMeta {
	f.fileMetaDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f fileMeta) replaceDB(db *gorm.DB) fileMeta {
	f.fileMetaDo.ReplaceDB(db)
	return f
}

type fileMetaDo struct{ gen.DO }

type IFileMetaDo interface {
	gen.SubQuery
	Debug() IFileMetaDo
	WithContext(ctx context.Context) IFileMetaDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFileMetaDo
	WriteDB() IFileMetaDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFileMetaDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFileMetaDo
	Not(conds ...gen.Condition) IFileMetaDo
	Or(conds ...gen.Condition) IFileMetaDo
	Select(conds ...field.Expr) IFileMetaDo
	Where(conds ...gen.Condition) IFileMetaDo
	Order(conds ...field.Expr) IFileMetaDo
	Distinct(cols ...field.Expr) IFileMetaDo
	Omit(cols ...field.Expr) IFileMetaDo
	Join(table schema.Tabler, on ...field.Expr) IFileMetaDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFileMetaDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFileMetaDo
	Group(cols ...field.Expr) IFileMetaDo
	Having(conds ...gen.Condition) IFileMetaDo
	Limit(limit int) IFileMetaDo
	Offset(offset int) IFileMetaDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFileMetaDo
	Unscoped() IFileMetaDo
	Create(values ...*model.FileMeta) error
	CreateInBatches(values []*model.FileMeta, batchSize int) error
	Save(values ...*model.FileMeta) error
	First() (*model.FileMeta, error)
	Take() (*model.FileMeta, error)
	Last() (*model.FileMeta, error)
	Find() ([]*model.FileMeta, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FileMeta, err error)
	FindInBatches(result *[]*model.FileMeta, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.FileMeta) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFileMetaDo
	Assign(attrs ...field.AssignExpr) IFileMetaDo
	Joins(fields ...field.RelationField) IFileMetaDo
	Preload(fields ...field.RelationField) IFileMetaDo
	FirstOrInit() (*model.FileMeta, error)
	FirstOrCreate() (*model.FileMeta, error)
	FindByPage(offset int, limit int) (result []*model.FileMeta, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFileMetaDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f fileMetaDo) Debug() IFileMetaDo {
	return f.withDO(f.DO.Debug())
}

func (f fileMetaDo) WithContext(ctx context.Context) IFileMetaDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f fileMetaDo) ReadDB() IFileMetaDo {
	return f.Clauses(dbresolver.Read)
}

func (f fileMetaDo) WriteDB() IFileMetaDo {
	return f.Clauses(dbresolver.Write)
}

func (f fileMetaDo) Session(config *gorm.Session) IFileMetaDo {
	return f.withDO(f.DO.Session(config))
}

func (f fileMetaDo) Clauses(conds ...clause.Expression) IFileMetaDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f fileMetaDo) Returning(value interface{}, columns ...string) IFileMetaDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f fileMetaDo) Not(conds ...gen.Condition) IFileMetaDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f fileMetaDo) Or(conds ...gen.Condition) IFileMetaDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f fileMetaDo) Select(conds ...field.Expr) IFileMetaDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f fileMetaDo) Where(conds ...gen.Condition) IFileMetaDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f fileMetaDo) Order(conds ...field.Expr) IFileMetaDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f fileMetaDo) Distinct(cols ...field.Expr) IFileMetaDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f fileMetaDo) Omit(cols ...field.Expr) IFileMetaDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f fileMetaDo) Join(table schema.Tabler, on ...field.Expr) IFileMetaDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f fileMetaDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFileMetaDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f fileMetaDo) RightJoin(table schema.Tabler, on ...field.Expr) IFileMetaDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f fileMetaDo) Group(cols ...field.Expr) IFileMetaDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f fileMetaDo) Having(conds ...gen.Condition) IFileMetaDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f fileMetaDo) Limit(limit int) IFileMetaDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f fileMetaDo) Offset(offset int) IFileMetaDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f fileMetaDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFileMetaDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f fileMetaDo) Unscoped() IFileMetaDo {
	return f.withDO(f.DO.Unscoped())
}

func (f fileMetaDo) Create(values ...*model.FileMeta) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f fileMetaDo) CreateInBatches(values []*model.FileMeta, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f fileMetaDo) Save(values ...*model.FileMeta) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f fileMetaDo) First() (*model.FileMeta, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FileMeta), nil
	}
}

func (f fileMetaDo) Take() (*model.FileMeta, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FileMeta), nil
	}
}

func (f fileMetaDo) Last() (*model.FileMeta, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FileMeta), nil
	}
}

func (f fileMetaDo) Find() ([]*model.FileMeta, error) {
	result, err := f.DO.Find()
	return result.([]*model.FileMeta), err
}

func (f fileMetaDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FileMeta, err error) {
	buf := make([]*model.FileMeta, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f fileMetaDo) FindInBatches(result *[]*model.FileMeta, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f fileMetaDo) Attrs(attrs ...field.AssignExpr) IFileMetaDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f fileMetaDo) Assign(attrs ...field.AssignExpr) IFileMetaDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f fileMetaDo) Joins(fields ...field.RelationField) IFileMetaDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f fileMetaDo) Preload(fields ...field.RelationField) IFileMetaDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f fileMetaDo) FirstOrInit() (*model.FileMeta, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FileMeta), nil
	}
}

func (f fileMetaDo) FirstOrCreate() (*model.FileMeta, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FileMeta), nil
	}
}

func (f fileMetaDo) FindByPage(offset int, limit int) (result []*model.FileMeta, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f fileMetaDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f fileMetaDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f fileMetaDo) Delete(models ...*model.FileMeta) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *fileMetaDo) withDO(do gen.Dao) *fileMetaDo {
	f.DO = *do.(*gen.DO)
	return f
}