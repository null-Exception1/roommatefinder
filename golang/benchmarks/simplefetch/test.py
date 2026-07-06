import psycopg2
import random
import string
from datetime import datetime
from faker import Faker

fake = Faker()

def random_hash(length=8):
    return ''.join(random.choices(string.ascii_lowercase + string.digits, k=length))

def seed_users(n=1000):
    conn = psycopg2.connect(
        dbname="test_db",
        user="devuser",
        password="devpass",
        host="localhost",
        port="5432"
    )
    cur = conn.cursor()

    for i in range(n):
        admn_hash = random_hash(12)
        name = fake.name()
        social = fake.user_name()
        socialtype = random.choice(["Discord", "Twitter", "Instagram"])
        roomno = random.randint(1, 200)
        blockno = random.randint(1, 20)
        created_at = datetime.now()

        cur.execute("""
            INSERT INTO people (admn_hash, name, social, socialtype, roomno, blockno, created_at)
            VALUES (%s, %s, %s, %s, %s, %s, %s)
        """, (admn_hash, name, social, socialtype, roomno, blockno, created_at))

    conn.commit()
    cur.close()
    conn.close()
    print(f"Seeded {n} users into people table.")

if __name__ == "__main__":
    seed_users(1000)
