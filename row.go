package tsv

import (
	"github.com/smartwalle/conv4go"
	"strings"
)

type Row struct {
	values []string
}

func (this *Row) Len() int {
	return len(this.values)
}

func (this *Row) Value(index int) string {
	if index >= len(this.values) {
		return ""
	}
	return strings.TrimSpace(this.values[index])
}

func (this *Row) String(index int) string {
	return this.Value(index)
}

func (this *Row) Int(index int) int {
	var v = this.Value(index)
	return conv4go.Int(v)
}

func (this *Row) Int8(index int) int8 {
	var v = this.Value(index)
	return conv4go.Int8(v)
}

func (this *Row) Int16(index int) int16 {
	var v = this.Value(index)
	return conv4go.Int16(v)
}

func (this *Row) Int32(index int) int32 {
	var v = this.Value(index)
	return conv4go.Int32(v)
}

func (this *Row) Int64(index int) int64 {
	var v = this.Value(index)
	return conv4go.Int64(v)
}

func (this *Row) Uint(index int) uint {
	var v = this.Value(index)
	return conv4go.Uint(v)
}

func (this *Row) Uint8(index int) uint8 {
	var v = this.Value(index)
	return conv4go.Uint8(v)
}

func (this *Row) Uint16(index int) uint16 {
	var v = this.Value(index)
	return conv4go.Uint16(v)
}

func (this *Row) Uint32(index int) uint32 {
	var v = this.Value(index)
	return conv4go.Uint32(v)
}

func (this *Row) Uint64(index int) uint64 {
	var v = this.Value(index)
	return conv4go.Uint64(v)
}

func (this *Row) Bool(index int) bool {
	var v = this.Value(index)
	return conv4go.Bool(v)
}

func (this *Row) Float32(index int) float32 {
	var v = this.Value(index)
	return conv4go.Float32(v)
}

func (this *Row) Float64(index int) float64 {
	var v = this.Value(index)
	return conv4go.Float64(v)
}
