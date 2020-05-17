# Thrift GraphQL Demo

**This repo is in my [blog post](https://www.nisheetsinvhal.com/microservices-thrift-nodejs-golang-graphql/).**

This repo has 3 services

- service-gql
  - Written in Nodejs
  - Uses Thrift clients to communicate to other services
  - Apollo GraphQL server
- service-post
  - Written in Golang
  - Thrift server for Posts data
- service-user
  - Written in Golang
  - Thrift server for User data

Each service has a `/deployments` folder which contains a _Dockerfile_ to build the image.

# Get started

## Kubernetes

In the root of the directory and `kubectl` configured for the kubernetes cluster

```shell
$ kubectl apply -R -f k8s/
```

## Docker Compose

In the root of the directory and having `docker-compose` installed

```shell
$ docker-compose up
```

## Native

In each of the service-post and service-user folder

```shell
$ go mod download
$ go run cmd/thrift/main.go -addr :<port>
```

By default the `port` is 9090 so be sure to change them for any one service

In the service-gql folder

```shell
$ npm install
```

Create `.env` file in the folder and add the following lines

```text
GRAPHQL_PORT=3000
SERVICE_USER_HOST=localhost
SERVICE_USER_PORT=9090
SERVICE_POST_HOST=localhost
SERVICE_POST_PORT=9091
```

Or, alternatively - set the values as environment variables

And finally

```shell
$ npm start
```

### Check

```shell
$ curl -H 'Content-Type: application/json' -X POST -d '{"query": "query { ping { message } postPing { message } userPing { message  } }"}' http://localhost:3000/graphql
{"data":{"ping":{"message":"ping"},"postPing":{"message":"ping"},"userPing":{"message":"ping"}}}
```

You can see the logs accordingly for each service

## Modifying Thrift

You can modify the `.thrift` files in any service-{user,post} and be sure to copy them over to service-gql.
Then simply run the following service in the folder.

```shell
$ ./build/thrift-gen.sh
```

Don't forget to implement the changes!

Don't forget to change the schema!

## Future

- [ ] Further implementation
  - [ ] CRUD for service-user
  - [ ] CRUD for service-post
  - [ ] Healthchecks for all services
- [ ] Istio
  - [ ] Make another image with `_version` as 2 (Enables Traffic Shifting, etc)
  - [ ] Metrics
