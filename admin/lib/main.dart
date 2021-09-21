import 'package:admin/page/candidate_add.dart';
import 'package:admin/page/candidate_detail.dart';
import 'package:admin/page/candidate_list.dart';
import 'package:admin/page/index.dart';
import 'package:admin/page/menu.dart';
import 'package:flutter/material.dart';
import 'package:get/route_manager.dart';

void main() {
  runApp(App());
}

class App extends StatelessWidget {
  const App({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return GetMaterialApp(
      title: '관리 앱',
      initialRoute: '/',
      getPages: [
        GetPage(name: indexRoute, page: () => IndexPage()),
        GetPage(name: menuRoute, page: () => MenuPage()),
        GetPage(name: candidateAddRoute, page: () => CandidateAddPage()),
        GetPage(name: candidateListRoute, page: () => CandidateListPage()),
        GetPage(name: candidateDetailRoute, page: () => CandidateDetailPage()),
      ],
    );
  }
}
