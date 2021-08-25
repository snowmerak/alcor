import 'dart:typed_data';

import 'package:admin/controller/index.dart';
import 'package:admin/page/menu.dart';
import 'package:admin/proto/protobuf/capsule.pb.dart';
import 'package:admin/service/handshake.dart';
import 'package:get/get.dart';
import 'package:http/http.dart';

Future<void> connect() async {
  var indexController = Get.find<IndexController>();
  try {
    await handshake();

    var data = await indexController.aead.encrypt(
        Uint8List.fromList(indexController.password.codeUnits),
        secretKey: indexController.secretKey,
        nonce: indexController.nonce);

    var capsule = Capsule(
      iD: indexController.id,
      data: data.cipherText + data.mac.bytes,
    );

    var value = await post(
        Uri.parse('http://${Get.find<IndexController>().domain}:9999/admin'),
        body: capsule.writeToBuffer());
    if (value.statusCode == 200) {
      await Get.defaultDialog(title: '확인되었습니다', middleText: '메뉴로 넘어갑니다.');
      Get.toNamed(menuRoute);
    } else {
      await Get.defaultDialog(
          title: '실패하였습니다.', middleText: value.statusCode.toString());
    }
  } catch (err) {
    await Get.defaultDialog(title: '통신 에러', middleText: err.toString());
  }
}
