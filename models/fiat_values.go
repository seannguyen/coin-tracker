// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// FiatValue is an object representing the database table.
type FiatValue struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	BalanceID int       `boil:"balance_id" json:"balance_id" toml:"balance_id" yaml:"balance_id"`
	Currency  string    `boil:"currency" json:"currency" toml:"currency" yaml:"currency"`
	Amount    float64   `boil:"amount" json:"amount" toml:"amount" yaml:"amount"`

	R *fiatValueR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L fiatValueL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var FiatValueColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	BalanceID string
	Currency  string
	Amount    string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	BalanceID: "balance_id",
	Currency:  "currency",
	Amount:    "amount",
}

// fiatValueR is where relationships are stored.
type fiatValueR struct {
	Balance *Balance
}

// fiatValueL is where Load methods for each relationship are stored.
type fiatValueL struct{}

var (
	fiatValueColumns               = []string{"id", "created_at", "updated_at", "balance_id", "currency", "amount"}
	fiatValueColumnsWithoutDefault = []string{"created_at", "updated_at", "balance_id", "currency", "amount"}
	fiatValueColumnsWithDefault    = []string{"id"}
	fiatValuePrimaryKeyColumns     = []string{"id"}
)

type (
	// FiatValueSlice is an alias for a slice of pointers to FiatValue.
	// This should generally be used opposed to []FiatValue.
	FiatValueSlice []*FiatValue
	// FiatValueHook is the signature for custom FiatValue hook methods
	FiatValueHook func(boil.Executor, *FiatValue) error

	fiatValueQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	fiatValueType                 = reflect.TypeOf(&FiatValue{})
	fiatValueMapping              = queries.MakeStructMapping(fiatValueType)
	fiatValuePrimaryKeyMapping, _ = queries.BindMapping(fiatValueType, fiatValueMapping, fiatValuePrimaryKeyColumns)
	fiatValueInsertCacheMut       sync.RWMutex
	fiatValueInsertCache          = make(map[string]insertCache)
	fiatValueUpdateCacheMut       sync.RWMutex
	fiatValueUpdateCache          = make(map[string]updateCache)
	fiatValueUpsertCacheMut       sync.RWMutex
	fiatValueUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var fiatValueBeforeInsertHooks []FiatValueHook
var fiatValueBeforeUpdateHooks []FiatValueHook
var fiatValueBeforeDeleteHooks []FiatValueHook
var fiatValueBeforeUpsertHooks []FiatValueHook

var fiatValueAfterInsertHooks []FiatValueHook
var fiatValueAfterSelectHooks []FiatValueHook
var fiatValueAfterUpdateHooks []FiatValueHook
var fiatValueAfterDeleteHooks []FiatValueHook
var fiatValueAfterUpsertHooks []FiatValueHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *FiatValue) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range fiatValueBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *FiatValue) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range fiatValueBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *FiatValue) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range fiatValueBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *FiatValue) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range fiatValueBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *FiatValue) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range fiatValueAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *FiatValue) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range fiatValueAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *FiatValue) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range fiatValueAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *FiatValue) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range fiatValueAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *FiatValue) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range fiatValueAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFiatValueHook registers your hook function for all future operations.
func AddFiatValueHook(hookPoint boil.HookPoint, fiatValueHook FiatValueHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		fiatValueBeforeInsertHooks = append(fiatValueBeforeInsertHooks, fiatValueHook)
	case boil.BeforeUpdateHook:
		fiatValueBeforeUpdateHooks = append(fiatValueBeforeUpdateHooks, fiatValueHook)
	case boil.BeforeDeleteHook:
		fiatValueBeforeDeleteHooks = append(fiatValueBeforeDeleteHooks, fiatValueHook)
	case boil.BeforeUpsertHook:
		fiatValueBeforeUpsertHooks = append(fiatValueBeforeUpsertHooks, fiatValueHook)
	case boil.AfterInsertHook:
		fiatValueAfterInsertHooks = append(fiatValueAfterInsertHooks, fiatValueHook)
	case boil.AfterSelectHook:
		fiatValueAfterSelectHooks = append(fiatValueAfterSelectHooks, fiatValueHook)
	case boil.AfterUpdateHook:
		fiatValueAfterUpdateHooks = append(fiatValueAfterUpdateHooks, fiatValueHook)
	case boil.AfterDeleteHook:
		fiatValueAfterDeleteHooks = append(fiatValueAfterDeleteHooks, fiatValueHook)
	case boil.AfterUpsertHook:
		fiatValueAfterUpsertHooks = append(fiatValueAfterUpsertHooks, fiatValueHook)
	}
}

