// Copyright 2022 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: sigstore_bundle.proto

package v1

import (
	v1 "github.com/sigstore/protobuf-specs/gen/pb-go/common/v1"
	dsse "github.com/sigstore/protobuf-specs/gen/pb-go/dsse"
	v11 "github.com/sigstore/protobuf-specs/gen/pb-go/rekor/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Various timestamped counter signatures over the artifacts signature.
// Currently only RFC3161 signatures are provided. More formats may be added
// in the future.
type TimestampVerificationData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of RFC3161 signed timestamps provided by the user.
	// This can be used when the entry has not been stored on a
	// transparency log, or in conjunction for a stronger trust model.
	// Clients MUST verify the hashed message in the message imprint
	// against the signature in the bundle.
	Rfc3161Timestamps []*v1.RFC3161SignedTimestamp `protobuf:"bytes,1,rep,name=rfc3161_timestamps,json=rfc3161Timestamps,proto3" json:"rfc3161_timestamps,omitempty"`
}

func (x *TimestampVerificationData) Reset() {
	*x = TimestampVerificationData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sigstore_bundle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimestampVerificationData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimestampVerificationData) ProtoMessage() {}

func (x *TimestampVerificationData) ProtoReflect() protoreflect.Message {
	mi := &file_sigstore_bundle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimestampVerificationData.ProtoReflect.Descriptor instead.
func (*TimestampVerificationData) Descriptor() ([]byte, []int) {
	return file_sigstore_bundle_proto_rawDescGZIP(), []int{0}
}

func (x *TimestampVerificationData) GetRfc3161Timestamps() []*v1.RFC3161SignedTimestamp {
	if x != nil {
		return x.Rfc3161Timestamps
	}
	return nil
}

// VerificationMaterial captures details on the materials used to verify
// signatures. This message may be embedded in a DSSE envelope as a signature
// extension. Specifically, the `ext` field of the extension will expect this
// message when the signature extension is for Sigstore. This is identified by
// the `kind` field in the extension, which must be set to
// application/vnd.dev.sigstore.verificationmaterial;version=0.1 for Sigstore.
// When used as a DSSE extension, if the `public_key` field is used to indicate
// the key identifier, it MUST match the `keyid` field of the signature the
// extension is attached to.
type VerificationMaterial struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The key material for verification purposes.
	//
	// This allows key material to be conveyed in one of three forms:
	//
	//  1. An unspecified public key identifier, for retrieving a key
	//     from an out-of-band mechanism (such as a keyring);
	//
	//  2. A sequence of one or more X.509 certificates, of which the first member
	//     MUST be a leaf certificate conveying the signing key. Subsequent members
	//     SHOULD be in issuing order, meaning that `n + 1` should be an issuer for `n`.
	//
	//     Signers MUST NOT include root CA certificates in bundles, and SHOULD NOT
	//     include intermediate CA certificates that appear in an independent root of trust
	//     (such as the Public Good Instance's trusted root).
	//
	//     Verifiers MUST validate the chain carefully to ensure that it chains up
	//     to a CA certificate that they independently trust. Verifiers SHOULD
	//     handle old or non-complying bundles that have superfluous intermediate and/or
	//     root CA certificates by either ignoring them or explicitly considering them
	//     untrusted for the purposes of chain building.
	//
	//  3. A single X.509 certificate, which MUST be a leaf certificate conveying
	//     the signing key.
	//
	// When used with the Public Good Instance (PGI) of Sigstore for "keyless" signing
	// via Fulcio, form (1) MUST NOT be used, regardless of bundle version. Form (1)
	// MAY be used with the PGI for self-managed keys.
	//
	// When used in a `0.1` or `0.2` bundle with the PGI and "keyless" signing,
	// form (2) MUST be used.
	//
	// When used in a `0.3` bundle with the PGI and "keyless" signing,
	// form (3) MUST be used.
	//
	// Types that are assignable to Content:
	//
	//	*VerificationMaterial_PublicKey
	//	*VerificationMaterial_X509CertificateChain
	//	*VerificationMaterial_Certificate
	Content isVerificationMaterial_Content `protobuf_oneof:"content"`
	// An inclusion proof and an optional signed timestamp from the log.
	// Client verification libraries MAY provide an option to support v0.1
	// bundles for backwards compatibility, which may contain an inclusion
	// promise and not an inclusion proof. In this case, the client MUST
	// validate the promise.
	// Verifiers SHOULD NOT allow v0.1 bundles if they're used in an
	// ecosystem which never produced them.
	TlogEntries []*v11.TransparencyLogEntry `protobuf:"bytes,3,rep,name=tlog_entries,json=tlogEntries,proto3" json:"tlog_entries,omitempty"`
	// Timestamp may also come from
	// tlog_entries.inclusion_promise.signed_entry_timestamp.
	TimestampVerificationData *TimestampVerificationData `protobuf:"bytes,4,opt,name=timestamp_verification_data,json=timestampVerificationData,proto3" json:"timestamp_verification_data,omitempty"`
}

