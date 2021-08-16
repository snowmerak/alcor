import 'package:alcor/proto/protobuf/vote.pb.dart';
import 'package:channel/channel.dart';
import 'package:get/get.dart';

class VoteController extends GetxController {
  Paper paper = Paper();
  Channel<bool> receiveKey = Channel<bool>();
}