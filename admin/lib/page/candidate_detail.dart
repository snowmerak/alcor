import 'package:admin/controller/candidate.dart';
import 'package:admin/service/candidate.dart';
import 'package:flutter/material.dart';
import 'package:flutter_markdown/flutter_markdown.dart';
import 'package:get/get.dart';

final candidateDetailRoute = '/candidate/detail';

class CandidateDetailPage extends StatelessWidget {
  const CandidateDetailPage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    var detailController = Get.put(CandidateDetailController());
    detailController.name = Get.find<CandidateListController>().selected;
    getCandidate();
    return Scaffold(
      appBar: AppBar(
        title: Text(detailController.name),
      ),
      body: GetBuilder<CandidateDetailController>(
        builder: (controller) {
          return Container(
            padding: EdgeInsets.all(24),
            child: Column(
              children: [
                Expanded(
                  child: controller.nowEditing
                      ? TextField(
                          controller: TextEditingController()
                            ..text = controller.markdown,
                          textInputAction: TextInputAction.newline,
                          keyboardType: TextInputType.multiline,
                          maxLines: null,
                          onChanged: (value) => controller.markdown = value,
                        )
                      : Markdown(
                          data: detailController.markdown,
                        ),
                ),
                Row(
                  mainAxisAlignment: MainAxisAlignment.end,
                  children: [
                    ElevatedButton(
                      onPressed: () {
                        if (!controller.nowEditing) {
                          controller.nowEditing = true;
                          controller.update();
                          return;
                        }
                        editCandidate();
                      },
                      child: controller.nowEditing ? Text('등록') : Text('수정'),
                    ),
                    SizedBox(
                      width: 16,
                    ),
                    controller.nowEditing
                        ? SizedBox()
                        : ElevatedButton(
                            onPressed: () => removeCandidate(),
                            child: Text('삭제'),
                          ),
                  ],
                ),
              ],
            ),
          );
        },
      ),
    );
  }
}
