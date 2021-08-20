import 'package:alcor/controllers/selection.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';

final selectionRoute = '/survey/selection';

class SurveySelectionPage extends StatelessWidget {
  const SurveySelectionPage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(Get.find<SelectionController>().title),
      ),
      body: Container(
        padding: EdgeInsets.all(24),
        child: Column(
          children: [
            TextField(
              decoration: InputDecoration(
                border: OutlineInputBorder(),
                labelText: 'Name',
              ),
              onSubmitted: (String value) {
                Get.find<SelectionController>().filter(value);
              },
            ),
            GetBuilder<SelectionController>(
              builder: (controller) => Expanded(
                child: ListView.separated(
                  itemBuilder: (context, index) => TextButton(
                    style: TextButton.styleFrom(
                      primary: Colors.black87,
                    ),
                    onPressed: () {
                      controller.selected.send(controller.filtered[index]);
                      Get.back();
                    },
                    child: Text(controller.filtered[index]),
                  ),
                  separatorBuilder: (context, index) => Divider(),
                  itemCount: controller.filtered.length,
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
