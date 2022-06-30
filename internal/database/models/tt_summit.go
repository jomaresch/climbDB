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

// TTSummit is an object representing the database table.
type TTSummit struct {
	ID          string  `boil:"id" json:"id" toml:"id" yaml:"id"`
	GuideID     string  `boil:"guide_id" json:"guide_id" toml:"guide_id" yaml:"guide_id"`
	RegionID    string  `boil:"region_id" json:"region_id" toml:"region_id" yaml:"region_id"`
	DisplayName string  `boil:"display_name" json:"display_name" toml:"display_name" yaml:"display_name"`
	Lat         float64 `boil:"lat" json:"lat" toml:"lat" yaml:"lat"`
	Long        float64 `boil:"long" json:"long" toml:"long" yaml:"long"`

	R *ttSummitR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L ttSummitL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TTSummitColumns = struct {
	ID          string
	GuideID     string
	RegionID    string
	DisplayName string
	Lat         string
	Long        string
}{
	ID:          "id",
	GuideID:     "guide_id",
	RegionID:    "region_id",
	DisplayName: "display_name",
	Lat:         "lat",
	Long:        "long",
}

var TTSummitTableColumns = struct {
	ID          string
	GuideID     string
	RegionID    string
	DisplayName string
	Lat         string
	Long        string
}{
	ID:          "tt_summit.id",
	GuideID:     "tt_summit.guide_id",
	RegionID:    "tt_summit.region_id",
	DisplayName: "tt_summit.display_name",
	Lat:         "tt_summit.lat",
	Long:        "tt_summit.long",
}

// Generated where

type whereHelperfloat64 struct{ field string }

func (w whereHelperfloat64) EQ(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperfloat64) NEQ(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperfloat64) LT(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperfloat64) LTE(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperfloat64) GT(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperfloat64) GTE(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelperfloat64) IN(slice []float64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperfloat64) NIN(slice []float64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var TTSummitWhere = struct {
	ID          whereHelperstring
	GuideID     whereHelperstring
	RegionID    whereHelperstring
	DisplayName whereHelperstring
	Lat         whereHelperfloat64
	Long        whereHelperfloat64
}{
	ID:          whereHelperstring{field: "\"tt_summit\".\"id\""},
	GuideID:     whereHelperstring{field: "\"tt_summit\".\"guide_id\""},
	RegionID:    whereHelperstring{field: "\"tt_summit\".\"region_id\""},
	DisplayName: whereHelperstring{field: "\"tt_summit\".\"display_name\""},
	Lat:         whereHelperfloat64{field: "\"tt_summit\".\"lat\""},
	Long:        whereHelperfloat64{field: "\"tt_summit\".\"long\""},
}

// TTSummitRels is where relationship names are stored.
var TTSummitRels = struct {
}{}

// ttSummitR is where relationships are stored.
type ttSummitR struct {
}

// NewStruct creates a new relationship struct
func (*ttSummitR) NewStruct() *ttSummitR {
	return &ttSummitR{}
}

// ttSummitL is where Load methods for each relationship are stored.
type ttSummitL struct{}

var (
	ttSummitAllColumns            = []string{"id", "guide_id", "region_id", "display_name", "lat", "long"}
	ttSummitColumnsWithoutDefault = []string{"id", "guide_id", "region_id", "display_name", "lat", "long"}
	ttSummitColumnsWithDefault    = []string{}
	ttSummitPrimaryKeyColumns     = []string{"id"}
)

type (
	// TTSummitSlice is an alias for a slice of pointers to TTSummit.
	// This should almost always be used instead of []TTSummit.
	TTSummitSlice []*TTSummit
	// TTSummitHook is the signature for custom TTSummit hook methods
	TTSummitHook func(context.Context, boil.ContextExecutor, *TTSummit) error

	ttSummitQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	ttSummitType                 = reflect.TypeOf(&TTSummit{})
	ttSummitMapping              = queries.MakeStructMapping(ttSummitType)
	ttSummitPrimaryKeyMapping, _ = queries.BindMapping(ttSummitType, ttSummitMapping, ttSummitPrimaryKeyColumns)
	ttSummitInsertCacheMut       sync.RWMutex
	ttSummitInsertCache          = make(map[string]insertCache)
	ttSummitUpdateCacheMut       sync.RWMutex
	ttSummitUpdateCache          = make(map[string]updateCache)
	ttSummitUpsertCacheMut       sync.RWMutex
	ttSummitUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var ttSummitBeforeInsertHooks []TTSummitHook
var ttSummitBeforeUpdateHooks []TTSummitHook
var ttSummitBeforeDeleteHooks []TTSummitHook
var ttSummitBeforeUpsertHooks []TTSummitHook

var ttSummitAfterInsertHooks []TTSummitHook
var ttSummitAfterSelectHooks []TTSummitHook
var ttSummitAfterUpdateHooks []TTSummitHook
var ttSummitAfterDeleteHooks []TTSummitHook
var ttSummitAfterUpsertHooks []TTSummitHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TTSummit) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttSummitBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TTSummit) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttSummitBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TTSummit) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttSummitBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TTSummit) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttSummitBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TTSummit) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttSummitAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TTSummit) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttSummitAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TTSummit) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttSummitAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TTSummit) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttSummitAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TTSummit) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttSummitAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTTSummitHook registers your hook function for all future operations.
func AddTTSummitHook(hookPoint boil.HookPoint, ttSummitHook TTSummitHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		ttSummitBeforeInsertHooks = append(ttSummitBeforeInsertHooks, ttSummitHook)
	case boil.BeforeUpdateHook:
		ttSummitBeforeUpdateHooks = append(ttSummitBeforeUpdateHooks, ttSummitHook)
	case boil.BeforeDeleteHook:
		ttSummitBeforeDeleteHooks = append(ttSummitBeforeDeleteHooks, ttSummitHook)
	case boil.BeforeUpsertHook:
		ttSummitBeforeUpsertHooks = append(ttSummitBeforeUpsertHooks, ttSummitHook)
	case boil.AfterInsertHook:
		ttSummitAfterInsertHooks = append(ttSummitAfterInsertHooks, ttSummitHook)
	case boil.AfterSelectHook:
		ttSummitAfterSelectHooks = append(ttSummitAfterSelectHooks, ttSummitHook)
	case boil.AfterUpdateHook:
		ttSummitAfterUpdateHooks = append(ttSummitAfterUpdateHooks, ttSummitHook)
	case boil.AfterDeleteHook:
		ttSummitAfterDeleteHooks = append(ttSummitAfterDeleteHooks, ttSummitHook)
	case boil.AfterUpsertHook:
		ttSummitAfterUpsertHooks = append(ttSummitAfterUpsertHooks, ttSummitHook)
	}
}

