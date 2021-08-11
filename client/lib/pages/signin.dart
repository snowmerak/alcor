import 'package:alcor/controllers/config.dart';
import 'package:alcor/controllers/signin.dart';
import 'package:alcor/service/signin.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:get/get.dart';

final String signinRoute = '/signin';

class SignIn extends StatefulWidget {
  const SignIn({Key? key}) : super(key: key);

  @override
  _SignInState createState() => _SignInState();
}

class _SignInState extends State<SignIn> {
  final double breakSize = 24;

  @override
  Widget build(BuildContext context) {
    var signinController = Get.put(SigninController());
    return Scaffold(
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
                    onChanged: (value) => {signinController.name = value},
                  ),
                ),
                ToggleButtons(
                  children: <Widget>[
                    Icon(Icons.male),
                    Icon(Icons.female),
                  ],
                  isSelected: signinController.gender,
                  onPressed: (index) => {
                    this.setState(() {
                      signinController.gender = signinController.gender.map((e) => false).toList();
                      signinController.gender[index] = true;
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
              onChanged: (value) => {signinController.phoneNumber = value},
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
              onChanged: (value) => {signinController.birth = value},
            ),
            SizedBox(
              height: breakSize,
            ),
            ElevatedButton(
              onPressed: () {
                enrollVoter(context);
                Get.defaultDialog(title: '요청 시도', middleText: '인증 번호를 요청하였습니다.');
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
              onChanged: (value) => {signinController.certCode = value},
            ),
            SizedBox(
              height: breakSize,
            ),
            ElevatedButton(
              onPressed: () => {signinController.signal.send(signinController.certCode)},
              child: Text('인증 번호 전송'),
            ),
          ],
        ),
      ),
    );
  }
}
