import 'package:flutter/material.dart';

class Enroll extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Column(children: [
      Text(
        'Input Account',
        style: TextStyle(
          fontWeight: FontWeight.bold,
          fontSize: 24,
        ),
      ),
      TextField(
        decoration: InputDecoration(
          border: OutlineInputBorder(),
          labelText: 'Name',
        ),
      ),
      OutlinedButton(
        onPressed: () => {},
        child: Text('Submit'),
      )
    ]);
  }
}
