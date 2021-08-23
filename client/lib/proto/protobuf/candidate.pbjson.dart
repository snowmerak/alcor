///
//  Generated code. Do not modify.
//  source: proto/protobuf/candidate.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields,deprecated_member_use_from_same_package

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use requestDescriptor instead')
const Request$json = const {
  '1': 'Request',
  '2': const [
    const {'1': 'VoterID', '3': 1, '4': 1, '5': 12, '10': 'VoterID'},
    const {'1': 'Signature', '3': 2, '4': 1, '5': 12, '10': 'Signature'},
    const {'1': 'KeyBytes', '3': 3, '4': 1, '5': 12, '10': 'KeyBytes'},
    const {'1': 'Name', '3': 4, '4': 1, '5': 9, '10': 'Name'},
    const {'1': 'Markdown', '3': 5, '4': 1, '5': 9, '10': 'Markdown'},
    const {'1': 'Hash', '3': 6, '4': 1, '5': 12, '10': 'Hash'},
  ],
};

/// Descriptor for `Request`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List requestDescriptor = $convert.base64Decode('CgdSZXF1ZXN0EhgKB1ZvdGVySUQYASABKAxSB1ZvdGVySUQSHAoJU2lnbmF0dXJlGAIgASgMUglTaWduYXR1cmUSGgoIS2V5Qnl0ZXMYAyABKAxSCEtleUJ5dGVzEhIKBE5hbWUYBCABKAlSBE5hbWUSGgoITWFya2Rvd24YBSABKAlSCE1hcmtkb3duEhIKBEhhc2gYBiABKAxSBEhhc2g=');
@$core.Deprecated('Use resultDescriptor instead')
const Result$json = const {
  '1': 'Result',
  '2': const [
    const {'1': 'OK', '3': 1, '4': 1, '5': 8, '10': 'OK'},
    const {'1': 'Error', '3': 2, '4': 1, '5': 9, '10': 'Error'},
  ],
};

/// Descriptor for `Result`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List resultDescriptor = $convert.base64Decode('CgZSZXN1bHQSDgoCT0sYASABKAhSAk9LEhQKBUVycm9yGAIgASgJUgVFcnJvcg==');
@$core.Deprecated('Use candidateDescriptor instead')
const Candidate$json = const {
  '1': 'Candidate',
  '2': const [
    const {'1': 'Name', '3': 1, '4': 1, '5': 9, '10': 'Name'},
    const {'1': 'Markdown', '3': 2, '4': 1, '5': 9, '10': 'Markdown'},
    const {'1': 'Password', '3': 3, '4': 1, '5': 9, '10': 'Password'},
  ],
};

/// Descriptor for `Candidate`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List candidateDescriptor = $convert.base64Decode('CglDYW5kaWRhdGUSEgoETmFtZRgBIAEoCVIETmFtZRIaCghNYXJrZG93bhgCIAEoCVIITWFya2Rvd24SGgoIUGFzc3dvcmQYAyABKAlSCFBhc3N3b3Jk');
@$core.Deprecated('Use list_Descriptor instead')
const List_$json = const {
  '1': 'List',
  '2': const [
    const {'1': 'Error', '3': 1, '4': 1, '5': 9, '10': 'Error'},
    const {'1': 'Names', '3': 2, '4': 3, '5': 9, '10': 'Names'},
  ],
};

/// Descriptor for `List`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List list_Descriptor = $convert.base64Decode('CgRMaXN0EhQKBUVycm9yGAEgASgJUgVFcnJvchIUCgVOYW1lcxgCIAMoCVIFTmFtZXM=');
