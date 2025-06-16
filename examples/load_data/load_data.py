import json
import urllib.request
import urllib.error

# your API key
API_KEY = "GcKtft9raN8M6jtD3yV6Lgmw24kVdEfCF1y96Pl60Gn3f5mhDY6OsyTATHXf2SRSB1wqy7Wd2wiFyu2dwtlul9FQfaGTnyNfL2dsIxdEWXIOd02wkba/oXMVdhzVcJj/rQcCGgDKRk1wHJi5xXPRFFbbUicpHC6WSzX9ZwaOdkg"
COLLECTION = "dev_coll"

# API gateway URL
url = 'http://localhost:8080/api/v1/storage/{collection}?api-key={api_key}'
url = url.format(collection=COLLECTION, api_key=API_KEY)

# your data (or data parser function)
file_path = 'data.json'

# send request
try:
    with open(file_path, 'r', encoding='utf-8') as f:
        data = json.load(f)

    json_data = json.dumps(data).encode('utf-8')

    req = urllib.request.Request(url, data=json_data, method="POST")
    req.add_header('Content-Type', 'application/json')
    req.add_header('Accept', 'application/json')

    try:
        with urllib.request.urlopen(req) as response:
            response_text = response.read().decode('utf-8')
            print("Status:", response.getcode())
    except urllib.error.HTTPError as e:
        print("HTTP Error:", e.code)
        print("Body:", e.read().decode('utf-8'))

except FileNotFoundError:
    print(f"File {file_path} not found.")