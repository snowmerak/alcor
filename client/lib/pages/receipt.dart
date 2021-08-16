import 'dart:convert';

import 'package:alcor/controllers/receipt.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:qr_flutter/qr_flutter.dart';

final receiptRoute = '/receipt';

class ReceiptPage extends StatelessWidget {
  const ReceiptPage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    var dic = Map<String, dynamic>();
    var controller = Get.find<ReceiptController>();
    dic['voter_id'] = controller.voterID.map((e) => e.toRadixString(16)).join();
    dic['private_key'] =
        controller.privateKey.map((e) => e.toRadixString(16)).join();
    dic['elliptic'] = controller.elliptic;
    dic['time_stamp'] = controller.time.millisecondsSinceEpoch.toRadixString(16);
    dic['paper_hash'] = controller.hash.map((e) => e.toRadixString(16)).join();
    dic['candidate'] = controller.candidate;
    dic['signature'] =
        controller.signature.map((e) => e.toRadixString(16)).join();
    return Scaffold(
      appBar: AppBar(
        title: Text('투표 내역'),
        automaticallyImplyLeading: false,
      ),
      body: GetBuilder<ReceiptController>(
        builder: (controller) => Center(
          child: Container(
            padding: EdgeInsets.all(24),
            child: Card(
              child: QrImage(
                data: jsonEncode(dic),
              ),
            ),
          ),
        ),
      ),
    );
  }
}
