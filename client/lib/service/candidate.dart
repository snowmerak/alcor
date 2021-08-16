import 'dart:convert';

import 'package:alcor/controllers/candidate.dart';
import 'package:alcor/controllers/config.dart';
import 'package:alcor/proto/protobuf/candidate.pb.dart';
import 'package:get/get.dart';
import 'package:http/http.dart';

Future<void> getCandidates() async {
  var configController = Get.find<ConfigController>();
  var candidateController = Get.find<CandidateController>();
  get(Uri.parse('http://${configController.domain}:9999/candidate'))
    .then((value) async {
      if (value.statusCode == 200) {
        var list = List_.fromBuffer(value.bodyBytes);
        candidateController.updateNames(list.names);
      } else {
        await Get.defaultDialog(title: '후보 목록을 불러올 수 없습니다.', middleText: value.statusCode.toString());
      }
  })
    .onError((error, stackTrace) async => await Get.defaultDialog(title: '후보 목록을 불러올 수 없습니다.', middleText: error.toString()));
}

Future<void> getCandidate() async {
  var configController = Get.find<ConfigController>();
  var candidateDetailController = Get.find<CandidateDetailController>();

  var name = base64Url.encode(utf8.encoder.convert(candidateDetailController.name));

  get(Uri.parse('http://${configController.domain}:9999/candidate/$name'))
      .then((value) async {
        if (value.statusCode == 200) {
          var candidate = Candidate.fromBuffer(value.bodyBytes);
          candidateDetailController.updateMarkdown(candidate.markdown);
        } else {
          await Get.defaultDialog(title: '후보를 불러올 수 없습니다.', middleText: value.statusCode.toString());
        }
      })
      .onError((error, stackTrace) async => await Get.defaultDialog(title: '후보를 불러올 수 없습니다.', middleText: error.toString()));
}