package pokecache

import (
	"reflect"
	"testing"
)

func TestNewCache(t *testing.T) {
	fail := false

	c := NewCache()

	expected := reflect.TypeFor[*Cache]()
	actual := reflect.TypeOf(c)

	if actual != expected {
		t.Log("NewCache() did not create cache pointer")
		fail = true
	}

	expected = reflect.TypeFor[map[string][]byte]()
	actual = reflect.TypeOf(c.cache)

	if actual != expected {
		t.Log("Cache.cache is not correct type")
		fail = true
	}

	if len(c.cache) != 0 {
		t.Log("Cache.cache is not initizaling empty")
		fail = true
	}

	if fail {
		t.FailNow()
	}
}

type addition struct {
	key string
	val []byte
}

func TestCacheAdd(t *testing.T) {
	tests := map[string]struct {
		additions []addition
		expected  map[string][]byte
	}{
		"addOne": {
			additions: []addition{
				{key: "pika", val: []byte("electric")},
			},
			expected: map[string][]byte{
				"pika": []byte("electric"),
			},
		},
		"addTwo": {
			additions: []addition{
				{key: "pika", val: []byte("electric")},
				{key: "squirt", val: []byte("water")},
			},
			expected: map[string][]byte{
				"pika":   []byte("electric"),
				"squirt": []byte("water"),
			},
		},
		"addDuplicate": {
			additions: []addition{
				{key: "pika", val: []byte("electric")},
				{key: "pika", val: []byte("water")},
			},
			expected: map[string][]byte{"pika": []byte("water")},
		},
		"addDuplicateWithMore": {
			additions: []addition{
				{key: "pika", val: []byte("electric")},
				{key: "pika", val: []byte("water")},
				{key: "pika2", val: []byte("e'lectric")},
			},
			expected: map[string][]byte{
				"pika":  []byte("water"),
				"pika2": []byte("e'lectric"),
			},
		},
		"addEmpty": {
			additions: []addition{{key: "", val: []byte("")}},
			expected:  map[string][]byte{},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := NewCache()

			for _, addition := range tt.additions {
				c.Add(addition.key, addition.val)
			}

			actual := c.cache

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Fatalf("expected: %v, got: %v", tt.expected, actual)
			}
		})
	}
}

func TestCacheGet(t *testing.T) {
	tests := map[string]struct {
		cache map[string][]byte
		get   string
		val   []byte
		flag  bool
	}{
		"singleItem": {
			cache: map[string][]byte{
				"pika": []byte("electric"),
			},
			get:  "pika",
			val:  []byte("electric"),
			flag: true,
		},
		"doubleItem": {
			cache: map[string][]byte{
				"pika":   []byte("electric"),
				"squirt": []byte("water"),
			},
			get:  "squirt",
			val:  []byte("water"),
			flag: true,
		},
		"noKey": {
			cache: map[string][]byte{
				"pika":   []byte("electric"),
				"squirt": []byte("water"),
			},
			get:  "char",
			val:  nil,
			flag: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := Cache{cache: tt.cache}

			v, f := c.Get(tt.get)

			if !reflect.DeepEqual(v, tt.val) {
				t.Fatalf("value not the same, expected: %v, got: %v", tt.val, v)
			}

			if f != tt.flag {
				t.Fatalf("success flag doesn't match, expected: %v, got: %v", tt.flag, f)
			}
		})
	}
}
