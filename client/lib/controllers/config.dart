import 'package:cryptography/cryptography.dart';
import 'package:get/get.dart';

class ConfigController extends GetxController {
  String domain = '';
  String id = '';
  AesGcm aead = AesGcm.with256bits();
  SecretKey secretKey = SecretKey([]);
  List<int> nonce = [];
}