import 'package:alcor/controllers/candidate.dart';
import 'package:alcor/controllers/receipt.dart';
import 'package:alcor/pages/candidateDetail.dart';
import 'package:alcor/service/candidate.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';

final candidateRoute = '/candidate';

class CandidatePage extends StatelessWidget {
  const CandidatePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    Get.put(CandidateController());
    var detailController = Get.put(CandidateDetailController());
    Get.put(ReceiptController());
    getCandidates();
    return Scaffold(
      appBar: AppBar(
        title: Text('후보 목록'),
        automaticallyImplyLeading: false,
      ),
      body: GetBuilder<CandidateController>(
        builder: (controller) {
          return ListView.separated(
              itemBuilder: (context, index) {
                return TextButton(
                  style: TextButton.styleFrom(
                    primary: Colors.black87,
                  ),
                  onPressed: () {
                    detailController.updateName(controller.names[index]);
                    Get.toNamed(candidateDetailRoute);
                  },
                  child: Text(
                      controller.names[index],
                      style: TextStyle(
                        fontWeight: FontWeight.bold,
                        fontSize: 24,
                      ),
                  ),
                );
              },
              separatorBuilder: (context, index) {
                return Divider();
              },
              itemCount: controller.names.length);
        },
      ),
    );
  }
}
