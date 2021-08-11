import 'dart:typed_data';

import 'package:alcor/controllers/config.dart';
import 'package:alcor/pages/candidate.dart';
import 'package:alcor/pages/signin.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:hive/hive.dart';

final String indexRoute = '/';

class Index extends StatelessWidget {
  const Index({Key? key}) : super(key: key);

  final double breakSize = 24;

  @override
  Widget build(BuildContext context) {
    final configController = Get.put(ConfigController());
    return Scaffold(
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
                  await Hive.openBox<Uint8List>(configController.domain);
                  var id = Hive.box<Uint8List>(configController.domain).get('id');
                  if (id == null) {
                    Get.toNamed(signinRoute);
                  } else {
                    Get.toNamed(candidateListRoute);
                  }
                },
                child: Text('저장'),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
