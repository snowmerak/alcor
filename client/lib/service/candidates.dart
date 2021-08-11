import 'dart:convert';
import 'dart:typed_data';

import 'package:alcor/controllers/candidates.dart';
import 'package:alcor/protobuf/candidate.pb.dart';
import 'package:flutter/services.dart';
import 'package:hive/hive.dart';
import 'package:http/http.dart';

Future<void> initList(CandidatesController controller, String domain) async {
  var res = await get(Uri.parse(
      'http://$domain:9999/candidate'));
  controller.updateCandidates(CandidateList.fromBuffer(res.bodyBytes).names);
}

Future<Candidate> getCandidate(String selected, String domain) async {
  try {
    var name = base64.encode(utf8.encode(selected));
    var res = await get(Uri.parse('http://$domain:9999/candidate/$name'));
    return Candidate.fromBuffer(res.bodyBytes);
  } catch (e) {
    throw e;
  }
}