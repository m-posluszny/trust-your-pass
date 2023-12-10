import requests
import random

a = requests.post("http://localhost:10000/api/v1/passwords",json={"password":f"{random.randint(10000,100000)}"})
print(a.text)
print(a.json())
id = a.json()['id']
print(id)
print(requests.get("http://localhost:10000/api/v1/passwords/"+id).json())