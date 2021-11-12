## Mongo on docker

create mongo container:
`docker run -d  --name mongo-on-docker  -p 27888:27017 -e MONGO_INITDB_ROOT_USERNAME=mongoadmin -e MONGO_INITDB_ROOT_PASSWORD=secret mongo`

connection string: `mongodb://<username>:<password>@<host>:<port>/?authSource=admin`

use mongoshell
`sudo docker exec -it mongodb bash`

`mongo -u mongo -p monogoadmin --authenticationDatabase secret`