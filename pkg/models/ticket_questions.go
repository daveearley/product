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

// TicketQuestion is an object representing the database table.
type TicketQuestion struct {
	ID         int `boil:"id" json:"id" toml:"id" yaml:"id"`
	TicketID   int `boil:"ticket_id" json:"ticket_id" toml:"ticket_id" yaml:"ticket_id"`
	QuestionID int `boil:"question_id" json:"question_id" toml:"question_id" yaml:"question_id"`

	R *ticketQuestionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L ticketQuestionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TicketQuestionColumns = struct {
	ID         string
	TicketID   string
	QuestionID string
}{
	ID:         "id",
	TicketID:   "ticket_id",
	QuestionID: "question_id",
}

// TicketQuestionRels is where relationship names are stored.
var TicketQuestionRels = struct {
	Ticket   string
	Question string
}{
	Ticket:   "Ticket",
	Question: "Question",
}

// ticketQuestionR is where relationships are stored.
type ticketQuestionR struct {
	Ticket   *Ticket
	Question *Question
}

// NewStruct creates a new relationship struct
func (*ticketQuestionR) NewStruct() *ticketQuestionR {
	return &ticketQuestionR{}
}

// ticketQuestionL is where Load methods for each relationship are stored.
type ticketQuestionL struct{}

var (
	ticketQuestionColumns               = []string{"id", "ticket_id", "question_id"}
	ticketQuestionColumnsWithoutDefault = []string{"ticket_id", "question_id"}
	ticketQuestionColumnsWithDefault    = []string{"id"}
	ticketQuestionPrimaryKeyColumns     = []string{"id"}
)

type (
	// TicketQuestionSlice is an alias for a slice of pointers to TicketQuestion.
	// This should generally be used opposed to []TicketQuestion.
	TicketQuestionSlice []*TicketQuestion
	// TicketQuestionHook is the signature for custom TicketQuestion hook methods
	TicketQuestionHook func(boil.Executor, *TicketQuestion) error

	ticketQuestionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	ticketQuestionType                 = reflect.TypeOf(&TicketQuestion{})
	ticketQuestionMapping              = queries.MakeStructMapping(ticketQuestionType)
	ticketQuestionPrimaryKeyMapping, _ = queries.BindMapping(ticketQuestionType, ticketQuestionMapping, ticketQuestionPrimaryKeyColumns)
	ticketQuestionInsertCacheMut       sync.RWMutex
	ticketQuestionInsertCache          = make(map[string]insertCache)
	ticketQuestionUpdateCacheMut       sync.RWMutex
	ticketQuestionUpdateCache          = make(map[string]updateCache)
	ticketQuestionUpsertCacheMut       sync.RWMutex
	ticketQuestionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var ticketQuestionBeforeInsertHooks []TicketQuestionHook
var ticketQuestionBeforeUpdateHooks []TicketQuestionHook
var ticketQuestionBeforeDeleteHooks []TicketQuestionHook
var ticketQuestionBeforeUpsertHooks []TicketQuestionHook

var ticketQuestionAfterInsertHooks []TicketQuestionHook
var ticketQuestionAfterSelectHooks []TicketQuestionHook
var ticketQuestionAfterUpdateHooks []TicketQuestionHook
var ticketQuestionAfterDeleteHooks []TicketQuestionHook
var ticketQuestionAfterUpsertHooks []TicketQuestionHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TicketQuestion) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range ticketQuestionBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TicketQuestion) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range ticketQuestionBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TicketQuestion) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range ticketQuestionBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TicketQuestion) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range ticketQuestionBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TicketQuestion) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range ticketQuestionAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TicketQuestion) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range ticketQuestionAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TicketQuestion) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range ticketQuestionAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TicketQuestion) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range ticketQuestionAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TicketQuestion) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range ticketQuestionAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTicketQuestionHook registers your hook function for all future operations.
func AddTicketQuestionHook(hookPoint boil.HookPoint, ticketQuestionHook TicketQuestionHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		ticketQuestionBeforeInsertHooks = append(ticketQuestionBeforeInsertHooks, ticketQuestionHook)
	case boil.BeforeUpdateHook:
		ticketQuestionBeforeUpdateHooks = append(ticketQuestionBeforeUpdateHooks, ticketQuestionHook)
	case boil.BeforeDeleteHook:
		ticketQuestionBeforeDeleteHooks = append(ticketQuestionBeforeDeleteHooks, ticketQuestionHook)
	case boil.BeforeUpsertHook:
		ticketQuestionBeforeUpsertHooks = append(ticketQuestionBeforeUpsertHooks, ticketQuestionHook)
	case boil.AfterInsertHook:
		ticketQuestionAfterInsertHooks = append(ticketQuestionAfterInsertHooks, ticketQuestionHook)
	case boil.AfterSelectHook:
		ticketQuestionAfterSelectHooks = append(ticketQuestionAfterSelectHooks, ticketQuestionHook)
	case boil.AfterUpdateHook:
		ticketQuestionAfterUpdateHooks = append(ticketQuestionAfterUpdateHooks, ticketQuestionHook)
	case boil.AfterDeleteHook:
		ticketQuestionAfterDeleteHooks = append(ticketQuestionAfterDeleteHooks, ticketQuestionHook)
	case boil.AfterUpsertHook:
		ticketQuestionAfterUpsertHooks = append(ticketQuestionAfterUpsertHooks, ticketQuestionHook)
	}
}

