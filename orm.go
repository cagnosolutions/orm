package orm

import (
    "reflect"
)

// main orm interface
type Orm interface {
    Find(id []byte, dest interface{}) error
    FindAll(dests ...interface{}) (int, error)
    FindAllLimit(limit int, dest ...interface{}) (int, error)
    Delete(id []byte) error
    Insert(data ...interface{}) (int, error)
    Update(id []byte, updated interface{}) error
    QueryModel(query string, args ...interface{}) ([]*Model, int)
}

type Model struct {
    Typ reflect.Type
    Data interface{}
}

func NewModel(v interface{}) *Model {
    m := &Model{
        // fill out new model
        // return pointer to it
    }
    return m
}

func (m *Model) Get(dest interface{}) {
    // build / unmarshal / fill-out
}
