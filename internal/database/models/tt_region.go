// Code generated by SQLBoiler 4.8.3 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// TTRegion is an object representing the database table.
type TTRegion struct {
	ID          string `boil:"id" json:"id" toml:"id" yaml:"id"`
	DisplayName string `boil:"display_name" json:"display_name" toml:"display_name" yaml:"display_name"`

	R *ttRegionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L ttRegionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TTRegionColumns = struct {
	ID          string
	DisplayName string
}{
	ID:          "id",
	DisplayName: "display_name",
}

var TTRegionTableColumns = struct {
	ID          string
	DisplayName string
}{
	ID:          "tt_region.id",
	DisplayName: "tt_region.display_name",
}

// Generated where

var TTRegionWhere = struct {
	ID          whereHelperstring
	DisplayName whereHelperstring
}{
	ID:          whereHelperstring{field: "\"tt_region\".\"id\""},
	DisplayName: whereHelperstring{field: "\"tt_region\".\"display_name\""},
}

// TTRegionRels is where relationship names are stored.
var TTRegionRels = struct {
}{}

// ttRegionR is where relationships are stored.
type ttRegionR struct {
}

// NewStruct creates a new relationship struct
func (*ttRegionR) NewStruct() *ttRegionR {
	return &ttRegionR{}
}

// ttRegionL is where Load methods for each relationship are stored.
type ttRegionL struct{}

var (
	ttRegionAllColumns            = []string{"id", "display_name"}
	ttRegionColumnsWithoutDefault = []string{"id", "display_name"}
	ttRegionColumnsWithDefault    = []string{}
	ttRegionPrimaryKeyColumns     = []string{"id"}
)

type (
	// TTRegionSlice is an alias for a slice of pointers to TTRegion.
	// This should almost always be used instead of []TTRegion.
	TTRegionSlice []*TTRegion
	// TTRegionHook is the signature for custom TTRegion hook methods
	TTRegionHook func(context.Context, boil.ContextExecutor, *TTRegion) error

	ttRegionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	ttRegionType                 = reflect.TypeOf(&TTRegion{})
	ttRegionMapping              = queries.MakeStructMapping(ttRegionType)
	ttRegionPrimaryKeyMapping, _ = queries.BindMapping(ttRegionType, ttRegionMapping, ttRegionPrimaryKeyColumns)
	ttRegionInsertCacheMut       sync.RWMutex
	ttRegionInsertCache          = make(map[string]insertCache)
	ttRegionUpdateCacheMut       sync.RWMutex
	ttRegionUpdateCache          = make(map[string]updateCache)
	ttRegionUpsertCacheMut       sync.RWMutex
	ttRegionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var ttRegionBeforeInsertHooks []TTRegionHook
var ttRegionBeforeUpdateHooks []TTRegionHook
var ttRegionBeforeDeleteHooks []TTRegionHook
var ttRegionBeforeUpsertHooks []TTRegionHook

var ttRegionAfterInsertHooks []TTRegionHook
var ttRegionAfterSelectHooks []TTRegionHook
var ttRegionAfterUpdateHooks []TTRegionHook
var ttRegionAfterDeleteHooks []TTRegionHook
var ttRegionAfterUpsertHooks []TTRegionHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TTRegion) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttRegionBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TTRegion) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttRegionBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TTRegion) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttRegionBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TTRegion) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttRegionBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TTRegion) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttRegionAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TTRegion) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttRegionAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TTRegion) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttRegionAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TTRegion) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttRegionAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TTRegion) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttRegionAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTTRegionHook registers your hook function for all future operations.
func AddTTRegionHook(hookPoint boil.HookPoint, ttRegionHook TTRegionHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		ttRegionBeforeInsertHooks = append(ttRegionBeforeInsertHooks, ttRegionHook)
	case boil.BeforeUpdateHook:
		ttRegionBeforeUpdateHooks = append(ttRegionBeforeUpdateHooks, ttRegionHook)
	case boil.BeforeDeleteHook:
		ttRegionBeforeDeleteHooks = append(ttRegionBeforeDeleteHooks, ttRegionHook)
	case boil.BeforeUpsertHook:
		ttRegionBeforeUpsertHooks = append(ttRegionBeforeUpsertHooks, ttRegionHook)
	case boil.AfterInsertHook:
		ttRegionAfterInsertHooks = append(ttRegionAfterInsertHooks, ttRegionHook)
	case boil.AfterSelectHook:
		ttRegionAfterSelectHooks = append(ttRegionAfterSelectHooks, ttRegionHook)
	case boil.AfterUpdateHook:
		ttRegionAfterUpdateHooks = append(ttRegionAfterUpdateHooks, ttRegionHook)
	case boil.AfterDeleteHook:
		ttRegionAfterDeleteHooks = append(ttRegionAfterDeleteHooks, ttRegionHook)
	case boil.AfterUpsertHook:
		ttRegionAfterUpsertHooks = append(ttRegionAfterUpsertHooks, ttRegionHook)
	}
}