// OneP returns a single fiatValue record from the query, and panics on error.
func (q fiatValueQuery) OneP() *FiatValue {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single fiatValue record from the query.
func (q fiatValueQuery) One() (*FiatValue, error) {
	o := &FiatValue{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for fiat_values")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all FiatValue records from the query, and panics on error.
func (q fiatValueQuery) AllP() FiatValueSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all FiatValue records from the query.
func (q fiatValueQuery) All() (FiatValueSlice, error) {
	var o []*FiatValue

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to FiatValue slice")
	}

	if len(fiatValueAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all FiatValue records in the query, and panics on error.
func (q fiatValueQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all FiatValue records in the query.
func (q fiatValueQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count fiat_values rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q fiatValueQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q fiatValueQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if fiat_values exists")
	}

	return count > 0, nil
}

// BalanceG pointed to by the foreign key.
func (o *FiatValue) BalanceG(mods ...qm.QueryMod) balanceQuery {
	return o.Balance(boil.GetDB(), mods...)
}

// Balance pointed to by the foreign key.
func (o *FiatValue) Balance(exec boil.Executor, mods ...qm.QueryMod) balanceQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.BalanceID),
	}

	queryMods = append(queryMods, mods...)

	query := Balances(exec, queryMods...)
	queries.SetFrom(query.Query, "\"balances\"")

	return query
} // LoadBalance allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (fiatValueL) LoadBalance(e boil.Executor, singular bool, maybeFiatValue interface{}) error {
	var slice []*FiatValue
	var object *FiatValue

	count := 1
	if singular {
		object = maybeFiatValue.(*FiatValue)
	} else {
		slice = *maybeFiatValue.(*[]*FiatValue)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &fiatValueR{}
		}
		args[0] = object.BalanceID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &fiatValueR{}
			}
			args[i] = obj.BalanceID
		}
	}

	query := fmt.Sprintf(
		"select * from \"balances\" where \"id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Balance")
	}
	defer results.Close()

	var resultSlice []*Balance
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Balance")
	}

	if len(fiatValueAfterSelectHooks) != 0 {
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
		object.R.Balance = resultSlice[0]
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.BalanceID == foreign.ID {
				local.R.Balance = foreign
				break
			}
		}
	}

	return nil
}

// SetBalanceG of the fiat_value to the related item.
// Sets o.R.Balance to related.
// Adds o to related.R.FiatValues.
// Uses the global database handle.
func (o *FiatValue) SetBalanceG(insert bool, related *Balance) error {
	return o.SetBalance(boil.GetDB(), insert, related)
}

