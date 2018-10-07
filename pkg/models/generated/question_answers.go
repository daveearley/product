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

// QuestionAnswer is an object representing the database table.
type QuestionAnswer struct {
	ID         int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Answer     string    `boil:"answer" json:"answer" toml:"answer" yaml:"answer"`
	QuestionID int       `boil:"question_id" json:"question_id" toml:"question_id" yaml:"question_id"`
	CreatedAt  time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt  time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	DeletedAt  time.Time `boil:"deleted_at" json:"deleted_at" toml:"deleted_at" yaml:"deleted_at"`

	R *questionAnswerR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L questionAnswerL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var QuestionAnswerColumns = struct {
	ID         string
	Answer     string
	QuestionID string
	CreatedAt  string
	UpdatedAt  string
	DeletedAt  string
}{
	ID:         "id",
	Answer:     "answer",
	QuestionID: "question_id",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
}

// QuestionAnswerRels is where relationship names are stored.
var QuestionAnswerRels = struct {
	Question string
}{
	Question: "Question",
}

// questionAnswerR is where relationships are stored.
type questionAnswerR struct {
	Question *Question
}

// NewStruct creates a new relationship struct
func (*questionAnswerR) NewStruct() *questionAnswerR {
	return &questionAnswerR{}
}

// questionAnswerL is where Load methods for each relationship are stored.
type questionAnswerL struct{}

var (
	questionAnswerColumns               = []string{"id", "answer", "question_id", "created_at", "updated_at", "deleted_at"}
	questionAnswerColumnsWithoutDefault = []string{"answer", "question_id", "created_at", "updated_at", "deleted_at"}
	questionAnswerColumnsWithDefault    = []string{"id"}
	questionAnswerPrimaryKeyColumns     = []string{"id"}
)

type (
	// QuestionAnswerSlice is an alias for a slice of pointers to QuestionAnswer.
	// This should generally be used opposed to []QuestionAnswer.
	QuestionAnswerSlice []*QuestionAnswer
	// QuestionAnswerHook is the signature for custom QuestionAnswer hook methods
	QuestionAnswerHook func(boil.Executor, *QuestionAnswer) error

	questionAnswerQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	questionAnswerType                 = reflect.TypeOf(&QuestionAnswer{})
	questionAnswerMapping              = queries.MakeStructMapping(questionAnswerType)
	questionAnswerPrimaryKeyMapping, _ = queries.BindMapping(questionAnswerType, questionAnswerMapping, questionAnswerPrimaryKeyColumns)
	questionAnswerInsertCacheMut       sync.RWMutex
	questionAnswerInsertCache          = make(map[string]insertCache)
	questionAnswerUpdateCacheMut       sync.RWMutex
	questionAnswerUpdateCache          = make(map[string]updateCache)
	questionAnswerUpsertCacheMut       sync.RWMutex
	questionAnswerUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var questionAnswerBeforeInsertHooks []QuestionAnswerHook
var questionAnswerBeforeUpdateHooks []QuestionAnswerHook
var questionAnswerBeforeDeleteHooks []QuestionAnswerHook
var questionAnswerBeforeUpsertHooks []QuestionAnswerHook

var questionAnswerAfterInsertHooks []QuestionAnswerHook
var questionAnswerAfterSelectHooks []QuestionAnswerHook
var questionAnswerAfterUpdateHooks []QuestionAnswerHook
var questionAnswerAfterDeleteHooks []QuestionAnswerHook
var questionAnswerAfterUpsertHooks []QuestionAnswerHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *QuestionAnswer) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range questionAnswerBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *QuestionAnswer) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range questionAnswerBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *QuestionAnswer) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range questionAnswerBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *QuestionAnswer) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range questionAnswerBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *QuestionAnswer) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range questionAnswerAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *QuestionAnswer) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range questionAnswerAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *QuestionAnswer) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range questionAnswerAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *QuestionAnswer) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range questionAnswerAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *QuestionAnswer) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range questionAnswerAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddQuestionAnswerHook registers your hook function for all future operations.
func AddQuestionAnswerHook(hookPoint boil.HookPoint, questionAnswerHook QuestionAnswerHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		questionAnswerBeforeInsertHooks = append(questionAnswerBeforeInsertHooks, questionAnswerHook)
	case boil.BeforeUpdateHook:
		questionAnswerBeforeUpdateHooks = append(questionAnswerBeforeUpdateHooks, questionAnswerHook)
	case boil.BeforeDeleteHook:
		questionAnswerBeforeDeleteHooks = append(questionAnswerBeforeDeleteHooks, questionAnswerHook)
	case boil.BeforeUpsertHook:
		questionAnswerBeforeUpsertHooks = append(questionAnswerBeforeUpsertHooks, questionAnswerHook)
	case boil.AfterInsertHook:
		questionAnswerAfterInsertHooks = append(questionAnswerAfterInsertHooks, questionAnswerHook)
	case boil.AfterSelectHook:
		questionAnswerAfterSelectHooks = append(questionAnswerAfterSelectHooks, questionAnswerHook)
	case boil.AfterUpdateHook:
		questionAnswerAfterUpdateHooks = append(questionAnswerAfterUpdateHooks, questionAnswerHook)
	case boil.AfterDeleteHook:
		questionAnswerAfterDeleteHooks = append(questionAnswerAfterDeleteHooks, questionAnswerHook)
	case boil.AfterUpsertHook:
		questionAnswerAfterUpsertHooks = append(questionAnswerAfterUpsertHooks, questionAnswerHook)
	}
}

