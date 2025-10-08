import requests
import sys
import json

if len(sys.argv) < 2:
    print("No IDs specified")
    sys.exit(1)

for id in sys.argv[1:]:
    r = requests.get(f"https://www.ecs.soton.ac.uk/people/{id}")
    if r.status_code != 200:
        print(f"Invalid status code ({r.status_code}) when trying to query {id}\n")
        continue

    if "appear in the public directory" in r.text:
        print(f"ID {id} does not exist\n")
        continue

    user = json.loads(r.text.split("}\n        },\n        ")[1].split("]")[0])

    print(f"""Name: {user["name"]}
Email: {user["email"]}
Description: {user["description"]}
Telephone: {user["telephone"]}
""")
