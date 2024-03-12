// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"suwasystem/backend/ent/community"
	"suwasystem/backend/ent/predicate"
	"suwasystem/backend/ent/transaction"
	"suwasystem/backend/ent/user"
	"suwasystem/backend/ent/wallet"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CommunityUpdate is the builder for updating Community entities.
type CommunityUpdate struct {
	config
	hooks    []Hook
	mutation *CommunityMutation
}

// Where appends a list predicates to the CommunityUpdate builder.
func (cu *CommunityUpdate) Where(ps ...predicate.Community) *CommunityUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *CommunityUpdate) SetName(s string) *CommunityUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cu *CommunityUpdate) SetNillableName(s *string) *CommunityUpdate {
	if s != nil {
		cu.SetName(*s)
	}
	return cu
}

// SetDescription sets the "description" field.
func (cu *CommunityUpdate) SetDescription(s string) *CommunityUpdate {
	cu.mutation.SetDescription(s)
	return cu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cu *CommunityUpdate) SetNillableDescription(s *string) *CommunityUpdate {
	if s != nil {
		cu.SetDescription(*s)
	}
	return cu
}

// ClearDescription clears the value of the "description" field.
func (cu *CommunityUpdate) ClearDescription() *CommunityUpdate {
	cu.mutation.ClearDescription()
	return cu
}

// SetRate sets the "rate" field.
func (cu *CommunityUpdate) SetRate(f float32) *CommunityUpdate {
	cu.mutation.ResetRate()
	cu.mutation.SetRate(f)
	return cu
}

// SetNillableRate sets the "rate" field if the given value is not nil.
func (cu *CommunityUpdate) SetNillableRate(f *float32) *CommunityUpdate {
	if f != nil {
		cu.SetRate(*f)
	}
	return cu
}

// AddRate adds f to the "rate" field.
func (cu *CommunityUpdate) AddRate(f float32) *CommunityUpdate {
	cu.mutation.AddRate(f)
	return cu
}

// SetTax sets the "tax" field.
func (cu *CommunityUpdate) SetTax(f float32) *CommunityUpdate {
	cu.mutation.ResetTax()
	cu.mutation.SetTax(f)
	return cu
}

// SetNillableTax sets the "tax" field if the given value is not nil.
func (cu *CommunityUpdate) SetNillableTax(f *float32) *CommunityUpdate {
	if f != nil {
		cu.SetTax(*f)
	}
	return cu
}

// AddTax adds f to the "tax" field.
func (cu *CommunityUpdate) AddTax(f float32) *CommunityUpdate {
	cu.mutation.AddTax(f)
	return cu
}

// SetCreatedAt sets the "created_at" field.
func (cu *CommunityUpdate) SetCreatedAt(t time.Time) *CommunityUpdate {
	cu.mutation.SetCreatedAt(t)
	return cu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cu *CommunityUpdate) SetNillableCreatedAt(t *time.Time) *CommunityUpdate {
	if t != nil {
		cu.SetCreatedAt(*t)
	}
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CommunityUpdate) SetUpdatedAt(t time.Time) *CommunityUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetOwnerUserID sets the "owner_user_id" field.
func (cu *CommunityUpdate) SetOwnerUserID(u uuid.UUID) *CommunityUpdate {
	cu.mutation.SetOwnerUserID(u)
	return cu
}

// SetNillableOwnerUserID sets the "owner_user_id" field if the given value is not nil.
func (cu *CommunityUpdate) SetNillableOwnerUserID(u *uuid.UUID) *CommunityUpdate {
	if u != nil {
		cu.SetOwnerUserID(*u)
	}
	return cu
}

// ClearOwnerUserID clears the value of the "owner_user_id" field.
func (cu *CommunityUpdate) ClearOwnerUserID() *CommunityUpdate {
	cu.mutation.ClearOwnerUserID()
	return cu
}

