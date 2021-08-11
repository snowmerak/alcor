import 'dart:typed_data';

import 'package:alcor/controllers/config.dart';
import 'package:alcor/controllers/signin.dart';
import 'package:alcor/pages/candidate.dart';
import 'package:alcor/protobuf/cert.pb.dart';
import 'package:alcor/protobuf/auth.pb.dart' as auth;
import 'package:cryptography/cryptography.dart';
import 'package:elliptic/elliptic.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:hive/hive.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

Future<void> enrollVoter(BuildContext context) async {
  var configController = Get.find<ConfigController>();
  var signinController = Get.find<SigninController>();
  var domain = configController.domain;
  var channel =
      WebSocketChannel.connect(Uri.parse('ws://$domain:9999/enroll/phone'));
  List<int>? secret;
  var state = 0;
  X25519().newKeyPair().then((keyPair) async {
    AesGcm aead = AesGcm.with256bits();
    SecretKey secretKey = SecretKey([]);
    var privateKey = getP521().generatePrivateKey();
    List<int> nonce = [];
    channel.sink
        .add(Uint8List.fromList((await keyPair.extractPublicKey()).bytes));
    channel.stream.listen((event) async {
      switch (state) {
        case 0:
          var pub = event as Uint8List;
          var remoteKey =
              SimplePublicKey(pub.toList(), type: KeyPairType.x25519);
          var sec = await X25519()
              .sharedSecretKey(keyPair: keyPair, remotePublicKey: remoteKey);
          secret = await sec.extractBytes();
          aead = AesGcm.with256bits();
          secretKey = SecretKey(secret!.toList());
          nonce = secret!.sublist(0, 12);

          state = 2;

          var phone = Phone();
          phone.gender = true;
          if (signinController.gender[0]) {
            phone.gender = false;
          }
          phone.name = signinController.name;
          phone.number = signinController.phoneNumber;
          phone.rRN = signinController.birth;
          var box = await aead.encrypt(phone.writeToBuffer().toList(),
              secretKey: secretKey, nonce: nonce);
          channel.sink.add(Uint8List.fromList(box.cipherText + box.mac.bytes));
          break;
        case 1:
          state = 2;
          break;
        case 2:
          var buf = (event as Uint8List);
          var box = SecretBox(buf.sublist(0, buf.length - 16).toList(),
              nonce: nonce, mac: Mac(buf.sublist(buf.length - 16, buf.length)));
          var result = Result.fromBuffer(Uint8List.fromList(
              await aead.decrypt(box, secretKey: secretKey)));
          if (!result.ok) {
            Get.defaultDialog(title: '에러 발생', middleText: result.error);
            channel.sink.close();
            break;
          }

          var code = (await signinController.signal.receive()).data;
          box = await aead.encrypt(Uint8List.fromList(code!.codeUnits),
              secretKey: secretKey, nonce: nonce);
          channel.sink.add(
              Uint8List.fromList((box.cipherText + box.mac.bytes).toList()));

          state = 3;
          break;
        case 3:
          var buf = (event as Uint8List);
          var box = SecretBox(buf.sublist(0, buf.length - 16).toList(),
              nonce: nonce, mac: Mac(buf.sublist(buf.length - 16, buf.length)));
          var result = Result.fromBuffer(Uint8List.fromList(
              await aead.decrypt(box, secretKey: secretKey)));
          if (result.error != '') {
            Get.defaultDialog(title: '에러 발생', middleText: result.error);
            channel.sink.close();
            break;
          }

          var publicKey = auth.PublicKey();
          publicKey.x = privateKey.publicKey.X.toRadixString(16);
          publicKey.y = privateKey.publicKey.Y.toRadixString(16);

          box = await aead.encrypt(publicKey.writeToBuffer(),
              secretKey: secretKey, nonce: nonce);
          channel.sink.add(
              Uint8List.fromList((box.cipherText + box.mac.bytes).toList()));

          state = 4;
          break;
        case 4:
          var buf = (event as Uint8List);
          var box = SecretBox(buf.sublist(0, buf.length - 16).toList(),
              nonce: nonce, mac: Mac(buf.sublist(buf.length - 16, buf.length)));
          var result = ID.fromBuffer(Uint8List.fromList(
              await aead.decrypt(box, secretKey: secretKey)));
          if (result.error != '') {
            Get.defaultDialog(title: '에러 발생', middleText: result.error);
            channel.sink.close();
            break;
          }
          await Get.defaultDialog(
              title: '등록 성공',
              middleText:
                  'your new id is ${result.iD.map((e) => e.toRadixString(16)).join()}',
              confirm: ElevatedButton(
                onPressed: () {
                  Get.back();
                },
                child: Text('확인'),
              ));
          await Hive.box<Uint8List>(domain)
              .put('id', Uint8List.fromList(result.iD));
          await Hive.box<Uint8List>(domain)
              .put('private_key', Uint8List.fromList(privateKey.bytes));
          channel.sink.close();
          Get.toNamed(candidateListRoute);
          break;
      }
    });
  });
  channel.stream.timeout(Duration(minutes: 2), onTimeout: (event) {
    Get.defaultDialog(title: '시간 초과', middleText: '서버에서 에러가 발생했습니다.');
  });
}
