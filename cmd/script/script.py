import requests
import json

url = "https://suzi-backend.onrender.com/entries/apjDVdoq7ZNpqzDXL8ELsR8uhOT2"

body = {
    "Timestamp": "2022-03-01T00:00:00Z",
    "FirebaseID": "apjDVdoq7ZNpqzDXL8ELsR8uhOT2",
    "Latitude": 19.0760,
    "Longitude": 72.8777,
    "Altitude": 0.0,
    "Category": "test",
    "Rating": 1
}

for _ in range(1000):
        body["Latitude"] += 0.000001
        body["Longitude"] += 0.000001
        response = requests.post(url, data=json.dumps(body), headers={'Content-Type': 'application/json'})
        print(response.status_code)