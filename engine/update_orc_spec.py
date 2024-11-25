import json
import urllib.request

url = "https://data.orc.org/public/WPub.dll?action=DownBoatRMS&SailNo=GRE-016&Family=DH&ext=json"

with urllib.request.urlopen(url) as response:
  data = json.load(response)


ships = data["rms"]

if isinstance(ships, list):
  print(len(ships))
  print(json.dumps(ships, indent=2))