import 'package:alcor/controllers/config.dart';
import 'package:alcor/proto/protobuf/capsule.pb.dart';
import 'package:cryptography/cryptography.dart';
import 'package:get/get.dart';

Future<List<int>> decapsulate(List<int> data) async {
  var configController = Get.find<ConfigController>();
  var capsule = Capsule.fromBuffer(data);
  data = capsule.data;
  var mac = Mac(data.sublist(data.length - 16, data.length));
  data = data.sublist(0, data.length - 16);
  var content = await configController.aead.decrypt(
      SecretBox(data, nonce: configController.nonce, mac: mac),
      secretKey: configController.secretKey);
  return content;
}

Future<Capsule> encapsulate(List<int> data) async {
  var configController = Get.find<ConfigController>();
  var capsule = Capsule();
  capsule.iD = configController.id;
  var box = await configController.aead.encrypt(data,
      secretKey: configController.secretKey, nonce: configController.nonce);
  capsule.data = box.cipherText + box.mac.bytes;
  return capsule;
}
