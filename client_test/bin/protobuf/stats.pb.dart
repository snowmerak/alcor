///
//  Generated code. Do not modify.
//  source: proto/protobuf/stats.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

class Data extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Data', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stats'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Candidate', protoName: 'Candidate')
    ..aOB(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Gender', protoName: 'Gender')
    ..a<$core.int>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Age', $pb.PbFieldType.O3, protoName: 'Age')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Region', protoName: 'Region')
    ..aOS(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Job', protoName: 'Job')
    ..aOS(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Education', protoName: 'Education')
    ..aOB(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Married', protoName: 'Married')
    ..aOB(8, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Divorced', protoName: 'Divorced')
    ..aOB(9, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'HasCar', protoName: 'HasCar')
    ..aOS(10, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'HouseType', protoName: 'HouseType')
    ..a<$core.int>(11, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Salary', $pb.PbFieldType.O3, protoName: 'Salary')
    ..aOB(12, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'HasDebt', protoName: 'HasDebt')
    ..aOS(13, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Ideology', protoName: 'Ideology')
    ..a<$core.List<$core.int>>(14, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'VoterID', $pb.PbFieldType.OY, protoName: 'VoterID')
    ..aInt64(15, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Number', protoName: 'Number')
    ..hasRequiredFields = false
  ;

  Data._() : super();
  factory Data({
    $core.String? candidate,
    $core.bool? gender,
    $core.int? age,
    $core.String? region,
    $core.String? job,
    $core.String? education,
    $core.bool? married,
    $core.bool? divorced,
    $core.bool? hasCar,
    $core.String? houseType,
    $core.int? salary,
    $core.bool? hasDebt,
    $core.String? ideology,
    $core.List<$core.int>? voterID,
    $fixnum.Int64? number,
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
    if (number != null) {
      _result.number = number;
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
  $core.String get region => $_getSZ(3);
  @$pb.TagNumber(4)
  set region($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasRegion() => $_has(3);
  @$pb.TagNumber(4)
  void clearRegion() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get job => $_getSZ(4);
  @$pb.TagNumber(5)
  set job($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasJob() => $_has(4);
  @$pb.TagNumber(5)
  void clearJob() => clearField(5);

  @$pb.TagNumber(6)
  $core.String get education => $_getSZ(5);
  @$pb.TagNumber(6)
  set education($core.String v) { $_setString(5, v); }
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
  $core.String get houseType => $_getSZ(9);
  @$pb.TagNumber(10)
  set houseType($core.String v) { $_setString(9, v); }
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
  $core.String get ideology => $_getSZ(12);
  @$pb.TagNumber(13)
  set ideology($core.String v) { $_setString(12, v); }
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

  @$pb.TagNumber(15)
  $fixnum.Int64 get number => $_getI64(14);
  @$pb.TagNumber(15)
  set number($fixnum.Int64 v) { $_setInt64(14, v); }
  @$pb.TagNumber(15)
  $core.bool hasNumber() => $_has(14);
  @$pb.TagNumber(15)
  void clearNumber() => clearField(15);
}

class List_ extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'List', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stats'), createEmptyInstance: create)
    ..pPS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Data', protoName: 'Data')
    ..hasRequiredFields = false
  ;

  List_._() : super();
  factory List_({
    $core.Iterable<$core.String>? data,
  }) {
    final _result = create();
    if (data != null) {
      _result.data.addAll(data);
    }
    return _result;
  }
  factory List_.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory List_.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  List_ clone() => List_()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  List_ copyWith(void Function(List_) updates) => super.copyWith((message) => updates(message as List_)) as List_; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static List_ create() => List_._();
  List_ createEmptyInstance() => create();
  static $pb.PbList<List_> createRepeated() => $pb.PbList<List_>();
  @$core.pragma('dart2js:noInline')
  static List_ getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<List_>(create);
  static List_? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$core.String> get data => $_getList(0);
}

class NumberList extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'NumberList', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stats'), createEmptyInstance: create)
    ..p<$fixnum.Int64>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Data', $pb.PbFieldType.P6, protoName: 'Data')
    ..hasRequiredFields = false
  ;

  NumberList._() : super();
  factory NumberList({
    $core.Iterable<$fixnum.Int64>? data,
  }) {
    final _result = create();
    if (data != null) {
      _result.data.addAll(data);
    }
    return _result;
  }
  factory NumberList.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory NumberList.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  NumberList clone() => NumberList()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  NumberList copyWith(void Function(NumberList) updates) => super.copyWith((message) => updates(message as NumberList)) as NumberList; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static NumberList create() => NumberList._();
  NumberList createEmptyInstance() => create();
  static $pb.PbList<NumberList> createRepeated() => $pb.PbList<NumberList>();
  @$core.pragma('dart2js:noInline')
  static NumberList getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<NumberList>(create);
  static NumberList? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$fixnum.Int64> get data => $_getList(0);
}

