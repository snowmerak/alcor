import 'package:get/get.dart';
import 'package:cryptography/cryptography.dart';

class IndexController extends GetxController {
  String domain = '';
  String password = '';
  String id = '';
  AesGcm aead = AesGcm.with256bits();
  SecretKey secretKey = SecretKey([]);
  List<int> nonce = [];
}
