import 'package:get/get.dart';

class CandidateController extends GetxController {
  List<String> names = [];

  void updateNames(List<String> names) {
    this.names = names;
    update();
  }
}

class CandidateDetailController extends GetxController {
  String name = '';
  String markdown = '';

  void updateName(String name) {
    this.name = name;
    update();
  }

  void updateMarkdown(String markdown) {
    this.markdown = markdown;
    update();
  }
}