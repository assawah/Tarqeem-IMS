// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tarqeem/ims/ent/file"
	"github.com/tarqeem/ims/ent/issue"
	"github.com/tarqeem/ims/ent/predicate"
)

// FileUpdate is the builder for updating File entities.
type FileUpdate struct {
	config
	hooks    []Hook
	mutation *FileMutation
}

// Where appends a list predicates to the FileUpdate builder.
func (fu *FileUpdate) Where(ps ...predicate.File) *FileUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetFilePath sets the "file_path" field.
func (fu *FileUpdate) SetFilePath(s string) *FileUpdate {
	fu.mutation.SetFilePath(s)
	return fu
}

// SetNillableFilePath sets the "file_path" field if the given value is not nil.
func (fu *FileUpdate) SetNillableFilePath(s *string) *FileUpdate {
	if s != nil {
		fu.SetFilePath(*s)
	}
	return fu
}

// ClearFilePath clears the value of the "file_path" field.
func (fu *FileUpdate) ClearFilePath() *FileUpdate {
	fu.mutation.ClearFilePath()
	return fu
}

// SetFileName sets the "file_name" field.
func (fu *FileUpdate) SetFileName(s string) *FileUpdate {
	fu.mutation.SetFileName(s)
	return fu
}

// SetNillableFileName sets the "file_name" field if the given value is not nil.
func (fu *FileUpdate) SetNillableFileName(s *string) *FileUpdate {
	if s != nil {
		fu.SetFileName(*s)
	}
	return fu
}

// ClearFileName clears the value of the "file_name" field.
func (fu *FileUpdate) ClearFileName() *FileUpdate {
	fu.mutation.ClearFileName()
	return fu
}

// SetFileSize sets the "file_size" field.
func (fu *FileUpdate) SetFileSize(i int64) *FileUpdate {
	fu.mutation.ResetFileSize()
	fu.mutation.SetFileSize(i)
	return fu
}

// SetNillableFileSize sets the "file_size" field if the given value is not nil.
func (fu *FileUpdate) SetNillableFileSize(i *int64) *FileUpdate {
	if i != nil {
		fu.SetFileSize(*i)
	}
	return fu
}

// AddFileSize adds i to the "file_size" field.
func (fu *FileUpdate) AddFileSize(i int64) *FileUpdate {
	fu.mutation.AddFileSize(i)
	return fu
}

// ClearFileSize clears the value of the "file_size" field.
func (fu *FileUpdate) ClearFileSize() *FileUpdate {
	fu.mutation.ClearFileSize()
	return fu
}

// AddIssueIDs adds the "issue" edge to the Issue entity by IDs.
func (fu *FileUpdate) AddIssueIDs(ids ...int) *FileUpdate {
	fu.mutation.AddIssueIDs(ids...)
	return fu
}

// AddIssue adds the "issue" edges to the Issue entity.
func (fu *FileUpdate) AddIssue(i ...*Issue) *FileUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return fu.AddIssueIDs(ids...)
}

// Mutation returns the FileMutation object of the builder.
func (fu *FileUpdate) Mutation() *FileMutation {
	return fu.mutation
}

// ClearIssue clears all "issue" edges to the Issue entity.
func (fu *FileUpdate) ClearIssue() *FileUpdate {
	fu.mutation.ClearIssue()
	return fu
}

// RemoveIssueIDs removes the "issue" edge to Issue entities by IDs.
func (fu *FileUpdate) RemoveIssueIDs(ids ...int) *FileUpdate {
	fu.mutation.RemoveIssueIDs(ids...)
	return fu
}

// RemoveIssue removes "issue" edges to Issue entities.
func (fu *FileUpdate) RemoveIssue(i ...*Issue) *FileUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return fu.RemoveIssueIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FileUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, fu.sqlSave, fu.mutation, fu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FileUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FileUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FileUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fu *FileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(file.Table, file.Columns, sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt))
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.FilePath(); ok {
		_spec.SetField(file.FieldFilePath, field.TypeString, value)
	}
	if fu.mutation.FilePathCleared() {
		_spec.ClearField(file.FieldFilePath, field.TypeString)
	}
	if value, ok := fu.mutation.FileName(); ok {
		_spec.SetField(file.FieldFileName, field.TypeString, value)
	}
	if fu.mutation.FileNameCleared() {
		_spec.ClearField(file.FieldFileName, field.TypeString)
	}
	if value, ok := fu.mutation.FileSize(); ok {
		_spec.SetField(file.FieldFileSize, field.TypeInt64, value)
	}
	if value, ok := fu.mutation.AddedFileSize(); ok {
		_spec.AddField(file.FieldFileSize, field.TypeInt64, value)
	}
	if fu.mutation.FileSizeCleared() {
		_spec.ClearField(file.FieldFileSize, field.TypeInt64)
	}
	if fu.mutation.IssueCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   file.IssueTable,
			Columns: file.IssuePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(issue.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.RemovedIssueIDs(); len(nodes) > 0 && !fu.mutation.IssueCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   file.IssueTable,
			Columns: file.IssuePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(issue.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.IssueIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   file.IssueTable,
			Columns: file.IssuePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(issue.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{file.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fu.mutation.done = true
	return n, nil
}

// FileUpdateOne is the builder for updating a single File entity.
type FileUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FileMutation
}

// SetFilePath sets the "file_path" field.
func (fuo *FileUpdateOne) SetFilePath(s string) *FileUpdateOne {
	fuo.mutation.SetFilePath(s)
	return fuo
}

// SetNillableFilePath sets the "file_path" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableFilePath(s *string) *FileUpdateOne {
	if s != nil {
		fuo.SetFilePath(*s)
	}
	return fuo
}

// ClearFilePath clears the value of the "file_path" field.
func (fuo *FileUpdateOne) ClearFilePath() *FileUpdateOne {
	fuo.mutation.ClearFilePath()
	return fuo
}

// SetFileName sets the "file_name" field.
func (fuo *FileUpdateOne) SetFileName(s string) *FileUpdateOne {
	fuo.mutation.SetFileName(s)
	return fuo
}

// SetNillableFileName sets the "file_name" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableFileName(s *string) *FileUpdateOne {
	if s != nil {
		fuo.SetFileName(*s)
	}
	return fuo
}

