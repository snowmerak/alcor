import 'package:channel/channel.dart';
import 'package:get/get.dart';

class SelectionController extends GetxController {
  String title = '';
  List<String> list = [];
  List<String> filtered = [];
  Channel<String> selected = Channel();

  void filter(String value) {
    filtered = list.where((element) => element.contains(value)).toList();
    update();
  }
}