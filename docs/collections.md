<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# collections

```go
import "github.com/avivatedgi/go-rust-std/collections"
```

## Index

- [func Dedup[T comparable](vec *Vec[T])](<#func-dedup>)
- [func DedupByKey[T comparable](vec *Vec[T], key func(T) T)](<#func-dedupbykey>)
- [type Iterator](<#type-iterator>)
  - [func (ch *Iterator[T]) Close()](<#func-iteratort-close>)
  - [func (ch *Iterator[T]) IntoVector() *Vec[T]](<#func-iteratort-intovector>)
  - [func (ch *Iterator[T]) Push(value *T)](<#func-iteratort-push>)
- [type Map](<#type-map>)
  - [func (m *Map[K, V]) Clear()](<#func-mapk-v-clear>)
  - [func (m Map[K, V]) ContainsKey(key K) bool](<#func-mapk-v-containskey>)
  - [func (m *Map[K, V]) Drain() Iterator[Pair[K, V]]](<#func-mapk-v-drain>)
  - [func (m *Map[K, V]) Entry(key K) MapEntry[K, V]](<#func-mapk-v-entry>)
  - [func (m *Map[K, V]) ForEach(f func(*K, *V) bool)](<#func-mapk-v-foreach>)
  - [func (m Map[K, V]) Get(key K) option.Option[V]](<#func-mapk-v-get>)
  - [func (m Map[K, V]) GetKeyValue(key K) option.Option[Pair[K, V]]](<#func-mapk-v-getkeyvalue>)
  - [func (m *Map[K, V]) Insert(key K, value V) option.Option[V]](<#func-mapk-v-insert>)
  - [func (m Map[K, V]) Iter() Iterator[Pair[K, V]]](<#func-mapk-v-iter>)
  - [func (m Map[K, _]) Keys() Iterator[K]](<#func-mapk-_-keys>)
  - [func (m Map[_, V]) Values() Iterator[V]](<#func-map_-v-values>)
- [type MapEntry](<#type-mapentry>)
  - [func (m MapEntry[K, V]) AndModify(f func(*V)) MapEntry[K, V]](<#func-mapentryk-v-andmodify>)
  - [func (m MapEntry[K, V]) Key() K](<#func-mapentryk-v-key>)
  - [func (m MapEntry[K, V]) OrDefault() V](<#func-mapentryk-v-ordefault>)
  - [func (m MapEntry[K, V]) OrInsert(value V) V](<#func-mapentryk-v-orinsert>)
  - [func (m MapEntry[K, V]) OrInsertWith(f func() V) V](<#func-mapentryk-v-orinsertwith>)
  - [func (m MapEntry[K, V]) OrInsertWithKey(f func(K) V) V](<#func-mapentryk-v-orinsertwithkey>)
- [type Pair](<#type-pair>)
- [type Vec](<#type-vec>)
  - [func (vec *Vec[T]) Append(other *Vec[T])](<#func-vect-append>)
  - [func (vec Vec[T]) Capacity() int](<#func-vect-capacity>)
  - [func (vec *Vec[T]) Clear()](<#func-vect-clear>)
  - [func (vec *Vec[T]) DedupBy(f func(T, T) bool)](<#func-vect-dedupby>)
  - [func (vec *Vec[T]) Drain(start, end int) Iterator[T]](<#func-vect-drain>)
  - [func (vec *Vec[T]) Extend(other *Vec[T])](<#func-vect-extend>)
  - [func (vec *Vec[T]) Insert(index int, item T)](<#func-vect-insert>)
  - [func (vec Vec[T]) IsEmpty() bool](<#func-vect-isempty>)
  - [func (vec *Vec[T]) Iter() Iterator[T]](<#func-vect-iter>)
  - [func (vec Vec[T]) Len() int](<#func-vect-len>)
  - [func (vec *Vec[T]) Pop() option.Option[T]](<#func-vect-pop>)
  - [func (vec *Vec[T]) Push(item T)](<#func-vect-push>)
  - [func (vec *Vec[T]) Remove(index int) T](<#func-vect-remove>)
  - [func (vec *Vec[T]) Resize(newLength int, value T)](<#func-vect-resize>)
  - [func (vec *Vec[T]) ResizeWith(newLength int, f func() T)](<#func-vect-resizewith>)
  - [func (vec *Vec[T]) Retain(f func(T) bool)](<#func-vect-retain>)
  - [func (vec *Vec[T]) Splice(start, end int, replaceWith Vec[T]) Vec[T]](<#func-vect-splice>)
  - [func (vec Vec[T]) SplitOff(at int) Vec[T]](<#func-vect-splitoff>)
  - [func (vec *Vec[T]) SwapRemove(index int) T](<#func-vect-swapremove>)
  - [func (vec *Vec[T]) Truncate(len int)](<#func-vect-truncate>)


## func Dedup

```go
func Dedup[T comparable](vec *Vec[T])
```

Removes consecutive repeated elements in the vector\. If the vector is sorted\, this removes all duplicates\.

NOTE: This function isn't a method of the vector because it can only work on comparable types\, and I didn't wanted to limit the Vector structure to hold only comparable types\.

## func DedupByKey

```go
func DedupByKey[T comparable](vec *Vec[T], key func(T) T)
```

Removes all but the first of consecutive elements in the vector that resolve to the same key\. If the vector is sorted\, this removes all duplicates\.

NOTE: This function isn't a method of the vector because it can only work on comparable types\, and I didn't wanted to limit the Vector structure to hold only comparable types\.

## type Iterator

A channel based iterator\, used to abstarct a channel as an iterator\. Usage example \(taken from collections\.Vector\[T\]\): for value := range vec\.Iter\(\) \{ fmt\.Println\(value\) \}

```go
type Iterator[T any] chan T
```

### func \(\*Iterator\[T\]\) Close

```go
func (ch *Iterator[T]) Close()
```

Close the channel\, this function MUST be called after we are done use Push for all of the values we want to iterate\.

### func \(\*Iterator\[T\]\) IntoVector

```go
func (ch *Iterator[T]) IntoVector() *Vec[T]
```

Convert an iterator into a vector of the same type\.

### func \(\*Iterator\[T\]\) Push

```go
func (ch *Iterator[T]) Push(value *T)
```

Push a value into the channel\, passing nil as the value will be ignored\.

## type Map

```go
type Map[K comparable, V any] map[K]V
```

### func \(\*Map\[K\, V\]\) Clear

```go
func (m *Map[K, V]) Clear()
```

Clears the map\, removing all key\-value pairs\.

### func \(Map\[K\, V\]\) ContainsKey

```go
func (m Map[K, V]) ContainsKey(key K) bool
```

Returns true if the map contains a value for the specified key\.

### func \(\*Map\[K\, V\]\) Drain

```go
func (m *Map[K, V]) Drain() Iterator[Pair[K, V]]
```

Clears the map\, returning all key\-value pairs as an iterator\.

### func \(\*Map\[K\, V\]\) Entry

```go
func (m *Map[K, V]) Entry(key K) MapEntry[K, V]
```

Gets the given key’s corresponding entry in the map for in\-place manipulation\. WARNING: In difference from rust\, this method does not return a reference to the value \(\!\)\. But\, it does match the signatures of MapEntry in rust\.

### func \(\*Map\[K\, V\]\) ForEach

```go
func (m *Map[K, V]) ForEach(f func(*K, *V) bool)
```

Executes the f function once for each map entry\. There is an option to stop the iteration in the middle\, if the handler function returns false\.

### func \(Map\[K\, V\]\) Get

```go
func (m Map[K, V]) Get(key K) option.Option[V]
```

Returns the value corresponding to the key\.

### func \(Map\[K\, V\]\) GetKeyValue

```go
func (m Map[K, V]) GetKeyValue(key K) option.Option[Pair[K, V]]
```

Returns the key\-value pair corresponding to the supplied key in a Pair \(first is key\, second is value\)\.

### func \(\*Map\[K\, V\]\) Insert

```go
func (m *Map[K, V]) Insert(key K, value V) option.Option[V]
```

Inserts a key\-value pair into the map\. If the map did not have this key present\, option\.None is returned\. If the map did have this key present\, the value is updated\, and the old value is returned\. The key is not updated\, though;

### func \(Map\[K\, V\]\) Iter

```go
func (m Map[K, V]) Iter() Iterator[Pair[K, V]]
```

Return the map iterator\.

### func \(Map\[K\, \_\]\) Keys

```go
func (m Map[K, _]) Keys() Iterator[K]
```

Retreive all the keys of the map\.

### func \(Map\[\_\, V\]\) Values

```go
func (m Map[_, V]) Values() Iterator[V]
```

Retreive all the values of the map

## type MapEntry

This struct is constructed from the Entry method on Map\.

```go
type MapEntry[K comparable, V any] struct {
    // contains filtered or unexported fields
}
```

### func \(MapEntry\[K\, V\]\) AndModify

```go
func (m MapEntry[K, V]) AndModify(f func(*V)) MapEntry[K, V]
```

Provides access to an occupied entry before any potential inserts into the map\.

### func \(MapEntry\[K\, V\]\) Key

```go
func (m MapEntry[K, V]) Key() K
```

Returns this entry’s key\.

### func \(MapEntry\[K\, V\]\) OrDefault

```go
func (m MapEntry[K, V]) OrDefault() V
```

Ensures a value is in the entry by inserting the default value if empty\. Returns the entry's value\.

### func \(MapEntry\[K\, V\]\) OrInsert

```go
func (m MapEntry[K, V]) OrInsert(value V) V
```

Ensures a value is in the entry by inserting the default if empty\. Returns the entry’s value\.

### func \(MapEntry\[K\, V\]\) OrInsertWith

```go
func (m MapEntry[K, V]) OrInsertWith(f func() V) V
```

Ensures a value is in the entry by inserting the result of the default function if empty\. Returns the entry’s value\.

### func \(MapEntry\[K\, V\]\) OrInsertWithKey

```go
func (m MapEntry[K, V]) OrInsertWithKey(f func(K) V) V
```

Ensures a value is in the entry by inserting\, if empty\, the result of the default function\.

## type Pair

A struct that represents a pair of values\.

```go
type Pair[T any, U any] struct {
    First  T
    Second U
}
```

## type Vec

```go
type Vec[T any] []T
```

### func \(\*Vec\[T\]\) Append

```go
func (vec *Vec[T]) Append(other *Vec[T])
```

Moves all the elements of other into vec\, leaving other empty\.

### func \(Vec\[T\]\) Capacity

```go
func (vec Vec[T]) Capacity() int
```

Returns the number of elements the vector can hold without reallocating\.

### func \(\*Vec\[T\]\) Clear

```go
func (vec *Vec[T]) Clear()
```

Clears the vector\, removing all values\. Note that this method has no effect on the allocated capacity of the vector\.

### func \(\*Vec\[T\]\) DedupBy

```go
func (vec *Vec[T]) DedupBy(f func(T, T) bool)
```

Removes all but the first of consecutive elements in the vector satisfying a given equality relation\. The f function is passed the two elements from the vector and must determine if the elements compare equal\. The elements are passed in opposite order from their order in the slice\, so if f\(a\, b\) returns true\, a is removed\.

### func \(\*Vec\[T\]\) Drain

```go
func (vec *Vec[T]) Drain(start, end int) Iterator[T]
```

Removes the specified range from the vector in bulk\, returning all removed elements as an Iterator\.

### func \(\*Vec\[T\]\) Extend

```go
func (vec *Vec[T]) Extend(other *Vec[T])
```

Appends all elements in a slice to the Vec\.

### func \(\*Vec\[T\]\) Insert

```go
func (vec *Vec[T]) Insert(index int, item T)
```

Inserts an element at position index within the vector\, shifting all elements after it to the right\. Panics if index \> len\.

### func \(Vec\[T\]\) IsEmpty

```go
func (vec Vec[T]) IsEmpty() bool
```

Returns true if the vector contains no elements\.

### func \(\*Vec\[T\]\) Iter

```go
func (vec *Vec[T]) Iter() Iterator[T]
```

Returns an Iterator to the vector elements\.

### func \(Vec\[T\]\) Len

```go
func (vec Vec[T]) Len() int
```

Returns the number of elements in the vector\, also referred to as its ‘length’\.

### func \(\*Vec\[T\]\) Pop

```go
func (vec *Vec[T]) Pop() option.Option[T]
```

Removes the last element from a vector and returns it\, or None if it is empty\.

### func \(\*Vec\[T\]\) Push

```go
func (vec *Vec[T]) Push(item T)
```

Appends an element to the back of a collection\.

### func \(\*Vec\[T\]\) Remove

```go
func (vec *Vec[T]) Remove(index int) T
```

Removes and returns the element at position index within the vector\, shifting all elements after it to the left\. Note: Because this shifts over the remaining elements\, it has a worst\-case performance of O\(n\)\. If you don’t need the order of elements to be preserved\, use SwapRemove instead\. Panics if index is out of bounds\.

### func \(\*Vec\[T\]\) Resize

```go
func (vec *Vec[T]) Resize(newLength int, value T)
```

Resizes the Vec in\-place so that len is equal to newLength\. If newLength is greater than len\, the Vec is extended by the difference\, with each additional slot filled with value\. If newLength is less than len\, the Vec is simply truncated\. Panics if index is less than zero

### func \(\*Vec\[T\]\) ResizeWith

```go
func (vec *Vec[T]) ResizeWith(newLength int, f func() T)
```

Resizes the Vec in\-place so that len is equal to newLength\. If newLength is greater than len\, the Vec is extended by the difference\, with each additional slot filled with the result of calling the function f\. The return values from f will end up in the Vec in the order they have been generated\.

If new\_len is less than len\, the Vec is simply truncated\. Panics if index is less than zero

### func \(\*Vec\[T\]\) Retain

```go
func (vec *Vec[T]) Retain(f func(T) bool)
```

Retains only the elements specified by the predicate\. In other words\, remove all elements e such that f\(&e\) returns false\. This method operates in place\, visiting each element exactly once in the original order\, and preserves the order of the retained elements\.

### func \(\*Vec\[T\]\) Splice

```go
func (vec *Vec[T]) Splice(start, end int, replaceWith Vec[T]) Vec[T]
```

Creates a splicing vector that replaces the specified range in the vector with the given replaceWith vector and returns the removed items\. replaceWith does not need to be the same length as range\. range is removed even if the vector is not consumed until the end\.

### func \(Vec\[T\]\) SplitOff

```go
func (vec Vec[T]) SplitOff(at int) Vec[T]
```

Splits the collection into two at the given index\. Returns a newly allocated vector containing the elements in the range \[at\, len\]\. After the call\, the original vector will be left containing the elements \[0\, at\] with its previous capacity unchanged\.

### func \(\*Vec\[T\]\) SwapRemove

```go
func (vec *Vec[T]) SwapRemove(index int) T
```

Removes an element from the vector and returns it\. The removed element is replaced by the last element of the vector\. This does not preserve ordering\, but is O\(1\)\. If you need to preserve the element order\, use remove instead\. Panics if index is out of bounds\.

### func \(\*Vec\[T\]\) Truncate

```go
func (vec *Vec[T]) Truncate(len int)
```

Shortens the vector\, keeping the first len elements and dropping the rest\. If len is greater than the vector’s current length\, this has no effect\. The drain method can emulate truncate\, but causes the excess elements to be returned instead of dropped\. Note that this method has no effect on the allocated capacity of the vector\. Panics if index is negative\.



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
