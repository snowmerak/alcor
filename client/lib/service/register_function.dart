import 'dart:typed_data';

import 'package:alcor/protobuf/auth.pbserver.dart' as auth;
import 'package:alcor/protobuf/register.pbserver.dart';
import 'package:cryptography/cryptography.dart';
import 'package:ecdsa/ecdsa.dart';
import 'package:elliptic/elliptic.dart';
import 'package:flutter/material.dart';
import 'package:hive/hive.dart';
import 'package:rflutter_alert/rflutter_alert.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

Future<void> registerCandidate(
    BuildContext context, String name, String markdown) async {
  var domain =
      String.fromCharCodes(Hive.box<Uint8List>('config').get('server')!);
  var channel =
      WebSocketChannel.connect(Uri.parse('ws://$domain:9999/register'));
  List<int>? secret;
  var state = 0;
  var keyPair = await X25519().newKeyPair();
  AesGcm aead = AesGcm.with256bits();
  SecretKey secretKey = SecretKey([]);
  List<int> nonce = [];
  var keyBytes = <int>[];
  channel.sink
      .add(Uint8List.fromList((await keyPair.extractPublicKey()).bytes));
  channel.stream.listen((event) async {
    print(state);
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
        keyBytes = await aead.decrypt(box, secretKey: secretKey);

        var request = Request();
        request.keyBytes = keyBytes;
        request.name = name;
        request.markdown = markdown;
        request.voterID = Hive.box<Uint8List>(domain).get('id')!.toList();
        var privateKey = PrivateKey.fromBytes(getP521(),
            Hive.box<Uint8List>(domain).get('private_key')!.toList());
        var sha = Sha256();
        var hash = await sha.hash(request.keyBytes +
            request.name.codeUnits +
            request.markdown.codeUnits +
            request.voterID);
        request.hash = hash.bytes;
        var sign = signature(privateKey, hash.bytes);
        var sig = auth.Signature();
        sig.r = sign.R.toRadixString(16);
        sig.s = sign.S.toRadixString(16);
        request.signature = sig.writeToBuffer().toList();

        box = await aead.encrypt(request.writeToBuffer().toList(),
            secretKey: secretKey, nonce: nonce);
        channel.sink.add(Uint8List.fromList(box.cipherText + box.mac.bytes));

        state = 2;
        break;
      case 2:
        var buf = (event as Uint8List);
        var box = SecretBox(buf.sublist(0, buf.length - 16).toList(),
            nonce: nonce, mac: Mac(buf.sublist(buf.length - 16, buf.length)));
        var result =
            Result.fromBuffer(await aead.decrypt(box, secretKey: secretKey));

        if (!result.oK) {
          Alert(context: context, title: 'Error!', desc: result.error).show();
          channel.sink.close();
          break;
        }
        Alert(context: context, title: 'Succeed!', desc: 'Registered').show();
        channel.sink.close();
        break;
    }
  });
  channel.stream.timeout(Duration(minutes: 2), onTimeout: (event) {
    Alert(
            context: context,
            title: 'TimeOut!',
            desc: 'an error occurred in server')
        .show();
  });
}
