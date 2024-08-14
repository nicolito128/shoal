package shoal

// Morph0 is a static transformation of CurrentType.
type Morph0[CurrentType any] func(CurrentType)

// Morph is a CurrentType to NewType transformation.
type Morph[CurrentType, NewType any] func(CurrentType) NewType

// Morph2 is a KeyType and ValueType to NewValueType transformation.
type Morph2[KeyType, ValueType, NewValueType any] func(KeyType, ValueType) NewValueType

// MorphCond is a conditional transformation of T-type.
type MorphCond[T comparable] func(T) bool

// MorphCond2 is a conditional transformation of K-type and T-type.
type MorphCond2[K, V comparable] func(K, V) bool