// One returns a single ticketQuestion record from the query.
func (q ticketQuestionQuery) One(exec boil.Executor) (*TicketQuestion, error) {
	o := &TicketQuestion{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for ticket_questions")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all TicketQuestion records from the query.
func (q ticketQuestionQuery) All(exec boil.Executor) (TicketQuestionSlice, error) {
	var o []*TicketQuestion

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TicketQuestion slice")
	}

	if len(ticketQuestionAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all TicketQuestion records in the query.
func (q ticketQuestionQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count ticket_questions rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q ticketQuestionQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if ticket_questions exists")
	}

	return count > 0, nil
}

// Ticket pointed to by the foreign key.
func (o *TicketQuestion) Ticket(mods ...qm.QueryMod) ticketQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.TicketID),
	}

	queryMods = append(queryMods, mods...)

	query := Tickets(queryMods...)
	queries.SetFrom(query.Query, "\"tickets\"")

	return query
}

// Question pointed to by the foreign key.
func (o *TicketQuestion) Question(mods ...qm.QueryMod) questionQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.QuestionID),
	}

	queryMods = append(queryMods, mods...)

	query := Questions(queryMods...)
	queries.SetFrom(query.Query, "\"questions\"")

	return query
}

// LoadTicket allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (ticketQuestionL) LoadTicket(e boil.Executor, singular bool, maybeTicketQuestion interface{}, mods queries.Applicator) error {
	var slice []*TicketQuestion
	var object *TicketQuestion

	if singular {
		object = maybeTicketQuestion.(*TicketQuestion)
	} else {
		slice = *maybeTicketQuestion.(*[]*TicketQuestion)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &ticketQuestionR{}
		}
		args = append(args, object.TicketID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &ticketQuestionR{}
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

	if len(ticketQuestionAfterSelectHooks) != 0 {
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
		foreign.R.TicketQuestions = append(foreign.R.TicketQuestions, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.TicketID == foreign.ID {
				local.R.Ticket = foreign
				if foreign.R == nil {
					foreign.R = &ticketR{}
				}
				foreign.R.TicketQuestions = append(foreign.R.TicketQuestions, local)
				break
			}
		}
	}

	return nil
}

// LoadQuestion allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (ticketQuestionL) LoadQuestion(e boil.Executor, singular bool, maybeTicketQuestion interface{}, mods queries.Applicator) error {
	var slice []*TicketQuestion
	var object *TicketQuestion

	if singular {
		object = maybeTicketQuestion.(*TicketQuestion)
	} else {
		slice = *maybeTicketQuestion.(*[]*TicketQuestion)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &ticketQuestionR{}
		}
		args = append(args, object.QuestionID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &ticketQuestionR{}
			}

			for _, a := range args {
				if a == obj.QuestionID {
					continue Outer
				}
			}

			args = append(args, obj.QuestionID)
		}
	}

	query := NewQuery(qm.From(`questions`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Question")
	}

	var resultSlice []*Question
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Question")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for questions")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for questions")
	}

	if len(ticketQuestionAfterSelectHooks) != 0 {
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
		object.R.Question = foreign
		if foreign.R == nil {
			foreign.R = &questionR{}
		}
		foreign.R.TicketQuestions = append(foreign.R.TicketQuestions, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.QuestionID == foreign.ID {
				local.R.Question = foreign
				if foreign.R == nil {
					foreign.R = &questionR{}
				}
				foreign.R.TicketQuestions = append(foreign.R.TicketQuestions, local)
				break
			}
		}
	}

	return nil
}

// SetTicket of the ticketQuestion to the related item.
// Sets o.R.Ticket to related.
// Adds o to related.R.TicketQuestions.
func (o *TicketQuestion) SetTicket(exec boil.Executor, insert bool, related *Ticket) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"ticket_questions\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"ticket_id"}),
		strmangle.WhereClause("\"", "\"", 2, ticketQuestionPrimaryKeyColumns),
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
		o.R = &ticketQuestionR{
			Ticket: related,
		}
	} else {
		o.R.Ticket = related
	}

	if related.R == nil {
		related.R = &ticketR{
			TicketQuestions: TicketQuestionSlice{o},
		}
	} else {
		related.R.TicketQuestions = append(related.R.TicketQuestions, o)
	}

	return nil
}

