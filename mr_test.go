package mr

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type tPerson struct {
	ID   int
	Name string
}

var (
	test_persons = []tPerson{
		{1, "tom"},
		{2, "jerry"},
		{3, "lili"},
	}

	test_persons2 = []*tPerson{
		{1, "tom"},
		{2, "jerry"},
		{3, "lili"},
	}

	test_int_arr    = []int{1, 2, 3, 4, 5}
	test_string_arr = []string{"a", "b", "c"}
)

func TestMap(t *testing.T) {
	f := func(i int) int {
		return i * 2
	}
	newm := Map(test_int_arr, f)
	for i, v := range test_int_arr {
		assert.Equal(t, newm[i], f(v), "Map failed")
	}

	f2 := func(s string) string {
		return s + "1"
	}
	newm2 := Map(test_string_arr, f2)
	for i, v := range test_string_arr {
		assert.Equal(t, newm2[i], f2(v), "Map failed")
	}

	newm3 := Map(test_persons, func(s tPerson) string {
		return s.Name
	})
	assert.Equal(t, newm3[0], "tom", "Map failed")
	assert.Equal(t, newm3[1], "jerry", "Map failed")

	newm4 := Map(test_persons2, func(s *tPerson) string {
		return s.Name
	})
	assert.Equal(t, newm4[0], "tom", "Map failed")
	assert.Equal(t, newm4[1], "jerry", "Map failed")

	ts5 := []*tPerson{}
	newm5 := Map(ts5, func(s *tPerson) string {
		return s.Name
	})
	assert.Emptyf(t, newm5, "is not empty: %v", newm5)
}

func TestReduce(t *testing.T) {
	r1 := Reduce(test_int_arr, func(a, b int) int {
		return a + b
	}, 0)
	assert.Equal(t, r1, 15, "Reduce failed: %v", r1)

	r2 := Reduce(test_string_arr, func(a, b string) string {
		return a + b
	}, "")
	assert.Equal(t, r2, "abc", "Reduce failed: %v", r2)

	r3 := Reduce(test_persons, func(a, b tPerson) tPerson {
		if b.ID > a.ID {
			return b
		}
		return a
	}, tPerson{})
	assert.Equal(t, r3.Name, "lili", "Reduce failed: %v", r3)
}

func TestFilter(t *testing.T) {
	r1 := Filter(test_int_arr, func(i int) bool {
		return i > 2
	})
	assert.Lenf(t, r1, 3, "len: %v", r1)

	r2 := Filter(test_persons2, func(p *tPerson) bool {
		return p.Name == "tom"
	})

	assert.Lenf(t, r2, 1, "len: %v", r2)
}

func TestToMap(t *testing.T) {
	r1 := ToMap(test_int_arr, func(i int) int {
		return i
	})
	for k, v := range r1 {
		assert.Equalf(t, k, v, "ToMap failed: %v", v)
	}

	r2 := ToMap(test_persons2, func(p *tPerson) int {
		return p.ID
	})
	assert.Equalf(t, r2[1].Name, "tom", "ToMap failed: %v", r2)
}

func TestUnique(t *testing.T) {
	a := []int{1, 2, 1, 3, 4, 5, 1, 2, 3, 4}
	b := Unique(a, func(i int) int { return i })
	assert.Lenf(t, b, 5, "len: %v", b)
	assert.Equalf(t, b, []int{1, 2, 3, 4, 5}, "%v", b)

	test_person3 := test_persons2
	test_person3 = append(test_person3, &tPerson{1, "tom"})
	b2 := Unique(test_person3, func(i *tPerson) int { return i.ID })
	assert.Lenf(t, b2, 3, "len: %v", b2)
}

func TestDiff(t *testing.T) {
	// Test case 1: Empty slices
	ts1 := []int{}
	ts2 := []int{}
	f := func(i int) int { return i }
	expected := []int{}
	result := Diff(ts1, ts2, f)
	assert.Equal(t, expected, result)

	// Test case 2: No common elements
	ts1 = []int{1, 2, 3}
	ts2 = []int{4, 5, 6}
	f = func(i int) int { return i }
	expected = []int{1, 2, 3}
	result = Diff(ts1, ts2, f)
	assert.Equal(t, expected, result)

	// Test case 3: Some common elements
	ts1 = []int{1, 2, 3, 4, 5}
	ts2 = []int{4, 5, 6, 7, 8}
	f = func(i int) int { return i }
	expected = []int{1, 2, 3}
	result = Diff(ts1, ts2, f)
	assert.Equal(t, expected, result)

	// Test case 4: All elements in ts1 present in ts2
	ts1 = []int{1, 2, 3, 4, 5}
	ts2 = []int{1, 2, 3, 4, 5}
	f = func(i int) int { return i }
	expected = []int{}
	result = Diff(ts1, ts2, f)
	assert.Equal(t, expected, result)
}

