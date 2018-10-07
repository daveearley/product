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

func testAccounts(t *testing.T) {
	t.Parallel()

	query := Accounts()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testAccountsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
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

	count, err := Accounts().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAccountsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Accounts().DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Accounts().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAccountsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AccountSlice{o}

	if rowsAff, err := slice.DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Accounts().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAccountsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := AccountExists(tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Account exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AccountExists to return true, but got false.")
	}
}

func testAccountsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	accountFound, err := FindAccount(tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if accountFound == nil {
		t.Error("want a record, got nil")
	}
}

func testAccountsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Accounts().Bind(nil, tx, o); err != nil {
		t.Error(err)
	}
}

func testAccountsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Accounts().One(tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAccountsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	accountOne := &Account{}
	accountTwo := &Account{}
	if err = randomize.Struct(seed, accountOne, accountDBTypes, false, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}
	if err = randomize.Struct(seed, accountTwo, accountDBTypes, false, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = accountOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = accountTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Accounts().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAccountsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	accountOne := &Account{}
	accountTwo := &Account{}
	if err = randomize.Struct(seed, accountOne, accountDBTypes, false, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}
	if err = randomize.Struct(seed, accountTwo, accountDBTypes, false, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = accountOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = accountTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Accounts().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func accountBeforeInsertHook(e boil.Executor, o *Account) error {
	*o = Account{}
	return nil
}

func accountAfterInsertHook(e boil.Executor, o *Account) error {
	*o = Account{}
	return nil
}

func accountAfterSelectHook(e boil.Executor, o *Account) error {
	*o = Account{}
	return nil
}

func accountBeforeUpdateHook(e boil.Executor, o *Account) error {
	*o = Account{}
	return nil
}

func accountAfterUpdateHook(e boil.Executor, o *Account) error {
	*o = Account{}
	return nil
}

func accountBeforeDeleteHook(e boil.Executor, o *Account) error {
	*o = Account{}
	return nil
}

func accountAfterDeleteHook(e boil.Executor, o *Account) error {
	*o = Account{}
	return nil
}

func accountBeforeUpsertHook(e boil.Executor, o *Account) error {
	*o = Account{}
	return nil
}

func accountAfterUpsertHook(e boil.Executor, o *Account) error {
	*o = Account{}
	return nil
}

func testAccountsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Account{}
	o := &Account{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, accountDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Account object: %s", err)
	}

	AddAccountHook(boil.BeforeInsertHook, accountBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	accountBeforeInsertHooks = []AccountHook{}

	AddAccountHook(boil.AfterInsertHook, accountAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	accountAfterInsertHooks = []AccountHook{}

	AddAccountHook(boil.AfterSelectHook, accountAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	accountAfterSelectHooks = []AccountHook{}

	AddAccountHook(boil.BeforeUpdateHook, accountBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	accountBeforeUpdateHooks = []AccountHook{}

	AddAccountHook(boil.AfterUpdateHook, accountAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	accountAfterUpdateHooks = []AccountHook{}

	AddAccountHook(boil.BeforeDeleteHook, accountBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	accountBeforeDeleteHooks = []AccountHook{}

	AddAccountHook(boil.AfterDeleteHook, accountAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	accountAfterDeleteHooks = []AccountHook{}

	AddAccountHook(boil.BeforeUpsertHook, accountBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	accountBeforeUpsertHooks = []AccountHook{}

	AddAccountHook(boil.AfterUpsertHook, accountAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	accountAfterUpsertHooks = []AccountHook{}
}

func testAccountsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Accounts().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAccountsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Whitelist(accountColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Accounts().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAccountToManyAccountUsers(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Account
	var b, c AccountUser

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, accountUserDBTypes, false, accountUserColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, accountUserDBTypes, false, accountUserColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.AccountID = a.ID
	c.AccountID = a.ID

	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	accountUser, err := a.AccountUsers().All(tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range accountUser {
		if v.AccountID == b.AccountID {
			bFound = true
		}
		if v.AccountID == c.AccountID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := AccountSlice{&a}
	if err = a.L.LoadAccountUsers(tx, false, (*[]*Account)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AccountUsers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.AccountUsers = nil
	if err = a.L.LoadAccountUsers(tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AccountUsers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", accountUser)
	}
}

func testAccountToManyEvents(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Account
	var b, c Event

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, eventDBTypes, false, eventColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, eventDBTypes, false, eventColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.AccountID = a.ID
	c.AccountID = a.ID

	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	event, err := a.Events().All(tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range event {
		if v.AccountID == b.AccountID {
			bFound = true
		}
		if v.AccountID == c.AccountID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := AccountSlice{&a}
	if err = a.L.LoadEvents(tx, false, (*[]*Account)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Events); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Events = nil
	if err = a.L.LoadEvents(tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Events); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", event)
	}
}

func testAccountToManyUsers(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Account
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.AccountID = a.ID
	c.AccountID = a.ID

	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	user, err := a.Users().All(tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range user {
		if v.AccountID == b.AccountID {
			bFound = true
		}
		if v.AccountID == c.AccountID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := AccountSlice{&a}
	if err = a.L.LoadUsers(tx, false, (*[]*Account)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Users); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Users = nil
	if err = a.L.LoadUsers(tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Users); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", user)
	}
}

func testAccountToManyAddOpAccountUsers(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Account
	var b, c, d, e AccountUser

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, accountDBTypes, false, strmangle.SetComplement(accountPrimaryKeyColumns, accountColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*AccountUser{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, accountUserDBTypes, false, strmangle.SetComplement(accountUserPrimaryKeyColumns, accountUserColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*AccountUser{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddAccountUsers(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.AccountID {
			t.Error("foreign key was wrong value", a.ID, first.AccountID)
		}
		if a.ID != second.AccountID {
			t.Error("foreign key was wrong value", a.ID, second.AccountID)
		}

		if first.R.Account != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Account != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.AccountUsers[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.AccountUsers[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.AccountUsers().Count(tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testAccountToManyAddOpEvents(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Account
	var b, c, d, e Event

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, accountDBTypes, false, strmangle.SetComplement(accountPrimaryKeyColumns, accountColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Event{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, eventDBTypes, false, strmangle.SetComplement(eventPrimaryKeyColumns, eventColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Event{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddEvents(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.AccountID {
			t.Error("foreign key was wrong value", a.ID, first.AccountID)
		}
		if a.ID != second.AccountID {
			t.Error("foreign key was wrong value", a.ID, second.AccountID)
		}

		if first.R.Account != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Account != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Events[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Events[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Events().Count(tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testAccountToManyAddOpUsers(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Account
	var b, c, d, e User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, accountDBTypes, false, strmangle.SetComplement(accountPrimaryKeyColumns, accountColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*User{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*User{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUsers(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.AccountID {
			t.Error("foreign key was wrong value", a.ID, first.AccountID)
		}
		if a.ID != second.AccountID {
			t.Error("foreign key was wrong value", a.ID, second.AccountID)
		}

		if first.R.Account != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Account != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Users[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Users[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Users().Count(tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testAccountsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
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

func testAccountsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AccountSlice{o}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}

func testAccountsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Accounts().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	accountDBTypes = map[string]string{`CreatedAt`: `timestamp with time zone`, `DeletedAt`: `timestamp with time zone`, `Email`: `text`, `FirstName`: `text`, `ID`: `integer`, `LastName`: `text`, `UpdatedAt`: `timestamp with time zone`}
	_              = bytes.MinRead
)

func testAccountsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(accountPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(accountColumns) == len(accountPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Accounts().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, accountDBTypes, true, accountPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	if rowsAff, err := o.Update(tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testAccountsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(accountColumns) == len(accountPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Accounts().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, accountDBTypes, true, accountPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(accountColumns, accountPrimaryKeyColumns) {
		fields = accountColumns
	} else {
		fields = strmangle.SetComplement(
			accountColumns,
			accountPrimaryKeyColumns,
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

	slice := AccountSlice{o}
	if rowsAff, err := slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testAccountsUpsert(t *testing.T) {
	t.Parallel()

	if len(accountColumns) == len(accountPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Account{}
	if err = randomize.Struct(seed, &o, accountDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Account: %s", err)
	}

	count, err := Accounts().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, accountDBTypes, false, accountPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	if err = o.Upsert(tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Account: %s", err)
	}

	count, err = Accounts().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
