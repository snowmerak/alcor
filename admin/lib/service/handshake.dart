import 'dart:typed_data';

import 'package:admin/controller/index.dart';
import 'package:admin/proto/protobuf/security.pbserver.dart';
import 'package:cryptography/cryptography.dart';
import 'package:get/get.dart';
import 'package:http/http.dart';

Future<void> handshake() async {
  var configController = Get.find<IndexController>();
  var domain = configController.domain;
  var keyPair = await X25519().newKeyPair();
  var publicKey = await keyPair.extractPublicKey();
  try {
    var value = await post(Uri.parse('http://$domain:9999/handshake'),
        body: Uint8List.fromList(publicKey.bytes));
    var pub = Reply.fromBuffer(value.bodyBytes);
    var remoteKey = SimplePublicKey(pub.publicKey, type: KeyPairType.x25519);
    var sec = await X25519()
        .sharedSecretKey(keyPair: keyPair, remotePublicKey: remoteKey);
    var secret = await sec.extractBytes();
    configController.id = pub.iD;
    configController.aead = AesGcm.with256bits();
    configController.secretKey = SecretKey(secret.toList());
    configController.nonce = secret.sublist(0, 12);
  } catch (err) {
    await Get.defaultDialog(title: '통신 에러', middleText: err.toString());
  }
}
