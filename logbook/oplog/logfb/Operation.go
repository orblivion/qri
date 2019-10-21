// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package logfb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Operation struct {
	_tab flatbuffers.Table
}

func GetRootAsOperation(buf []byte, offset flatbuffers.UOffsetT) *Operation {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Operation{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Operation) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Operation) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Operation) Type() OpType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return OpType(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Operation) MutateType(n OpType) bool {
	return rcv._tab.MutateInt8Slot(4, int8(n))
}

func (rcv *Operation) Model() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Operation) MutateModel(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

func (rcv *Operation) Ref() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Operation) Prev() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Operation) Relations(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *Operation) RelationsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Operation) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Operation) AuthorID() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Operation) Timestamp() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Operation) MutateTimestamp(n int64) bool {
	return rcv._tab.MutateInt64Slot(18, n)
}

func (rcv *Operation) Size() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Operation) MutateSize(n uint64) bool {
	return rcv._tab.MutateUint64Slot(20, n)
}

func (rcv *Operation) Note() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func OperationStart(builder *flatbuffers.Builder) {
	builder.StartObject(10)
}
func OperationAddType(builder *flatbuffers.Builder, type_ OpType) {
	builder.PrependInt8Slot(0, int8(type_), 0)
}
func OperationAddModel(builder *flatbuffers.Builder, model uint32) {
	builder.PrependUint32Slot(1, model, 0)
}
func OperationAddRef(builder *flatbuffers.Builder, ref flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(ref), 0)
}
func OperationAddPrev(builder *flatbuffers.Builder, prev flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(prev), 0)
}
func OperationAddRelations(builder *flatbuffers.Builder, relations flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(relations), 0)
}
func OperationStartRelationsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func OperationAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(name), 0)
}
func OperationAddAuthorID(builder *flatbuffers.Builder, authorID flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(authorID), 0)
}
func OperationAddTimestamp(builder *flatbuffers.Builder, timestamp int64) {
	builder.PrependInt64Slot(7, timestamp, 0)
}
func OperationAddSize(builder *flatbuffers.Builder, size uint64) {
	builder.PrependUint64Slot(8, size, 0)
}
func OperationAddNote(builder *flatbuffers.Builder, note flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(9, flatbuffers.UOffsetT(note), 0)
}
func OperationEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}