import 'dart:typed_data';

import 'package:alcor/controllers/config.dart';
import 'package:alcor/controllers/register.dart';
import 'package:alcor/pages/candidate.dart';
import 'package:alcor/proto/protobuf/register.pb.dart';
import 'package:alcor/proto/protobuf/auth.pb.dart' as auth;
import 'package:alcor/service/capsulate.dart';
import 'package:elliptic/elliptic.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:hive/hive.dart';
import 'package:http/http.dart';

Future<void> registerCall(BuildContext context) async {
  var configController = Get.find<ConfigController>();
  var registerController = Get.find<RegisterController>();
  var phone = Phone();
  phone.name = registerController.name;
  phone.number = registerController.phoneNumber;
  phone.gender = false;
  if (registerController.gender[0]) {
    phone.gender = true;
  }
  phone.rRN = registerController.birth;
  var capsule = await encapsulate(phone.writeToBuffer());
  post(
          Uri.parse(
              'http://${configController.domain}:9999/voter/register/request'),
          body: capsule.writeToBuffer())
      .then((value) async {
    if (value.statusCode == 200) {
      await Get.defaultDialog(
          title: '인증 번호를 요청하였습니다', middleText: '인증 번호를 입력해주세요');
    } else {
      await Get.defaultDialog(
          title: '인증 번호를 요청에 실패하였습니다',
          middleText: '${value.statusCode}: 입력 정보를 확인해주세요');
    }
  }).onError((error, stackTrace) async {
    await Get.defaultDialog(
        title: '알 수 없는 에러가 발생했습니다', middleText: error.toString());
  });
}

Future<void> registerCert(BuildContext context) async {
  var configController = Get.find<ConfigController>();
  var registerController = Get.find<RegisterController>();
  var privateKey = getP256().generatePrivateKey();
  var publicKey = privateKey.publicKey;
  var cert = CertCode();
  cert.code = registerController.certCode.codeUnits;
  cert.publicKey = auth.PublicKey(
          x: publicKey.X.toRadixString(16), y: publicKey.Y.toRadixString(16))
      .writeToBuffer();
  encapsulate(cert.writeToBuffer()).then((capsule) {
    post(
            Uri.parse(
                'http://${configController.domain}:9999/voter/register/cert'),
            body: capsule.writeToBuffer())
        .then((value) async {
      var result = Result.fromBuffer(await decapsulate(value.bodyBytes));
      if (value.statusCode == 200) {
        await Hive.box<Uint8List>(configController.domain)
            .put('id', Uint8List.fromList(result.iD));
        await Hive.box<Uint8List>(configController.domain)
            .put('private_key', Uint8List.fromList(privateKey.bytes));
        await Get.defaultDialog(
            title: '인증에 성공하였습니다', middleText: '다음으로 넘어가 투표를 진행합니다.');
        Get.toNamed(candidateRoute);
      } else {
        await Get.defaultDialog(
            title: '인증에 실패하였습니다',
            middleText: '${value.statusCode}: ${result.error}');
      }
    });
  }).onError((error, stackTrace) async => await Get.defaultDialog(
      title: '알 수 없는 에러가 발생했습니다', middleText: error.toString()));
}