// AddWalletIDs adds the "wallets" edge to the Wallet entity by IDs.
func (cu *CommunityUpdate) AddWalletIDs(ids ...int) *CommunityUpdate {
	cu.mutation.AddWalletIDs(ids...)
	return cu
}

// AddWallets adds the "wallets" edges to the Wallet entity.
func (cu *CommunityUpdate) AddWallets(w ...*Wallet) *CommunityUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return cu.AddWalletIDs(ids...)
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (cu *CommunityUpdate) SetOwnerID(id uuid.UUID) *CommunityUpdate {
	cu.mutation.SetOwnerID(id)
	return cu
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (cu *CommunityUpdate) SetNillableOwnerID(id *uuid.UUID) *CommunityUpdate {
	if id != nil {
		cu = cu.SetOwnerID(*id)
	}
	return cu
}

// SetOwner sets the "owner" edge to the User entity.
func (cu *CommunityUpdate) SetOwner(u *User) *CommunityUpdate {
	return cu.SetOwnerID(u.ID)
}

// AddFromTransactionIDs adds the "from_transactions" edge to the Transaction entity by IDs.
func (cu *CommunityUpdate) AddFromTransactionIDs(ids ...uuid.UUID) *CommunityUpdate {
	cu.mutation.AddFromTransactionIDs(ids...)
	return cu
}

// AddFromTransactions adds the "from_transactions" edges to the Transaction entity.
func (cu *CommunityUpdate) AddFromTransactions(t ...*Transaction) *CommunityUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cu.AddFromTransactionIDs(ids...)
}

// AddToTransactionIDs adds the "to_transactions" edge to the Transaction entity by IDs.
func (cu *CommunityUpdate) AddToTransactionIDs(ids ...uuid.UUID) *CommunityUpdate {
	cu.mutation.AddToTransactionIDs(ids...)
	return cu
}

// AddToTransactions adds the "to_transactions" edges to the Transaction entity.
func (cu *CommunityUpdate) AddToTransactions(t ...*Transaction) *CommunityUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cu.AddToTransactionIDs(ids...)
}

// Mutation returns the CommunityMutation object of the builder.
func (cu *CommunityUpdate) Mutation() *CommunityMutation {
	return cu.mutation
}

// ClearWallets clears all "wallets" edges to the Wallet entity.
func (cu *CommunityUpdate) ClearWallets() *CommunityUpdate {
	cu.mutation.ClearWallets()
	return cu
}

// RemoveWalletIDs removes the "wallets" edge to Wallet entities by IDs.
func (cu *CommunityUpdate) RemoveWalletIDs(ids ...int) *CommunityUpdate {
	cu.mutation.RemoveWalletIDs(ids...)
	return cu
}

// RemoveWallets removes "wallets" edges to Wallet entities.
func (cu *CommunityUpdate) RemoveWallets(w ...*Wallet) *CommunityUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return cu.RemoveWalletIDs(ids...)
}

// ClearOwner clears the "owner" edge to the User entity.
func (cu *CommunityUpdate) ClearOwner() *CommunityUpdate {
	cu.mutation.ClearOwner()
	return cu
}

// ClearFromTransactions clears all "from_transactions" edges to the Transaction entity.
func (cu *CommunityUpdate) ClearFromTransactions() *CommunityUpdate {
	cu.mutation.ClearFromTransactions()
	return cu
}

// RemoveFromTransactionIDs removes the "from_transactions" edge to Transaction entities by IDs.
func (cu *CommunityUpdate) RemoveFromTransactionIDs(ids ...uuid.UUID) *CommunityUpdate {
	cu.mutation.RemoveFromTransactionIDs(ids...)
	return cu
}

// RemoveFromTransactions removes "from_transactions" edges to Transaction entities.
func (cu *CommunityUpdate) RemoveFromTransactions(t ...*Transaction) *CommunityUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cu.RemoveFromTransactionIDs(ids...)
}

