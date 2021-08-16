import 'package:channel/channel.dart';
import 'package:get/get.dart';

class RegisterController extends GetxController {
  String name = "";
  String phoneNumber = "";
  String birth = "";
  String certCode = "";
  List<bool> gender = [false, false];

  var signal = Channel<String>();
}