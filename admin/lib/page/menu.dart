import 'package:admin/page/candidate_add.dart';
import 'package:admin/page/candidate_list.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';

final menuRoute = '/menu';

class MenuPage extends StatelessWidget {
  const MenuPage({Key? key}) : super(key: key);

  final double breaksize = 24;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('메뉴'),
      ),
      body: Container(
        padding: EdgeInsets.all(24),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.stretch,
          children: [
            ElevatedButton(
              onPressed: () => Get.toNamed(candidateAddRoute),
              child: Text('후보 추가'),
            ),
            SizedBox(
              height: breaksize,
            ),
            ElevatedButton(
              onPressed: () => Get.toNamed(candidateListRoute),
              child: Text('후보 수정 및 삭제'),
            ),
            SizedBox(
              height: breaksize,
            ),
          ],
        ),
      ),
    );
  }
}
