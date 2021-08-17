import 'dart:convert';

import 'package:alcor/controllers/receipt.dart';
import 'package:alcor/pages/survey.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:hive/hive.dart';
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
    dic['time_stamp'] =
        controller.time.millisecondsSinceEpoch.toRadixString(16);
    dic['paper_hash'] = controller.hash.map((e) => e.toRadixString(16)).join();
    dic['candidate'] = controller.candidate;
    dic['signature'] =
        controller.signature.map((e) => e.toRadixString(16)).join();
    return Scaffold(
      appBar: AppBar(
        title: Text('투표 내역'),
        automaticallyImplyLeading: false,
      ),
      body: SingleChildScrollView(
        child: Column(
          children: [
            GetBuilder<ReceiptController>(
              builder: (controller) => Center(
                child: Expanded(
                  child: Container(
                    padding: EdgeInsets.all(24),
                    child: Card(
                      child: QrImage(
                        data: jsonEncode(dic),
                        size: MediaQuery.of(context).size.width <
                                MediaQuery.of(context).size.height
                            ? MediaQuery.of(context).size.width - 200
                            : MediaQuery.of(context).size.height - 200,
                      ),
                    ),
                  ),
                ),
              ),
            ),
            Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                ElevatedButton(
                  child: Text('설문조사에 참여합니다'),
                  onPressed: () {
                    Get.toNamed(surveyRoute);
                  },
                ),
                SizedBox(
                  width: 24,
                ),
                ElevatedButton(
                  onPressed: () async {
                    await Hive.deleteFromDisk();
                    Get.offAndToNamed('/');
                    await Get.defaultDialog(
                      title: '투표를 종료하였습니다.',
                      middleText: '모든 데이터를 지우고 투표 과정을 초기화하였습니다.',
                    );
                  },
                  child: Text('투표를 종료합니다'),
                ),
              ],
            ),
          ],
        ),
      ),
    );
  }
}
