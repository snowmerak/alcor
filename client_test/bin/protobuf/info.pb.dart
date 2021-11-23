///
//  Generated code. Do not modify.
//  source: proto/protobuf/info.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class Voter extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Voter', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'info'), createEmptyInstance: create)
    ..a<$core.List<$core.int>>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'HashID', $pb.PbFieldType.OY, protoName: 'HashID')
    ..a<$core.List<$core.int>>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'PublicKey', $pb.PbFieldType.OY, protoName: 'PublicKey')
    ..a<$core.List<$core.int>>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'RandomBytes', $pb.PbFieldType.OY, protoName: 'RandomBytes')
    ..a<$core.List<$core.int>>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Timestamp', $pb.PbFieldType.OY, protoName: 'Timestamp')
    ..aOB(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Voted', protoName: 'Voted')
    ..aOB(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'IsCandidate', protoName: 'IsCandidate')
    ..aOB(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'RepliedSurvey', protoName: 'RepliedSurvey')
    ..hasRequiredFields = false
  ;

  Voter._() : super();
  factory Voter({
    $core.List<$core.int>? hashID,
    $core.List<$core.int>? publicKey,
    $core.List<$core.int>? randomBytes,
    $core.List<$core.int>? timestamp,
    $core.bool? voted,
    $core.bool? isCandidate,
    $core.bool? repliedSurvey,
  }) {
    final _result = create();
    if (hashID != null) {
      _result.hashID = hashID;
    }
    if (publicKey != null) {
      _result.publicKey = publicKey;
    }
    if (randomBytes != null) {
      _result.randomBytes = randomBytes;
    }
    if (timestamp != null) {
      _result.timestamp = timestamp;
    }
    if (voted != null) {
      _result.voted = voted;
    }
    if (isCandidate != null) {
      _result.isCandidate = isCandidate;
    }
    if (repliedSurvey != null) {
      _result.repliedSurvey = repliedSurvey;
    }
    return _result;
  }
  factory Voter.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Voter.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Voter clone() => Voter()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Voter copyWith(void Function(Voter) updates) => super.copyWith((message) => updates(message as Voter)) as Voter; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Voter create() => Voter._();
  Voter createEmptyInstance() => create();
  static $pb.PbList<Voter> createRepeated() => $pb.PbList<Voter>();
  @$core.pragma('dart2js:noInline')
  static Voter getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Voter>(create);
  static Voter? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$core.int> get hashID => $_getN(0);
  @$pb.TagNumber(1)
  set hashID($core.List<$core.int> v) { $_setBytes(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasHashID() => $_has(0);
  @$pb.TagNumber(1)
  void clearHashID() => clearField(1);

  @$pb.TagNumber(2)
  $core.List<$core.int> get publicKey => $_getN(1);
  @$pb.TagNumber(2)
  set publicKey($core.List<$core.int> v) { $_setBytes(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasPublicKey() => $_has(1);
  @$pb.TagNumber(2)
  void clearPublicKey() => clearField(2);

  @$pb.TagNumber(3)
  $core.List<$core.int> get randomBytes => $_getN(2);
  @$pb.TagNumber(3)
  set randomBytes($core.List<$core.int> v) { $_setBytes(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasRandomBytes() => $_has(2);
  @$pb.TagNumber(3)
  void clearRandomBytes() => clearField(3);

  @$pb.TagNumber(4)
  $core.List<$core.int> get timestamp => $_getN(3);
  @$pb.TagNumber(4)
  set timestamp($core.List<$core.int> v) { $_setBytes(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasTimestamp() => $_has(3);
  @$pb.TagNumber(4)
  void clearTimestamp() => clearField(4);

  @$pb.TagNumber(5)
  $core.bool get voted => $_getBF(4);
  @$pb.TagNumber(5)
  set voted($core.bool v) { $_setBool(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasVoted() => $_has(4);
  @$pb.TagNumber(5)
  void clearVoted() => clearField(5);

  @$pb.TagNumber(6)
  $core.bool get isCandidate => $_getBF(5);
  @$pb.TagNumber(6)
  set isCandidate($core.bool v) { $_setBool(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasIsCandidate() => $_has(5);
  @$pb.TagNumber(6)
  void clearIsCandidate() => clearField(6);

  @$pb.TagNumber(7)
  $core.bool get repliedSurvey => $_getBF(6);
  @$pb.TagNumber(7)
  set repliedSurvey($core.bool v) { $_setBool(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasRepliedSurvey() => $_has(6);
  @$pb.TagNumber(7)
  void clearRepliedSurvey() => clearField(7);
}

class Paper extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Paper', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'info'), createEmptyInstance: create)
    ..a<$core.List<$core.int>>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'VoterID', $pb.PbFieldType.OY, protoName: 'VoterID')
    ..a<$core.List<$core.int>>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Timestamp', $pb.PbFieldType.OY, protoName: 'Timestamp')
    ..a<$core.List<$core.int>>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Message', $pb.PbFieldType.OY, protoName: 'Message')
    ..a<$core.List<$core.int>>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'RandomBytes', $pb.PbFieldType.OY, protoName: 'RandomBytes')
    ..a<$core.List<$core.int>>(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Hash', $pb.PbFieldType.OY, protoName: 'Hash')
    ..a<$core.List<$core.int>>(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Signature', $pb.PbFieldType.OY, protoName: 'Signature')
    ..hasRequiredFields = false
  ;

  Paper._() : super();
  factory Paper({
    $core.List<$core.int>? voterID,
    $core.List<$core.int>? timestamp,
    $core.List<$core.int>? message,
    $core.List<$core.int>? randomBytes,
    $core.List<$core.int>? hash,
    $core.List<$core.int>? signature,
  }) {
    final _result = create();
    if (voterID != null) {
      _result.voterID = voterID;
    }
    if (timestamp != null) {
      _result.timestamp = timestamp;
    }
    if (message != null) {
      _result.message = message;
    }
    if (randomBytes != null) {
      _result.randomBytes = randomBytes;
    }
    if (hash != null) {
      _result.hash = hash;
    }
    if (signature != null) {
      _result.signature = signature;
    }
    return _result;
  }
  factory Paper.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Paper.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Paper clone() => Paper()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Paper copyWith(void Function(Paper) updates) => super.copyWith((message) => updates(message as Paper)) as Paper; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Paper create() => Paper._();
  Paper createEmptyInstance() => create();
  static $pb.PbList<Paper> createRepeated() => $pb.PbList<Paper>();
  @$core.pragma('dart2js:noInline')
  static Paper getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Paper>(create);
  static Paper? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$core.int> get voterID => $_getN(0);
  @$pb.TagNumber(1)
  set voterID($core.List<$core.int> v) { $_setBytes(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasVoterID() => $_has(0);
  @$pb.TagNumber(1)
  void clearVoterID() => clearField(1);

  @$pb.TagNumber(2)
  $core.List<$core.int> get timestamp => $_getN(1);
  @$pb.TagNumber(2)
  set timestamp($core.List<$core.int> v) { $_setBytes(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasTimestamp() => $_has(1);
  @$pb.TagNumber(2)
  void clearTimestamp() => clearField(2);

  @$pb.TagNumber(3)
  $core.List<$core.int> get message => $_getN(2);
  @$pb.TagNumber(3)
  set message($core.List<$core.int> v) { $_setBytes(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasMessage() => $_has(2);
  @$pb.TagNumber(3)
  void clearMessage() => clearField(3);

  @$pb.TagNumber(4)
  $core.List<$core.int> get randomBytes => $_getN(3);
  @$pb.TagNumber(4)
  set randomBytes($core.List<$core.int> v) { $_setBytes(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasRandomBytes() => $_has(3);
  @$pb.TagNumber(4)
  void clearRandomBytes() => clearField(4);

  @$pb.TagNumber(5)
  $core.List<$core.int> get hash => $_getN(4);
  @$pb.TagNumber(5)
  set hash($core.List<$core.int> v) { $_setBytes(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasHash() => $_has(4);
  @$pb.TagNumber(5)
  void clearHash() => clearField(5);

  @$pb.TagNumber(6)
  $core.List<$core.int> get signature => $_getN(5);
  @$pb.TagNumber(6)
  set signature($core.List<$core.int> v) { $_setBytes(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasSignature() => $_has(5);
  @$pb.TagNumber(6)
  void clearSignature() => clearField(6);
}

class Bundle extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Bundle', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'info'), createEmptyInstance: create)
    ..a<$core.List<$core.int>>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Hash', $pb.PbFieldType.OY, protoName: 'Hash')
    ..a<$core.List<$core.int>>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Prev', $pb.PbFieldType.OY, protoName: 'Prev')
    ..p<$core.List<$core.int>>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'SubHashes', $pb.PbFieldType.PY, protoName: 'SubHashes')
    ..hasRequiredFields = false
  ;

  Bundle._() : super();
  factory Bundle({
    $core.List<$core.int>? hash,
    $core.List<$core.int>? prev,
    $core.Iterable<$core.List<$core.int>>? subHashes,
  }) {
    final _result = create();
    if (hash != null) {
      _result.hash = hash;
    }
    if (prev != null) {
      _result.prev = prev;
    }
    if (subHashes != null) {
      _result.subHashes.addAll(subHashes);
    }
    return _result;
  }
  factory Bundle.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Bundle.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Bundle clone() => Bundle()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Bundle copyWith(void Function(Bundle) updates) => super.copyWith((message) => updates(message as Bundle)) as Bundle; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Bundle create() => Bundle._();
  Bundle createEmptyInstance() => create();
  static $pb.PbList<Bundle> createRepeated() => $pb.PbList<Bundle>();
  @$core.pragma('dart2js:noInline')
  static Bundle getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Bundle>(create);
  static Bundle? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$core.int> get hash => $_getN(0);
  @$pb.TagNumber(1)
  set hash($core.List<$core.int> v) { $_setBytes(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasHash() => $_has(0);
  @$pb.TagNumber(1)
  void clearHash() => clearField(1);

  @$pb.TagNumber(2)
  $core.List<$core.int> get prev => $_getN(1);
  @$pb.TagNumber(2)
  set prev($core.List<$core.int> v) { $_setBytes(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasPrev() => $_has(1);
  @$pb.TagNumber(2)
  void clearPrev() => clearField(2);

  @$pb.TagNumber(3)
  $core.List<$core.List<$core.int>> get subHashes => $_getList(2);
}

