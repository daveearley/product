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
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// QuestionOption is an object representing the database table.
type QuestionOption struct {
	ID         int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Title      string    `boil:"title" json:"title" toml:"title" yaml:"title"`
	QuestionID int       `boil:"question_id" json:"question_id" toml:"question_id" yaml:"question_id"`
	CreatedAt  time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt  null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *questionOptionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L questionOptionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var QuestionOptionColumns = struct {
	ID         string
	Title      string
	QuestionID string
	CreatedAt  string
	UpdatedAt  string
}{
	ID:         "id",
	Title:      "title",
	QuestionID: "question_id",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// QuestionOptionRels is where relationship names are stored.
var QuestionOptionRels = struct {
	Question string
}{
	Question: "Question",
}

// questionOptionR is where relationships are stored.
type questionOptionR struct {
	Question *Question
}

// NewStruct creates a new relationship struct
func (*questionOptionR) NewStruct() *questionOptionR {
	return &questionOptionR{}
}

// questionOptionL is where Load methods for each relationship are stored.
type questionOptionL struct{}

var (
	questionOptionColumns               = []string{"id", "title", "question_id", "created_at", "updated_at"}
	questionOptionColumnsWithoutDefault = []string{"title", "question_id", "created_at", "updated_at"}
	questionOptionColumnsWithDefault    = []string{"id"}
	questionOptionPrimaryKeyColumns     = []string{"id"}
)

type (
	// QuestionOptionSlice is an alias for a slice of pointers to QuestionOption.
	// This should generally be used opposed to []QuestionOption.
	QuestionOptionSlice []*QuestionOption
	// QuestionOptionHook is the signature for custom QuestionOption hook methods
	QuestionOptionHook func(boil.Executor, *QuestionOption) error

	questionOptionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	questionOptionType                 = reflect.TypeOf(&QuestionOption{})
	questionOptionMapping              = queries.MakeStructMapping(questionOptionType)
	questionOptionPrimaryKeyMapping, _ = queries.BindMapping(questionOptionType, questionOptionMapping, questionOptionPrimaryKeyColumns)
	questionOptionInsertCacheMut       sync.RWMutex
	questionOptionInsertCache          = make(map[string]insertCache)
	questionOptionUpdateCacheMut       sync.RWMutex
	questionOptionUpdateCache          = make(map[string]updateCache)
	questionOptionUpsertCacheMut       sync.RWMutex
	questionOptionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var questionOptionBeforeInsertHooks []QuestionOptionHook
var questionOptionBeforeUpdateHooks []QuestionOptionHook
var questionOptionBeforeDeleteHooks []QuestionOptionHook
var questionOptionBeforeUpsertHooks []QuestionOptionHook

var questionOptionAfterInsertHooks []QuestionOptionHook
var questionOptionAfterSelectHooks []QuestionOptionHook
var questionOptionAfterUpdateHooks []QuestionOptionHook
var questionOptionAfterDeleteHooks []QuestionOptionHook
var questionOptionAfterUpsertHooks []QuestionOptionHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *QuestionOption) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range questionOptionBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *QuestionOption) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range questionOptionBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *QuestionOption) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range questionOptionBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *QuestionOption) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range questionOptionBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *QuestionOption) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range questionOptionAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *QuestionOption) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range questionOptionAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *QuestionOption) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range questionOptionAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *QuestionOption) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range questionOptionAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *QuestionOption) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range questionOptionAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddQuestionOptionHook registers your hook function for all future operations.
func AddQuestionOptionHook(hookPoint boil.HookPoint, questionOptionHook QuestionOptionHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		questionOptionBeforeInsertHooks = append(questionOptionBeforeInsertHooks, questionOptionHook)
	case boil.BeforeUpdateHook:
		questionOptionBeforeUpdateHooks = append(questionOptionBeforeUpdateHooks, questionOptionHook)
	case boil.BeforeDeleteHook:
		questionOptionBeforeDeleteHooks = append(questionOptionBeforeDeleteHooks, questionOptionHook)
	case boil.BeforeUpsertHook:
		questionOptionBeforeUpsertHooks = append(questionOptionBeforeUpsertHooks, questionOptionHook)
	case boil.AfterInsertHook:
		questionOptionAfterInsertHooks = append(questionOptionAfterInsertHooks, questionOptionHook)
	case boil.AfterSelectHook:
		questionOptionAfterSelectHooks = append(questionOptionAfterSelectHooks, questionOptionHook)
	case boil.AfterUpdateHook:
		questionOptionAfterUpdateHooks = append(questionOptionAfterUpdateHooks, questionOptionHook)
	case boil.AfterDeleteHook:
		questionOptionAfterDeleteHooks = append(questionOptionAfterDeleteHooks, questionOptionHook)
	case boil.AfterUpsertHook:
		questionOptionAfterUpsertHooks = append(questionOptionAfterUpsertHooks, questionOptionHook)
	}
}

// One returns a single questionOption record from the query.
func (q questionOptionQuery) One(exec boil.Executor) (*QuestionOption, error) {
	o := &QuestionOption{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for question_options")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all QuestionOption records from the query.
func (q questionOptionQuery) All(exec boil.Executor) (QuestionOptionSlice, error) {
	var o []*QuestionOption

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to QuestionOption slice")
	}

	if len(questionOptionAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all QuestionOption records in the query.
func (q questionOptionQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count question_options rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q questionOptionQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if question_options exists")
	}

	return count > 0, nil
}

// Question pointed to by the foreign key.
func (o *QuestionOption) Question(mods ...qm.QueryMod) questionQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.QuestionID),
	}

	queryMods = append(queryMods, mods...)

	query := Questions(queryMods...)
	queries.SetFrom(query.Query, "\"questions\"")

	return query
}

