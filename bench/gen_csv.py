"""Generates the benchmark file used in the README: 15M rows, ~1.3 GB, seeded so it is reproducible.

    python3 bench/gen_csv.py complaints.csv
    echo '{"label":"state","words":["Jalisco","Nuevo León"]}' > keywords.json
    /usr/bin/time -v hankotsu complaints.csv keywords.json
"""
import random
import sys

random.seed(7)
states = ["Aguascalientes", "Baja California", "Chiapas", "Chihuahua", "Ciudad de México", "Coahuila",
          "Guanajuato", "Guerrero", "Jalisco", "Estado de México", "Michoacán", "Nuevo León", "Oaxaca",
          "Puebla", "Querétaro", "Quintana Roo", "Sinaloa", "Sonora", "Tabasco", "Veracruz", "Yucatán",
          "Zacatecas"]
cats = ["telecom", "banking", "retail", "energy", "transport", "health", "insurance", "education"]
words = "service charge delay refund contract signal outage invoice branch account claim delivery".split()

with open(sys.argv[1] if len(sys.argv) > 1 else "complaints.csv", "w", buffering=1 << 22) as f:
    f.write("id,date,state,category,amount,description\n")
    for i in range(15_000_000):
        f.write(f"{i},2025-{random.randint(1,12):02d}-{random.randint(1,28):02d},{random.choice(states)},"
                f"{random.choice(cats)},{random.randint(50,90000)},{' '.join(random.choices(words, k=6))}\n")
