
import 'package:alcor/proto/protobuf/stats.pb.dart';
import 'package:get/get.dart';

class SurveyController extends GetxController {
  Data survey = Data(
    region: '',
    job: '',
    education: '',
    houseType: '',
    ideology: '',
  );
  List<String> regions = [];
  List<String> jobs = [];
  List<String> educations = [];
  List<String> houseTypes = [];
  List<String> ideologys = [];
}