// ClearFileName clears the value of the "file_name" field.
func (fuo *FileUpdateOne) ClearFileName() *FileUpdateOne {
	fuo.mutation.ClearFileName()
	return fuo
}

// SetFileSize sets the "file_size" field.
func (fuo *FileUpdateOne) SetFileSize(i int64) *FileUpdateOne {
	fuo.mutation.ResetFileSize()
	fuo.mutation.SetFileSize(i)
	return fuo
}

// SetNillableFileSize sets the "file_size" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableFileSize(i *int64) *FileUpdateOne {
	if i != nil {
		fuo.SetFileSize(*i)
	}
	return fuo
}

// AddFileSize adds i to the "file_size" field.
func (fuo *FileUpdateOne) AddFileSize(i int64) *FileUpdateOne {
	fuo.mutation.AddFileSize(i)
	return fuo
}

// ClearFileSize clears the value of the "file_size" field.
func (fuo *FileUpdateOne) ClearFileSize() *FileUpdateOne {
	fuo.mutation.ClearFileSize()
	return fuo
}

// AddIssueIDs adds the "issue" edge to the Issue entity by IDs.
func (fuo *FileUpdateOne) AddIssueIDs(ids ...int) *FileUpdateOne {
	fuo.mutation.AddIssueIDs(ids...)
	return fuo
}

// AddIssue adds the "issue" edges to the Issue entity.
func (fuo *FileUpdateOne) AddIssue(i ...*Issue) *FileUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return fuo.AddIssueIDs(ids...)
}

// Mutation returns the FileMutation object of the builder.
func (fuo *FileUpdateOne) Mutation() *FileMutation {
	return fuo.mutation
}

// ClearIssue clears all "issue" edges to the Issue entity.
func (fuo *FileUpdateOne) ClearIssue() *FileUpdateOne {
	fuo.mutation.ClearIssue()
	return fuo
}

// RemoveIssueIDs removes the "issue" edge to Issue entities by IDs.
func (fuo *FileUpdateOne) RemoveIssueIDs(ids ...int) *FileUpdateOne {
	fuo.mutation.RemoveIssueIDs(ids...)
	return fuo
}

// RemoveIssue removes "issue" edges to Issue entities.
func (fuo *FileUpdateOne) RemoveIssue(i ...*Issue) *FileUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return fuo.RemoveIssueIDs(ids...)
}

// Where appends a list predicates to the FileUpdate builder.
func (fuo *FileUpdateOne) Where(ps ...predicate.File) *FileUpdateOne {
	fuo.mutation.Where(ps...)
	return fuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FileUpdateOne) Select(field string, fields ...string) *FileUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated File entity.
func (fuo *FileUpdateOne) Save(ctx context.Context) (*File, error) {
	return withHooks(ctx, fuo.sqlSave, fuo.mutation, fuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FileUpdateOne) SaveX(ctx context.Context) *File {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FileUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FileUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fuo *FileUpdateOne) sqlSave(ctx context.Context) (_node *File, err error) {
	_spec := sqlgraph.NewUpdateSpec(file.Table, file.Columns, sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt))
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "File.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, file.FieldID)
		for _, f := range fields {
			if !file.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != file.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fuo.mutation.FilePath(); ok {
		_spec.SetField(file.FieldFilePath, field.TypeString, value)
	}
	if fuo.mutation.FilePathCleared() {
		_spec.ClearField(file.FieldFilePath, field.TypeString)
	}
	if value, ok := fuo.mutation.FileName(); ok {
		_spec.SetField(file.FieldFileName, field.TypeString, value)
	}
	if fuo.mutation.FileNameCleared() {
		_spec.ClearField(file.FieldFileName, field.TypeString)
	}
	if value, ok := fuo.mutation.FileSize(); ok {
		_spec.SetField(file.FieldFileSize, field.TypeInt64, value)
	}
	if value, ok := fuo.mutation.AddedFileSize(); ok {
		_spec.AddField(file.FieldFileSize, field.TypeInt64, value)
	}
	if fuo.mutation.FileSizeCleared() {
		_spec.ClearField(file.FieldFileSize, field.TypeInt64)
	}
	if fuo.mutation.IssueCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   file.IssueTable,
			Columns: file.IssuePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(issue.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.RemovedIssueIDs(); len(nodes) > 0 && !fuo.mutation.IssueCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   file.IssueTable,
			Columns: file.IssuePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(issue.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.IssueIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   file.IssueTable,
			Columns: file.IssuePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(issue.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &File{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{file.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fuo.mutation.done = true
	return _node, nil
}
