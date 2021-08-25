import 'package:admin/controller/candidate.dart';
import 'package:admin/page/candidate_detail.dart';
import 'package:admin/service/candidate.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';

final candidateListRoute = '/candidate/list';

class CandidateListPage extends StatelessWidget {
  const CandidateListPage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    Get.put(CandidateListController());
    getCandidates();
    return Scaffold(
      appBar: AppBar(
        title: Text('후보 목록'),
      ),
      body: GetBuilder<CandidateListController>(
        builder: (controller) => Container(
          padding: EdgeInsets.all(24),
          child: ListView.separated(
              itemBuilder: (context, index) => TextButton(
                    style: TextButton.styleFrom(
                      primary: Colors.black87,
                    ),
                    onPressed: () {
                      controller.selected = controller.list[index];
                      Get.toNamed(candidateDetailRoute);
                    },
                    child: Text(
                      controller.list[index],
                      style: TextStyle(
                        fontWeight: FontWeight.bold,
                        fontSize: 24,
                      ),
                    ),
                  ),
              separatorBuilder: (context, index) => Divider(),
              itemCount: controller.list.length),
        ),
      ),
    );
  }
}