// LoadQuestion allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (questionOptionL) LoadQuestion(e boil.Executor, singular bool, maybeQuestionOption interface{}, mods queries.Applicator) error {
	var slice []*QuestionOption
	var object *QuestionOption

	if singular {
		object = maybeQuestionOption.(*QuestionOption)
	} else {
		slice = *maybeQuestionOption.(*[]*QuestionOption)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &questionOptionR{}
		}
		args = append(args, object.QuestionID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &questionOptionR{}
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

	if len(questionOptionAfterSelectHooks) != 0 {
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
		foreign.R.QuestionOptions = append(foreign.R.QuestionOptions, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.QuestionID == foreign.ID {
				local.R.Question = foreign
				if foreign.R == nil {
					foreign.R = &questionR{}
				}
				foreign.R.QuestionOptions = append(foreign.R.QuestionOptions, local)
				break
			}
		}
	}

	return nil
}

// SetQuestion of the questionOption to the related item.
// Sets o.R.Question to related.
// Adds o to related.R.QuestionOptions.
func (o *QuestionOption) SetQuestion(exec boil.Executor, insert bool, related *Question) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"question_options\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"question_id"}),
		strmangle.WhereClause("\"", "\"", 2, questionOptionPrimaryKeyColumns),
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
		o.R = &questionOptionR{
			Question: related,
		}
	} else {
		o.R.Question = related
	}

	if related.R == nil {
		related.R = &questionR{
			QuestionOptions: QuestionOptionSlice{o},
		}
	} else {
		related.R.QuestionOptions = append(related.R.QuestionOptions, o)
	}

	return nil
}

// QuestionOptions retrieves all the records using an executor.
func QuestionOptions(mods ...qm.QueryMod) questionOptionQuery {
	mods = append(mods, qm.From("\"question_options\""))
	return questionOptionQuery{NewQuery(mods...)}
}

// FindQuestionOption retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindQuestionOption(exec boil.Executor, iD int, selectCols ...string) (*QuestionOption, error) {
	questionOptionObj := &QuestionOption{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"question_options\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, questionOptionObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from question_options")
	}

	return questionOptionObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *QuestionOption) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no question_options provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	if queries.MustTime(o.UpdatedAt).IsZero() {
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(questionOptionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	questionOptionInsertCacheMut.RLock()
	cache, cached := questionOptionInsertCache[key]
	questionOptionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			questionOptionColumns,
			questionOptionColumnsWithDefault,
			questionOptionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(questionOptionType, questionOptionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(questionOptionType, questionOptionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"question_options\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"question_options\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into question_options")
	}

	if !cached {
		questionOptionInsertCacheMut.Lock()
		questionOptionInsertCache[key] = cache
		questionOptionInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// Update uses an executor to update the QuestionOption.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *QuestionOption) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	queries.SetScanner(&o.UpdatedAt, currTime)

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	questionOptionUpdateCacheMut.RLock()
	cache, cached := questionOptionUpdateCache[key]
	questionOptionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			questionOptionColumns,
			questionOptionPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update question_options, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"question_options\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, questionOptionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(questionOptionType, questionOptionMapping, append(wl, questionOptionPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update question_options row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for question_options")
	}

	if !cached {
		questionOptionUpdateCacheMut.Lock()
		questionOptionUpdateCache[key] = cache
		questionOptionUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAll updates all rows with the specified column values.
func (q questionOptionQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for question_options")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for question_options")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o QuestionOptionSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), questionOptionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"question_options\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, questionOptionPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in questionOption slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all questionOption")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *QuestionOption) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no question_options provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	queries.SetScanner(&o.UpdatedAt, currTime)

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(questionOptionColumnsWithDefault, o)

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

	questionOptionUpsertCacheMut.RLock()
	cache, cached := questionOptionUpsertCache[key]
	questionOptionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			questionOptionColumns,
			questionOptionColumnsWithDefault,
			questionOptionColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			questionOptionColumns,
			questionOptionPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert question_options, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(questionOptionPrimaryKeyColumns))
			copy(conflict, questionOptionPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"question_options\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(questionOptionType, questionOptionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(questionOptionType, questionOptionMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert question_options")
	}

	if !cached {
		questionOptionUpsertCacheMut.Lock()
		questionOptionUpsertCache[key] = cache
		questionOptionUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// Delete deletes a single QuestionOption record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *QuestionOption) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no QuestionOption provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), questionOptionPrimaryKeyMapping)
	sql := "DELETE FROM \"question_options\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from question_options")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for question_options")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q questionOptionQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no questionOptionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from question_options")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for question_options")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o QuestionOptionSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no QuestionOption slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(questionOptionBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), questionOptionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"question_options\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, questionOptionPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from questionOption slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for question_options")
	}

	if len(questionOptionAfterDeleteHooks) != 0 {
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
func (o *QuestionOption) Reload(exec boil.Executor) error {
	ret, err := FindQuestionOption(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *QuestionOptionSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := QuestionOptionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), questionOptionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"question_options\".* FROM \"question_options\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, questionOptionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in QuestionOptionSlice")
	}

	*o = slice

	return nil
}

// QuestionOptionExists checks if the QuestionOption row exists.
func QuestionOptionExists(exec boil.Executor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"question_options\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if question_options exists")
	}

	return exists, nil
}
