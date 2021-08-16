import 'dart:typed_data';

import 'package:alcor/pages/candidate.dart';
import 'package:alcor/pages/index.dart';
import 'package:alcor/pages/signin.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:get/get_navigation/src/root/get_material_app.dart';
import 'package:hive_flutter/hive_flutter.dart';

Future<void> main() async {
  await Hive.initFlutter();
  await Hive.deleteFromDisk();
  await Hive.openBox<Uint8List>('account');
  await Hive.openBox<Uint8List>('config');
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return GetMaterialApp(
      title: 'ALCOR',
      initialRoute: '/',
      getPages: [
        GetPage(name: indexRoute, page: () => Index()),
        GetPage(name: signinRoute, page: () => SignIn()),
        GetPage(name: candidateListRoute, page: () => Candidates()),
      ],
    );
  }
}