func (x *VerificationMaterial) Reset() {
	*x = VerificationMaterial{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sigstore_bundle_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerificationMaterial) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerificationMaterial) ProtoMessage() {}

func (x *VerificationMaterial) ProtoReflect() protoreflect.Message {
	mi := &file_sigstore_bundle_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerificationMaterial.ProtoReflect.Descriptor instead.
func (*VerificationMaterial) Descriptor() ([]byte, []int) {
	return file_sigstore_bundle_proto_rawDescGZIP(), []int{1}
}

func (m *VerificationMaterial) GetContent() isVerificationMaterial_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *VerificationMaterial) GetPublicKey() *v1.PublicKeyIdentifier {
	if x, ok := x.GetContent().(*VerificationMaterial_PublicKey); ok {
		return x.PublicKey
	}
	return nil
}

func (x *VerificationMaterial) GetX509CertificateChain() *v1.X509CertificateChain {
	if x, ok := x.GetContent().(*VerificationMaterial_X509CertificateChain); ok {
		return x.X509CertificateChain
	}
	return nil
}

func (x *VerificationMaterial) GetCertificate() *v1.X509Certificate {
	if x, ok := x.GetContent().(*VerificationMaterial_Certificate); ok {
		return x.Certificate
	}
	return nil
}

func (x *VerificationMaterial) GetTlogEntries() []*v11.TransparencyLogEntry {
	if x != nil {
		return x.TlogEntries
	}
	return nil
}

func (x *VerificationMaterial) GetTimestampVerificationData() *TimestampVerificationData {
	if x != nil {
		return x.TimestampVerificationData
	}
	return nil
}

type isVerificationMaterial_Content interface {
	isVerificationMaterial_Content()
}

type VerificationMaterial_PublicKey struct {
	PublicKey *v1.PublicKeyIdentifier `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3,oneof"`
}

type VerificationMaterial_X509CertificateChain struct {
	X509CertificateChain *v1.X509CertificateChain `protobuf:"bytes,2,opt,name=x509_certificate_chain,json=x509CertificateChain,proto3,oneof"`
}

type VerificationMaterial_Certificate struct {
	Certificate *v1.X509Certificate `protobuf:"bytes,5,opt,name=certificate,proto3,oneof"`
}

func (*VerificationMaterial_PublicKey) isVerificationMaterial_Content() {}

func (*VerificationMaterial_X509CertificateChain) isVerificationMaterial_Content() {}

func (*VerificationMaterial_Certificate) isVerificationMaterial_Content() {}

type Bundle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// MUST be application/vnd.dev.sigstore.bundle.v0.3+json when
	// when encoded as JSON.
	// Clients must to be able to accept media type using the previously
	// defined formats:
	// * application/vnd.dev.sigstore.bundle+json;version=0.1
	// * application/vnd.dev.sigstore.bundle+json;version=0.2
	// * application/vnd.dev.sigstore.bundle+json;version=0.3
	MediaType string `protobuf:"bytes,1,opt,name=media_type,json=mediaType,proto3" json:"media_type,omitempty"`
	// When a signer is identified by a X.509 certificate, a verifier MUST
	// verify that the signature was computed at the time the certificate
	// was valid as described in the Sigstore client spec: "Verification
	// using a Bundle".
	// <https://docs.google.com/document/d/1kbhK2qyPPk8SLavHzYSDM8-Ueul9_oxIMVFuWMWKz0E/edit#heading=h.x8bduppe89ln>
	// If the verification material contains a public key identifier
	// (key hint) and the `content` is a DSSE envelope, the key hints
	// MUST be exactly the same in the verification material and in the
	// DSSE envelope.
	VerificationMaterial *VerificationMaterial `protobuf:"bytes,2,opt,name=verification_material,json=verificationMaterial,proto3" json:"verification_material,omitempty"`
	// Types that are assignable to Content:
	//
	//	*Bundle_MessageSignature
	//	*Bundle_DsseEnvelope
	Content isBundle_Content `protobuf_oneof:"content"`
}

func (x *Bundle) Reset() {
	*x = Bundle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sigstore_bundle_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bundle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bundle) ProtoMessage() {}

