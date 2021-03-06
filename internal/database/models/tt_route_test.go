// Code generated by SQLBoiler 4.8.3 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testTTRoutes(t *testing.T) {
	t.Parallel()

	query := TTRoutes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTTRoutesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TTRoute{}
	if err = randomize.Struct(seed, o, ttRouteDBTypes, true, ttRouteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TTRoute struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TTRoutes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTTRoutesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TTRoute{}
	if err = randomize.Struct(seed, o, ttRouteDBTypes, true, ttRouteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TTRoute struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := TTRoutes().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TTRoutes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTTRoutesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TTRoute{}
	if err = randomize.Struct(seed, o, ttRouteDBTypes, true, ttRouteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TTRoute struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TTRouteSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TTRoutes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTTRoutesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TTRoute{}
	if err = randomize.Struct(seed, o, ttRouteDBTypes, true, ttRouteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TTRoute struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := TTRouteExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if TTRoute exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TTRouteExists to return true, but got false.")
	}
}

func testTTRoutesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TTRoute{}
	if err = randomize.Struct(seed, o, ttRouteDBTypes, true, ttRouteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TTRoute struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	ttRouteFound, err := FindTTRoute(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if ttRouteFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTTRoutesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TTRoute{}
	if err = randomize.Struct(seed, o, ttRouteDBTypes, true, ttRouteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TTRoute struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = TTRoutes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testTTRoutesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TTRoute{}
	if err = randomize.Struct(seed, o, ttRouteDBTypes, true, ttRouteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TTRoute struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := TTRoutes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTTRoutesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ttRouteOne := &TTRoute{}
	ttRouteTwo := &TTRoute{}
	if err = randomize.Struct(seed, ttRouteOne, ttRouteDBTypes, false, ttRouteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TTRoute struct: %s", err)
	}
	if err = randomize.Struct(seed, ttRouteTwo, ttRouteDBTypes, false, ttRouteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TTRoute struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = ttRouteOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = ttRouteTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := TTRoutes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTTRoutesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	ttRouteOne := &TTRoute{}
	ttRouteTwo := &TTRoute{}
	if err = randomize.Struct(seed, ttRouteOne, ttRouteDBTypes, false, ttRouteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TTRoute struct: %s", err)
	}
	if err = randomize.Struct(seed, ttRouteTwo, ttRouteDBTypes, false, ttRouteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TTRoute struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = ttRouteOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = ttRouteTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TTRoutes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func ttRouteBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *TTRoute) error {
	*o = TTRoute{}
	return nil
}

func ttRouteAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *TTRoute) error {
	*o = TTRoute{}
	return nil
}

func ttRouteAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *TTRoute) error {
	*o = TTRoute{}
	return nil
}

func ttRouteBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *TTRoute) error {
	*o = TTRoute{}
	return nil
}

func ttRouteAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *TTRoute) error {
	*o = TTRoute{}
	return nil
}

func ttRouteBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *TTRoute) error {
	*o = TTRoute{}
	return nil
}

func ttRouteAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *TTRoute) error {
	*o = TTRoute{}
	return nil
}

func ttRouteBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *TTRoute) error {
	*o = TTRoute{}
	return nil
}

func ttRouteAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *TTRoute) error {
	*o = TTRoute{}
	return nil
}

func testTTRoutesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &TTRoute{}
	o := &TTRoute{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, ttRouteDBTypes, false); err != nil {
		t.Errorf("Unable to randomize TTRoute object: %s", err)
	}

	AddTTRouteHook(boil.BeforeInsertHook, ttRouteBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	ttRouteBeforeInsertHooks = []TTRouteHook{}

	AddTTRouteHook(boil.AfterInsertHook, ttRouteAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	ttRouteAfterInsertHooks = []TTRouteHook{}

	AddTTRouteHook(boil.AfterSelectHook, ttRouteAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	ttRouteAfterSelectHooks = []TTRouteHook{}

	AddTTRouteHook(boil.BeforeUpdateHook, ttRouteBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	ttRouteBeforeUpdateHooks = []TTRouteHook{}

	AddTTRouteHook(boil.AfterUpdateHook, ttRouteAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	ttRouteAfterUpdateHooks = []TTRouteHook{}

	AddTTRouteHook(boil.BeforeDeleteHook, ttRouteBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	ttRouteBeforeDeleteHooks = []TTRouteHook{}

	AddTTRouteHook(boil.AfterDeleteHook, ttRouteAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	ttRouteAfterDeleteHooks = []TTRouteHook{}

	AddTTRouteHook(boil.BeforeUpsertHook, ttRouteBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	ttRouteBeforeUpsertHooks = []TTRouteHook{}

	AddTTRouteHook(boil.AfterUpsertHook, ttRouteAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	ttRouteAfterUpsertHooks = []TTRouteHook{}
}

func testTTRoutesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TTRoute{}
	if err = randomize.Struct(seed, o, ttRouteDBTypes, true, ttRouteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TTRoute struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TTRoutes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTTRoutesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TTRoute{}
	if err = randomize.Struct(seed, o, ttRouteDBTypes, true); err != nil {
		t.Errorf("Unable to randomize TTRoute struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(ttRouteColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := TTRoutes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTTRoutesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TTRoute{}
	if err = randomize.Struct(seed, o, ttRouteDBTypes, true, ttRouteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TTRoute struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTTRoutesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TTRoute{}
	if err = randomize.Struct(seed, o, ttRouteDBTypes, true, ttRouteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TTRoute struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TTRouteSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTTRoutesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TTRoute{}
	if err = randomize.Struct(seed, o, ttRouteDBTypes, true, ttRouteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TTRoute struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := TTRoutes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	ttRouteDBTypes = map[string]string{`ID`: `TEXT`, `SummitID`: `TEXT`, `DisplayName`: `TEXT`, `AvgRating`: `INTEGER`, `GuideRating`: `INTEGER`, `GradeID`: `INTEGER`, `SuggestedGradeID`: `INTEGER`, `RedPointGradeID`: `INTEGER`, `JumpGrade`: `INTEGER`}
	_              = bytes.MinRead
)

func testTTRoutesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(ttRoutePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(ttRouteAllColumns) == len(ttRoutePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &TTRoute{}
	if err = randomize.Struct(seed, o, ttRouteDBTypes, true, ttRouteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TTRoute struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TTRoutes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, ttRouteDBTypes, true, ttRoutePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TTRoute struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testTTRoutesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(ttRouteAllColumns) == len(ttRoutePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &TTRoute{}
	if err = randomize.Struct(seed, o, ttRouteDBTypes, true, ttRouteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TTRoute struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TTRoutes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, ttRouteDBTypes, true, ttRoutePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TTRoute struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(ttRouteAllColumns, ttRoutePrimaryKeyColumns) {
		fields = ttRouteAllColumns
	} else {
		fields = strmangle.SetComplement(
			ttRouteAllColumns,
			ttRoutePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := TTRouteSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}
