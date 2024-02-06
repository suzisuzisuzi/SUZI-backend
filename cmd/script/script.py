import requests
import json
import random

url = "https://suzi-backend.onrender.com/entries/apjDVdoq7ZNpqzDXL8ELsR8uhOT2"

body = {
    "Timestamp": "2022-03-01T00:00:00Z",
    "FirebaseID": "apjDVdoq7ZNpqzDXL8ELsR8uhOT2",
    "Latitude": 0,
    "Longitude": 0,
    "Altitude": 0.0,
    "Category": "test",
    "Rating": 1
}

# west bengal coordinates lol
sw = [21.53, 85.82]  
ne = [27.22, 89.83] 
nw = [21.53, 89.83]  
se = [27.22, 85.82]  

num_points = 2000

for i in range(num_points):
    body["Latitude"] = sw[0] + (ne[0] - sw[0]) * i / num_points
    body["Longitude"] = sw[1] + (ne[1] - sw[1]) * i / num_points
    body["Rating"] = random.uniform(-1, 1)
    response = requests.post(url, data=json.dumps(body), headers={'Content-Type': 'application/json'})
    print(response.status_code)

for i in range(num_points):
    body["Latitude"] = nw[0] + (se[0] - nw[0]) * i / num_points
    body["Longitude"] = nw[1] + (se[1] - nw[1]) * i / num_points
    body["Rating"] = random.uniform(-1, 1)
    response = requests.post(url, data=json.dumps(body), headers={'Content-Type': 'application/json'})
    print(response.status_code)