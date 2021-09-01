from edgepy.proto.protobuf.stats_pb2 import Data
import edgedb # must use edgedb==0.16.0

id = 'merak'
pw = 'test'
uri = 'voters'
port = '16419'

def connect(id: str, pw: str, uri: str, port: str) -> edgedb.BlockingIOConnection:
    edgedb.connect(
        f'edgedb://{id}:{pw}@{uri}:{port}')

def insert_survey(conn: edgedb.BlockingIOConnection, survey: Data):
	return conn.execute('''
	insert Data {
		number := <int64>$number,
		candidate := <str>$candidate,
		gender := <bool>$gender,
		age := <int16>$age,
		region := <str>$region,
		job := <str>$job,
		education := <str>$education,
		married := <bool>$married,
		divorced := <bool>$divorced,
		has_car := <bool>$has_car,
		house_type := <str>$house_type,
		salary := <int16>$salary,
		has_debt := <bool>$has_debt,
		ideology := <str>$ideology,
	}
	''', number=survey.Number, candidate=survey.Number, gender=survey.Gender, age=survey.Age, region=survey.Region, job=survey.Job, education=survey.Education, married=survey.Married, divorced=survey.Divorced, has_car=survey.HasCar, house_type=survey.HouseType, salary=survey.Salary, has_debt=survey.HasDebt, ideology=survey.Ideology)

def select_survey(conn: edgedb.BlockingIOConnection, survey_id: int):
	return conn.query_one('''
    select Data {
		candidate,
		gender,
		age,
		region,
		job,
		education,
		married,
		divorced,
		has_car,
		house_type,
		salary,
		has_debt,
		ideology,
	}
	filter Data.number = <int64>$number
	limit 1
    ''', number=survey_id)

def select_numbers(conn: edgedb.BlockingIOConnection):
	return conn.query('''
	select Data {
		number,
	}
	''')