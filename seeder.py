import psycopg2
import random
import string
from datetime import datetime
import uuid

# Establish a connection to the database
conn_client = psycopg2.connect(
    dbname="simtix-client-db",
    user="simtix",
    password="simtix",
    host="localhost",
    port="8001"
)
conn_ticket = psycopg2.connect(
    dbname="simtix-ticketing",
    user="simtix",
    password="simtix",
    host="localhost",
    port="8006"
)

cur_client = conn_client.cursor()
cur_ticket = conn_ticket.cursor()

# Insert 5 users into the 'users' table
for i in range(5):
    email = f'user{i+1}@example.com'
    name = f'User {i+1}'
    password = '$2b$10$fBFDVg/UXudJ82miGq3c5OSc.NQKuGTFVWUxSpNzx19KiABm6TJ0i'

    cur_client.execute(
        "INSERT INTO users (email, name, password) VALUES (%s, %s, %s)",
        (email, name, password)
    )

# Insert 3 events into the 'events' table, each with 8 seats
from datetime import datetime

for i in range(3):
    event_id = str(uuid.uuid4())  # Generate a unique id
    event_name = f'Event {i+1}'
    event_time = f"{i+1} Mar 24 08:00 SGT"

    cur_ticket.execute(
        "INSERT INTO events (id, event_name, event_time, created_at, updated_at) VALUES (%(id)s, %(name)s, %(time)s, %(created)s, %(updated)s)",
        {'id': event_id, 'name': event_name, 'time': event_time, 'created': datetime.now(), 'updated': datetime.now()}
    )

    for j in range(8):
        seat_id = str(uuid.uuid4())
        status = 'OPEN'
        seat_row = 'A'
        seat_number = str(j+1)
        price = str((j+1) * 100000)

        cur_ticket.execute(
            "INSERT INTO seats (id, event_id, status, seat_row, seat_number, price, created_at, updated_at) VALUES (%(id)s, %(event_id)s, %(status)s, %(row)s, %(number)s, %(price)s, %(created)s, %(updated)s)",
            {'id': seat_id, 'event_id': event_id, 'status': status, 'row': seat_row, 'number': seat_number, 'price': price, 'created': datetime.now(), 'updated': datetime.now()}
        )


conn_client.commit()
cur_client.close()
conn_client.close()

conn_ticket.commit()
cur_ticket.close()
conn_ticket.close()