// ClearToTransactions clears all "to_transactions" edges to the Transaction entity.
func (cu *CommunityUpdate) ClearToTransactions() *CommunityUpdate {
	cu.mutation.ClearToTransactions()
	return cu
}

// RemoveToTransactionIDs removes the "to_transactions" edge to Transaction entities by IDs.
func (cu *CommunityUpdate) RemoveToTransactionIDs(ids ...uuid.UUID) *CommunityUpdate {
	cu.mutation.RemoveToTransactionIDs(ids...)
	return cu
}

// RemoveToTransactions removes "to_transactions" edges to Transaction entities.
func (cu *CommunityUpdate) RemoveToTransactions(t ...*Transaction) *CommunityUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cu.RemoveToTransactionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CommunityUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CommunityUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CommunityUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CommunityUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CommunityUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := community.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CommunityUpdate) check() error {
	if v, ok := cu.mutation.Rate(); ok {
		if err := community.RateValidator(v); err != nil {
			return &ValidationError{Name: "rate", err: fmt.Errorf(`ent: validator failed for field "Community.rate": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Tax(); ok {
		if err := community.TaxValidator(v); err != nil {
			return &ValidationError{Name: "tax", err: fmt.Errorf(`ent: validator failed for field "Community.tax": %w`, err)}
		}
	}
	return nil
}

func (cu *CommunityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(community.Table, community.Columns, sqlgraph.NewFieldSpec(community.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(community.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.Description(); ok {
		_spec.SetField(community.FieldDescription, field.TypeString, value)
	}
	if cu.mutation.DescriptionCleared() {
		_spec.ClearField(community.FieldDescription, field.TypeString)
	}
	if value, ok := cu.mutation.Rate(); ok {
		_spec.SetField(community.FieldRate, field.TypeFloat32, value)
	}
	if value, ok := cu.mutation.AddedRate(); ok {
		_spec.AddField(community.FieldRate, field.TypeFloat32, value)
	}
	if value, ok := cu.mutation.Tax(); ok {
		_spec.SetField(community.FieldTax, field.TypeFloat32, value)
	}
	if value, ok := cu.mutation.AddedTax(); ok {
		_spec.AddField(community.FieldTax, field.TypeFloat32, value)
	}
	if value, ok := cu.mutation.CreatedAt(); ok {
		_spec.SetField(community.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(community.FieldUpdatedAt, field.TypeTime, value)
	}
	if cu.mutation.WalletsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   community.WalletsTable,
			Columns: []string{community.WalletsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(wallet.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedWalletsIDs(); len(nodes) > 0 && !cu.mutation.WalletsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   community.WalletsTable,
			Columns: []string{community.WalletsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(wallet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.WalletsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   community.WalletsTable,
			Columns: []string{community.WalletsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(wallet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   community.OwnerTable,
			Columns: []string{community.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   community.OwnerTable,
			Columns: []string{community.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.FromTransactionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   community.FromTransactionsTable,
			Columns: []string{community.FromTransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedFromTransactionsIDs(); len(nodes) > 0 && !cu.mutation.FromTransactionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   community.FromTransactionsTable,
			Columns: []string{community.FromTransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.FromTransactionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   community.FromTransactionsTable,
			Columns: []string{community.FromTransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.ToTransactionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   community.ToTransactionsTable,
			Columns: []string{community.ToTransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedToTransactionsIDs(); len(nodes) > 0 && !cu.mutation.ToTransactionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   community.ToTransactionsTable,
			Columns: []string{community.ToTransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ToTransactionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   community.ToTransactionsTable,
			Columns: []string{community.ToTransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{community.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CommunityUpdateOne is the builder for updating a single Community entity.
type CommunityUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CommunityMutation
}

// SetName sets the "name" field.
func (cuo *CommunityUpdateOne) SetName(s string) *CommunityUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cuo *CommunityUpdateOne) SetNillableName(s *string) *CommunityUpdateOne {
	if s != nil {
		cuo.SetName(*s)
	}
	return cuo
}

// SetDescription sets the "description" field.
func (cuo *CommunityUpdateOne) SetDescription(s string) *CommunityUpdateOne {
	cuo.mutation.SetDescription(s)
	return cuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cuo *CommunityUpdateOne) SetNillableDescription(s *string) *CommunityUpdateOne {
	if s != nil {
		cuo.SetDescription(*s)
	}
	return cuo
}

// ClearDescription clears the value of the "description" field.
func (cuo *CommunityUpdateOne) ClearDescription() *CommunityUpdateOne {
	cuo.mutation.ClearDescription()
	return cuo
}

// SetRate sets the "rate" field.
func (cuo *CommunityUpdateOne) SetRate(f float32) *CommunityUpdateOne {
	cuo.mutation.ResetRate()
	cuo.mutation.SetRate(f)
	return cuo
}

// SetNillableRate sets the "rate" field if the given value is not nil.
func (cuo *CommunityUpdateOne) SetNillableRate(f *float32) *CommunityUpdateOne {
	if f != nil {
		cuo.SetRate(*f)
	}
	return cuo
}

// AddRate adds f to the "rate" field.
func (cuo *CommunityUpdateOne) AddRate(f float32) *CommunityUpdateOne {
	cuo.mutation.AddRate(f)
	return cuo
}

// SetTax sets the "tax" field.
func (cuo *CommunityUpdateOne) SetTax(f float32) *CommunityUpdateOne {
	cuo.mutation.ResetTax()
	cuo.mutation.SetTax(f)
	return cuo
}

// SetNillableTax sets the "tax" field if the given value is not nil.
func (cuo *CommunityUpdateOne) SetNillableTax(f *float32) *CommunityUpdateOne {
	if f != nil {
		cuo.SetTax(*f)
	}
	return cuo
}

// AddTax adds f to the "tax" field.
func (cuo *CommunityUpdateOne) AddTax(f float32) *CommunityUpdateOne {
	cuo.mutation.AddTax(f)
	return cuo
}

// SetCreatedAt sets the "created_at" field.
func (cuo *CommunityUpdateOne) SetCreatedAt(t time.Time) *CommunityUpdateOne {
	cuo.mutation.SetCreatedAt(t)
	return cuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cuo *CommunityUpdateOne) SetNillableCreatedAt(t *time.Time) *CommunityUpdateOne {
	if t != nil {
		cuo.SetCreatedAt(*t)
	}
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CommunityUpdateOne) SetUpdatedAt(t time.Time) *CommunityUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetOwnerUserID sets the "owner_user_id" field.
func (cuo *CommunityUpdateOne) SetOwnerUserID(u uuid.UUID) *CommunityUpdateOne {
	cuo.mutation.SetOwnerUserID(u)
	return cuo
}

// SetNillableOwnerUserID sets the "owner_user_id" field if the given value is not nil.
func (cuo *CommunityUpdateOne) SetNillableOwnerUserID(u *uuid.UUID) *CommunityUpdateOne {
	if u != nil {
		cuo.SetOwnerUserID(*u)
	}
	return cuo
}

// ClearOwnerUserID clears the value of the "owner_user_id" field.
func (cuo *CommunityUpdateOne) ClearOwnerUserID() *CommunityUpdateOne {
	cuo.mutation.ClearOwnerUserID()
	return cuo
}

// AddWalletIDs adds the "wallets" edge to the Wallet entity by IDs.
func (cuo *CommunityUpdateOne) AddWalletIDs(ids ...int) *CommunityUpdateOne {
	cuo.mutation.AddWalletIDs(ids...)
	return cuo
}

// AddWallets adds the "wallets" edges to the Wallet entity.
func (cuo *CommunityUpdateOne) AddWallets(w ...*Wallet) *CommunityUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return cuo.AddWalletIDs(ids...)
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (cuo *CommunityUpdateOne) SetOwnerID(id uuid.UUID) *CommunityUpdateOne {
	cuo.mutation.SetOwnerID(id)
	return cuo
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (cuo *CommunityUpdateOne) SetNillableOwnerID(id *uuid.UUID) *CommunityUpdateOne {
	if id != nil {
		cuo = cuo.SetOwnerID(*id)
	}
	return cuo
}

// SetOwner sets the "owner" edge to the User entity.
func (cuo *CommunityUpdateOne) SetOwner(u *User) *CommunityUpdateOne {
	return cuo.SetOwnerID(u.ID)
}

// AddFromTransactionIDs adds the "from_transactions" edge to the Transaction entity by IDs.
func (cuo *CommunityUpdateOne) AddFromTransactionIDs(ids ...uuid.UUID) *CommunityUpdateOne {
	cuo.mutation.AddFromTransactionIDs(ids...)
	return cuo
}

// AddFromTransactions adds the "from_transactions" edges to the Transaction entity.
func (cuo *CommunityUpdateOne) AddFromTransactions(t ...*Transaction) *CommunityUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cuo.AddFromTransactionIDs(ids...)
}

// AddToTransactionIDs adds the "to_transactions" edge to the Transaction entity by IDs.
func (cuo *CommunityUpdateOne) AddToTransactionIDs(ids ...uuid.UUID) *CommunityUpdateOne {
	cuo.mutation.AddToTransactionIDs(ids...)
	return cuo
}

// AddToTransactions adds the "to_transactions" edges to the Transaction entity.
func (cuo *CommunityUpdateOne) AddToTransactions(t ...*Transaction) *CommunityUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cuo.AddToTransactionIDs(ids...)
}

// Mutation returns the CommunityMutation object of the builder.
func (cuo *CommunityUpdateOne) Mutation() *CommunityMutation {
	return cuo.mutation
}

// ClearWallets clears all "wallets" edges to the Wallet entity.
func (cuo *CommunityUpdateOne) ClearWallets() *CommunityUpdateOne {
	cuo.mutation.ClearWallets()
	return cuo
}

// RemoveWalletIDs removes the "wallets" edge to Wallet entities by IDs.
func (cuo *CommunityUpdateOne) RemoveWalletIDs(ids ...int) *CommunityUpdateOne {
	cuo.mutation.RemoveWalletIDs(ids...)
	return cuo
}

// RemoveWallets removes "wallets" edges to Wallet entities.
func (cuo *CommunityUpdateOne) RemoveWallets(w ...*Wallet) *CommunityUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return cuo.RemoveWalletIDs(ids...)
}

// ClearOwner clears the "owner" edge to the User entity.
func (cuo *CommunityUpdateOne) ClearOwner() *CommunityUpdateOne {
	cuo.mutation.ClearOwner()
	return cuo
}

// ClearFromTransactions clears all "from_transactions" edges to the Transaction entity.
func (cuo *CommunityUpdateOne) ClearFromTransactions() *CommunityUpdateOne {
	cuo.mutation.ClearFromTransactions()
	return cuo
}

// RemoveFromTransactionIDs removes the "from_transactions" edge to Transaction entities by IDs.
func (cuo *CommunityUpdateOne) RemoveFromTransactionIDs(ids ...uuid.UUID) *CommunityUpdateOne {
	cuo.mutation.RemoveFromTransactionIDs(ids...)
	return cuo
}

// RemoveFromTransactions removes "from_transactions" edges to Transaction entities.
func (cuo *CommunityUpdateOne) RemoveFromTransactions(t ...*Transaction) *CommunityUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cuo.RemoveFromTransactionIDs(ids...)
}

// ClearToTransactions clears all "to_transactions" edges to the Transaction entity.
func (cuo *CommunityUpdateOne) ClearToTransactions() *CommunityUpdateOne {
	cuo.mutation.ClearToTransactions()
	return cuo
}

// RemoveToTransactionIDs removes the "to_transactions" edge to Transaction entities by IDs.
func (cuo *CommunityUpdateOne) RemoveToTransactionIDs(ids ...uuid.UUID) *CommunityUpdateOne {
	cuo.mutation.RemoveToTransactionIDs(ids...)
	return cuo
}

// RemoveToTransactions removes "to_transactions" edges to Transaction entities.
func (cuo *CommunityUpdateOne) RemoveToTransactions(t ...*Transaction) *CommunityUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cuo.RemoveToTransactionIDs(ids...)
}

// Where appends a list predicates to the CommunityUpdate builder.
func (cuo *CommunityUpdateOne) Where(ps ...predicate.Community) *CommunityUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CommunityUpdateOne) Select(field string, fields ...string) *CommunityUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Community entity.
func (cuo *CommunityUpdateOne) Save(ctx context.Context) (*Community, error) {
	cuo.defaults()
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CommunityUpdateOne) SaveX(ctx context.Context) *Community {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CommunityUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CommunityUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CommunityUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := community.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CommunityUpdateOne) check() error {
	if v, ok := cuo.mutation.Rate(); ok {
		if err := community.RateValidator(v); err != nil {
			return &ValidationError{Name: "rate", err: fmt.Errorf(`ent: validator failed for field "Community.rate": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Tax(); ok {
		if err := community.TaxValidator(v); err != nil {
			return &ValidationError{Name: "tax", err: fmt.Errorf(`ent: validator failed for field "Community.tax": %w`, err)}
		}
	}
	return nil
}

func (cuo *CommunityUpdateOne) sqlSave(ctx context.Context) (_node *Community, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(community.Table, community.Columns, sqlgraph.NewFieldSpec(community.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Community.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, community.FieldID)
		for _, f := range fields {
			if !community.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != community.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(community.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Description(); ok {
		_spec.SetField(community.FieldDescription, field.TypeString, value)
	}
	if cuo.mutation.DescriptionCleared() {
		_spec.ClearField(community.FieldDescription, field.TypeString)
	}
	if value, ok := cuo.mutation.Rate(); ok {
		_spec.SetField(community.FieldRate, field.TypeFloat32, value)
	}
	if value, ok := cuo.mutation.AddedRate(); ok {
		_spec.AddField(community.FieldRate, field.TypeFloat32, value)
	}
	if value, ok := cuo.mutation.Tax(); ok {
		_spec.SetField(community.FieldTax, field.TypeFloat32, value)
	}
	if value, ok := cuo.mutation.AddedTax(); ok {
		_spec.AddField(community.FieldTax, field.TypeFloat32, value)
	}
	if value, ok := cuo.mutation.CreatedAt(); ok {
		_spec.SetField(community.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(community.FieldUpdatedAt, field.TypeTime, value)
	}
	if cuo.mutation.WalletsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   community.WalletsTable,
			Columns: []string{community.WalletsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(wallet.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedWalletsIDs(); len(nodes) > 0 && !cuo.mutation.WalletsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   community.WalletsTable,
			Columns: []string{community.WalletsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(wallet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.WalletsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   community.WalletsTable,
			Columns: []string{community.WalletsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(wallet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   community.OwnerTable,
			Columns: []string{community.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   community.OwnerTable,
			Columns: []string{community.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.FromTransactionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   community.FromTransactionsTable,
			Columns: []string{community.FromTransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedFromTransactionsIDs(); len(nodes) > 0 && !cuo.mutation.FromTransactionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   community.FromTransactionsTable,
			Columns: []string{community.FromTransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.FromTransactionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   community.FromTransactionsTable,
			Columns: []string{community.FromTransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.ToTransactionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   community.ToTransactionsTable,
			Columns: []string{community.ToTransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedToTransactionsIDs(); len(nodes) > 0 && !cuo.mutation.ToTransactionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   community.ToTransactionsTable,
			Columns: []string{community.ToTransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ToTransactionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   community.ToTransactionsTable,
			Columns: []string{community.ToTransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Community{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{community.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
