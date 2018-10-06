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

// TicketAttribute is an object representing the database table.
type TicketAttribute struct {
	ID          int `boil:"id" json:"id" toml:"id" yaml:"id"`
	TicketID    int `boil:"ticket_id" json:"ticket_id" toml:"ticket_id" yaml:"ticket_id"`
	AttributeID int `boil:"attribute_id" json:"attribute_id" toml:"attribute_id" yaml:"attribute_id"`

	R *ticketAttributeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L ticketAttributeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TicketAttributeColumns = struct {
	ID          string
	TicketID    string
	AttributeID string
}{
	ID:          "id",
	TicketID:    "ticket_id",
	AttributeID: "attribute_id",
}

// TicketAttributeRels is where relationship names are stored.
var TicketAttributeRels = struct {
	Ticket    string
	Attribute string
}{
	Ticket:    "Ticket",
	Attribute: "Attribute",
}

// ticketAttributeR is where relationships are stored.
type ticketAttributeR struct {
	Ticket    *Ticket
	Attribute *Attribute
}

// NewStruct creates a new relationship struct
func (*ticketAttributeR) NewStruct() *ticketAttributeR {
	return &ticketAttributeR{}
}

// ticketAttributeL is where Load methods for each relationship are stored.
type ticketAttributeL struct{}

var (
	ticketAttributeColumns               = []string{"id", "ticket_id", "attribute_id"}
	ticketAttributeColumnsWithoutDefault = []string{"ticket_id", "attribute_id"}
	ticketAttributeColumnsWithDefault    = []string{"id"}
	ticketAttributePrimaryKeyColumns     = []string{"id"}
)

type (
	// TicketAttributeSlice is an alias for a slice of pointers to TicketAttribute.
	// This should generally be used opposed to []TicketAttribute.
	TicketAttributeSlice []*TicketAttribute
	// TicketAttributeHook is the signature for custom TicketAttribute hook methods
	TicketAttributeHook func(boil.Executor, *TicketAttribute) error

	ticketAttributeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	ticketAttributeType                 = reflect.TypeOf(&TicketAttribute{})
	ticketAttributeMapping              = queries.MakeStructMapping(ticketAttributeType)
	ticketAttributePrimaryKeyMapping, _ = queries.BindMapping(ticketAttributeType, ticketAttributeMapping, ticketAttributePrimaryKeyColumns)
	ticketAttributeInsertCacheMut       sync.RWMutex
	ticketAttributeInsertCache          = make(map[string]insertCache)
	ticketAttributeUpdateCacheMut       sync.RWMutex
	ticketAttributeUpdateCache          = make(map[string]updateCache)
	ticketAttributeUpsertCacheMut       sync.RWMutex
	ticketAttributeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var ticketAttributeBeforeInsertHooks []TicketAttributeHook
var ticketAttributeBeforeUpdateHooks []TicketAttributeHook
var ticketAttributeBeforeDeleteHooks []TicketAttributeHook
var ticketAttributeBeforeUpsertHooks []TicketAttributeHook

var ticketAttributeAfterInsertHooks []TicketAttributeHook
var ticketAttributeAfterSelectHooks []TicketAttributeHook
var ticketAttributeAfterUpdateHooks []TicketAttributeHook
var ticketAttributeAfterDeleteHooks []TicketAttributeHook
var ticketAttributeAfterUpsertHooks []TicketAttributeHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TicketAttribute) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range ticketAttributeBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TicketAttribute) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range ticketAttributeBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TicketAttribute) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range ticketAttributeBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TicketAttribute) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range ticketAttributeBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TicketAttribute) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range ticketAttributeAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TicketAttribute) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range ticketAttributeAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TicketAttribute) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range ticketAttributeAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TicketAttribute) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range ticketAttributeAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TicketAttribute) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range ticketAttributeAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTicketAttributeHook registers your hook function for all future operations.
func AddTicketAttributeHook(hookPoint boil.HookPoint, ticketAttributeHook TicketAttributeHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		ticketAttributeBeforeInsertHooks = append(ticketAttributeBeforeInsertHooks, ticketAttributeHook)
	case boil.BeforeUpdateHook:
		ticketAttributeBeforeUpdateHooks = append(ticketAttributeBeforeUpdateHooks, ticketAttributeHook)
	case boil.BeforeDeleteHook:
		ticketAttributeBeforeDeleteHooks = append(ticketAttributeBeforeDeleteHooks, ticketAttributeHook)
	case boil.BeforeUpsertHook:
		ticketAttributeBeforeUpsertHooks = append(ticketAttributeBeforeUpsertHooks, ticketAttributeHook)
	case boil.AfterInsertHook:
		ticketAttributeAfterInsertHooks = append(ticketAttributeAfterInsertHooks, ticketAttributeHook)
	case boil.AfterSelectHook:
		ticketAttributeAfterSelectHooks = append(ticketAttributeAfterSelectHooks, ticketAttributeHook)
	case boil.AfterUpdateHook:
		ticketAttributeAfterUpdateHooks = append(ticketAttributeAfterUpdateHooks, ticketAttributeHook)
	case boil.AfterDeleteHook:
		ticketAttributeAfterDeleteHooks = append(ticketAttributeAfterDeleteHooks, ticketAttributeHook)
	case boil.AfterUpsertHook:
		ticketAttributeAfterUpsertHooks = append(ticketAttributeAfterUpsertHooks, ticketAttributeHook)
	}
}

