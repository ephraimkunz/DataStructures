// Package hashtable implements a hashtable using a base array with linked lists for
// collisions. Keys are strings, values are ints, and a key is hashed using hash/fnv.
package hashtable

import (
	"hash/fnv"
)

type Hashtable struct {
	table []*node // Array of pointer to heads of linked lists
}

type node struct {
	key   string
	value int
	next  *node
}

const defaultBuckets = 50

// NewHashtable creates and returns a new hashtable.
func NewHashtable() Hashtable {
	table := make([]*node, defaultBuckets)
	hashtable := Hashtable{table}
	return hashtable
}

func hashKey(key string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32()
}

func (ht *Hashtable) numBuckets() int {
	return len(ht.table)
}

// Insert inserts a new key with the given value if not exists,
// or changes an existing key to the given value.
func (ht *Hashtable) Insert(key string, value int) {
	bucket := hashKey(key) % uint32(ht.numBuckets())

	ll := ht.table[bucket]

	if ll == nil { // First item in this bucket
		newNode := &node{key, value, nil}
		ht.table[bucket] = newNode
		return
	}

	for {
		if ll.key == key {
			ll.value = value
			return
		}

		if ll.next == nil {
			break
		}
		ll = ll.next
	}

	ll.next = &node{key, value, nil}
}

// Remove removes the key and value associated with key.
// Returns true if remove was successful, else false
func (ht *Hashtable) Remove(key string) bool {
	bucket := hashKey(key) % uint32(ht.numBuckets())

	ll := ht.table[bucket]

	if ll == nil { // No items in bucket
		return false
	}

	var prev *node // Nil by default
	for ; ll != nil; prev, ll = ll, ll.next {
		if ll.key == key {
			if prev == nil {
				ht.table[bucket] = ll.next
			} else {
				prev.next = ll.next
			}
			return true
		}
	}

	return false
}

// Get returns a value for a key, and returns false if the
// key was not present.
func (ht *Hashtable) Get(key string) (int, bool) {
	bucket := hashKey(key) % uint32(ht.numBuckets())

	ll := ht.table[bucket]

	if ll == nil {
		return 0, false
	}

	for ; ll != nil; ll = ll.next {
		if ll.key == key {
			return ll.value, true
		}
	}

	return 0, false
}
