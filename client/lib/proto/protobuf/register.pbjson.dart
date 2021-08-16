///
//  Generated code. Do not modify.
//  source: proto/protobuf/register.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields,deprecated_member_use_from_same_package

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use phoneDescriptor instead')
const Phone$json = const {
  '1': 'Phone',
  '2': const [
    const {'1': 'Number', '3': 2, '4': 1, '5': 9, '10': 'Number'},
    const {'1': 'Name', '3': 3, '4': 1, '5': 9, '10': 'Name'},
    const {'1': 'RRN', '3': 4, '4': 1, '5': 9, '10': 'RRN'},
    const {'1': 'Gender', '3': 1, '4': 1, '5': 8, '10': 'Gender'},
  ],
};

/// Descriptor for `Phone`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List phoneDescriptor = $convert.base64Decode('CgVQaG9uZRIWCgZOdW1iZXIYAiABKAlSBk51bWJlchISCgROYW1lGAMgASgJUgROYW1lEhAKA1JSThgEIAEoCVIDUlJOEhYKBkdlbmRlchgBIAEoCFIGR2VuZGVy');
@$core.Deprecated('Use cacheDescriptor instead')
const Cache$json = const {
  '1': 'Cache',
  '2': const [
    const {'1': 'Number', '3': 2, '4': 1, '5': 9, '10': 'Number'},
    const {'1': 'Name', '3': 3, '4': 1, '5': 9, '10': 'Name'},
    const {'1': 'RRN', '3': 4, '4': 1, '5': 9, '10': 'RRN'},
    const {'1': 'Gender', '3': 1, '4': 1, '5': 8, '10': 'Gender'},
    const {'1': 'Code', '3': 5, '4': 1, '5': 12, '10': 'Code'},
  ],
};

/// Descriptor for `Cache`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List cacheDescriptor = $convert.base64Decode('CgVDYWNoZRIWCgZOdW1iZXIYAiABKAlSBk51bWJlchISCgROYW1lGAMgASgJUgROYW1lEhAKA1JSThgEIAEoCVIDUlJOEhYKBkdlbmRlchgBIAEoCFIGR2VuZGVyEhIKBENvZGUYBSABKAxSBENvZGU=');
@$core.Deprecated('Use certCodeDescriptor instead')
const CertCode$json = const {
  '1': 'CertCode',
  '2': const [
    const {'1': 'Code', '3': 1, '4': 1, '5': 12, '10': 'Code'},
    const {'1': 'PublicKey', '3': 2, '4': 1, '5': 12, '10': 'PublicKey'},
  ],
};

/// Descriptor for `CertCode`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List certCodeDescriptor = $convert.base64Decode('CghDZXJ0Q29kZRISCgRDb2RlGAEgASgMUgRDb2RlEhwKCVB1YmxpY0tleRgCIAEoDFIJUHVibGljS2V5');
@$core.Deprecated('Use resultDescriptor instead')
const Result$json = const {
  '1': 'Result',
  '2': const [
    const {'1': 'ID', '3': 1, '4': 1, '5': 12, '10': 'ID'},
    const {'1': 'Error', '3': 2, '4': 1, '5': 9, '10': 'Error'},
  ],
};

/// Descriptor for `Result`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List resultDescriptor = $convert.base64Decode('CgZSZXN1bHQSDgoCSUQYASABKAxSAklEEhQKBUVycm9yGAIgASgJUgVFcnJvcg==');
@$core.Deprecated('Use accountDescriptor instead')
const Account$json = const {
  '1': 'Account',
  '2': const [
    const {'1': 'HashID', '3': 1, '4': 1, '5': 12, '10': 'HashID'},
    const {'1': 'PrivateKey', '3': 2, '4': 1, '5': 12, '10': 'PrivateKey'},
  ],
};

/// Descriptor for `Account`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List accountDescriptor = $convert.base64Decode('CgdBY2NvdW50EhYKBkhhc2hJRBgBIAEoDFIGSGFzaElEEh4KClByaXZhdGVLZXkYAiABKAxSClByaXZhdGVLZXk=');
