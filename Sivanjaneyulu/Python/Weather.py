import requests

api_address = 'http://api.openweathermap.org/data/2.5/weather?appid=02af5f60eb01969c48620324affb8724&q='

city = input("City name: ")

url = api_address + city

response = requests.get(url).json()

print(response)