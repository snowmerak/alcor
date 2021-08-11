import 'package:alcor/service/candidates.dart';
import 'package:get/get.dart';

class CandidatesController extends GetxController {
  List<String> candidates = [];
  List<String> filtered = [];

  String inputName = '';

  String selected = '';

  String detailName = '';
  String markdown = '';

  bool isDetail = false;

  void updateCandidates(List<String> names) {
    candidates.clear();
    candidates.addAll(names);
    filtered.clear();
    filtered.addAll(candidates);
    update();
  }

  void updateSelected(String data, String domain) {
    selected = data;
    isDetail = true;
    getCandidate(selected, domain).then((value) {
      detailName = value.name;
      markdown = value.markdown;
      update();
    }).onError((error, stackTrace) => Get.defaultDialog(title: 'Server Error!', middleText: error.toString()));
  }

  void resetSelected() {
    selected = '';
    isDetail = false;
    update();
  }

  void updateInputName(String name) {
    inputName = name;
    filtered = candidates.where((element) => element.contains(inputName)).toList();
    update();
  }
}