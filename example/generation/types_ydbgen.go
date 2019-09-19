// Code generated by ydbgen; DO NOT EDIT.

package main

import (
	"strconv"

	"github.com/yandex-cloud/ydb-go-sdk"
	"github.com/yandex-cloud/ydb-go-sdk/table"
)

var (
	_ = strconv.Itoa
	_ = ydb.StringValue
	_ = table.NewQueryParameters
)

func (u *User) Scan(res *table.Result) (err error) {
	res.SeekItem("id")
	u.ID = res.OUint64()

	res.SeekItem("username")
	u.Username = res.OUTF8()

	res.SeekItem("mode")
	u.Mode = ydbConvU64ToU8(res.OUint64())

	res.SeekItem("magic")
	u.Magic = uint(res.OUint32())

	res.SeekItem("score")
	res.Unwrap()
	if !res.IsNull() {
		x0 := res.Int64()
		u.Score.Set(x0)
	}

	res.SeekItem("updated")
	res.Unwrap()
	if !res.IsNull() {
		x0 := res.Timestamp()
		err := (*ydb.Time)(&u.Updated).FromTimestamp(x0)
		if err != nil {
			panic("ydbgen: date type conversion failed: " + err.Error())
		}
	}

	return res.Err()
}

func (u *User) QueryParameters() *table.QueryParameters {
	var v0 ydb.Value
	{
		vp0 := ydb.OptionalValue(ydb.Uint64Value(u.ID))
		v0 = vp0
	}
	var v1 ydb.Value
	{
		vp0 := ydb.OptionalValue(ydb.UTF8Value(u.Username))
		v1 = vp0
	}
	var v2 ydb.Value
	{
		vp0 := ydb.OptionalValue(ydb.Uint64Value(uint64(u.Mode)))
		v2 = vp0
	}
	var v3 ydb.Value
	{
		vp0 := ydb.OptionalValue(ydb.Uint32Value(uint32(u.Magic)))
		v3 = vp0
	}
	var v4 ydb.Value
	{
		var v5 ydb.Value
		x0, ok0 := u.Score.Get()
		if ok0 {
			v5 = ydb.OptionalValue(ydb.Int64Value(x0))
		} else {
			v5 = ydb.NullValue(ydb.TypeInt64)
		}
		v4 = v5
	}
	var v5 ydb.Value
	{
		var v6 ydb.Value
		var x0 uint64
		ok0 := !u.Updated.IsZero()
		if ok0 {
			x0 = ydb.Time(u.Updated).Timestamp()
		}
		if ok0 {
			v6 = ydb.OptionalValue(ydb.TimestampValue(x0))
		} else {
			v6 = ydb.NullValue(ydb.TypeTimestamp)
		}
		v5 = v6
	}
	return table.NewQueryParameters(
		table.ValueParam("$id", v0),
		table.ValueParam("$username", v1),
		table.ValueParam("$mode", v2),
		table.ValueParam("$magic", v3),
		table.ValueParam("$score", v4),
		table.ValueParam("$updated", v5),
	)
}

func (u *User) StructValue() ydb.Value {
	var v0 ydb.Value
	{
		var v1 ydb.Value
		{
			vp0 := ydb.OptionalValue(ydb.Uint64Value(u.ID))
			v1 = vp0
		}
		var v2 ydb.Value
		{
			vp0 := ydb.OptionalValue(ydb.UTF8Value(u.Username))
			v2 = vp0
		}
		var v3 ydb.Value
		{
			vp0 := ydb.OptionalValue(ydb.Uint64Value(uint64(u.Mode)))
			v3 = vp0
		}
		var v4 ydb.Value
		{
			vp0 := ydb.OptionalValue(ydb.Uint32Value(uint32(u.Magic)))
			v4 = vp0
		}
		var v5 ydb.Value
		{
			var v6 ydb.Value
			x0, ok0 := u.Score.Get()
			if ok0 {
				v6 = ydb.OptionalValue(ydb.Int64Value(x0))
			} else {
				v6 = ydb.NullValue(ydb.TypeInt64)
			}
			v5 = v6
		}
		var v6 ydb.Value
		{
			var v7 ydb.Value
			var x0 uint64
			ok0 := !u.Updated.IsZero()
			if ok0 {
				x0 = ydb.Time(u.Updated).Timestamp()
			}
			if ok0 {
				v7 = ydb.OptionalValue(ydb.TimestampValue(x0))
			} else {
				v7 = ydb.NullValue(ydb.TypeTimestamp)
			}
			v6 = v7
		}
		v0 = ydb.StructValue(
			ydb.StructFieldValue("id", v1),
			ydb.StructFieldValue("username", v2),
			ydb.StructFieldValue("mode", v3),
			ydb.StructFieldValue("magic", v4),
			ydb.StructFieldValue("score", v5),
			ydb.StructFieldValue("updated", v6),
		)
	}
	return v0
}

