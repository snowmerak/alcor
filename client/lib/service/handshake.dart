import 'dart:typed_data';

import 'package:alcor/controllers/config.dart';
import 'package:alcor/pages/candidate.dart';
import 'package:alcor/pages/register.dart';
import 'package:alcor/proto/protobuf/security.pbserver.dart';
import 'package:cryptography/cryptography.dart';
import 'package:get/get.dart';
import 'package:hive/hive.dart';
import 'package:http/http.dart';

Future<void> handshake() async {
  var configController = Get.find<ConfigController>();
  var domain = configController.domain;
  var keyPair = await X25519().newKeyPair();
  var publicKey = await keyPair.extractPublicKey();
  post(Uri.parse('http://$domain:9999/handshake'), body: Uint8List.fromList(publicKey.bytes)).then((value) async {
    var pub = Reply.fromBuffer(value.bodyBytes);
    var remoteKey =
    SimplePublicKey(pub.publicKey, type: KeyPairType.x25519);
    var sec = await X25519()
        .sharedSecretKey(keyPair: keyPair, remotePublicKey: remoteKey);
    var secret = await sec.extractBytes();
    configController.id = pub.iD;
    configController.aead = AesGcm.with256bits();
    configController.secretKey = SecretKey(secret.toList());
    configController.nonce = secret.sublist(0, 12);
    Get.defaultDialog(title: '연결 했습니다', middleText: '서버에 연결되었습니다!');
    await Hive.openBox<Uint8List>(configController.domain);
    var id = Hive.box<Uint8List>(configController.domain).get('id');
    if (id == null) {
      Get.toNamed(registerRoute);
    } else {
      Get.toNamed(candidateRoute);
    }
    Get.toNamed(registerRoute);
  }).catchError((err) {
    Get.defaultDialog(title: '에러가 발생하였습니다', middleText: err.toString());
  });
}
