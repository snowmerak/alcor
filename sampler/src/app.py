import requests
import proto.protobuf.stats_pb2
import os, csv

uri = "127.0.0.1"

last_number = int(requests.get(f"http://{uri}:9999/survey/lastnumber").text)

def convertBool(b: bool) -> str:
    if b == True:
        return "true"
    return "false"

with open("./elk-final/vote_result.csv", "w", encoding="utf-8", newline="") as f:
    wr = csv.writer(f)
    wr.writerow(["candidate", "gender", "age", "region", "job", "education", "married", "divorced", "has_car" , "house_type" , "salary" ,"has_debt", "ideology"])

    year = 2021

    for i in range(last_number):
        r = requests.get(f"http://{uri}:9999/survey/{i}")
        survey = proto.protobuf.stats_pb2.Data()
        survey.ParseFromString(r.content)
        wr.writerow([survey.Candidate,convertBool(survey.Gender),year-survey.Age,survey.Region,survey.Job,survey.Education,convertBool(survey.Married),convertBool(survey.Divorced),convertBool(survey.HasCar),survey.HouseType,survey.Salary,convertBool(survey.HasDebt),survey.Ideology])

