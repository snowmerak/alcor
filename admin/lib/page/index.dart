import 'package:admin/controller/index.dart';
import 'package:admin/service/check.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';

final String indexRoute = '/';

class IndexPage extends StatelessWidget {
  const IndexPage({Key? key}) : super(key: key);

  final double breaksize = 24;

  @override
  Widget build(BuildContext context) {
    var controller = Get.put(IndexController());
    return Scaffold(
      appBar: AppBar(
        title: Text('서버 연결'),
      ),
      body: Container(
        padding: EdgeInsets.all(24),
        child: Column(
          children: [
            TextField(
              decoration: InputDecoration(
                border: OutlineInputBorder(),
                labelText: 'Server Domain',
              ),
              onChanged: (value) => {controller.domain = value},
            ),
            SizedBox(
              height: breaksize,
            ),
            TextField(
              decoration: InputDecoration(
                border: OutlineInputBorder(),
                labelText: 'Server password',
              ),
              onChanged: (value) => {controller.password = value},
            ),
            SizedBox(
              height: breaksize,
            ),
            ElevatedButton(
              onPressed: () => connect(),
              child: Text('연결'),
            ),
          ],
        ),
      ),
    );
  }
}
