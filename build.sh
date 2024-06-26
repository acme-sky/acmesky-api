read -p "POSTGRES_USER [acme]: " pg_user
pg_user=${pg_user:-'acme'}

read -p "POSTGRES_PASSWORD [pass]: " pg_pass
pg_pass=${pg_pass:-'pass'}

read -p "POSTGRES_DB [db]: " pg_db
pg_db=${pg_db:-'db'}

read -p "JWT_TOKEN [un1b0!!Tok3n]: " jwt_token
jwt_token=${jwt_token:-'un1b0!!Tok3n'}

read -p "RABBITMQ [amqp://guest:guest@rabbit:5672/]: " rabbitmq
rabbitmq=${rabbitmq:-'amqp://guest:guest@rabbit:5672/'}

read -p "SERVER_URL [0.0.0.0:8080]: " server_url
server_url=${server_url:-'0.0.0.0:8080'}

export POSTGRES_USER="$pg_user"
export POSTGRES_PASSWORD="$pg_pass"
export POSTGRES_DB="$pg_db"
export DATABASE_DSN="host=acmesky-postgres user=$pg_user password=$pg_pass dbname=$pg_db port=5432"
export JWT_TOKEN="$jwt_token"
export RABBITMQ="$rabbitmq"
export SERVER_URL="$server_url"

docker build -t acmesky-api .

docker compose up
