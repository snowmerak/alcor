///
//  Generated code. Do not modify.
//  source: proto/protobuf/info.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields,deprecated_member_use_from_same_package

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use voterDescriptor instead')
const Voter$json = const {
  '1': 'Voter',
  '2': const [
    const {'1': 'HashID', '3': 1, '4': 1, '5': 12, '10': 'HashID'},
    const {'1': 'PublicKey', '3': 2, '4': 1, '5': 12, '10': 'PublicKey'},
    const {'1': 'RandomBytes', '3': 3, '4': 1, '5': 12, '10': 'RandomBytes'},
    const {'1': 'Timestamp', '3': 4, '4': 1, '5': 12, '10': 'Timestamp'},
    const {'1': 'Voted', '3': 5, '4': 1, '5': 8, '10': 'Voted'},
    const {'1': 'IsCandidate', '3': 6, '4': 1, '5': 8, '10': 'IsCandidate'},
    const {'1': 'RepliedSurvey', '3': 7, '4': 1, '5': 8, '10': 'RepliedSurvey'},
  ],
};

/// Descriptor for `Voter`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List voterDescriptor = $convert.base64Decode('CgVWb3RlchIWCgZIYXNoSUQYASABKAxSBkhhc2hJRBIcCglQdWJsaWNLZXkYAiABKAxSCVB1YmxpY0tleRIgCgtSYW5kb21CeXRlcxgDIAEoDFILUmFuZG9tQnl0ZXMSHAoJVGltZXN0YW1wGAQgASgMUglUaW1lc3RhbXASFAoFVm90ZWQYBSABKAhSBVZvdGVkEiAKC0lzQ2FuZGlkYXRlGAYgASgIUgtJc0NhbmRpZGF0ZRIkCg1SZXBsaWVkU3VydmV5GAcgASgIUg1SZXBsaWVkU3VydmV5');
@$core.Deprecated('Use paperDescriptor instead')
const Paper$json = const {
  '1': 'Paper',
  '2': const [
    const {'1': 'VoterID', '3': 1, '4': 1, '5': 12, '10': 'VoterID'},
    const {'1': 'Timestamp', '3': 2, '4': 1, '5': 12, '10': 'Timestamp'},
    const {'1': 'Message', '3': 3, '4': 1, '5': 12, '10': 'Message'},
    const {'1': 'RandomBytes', '3': 4, '4': 1, '5': 12, '10': 'RandomBytes'},
    const {'1': 'Hash', '3': 5, '4': 1, '5': 12, '10': 'Hash'},
    const {'1': 'Signature', '3': 6, '4': 1, '5': 12, '10': 'Signature'},
  ],
};

/// Descriptor for `Paper`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List paperDescriptor = $convert.base64Decode('CgVQYXBlchIYCgdWb3RlcklEGAEgASgMUgdWb3RlcklEEhwKCVRpbWVzdGFtcBgCIAEoDFIJVGltZXN0YW1wEhgKB01lc3NhZ2UYAyABKAxSB01lc3NhZ2USIAoLUmFuZG9tQnl0ZXMYBCABKAxSC1JhbmRvbUJ5dGVzEhIKBEhhc2gYBSABKAxSBEhhc2gSHAoJU2lnbmF0dXJlGAYgASgMUglTaWduYXR1cmU=');
@$core.Deprecated('Use bundleDescriptor instead')
const Bundle$json = const {
  '1': 'Bundle',
  '2': const [
    const {'1': 'Hash', '3': 1, '4': 1, '5': 12, '10': 'Hash'},
    const {'1': 'Prev', '3': 2, '4': 1, '5': 12, '10': 'Prev'},
    const {'1': 'SubHashes', '3': 3, '4': 3, '5': 12, '10': 'SubHashes'},
  ],
};

/// Descriptor for `Bundle`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List bundleDescriptor = $convert.base64Decode('CgZCdW5kbGUSEgoESGFzaBgBIAEoDFIESGFzaBISCgRQcmV2GAIgASgMUgRQcmV2EhwKCVN1Ykhhc2hlcxgDIAMoDFIJU3ViSGFzaGVz');
