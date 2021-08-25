import 'package:get/get.dart';

class CandidateController extends GetxController {
  String name = '';
  String markdown = '';
}

class CandidateDetailController extends GetxController {
  String name = '';
  String markdown = '';
  bool nowEditing = false;

  void updateName(String name) {
    this.name = name;
    update();
  }

  void updateMarkdown(String markdown) {
    this.markdown = markdown;
    update();
  }
}

class CandidateListController extends GetxController {
  List<String> list = [];
  String selected = '';
}
