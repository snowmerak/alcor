import 'package:alcor/controllers/candidates.dart';
import 'package:alcor/controllers/config.dart';
import 'package:alcor/service/candidates.dart';
import 'package:alcor/service/vote.dart';
import 'package:flutter/material.dart';
import 'package:flutter_markdown/flutter_markdown.dart';
import 'package:get/get.dart';

final String candidateListRoute = '/candidate';

class Candidates extends StatelessWidget {
  const Candidates({Key? key}) : super(key: key);

  final double breakSize = 20;

  @override
  Widget build(BuildContext context) {
    var domain = Get.find<ConfigController>().domain;
    var candidatesController = Get.put(CandidatesController());
    initList(candidatesController, domain);
    return GetBuilder<CandidatesController>(
      builder: (_) => Scaffold(
        body: Container(
          padding: EdgeInsets.all(24),
          child: _.selected == ''
              ? Column(
                  crossAxisAlignment: CrossAxisAlignment.center,
                  children: [
                    TextField(
                      decoration: InputDecoration(
                        border: OutlineInputBorder(),
                        labelText: 'Search Name',
                      ),
                      onSubmitted: (value) => {_.updateInputName(value)},
                    ),
                    Expanded(
                      child: ListView.separated(
                        padding: EdgeInsets.all(8),
                        itemCount: _.filtered.length,
                        itemBuilder: (context, index) {
                          return TextButton(
                            child: Text(_.filtered[index]),
                            onPressed: () {
                              _.updateSelected(_.filtered[index], domain);
                            },
                          );
                        },
                        separatorBuilder: (BuildContext context, int index) {
                          return Divider();
                        },
                      ),
                    ),
                  ],
                )
              : Column(
                  children: [
                    Text(
                      _.detailName,
                      style: TextStyle(
                        fontWeight: FontWeight.bold,
                        fontSize: 24,
                      ),
                    ),
                    Expanded(
                      child: Markdown(data: _.markdown),
                    ),
                    Row(
                      children: [
                        ElevatedButton(onPressed: () {
                          vote(context);
                        }, child: Text('투표')),
                      ],
                    ),
                  ],
                ),
        ),
        floatingActionButton: _.isDetail
            ? FloatingActionButton(
                onPressed: () {
                  _.resetSelected();
                },
                child: Icon(Icons.arrow_back),
              )
            : FloatingActionButton(
                onPressed: () {
                  initList(candidatesController, domain);
                },
                child: Icon(Icons.wifi_protected_setup),
              ),
      ),
    );
  }
}
