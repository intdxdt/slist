package slist

import (
	"fmt"
	"bytes"
	"github.com/intdxdt/cmp"
	"github.com/intdxdt/skiplist"
)

//SList - Sorted Set
type SList struct {
	list *skiplist.SkipList
}

//NewSList - New Sorted List Constructor
func NewSList(size int, comparator cmp.Compare) *SList {
	var list = skiplist.NewSkipList(size, true, comparator)
	return &SList{
		list: list,
	}
}

//Push item to list.
func (sl *SList) Add(item interface{}) *SList {
	sl.list.Insert(item)
	return sl
}

//Contains - checks for the presence of a value in the sorted list.
func (sl *SList) Contains(item interface{}) bool {
	v := sl.list.Search(item)
	return v != nil
}

//ContainsAll - checks if all items are in list.
func (sl *SList) ContainsAll(items []interface{}) bool {
	var bln = true
	var n = len(items)
	for i := 0; bln && i < n; i++ {
		bln = sl.Contains(items[i])
	}
	return bln
}

//Stringer method
func (sl *SList) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	sl.list.Each(func(o interface{}, i int) {
		buffer.WriteString(fmt.Sprintf("%v", o) + ",")
	})
	buffer.WriteString("]")
	return buffer.String()
}

//Removes an item by value from the sorted list.
func (sl *SList) Remove(value interface{}) *SList {
	sl.list.Remove(value)
	return sl
}

//IsEmpty checks if sorted list is empty ?
func (sl *SList) IsEmpty() bool {
	return sl.list.IsEmpty()
}

//First item in sorted list.
func (sl *SList) First() interface{} {
	return sl.list.First()
}

//Last interface{} in list
func (sl *SList) Last() interface{} {
	return sl.list.Last()
}

//Each - iterates over each item with call to func(interface{}, int).
func (sl *SList) Each(fn func(interface{}, int)) {
	sl.list.Each(fn)
}

//Filters items based on predicate : func (item interface{}, i int) bool
func (sl *SList) Filter(fn func(item interface{}) bool) []interface{} {
	var items = make([]interface{}, 0)
	sl.Each(func(v interface{}, i int) {
		if fn(v) {
			items = append(items, v)
		}
	})
	return items
}

//Pops item from the end of the sorted list.
func (sl *SList) Pop() interface{} {
	if !sl.IsEmpty() {
		v := sl.list.Last()
		sl.list.Remove(v)
		return v
	}
	return nil
}

//PopLeft - removes item from the beginning of the sorted list.
func (sl *SList) Shift() interface{} {
	if !sl.IsEmpty() {
		v := sl.list.First()
		sl.list.Remove(v)
		return v
	}
	return nil
}

//Empty empties a sorted list.
func (sl *SList) Empty() *SList {
	sl.list.Empty()
	return sl
}
