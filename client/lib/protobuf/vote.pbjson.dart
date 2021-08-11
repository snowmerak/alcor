///
//  Generated code. Do not modify.
//  source: protobuf/vote.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields,deprecated_member_use_from_same_package

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use keyBytesDescriptor instead')
const KeyBytes$json = const {
  '1': 'KeyBytes',
  '2': const [
    const {'1': 'Data', '3': 1, '4': 1, '5': 12, '10': 'Data'},
  ],
};

/// Descriptor for `KeyBytes`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List keyBytesDescriptor = $convert.base64Decode('CghLZXlCeXRlcxISCgREYXRhGAEgASgMUgREYXRh');
@$core.Deprecated('Use paperDescriptor instead')
const Paper$json = const {
  '1': 'Paper',
  '2': const [
    const {'1': 'VoterId', '3': 2, '4': 1, '5': 12, '10': 'VoterId'},
    const {'1': 'Timestamp', '3': 3, '4': 1, '5': 12, '10': 'Timestamp'},
    const {'1': 'Message', '3': 4, '4': 1, '5': 12, '10': 'Message'},
    const {'1': 'RandomBytes', '3': 5, '4': 1, '5': 12, '10': 'RandomBytes'},
    const {'1': 'Hash', '3': 6, '4': 1, '5': 12, '10': 'Hash'},
    const {'1': 'Signature', '3': 7, '4': 1, '5': 12, '10': 'Signature'},
  ],
};

/// Descriptor for `Paper`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List paperDescriptor = $convert.base64Decode('CgVQYXBlchIYCgdWb3RlcklkGAIgASgMUgdWb3RlcklkEhwKCVRpbWVzdGFtcBgDIAEoDFIJVGltZXN0YW1wEhgKB01lc3NhZ2UYBCABKAxSB01lc3NhZ2USIAoLUmFuZG9tQnl0ZXMYBSABKAxSC1JhbmRvbUJ5dGVzEhIKBEhhc2gYBiABKAxSBEhhc2gSHAoJU2lnbmF0dXJlGAcgASgMUglTaWduYXR1cmU=');
@$core.Deprecated('Use resultDescriptor instead')
const Result$json = const {
  '1': 'Result',
  '2': const [
    const {'1': 'Ok', '3': 1, '4': 1, '5': 8, '10': 'Ok'},
    const {'1': 'Error', '3': 2, '4': 1, '5': 9, '10': 'Error'},
  ],
};

/// Descriptor for `Result`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List resultDescriptor = $convert.base64Decode('CgZSZXN1bHQSDgoCT2sYASABKAhSAk9rEhQKBUVycm9yGAIgASgJUgVFcnJvcg==');
