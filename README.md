### asagidaki kimi postgres elave etdim
docker run --name=auth2-db -e POSTGRES_PASSWORD='fuadfuad' -p 5437:5432 -d --rm postgres

### sekret key for token
export SECRET_KEY=supersecretkey

### db password
export DB_PASSWORD=123456

### productionda connection.go host adi go-auth_db_1 / development-de localhost