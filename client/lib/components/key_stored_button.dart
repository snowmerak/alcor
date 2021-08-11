import 'package:flutter/material.dart';

class StringStoredButton extends TextButton {
  String data = '';

  StringStoredButton(Widget child, Function() onPressed, String data): super(child: child, onPressed: onPressed) {
    this.data = data;
  }
}