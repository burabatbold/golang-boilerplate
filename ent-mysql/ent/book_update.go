// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"http/ent/book"
	"http/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BookUpdate is the builder for updating Book entities.
type BookUpdate struct {
	config
	hooks    []Hook
	mutation *BookMutation
}

// Where appends a list predicates to the BookUpdate builder.
func (bu *BookUpdate) Where(ps ...predicate.Book) *BookUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetTitle sets the "title" field.
func (bu *BookUpdate) SetTitle(s string) *BookUpdate {
	bu.mutation.SetTitle(s)
	return bu
}

// SetAuthor sets the "author" field.
func (bu *BookUpdate) SetAuthor(s string) *BookUpdate {
	bu.mutation.SetAuthor(s)
	return bu
}

// Mutation returns the BookMutation object of the builder.
func (bu *BookUpdate) Mutation() *BookMutation {
	return bu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BookUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BookUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BookUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BookUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bu *BookUpdate) check() error {
	if v, ok := bu.mutation.Title(); ok {
		if err := book.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Book.title": %w`, err)}
		}
	}
	if v, ok := bu.mutation.Author(); ok {
		if err := book.AuthorValidator(v); err != nil {
			return &ValidationError{Name: "author", err: fmt.Errorf(`ent: validator failed for field "Book.author": %w`, err)}
		}
	}
	return nil
}

func (bu *BookUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := bu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(book.Table, book.Columns, sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.Title(); ok {
		_spec.SetField(book.FieldTitle, field.TypeString, value)
	}
	if value, ok := bu.mutation.Author(); ok {
		_spec.SetField(book.FieldAuthor, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{book.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BookUpdateOne is the builder for updating a single Book entity.
type BookUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BookMutation
}

// SetTitle sets the "title" field.
func (buo *BookUpdateOne) SetTitle(s string) *BookUpdateOne {
	buo.mutation.SetTitle(s)
	return buo
}

// SetAuthor sets the "author" field.
func (buo *BookUpdateOne) SetAuthor(s string) *BookUpdateOne {
	buo.mutation.SetAuthor(s)
	return buo
}

// Mutation returns the BookMutation object of the builder.
func (buo *BookUpdateOne) Mutation() *BookMutation {
	return buo.mutation
}

// Where appends a list predicates to the BookUpdate builder.
func (buo *BookUpdateOne) Where(ps ...predicate.Book) *BookUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BookUpdateOne) Select(field string, fields ...string) *BookUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Book entity.
func (buo *BookUpdateOne) Save(ctx context.Context) (*Book, error) {
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BookUpdateOne) SaveX(ctx context.Context) *Book {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BookUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BookUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (buo *BookUpdateOne) check() error {
	if v, ok := buo.mutation.Title(); ok {
		if err := book.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Book.title": %w`, err)}
		}
	}
	if v, ok := buo.mutation.Author(); ok {
		if err := book.AuthorValidator(v); err != nil {
			return &ValidationError{Name: "author", err: fmt.Errorf(`ent: validator failed for field "Book.author": %w`, err)}
		}
	}
	return nil
}

func (buo *BookUpdateOne) sqlSave(ctx context.Context) (_node *Book, err error) {
	if err := buo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(book.Table, book.Columns, sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Book.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, book.FieldID)
		for _, f := range fields {
			if !book.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != book.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.Title(); ok {
		_spec.SetField(book.FieldTitle, field.TypeString, value)
	}
	if value, ok := buo.mutation.Author(); ok {
		_spec.SetField(book.FieldAuthor, field.TypeString, value)
	}
	_node = &Book{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{book.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
