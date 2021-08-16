import 'dart:convert';
import 'dart:typed_data';

import 'package:alcor/controllers/candidate.dart';
import 'package:alcor/controllers/config.dart';
import 'package:alcor/controllers/receipt.dart';
import 'package:alcor/controllers/vote.dart';
import 'package:alcor/pages/receipt.dart';
import 'package:alcor/service/capsulate.dart';
import 'package:crypto/crypto.dart';
import 'package:ecdsa/ecdsa.dart';
import 'package:elliptic/elliptic.dart';
import 'package:get/get.dart';
import 'package:hive/hive.dart';
import 'package:http/http.dart';
import 'package:alcor/proto/protobuf/auth.pb.dart' as auth;

Future<void> requestKeybytes() async {
  var configController = Get.find<ConfigController>();
  var voteController = Get.find<VoteController>();
  var hashID = Hive.box<Uint8List>(configController.domain).get('id')!;
  var capsule = await encapsulate(hashID.toList());

  post(Uri.parse('http://${configController.domain}:9999/vote/request'),
          body: capsule.writeToBuffer())
      .then((value) async {
    if (value.statusCode == 200) {
      voteController.paper.randomBytes = await decapsulate(value.bodyBytes.toList());
      voteController.receiveKey.send(true);
    } else {
      await Get.defaultDialog(
          title: '투표를 진행할 수 없습니다.', middleText: value.statusCode.toString());
      voteController.receiveKey.send(false);
    }
  }).onError((error, stackTrace) async {
    await Get.defaultDialog(
        title: '투표를 진행할 수 없습니다.', middleText: error.toString());
    voteController.receiveKey.send(false);
  });
}

Future<void> submitPaper() async {
  var configController = Get.find<ConfigController>();
  var receiptController = Get.find<ReceiptController>();
  requestKeybytes();
  var rec = await Get.find<VoteController>().receiveKey.receive();
  if (!rec.data!) {
    return;
  }
  var paper = Get.find<VoteController>().paper;
  var now = DateTime.now();
  paper.timestamp =
      now.millisecondsSinceEpoch.toRadixString(16).codeUnits;
  receiptController.time = now;
  paper.voterId = Hive.box<Uint8List>(configController.domain).get('id')!;
  receiptController.voterID = paper.voterId;
  paper.message =
      utf8.encoder.convert(Get.find<CandidateDetailController>().name);
  receiptController.candidate = Get.find<CandidateDetailController>().name;
  paper.hash = sha512
      .convert(
          paper.randomBytes + paper.voterId + paper.timestamp + paper.message)
      .bytes;
  var privateKey = PrivateKey.fromBytes(
      getP521(),
      Hive.box<Uint8List>(configController.domain)
          .get('private_key')!
          .toList());
  receiptController.privateKey = privateKey.bytes;
  var sig = signature(privateKey, paper.hash);
  paper.signature =
      auth.Signature(r: sig.R.toRadixString(16), s: sig.S.toRadixString(16))
          .writeToBuffer();
  receiptController.signature = paper.signature;
  var capsule = await encapsulate(paper.writeToBuffer());
  post(Uri.parse('http://${configController.domain}:9999/vote/submit'),
          body: capsule.writeToBuffer())
      .then((value) async {
    if (value.statusCode == 200) {
      var id = await decapsulate(value.bodyBytes);
      await Get.defaultDialog(
          title: '투표를 등록하였습니다.',
          middleText: 'paper id: ${id.toString()}');
      receiptController.hash = id;
      Get.toNamed(receiptRoute);
    } else {
      await Get.defaultDialog(
          title: '투표를 완료할 수 없습니다.', middleText: value.statusCode.toString() + ' ' + value.body);
    }
  }).onError((error, stackTrace) async => await Get.defaultDialog(
          title: '투표를 완료할 수 없습니다.', middleText: error.toString()));
}
