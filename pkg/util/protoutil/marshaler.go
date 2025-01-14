// Copyright 2016 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package protoutil

import (
	"github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"
)

// ProtoPb is a gwruntime.Marshaler that uses github.com/gogo/protobuf/proto.
type ProtoPb struct{}

// ContentType implements gwruntime.Marshaler.
func (*ProtoPb) ContentType() string {
	// NB: This is the same as httputil.ProtoContentType which we can't use due
	// to an import cycle.
	const ProtoContentType = "application/x-protobuf"
	return ProtoContentType
}

// Marshal implements gwruntime.Marshaler.
func (*ProtoPb) Marshal(v interface{}) ([]byte, error) {
	// NB: we use proto.Message here because grpc-gateway passes us protos that
	// we don't control and thus don't implement protoutil.Message.
	if p, ok := v.(proto.Message); ok {
		return proto.Marshal(p)
	}
	return nil, errors.Errorf("unexpected type %T does not implement %s", v, typeProtoMessage)
}

// Unmarshal implements gwruntime.Marshaler.
func (*ProtoPb) Unmarshal(data []byte, v interface{}) error {
	// NB: we use proto.Message here because grpc-gateway passes us protos that
	// we don't control and thus don't implement protoutil.Message.
	if p, ok := v.(proto.Message); ok {
		return proto.Unmarshal(data, p)
	}
	return errors.Errorf("unexpected type %T does not implement %s", v, typeProtoMessage)
}

// Delimiter implements gwruntime.Delimited.
func (*ProtoPb) Delimiter() []byte {
	return nil
}