func (x *Bundle) ProtoReflect() protoreflect.Message {
	mi := &file_sigstore_bundle_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bundle.ProtoReflect.Descriptor instead.
func (*Bundle) Descriptor() ([]byte, []int) {
	return file_sigstore_bundle_proto_rawDescGZIP(), []int{2}
}

func (x *Bundle) GetMediaType() string {
	if x != nil {
		return x.MediaType
	}
	return ""
}

func (x *Bundle) GetVerificationMaterial() *VerificationMaterial {
	if x != nil {
		return x.VerificationMaterial
	}
	return nil
}

func (m *Bundle) GetContent() isBundle_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *Bundle) GetMessageSignature() *v1.MessageSignature {
	if x, ok := x.GetContent().(*Bundle_MessageSignature); ok {
		return x.MessageSignature
	}
	return nil
}

func (x *Bundle) GetDsseEnvelope() *dsse.Envelope {
	if x, ok := x.GetContent().(*Bundle_DsseEnvelope); ok {
		return x.DsseEnvelope
	}
	return nil
}

type isBundle_Content interface {
	isBundle_Content()
}

type Bundle_MessageSignature struct {
	MessageSignature *v1.MessageSignature `protobuf:"bytes,3,opt,name=message_signature,json=messageSignature,proto3,oneof"`
}

type Bundle_DsseEnvelope struct {
	// A DSSE envelope can contain arbitrary payloads.
	// Verifiers must verify that the payload type is a
	// supported and expected type. This is part of the DSSE
	// protocol which is defined here:
	// <https://github.com/secure-systems-lab/dsse/blob/master/protocol.md>
	// DSSE envelopes in a bundle MUST have exactly one signture.
	// This is a limitation from the DSSE spec, as it can contain
	// multiple signatures. There are two primary reasons:
	//  1. It simplfies the verification logic and policy
	//  2. The bundle (currently) can only contain a single
	//     instance of the required verification materials
	//
	// During verification a client MUST reject an envelope if
	// the number of signatures is not equal to one.
	DsseEnvelope *dsse.Envelope `protobuf:"bytes,4,opt,name=dsse_envelope,json=dsseEnvelope,proto3,oneof"`
}

func (*Bundle_MessageSignature) isBundle_Content() {}

func (*Bundle_DsseEnvelope) isBundle_Content() {}

var File_sigstore_bundle_proto protoreflect.FileDescriptor

var file_sigstore_bundle_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x69, 0x67, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x62, 0x75, 0x6e, 0x64, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x69, 0x67,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0e, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x73, 0x69, 0x67, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x69, 0x67, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x5f, 0x72, 0x65, 0x6b, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a,
	0x19, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x5d, 0x0a, 0x12, 0x72, 0x66,
	0x63, 0x33, 0x31, 0x36, 0x31, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x69, 0x67,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x46, 0x43, 0x33, 0x31, 0x36, 0x31, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x11, 0x72, 0x66, 0x63, 0x33, 0x31, 0x36, 0x31, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x73, 0x22, 0xf4, 0x03, 0x0a, 0x14, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x12, 0x51, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x69, 0x67,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x48, 0x00, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x69, 0x0a, 0x16, 0x78, 0x35, 0x30, 0x39, 0x5f, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x69, 0x67, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x58,
	0x35, 0x30, 0x39, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x43, 0x68,
	0x61, 0x69, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x48, 0x00, 0x52, 0x14, 0x78, 0x35, 0x30, 0x39,
	0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x12, 0x50, 0x0a, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x69, 0x67, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x58,
	0x35, 0x30, 0x39, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x48, 0x00, 0x52, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x12, 0x4e, 0x0a, 0x0c, 0x74, 0x6c, 0x6f, 0x67, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73,
	0x69, 0x67, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x65, 0x6b, 0x6f, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4c, 0x6f, 0x67,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x74, 0x6c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x71, 0x0a, 0x1b, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x69,
	0x67, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x19, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x61, 0x74, 0x61, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x22, 0xbf, 0x02, 0x0a, 0x06, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x66, 0x0a, 0x15, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x64, 0x65, 0x76, 0x2e,
	0x73, 0x69, 0x67, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x14, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x12, 0x5c, 0x0a, 0x11, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x64, 0x65, 0x76, 0x2e, 0x73, 0x69, 0x67, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x48, 0x00, 0x52, 0x10,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x3f, 0x0a, 0x0d, 0x64, 0x73, 0x73, 0x65, 0x5f, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x69, 0x6f, 0x2e, 0x69, 0x6e, 0x74,
	0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41,
	0x02, 0x48, 0x00, 0x52, 0x0c, 0x64, 0x73, 0x73, 0x65, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70,
	0x65, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4a, 0x04, 0x08, 0x05,
	0x10, 0x33, 0x42, 0x7c, 0x0a, 0x1c, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x69, 0x67, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x42, 0x0b, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69,
	0x67, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d,
	0x73, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x2d, 0x67, 0x6f, 0x2f,
	0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0xea, 0x02, 0x14, 0x53, 0x69, 0x67, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x3a, 0x3a, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sigstore_bundle_proto_rawDescOnce sync.Once
	file_sigstore_bundle_proto_rawDescData = file_sigstore_bundle_proto_rawDesc
)

