package slist

import (
	"fmt"
	"testing"
	"github.com/intdxdt/cmp"
	"github.com/franela/goblin"
)

func TestSList(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("SSet", func() {
		var array = []int{9, 5, 3, 2, 8, 6, 4, 6, 1, 2, 3};
		var obj = NewSList(len(array), cmp.Int)

		for _, v := range array {
			obj.Add(v)
		}
		//print
		fmt.Println(obj)

		g.It("should test Delete & Search & Iter item from list", func() {
			var f6 = 6
			var f3 = 3
			var search_items = []interface{}{f6, f3}
			g.Assert(obj.Contains(f6)).IsTrue()
			g.Assert(obj.ContainsAll(search_items)).IsTrue()
			g.Assert(obj.Remove(f6).Contains(f6)).IsTrue()
			g.Assert(obj.Remove(f6).Contains(f6)).IsFalse()
			g.Assert(obj.IsEmpty()).IsFalse()

			var each_item = make([]int, 0)
			obj.Each(func(o interface{}, i int) {
				each_item = append(each_item, o.(int))
			})
			g.Assert(each_item).Eql([]int{1, 2, 2, 3, 3, 4, 5, 8, 9,})
			var odd_list = obj.Filter(func(o interface{}) bool {
				return o.(int)%2 == 1
			})
			g.Assert(len(odd_list)).Eql(5)
			var odds = make([]int, 0)
			for _, v := range odd_list {
				odds = append(odds, v.(int))
			}
			g.Assert(odds).Eql([]int{1, 3, 3, 5, 9})
			g.Assert(obj.First()).Equal(int(1))
			g.Assert(obj.Last()).Equal(int(9))
			g.Assert(obj.Shift()).Equal(int(1))
			g.Assert(obj.Pop()).Equal(int(9))

			g.Assert(obj.First()).Equal(int(2))
			g.Assert(obj.Last()).Equal(int(8))
			obj.Empty()
			g.Assert(obj.Shift()).Equal(nil)
			g.Assert(obj.Pop()).Equal(nil)
			g.Assert(obj.IsEmpty()).IsTrue()

			//print
			fmt.Println(obj)
		})
	})
}
