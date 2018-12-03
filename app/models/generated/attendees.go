// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Attendee is an object representing the database table.
type Attendee struct {
	ID            int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	TransactionID int       `boil:"transaction_id" json:"transaction_id" toml:"transaction_id" yaml:"transaction_id"`
	TicketID      int       `boil:"ticket_id" json:"ticket_id" toml:"ticket_id" yaml:"ticket_id"`
	Email         string    `boil:"email" json:"email" toml:"email" yaml:"email"`
	Status        string    `boil:"status" json:"status" toml:"status" yaml:"status"`
	CreatedAt     time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt     time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	DeletedAt     time.Time `boil:"deleted_at" json:"deleted_at" toml:"deleted_at" yaml:"deleted_at"`
	EventID       int       `boil:"event_id" json:"event_id" toml:"event_id" yaml:"event_id"`

	R *attendeeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L attendeeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AttendeeColumns = struct {
	ID            string
	TransactionID string
	TicketID      string
	Email         string
	Status        string
	CreatedAt     string
	UpdatedAt     string
	DeletedAt     string
	EventID       string
}{
	ID:            "id",
	TransactionID: "transaction_id",
	TicketID:      "ticket_id",
	Email:         "email",
	Status:        "status",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
	EventID:       "event_id",
}

// AttendeeRels is where relationship names are stored.
var AttendeeRels = struct {
	Ticket      string
	Event       string
	Transaction string
}{
	Ticket:      "Ticket",
	Event:       "Event",
	Transaction: "Transaction",
}

// attendeeR is where relationships are stored.
type attendeeR struct {
	Ticket      *Ticket
	Event       *Event
	Transaction *Transaction
}

// NewStruct creates a new relationship struct
func (*attendeeR) NewStruct() *attendeeR {
	return &attendeeR{}
}

// attendeeL is where Load methods for each relationship are stored.
type attendeeL struct{}

var (
	attendeeColumns               = []string{"id", "transaction_id", "ticket_id", "email", "status", "created_at", "updated_at", "deleted_at", "event_id"}
	attendeeColumnsWithoutDefault = []string{"transaction_id", "ticket_id", "email", "status", "created_at", "updated_at", "deleted_at", "event_id"}
	attendeeColumnsWithDefault    = []string{"id"}
	attendeePrimaryKeyColumns     = []string{"id"}
)

type (
	// AttendeeSlice is an alias for a slice of pointers to Attendee.
	// This should generally be used opposed to []Attendee.
	AttendeeSlice []*Attendee
	// AttendeeHook is the signature for custom Attendee hook methods
	AttendeeHook func(boil.Executor, *Attendee) error

	attendeeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	attendeeType                 = reflect.TypeOf(&Attendee{})
	attendeeMapping              = queries.MakeStructMapping(attendeeType)
	attendeePrimaryKeyMapping, _ = queries.BindMapping(attendeeType, attendeeMapping, attendeePrimaryKeyColumns)
	attendeeInsertCacheMut       sync.RWMutex
	attendeeInsertCache          = make(map[string]insertCache)
	attendeeUpdateCacheMut       sync.RWMutex
	attendeeUpdateCache          = make(map[string]updateCache)
	attendeeUpsertCacheMut       sync.RWMutex
	attendeeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var attendeeBeforeInsertHooks []AttendeeHook
var attendeeBeforeUpdateHooks []AttendeeHook
var attendeeBeforeDeleteHooks []AttendeeHook
var attendeeBeforeUpsertHooks []AttendeeHook

var attendeeAfterInsertHooks []AttendeeHook
var attendeeAfterSelectHooks []AttendeeHook
var attendeeAfterUpdateHooks []AttendeeHook
var attendeeAfterDeleteHooks []AttendeeHook
var attendeeAfterUpsertHooks []AttendeeHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Attendee) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range attendeeBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Attendee) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range attendeeBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Attendee) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range attendeeBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Attendee) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range attendeeBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Attendee) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range attendeeAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Attendee) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range attendeeAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Attendee) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range attendeeAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Attendee) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range attendeeAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Attendee) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range attendeeAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAttendeeHook registers your hook function for all future operations.
func AddAttendeeHook(hookPoint boil.HookPoint, attendeeHook AttendeeHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		attendeeBeforeInsertHooks = append(attendeeBeforeInsertHooks, attendeeHook)
	case boil.BeforeUpdateHook:
		attendeeBeforeUpdateHooks = append(attendeeBeforeUpdateHooks, attendeeHook)
	case boil.BeforeDeleteHook:
		attendeeBeforeDeleteHooks = append(attendeeBeforeDeleteHooks, attendeeHook)
	case boil.BeforeUpsertHook:
		attendeeBeforeUpsertHooks = append(attendeeBeforeUpsertHooks, attendeeHook)
	case boil.AfterInsertHook:
		attendeeAfterInsertHooks = append(attendeeAfterInsertHooks, attendeeHook)
	case boil.AfterSelectHook:
		attendeeAfterSelectHooks = append(attendeeAfterSelectHooks, attendeeHook)
	case boil.AfterUpdateHook:
		attendeeAfterUpdateHooks = append(attendeeAfterUpdateHooks, attendeeHook)
	case boil.AfterDeleteHook:
		attendeeAfterDeleteHooks = append(attendeeAfterDeleteHooks, attendeeHook)
	case boil.AfterUpsertHook:
		attendeeAfterUpsertHooks = append(attendeeAfterUpsertHooks, attendeeHook)
	}
}