func file_sigstore_bundle_proto_rawDescGZIP() []byte {
	file_sigstore_bundle_proto_rawDescOnce.Do(func() {
		file_sigstore_bundle_proto_rawDescData = protoimpl.X.CompressGZIP(file_sigstore_bundle_proto_rawDescData)
	})
	return file_sigstore_bundle_proto_rawDescData
}

var file_sigstore_bundle_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sigstore_bundle_proto_goTypes = []interface{}{
	(*TimestampVerificationData)(nil), // 0: dev.sigstore.bundle.v1.TimestampVerificationData
	(*VerificationMaterial)(nil),      // 1: dev.sigstore.bundle.v1.VerificationMaterial
	(*Bundle)(nil),                    // 2: dev.sigstore.bundle.v1.Bundle
	(*v1.RFC3161SignedTimestamp)(nil), // 3: dev.sigstore.common.v1.RFC3161SignedTimestamp
	(*v1.PublicKeyIdentifier)(nil),    // 4: dev.sigstore.common.v1.PublicKeyIdentifier
	(*v1.X509CertificateChain)(nil),   // 5: dev.sigstore.common.v1.X509CertificateChain
	(*v1.X509Certificate)(nil),        // 6: dev.sigstore.common.v1.X509Certificate
	(*v11.TransparencyLogEntry)(nil),  // 7: dev.sigstore.rekor.v1.TransparencyLogEntry
	(*v1.MessageSignature)(nil),       // 8: dev.sigstore.common.v1.MessageSignature
	(*dsse.Envelope)(nil),             // 9: io.intoto.Envelope
}
var file_sigstore_bundle_proto_depIdxs = []int32{
	3, // 0: dev.sigstore.bundle.v1.TimestampVerificationData.rfc3161_timestamps:type_name -> dev.sigstore.common.v1.RFC3161SignedTimestamp
	4, // 1: dev.sigstore.bundle.v1.VerificationMaterial.public_key:type_name -> dev.sigstore.common.v1.PublicKeyIdentifier
	5, // 2: dev.sigstore.bundle.v1.VerificationMaterial.x509_certificate_chain:type_name -> dev.sigstore.common.v1.X509CertificateChain
	6, // 3: dev.sigstore.bundle.v1.VerificationMaterial.certificate:type_name -> dev.sigstore.common.v1.X509Certificate
	7, // 4: dev.sigstore.bundle.v1.VerificationMaterial.tlog_entries:type_name -> dev.sigstore.rekor.v1.TransparencyLogEntry
	0, // 5: dev.sigstore.bundle.v1.VerificationMaterial.timestamp_verification_data:type_name -> dev.sigstore.bundle.v1.TimestampVerificationData
	1, // 6: dev.sigstore.bundle.v1.Bundle.verification_material:type_name -> dev.sigstore.bundle.v1.VerificationMaterial
	8, // 7: dev.sigstore.bundle.v1.Bundle.message_signature:type_name -> dev.sigstore.common.v1.MessageSignature
	9, // 8: dev.sigstore.bundle.v1.Bundle.dsse_envelope:type_name -> io.intoto.Envelope
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_sigstore_bundle_proto_init() }
func file_sigstore_bundle_proto_init() {
	if File_sigstore_bundle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sigstore_bundle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimestampVerificationData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sigstore_bundle_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerificationMaterial); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sigstore_bundle_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bundle); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_sigstore_bundle_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*VerificationMaterial_PublicKey)(nil),
		(*VerificationMaterial_X509CertificateChain)(nil),
		(*VerificationMaterial_Certificate)(nil),
	}
	file_sigstore_bundle_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Bundle_MessageSignature)(nil),
		(*Bundle_DsseEnvelope)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sigstore_bundle_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sigstore_bundle_proto_goTypes,
		DependencyIndexes: file_sigstore_bundle_proto_depIdxs,
		MessageInfos:      file_sigstore_bundle_proto_msgTypes,
	}.Build()
	File_sigstore_bundle_proto = out.File
	file_sigstore_bundle_proto_rawDesc = nil
	file_sigstore_bundle_proto_goTypes = nil
	file_sigstore_bundle_proto_depIdxs = nil
}
