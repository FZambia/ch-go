// Code generated by ./cmd/ch-gen-col, DO NOT EDIT.

package proto

// ColInt256 represents Int256 column.
type ColInt256 []Int256

// Compile-time assertions for ColInt256.
var (
	_ ColInput  = ColInt256{}
	_ ColResult = (*ColInt256)(nil)
	_ Column    = (*ColInt256)(nil)
)

// Type returns ColumnType of Int256.
func (ColInt256) Type() ColumnType {
	return ColumnTypeInt256
}

// Rows returns count of rows in column.
func (c ColInt256) Rows() int {
	return len(c)
}

// Row returns i-th row of column.
func (c ColInt256) Row(i int) Int256 {
	return c[i]
}

// Append Int256 to column.
func (c *ColInt256) Append(v Int256) {
	*c = append(*c, v)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColInt256) Reset() {
	*c = (*c)[:0]
}

// LowCardinality returns LowCardinality for Int256 .
func (c *ColInt256) LowCardinality() *ColLowCardinality[Int256] {
	return &ColLowCardinality[Int256]{
		index: c,
	}
}

// Array is helper that creates Array of Int256.
func (c *ColInt256) Array() *ColArr[Int256] {
	return &ColArr[Int256]{
		Data: c,
	}
}

// Nullable is helper that creates Nullable(Int256).
func (c *ColInt256) Nullable() *ColNullable[Int256] {
	return &ColNullable[Int256]{
		Values: c,
	}
}

// NewArrInt256 returns new Array(Int256).
func NewArrInt256() *ColArr[Int256] {
	return &ColArr[Int256]{
		Data: new(ColInt256),
	}
}
