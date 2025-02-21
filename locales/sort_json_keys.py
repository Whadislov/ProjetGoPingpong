import json
import os
print("Répertoire courant :", os.getcwd())

# Quick sort code 
input_file = 'locales/en.json'
output_file = 'locales/en.json'

with open(input_file, 'r', encoding='utf-8') as f:
    data = json.load(f)

sorted_data = {k: data[k] for k in sorted(data.keys())}

with open(output_file, 'w', encoding='utf-8') as f:
    json.dump(sorted_data, f, indent=4, ensure_ascii=False)

print(f"Sorted JSON done : {output_file}")