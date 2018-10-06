// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testQuestionAnswers(t *testing.T) {
	t.Parallel()

	query := QuestionAnswers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testQuestionAnswersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &QuestionAnswer{}
	if err = randomize.Struct(seed, o, questionAnswerDBTypes, true, questionAnswerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QuestionAnswer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := QuestionAnswers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testQuestionAnswersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &QuestionAnswer{}
	if err = randomize.Struct(seed, o, questionAnswerDBTypes, true, questionAnswerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QuestionAnswer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := QuestionAnswers().DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := QuestionAnswers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testQuestionAnswersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &QuestionAnswer{}
	if err = randomize.Struct(seed, o, questionAnswerDBTypes, true, questionAnswerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QuestionAnswer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := QuestionAnswerSlice{o}

	if rowsAff, err := slice.DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := QuestionAnswers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testQuestionAnswersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &QuestionAnswer{}
	if err = randomize.Struct(seed, o, questionAnswerDBTypes, true, questionAnswerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QuestionAnswer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := QuestionAnswerExists(tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if QuestionAnswer exists: %s", err)
	}
	if !e {
		t.Errorf("Expected QuestionAnswerExists to return true, but got false.")
	}
}

func testQuestionAnswersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &QuestionAnswer{}
	if err = randomize.Struct(seed, o, questionAnswerDBTypes, true, questionAnswerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QuestionAnswer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	questionAnswerFound, err := FindQuestionAnswer(tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if questionAnswerFound == nil {
		t.Error("want a record, got nil")
	}
}

func testQuestionAnswersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &QuestionAnswer{}
	if err = randomize.Struct(seed, o, questionAnswerDBTypes, true, questionAnswerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QuestionAnswer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = QuestionAnswers().Bind(nil, tx, o); err != nil {
		t.Error(err)
	}
}

func testQuestionAnswersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &QuestionAnswer{}
	if err = randomize.Struct(seed, o, questionAnswerDBTypes, true, questionAnswerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QuestionAnswer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := QuestionAnswers().One(tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testQuestionAnswersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	questionAnswerOne := &QuestionAnswer{}
	questionAnswerTwo := &QuestionAnswer{}
	if err = randomize.Struct(seed, questionAnswerOne, questionAnswerDBTypes, false, questionAnswerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QuestionAnswer struct: %s", err)
	}
	if err = randomize.Struct(seed, questionAnswerTwo, questionAnswerDBTypes, false, questionAnswerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QuestionAnswer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = questionAnswerOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = questionAnswerTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := QuestionAnswers().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testQuestionAnswersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	questionAnswerOne := &QuestionAnswer{}
	questionAnswerTwo := &QuestionAnswer{}
	if err = randomize.Struct(seed, questionAnswerOne, questionAnswerDBTypes, false, questionAnswerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QuestionAnswer struct: %s", err)
	}
	if err = randomize.Struct(seed, questionAnswerTwo, questionAnswerDBTypes, false, questionAnswerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QuestionAnswer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = questionAnswerOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = questionAnswerTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := QuestionAnswers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func questionAnswerBeforeInsertHook(e boil.Executor, o *QuestionAnswer) error {
	*o = QuestionAnswer{}
	return nil
}

func questionAnswerAfterInsertHook(e boil.Executor, o *QuestionAnswer) error {
	*o = QuestionAnswer{}
	return nil
}

func questionAnswerAfterSelectHook(e boil.Executor, o *QuestionAnswer) error {
	*o = QuestionAnswer{}
	return nil
}

func questionAnswerBeforeUpdateHook(e boil.Executor, o *QuestionAnswer) error {
	*o = QuestionAnswer{}
	return nil
}

func questionAnswerAfterUpdateHook(e boil.Executor, o *QuestionAnswer) error {
	*o = QuestionAnswer{}
	return nil
}

func questionAnswerBeforeDeleteHook(e boil.Executor, o *QuestionAnswer) error {
	*o = QuestionAnswer{}
	return nil
}

func questionAnswerAfterDeleteHook(e boil.Executor, o *QuestionAnswer) error {
	*o = QuestionAnswer{}
	return nil
}

func questionAnswerBeforeUpsertHook(e boil.Executor, o *QuestionAnswer) error {
	*o = QuestionAnswer{}
	return nil
}

func questionAnswerAfterUpsertHook(e boil.Executor, o *QuestionAnswer) error {
	*o = QuestionAnswer{}
	return nil
}

func testQuestionAnswersHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &QuestionAnswer{}
	o := &QuestionAnswer{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, questionAnswerDBTypes, false); err != nil {
		t.Errorf("Unable to randomize QuestionAnswer object: %s", err)
	}

	AddQuestionAnswerHook(boil.BeforeInsertHook, questionAnswerBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	questionAnswerBeforeInsertHooks = []QuestionAnswerHook{}

	AddQuestionAnswerHook(boil.AfterInsertHook, questionAnswerAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	questionAnswerAfterInsertHooks = []QuestionAnswerHook{}

	AddQuestionAnswerHook(boil.AfterSelectHook, questionAnswerAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	questionAnswerAfterSelectHooks = []QuestionAnswerHook{}

	AddQuestionAnswerHook(boil.BeforeUpdateHook, questionAnswerBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	questionAnswerBeforeUpdateHooks = []QuestionAnswerHook{}

	AddQuestionAnswerHook(boil.AfterUpdateHook, questionAnswerAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	questionAnswerAfterUpdateHooks = []QuestionAnswerHook{}

	AddQuestionAnswerHook(boil.BeforeDeleteHook, questionAnswerBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	questionAnswerBeforeDeleteHooks = []QuestionAnswerHook{}

	AddQuestionAnswerHook(boil.AfterDeleteHook, questionAnswerAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	questionAnswerAfterDeleteHooks = []QuestionAnswerHook{}

	AddQuestionAnswerHook(boil.BeforeUpsertHook, questionAnswerBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	questionAnswerBeforeUpsertHooks = []QuestionAnswerHook{}

	AddQuestionAnswerHook(boil.AfterUpsertHook, questionAnswerAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	questionAnswerAfterUpsertHooks = []QuestionAnswerHook{}
}

func testQuestionAnswersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &QuestionAnswer{}
	if err = randomize.Struct(seed, o, questionAnswerDBTypes, true, questionAnswerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QuestionAnswer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := QuestionAnswers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testQuestionAnswersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &QuestionAnswer{}
	if err = randomize.Struct(seed, o, questionAnswerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize QuestionAnswer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Whitelist(questionAnswerColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := QuestionAnswers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testQuestionAnswerToOneQuestionUsingQuestion(t *testing.T) {

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var local QuestionAnswer
	var foreign Question

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, questionAnswerDBTypes, false, questionAnswerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QuestionAnswer struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, questionDBTypes, false, questionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Question struct: %s", err)
	}

	if err := foreign.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.QuestionID = foreign.ID
	if err := local.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Question().One(tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := QuestionAnswerSlice{&local}
	if err = local.L.LoadQuestion(tx, false, (*[]*QuestionAnswer)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Question == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Question = nil
	if err = local.L.LoadQuestion(tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Question == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testQuestionAnswerToOneSetOpQuestionUsingQuestion(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a QuestionAnswer
	var b, c Question

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, questionAnswerDBTypes, false, strmangle.SetComplement(questionAnswerPrimaryKeyColumns, questionAnswerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, questionDBTypes, false, strmangle.SetComplement(questionPrimaryKeyColumns, questionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, questionDBTypes, false, strmangle.SetComplement(questionPrimaryKeyColumns, questionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Question{&b, &c} {
		err = a.SetQuestion(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Question != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.QuestionAnswers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.QuestionID != x.ID {
			t.Error("foreign key was wrong value", a.QuestionID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.QuestionID))
		reflect.Indirect(reflect.ValueOf(&a.QuestionID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.QuestionID != x.ID {
			t.Error("foreign key was wrong value", a.QuestionID, x.ID)
		}
	}
}

func testQuestionAnswersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &QuestionAnswer{}
	if err = randomize.Struct(seed, o, questionAnswerDBTypes, true, questionAnswerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QuestionAnswer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testQuestionAnswersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &QuestionAnswer{}
	if err = randomize.Struct(seed, o, questionAnswerDBTypes, true, questionAnswerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QuestionAnswer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := QuestionAnswerSlice{o}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}

func testQuestionAnswersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &QuestionAnswer{}
	if err = randomize.Struct(seed, o, questionAnswerDBTypes, true, questionAnswerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QuestionAnswer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := QuestionAnswers().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	questionAnswerDBTypes = map[string]string{`Answer`: `text`, `CreatedAt`: `timestamp without time zone`, `DeletedAt`: `timestamp without time zone`, `ID`: `integer`, `QuestionID`: `integer`, `UpdatedAt`: `timestamp without time zone`}
	_                     = bytes.MinRead
)

func testQuestionAnswersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(questionAnswerPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(questionAnswerColumns) == len(questionAnswerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &QuestionAnswer{}
	if err = randomize.Struct(seed, o, questionAnswerDBTypes, true, questionAnswerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QuestionAnswer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := QuestionAnswers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, questionAnswerDBTypes, true, questionAnswerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize QuestionAnswer struct: %s", err)
	}

	if rowsAff, err := o.Update(tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testQuestionAnswersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(questionAnswerColumns) == len(questionAnswerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &QuestionAnswer{}
	if err = randomize.Struct(seed, o, questionAnswerDBTypes, true, questionAnswerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QuestionAnswer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := QuestionAnswers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, questionAnswerDBTypes, true, questionAnswerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize QuestionAnswer struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(questionAnswerColumns, questionAnswerPrimaryKeyColumns) {
		fields = questionAnswerColumns
	} else {
		fields = strmangle.SetComplement(
			questionAnswerColumns,
			questionAnswerPrimaryKeyColumns,
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

	slice := QuestionAnswerSlice{o}
	if rowsAff, err := slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testQuestionAnswersUpsert(t *testing.T) {
	t.Parallel()

	if len(questionAnswerColumns) == len(questionAnswerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := QuestionAnswer{}
	if err = randomize.Struct(seed, &o, questionAnswerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize QuestionAnswer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert QuestionAnswer: %s", err)
	}

	count, err := QuestionAnswers().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, questionAnswerDBTypes, false, questionAnswerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize QuestionAnswer struct: %s", err)
	}

	if err = o.Upsert(tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert QuestionAnswer: %s", err)
	}

	count, err = QuestionAnswers().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
