import 'dart:typed_data';

import 'package:alcor/controllers/config.dart';
import 'package:alcor/pages/register.dart';
import 'package:alcor/service/handshake.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:hive/hive.dart';

import 'candidate.dart';

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
                  await handshake();
                  await Hive.openBox<Uint8List>(configController.domain);
                  var id = Hive.box<Uint8List>(configController.domain).get('id');
                  if (id == null) {
                    Get.toNamed(registerRoute);
                  } else {
                    Get.toNamed(candidateRoute);
                  }
                  Get.toNamed(registerRoute);
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
