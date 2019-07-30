// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
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
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
	"github.com/volatiletech/sqlboiler/types"
)

// MessageLogs2 is an object representing the database table.
type MessageLogs2 struct {
	ID             int              `boil:"id" json:"id" toml:"id" yaml:"id"`
	LegacyID       int              `boil:"legacy_id" json:"legacy_id" toml:"legacy_id" yaml:"legacy_id"`
	GuildID        int64            `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	CreatedAt      time.Time        `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt      time.Time        `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	ChannelName    string           `boil:"channel_name" json:"channel_name" toml:"channel_name" yaml:"channel_name"`
	ChannelID      int64            `boil:"channel_id" json:"channel_id" toml:"channel_id" yaml:"channel_id"`
	AuthorID       int64            `boil:"author_id" json:"author_id" toml:"author_id" yaml:"author_id"`
	AuthorUsername string           `boil:"author_username" json:"author_username" toml:"author_username" yaml:"author_username"`
	Messages       types.Int64Array `boil:"messages" json:"messages,omitempty" toml:"messages" yaml:"messages,omitempty"`

	R *messageLogs2R `boil:"-" json:"-" toml:"-" yaml:"-"`
	L messageLogs2L  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MessageLogs2Columns = struct {
	ID             string
	LegacyID       string
	GuildID        string
	CreatedAt      string
	UpdatedAt      string
	ChannelName    string
	ChannelID      string
	AuthorID       string
	AuthorUsername string
	Messages       string
}{
	ID:             "id",
	LegacyID:       "legacy_id",
	GuildID:        "guild_id",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	ChannelName:    "channel_name",
	ChannelID:      "channel_id",
	AuthorID:       "author_id",
	AuthorUsername: "author_username",
	Messages:       "messages",
}

// Generated where

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var MessageLogs2Where = struct {
	ID             whereHelperint
	LegacyID       whereHelperint
	GuildID        whereHelperint64
	CreatedAt      whereHelpertime_Time
	UpdatedAt      whereHelpertime_Time
	ChannelName    whereHelperstring
	ChannelID      whereHelperint64
	AuthorID       whereHelperint64
	AuthorUsername whereHelperstring
	Messages       whereHelpertypes_Int64Array
}{
	ID:             whereHelperint{field: "\"message_logs2\".\"id\""},
	LegacyID:       whereHelperint{field: "\"message_logs2\".\"legacy_id\""},
	GuildID:        whereHelperint64{field: "\"message_logs2\".\"guild_id\""},
	CreatedAt:      whereHelpertime_Time{field: "\"message_logs2\".\"created_at\""},
	UpdatedAt:      whereHelpertime_Time{field: "\"message_logs2\".\"updated_at\""},
	ChannelName:    whereHelperstring{field: "\"message_logs2\".\"channel_name\""},
	ChannelID:      whereHelperint64{field: "\"message_logs2\".\"channel_id\""},
	AuthorID:       whereHelperint64{field: "\"message_logs2\".\"author_id\""},
	AuthorUsername: whereHelperstring{field: "\"message_logs2\".\"author_username\""},
	Messages:       whereHelpertypes_Int64Array{field: "\"message_logs2\".\"messages\""},
}

// MessageLogs2Rels is where relationship names are stored.
var MessageLogs2Rels = struct {
}{}

// messageLogs2R is where relationships are stored.
type messageLogs2R struct {
}

// NewStruct creates a new relationship struct
func (*messageLogs2R) NewStruct() *messageLogs2R {
	return &messageLogs2R{}
}

// messageLogs2L is where Load methods for each relationship are stored.
type messageLogs2L struct{}

var (
	messageLogs2AllColumns            = []string{"id", "legacy_id", "guild_id", "created_at", "updated_at", "channel_name", "channel_id", "author_id", "author_username", "messages"}
	messageLogs2ColumnsWithoutDefault = []string{"id", "legacy_id", "guild_id", "created_at", "updated_at", "channel_name", "channel_id", "author_id", "author_username", "messages"}
	messageLogs2ColumnsWithDefault    = []string{}
	messageLogs2PrimaryKeyColumns     = []string{"id"}
)

type (
	// MessageLogs2Slice is an alias for a slice of pointers to MessageLogs2.
	// This should generally be used opposed to []MessageLogs2.
	MessageLogs2Slice []*MessageLogs2

	messageLogs2Query struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	messageLogs2Type                 = reflect.TypeOf(&MessageLogs2{})
	messageLogs2Mapping              = queries.MakeStructMapping(messageLogs2Type)
	messageLogs2PrimaryKeyMapping, _ = queries.BindMapping(messageLogs2Type, messageLogs2Mapping, messageLogs2PrimaryKeyColumns)
	messageLogs2InsertCacheMut       sync.RWMutex
	messageLogs2InsertCache          = make(map[string]insertCache)
	messageLogs2UpdateCacheMut       sync.RWMutex
	messageLogs2UpdateCache          = make(map[string]updateCache)
	messageLogs2UpsertCacheMut       sync.RWMutex
	messageLogs2UpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single messageLogs2 record from the query using the global executor.
func (q messageLogs2Query) OneG(ctx context.Context) (*MessageLogs2, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single messageLogs2 record from the query.
func (q messageLogs2Query) One(ctx context.Context, exec boil.ContextExecutor) (*MessageLogs2, error) {
	o := &MessageLogs2{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for message_logs2")
	}

	return o, nil
}

// AllG returns all MessageLogs2 records from the query using the global executor.
func (q messageLogs2Query) AllG(ctx context.Context) (MessageLogs2Slice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all MessageLogs2 records from the query.
func (q messageLogs2Query) All(ctx context.Context, exec boil.ContextExecutor) (MessageLogs2Slice, error) {
	var o []*MessageLogs2

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to MessageLogs2 slice")
	}

	return o, nil
}

// CountG returns the count of all MessageLogs2 records in the query, and panics on error.
func (q messageLogs2Query) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all MessageLogs2 records in the query.
func (q messageLogs2Query) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count message_logs2 rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q messageLogs2Query) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q messageLogs2Query) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if message_logs2 exists")
	}

	return count > 0, nil
}

// MessageLogs2s retrieves all the records using an executor.
func MessageLogs2s(mods ...qm.QueryMod) messageLogs2Query {
	mods = append(mods, qm.From("\"message_logs2\""))
	return messageLogs2Query{NewQuery(mods...)}
}

// FindMessageLogs2G retrieves a single record by ID.
func FindMessageLogs2G(ctx context.Context, iD int, selectCols ...string) (*MessageLogs2, error) {
	return FindMessageLogs2(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindMessageLogs2 retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMessageLogs2(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*MessageLogs2, error) {
	messageLogs2Obj := &MessageLogs2{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"message_logs2\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, messageLogs2Obj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from message_logs2")
	}

	return messageLogs2Obj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *MessageLogs2) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *MessageLogs2) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no message_logs2 provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(messageLogs2ColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	messageLogs2InsertCacheMut.RLock()
	cache, cached := messageLogs2InsertCache[key]
	messageLogs2InsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			messageLogs2AllColumns,
			messageLogs2ColumnsWithDefault,
			messageLogs2ColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(messageLogs2Type, messageLogs2Mapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(messageLogs2Type, messageLogs2Mapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"message_logs2\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"message_logs2\" %sDEFAULT VALUES%s"
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
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into message_logs2")
	}

	if !cached {
		messageLogs2InsertCacheMut.Lock()
		messageLogs2InsertCache[key] = cache
		messageLogs2InsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single MessageLogs2 record using the global executor.
// See Update for more documentation.
func (o *MessageLogs2) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the MessageLogs2.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *MessageLogs2) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	messageLogs2UpdateCacheMut.RLock()
	cache, cached := messageLogs2UpdateCache[key]
	messageLogs2UpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			messageLogs2AllColumns,
			messageLogs2PrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update message_logs2, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"message_logs2\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, messageLogs2PrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(messageLogs2Type, messageLogs2Mapping, append(wl, messageLogs2PrimaryKeyColumns...))
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
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update message_logs2 row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for message_logs2")
	}

	if !cached {
		messageLogs2UpdateCacheMut.Lock()
		messageLogs2UpdateCache[key] = cache
		messageLogs2UpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q messageLogs2Query) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q messageLogs2Query) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for message_logs2")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for message_logs2")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o MessageLogs2Slice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MessageLogs2Slice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), messageLogs2PrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"message_logs2\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, messageLogs2PrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in messageLogs2 slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all messageLogs2")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *MessageLogs2) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *MessageLogs2) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no message_logs2 provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(messageLogs2ColumnsWithDefault, o)

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

	messageLogs2UpsertCacheMut.RLock()
	cache, cached := messageLogs2UpsertCache[key]
	messageLogs2UpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			messageLogs2AllColumns,
			messageLogs2ColumnsWithDefault,
			messageLogs2ColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			messageLogs2AllColumns,
			messageLogs2PrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert message_logs2, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(messageLogs2PrimaryKeyColumns))
			copy(conflict, messageLogs2PrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"message_logs2\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(messageLogs2Type, messageLogs2Mapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(messageLogs2Type, messageLogs2Mapping, ret)
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
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert message_logs2")
	}

	if !cached {
		messageLogs2UpsertCacheMut.Lock()
		messageLogs2UpsertCache[key] = cache
		messageLogs2UpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single MessageLogs2 record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *MessageLogs2) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single MessageLogs2 record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *MessageLogs2) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no MessageLogs2 provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), messageLogs2PrimaryKeyMapping)
	sql := "DELETE FROM \"message_logs2\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from message_logs2")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for message_logs2")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q messageLogs2Query) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no messageLogs2Query provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from message_logs2")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for message_logs2")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o MessageLogs2Slice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MessageLogs2Slice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), messageLogs2PrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"message_logs2\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, messageLogs2PrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from messageLogs2 slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for message_logs2")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *MessageLogs2) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no MessageLogs2 provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *MessageLogs2) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMessageLogs2(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MessageLogs2Slice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty MessageLogs2Slice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MessageLogs2Slice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MessageLogs2Slice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), messageLogs2PrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"message_logs2\".* FROM \"message_logs2\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, messageLogs2PrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in MessageLogs2Slice")
	}

	*o = slice

	return nil
}

// MessageLogs2ExistsG checks if the MessageLogs2 row exists.
func MessageLogs2ExistsG(ctx context.Context, iD int) (bool, error) {
	return MessageLogs2Exists(ctx, boil.GetContextDB(), iD)
}

// MessageLogs2Exists checks if the MessageLogs2 row exists.
func MessageLogs2Exists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"message_logs2\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if message_logs2 exists")
	}

	return exists, nil
}