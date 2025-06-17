import psycopg2
import os
import sys

def apply_migrations(db_connection, migrations_folder):
    # Get migration files
    migration_files = [f for f in os.listdir(migrations_folder) 
                      if f.endswith('.sql') and os.path.isfile(os.path.join(migrations_folder, f))]
    
    if not migration_files:
        print("No migration files found in the directory")
        return

    # Sort files naturally (e.g., 001_init.sql, 002_table.sql)
    
    print(f"Found {len(migration_files)} migration files to execute")
    
    # Execute each migration file
    for filename in migration_files:
        filepath = os.path.join(migrations_folder, filename)
        print(f"Executing: {filename}")
        
        try:
            # Read SQL file content
            with open(filepath, 'r') as sql_file:
                sql = sql_file.read()
            
            # Execute SQL commands
            with db_connection.cursor() as cursor:
                cursor.execute(sql)
            db_connection.commit()
            print(f"Success: {filename}")
            
        except Exception as e:
            db_connection.rollback()
            print(f"ERROR executing {filename}: {str(e)}")
            print("Stopping migration process due to error")
            raise

if __name__ == "__main__":
    if len(sys.argv) != 3:
        print("Usage: python migrate.py <db_connection_string> <migrations_folder>")
        print("Example: python migrate.py \"host=localhost dbname=mydb user=postgres password=pass\" ./migrations")
        sys.exit(1)
    
    conn_str = sys.argv[1]
    migrations_dir = sys.argv[2]
    
    # Validate migrations directory
    if not os.path.isdir(migrations_dir):
        print(f"Error: Directory not found: {migrations_dir}")
        sys.exit(1)
    
    try:
        # Connect to PostgreSQL
        conn = psycopg2.connect(conn_str)
        print(f"Connected to database: {conn.dsn}")
        
        # Apply migrations
        apply_migrations(conn, migrations_dir)
        print("All migrations executed successfully")
        
    except psycopg2.Error as e:
        print(f"Database connection failed: {str(e)}")
        sys.exit(1)
    finally:
        if 'conn' in locals():
            conn.close()