// One returns a single attendee record from the query.
func (q attendeeQuery) One(exec boil.Executor) (*Attendee, error) {
	o := &Attendee{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for attendees")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Attendee records from the query.
func (q attendeeQuery) All(exec boil.Executor) (AttendeeSlice, error) {
	var o []*Attendee

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Attendee slice")
	}

	if len(attendeeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Attendee records in the query.
func (q attendeeQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count attendees rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q attendeeQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if attendees exists")
	}

	return count > 0, nil
}

// Ticket pointed to by the foreign key.
func (o *Attendee) Ticket(mods ...qm.QueryMod) ticketQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.TicketID),
	}

	queryMods = append(queryMods, mods...)

	query := Tickets(queryMods...)
	queries.SetFrom(query.Query, "\"tickets\"")

	return query
}

// Event pointed to by the foreign key.
func (o *Attendee) Event(mods ...qm.QueryMod) eventQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.EventID),
	}

	queryMods = append(queryMods, mods...)

	query := Events(queryMods...)
	queries.SetFrom(query.Query, "\"events\"")

	return query
}

// Transaction pointed to by the foreign key.
func (o *Attendee) Transaction(mods ...qm.QueryMod) transactionQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.TransactionID),
	}

	queryMods = append(queryMods, mods...)

	query := Transactions(queryMods...)
	queries.SetFrom(query.Query, "\"transactions\"")

	return query
}

// LoadTicket allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (attendeeL) LoadTicket(e boil.Executor, singular bool, maybeAttendee interface{}, mods queries.Applicator) error {
	var slice []*Attendee
	var object *Attendee

	if singular {
		object = maybeAttendee.(*Attendee)
	} else {
		slice = *maybeAttendee.(*[]*Attendee)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &attendeeR{}
		}
		args = append(args, object.TicketID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &attendeeR{}
			}

			for _, a := range args {
				if a == obj.TicketID {
					continue Outer
				}
			}

			args = append(args, obj.TicketID)
		}
	}

	query := NewQuery(qm.From(`tickets`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Ticket")
	}

	var resultSlice []*Ticket
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Ticket")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for tickets")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for tickets")
	}

	if len(attendeeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Ticket = foreign
		if foreign.R == nil {
			foreign.R = &ticketR{}
		}
		foreign.R.Attendees = append(foreign.R.Attendees, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.TicketID == foreign.ID {
				local.R.Ticket = foreign
				if foreign.R == nil {
					foreign.R = &ticketR{}
				}
				foreign.R.Attendees = append(foreign.R.Attendees, local)
				break
			}
		}
	}

	return nil
}