// One returns a single ticketAttribute record from the query.
func (q ticketAttributeQuery) One(exec boil.Executor) (*TicketAttribute, error) {
	o := &TicketAttribute{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for ticket_attributes")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all TicketAttribute records from the query.
func (q ticketAttributeQuery) All(exec boil.Executor) (TicketAttributeSlice, error) {
	var o []*TicketAttribute

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TicketAttribute slice")
	}

	if len(ticketAttributeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all TicketAttribute records in the query.
func (q ticketAttributeQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count ticket_attributes rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q ticketAttributeQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if ticket_attributes exists")
	}

	return count > 0, nil
}

// Ticket pointed to by the foreign key.
func (o *TicketAttribute) Ticket(mods ...qm.QueryMod) ticketQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.TicketID),
	}

	queryMods = append(queryMods, mods...)

	query := Tickets(queryMods...)
	queries.SetFrom(query.Query, "\"tickets\"")

	return query
}

// Attribute pointed to by the foreign key.
func (o *TicketAttribute) Attribute(mods ...qm.QueryMod) attributeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.AttributeID),
	}

	queryMods = append(queryMods, mods...)

	query := Attributes(queryMods...)
	queries.SetFrom(query.Query, "\"attributes\"")

	return query
}

// LoadTicket allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (ticketAttributeL) LoadTicket(e boil.Executor, singular bool, maybeTicketAttribute interface{}, mods queries.Applicator) error {
	var slice []*TicketAttribute
	var object *TicketAttribute

	if singular {
		object = maybeTicketAttribute.(*TicketAttribute)
	} else {
		slice = *maybeTicketAttribute.(*[]*TicketAttribute)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &ticketAttributeR{}
		}
		args = append(args, object.TicketID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &ticketAttributeR{}
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

	if len(ticketAttributeAfterSelectHooks) != 0 {
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
		foreign.R.TicketAttributes = append(foreign.R.TicketAttributes, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.TicketID == foreign.ID {
				local.R.Ticket = foreign
				if foreign.R == nil {
					foreign.R = &ticketR{}
				}
				foreign.R.TicketAttributes = append(foreign.R.TicketAttributes, local)
				break
			}
		}
	}

	return nil
}

// LoadAttribute allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (ticketAttributeL) LoadAttribute(e boil.Executor, singular bool, maybeTicketAttribute interface{}, mods queries.Applicator) error {
	var slice []*TicketAttribute
	var object *TicketAttribute

	if singular {
		object = maybeTicketAttribute.(*TicketAttribute)
	} else {
		slice = *maybeTicketAttribute.(*[]*TicketAttribute)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &ticketAttributeR{}
		}
		args = append(args, object.AttributeID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &ticketAttributeR{}
			}

			for _, a := range args {
				if a == obj.AttributeID {
					continue Outer
				}
			}

			args = append(args, obj.AttributeID)
		}
	}

	query := NewQuery(qm.From(`attributes`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Attribute")
	}

	var resultSlice []*Attribute
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Attribute")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for attributes")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for attributes")
	}

	if len(ticketAttributeAfterSelectHooks) != 0 {
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
		object.R.Attribute = foreign
		if foreign.R == nil {
			foreign.R = &attributeR{}
		}
		foreign.R.TicketAttributes = append(foreign.R.TicketAttributes, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.AttributeID == foreign.ID {
				local.R.Attribute = foreign
				if foreign.R == nil {
					foreign.R = &attributeR{}
				}
				foreign.R.TicketAttributes = append(foreign.R.TicketAttributes, local)
				break
			}
		}
	}

	return nil
}

// SetTicket of the ticketAttribute to the related item.
// Sets o.R.Ticket to related.
// Adds o to related.R.TicketAttributes.
func (o *TicketAttribute) SetTicket(exec boil.Executor, insert bool, related *Ticket) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"ticket_attributes\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"ticket_id"}),
		strmangle.WhereClause("\"", "\"", 2, ticketAttributePrimaryKeyColumns),
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
		o.R = &ticketAttributeR{
			Ticket: related,
		}
	} else {
		o.R.Ticket = related
	}

	if related.R == nil {
		related.R = &ticketR{
			TicketAttributes: TicketAttributeSlice{o},
		}
	} else {
		related.R.TicketAttributes = append(related.R.TicketAttributes, o)
	}

	return nil
}

