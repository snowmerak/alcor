import requests
from edgepy.proto.protobuf.stats_pb2 import Data

def get_survey_last_number(uri: str) -> int:
    res = requests.get(f'http://{uri}:9999/survey/last_number')
    return int(res.content)

def get_survey_from_index(uri: str, idx: int) -> Data:
    survey = Data()
    res = requests.get(f'http://{uri}:9999/survey/{idx}')
    survey.ParseFromString(res.content)
    return survey