// One returns a single ttSummit record from the query.
func (q ttSummitQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TTSummit, error) {
	o := &TTSummit{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for tt_summit")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all TTSummit records from the query.
func (q ttSummitQuery) All(ctx context.Context, exec boil.ContextExecutor) (TTSummitSlice, error) {
	var o []*TTSummit

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TTSummit slice")
	}

	if len(ttSummitAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all TTSummit records in the query.
func (q ttSummitQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count tt_summit rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q ttSummitQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if tt_summit exists")
	}

	return count > 0, nil
}

// TTSummits retrieves all the records using an executor.
func TTSummits(mods ...qm.QueryMod) ttSummitQuery {
	mods = append(mods, qm.From("\"tt_summit\""))
	return ttSummitQuery{NewQuery(mods...)}
}

// FindTTSummit retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTTSummit(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*TTSummit, error) {
	ttSummitObj := &TTSummit{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"tt_summit\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, ttSummitObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from tt_summit")
	}

	if err = ttSummitObj.doAfterSelectHooks(ctx, exec); err != nil {
		return ttSummitObj, err
	}

	return ttSummitObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TTSummit) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no tt_summit provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(ttSummitColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	ttSummitInsertCacheMut.RLock()
	cache, cached := ttSummitInsertCache[key]
	ttSummitInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			ttSummitAllColumns,
			ttSummitColumnsWithDefault,
			ttSummitColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(ttSummitType, ttSummitMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(ttSummitType, ttSummitMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"tt_summit\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"tt_summit\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into tt_summit")
	}

	if !cached {
		ttSummitInsertCacheMut.Lock()
		ttSummitInsertCache[key] = cache
		ttSummitInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the TTSummit.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TTSummit) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	ttSummitUpdateCacheMut.RLock()
	cache, cached := ttSummitUpdateCache[key]
	ttSummitUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			ttSummitAllColumns,
			ttSummitPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update tt_summit, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"tt_summit\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, ttSummitPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(ttSummitType, ttSummitMapping, append(wl, ttSummitPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update tt_summit row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for tt_summit")
	}

	if !cached {
		ttSummitUpdateCacheMut.Lock()
		ttSummitUpdateCache[key] = cache
		ttSummitUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q ttSummitQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for tt_summit")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for tt_summit")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TTSummitSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ttSummitPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"tt_summit\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, ttSummitPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in ttSummit slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all ttSummit")
	}
	return rowsAff, nil
}

// Delete deletes a single TTSummit record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TTSummit) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TTSummit provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), ttSummitPrimaryKeyMapping)
	sql := "DELETE FROM \"tt_summit\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from tt_summit")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for tt_summit")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q ttSummitQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no ttSummitQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from tt_summit")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for tt_summit")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TTSummitSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(ttSummitBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ttSummitPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"tt_summit\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, ttSummitPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from ttSummit slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for tt_summit")
	}

	if len(ttSummitAfterDeleteHooks) != 0 {
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
func (o *TTSummit) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTTSummit(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TTSummitSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TTSummitSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ttSummitPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"tt_summit\".* FROM \"tt_summit\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, ttSummitPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TTSummitSlice")
	}

	*o = slice

	return nil
}

// TTSummitExists checks if the TTSummit row exists.
func TTSummitExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"tt_summit\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if tt_summit exists")
	}

	return exists, nil
}