///
//  Generated code. Do not modify.
//  source: proto/protobuf/register.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class Phone extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Phone', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'register'), createEmptyInstance: create)
    ..aOB(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Gender', protoName: 'Gender')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Number', protoName: 'Number')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Name', protoName: 'Name')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'RRN', protoName: 'RRN')
    ..hasRequiredFields = false
  ;

  Phone._() : super();
  factory Phone({
    $core.bool? gender,
    $core.String? number,
    $core.String? name,
    $core.String? rRN,
  }) {
    final _result = create();
    if (gender != null) {
      _result.gender = gender;
    }
    if (number != null) {
      _result.number = number;
    }
    if (name != null) {
      _result.name = name;
    }
    if (rRN != null) {
      _result.rRN = rRN;
    }
    return _result;
  }
  factory Phone.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Phone.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Phone clone() => Phone()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Phone copyWith(void Function(Phone) updates) => super.copyWith((message) => updates(message as Phone)) as Phone; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Phone create() => Phone._();
  Phone createEmptyInstance() => create();
  static $pb.PbList<Phone> createRepeated() => $pb.PbList<Phone>();
  @$core.pragma('dart2js:noInline')
  static Phone getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Phone>(create);
  static Phone? _defaultInstance;

  @$pb.TagNumber(1)
  $core.bool get gender => $_getBF(0);
  @$pb.TagNumber(1)
  set gender($core.bool v) { $_setBool(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasGender() => $_has(0);
  @$pb.TagNumber(1)
  void clearGender() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get number => $_getSZ(1);
  @$pb.TagNumber(2)
  set number($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasNumber() => $_has(1);
  @$pb.TagNumber(2)
  void clearNumber() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get name => $_getSZ(2);
  @$pb.TagNumber(3)
  set name($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasName() => $_has(2);
  @$pb.TagNumber(3)
  void clearName() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get rRN => $_getSZ(3);
  @$pb.TagNumber(4)
  set rRN($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasRRN() => $_has(3);
  @$pb.TagNumber(4)
  void clearRRN() => clearField(4);
}

class Cache extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Cache', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'register'), createEmptyInstance: create)
    ..aOB(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Gender', protoName: 'Gender')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Number', protoName: 'Number')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Name', protoName: 'Name')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'RRN', protoName: 'RRN')
    ..a<$core.List<$core.int>>(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Code', $pb.PbFieldType.OY, protoName: 'Code')
    ..hasRequiredFields = false
  ;

  Cache._() : super();
  factory Cache({
    $core.bool? gender,
    $core.String? number,
    $core.String? name,
    $core.String? rRN,
    $core.List<$core.int>? code,
  }) {
    final _result = create();
    if (gender != null) {
      _result.gender = gender;
    }
    if (number != null) {
      _result.number = number;
    }
    if (name != null) {
      _result.name = name;
    }
    if (rRN != null) {
      _result.rRN = rRN;
    }
    if (code != null) {
      _result.code = code;
    }
    return _result;
  }
  factory Cache.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Cache.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Cache clone() => Cache()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Cache copyWith(void Function(Cache) updates) => super.copyWith((message) => updates(message as Cache)) as Cache; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Cache create() => Cache._();
  Cache createEmptyInstance() => create();
  static $pb.PbList<Cache> createRepeated() => $pb.PbList<Cache>();
  @$core.pragma('dart2js:noInline')
  static Cache getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Cache>(create);
  static Cache? _defaultInstance;

  @$pb.TagNumber(1)
  $core.bool get gender => $_getBF(0);
  @$pb.TagNumber(1)
  set gender($core.bool v) { $_setBool(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasGender() => $_has(0);
  @$pb.TagNumber(1)
  void clearGender() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get number => $_getSZ(1);
  @$pb.TagNumber(2)
  set number($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasNumber() => $_has(1);
  @$pb.TagNumber(2)
  void clearNumber() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get name => $_getSZ(2);
  @$pb.TagNumber(3)
  set name($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasName() => $_has(2);
  @$pb.TagNumber(3)
  void clearName() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get rRN => $_getSZ(3);
  @$pb.TagNumber(4)
  set rRN($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasRRN() => $_has(3);
  @$pb.TagNumber(4)
  void clearRRN() => clearField(4);

  @$pb.TagNumber(5)
  $core.List<$core.int> get code => $_getN(4);
  @$pb.TagNumber(5)
  set code($core.List<$core.int> v) { $_setBytes(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasCode() => $_has(4);
  @$pb.TagNumber(5)
  void clearCode() => clearField(5);
}

class CertCode extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'CertCode', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'register'), createEmptyInstance: create)
    ..a<$core.List<$core.int>>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Code', $pb.PbFieldType.OY, protoName: 'Code')
    ..a<$core.List<$core.int>>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'PublicKey', $pb.PbFieldType.OY, protoName: 'PublicKey')
    ..hasRequiredFields = false
  ;

  CertCode._() : super();
  factory CertCode({
    $core.List<$core.int>? code,
    $core.List<$core.int>? publicKey,
  }) {
    final _result = create();
    if (code != null) {
      _result.code = code;
    }
    if (publicKey != null) {
      _result.publicKey = publicKey;
    }
    return _result;
  }
  factory CertCode.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory CertCode.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  CertCode clone() => CertCode()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  CertCode copyWith(void Function(CertCode) updates) => super.copyWith((message) => updates(message as CertCode)) as CertCode; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static CertCode create() => CertCode._();
  CertCode createEmptyInstance() => create();
  static $pb.PbList<CertCode> createRepeated() => $pb.PbList<CertCode>();
  @$core.pragma('dart2js:noInline')
  static CertCode getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CertCode>(create);
  static CertCode? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$core.int> get code => $_getN(0);
  @$pb.TagNumber(1)
  set code($core.List<$core.int> v) { $_setBytes(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasCode() => $_has(0);
  @$pb.TagNumber(1)
  void clearCode() => clearField(1);

  @$pb.TagNumber(2)
  $core.List<$core.int> get publicKey => $_getN(1);
  @$pb.TagNumber(2)
  set publicKey($core.List<$core.int> v) { $_setBytes(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasPublicKey() => $_has(1);
  @$pb.TagNumber(2)
  void clearPublicKey() => clearField(2);
}

class Result extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Result', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'register'), createEmptyInstance: create)
    ..a<$core.List<$core.int>>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ID', $pb.PbFieldType.OY, protoName: 'ID')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Error', protoName: 'Error')
    ..hasRequiredFields = false
  ;

  Result._() : super();
  factory Result({
    $core.List<$core.int>? iD,
    $core.String? error,
  }) {
    final _result = create();
    if (iD != null) {
      _result.iD = iD;
    }
    if (error != null) {
      _result.error = error;
    }
    return _result;
  }
  factory Result.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Result.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Result clone() => Result()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Result copyWith(void Function(Result) updates) => super.copyWith((message) => updates(message as Result)) as Result; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Result create() => Result._();
  Result createEmptyInstance() => create();
  static $pb.PbList<Result> createRepeated() => $pb.PbList<Result>();
  @$core.pragma('dart2js:noInline')
  static Result getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Result>(create);
  static Result? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$core.int> get iD => $_getN(0);
  @$pb.TagNumber(1)
  set iD($core.List<$core.int> v) { $_setBytes(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasID() => $_has(0);
  @$pb.TagNumber(1)
  void clearID() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get error => $_getSZ(1);
  @$pb.TagNumber(2)
  set error($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasError() => $_has(1);
  @$pb.TagNumber(2)
  void clearError() => clearField(2);
}

class Account extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Account', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'register'), createEmptyInstance: create)
    ..a<$core.List<$core.int>>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'HashID', $pb.PbFieldType.OY, protoName: 'HashID')
    ..a<$core.List<$core.int>>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'PrivateKey', $pb.PbFieldType.OY, protoName: 'PrivateKey')
    ..hasRequiredFields = false
  ;

  Account._() : super();
  factory Account({
    $core.List<$core.int>? hashID,
    $core.List<$core.int>? privateKey,
  }) {
    final _result = create();
    if (hashID != null) {
      _result.hashID = hashID;
    }
    if (privateKey != null) {
      _result.privateKey = privateKey;
    }
    return _result;
  }
  factory Account.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Account.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Account clone() => Account()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Account copyWith(void Function(Account) updates) => super.copyWith((message) => updates(message as Account)) as Account; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Account create() => Account._();
  Account createEmptyInstance() => create();
  static $pb.PbList<Account> createRepeated() => $pb.PbList<Account>();
  @$core.pragma('dart2js:noInline')
  static Account getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Account>(create);
  static Account? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$core.int> get hashID => $_getN(0);
  @$pb.TagNumber(1)
  set hashID($core.List<$core.int> v) { $_setBytes(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasHashID() => $_has(0);
  @$pb.TagNumber(1)
  void clearHashID() => clearField(1);

  @$pb.TagNumber(2)
  $core.List<$core.int> get privateKey => $_getN(1);
  @$pb.TagNumber(2)
  set privateKey($core.List<$core.int> v) { $_setBytes(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasPrivateKey() => $_has(1);
  @$pb.TagNumber(2)
  void clearPrivateKey() => clearField(2);
}