// One returns a single questionAnswer record from the query.
func (q questionAnswerQuery) One(exec boil.Executor) (*QuestionAnswer, error) {
	o := &QuestionAnswer{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for question_answers")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all QuestionAnswer records from the query.
func (q questionAnswerQuery) All(exec boil.Executor) (QuestionAnswerSlice, error) {
	var o []*QuestionAnswer

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to QuestionAnswer slice")
	}

	if len(questionAnswerAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all QuestionAnswer records in the query.
func (q questionAnswerQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count question_answers rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q questionAnswerQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if question_answers exists")
	}

	return count > 0, nil
}

// Question pointed to by the foreign key.
func (o *QuestionAnswer) Question(mods ...qm.QueryMod) questionQuery {
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
func (questionAnswerL) LoadQuestion(e boil.Executor, singular bool, maybeQuestionAnswer interface{}, mods queries.Applicator) error {
	var slice []*QuestionAnswer
	var object *QuestionAnswer

	if singular {
		object = maybeQuestionAnswer.(*QuestionAnswer)
	} else {
		slice = *maybeQuestionAnswer.(*[]*QuestionAnswer)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &questionAnswerR{}
		}
		args = append(args, object.QuestionID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &questionAnswerR{}
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

	if len(questionAnswerAfterSelectHooks) != 0 {
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
		foreign.R.QuestionAnswers = append(foreign.R.QuestionAnswers, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.QuestionID == foreign.ID {
				local.R.Question = foreign
				if foreign.R == nil {
					foreign.R = &questionR{}
				}
				foreign.R.QuestionAnswers = append(foreign.R.QuestionAnswers, local)
				break
			}
		}
	}

	return nil
}

// SetQuestion of the questionAnswer to the related item.
// Sets o.R.Question to related.
// Adds o to related.R.QuestionAnswers.
func (o *QuestionAnswer) SetQuestion(exec boil.Executor, insert bool, related *Question) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"question_answers\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"question_id"}),
		strmangle.WhereClause("\"", "\"", 2, questionAnswerPrimaryKeyColumns),
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
		o.R = &questionAnswerR{
			Question: related,
		}
	} else {
		o.R.Question = related
	}

	if related.R == nil {
		related.R = &questionR{
			QuestionAnswers: QuestionAnswerSlice{o},
		}
	} else {
		related.R.QuestionAnswers = append(related.R.QuestionAnswers, o)
	}

	return nil
}

// QuestionAnswers retrieves all the records using an executor.
func QuestionAnswers(mods ...qm.QueryMod) questionAnswerQuery {
	mods = append(mods, qm.From("\"question_answers\""))
	return questionAnswerQuery{NewQuery(mods...)}
}

// FindQuestionAnswer retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindQuestionAnswer(exec boil.Executor, iD int, selectCols ...string) (*QuestionAnswer, error) {
	questionAnswerObj := &QuestionAnswer{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"question_answers\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, questionAnswerObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from question_answers")
	}

	return questionAnswerObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *QuestionAnswer) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no question_answers provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(questionAnswerColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	questionAnswerInsertCacheMut.RLock()
	cache, cached := questionAnswerInsertCache[key]
	questionAnswerInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			questionAnswerColumns,
			questionAnswerColumnsWithDefault,
			questionAnswerColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(questionAnswerType, questionAnswerMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(questionAnswerType, questionAnswerMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"question_answers\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"question_answers\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into question_answers")
	}

	if !cached {
		questionAnswerInsertCacheMut.Lock()
		questionAnswerInsertCache[key] = cache
		questionAnswerInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// Update uses an executor to update the QuestionAnswer.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *QuestionAnswer) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	questionAnswerUpdateCacheMut.RLock()
	cache, cached := questionAnswerUpdateCache[key]
	questionAnswerUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			questionAnswerColumns,
			questionAnswerPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update question_answers, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"question_answers\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, questionAnswerPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(questionAnswerType, questionAnswerMapping, append(wl, questionAnswerPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update question_answers row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for question_answers")
	}

	if !cached {
		questionAnswerUpdateCacheMut.Lock()
		questionAnswerUpdateCache[key] = cache
		questionAnswerUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAll updates all rows with the specified column values.
func (q questionAnswerQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for question_answers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for question_answers")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o QuestionAnswerSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), questionAnswerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"question_answers\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, questionAnswerPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in questionAnswer slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all questionAnswer")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *QuestionAnswer) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no question_answers provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(questionAnswerColumnsWithDefault, o)

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

	questionAnswerUpsertCacheMut.RLock()
	cache, cached := questionAnswerUpsertCache[key]
	questionAnswerUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			questionAnswerColumns,
			questionAnswerColumnsWithDefault,
			questionAnswerColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			questionAnswerColumns,
			questionAnswerPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert question_answers, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(questionAnswerPrimaryKeyColumns))
			copy(conflict, questionAnswerPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"question_answers\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(questionAnswerType, questionAnswerMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(questionAnswerType, questionAnswerMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert question_answers")
	}

	if !cached {
		questionAnswerUpsertCacheMut.Lock()
		questionAnswerUpsertCache[key] = cache
		questionAnswerUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// Delete deletes a single QuestionAnswer record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *QuestionAnswer) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no QuestionAnswer provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), questionAnswerPrimaryKeyMapping)
	sql := "DELETE FROM \"question_answers\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from question_answers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for question_answers")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q questionAnswerQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no questionAnswerQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from question_answers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for question_answers")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o QuestionAnswerSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no QuestionAnswer slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(questionAnswerBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), questionAnswerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"question_answers\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, questionAnswerPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from questionAnswer slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for question_answers")
	}

	if len(questionAnswerAfterDeleteHooks) != 0 {
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
func (o *QuestionAnswer) Reload(exec boil.Executor) error {
	ret, err := FindQuestionAnswer(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *QuestionAnswerSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := QuestionAnswerSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), questionAnswerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"question_answers\".* FROM \"question_answers\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, questionAnswerPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in QuestionAnswerSlice")
	}

	*o = slice

	return nil
}

// QuestionAnswerExists checks if the QuestionAnswer row exists.
func QuestionAnswerExists(exec boil.Executor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"question_answers\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if question_answers exists")
	}

	return exists, nil
}