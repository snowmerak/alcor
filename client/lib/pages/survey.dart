import 'package:alcor/controllers/survey.dart';
import 'package:alcor/service/survay.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:get/get.dart';

final surveyRoute = '/survey';

class SurveyPage extends StatelessWidget {
  const SurveyPage({Key? key}) : super(key: key);

  final double breakSize = 24;
  final double fontSize = 16;

  @override
  Widget build(BuildContext context) {
    Get.put(SurveyController());
    getRegions();
    getJobs();
    getEducations();
    getHouseTypes();
    getIdeologys();
    return Scaffold(
      appBar: AppBar(
        title: Text('설문 조사'),
      ),
      body: SingleChildScrollView(
        child: Container(
          padding: EdgeInsets.all(24),
          child: GetBuilder<SurveyController>(
            builder: (controller) => Column(
              mainAxisAlignment: MainAxisAlignment.start,
              children: [
                SizedBox(
                  height: breakSize,
                ),
                TextField(
                  keyboardType: TextInputType.number,
                  maxLength: 4,
                  inputFormatters: [
                    FilteringTextInputFormatter.digitsOnly,
                  ],
                  decoration: InputDecoration(
                    labelText: '귀하의 출생 연도가 언제입니까?',
                  ),
                  onChanged: (value) async {
                    var i = int.parse(value);
                    if (0 <= i && i <= DateTime.now().year) {
                      controller.survey.age = i;
                    } else {
                      await Get.defaultDialog(title: '출생 연도는 올해를 넘을 수 없습니다.');
                    }
                  },
                ),
                SizedBox(
                  height: breakSize,
                ),
                TextField(
                  keyboardType: TextInputType.number,
                  maxLength: 2,
                  inputFormatters: [
                    FilteringTextInputFormatter.digitsOnly,
                  ],
                  decoration: InputDecoration(
                    labelText: '귀하의 임금 수준은 상위 몇%라고 생각하십니까?',
                  ),
                  onChanged: (value) async {
                    var i = int.parse(value);
                    if (0 <= i && i <= 99) {
                      controller.survey.salary = i;
                    } else {
                      await Get.defaultDialog(
                          title: '최소 0%에서 최대 99%까지 입력할 수 있습니다.');
                    }
                  },
                ),
                SizedBox(
                  height: breakSize,
                ),
                Row(
                  children: [
                    Text(
                      '귀하는 어느 지역에 거주하고 계십니까?',
                      style: TextStyle(
                        fontSize: fontSize,
                      ),
                    ),
                    PopupMenuButton<String>(
                      onSelected: (value) {
                        controller.survey.region = value;
                        controller.update();
                      },
                      itemBuilder: (context) => controller.regions
                          .map((e) => PopupMenuItem(
                                child: Text(e),
                                value: e,
                              ))
                          .toList(),
                    ),
                    Text(
                      controller.survey.region,
                      style: TextStyle(
                        fontSize: fontSize,
                      ),
                    ),
                  ],
                ),
                SizedBox(
                  height: breakSize,
                ),
                Row(
                  children: [
                    Text(
                      '귀하는 어떤 직업을 가지고 계십니까?',
                      style: TextStyle(
                        fontSize: fontSize,
                      ),
                    ),
                    PopupMenuButton<String>(
                      onSelected: (value) {
                        controller.survey.job = value;
                        controller.update();
                      },
                      itemBuilder: (context) => controller.jobs
                          .map((e) => PopupMenuItem(
                                child: Text(e),
                                value: e,
                              ))
                          .toList(),
                    ),
                    Text(
                      controller.survey.job,
                      style: TextStyle(
                        fontSize: fontSize,
                      ),
                    ),
                  ],
                ),
                SizedBox(
                  height: breakSize,
                ),
                Row(
                  children: [
                    Text(
                      '최종 학력은 어느 정도 입니까?',
                      style: TextStyle(
                        fontSize: fontSize,
                      ),
                    ),
                    PopupMenuButton<String>(
                      onSelected: (value) {
                        controller.survey.education = value;
                        controller.update();
                      },
                      itemBuilder: (context) => controller.educations
                          .map((e) => PopupMenuItem(
                                child: Text(e),
                                value: e,
                              ))
                          .toList(),
                    ),
                    Text(
                      controller.survey.education,
                      style: TextStyle(
                        fontSize: fontSize,
                      ),
                    ),
                  ],
                ),
                SizedBox(
                  height: breakSize,
                ),
                Row(
                  children: [
                    Text(
                      '귀하는 결혼 이력이 있습니까?',
                      style: TextStyle(
                        fontSize: fontSize,
                      ),
                    ),
                    PopupMenuButton<bool>(
                      onSelected: (value) {
                        controller.survey.married = value;
                        controller.update();
                      },
                      itemBuilder: (context) => [
                        PopupMenuItem(
                          child: Text('예'),
                          value: true,
                        ),
                        PopupMenuItem(
                          child: Text('아니오'),
                          value: false,
                        ),
                      ],
                    ),
                    Text(
                      controller.survey.married ? '예' : '아니오',
                      style: TextStyle(
                        fontSize: fontSize,
                      ),
                    ),
                  ],
                ),
                SizedBox(
                  height: breakSize,
                ),
                Row(
                  children: [
                    Text(
                      '귀하는 이혼 이력이 있습니까?',
                      style: TextStyle(
                        fontSize: fontSize,
                      ),
                    ),
                    PopupMenuButton<bool>(
                      onSelected: (value) {
                        controller.survey.divorced = value;
                        controller.update();
                      },
                      itemBuilder: (context) => [
                        PopupMenuItem(
                          child: Text('예'),
                          value: true,
                        ),
                        PopupMenuItem(
                          child: Text('아니오'),
                          value: false,
                        ),
                      ],
                    ),
                    Text(
                      controller.survey.divorced ? '예' : '아니오',
                      style: TextStyle(
                        fontSize: fontSize,
                      ),
                    ),
                  ],
                ),
                SizedBox(
                  height: breakSize,
                ),
                Row(
                  children: [
                    Text(
                      '귀하의 거주지는 어떤 형태입니까?',
                      style: TextStyle(
                        fontSize: fontSize,
                      ),
                    ),
                    PopupMenuButton<String>(
                      onSelected: (value) {
                        controller.survey.houseType = value;
                        controller.update();
                      },
                      itemBuilder: (context) => controller.houseTypes
                          .map((e) => PopupMenuItem(
                                child: Text(e),
                                value: e,
                              ))
                          .toList(),
                    ),
                    Text(
                      controller.survey.houseType,
                      style: TextStyle(
                        fontSize: fontSize,
                      ),
                    ),
                  ],
                ),
                SizedBox(
                  height: breakSize,
                ),
                Row(
                  children: [
                    Text(
                      '귀하는 자차를 소유하고 있습니까?',
                      style: TextStyle(
                        fontSize: fontSize,
                      ),
                    ),
                    PopupMenuButton<bool>(
                      onSelected: (value) {
                        controller.survey.hasCar = value;
                        controller.update();
                      },
                      itemBuilder: (context) => [
                        PopupMenuItem(
                          child: Text('예'),
                          value: true,
                        ),
                        PopupMenuItem(
                          child: Text('아니오'),
                          value: false,
                        ),
                      ],
                    ),
                    Text(
                      controller.survey.hasCar ? '예' : '아니오',
                      style: TextStyle(
                        fontSize: fontSize,
                      ),
                    ),
                  ],
                ),
                SizedBox(
                  height: breakSize,
                ),
                Row(
                  children: [
                    Text(
                      '귀하는 빚이 있으십니까?',
                      style: TextStyle(
                        fontSize: fontSize,
                      ),
                    ),
                    PopupMenuButton<bool>(
                      onSelected: (value) {
                        controller.survey.hasDebt = value;
                        controller.update();
                      },
                      itemBuilder: (context) => [
                        PopupMenuItem(
                          child: Text('예'),
                          value: true,
                        ),
                        PopupMenuItem(
                          child: Text('아니오'),
                          value: false,
                        ),
                      ],
                    ),
                    Text(
                      controller.survey.hasDebt ? '예' : '아니오',
                      style: TextStyle(
                        fontSize: fontSize,
                      ),
                    ),
                  ],
                ),
                SizedBox(
                  height: breakSize,
                ),
                Row(
                  children: [
                    Text(
                      '귀하는 어떤 정치적 사상에 가깝습니까?',
                      style: TextStyle(
                        fontSize: fontSize,
                      ),
                    ),
                    PopupMenuButton<String>(
                      onSelected: (value) {
                        controller.survey.ideology = value;
                        controller.update();
                      },
                      itemBuilder: (context) => controller.ideologys
                          .map((e) => PopupMenuItem(
                                child: Text(e),
                                value: e,
                              ))
                          .toList(),
                    ),
                    Text(
                      controller.survey.ideology,
                      style: TextStyle(
                        fontSize: fontSize,
                      ),
                    ),
                  ],
                ),
                SizedBox(
                  height: breakSize,
                ),
                Text(
                  '설문조사에 응해주셔서 감사합니다.',
                  style: TextStyle(
                    fontSize: fontSize,
                  ),
                ),
                SizedBox(
                  height: breakSize,
                ),
                ElevatedButton(
                  onPressed: () => submitSurvey(),
                  child: Text('제출'),
                ),
                SizedBox(
                  height: breakSize,
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
