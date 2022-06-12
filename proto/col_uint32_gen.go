// Code generated by ./cmd/ch-gen-col, DO NOT EDIT.

package proto

// ColUInt32 represents UInt32 column.
type ColUInt32 []uint32

// Compile-time assertions for ColUInt32.
var (
	_ ColInput  = ColUInt32{}
	_ ColResult = (*ColUInt32)(nil)
	_ Column    = (*ColUInt32)(nil)
)

// Type returns ColumnType of UInt32.
func (ColUInt32) Type() ColumnType {
	return ColumnTypeUInt32
}

// Rows returns count of rows in column.
func (c ColUInt32) Rows() int {
	return len(c)
}

// Row returns i-th row of column.
func (c ColUInt32) Row(i int) uint32 {
	return c[i]
}

// Append uint32 to column.
func (c *ColUInt32) Append(v uint32) {
	*c = append(*c, v)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColUInt32) Reset() {
	*c = (*c)[:0]
}

// LowCardinality returns LowCardinality for UInt32 .
func (c *ColUInt32) LowCardinality() *ColLowCardinality[uint32] {
	return &ColLowCardinality[uint32]{
		index: c,
	}
}

// Array is helper that creates Array of uint32.
func (c *ColUInt32) Array() *ColArr[uint32] {
	return &ColArr[uint32]{
		Data: c,
	}
}

// Nullable is helper that creates Nullable(uint32).
func (c *ColUInt32) Nullable() *ColNullable[uint32] {
	return &ColNullable[uint32]{
		Values: c,
	}
}

// NewArrUInt32 returns new Array(UInt32).
func NewArrUInt32() *ColArr[uint32] {
	return &ColArr[uint32]{
		Data: new(ColUInt32),
	}
}