// One returns a single ttRegion record from the query.
func (q ttRegionQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TTRegion, error) {
	o := &TTRegion{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for tt_region")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all TTRegion records from the query.
func (q ttRegionQuery) All(ctx context.Context, exec boil.ContextExecutor) (TTRegionSlice, error) {
	var o []*TTRegion

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TTRegion slice")
	}

	if len(ttRegionAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all TTRegion records in the query.
func (q ttRegionQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count tt_region rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q ttRegionQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if tt_region exists")
	}

	return count > 0, nil
}

// TTRegions retrieves all the records using an executor.
func TTRegions(mods ...qm.QueryMod) ttRegionQuery {
	mods = append(mods, qm.From("\"tt_region\""))
	return ttRegionQuery{NewQuery(mods...)}
}

// FindTTRegion retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTTRegion(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*TTRegion, error) {
	ttRegionObj := &TTRegion{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"tt_region\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, ttRegionObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from tt_region")
	}

	if err = ttRegionObj.doAfterSelectHooks(ctx, exec); err != nil {
		return ttRegionObj, err
	}

	return ttRegionObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TTRegion) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no tt_region provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(ttRegionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	ttRegionInsertCacheMut.RLock()
	cache, cached := ttRegionInsertCache[key]
	ttRegionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			ttRegionAllColumns,
			ttRegionColumnsWithDefault,
			ttRegionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(ttRegionType, ttRegionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(ttRegionType, ttRegionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"tt_region\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"tt_region\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into tt_region")
	}

	if !cached {
		ttRegionInsertCacheMut.Lock()
		ttRegionInsertCache[key] = cache
		ttRegionInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the TTRegion.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TTRegion) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	ttRegionUpdateCacheMut.RLock()
	cache, cached := ttRegionUpdateCache[key]
	ttRegionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			ttRegionAllColumns,
			ttRegionPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update tt_region, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"tt_region\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, ttRegionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(ttRegionType, ttRegionMapping, append(wl, ttRegionPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update tt_region row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for tt_region")
	}

	if !cached {
		ttRegionUpdateCacheMut.Lock()
		ttRegionUpdateCache[key] = cache
		ttRegionUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q ttRegionQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for tt_region")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for tt_region")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TTRegionSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ttRegionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"tt_region\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, ttRegionPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in ttRegion slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all ttRegion")
	}
	return rowsAff, nil
}

// Delete deletes a single TTRegion record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TTRegion) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TTRegion provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), ttRegionPrimaryKeyMapping)
	sql := "DELETE FROM \"tt_region\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from tt_region")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for tt_region")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q ttRegionQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no ttRegionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from tt_region")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for tt_region")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TTRegionSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(ttRegionBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ttRegionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"tt_region\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, ttRegionPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from ttRegion slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for tt_region")
	}

	if len(ttRegionAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *TTRegion) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTTRegion(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TTRegionSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TTRegionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ttRegionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"tt_region\".* FROM \"tt_region\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, ttRegionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TTRegionSlice")
	}

	*o = slice

	return nil
}

// TTRegionExists checks if the TTRegion row exists.
func TTRegionExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"tt_region\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if tt_region exists")
	}

	return exists, nil
}
