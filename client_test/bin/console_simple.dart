import 'dart:convert';
import 'dart:math';
import 'dart:typed_data';

import 'package:crypto/crypto.dart';
import 'package:cryptography/cryptography.dart';
import 'package:ecdsa/ecdsa.dart';
import 'package:elliptic/elliptic.dart';
import 'package:http/http.dart';

import 'protobuf/stats.pb.dart';
import 'protobuf/capsule.pb.dart';
import 'protobuf/register.pb.dart';
import 'protobuf/security.pb.dart';
import 'protobuf/auth.pb.dart' as auth;
import 'protobuf/vote.pb.dart' as vote;

Future<void> main(List<String> arguments) async {
  var url = '127.0.0.1';
  var port = '9999';

  var names = <String>[
    '민준',
    '서준',
    '예준',
    '도윤',
    '시우',
    '주원',
    '하준',
    '도현',
    '서연',
    '서윤',
    '지우',
    '지윤',
    '은서',
    '수아',
    '지아',
    '예린',
    '유찬',
    '예진',
    '서아',
    '지원',
    '민규',
    '다인',
    '유빈',
    '현서',
  ];

  var candidates = <String>[
    'golang',
    'python',
    'kotlin',
    'julia',
    'rust',
    'dart',
    'java',
  ];

  for (var name in names) {
    print('$name start');
    var info = await handshake(url, port);
    print('$name handshake end: ${info.id}');
    var candidate = candidates[Random().nextInt(candidates.length)];
    var registred = await register(
      url,
      port,
      info,
      candidate,
    );
    print('$name register end: $registred');
    if (!registred) {
      continue;
    }
    var account = await registerCert(url, port, info);
    print('$name registerCert end: ${account!.hashID.toString()}');
    var randombytes = await requestKeybytes(url, port, info, account);
    print('$name requestKeybytes end: ${randombytes?.randomBytes}');
    var result = await submitPaper(
        url,
        port,
        info,
        account,
        randombytes!.randomBytes,
        candidates[Random().nextInt(max(0, candidates.length - 1))]);
    print('$name submitPaper end: $result');
    print('$name is voted');

    result = await submitSurvey(url, candidate, info, account);
    print('$name submitSurvey end: $result');
    print('$name end');
  }
}

class Info {
  late String id;
  late AesGcm aead;
  late SecretKey secretKey;
  late List<int> nonce;
}

Future<Info> handshake(String url, String port) async {
  var keyPair = await X25519().newKeyPair();
  var publicKey = await keyPair.extractPublicKey();

  var info = Info();

  var value = await post(Uri.parse('http://$url:$port/handshake'),
      body: Uint8List.fromList(publicKey.bytes));
  var pub = Reply.fromBuffer(value.bodyBytes);
  var remoteKey = SimplePublicKey(pub.publicKey, type: KeyPairType.x25519);
  var sec = await X25519()
      .sharedSecretKey(keyPair: keyPair, remotePublicKey: remoteKey);
  var secret = await sec.extractBytes();
  info.id = pub.iD;
  info.aead = AesGcm.with256bits();
  info.secretKey = SecretKey(secret.toList());
  info.nonce = secret.sublist(0, 12);
  print("connected");

  return info;
}

Future<List<int>> decapsulate(Info info, List<int> data) async {
  var capsule = Capsule.fromBuffer(data);
  data = capsule.data;
  var mac = Mac(data.sublist(data.length - 16, data.length));
  data = data.sublist(0, data.length - 16);
  var content = await info.aead.decrypt(
      SecretBox(data, nonce: info.nonce, mac: mac),
      secretKey: info.secretKey);
  return content;
}

Future<Capsule> encapsulate(Info info, List<int> data) async {
  var capsule = Capsule();
  capsule.iD = info.id;
  var box = await info.aead
      .encrypt(data, secretKey: info.secretKey, nonce: info.nonce);
  capsule.data = box.cipherText + box.mac.bytes;
  return capsule;
}

Future<bool> register(String url, String port, Info info, String name) async {
  var phone = Phone();
  phone.name = name;
  phone.number = '010-1234-5678';
  phone.gender = false;
  phone.rRN = '000000';
  var capsule = await encapsulate(info, phone.writeToBuffer());
  var value = await post(Uri.parse('http://$url:9999/voter/register/request'),
      body: capsule.writeToBuffer());

  if (value.statusCode == 200) {
    return true;
  }
  return false;
}

