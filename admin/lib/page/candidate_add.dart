import 'package:admin/controller/candidate.dart';
import 'package:admin/service/candidate.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';

final candidateAddRoute = '/candidate/add';

class CandidateAddPage extends StatelessWidget {
  const CandidateAddPage({Key? key}) : super(key: key);

  final double breaksize = 24;

  @override
  Widget build(BuildContext context) {
    var controller = Get.put(CandidateController());
    return Scaffold(
      appBar: AppBar(
        title: Text('후보자 추가'),
      ),
      body: Container(
        padding: EdgeInsets.all(24),
        child: Column(
          children: [
            TextField(
              decoration: InputDecoration(
                border: OutlineInputBorder(),
                labelText: '후보자 이름',
              ),
              onChanged: (value) => controller.name = value,
            ),
            SizedBox(
              height: breaksize,
            ),
            Expanded(
              child: TextField(
                textInputAction: TextInputAction.newline,
                keyboardType: TextInputType.multiline,
                maxLines: null,
                onChanged: (value) => controller.markdown = value,
              ),
            ),
            SizedBox(
              height: breaksize,
            ),
            ElevatedButton(
              onPressed: () => addCandidate(),
              child: Text('등록'),
            ),
          ],
        ),
      ),
    );
  }
}
