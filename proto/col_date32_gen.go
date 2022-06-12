// Code generated by ./cmd/ch-gen-col, DO NOT EDIT.

package proto

// ColDate32 represents Date32 column.
type ColDate32 []Date32

// Compile-time assertions for ColDate32.
var (
	_ ColInput  = ColDate32{}
	_ ColResult = (*ColDate32)(nil)
	_ Column    = (*ColDate32)(nil)
)

// Type returns ColumnType of Date32.
func (ColDate32) Type() ColumnType {
	return ColumnTypeDate32
}

// Rows returns count of rows in column.
func (c ColDate32) Rows() int {
	return len(c)
}

// Row returns i-th row of column.
func (c ColDate32) Row(i int) Date32 {
	return c[i]
}

// Append Date32 to column.
func (c *ColDate32) Append(v Date32) {
	*c = append(*c, v)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColDate32) Reset() {
	*c = (*c)[:0]
}

// LowCardinality returns LowCardinality for Date32 .
func (c *ColDate32) LowCardinality() *ColLowCardinality[Date32] {
	return &ColLowCardinality[Date32]{
		index: c,
	}
}

// Array is helper that creates Array of Date32.
func (c *ColDate32) Array() *ColArr[Date32] {
	return &ColArr[Date32]{
		Data: c,
	}
}

// Nullable is helper that creates Nullable(Date32).
func (c *ColDate32) Nullable() *ColNullable[Date32] {
	return &ColNullable[Date32]{
		Values: c,
	}
}

// NewArrDate32 returns new Array(Date32).
func NewArrDate32() *ColArr[Date32] {
	return &ColArr[Date32]{
		Data: new(ColDate32),
	}
}
