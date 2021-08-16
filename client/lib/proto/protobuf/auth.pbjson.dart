///
//  Generated code. Do not modify.
//  source: proto/protobuf/auth.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields,deprecated_member_use_from_same_package

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use privateKeyDescriptor instead')
const PrivateKey$json = const {
  '1': 'PrivateKey',
  '2': const [
    const {'1': 'X', '3': 1, '4': 1, '5': 9, '10': 'X'},
    const {'1': 'Y', '3': 2, '4': 1, '5': 9, '10': 'Y'},
  ],
};

/// Descriptor for `PrivateKey`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List privateKeyDescriptor = $convert.base64Decode('CgpQcml2YXRlS2V5EgwKAVgYASABKAlSAVgSDAoBWRgCIAEoCVIBWQ==');
@$core.Deprecated('Use publicKeyDescriptor instead')
const PublicKey$json = const {
  '1': 'PublicKey',
  '2': const [
    const {'1': 'X', '3': 1, '4': 1, '5': 9, '10': 'X'},
    const {'1': 'Y', '3': 2, '4': 1, '5': 9, '10': 'Y'},
  ],
};

/// Descriptor for `PublicKey`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List publicKeyDescriptor = $convert.base64Decode('CglQdWJsaWNLZXkSDAoBWBgBIAEoCVIBWBIMCgFZGAIgASgJUgFZ');
@$core.Deprecated('Use signatureDescriptor instead')
const Signature$json = const {
  '1': 'Signature',
  '2': const [
    const {'1': 'R', '3': 1, '4': 1, '5': 9, '10': 'R'},
    const {'1': 'S', '3': 2, '4': 1, '5': 9, '10': 'S'},
  ],
};

/// Descriptor for `Signature`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List signatureDescriptor = $convert.base64Decode('CglTaWduYXR1cmUSDAoBUhgBIAEoCVIBUhIMCgFTGAIgASgJUgFT');
