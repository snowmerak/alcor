import 'package:get/get.dart';

class ReceiptController extends GetxController {
  List<int> voterID = [];
  String candidate = '';
  DateTime time = DateTime.now();
  List<int> hash = [];
  String elliptic = 'P521';
  List<int> privateKey = [];
  List<int> signature = [];
}