func (u *User) StructType() ydb.Type {
	var t0 ydb.Type
	{
		fs0 := make([]ydb.StructOption, 6)
		var t1 ydb.Type
		{
			tp0 := ydb.TypeUint64
			t1 = ydb.Optional(tp0)
		}
		fs0[0] = ydb.StructField("id", t1)
		var t2 ydb.Type
		{
			tp0 := ydb.TypeUTF8
			t2 = ydb.Optional(tp0)
		}
		fs0[1] = ydb.StructField("username", t2)
		var t3 ydb.Type
		{
			tp0 := ydb.TypeUint64
			t3 = ydb.Optional(tp0)
		}
		fs0[2] = ydb.StructField("mode", t3)
		var t4 ydb.Type
		{
			tp0 := ydb.TypeUint32
			t4 = ydb.Optional(tp0)
		}
		fs0[3] = ydb.StructField("magic", t4)
		var t5 ydb.Type
		{
			tp0 := ydb.TypeInt64
			t5 = ydb.Optional(tp0)
		}
		fs0[4] = ydb.StructField("score", t5)
		var t6 ydb.Type
		{
			tp0 := ydb.TypeTimestamp
			t6 = ydb.Optional(tp0)
		}
		fs0[5] = ydb.StructField("updated", t6)
		t0 = ydb.Struct(fs0...)
	}
	return t0
}

func (m *MagicUsers) Scan(res *table.Result) (err error) {
	res.NextItem()
	m.Magic = uint(res.OUint32())

	res.NextItem()
	n0 := res.ListIn()
	xs0 := make([]string, n0)
	for i0 := 0; i0 < n0; i0++ {
		res.ListItem(i0)
		var x0 string
		x0 = res.UTF8()
		xs0[i0] = x0
	}
	m.Users = xs0
	res.ListOut()

	return res.Err()
}

func (us *Users) Scan(res *table.Result) (err error) {
	for res.NextRow() {
		var x0 User
		res.SeekItem("id")
		x0.ID = res.OUint64()

		res.SeekItem("username")
		x0.Username = res.OUTF8()

		res.SeekItem("mode")
		x0.Mode = ydbConvU64ToU8(res.OUint64())

		res.SeekItem("magic")
		x0.Magic = uint(res.OUint32())

		res.SeekItem("score")
		res.Unwrap()
		if !res.IsNull() {
			x1 := res.Int64()
			x0.Score.Set(x1)
		}

		res.SeekItem("updated")
		res.Unwrap()
		if !res.IsNull() {
			x1 := res.Timestamp()
			err := (*ydb.Time)(&x0.Updated).FromTimestamp(x1)
			if err != nil {
				panic("ydbgen: date type conversion failed: " + err.Error())
			}
		}

		if res.Err() == nil {
			*us = append(*us, x0)
		}
	}
	return res.Err()
}

