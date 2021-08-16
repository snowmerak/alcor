///
//  Generated code. Do not modify.
//  source: proto/protobuf/stats.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class Data extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Data', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stats'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Candidate', protoName: 'Candidate')
    ..aOB(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Gender', protoName: 'Gender')
    ..a<$core.int>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Age', $pb.PbFieldType.O3, protoName: 'Age')
    ..a<$core.int>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Region', $pb.PbFieldType.O3, protoName: 'Region')
    ..a<$core.int>(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Job', $pb.PbFieldType.O3, protoName: 'Job')
    ..a<$core.int>(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Education', $pb.PbFieldType.O3, protoName: 'Education')
    ..aOB(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Married', protoName: 'Married')
    ..aOB(8, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Divorced', protoName: 'Divorced')
    ..aOB(9, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'HasCar', protoName: 'HasCar')
    ..a<$core.int>(10, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'HouseType', $pb.PbFieldType.O3, protoName: 'HouseType')
    ..a<$core.int>(11, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Salary', $pb.PbFieldType.O3, protoName: 'Salary')
    ..aOB(12, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'HasDebt', protoName: 'HasDebt')
    ..a<$core.int>(13, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Ideology', $pb.PbFieldType.O3, protoName: 'Ideology')
    ..a<$core.List<$core.int>>(14, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'VoterID', $pb.PbFieldType.OY, protoName: 'VoterID')
    ..hasRequiredFields = false
  ;

  Data._() : super();
  factory Data({
    $core.String? candidate,
    $core.bool? gender,
    $core.int? age,
    $core.int? region,
    $core.int? job,
    $core.int? education,
    $core.bool? married,
    $core.bool? divorced,
    $core.bool? hasCar,
    $core.int? houseType,
    $core.int? salary,
    $core.bool? hasDebt,
    $core.int? ideology,
    $core.List<$core.int>? voterID,
  }) {
    final _result = create();
    if (candidate != null) {
      _result.candidate = candidate;
    }
    if (gender != null) {
      _result.gender = gender;
    }
    if (age != null) {
      _result.age = age;
    }
    if (region != null) {
      _result.region = region;
    }
    if (job != null) {
      _result.job = job;
    }
    if (education != null) {
      _result.education = education;
    }
    if (married != null) {
      _result.married = married;
    }
    if (divorced != null) {
      _result.divorced = divorced;
    }
    if (hasCar != null) {
      _result.hasCar = hasCar;
    }
    if (houseType != null) {
      _result.houseType = houseType;
    }
    if (salary != null) {
      _result.salary = salary;
    }
    if (hasDebt != null) {
      _result.hasDebt = hasDebt;
    }
    if (ideology != null) {
      _result.ideology = ideology;
    }
    if (voterID != null) {
      _result.voterID = voterID;
    }
    return _result;
  }
  factory Data.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Data.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Data clone() => Data()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Data copyWith(void Function(Data) updates) => super.copyWith((message) => updates(message as Data)) as Data; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Data create() => Data._();
  Data createEmptyInstance() => create();
  static $pb.PbList<Data> createRepeated() => $pb.PbList<Data>();
  @$core.pragma('dart2js:noInline')
  static Data getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Data>(create);
  static Data? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get candidate => $_getSZ(0);
  @$pb.TagNumber(1)
  set candidate($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasCandidate() => $_has(0);
  @$pb.TagNumber(1)
  void clearCandidate() => clearField(1);

  @$pb.TagNumber(2)
  $core.bool get gender => $_getBF(1);
  @$pb.TagNumber(2)
  set gender($core.bool v) { $_setBool(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasGender() => $_has(1);
  @$pb.TagNumber(2)
  void clearGender() => clearField(2);

  @$pb.TagNumber(3)
  $core.int get age => $_getIZ(2);
  @$pb.TagNumber(3)
  set age($core.int v) { $_setSignedInt32(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasAge() => $_has(2);
  @$pb.TagNumber(3)
  void clearAge() => clearField(3);

  @$pb.TagNumber(4)
  $core.int get region => $_getIZ(3);
  @$pb.TagNumber(4)
  set region($core.int v) { $_setSignedInt32(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasRegion() => $_has(3);
  @$pb.TagNumber(4)
  void clearRegion() => clearField(4);

  @$pb.TagNumber(5)
  $core.int get job => $_getIZ(4);
  @$pb.TagNumber(5)
  set job($core.int v) { $_setSignedInt32(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasJob() => $_has(4);
  @$pb.TagNumber(5)
  void clearJob() => clearField(5);

  @$pb.TagNumber(6)
  $core.int get education => $_getIZ(5);
  @$pb.TagNumber(6)
  set education($core.int v) { $_setSignedInt32(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasEducation() => $_has(5);
  @$pb.TagNumber(6)
  void clearEducation() => clearField(6);

  @$pb.TagNumber(7)
  $core.bool get married => $_getBF(6);
  @$pb.TagNumber(7)
  set married($core.bool v) { $_setBool(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasMarried() => $_has(6);
  @$pb.TagNumber(7)
  void clearMarried() => clearField(7);

  @$pb.TagNumber(8)
  $core.bool get divorced => $_getBF(7);
  @$pb.TagNumber(8)
  set divorced($core.bool v) { $_setBool(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasDivorced() => $_has(7);
  @$pb.TagNumber(8)
  void clearDivorced() => clearField(8);

  @$pb.TagNumber(9)
  $core.bool get hasCar => $_getBF(8);
  @$pb.TagNumber(9)
  set hasCar($core.bool v) { $_setBool(8, v); }
  @$pb.TagNumber(9)
  $core.bool hasHasCar() => $_has(8);
  @$pb.TagNumber(9)
  void clearHasCar() => clearField(9);

  @$pb.TagNumber(10)
  $core.int get houseType => $_getIZ(9);
  @$pb.TagNumber(10)
  set houseType($core.int v) { $_setSignedInt32(9, v); }
  @$pb.TagNumber(10)
  $core.bool hasHouseType() => $_has(9);
  @$pb.TagNumber(10)
  void clearHouseType() => clearField(10);

  @$pb.TagNumber(11)
  $core.int get salary => $_getIZ(10);
  @$pb.TagNumber(11)
  set salary($core.int v) { $_setSignedInt32(10, v); }
  @$pb.TagNumber(11)
  $core.bool hasSalary() => $_has(10);
  @$pb.TagNumber(11)
  void clearSalary() => clearField(11);

  @$pb.TagNumber(12)
  $core.bool get hasDebt => $_getBF(11);
  @$pb.TagNumber(12)
  set hasDebt($core.bool v) { $_setBool(11, v); }
  @$pb.TagNumber(12)
  $core.bool hasHasDebt() => $_has(11);
  @$pb.TagNumber(12)
  void clearHasDebt() => clearField(12);

  @$pb.TagNumber(13)
  $core.int get ideology => $_getIZ(12);
  @$pb.TagNumber(13)
  set ideology($core.int v) { $_setSignedInt32(12, v); }
  @$pb.TagNumber(13)
  $core.bool hasIdeology() => $_has(12);
  @$pb.TagNumber(13)
  void clearIdeology() => clearField(13);

  @$pb.TagNumber(14)
  $core.List<$core.int> get voterID => $_getN(13);
  @$pb.TagNumber(14)
  set voterID($core.List<$core.int> v) { $_setBytes(13, v); }
  @$pb.TagNumber(14)
  $core.bool hasVoterID() => $_has(13);
  @$pb.TagNumber(14)
  void clearVoterID() => clearField(14);
}