func TestIntersect(t *testing.T) {
	// Testing for intersecting elements in the two slices
	t.Run("Intersecting elements", func(t *testing.T) {
		ts1 := []int{1, 2, 3, 4, 5}
		ts2 := []int{4, 5, 6, 7, 8}
		expected := []int{4, 5}
		result := Intersect(ts1, ts2, func(i int) int { return i })
		assert.ElementsMatch(t, expected, result)
	})

	// Testing for no intersecting elements in the two slices
	t.Run("No intersecting elements", func(t *testing.T) {
		ts1 := []int{1, 2, 3}
		ts2 := []int{4, 5, 6}
		expected := []int{}
		result := Intersect(ts1, ts2, func(i int) int { return i })
		assert.ElementsMatch(t, expected, result)
	})

	// Testing for an empty slice as input
	t.Run("Empty slice as input", func(t *testing.T) {
		ts1 := []int{}
		ts2 := []int{4, 5, 6}
		expected := []int{}
		result := Intersect(ts1, ts2, func(i int) int { return i })
		assert.ElementsMatch(t, expected, result)
	})
}

func TestMerge(t *testing.T) {
	// Test case 1: Merge two empty slices
	ts1 := []int{}
	ts2 := []int{}
	expected := []int{}
	result := Merge(ts1, ts2, func(i int) int { return i })
	assert.Equal(t, expected, result)

	// Test case 2: Merge two slices with common elements
	ts1 = []int{1, 2, 3}
	ts2 = []int{3, 4, 5}
	expected = []int{1, 2, 3, 4, 5}
	result = Merge(ts1, ts2, func(i int) int { return i })
	assert.Equal(t, expected, result)

	// Test case 3: Merge two slices with no common elements
	ts1 = []int{1, 2, 3}
	ts2 = []int{4, 5, 6}
	expected = []int{1, 2, 3, 4, 5, 6}
	result = Merge(ts1, ts2, func(i int) int { return i })
	assert.Equal(t, expected, result)
}

func TestContains(t *testing.T) {
	// Test case 1: Empty slice, should return false
	ts := []int{}
	tVal := 5
	f := func(t int) int { return t }
	assert.False(t, Contains(ts, tVal, f), "Expected false, but got true")

	// Test case 2: Slice with no matching element, should return false
	ts = []int{1, 2, 3, 4}
	tVal = 5
	assert.False(t, Contains(ts, tVal, f), "Expected false, but got true")

	// Test case 3: Slice with a matching element, should return true
	ts = []int{1, 2, 3, 4, 5}
	tVal = 5
	assert.True(t, Contains(ts, tVal, f), "Expected true, but got false")

	// Test case 4: Slice with multiple matching elements, should return true
	ts = []int{1, 2, 3, 4, 5, 5, 5}
	tVal = 5
	assert.True(t, Contains(ts, tVal, f), "Expected true, but got false")
}

func TestPaginate(t *testing.T) {
	// Test case 1: Empty slice
	ts := []int{}
	expected := []int{}
	result := Paginate(ts, 1, 1)
	assert.Equal(t, expected, result)

	// Test case 2: Slice with 1 element
	ts = []int{1}
	expected = []int{}
	result = Paginate(ts, 2, 1)
	assert.Equal(t, expected, result)

	// Test case 3: Slice with 2 elements
	ts = []int{1, 2}
	expected = []int{2}
	result = Paginate(ts, 2, 1)
	assert.Equal(t, expected, result)

	// Test case 4: Slice with 3 elements
	ts = []int{1, 2, 3}
	expected = []int{1, 2}
	result = Paginate(ts, 1, 2)
	assert.Equal(t, expected, result)

	// Test case 5: Slice with 4 elements
	ts = []int{1, 2, 3, 4}
	expected = []int{3, 4}
	result = Paginate(ts, 2, 2)
	assert.Equal(t, expected, result)
}

func TestJoin(t *testing.T) {
	str := Join(test_int_arr, ",", func(i int) string { return strconv.Itoa(i) })
	assert.Equal(t, "1,2,3,4,5", str)

	ts := []int{}
	str = Join(ts, ",", func(i int) string { return strconv.Itoa(i) })
	assert.Empty(t, str)
}

func TestKeys(t *testing.T) {
	td1 := map[string]int64{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	actual := Keys(td1)
	expected := []string{"a", "b", "c"}
	assert.EqualValuesf(t, expected, actual, "Expected %v atutal %v", expected, actual)
}
