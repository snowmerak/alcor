import 'package:alcor/controllers/config.dart';
import 'package:alcor/controllers/vote.dart';
import 'package:alcor/service/handshake.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';

final String indexRoute = '/';

class Index extends StatelessWidget {
  const Index({Key? key}) : super(key: key);

  final double breakSize = 24;

  @override
  Widget build(BuildContext context) {
    final configController = Get.put(ConfigController());
    Get.put(VoteController());
    return Scaffold(
      appBar: AppBar(
        title: Text('서버 연결'),
        automaticallyImplyLeading: false,
      ),
      body: Container(
        padding: EdgeInsets.all(24),
        child: Center(
          child: Column(
            children: [
              TextField(
                decoration: InputDecoration(
                  border: OutlineInputBorder(),
                  labelText: 'Server Domain',
                ),
                onChanged: (value) => {configController.domain = value},
              ),
              SizedBox(
                height: breakSize,
              ),
              ElevatedButton(
                onPressed: () async {
                  await handshake();
                },
                child: Text('연결'),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