// SetBalanceP of the fiat_value to the related item.
// Sets o.R.Balance to related.
// Adds o to related.R.FiatValues.
// Panics on error.
func (o *FiatValue) SetBalanceP(exec boil.Executor, insert bool, related *Balance) {
	if err := o.SetBalance(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetBalanceGP of the fiat_value to the related item.
// Sets o.R.Balance to related.
// Adds o to related.R.FiatValues.
// Uses the global database handle and panics on error.
func (o *FiatValue) SetBalanceGP(insert bool, related *Balance) {
	if err := o.SetBalance(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetBalance of the fiat_value to the related item.
// Sets o.R.Balance to related.
// Adds o to related.R.FiatValues.
func (o *FiatValue) SetBalance(exec boil.Executor, insert bool, related *Balance) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"fiat_values\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"balance_id"}),
		strmangle.WhereClause("\"", "\"", 2, fiatValuePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.BalanceID = related.ID

	if o.R == nil {
		o.R = &fiatValueR{
			Balance: related,
		}
	} else {
		o.R.Balance = related
	}

	if related.R == nil {
		related.R = &balanceR{
			FiatValues: FiatValueSlice{o},
		}
	} else {
		related.R.FiatValues = append(related.R.FiatValues, o)
	}

	return nil
}

// FiatValuesG retrieves all records.
func FiatValuesG(mods ...qm.QueryMod) fiatValueQuery {
	return FiatValues(boil.GetDB(), mods...)
}

// FiatValues retrieves all the records using an executor.
func FiatValues(exec boil.Executor, mods ...qm.QueryMod) fiatValueQuery {
	mods = append(mods, qm.From("\"fiat_values\""))
	return fiatValueQuery{NewQuery(exec, mods...)}
}

// FindFiatValueG retrieves a single record by ID.
func FindFiatValueG(id int, selectCols ...string) (*FiatValue, error) {
	return FindFiatValue(boil.GetDB(), id, selectCols...)
}

// FindFiatValueGP retrieves a single record by ID, and panics on error.
func FindFiatValueGP(id int, selectCols ...string) *FiatValue {
	retobj, err := FindFiatValue(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindFiatValue retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFiatValue(exec boil.Executor, id int, selectCols ...string) (*FiatValue, error) {
	fiatValueObj := &FiatValue{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"fiat_values\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(fiatValueObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from fiat_values")
	}

	return fiatValueObj, nil
}

// FindFiatValueP retrieves a single record by ID with an executor, and panics on error.
func FindFiatValueP(exec boil.Executor, id int, selectCols ...string) *FiatValue {
	retobj, err := FindFiatValue(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *FiatValue) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *FiatValue) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *FiatValue) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *FiatValue) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no fiat_values provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(fiatValueColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	fiatValueInsertCacheMut.RLock()
	cache, cached := fiatValueInsertCache[key]
	fiatValueInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			fiatValueColumns,
			fiatValueColumnsWithDefault,
			fiatValueColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(fiatValueType, fiatValueMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(fiatValueType, fiatValueMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"fiat_values\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"fiat_values\" DEFAULT VALUES"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		if len(wl) != 0 {
			cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
		}
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
		return errors.Wrap(err, "models: unable to insert into fiat_values")
	}

	if !cached {
		fiatValueInsertCacheMut.Lock()
		fiatValueInsertCache[key] = cache
		fiatValueInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single FiatValue record. See Update for
// whitelist behavior description.
func (o *FiatValue) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single FiatValue record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *FiatValue) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the FiatValue, and panics on error.
// See Update for whitelist behavior description.
func (o *FiatValue) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the FiatValue.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *FiatValue) Update(exec boil.Executor, whitelist ...string) error {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	fiatValueUpdateCacheMut.RLock()
	cache, cached := fiatValueUpdateCache[key]
	fiatValueUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			fiatValueColumns,
			fiatValuePrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update fiat_values, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"fiat_values\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, fiatValuePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(fiatValueType, fiatValueMapping, append(wl, fiatValuePrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update fiat_values row")
	}

	if !cached {
		fiatValueUpdateCacheMut.Lock()
		fiatValueUpdateCache[key] = cache
		fiatValueUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q fiatValueQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q fiatValueQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for fiat_values")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o FiatValueSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o FiatValueSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o FiatValueSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FiatValueSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), fiatValuePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"fiat_values\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, fiatValuePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in fiatValue slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *FiatValue) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *FiatValue) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *FiatValue) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *FiatValue) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no fiat_values provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(fiatValueColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
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
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	fiatValueUpsertCacheMut.RLock()
	cache, cached := fiatValueUpsertCache[key]
	fiatValueUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			fiatValueColumns,
			fiatValueColumnsWithDefault,
			fiatValueColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			fiatValueColumns,
			fiatValuePrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert fiat_values, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(fiatValuePrimaryKeyColumns))
			copy(conflict, fiatValuePrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"fiat_values\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(fiatValueType, fiatValueMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(fiatValueType, fiatValueMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert fiat_values")
	}

	if !cached {
		fiatValueUpsertCacheMut.Lock()
		fiatValueUpsertCache[key] = cache
		fiatValueUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single FiatValue record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *FiatValue) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single FiatValue record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *FiatValue) DeleteG() error {
	if o == nil {
		return errors.New("models: no FiatValue provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single FiatValue record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *FiatValue) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single FiatValue record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *FiatValue) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no FiatValue provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), fiatValuePrimaryKeyMapping)
	sql := "DELETE FROM \"fiat_values\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from fiat_values")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q fiatValueQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q fiatValueQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no fiatValueQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from fiat_values")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o FiatValueSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o FiatValueSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no FiatValue slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o FiatValueSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FiatValueSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no FiatValue slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(fiatValueBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), fiatValuePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"fiat_values\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, fiatValuePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from fiatValue slice")
	}

	if len(fiatValueAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *FiatValue) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *FiatValue) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *FiatValue) ReloadG() error {
	if o == nil {
		return errors.New("models: no FiatValue provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *FiatValue) Reload(exec boil.Executor) error {
	ret, err := FindFiatValue(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FiatValueSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FiatValueSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FiatValueSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty FiatValueSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FiatValueSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	fiatValues := FiatValueSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), fiatValuePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"fiat_values\".* FROM \"fiat_values\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, fiatValuePrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&fiatValues)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in FiatValueSlice")
	}

	*o = fiatValues

	return nil
}

// FiatValueExists checks if the FiatValue row exists.
func FiatValueExists(exec boil.Executor, id int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"fiat_values\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if fiat_values exists")
	}

	return exists, nil
}

// FiatValueExistsG checks if the FiatValue row exists.
func FiatValueExistsG(id int) (bool, error) {
	return FiatValueExists(boil.GetDB(), id)
}

// FiatValueExistsGP checks if the FiatValue row exists. Panics on error.
func FiatValueExistsGP(id int) bool {
	e, err := FiatValueExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// FiatValueExistsP checks if the FiatValue row exists. Panics on error.
func FiatValueExistsP(exec boil.Executor, id int) bool {
	e, err := FiatValueExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
