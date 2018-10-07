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

func testAttendees(t *testing.T) {
	t.Parallel()

	query := Attendees()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testAttendeesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Attendee{}
	if err = randomize.Struct(seed, o, attendeeDBTypes, true, attendeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attendee struct: %s", err)
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

	count, err := Attendees().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAttendeesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Attendee{}
	if err = randomize.Struct(seed, o, attendeeDBTypes, true, attendeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attendee struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Attendees().DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Attendees().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAttendeesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Attendee{}
	if err = randomize.Struct(seed, o, attendeeDBTypes, true, attendeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attendee struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AttendeeSlice{o}

	if rowsAff, err := slice.DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Attendees().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAttendeesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Attendee{}
	if err = randomize.Struct(seed, o, attendeeDBTypes, true, attendeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attendee struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := AttendeeExists(tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Attendee exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AttendeeExists to return true, but got false.")
	}
}

func testAttendeesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Attendee{}
	if err = randomize.Struct(seed, o, attendeeDBTypes, true, attendeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attendee struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	attendeeFound, err := FindAttendee(tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if attendeeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testAttendeesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Attendee{}
	if err = randomize.Struct(seed, o, attendeeDBTypes, true, attendeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attendee struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Attendees().Bind(nil, tx, o); err != nil {
		t.Error(err)
	}
}

func testAttendeesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Attendee{}
	if err = randomize.Struct(seed, o, attendeeDBTypes, true, attendeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attendee struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Attendees().One(tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAttendeesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	attendeeOne := &Attendee{}
	attendeeTwo := &Attendee{}
	if err = randomize.Struct(seed, attendeeOne, attendeeDBTypes, false, attendeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attendee struct: %s", err)
	}
	if err = randomize.Struct(seed, attendeeTwo, attendeeDBTypes, false, attendeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attendee struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = attendeeOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = attendeeTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Attendees().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAttendeesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	attendeeOne := &Attendee{}
	attendeeTwo := &Attendee{}
	if err = randomize.Struct(seed, attendeeOne, attendeeDBTypes, false, attendeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attendee struct: %s", err)
	}
	if err = randomize.Struct(seed, attendeeTwo, attendeeDBTypes, false, attendeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attendee struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = attendeeOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = attendeeTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Attendees().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func attendeeBeforeInsertHook(e boil.Executor, o *Attendee) error {
	*o = Attendee{}
	return nil
}

func attendeeAfterInsertHook(e boil.Executor, o *Attendee) error {
	*o = Attendee{}
	return nil
}

func attendeeAfterSelectHook(e boil.Executor, o *Attendee) error {
	*o = Attendee{}
	return nil
}

func attendeeBeforeUpdateHook(e boil.Executor, o *Attendee) error {
	*o = Attendee{}
	return nil
}

func attendeeAfterUpdateHook(e boil.Executor, o *Attendee) error {
	*o = Attendee{}
	return nil
}

func attendeeBeforeDeleteHook(e boil.Executor, o *Attendee) error {
	*o = Attendee{}
	return nil
}

func attendeeAfterDeleteHook(e boil.Executor, o *Attendee) error {
	*o = Attendee{}
	return nil
}

func attendeeBeforeUpsertHook(e boil.Executor, o *Attendee) error {
	*o = Attendee{}
	return nil
}

func attendeeAfterUpsertHook(e boil.Executor, o *Attendee) error {
	*o = Attendee{}
	return nil
}

func testAttendeesHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Attendee{}
	o := &Attendee{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, attendeeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Attendee object: %s", err)
	}

	AddAttendeeHook(boil.BeforeInsertHook, attendeeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	attendeeBeforeInsertHooks = []AttendeeHook{}

	AddAttendeeHook(boil.AfterInsertHook, attendeeAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	attendeeAfterInsertHooks = []AttendeeHook{}

	AddAttendeeHook(boil.AfterSelectHook, attendeeAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	attendeeAfterSelectHooks = []AttendeeHook{}

	AddAttendeeHook(boil.BeforeUpdateHook, attendeeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	attendeeBeforeUpdateHooks = []AttendeeHook{}

	AddAttendeeHook(boil.AfterUpdateHook, attendeeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	attendeeAfterUpdateHooks = []AttendeeHook{}

	AddAttendeeHook(boil.BeforeDeleteHook, attendeeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	attendeeBeforeDeleteHooks = []AttendeeHook{}

	AddAttendeeHook(boil.AfterDeleteHook, attendeeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	attendeeAfterDeleteHooks = []AttendeeHook{}

	AddAttendeeHook(boil.BeforeUpsertHook, attendeeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	attendeeBeforeUpsertHooks = []AttendeeHook{}

	AddAttendeeHook(boil.AfterUpsertHook, attendeeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	attendeeAfterUpsertHooks = []AttendeeHook{}
}

func testAttendeesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Attendee{}
	if err = randomize.Struct(seed, o, attendeeDBTypes, true, attendeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attendee struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Attendees().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAttendeesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Attendee{}
	if err = randomize.Struct(seed, o, attendeeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Attendee struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Whitelist(attendeeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Attendees().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAttendeeToOneCustomerUsingCustomer(t *testing.T) {

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var local Attendee
	var foreign Customer

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, attendeeDBTypes, false, attendeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attendee struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, customerDBTypes, false, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	if err := foreign.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.CustomerID = foreign.ID
	if err := local.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Customer().One(tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := AttendeeSlice{&local}
	if err = local.L.LoadCustomer(tx, false, (*[]*Attendee)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Customer == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Customer = nil
	if err = local.L.LoadCustomer(tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Customer == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testAttendeeToOneTicketUsingTicket(t *testing.T) {

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var local Attendee
	var foreign Ticket

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, attendeeDBTypes, false, attendeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attendee struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, ticketDBTypes, false, ticketColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ticket struct: %s", err)
	}

	if err := foreign.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.TicketID = foreign.ID
	if err := local.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Ticket().One(tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := AttendeeSlice{&local}
	if err = local.L.LoadTicket(tx, false, (*[]*Attendee)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Ticket == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Ticket = nil
	if err = local.L.LoadTicket(tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Ticket == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testAttendeeToOneEventUsingEvent(t *testing.T) {

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var local Attendee
	var foreign Event

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, attendeeDBTypes, false, attendeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attendee struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, eventDBTypes, false, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	if err := foreign.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.EventID = foreign.ID
	if err := local.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Event().One(tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := AttendeeSlice{&local}
	if err = local.L.LoadEvent(tx, false, (*[]*Attendee)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Event == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Event = nil
	if err = local.L.LoadEvent(tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Event == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testAttendeeToOneSetOpCustomerUsingCustomer(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Attendee
	var b, c Customer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, attendeeDBTypes, false, strmangle.SetComplement(attendeePrimaryKeyColumns, attendeeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, customerDBTypes, false, strmangle.SetComplement(customerPrimaryKeyColumns, customerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, customerDBTypes, false, strmangle.SetComplement(customerPrimaryKeyColumns, customerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Customer{&b, &c} {
		err = a.SetCustomer(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Customer != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Attendees[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.CustomerID != x.ID {
			t.Error("foreign key was wrong value", a.CustomerID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.CustomerID))
		reflect.Indirect(reflect.ValueOf(&a.CustomerID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CustomerID != x.ID {
			t.Error("foreign key was wrong value", a.CustomerID, x.ID)
		}
	}
}
func testAttendeeToOneSetOpTicketUsingTicket(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Attendee
	var b, c Ticket

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, attendeeDBTypes, false, strmangle.SetComplement(attendeePrimaryKeyColumns, attendeeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, ticketDBTypes, false, strmangle.SetComplement(ticketPrimaryKeyColumns, ticketColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, ticketDBTypes, false, strmangle.SetComplement(ticketPrimaryKeyColumns, ticketColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Ticket{&b, &c} {
		err = a.SetTicket(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Ticket != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Attendees[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.TicketID != x.ID {
			t.Error("foreign key was wrong value", a.TicketID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.TicketID))
		reflect.Indirect(reflect.ValueOf(&a.TicketID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.TicketID != x.ID {
			t.Error("foreign key was wrong value", a.TicketID, x.ID)
		}
	}
}
func testAttendeeToOneSetOpEventUsingEvent(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Attendee
	var b, c Event

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, attendeeDBTypes, false, strmangle.SetComplement(attendeePrimaryKeyColumns, attendeeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, eventDBTypes, false, strmangle.SetComplement(eventPrimaryKeyColumns, eventColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, eventDBTypes, false, strmangle.SetComplement(eventPrimaryKeyColumns, eventColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Event{&b, &c} {
		err = a.SetEvent(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Event != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Attendees[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.EventID != x.ID {
			t.Error("foreign key was wrong value", a.EventID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.EventID))
		reflect.Indirect(reflect.ValueOf(&a.EventID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.EventID != x.ID {
			t.Error("foreign key was wrong value", a.EventID, x.ID)
		}
	}
}

func testAttendeesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Attendee{}
	if err = randomize.Struct(seed, o, attendeeDBTypes, true, attendeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attendee struct: %s", err)
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

func testAttendeesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Attendee{}
	if err = randomize.Struct(seed, o, attendeeDBTypes, true, attendeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attendee struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AttendeeSlice{o}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}

func testAttendeesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Attendee{}
	if err = randomize.Struct(seed, o, attendeeDBTypes, true, attendeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attendee struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Attendees().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	attendeeDBTypes = map[string]string{`CreatedAt`: `timestamp without time zone`, `CustomerID`: `integer`, `DeletedAt`: `timestamp without time zone`, `Email`: `character varying`, `EventID`: `integer`, `ID`: `integer`, `Status`: `enum.attendee_status('ACTIVE')`, `TicketID`: `integer`, `UpdatedAt`: `timestamp without time zone`}
	_               = bytes.MinRead
)

func testAttendeesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(attendeePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(attendeeColumns) == len(attendeePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Attendee{}
	if err = randomize.Struct(seed, o, attendeeDBTypes, true, attendeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attendee struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Attendees().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, attendeeDBTypes, true, attendeePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Attendee struct: %s", err)
	}

	if rowsAff, err := o.Update(tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testAttendeesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(attendeeColumns) == len(attendeePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Attendee{}
	if err = randomize.Struct(seed, o, attendeeDBTypes, true, attendeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attendee struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Attendees().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, attendeeDBTypes, true, attendeePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Attendee struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(attendeeColumns, attendeePrimaryKeyColumns) {
		fields = attendeeColumns
	} else {
		fields = strmangle.SetComplement(
			attendeeColumns,
			attendeePrimaryKeyColumns,
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

	slice := AttendeeSlice{o}
	if rowsAff, err := slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testAttendeesUpsert(t *testing.T) {
	t.Parallel()

	if len(attendeeColumns) == len(attendeePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Attendee{}
	if err = randomize.Struct(seed, &o, attendeeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Attendee struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Attendee: %s", err)
	}

	count, err := Attendees().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, attendeeDBTypes, false, attendeePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Attendee struct: %s", err)
	}

	if err = o.Upsert(tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Attendee: %s", err)
	}

	count, err = Attendees().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}