// LoadEvent allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (attendeeL) LoadEvent(e boil.Executor, singular bool, maybeAttendee interface{}, mods queries.Applicator) error {
	var slice []*Attendee
	var object *Attendee

	if singular {
		object = maybeAttendee.(*Attendee)
	} else {
		slice = *maybeAttendee.(*[]*Attendee)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &attendeeR{}
		}
		args = append(args, object.EventID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &attendeeR{}
			}

			for _, a := range args {
				if a == obj.EventID {
					continue Outer
				}
			}

			args = append(args, obj.EventID)
		}
	}

	query := NewQuery(qm.From(`events`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Event")
	}

	var resultSlice []*Event
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Event")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for events")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for events")
	}

	if len(attendeeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Event = foreign
		if foreign.R == nil {
			foreign.R = &eventR{}
		}
		foreign.R.Attendees = append(foreign.R.Attendees, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.EventID == foreign.ID {
				local.R.Event = foreign
				if foreign.R == nil {
					foreign.R = &eventR{}
				}
				foreign.R.Attendees = append(foreign.R.Attendees, local)
				break
			}
		}
	}

	return nil
}

// LoadTransaction allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (attendeeL) LoadTransaction(e boil.Executor, singular bool, maybeAttendee interface{}, mods queries.Applicator) error {
	var slice []*Attendee
	var object *Attendee

	if singular {
		object = maybeAttendee.(*Attendee)
	} else {
		slice = *maybeAttendee.(*[]*Attendee)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &attendeeR{}
		}
		args = append(args, object.TransactionID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &attendeeR{}
			}

			for _, a := range args {
				if a == obj.TransactionID {
					continue Outer
				}
			}

			args = append(args, obj.TransactionID)
		}
	}

	query := NewQuery(qm.From(`transactions`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Transaction")
	}

	var resultSlice []*Transaction
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Transaction")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for transactions")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for transactions")
	}

	if len(attendeeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Transaction = foreign
		if foreign.R == nil {
			foreign.R = &transactionR{}
		}
		foreign.R.Attendees = append(foreign.R.Attendees, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.TransactionID == foreign.ID {
				local.R.Transaction = foreign
				if foreign.R == nil {
					foreign.R = &transactionR{}
				}
				foreign.R.Attendees = append(foreign.R.Attendees, local)
				break
			}
		}
	}

	return nil
}

// SetTicket of the attendee to the related item.
// Sets o.R.Ticket to related.
// Adds o to related.R.Attendees.
func (o *Attendee) SetTicket(exec boil.Executor, insert bool, related *Ticket) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"attendees\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"ticket_id"}),
		strmangle.WhereClause("\"", "\"", 2, attendeePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TicketID = related.ID
	if o.R == nil {
		o.R = &attendeeR{
			Ticket: related,
		}
	} else {
		o.R.Ticket = related
	}

	if related.R == nil {
		related.R = &ticketR{
			Attendees: AttendeeSlice{o},
		}
	} else {
		related.R.Attendees = append(related.R.Attendees, o)
	}

	return nil
}

// SetEvent of the attendee to the related item.
// Sets o.R.Event to related.
// Adds o to related.R.Attendees.
func (o *Attendee) SetEvent(exec boil.Executor, insert bool, related *Event) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"attendees\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"event_id"}),
		strmangle.WhereClause("\"", "\"", 2, attendeePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.EventID = related.ID
	if o.R == nil {
		o.R = &attendeeR{
			Event: related,
		}
	} else {
		o.R.Event = related
	}

	if related.R == nil {
		related.R = &eventR{
			Attendees: AttendeeSlice{o},
		}
	} else {
		related.R.Attendees = append(related.R.Attendees, o)
	}

	return nil
}

