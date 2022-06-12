// Code generated by ./cmd/ch-gen-col, DO NOT EDIT.

package proto

// ColInt128 represents Int128 column.
type ColInt128 []Int128

// Compile-time assertions for ColInt128.
var (
	_ ColInput  = ColInt128{}
	_ ColResult = (*ColInt128)(nil)
	_ Column    = (*ColInt128)(nil)
)

// Type returns ColumnType of Int128.
func (ColInt128) Type() ColumnType {
	return ColumnTypeInt128
}

// Rows returns count of rows in column.
func (c ColInt128) Rows() int {
	return len(c)
}

// Row returns i-th row of column.
func (c ColInt128) Row(i int) Int128 {
	return c[i]
}

// Append Int128 to column.
func (c *ColInt128) Append(v Int128) {
	*c = append(*c, v)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColInt128) Reset() {
	*c = (*c)[:0]
}

// LowCardinality returns LowCardinality for Int128 .
func (c *ColInt128) LowCardinality() *ColLowCardinality[Int128] {
	return &ColLowCardinality[Int128]{
		index: c,
	}
}

// Array is helper that creates Array of Int128.
func (c *ColInt128) Array() *ColArr[Int128] {
	return &ColArr[Int128]{
		Data: c,
	}
}

// Nullable is helper that creates Nullable(Int128).
func (c *ColInt128) Nullable() *ColNullable[Int128] {
	return &ColNullable[Int128]{
		Values: c,
	}
}

// NewArrInt128 returns new Array(Int128).
func NewArrInt128() *ColArr[Int128] {
	return &ColArr[Int128]{
		Data: new(ColInt128),
	}
}