// SetQuestion of the ticketQuestion to the related item.
// Sets o.R.Question to related.
// Adds o to related.R.TicketQuestions.
func (o *TicketQuestion) SetQuestion(exec boil.Executor, insert bool, related *Question) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"ticket_questions\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"question_id"}),
		strmangle.WhereClause("\"", "\"", 2, ticketQuestionPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.QuestionID = related.ID
	if o.R == nil {
		o.R = &ticketQuestionR{
			Question: related,
		}
	} else {
		o.R.Question = related
	}

	if related.R == nil {
		related.R = &questionR{
			TicketQuestions: TicketQuestionSlice{o},
		}
	} else {
		related.R.TicketQuestions = append(related.R.TicketQuestions, o)
	}

	return nil
}

// TicketQuestions retrieves all the records using an executor.
func TicketQuestions(mods ...qm.QueryMod) ticketQuestionQuery {
	mods = append(mods, qm.From("\"ticket_questions\""))
	return ticketQuestionQuery{NewQuery(mods...)}
}

// FindTicketQuestion retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTicketQuestion(exec boil.Executor, iD int, selectCols ...string) (*TicketQuestion, error) {
	ticketQuestionObj := &TicketQuestion{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"ticket_questions\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, ticketQuestionObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from ticket_questions")
	}

	return ticketQuestionObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TicketQuestion) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no ticket_questions provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(ticketQuestionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	ticketQuestionInsertCacheMut.RLock()
	cache, cached := ticketQuestionInsertCache[key]
	ticketQuestionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			ticketQuestionColumns,
			ticketQuestionColumnsWithDefault,
			ticketQuestionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(ticketQuestionType, ticketQuestionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(ticketQuestionType, ticketQuestionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"ticket_questions\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"ticket_questions\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into ticket_questions")
	}

	if !cached {
		ticketQuestionInsertCacheMut.Lock()
		ticketQuestionInsertCache[key] = cache
		ticketQuestionInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// Update uses an executor to update the TicketQuestion.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TicketQuestion) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	ticketQuestionUpdateCacheMut.RLock()
	cache, cached := ticketQuestionUpdateCache[key]
	ticketQuestionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			ticketQuestionColumns,
			ticketQuestionPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update ticket_questions, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"ticket_questions\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, ticketQuestionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(ticketQuestionType, ticketQuestionMapping, append(wl, ticketQuestionPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update ticket_questions row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for ticket_questions")
	}

	if !cached {
		ticketQuestionUpdateCacheMut.Lock()
		ticketQuestionUpdateCache[key] = cache
		ticketQuestionUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAll updates all rows with the specified column values.
func (q ticketQuestionQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for ticket_questions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for ticket_questions")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TicketQuestionSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ticketQuestionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"ticket_questions\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, ticketQuestionPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in ticketQuestion slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all ticketQuestion")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TicketQuestion) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no ticket_questions provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(ticketQuestionColumnsWithDefault, o)

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

	ticketQuestionUpsertCacheMut.RLock()
	cache, cached := ticketQuestionUpsertCache[key]
	ticketQuestionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			ticketQuestionColumns,
			ticketQuestionColumnsWithDefault,
			ticketQuestionColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			ticketQuestionColumns,
			ticketQuestionPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert ticket_questions, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(ticketQuestionPrimaryKeyColumns))
			copy(conflict, ticketQuestionPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"ticket_questions\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(ticketQuestionType, ticketQuestionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(ticketQuestionType, ticketQuestionMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert ticket_questions")
	}

	if !cached {
		ticketQuestionUpsertCacheMut.Lock()
		ticketQuestionUpsertCache[key] = cache
		ticketQuestionUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// Delete deletes a single TicketQuestion record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TicketQuestion) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TicketQuestion provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), ticketQuestionPrimaryKeyMapping)
	sql := "DELETE FROM \"ticket_questions\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from ticket_questions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for ticket_questions")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q ticketQuestionQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no ticketQuestionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from ticket_questions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for ticket_questions")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TicketQuestionSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TicketQuestion slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(ticketQuestionBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ticketQuestionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"ticket_questions\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, ticketQuestionPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from ticketQuestion slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for ticket_questions")
	}

	if len(ticketQuestionAfterDeleteHooks) != 0 {
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
func (o *TicketQuestion) Reload(exec boil.Executor) error {
	ret, err := FindTicketQuestion(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TicketQuestionSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TicketQuestionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ticketQuestionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"ticket_questions\".* FROM \"ticket_questions\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, ticketQuestionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TicketQuestionSlice")
	}

	*o = slice

	return nil
}

// TicketQuestionExists checks if the TicketQuestion row exists.
func TicketQuestionExists(exec boil.Executor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"ticket_questions\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if ticket_questions exists")
	}

	return exists, nil
}