func (us Users) ListValue() ydb.Value {
	var list0 ydb.Value
	vs0 := make([]ydb.Value, len(us))
	for i0, item0 := range us {
		var v0 ydb.Value
		{
			var v1 ydb.Value
			{
				var v2 ydb.Value
				{
					vp0 := ydb.OptionalValue(ydb.Uint64Value(item0.ID))
					v2 = vp0
				}
				var v3 ydb.Value
				{
					vp0 := ydb.OptionalValue(ydb.UTF8Value(item0.Username))
					v3 = vp0
				}
				var v4 ydb.Value
				{
					vp0 := ydb.OptionalValue(ydb.Uint64Value(uint64(item0.Mode)))
					v4 = vp0
				}
				var v5 ydb.Value
				{
					vp0 := ydb.OptionalValue(ydb.Uint32Value(uint32(item0.Magic)))
					v5 = vp0
				}
				var v6 ydb.Value
				{
					var v7 ydb.Value
					x0, ok0 := item0.Score.Get()
					if ok0 {
						v7 = ydb.OptionalValue(ydb.Int64Value(x0))
					} else {
						v7 = ydb.NullValue(ydb.TypeInt64)
					}
					v6 = v7
				}
				var v7 ydb.Value
				{
					var v8 ydb.Value
					var x0 uint64
					ok0 := !item0.Updated.IsZero()
					if ok0 {
						x0 = ydb.Time(item0.Updated).Timestamp()
					}
					if ok0 {
						v8 = ydb.OptionalValue(ydb.TimestampValue(x0))
					} else {
						v8 = ydb.NullValue(ydb.TypeTimestamp)
					}
					v7 = v8
				}
				v1 = ydb.StructValue(
					ydb.StructFieldValue("id", v2),
					ydb.StructFieldValue("username", v3),
					ydb.StructFieldValue("mode", v4),
					ydb.StructFieldValue("magic", v5),
					ydb.StructFieldValue("score", v6),
					ydb.StructFieldValue("updated", v7),
				)
			}
			v0 = v1
		}
		vs0[i0] = v0
	}
	if len(vs0) == 0 {
		var t1 ydb.Type
		{
			var t2 ydb.Type
			{
				fs0 := make([]ydb.StructOption, 6)
				var t3 ydb.Type
				{
					tp0 := ydb.TypeUint64
					t3 = ydb.Optional(tp0)
				}
				fs0[0] = ydb.StructField("id", t3)
				var t4 ydb.Type
				{
					tp0 := ydb.TypeUTF8
					t4 = ydb.Optional(tp0)
				}
				fs0[1] = ydb.StructField("username", t4)
				var t5 ydb.Type
				{
					tp0 := ydb.TypeUint64
					t5 = ydb.Optional(tp0)
				}
				fs0[2] = ydb.StructField("mode", t5)
				var t6 ydb.Type
				{
					tp0 := ydb.TypeUint32
					t6 = ydb.Optional(tp0)
				}
				fs0[3] = ydb.StructField("magic", t6)
				var t7 ydb.Type
				{
					tp0 := ydb.TypeInt64
					t7 = ydb.Optional(tp0)
				}
				fs0[4] = ydb.StructField("score", t7)
				var t8 ydb.Type
				{
					tp0 := ydb.TypeTimestamp
					t8 = ydb.Optional(tp0)
				}
				fs0[5] = ydb.StructField("updated", t8)
				t2 = ydb.Struct(fs0...)
			}
			t1 = t2
		}
		t0 := ydb.List(t1)
		list0 = ydb.ZeroValue(t0)
	} else {
		list0 = ydb.ListValue(vs0...)
	}
	return list0
}

func (ms *MagicUsersList) Scan(res *table.Result) (err error) {
	for res.NextRow() {
		var x0 MagicUsers
		res.NextItem()
		x0.Magic = uint(res.OUint32())

		res.NextItem()
		n0 := res.ListIn()
		xs0 := make([]string, n0)
		for i0 := 0; i0 < n0; i0++ {
			res.ListItem(i0)
			var x1 string
			x1 = res.UTF8()
			xs0[i0] = x1
		}
		x0.Users = xs0
		res.ListOut()

		if res.Err() == nil {
			*ms = append(*ms, x0)
		}
	}
	return res.Err()
}

func ydbConvU64ToU8(x uint64) uint8 { 
	const (
		bits = 8
		mask = (1 << (bits)) - 1
	)
	abs := uint64(x)
	if abs&mask != abs {
		panic(
			"ydbgen: convassert: " + strconv.FormatUint(uint64(x), 10) +
				" (type uint64) overflows uint8",
		)
	}
	return uint8(x)
}

