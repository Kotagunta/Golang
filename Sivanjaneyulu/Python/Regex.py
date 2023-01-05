import re

string =  "AC*wv12n/:#e1234we2.45oin  (fwoi6n#a98nfwb+owi"

get_digits = re.findall(r'[\d\.\d]+', string)

result = []

for digit in get_digits:
    result.append(eval(digit))
    
print(sorted(result))