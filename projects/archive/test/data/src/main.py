import psycopg2
import uuid
import string
import random

def main():
    conn = psycopg2.connect(
        host="host.docker.internal",
        database="postgres",
        user="postgres",
        password="postgres"
    )

    # create a cursor
    cur = conn.cursor()
    
    # execute a statement
    print('PostgreSQL database version:')
    cur.execute('SELECT version()')

    # display the PostgreSQL database server version
    db_version = cur.fetchone()
    print(db_version)

    insert_statement = """INSERT INTO public."test-data"(
                id, name)
	            VALUES (%s, %s);"""

    index = 0
    while index < 1000000:
        user_id = str(uuid.uuid4())
        user_name =  ''.join(random.choice(string.ascii_lowercase) for i in range(10))
        
        cur.execute(insert_statement, (user_id, user_name))
        index+=1

    conn.commit()
    cur.close()
    print("Done.")




if __name__ == '__main__':
    main()