import 'dart:convert';
import 'dart:typed_data';

import 'package:alcor/controllers/candidates.dart';
import 'package:alcor/controllers/config.dart';
import 'package:alcor/protobuf/vote.pb.dart';
import 'package:alcor/protobuf/auth.pb.dart' as auth;
import 'package:cryptography/cryptography.dart';
import 'package:ecdsa/ecdsa.dart';
import 'package:elliptic/elliptic.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:hive/hive.dart';
import 'package:rflutter_alert/rflutter_alert.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

Future<void> vote(BuildContext context) async {
  var domain =
      Get.find<ConfigController>().domain;
  var channel =
      WebSocketChannel.connect(Uri.parse('ws://$domain:9999/vote'));
  List<int>? secret;
  var state = 0;
  var keyPair = await X25519().newKeyPair();
  AesGcm aead = AesGcm.with256bits();
  SecretKey secretKey = SecretKey([]);
  var privateKey = PrivateKey.fromBytes(getP521(), Hive.box<Uint8List>(domain).get('private_key')!);
  var paper = Paper();
  List<int> nonce = [];
  channel.sink
      .add(Uint8List.fromList((await keyPair.extractPublicKey()).bytes));
  channel.stream.listen((event) async {
    switch (state) {
      case 0:
        var pub = event as Uint8List;
        var remoteKey = SimplePublicKey(pub.toList(), type: KeyPairType.x25519);
        var sec = await X25519()
            .sharedSecretKey(keyPair: keyPair, remotePublicKey: remoteKey);
        secret = await sec.extractBytes();
        aead = AesGcm.with256bits();
        secretKey = SecretKey(secret!.toList());
        nonce = secret!.sublist(0, 12);

        state = 1;
        break;
      case 1:
        var buf = (event as Uint8List);
        var box = SecretBox(buf.sublist(0, buf.length - 16).toList(),
            nonce: nonce, mac: Mac(buf.sublist(buf.length - 16, buf.length)));
        paper.randomBytes = await aead.decrypt(box, secretKey: secretKey);
        paper.message = Utf8Encoder().convert(Get.find<CandidatesController>().detailName);
        var view = ByteData(8);
        view.setInt32(0, DateTime.now().millisecondsSinceEpoch.toInt());
        paper.timestamp = view.buffer.asUint8List().toList();
        paper.voterId = Hive.box<Uint8List>(domain).get('id')!.toList();
        var sha = Sha512();
        var hash = await sha.hash(paper.randomBytes + paper.voterId + paper.timestamp + paper.message);
        paper.hash = hash.bytes;
        var sig = signature(privateKey, paper.hash);
        paper.signature = auth.Signature(r: sig.R.toRadixString(16), s: sig.S.toRadixString(16)).writeToBuffer().toList();

        box = await aead.encrypt(paper.writeToBuffer(), secretKey: secretKey, nonce: nonce);
        channel.sink.add(Uint8List.fromList(box.cipherText + box.mac.bytes));
        state = 2;
        break;
      case 2:
        var buf = (event as Uint8List);
        var box = SecretBox(buf.sublist(0, buf.length - 16).toList(),
            nonce: nonce, mac: Mac(buf.sublist(buf.length - 16, buf.length)));
        var result = Result.fromBuffer(await aead.decrypt(box, secretKey: secretKey));
        if (!result.ok) {
          Get.defaultDialog(title: '에러 발생!', middleText: result.error);
          channel.sink.close();
          break;
        }
        Get.defaultDialog(title: '등록되었습니다!', middleText: '투표 내역이 등록되었습니다!');
        channel.sink.close();
        break;
    }
  });
  channel.stream.timeout(Duration(minutes: 2), onTimeout: (event) {
    Get.defaultDialog(title: 'TimeOut!', middleText: '서버의 응답이 너무 늦습니다. 잠시 후 다시 시도해주세요.');
  });
}
