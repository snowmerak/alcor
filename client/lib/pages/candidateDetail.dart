import 'package:alcor/controllers/candidate.dart';
import 'package:alcor/controllers/vote.dart';
import 'package:alcor/service/candidate.dart';
import 'package:alcor/service/vote.dart';
import 'package:flutter/material.dart';
import 'package:flutter_markdown/flutter_markdown.dart';
import 'package:get/get.dart';

final candidateDetailRoute = '/candidate/detail';

class CandidateDetailPage extends StatelessWidget {
  const CandidateDetailPage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    var voteController = Get.put(VoteController());
    var detailController = Get.find<CandidateDetailController>();
    getCandidate();
    return Scaffold(
      appBar: AppBar(
        title: Text(detailController.name),
      ),
      body: GetBuilder<CandidateDetailController>(
        builder: (controller) {
          return Container(
            child: Column(
              children: [
                Expanded(
                  child: Markdown(
                    data: detailController.markdown,
                  ),
                ),
                ElevatedButton(
                    onPressed: () {
                      submitPaper();
                    },
                    child: Text('투표'),
                ),
                SizedBox(
                  height: 24,
                ),
              ],
            ),
          );
        },
      ),
    );
  }
}
