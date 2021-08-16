import 'package:alcor/controllers/register.dart';
import 'package:alcor/service/register.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:get/get.dart';

final String registerRoute = '/register';

class RegisterPage extends StatefulWidget {
  const RegisterPage({Key? key}) : super(key: key);

  @override
  _RegisterPageState createState() => _RegisterPageState();
}

class _RegisterPageState extends State<RegisterPage> {
  final double breakSize = 24;

  @override
  Widget build(BuildContext context) {
    var registerController = Get.put(RegisterController());
    return Scaffold(
      appBar: AppBar(
        title: Text('투표권 등록'),
        automaticallyImplyLeading: false,
      ),
      body: Container(
        padding: EdgeInsets.all(24),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.center,
          children: [
            Row(
              mainAxisSize: MainAxisSize.min,
              mainAxisAlignment: MainAxisAlignment.start,
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Expanded(
                  child: TextField(
                    decoration: InputDecoration(
                      border: OutlineInputBorder(),
                      labelText: 'Name',
                    ),
                    onChanged: (value) => {registerController.name = value},
                  ),
                ),
                ToggleButtons(
                  children: <Widget>[
                    Icon(Icons.male),
                    Icon(Icons.female),
                  ],
                  isSelected: registerController.gender,
                  onPressed: (index) => {
                    this.setState(() {
                      registerController.gender = registerController.gender.map((e) => false).toList();
                      registerController.gender[index] = true;
                    })
                  },
                ),
              ],
            ),
            SizedBox(
              height: breakSize,
            ),
            TextField(
              keyboardType: TextInputType.phone,
              inputFormatters: [
                FilteringTextInputFormatter.digitsOnly,
              ],
              decoration: InputDecoration(
                  border: OutlineInputBorder(), labelText: 'Phone'),
              onChanged: (value) => {registerController.phoneNumber = value},
            ),
            SizedBox(
              height: breakSize,
            ),
            TextField(
              keyboardType: TextInputType.number,
              inputFormatters: [
                FilteringTextInputFormatter.digitsOnly,
              ],
              decoration: InputDecoration(
                  border: OutlineInputBorder(), labelText: 'Birthday (YYMMDD)'),
              onChanged: (value) => {registerController.birth = value},
            ),
            SizedBox(
              height: breakSize,
            ),
            ElevatedButton(
              onPressed: () {
                registerCall(context);
              },
              child: Text('인증 번호 요청'),
            ),
            SizedBox(
              height: breakSize,
            ),
            TextField(
              keyboardType: TextInputType.number,
              inputFormatters: [
                FilteringTextInputFormatter.digitsOnly,
              ],
              decoration: InputDecoration(
                  border: OutlineInputBorder(),
                  labelText: 'Certification Code'),
              onChanged: (value) => {registerController.certCode = value},
            ),
            SizedBox(
              height: breakSize,
            ),
            ElevatedButton(
              onPressed: () => {registerCert(context)},
              child: Text('인증 번호 전송'),
            ),
          ],
        ),
      ),
    );
  }
}