Future<Account?> registerCert(String url, String port, Info info) async {
  var privateKey = getP256().generatePrivateKey();
  var publicKey = privateKey.publicKey;
  var cert = CertCode();
  cert.code = '0000'.codeUnits;
  cert.publicKey = auth.PublicKey(
          x: publicKey.X.toRadixString(16), y: publicKey.Y.toRadixString(16))
      .writeToBuffer();
  var capsule = await encapsulate(info, cert.writeToBuffer());
  var value = await post(Uri.parse('http://$url:$port/voter/register/cert'),
      body: capsule.writeToBuffer());
  var result = Result.fromBuffer(await decapsulate(info, value.bodyBytes));
  var account = Account();
  if (value.statusCode == 200) {
    account.hashID = result.iD;
    account.privateKey = privateKey.bytes;
    return account;
  }
  return null;
}

class InnerPaper {
  late List<int> randomBytes;
}

Future<InnerPaper?> requestKeybytes(
    String url, String port, Info info, Account account) async {
  var hashID = account.hashID;
  var capsule = await encapsulate(info, hashID.toList());

  var paper = InnerPaper();

  var value = await post(Uri.parse('http://$url:9999/vote/request'),
      body: capsule.writeToBuffer());
  if (value.statusCode == 200) {
    paper.randomBytes = await decapsulate(info, value.bodyBytes.toList());
    return paper;
  }
  return null;
}

Future<bool> submitPaper(String url, String port, Info info, Account account,
    List<int> randombytes, String name) async {
  var now = DateTime.now();
  var paper = vote.Paper();
  paper.randomBytes = randombytes;
  paper.timestamp = now.millisecondsSinceEpoch.toRadixString(16).codeUnits;
  paper.voterId = account.hashID;
  paper.message = utf8.encoder.convert(name);
  paper.hash = sha256
      .convert(
          paper.randomBytes + paper.voterId + paper.timestamp + paper.message)
      .bytes;
  var privateKey = PrivateKey.fromBytes(getP256(), account.privateKey);
  var sig = signature(privateKey, paper.hash);
  paper.signature =
      auth.Signature(r: sig.R.toRadixString(16), s: sig.S.toRadixString(16))
          .writeToBuffer();
  var capsule = await encapsulate(info, paper.writeToBuffer());
  var value = await post(Uri.parse('http://$url:9999/vote/submit'),
      body: capsule.writeToBuffer());
  if (value.statusCode == 200) {
    var id = await decapsulate(info, value.bodyBytes);
    print('투표를 등록하였습니다. paper id: ${id.toString()}');
    return true;
  } else {
    print('투표를 완료할 수 없습니다. ${value.statusCode.toString() + ' ' + value.body}');
    return false;
  }
}

Future<bool> submitSurvey(
    String url, String name, Info info, Account account) async {
  var value = await get(Uri.parse('http://$url:9999/survey/regions'));
  var regions = List_.fromBuffer(value.bodyBytes);

  value = await get(Uri.parse('http://$url:9999/survey/jobs'));
  var jobs = List_.fromBuffer(value.bodyBytes);

  value = await get(Uri.parse('http://$url:9999/survey/educations'));
  var educations = List_.fromBuffer(value.bodyBytes);

  value = await get(Uri.parse('http://$url:9999/survey/house_types'));
  var houseTypes = List_.fromBuffer(value.bodyBytes);

  value = await get(Uri.parse('http://$url:9999/survey/ideologys'));
  var ideologys = List_.fromBuffer(value.bodyBytes);

  var rand = Random();
  var survey = Data();
  survey.salary = rand.nextInt(98) + 1;
  survey.age = rand.nextInt(81) + 1922;
  survey.candidate = name;
  survey.divorced = rand.nextBool();
  survey.education = educations.data[rand.nextInt(educations.data.length)];
  survey.houseType = houseTypes.data[rand.nextInt(houseTypes.data.length)];
  survey.ideology = ideologys.data[rand.nextInt(ideologys.data.length)];
  survey.job = jobs.data[rand.nextInt(jobs.data.length)];
  survey.region = regions.data[rand.nextInt(regions.data.length)];
  survey.hasCar = rand.nextBool();
  survey.hasDebt = rand.nextBool();
  survey.married = rand.nextBool();
  survey.voterID = account.hashID;
  survey.gender = rand.nextBool();

  var result = await post(Uri.parse('http://$url:9999/survey/reply'),
      body: survey.writeToBuffer());
  if (result.statusCode == 200) {
    return true;
  }

  return false;
}