// SetTransaction of the attendee to the related item.
// Sets o.R.Transaction to related.
// Adds o to related.R.Attendees.
func (o *Attendee) SetTransaction(exec boil.Executor, insert bool, related *Transaction) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"attendees\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"transaction_id"}),
		strmangle.WhereClause("\"", "\"", 2, attendeePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TransactionID = related.ID
	if o.R == nil {
		o.R = &attendeeR{
			Transaction: related,
		}
	} else {
		o.R.Transaction = related
	}

	if related.R == nil {
		related.R = &transactionR{
			Attendees: AttendeeSlice{o},
		}
	} else {
		related.R.Attendees = append(related.R.Attendees, o)
	}

	return nil
}

// Attendees retrieves all the records using an executor.
func Attendees(mods ...qm.QueryMod) attendeeQuery {
	mods = append(mods, qm.From("\"attendees\""))
	return attendeeQuery{NewQuery(mods...)}
}

// FindAttendee retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAttendee(exec boil.Executor, iD int, selectCols ...string) (*Attendee, error) {
	attendeeObj := &Attendee{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"attendees\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, attendeeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from attendees")
	}

	return attendeeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Attendee) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no attendees provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	if o.UpdatedAt.IsZero() {
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(attendeeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	attendeeInsertCacheMut.RLock()
	cache, cached := attendeeInsertCache[key]
	attendeeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			attendeeColumns,
			attendeeColumnsWithDefault,
			attendeeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(attendeeType, attendeeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(attendeeType, attendeeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"attendees\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"attendees\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into attendees")
	}

	if !cached {
		attendeeInsertCacheMut.Lock()
		attendeeInsertCache[key] = cache
		attendeeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// Update uses an executor to update the Attendee.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Attendee) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	attendeeUpdateCacheMut.RLock()
	cache, cached := attendeeUpdateCache[key]
	attendeeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			attendeeColumns,
			attendeePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update attendees, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"attendees\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, attendeePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(attendeeType, attendeeMapping, append(wl, attendeePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.Exec(cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update attendees row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for attendees")
	}

	if !cached {
		attendeeUpdateCacheMut.Lock()
		attendeeUpdateCache[key] = cache
		attendeeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAll updates all rows with the specified column values.
func (q attendeeQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for attendees")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for attendees")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AttendeeSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), attendeePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"attendees\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, attendeePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in attendee slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all attendee")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Attendee) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no attendees provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(attendeeColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	attendeeUpsertCacheMut.RLock()
	cache, cached := attendeeUpsertCache[key]
	attendeeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			attendeeColumns,
			attendeeColumnsWithDefault,
			attendeeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			attendeeColumns,
			attendeePrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert attendees, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(attendeePrimaryKeyColumns))
			copy(conflict, attendeePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"attendees\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(attendeeType, attendeeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(attendeeType, attendeeMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert attendees")
	}

	if !cached {
		attendeeUpsertCacheMut.Lock()
		attendeeUpsertCache[key] = cache
		attendeeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// Delete deletes a single Attendee record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Attendee) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Attendee provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), attendeePrimaryKeyMapping)
	sql := "DELETE FROM \"attendees\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from attendees")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for attendees")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q attendeeQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no attendeeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from attendees")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for attendees")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AttendeeSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Attendee slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(attendeeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), attendeePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"attendees\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, attendeePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from attendee slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for attendees")
	}

	if len(attendeeAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Attendee) Reload(exec boil.Executor) error {
	ret, err := FindAttendee(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AttendeeSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AttendeeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), attendeePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"attendees\".* FROM \"attendees\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, attendeePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AttendeeSlice")
	}

	*o = slice

	return nil
}

// AttendeeExists checks if the Attendee row exists.
func AttendeeExists(exec boil.Executor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"attendees\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if attendees exists")
	}

	return exists, nil
}