// SetAttribute of the ticketAttribute to the related item.
// Sets o.R.Attribute to related.
// Adds o to related.R.TicketAttributes.
func (o *TicketAttribute) SetAttribute(exec boil.Executor, insert bool, related *Attribute) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"ticket_attributes\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"attribute_id"}),
		strmangle.WhereClause("\"", "\"", 2, ticketAttributePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.AttributeID = related.ID
	if o.R == nil {
		o.R = &ticketAttributeR{
			Attribute: related,
		}
	} else {
		o.R.Attribute = related
	}

	if related.R == nil {
		related.R = &attributeR{
			TicketAttributes: TicketAttributeSlice{o},
		}
	} else {
		related.R.TicketAttributes = append(related.R.TicketAttributes, o)
	}

	return nil
}

// TicketAttributes retrieves all the records using an executor.
func TicketAttributes(mods ...qm.QueryMod) ticketAttributeQuery {
	mods = append(mods, qm.From("\"ticket_attributes\""))
	return ticketAttributeQuery{NewQuery(mods...)}
}

// FindTicketAttribute retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTicketAttribute(exec boil.Executor, iD int, selectCols ...string) (*TicketAttribute, error) {
	ticketAttributeObj := &TicketAttribute{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"ticket_attributes\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, ticketAttributeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from ticket_attributes")
	}

	return ticketAttributeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TicketAttribute) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no ticket_attributes provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(ticketAttributeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	ticketAttributeInsertCacheMut.RLock()
	cache, cached := ticketAttributeInsertCache[key]
	ticketAttributeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			ticketAttributeColumns,
			ticketAttributeColumnsWithDefault,
			ticketAttributeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(ticketAttributeType, ticketAttributeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(ticketAttributeType, ticketAttributeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"ticket_attributes\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"ticket_attributes\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into ticket_attributes")
	}

	if !cached {
		ticketAttributeInsertCacheMut.Lock()
		ticketAttributeInsertCache[key] = cache
		ticketAttributeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// Update uses an executor to update the TicketAttribute.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TicketAttribute) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	ticketAttributeUpdateCacheMut.RLock()
	cache, cached := ticketAttributeUpdateCache[key]
	ticketAttributeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			ticketAttributeColumns,
			ticketAttributePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update ticket_attributes, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"ticket_attributes\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, ticketAttributePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(ticketAttributeType, ticketAttributeMapping, append(wl, ticketAttributePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update ticket_attributes row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for ticket_attributes")
	}

	if !cached {
		ticketAttributeUpdateCacheMut.Lock()
		ticketAttributeUpdateCache[key] = cache
		ticketAttributeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAll updates all rows with the specified column values.
func (q ticketAttributeQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for ticket_attributes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for ticket_attributes")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TicketAttributeSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ticketAttributePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"ticket_attributes\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, ticketAttributePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in ticketAttribute slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all ticketAttribute")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TicketAttribute) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no ticket_attributes provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(ticketAttributeColumnsWithDefault, o)

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

	ticketAttributeUpsertCacheMut.RLock()
	cache, cached := ticketAttributeUpsertCache[key]
	ticketAttributeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			ticketAttributeColumns,
			ticketAttributeColumnsWithDefault,
			ticketAttributeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			ticketAttributeColumns,
			ticketAttributePrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert ticket_attributes, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(ticketAttributePrimaryKeyColumns))
			copy(conflict, ticketAttributePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"ticket_attributes\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(ticketAttributeType, ticketAttributeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(ticketAttributeType, ticketAttributeMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert ticket_attributes")
	}

	if !cached {
		ticketAttributeUpsertCacheMut.Lock()
		ticketAttributeUpsertCache[key] = cache
		ticketAttributeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// Delete deletes a single TicketAttribute record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TicketAttribute) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TicketAttribute provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), ticketAttributePrimaryKeyMapping)
	sql := "DELETE FROM \"ticket_attributes\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from ticket_attributes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for ticket_attributes")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q ticketAttributeQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no ticketAttributeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from ticket_attributes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for ticket_attributes")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TicketAttributeSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TicketAttribute slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(ticketAttributeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ticketAttributePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"ticket_attributes\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, ticketAttributePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from ticketAttribute slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for ticket_attributes")
	}

	if len(ticketAttributeAfterDeleteHooks) != 0 {
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
func (o *TicketAttribute) Reload(exec boil.Executor) error {
	ret, err := FindTicketAttribute(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TicketAttributeSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TicketAttributeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ticketAttributePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"ticket_attributes\".* FROM \"ticket_attributes\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, ticketAttributePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TicketAttributeSlice")
	}

	*o = slice

	return nil
}

// TicketAttributeExists checks if the TicketAttribute row exists.
func TicketAttributeExists(exec boil.Executor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"ticket_attributes\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if ticket_attributes exists")
	}

	return exists, nil
}
