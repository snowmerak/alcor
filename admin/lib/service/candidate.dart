import 'dart:convert';
import 'dart:typed_data';

import 'package:admin/controller/candidate.dart';
import 'package:admin/controller/index.dart';
import 'package:admin/proto/protobuf/candidate.pbserver.dart';
import 'package:admin/proto/protobuf/capsule.pb.dart';
import 'package:get/get.dart';
import 'package:http/http.dart';

Future<void> getCandidates() async {
  var indexController = Get.find<IndexController>();
  var candidateController = Get.find<CandidateListController>();
  get(Uri.parse('http://${indexController.domain}:9999/candidate'))
      .then((value) async {
    if (value.statusCode == 200) {
      var list = List_.fromBuffer(value.bodyBytes);
      candidateController.list = list.names;
      candidateController.update();
    } else {
      await Get.defaultDialog(
          title: '후보 목록을 불러올 수 없습니다.', middleText: value.statusCode.toString());
    }
  }).onError((error, stackTrace) async => await Get.defaultDialog(
          title: '후보 목록을 불러올 수 없습니다.', middleText: error.toString()));
}

Future<void> addCandidate() async {
  var indexController = Get.find<IndexController>();
  var candidateController = Get.find<CandidateController>();
  var candidate = Candidate();
  candidate.name = candidateController.name;
  candidate.markdown = candidateController.markdown;
  candidate.password = indexController.password;

  var box = await indexController.aead.encrypt(candidate.writeToBuffer(),
      secretKey: indexController.secretKey, nonce: indexController.nonce);
  var capsule = Capsule(
    iD: indexController.id,
    data: box.cipherText + box.mac.bytes,
  );

  post(Uri.parse('http://${indexController.domain}:9999/candidate'),
          body: capsule.writeToBuffer())
      .then((value) async {
    if (value.statusCode == 200) {
      await Get.defaultDialog(title: '등록하였습니다.');
      Get.back();
    } else {
      await Get.defaultDialog(
          title: '실패하였습니다.', middleText: value.statusCode.toString());
    }
  }).onError((error, stackTrace) async {
    await Get.defaultDialog(title: '통신 에러', middleText: error.toString());
  });
}

Future<void> getCandidate() async {
  var indexController = Get.find<IndexController>();
  var candidateDetailController = Get.find<CandidateDetailController>();

  var name =
      base64Url.encode(utf8.encoder.convert(candidateDetailController.name));

  get(Uri.parse('http://${indexController.domain}:9999/candidate/$name'))
      .then((value) async {
    if (value.statusCode == 200) {
      var candidate = Candidate.fromBuffer(value.bodyBytes);
      candidateDetailController.updateMarkdown(candidate.markdown);
    } else {
      await Get.defaultDialog(
          title: '후보를 불러올 수 없습니다.', middleText: value.statusCode.toString());
    }
  }).onError((error, stackTrace) async => await Get.defaultDialog(
          title: '후보를 불러올 수 없습니다.', middleText: error.toString()));
}

Future<void> removeCandidate() async {
  var indexController = Get.find<IndexController>();
  var candidateController = Get.find<CandidateListController>();
  var candidate = Candidate();
  candidate.name = candidateController.selected;
  candidate.password = indexController.password;

  var box = await indexController.aead.encrypt(candidate.writeToBuffer(),
      secretKey: indexController.secretKey, nonce: indexController.nonce);
  var capsule = Capsule(
    iD: indexController.id,
    data: box.cipherText + box.mac.bytes,
  );

  delete(Uri.parse('http://${indexController.domain}:9999/candidate'),
          body: capsule.writeToBuffer())
      .then((value) async {
    if (value.statusCode == 200) {
      await Get.defaultDialog(title: '삭제하였습니다.');
      getCandidates();
      Get.back();
    } else {
      await Get.defaultDialog(
          title: '실패하였습니다.', middleText: value.statusCode.toString());
    }
  }).onError((error, stackTrace) async {
    await Get.defaultDialog(title: '통신 에러', middleText: error.toString());
  });
}

Future<void> editCandidate() async {
  var indexController = Get.find<IndexController>();
  var candidateController = Get.find<CandidateDetailController>();
  var candidate = Candidate();
  candidate.name = candidateController.name;
  candidate.markdown = candidateController.markdown;
  candidate.password = indexController.password;

  var box = await indexController.aead.encrypt(candidate.writeToBuffer(),
      secretKey: indexController.secretKey, nonce: indexController.nonce);
  var capsule = Capsule(
    iD: indexController.id,
    data: box.cipherText + box.mac.bytes,
  );

  patch(Uri.parse('http://${indexController.domain}:9999/candidate'),
          body: capsule.writeToBuffer())
      .then((value) async {
    if (value.statusCode == 200) {
      await Get.defaultDialog(title: '수정하였습니다.');
      getCandidates();
      Get.back();
    } else {
      await Get.defaultDialog(
          title: '실패하였습니다.', middleText: value.statusCode.toString());
    }
  }).onError((error, stackTrace) async {
    await Get.defaultDialog(title: '통신 에러', middleText: error.toString());
  });
}
