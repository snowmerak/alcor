import 'package:alcor/controllers/config.dart';
import 'package:alcor/controllers/receipt.dart';
import 'package:alcor/controllers/survey.dart';
import 'package:alcor/controllers/vote.dart';
import 'package:alcor/proto/protobuf/stats.pbserver.dart';
import 'package:get/get.dart';
import 'package:hive/hive.dart';
import 'package:http/http.dart';

Future<void> submitSurvey() async {
  var receiptController = Get.find<ReceiptController>();
  var voteController = Get.find<VoteController>();
  var surveyController = Get.find<SurveyController>();

  surveyController.survey.candidate = receiptController.candidate;
  surveyController.survey.voterID = voteController.paper.voterId;

  var domain = Get.find<ConfigController>().domain;

  post(Uri.parse('http://$domain:9999/survey/reply'), body: surveyController.survey.writeToBuffer())
    .then((value) async {
      if (value.statusCode == 200) {
        await Get.defaultDialog(title: '제출하였습니다', middleText: '응답해 주셔서 감사합니다\n모든 데이터를 초기화합니다');
        await Hive.deleteFromDisk();
        Get.offAllNamed('/');
      } else {
        await Get.defaultDialog(title: '제출에 실패하였습니다', middleText: value.statusCode.toString());
      }
  })
    .onError((error, stackTrace) => null);
}

Future<void> getRegions() async {
  var domain = Get.find<ConfigController>().domain;
  var surveyController = Get.find<SurveyController>();

  get(Uri.parse('http://$domain:9999/survey/regions'))
    .then((value) {
      if (value.statusCode == 200) {
        var list = List_.fromBuffer(value.bodyBytes);
        surveyController.regions = list.data;
        surveyController.update();
      }
  }).onError((error, stackTrace) {
      Get.defaultDialog(title: '요청 실패', middleText: error.toString());
  });
}

Future<void> getJobs() async {
  var domain = Get.find<ConfigController>().domain;
  var surveyController = Get.find<SurveyController>();

  get(Uri.parse('http://$domain:9999/survey/jobs'))
      .then((value) {
    if (value.statusCode == 200) {
      var list = List_.fromBuffer(value.bodyBytes);
      surveyController.jobs = list.data;
      surveyController.update();
    }
  }).onError((error, stackTrace) {
    Get.defaultDialog(title: '요청 실패', middleText: error.toString());
  });
}

Future<void> getEducations() async {
  var domain = Get.find<ConfigController>().domain;
  var surveyController = Get.find<SurveyController>();

  get(Uri.parse('http://$domain:9999/survey/educations'))
      .then((value) {
    if (value.statusCode == 200) {
      var list = List_.fromBuffer(value.bodyBytes);
      surveyController.educations = list.data;
      surveyController.update();
    }
  }).onError((error, stackTrace) {
    Get.defaultDialog(title: '요청 실패', middleText: error.toString());
  });
}

Future<void> getHouseTypes() async {
  var domain = Get.find<ConfigController>().domain;
  var surveyController = Get.find<SurveyController>();

  get(Uri.parse('http://$domain:9999/survey/house_types'))
      .then((value) {
    if (value.statusCode == 200) {
      var list = List_.fromBuffer(value.bodyBytes);
      surveyController.houseTypes = list.data;
      surveyController.update();
    }
  }).onError((error, stackTrace) {
    Get.defaultDialog(title: '요청 실패', middleText: error.toString());
  });
}

Future<void> getIdeologys() async {
  var domain = Get.find<ConfigController>().domain;
  var surveyController = Get.find<SurveyController>();

  get(Uri.parse('http://$domain:9999/survey/ideologys'))
      .then((value) {
    if (value.statusCode == 200) {
      var list = List_.fromBuffer(value.bodyBytes);
      surveyController.ideologys = list.data;
      surveyController.update();
    }
  }).onError((error, stackTrace) {
    Get.defaultDialog(title: '요청 실패', middleText: error.toString());
  });
}