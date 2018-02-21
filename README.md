
## Install
1. Clone
```
git clone https://github.com/egig/resto-rest-api.git
```

2. Set GOPATH
```
cd resto-rest-api
export GOPATH="$PWD"
```

3. Create MySQL Database
4. Import file `resto-rest-api.sql` to your mysql database
5. Create file `config.json`, and edit it according to your system, see `config.json.example`


6. Install dependency
```
cd src/resto_rest_api
go get
```

## Run
```
go build
./resto-rest-api
```
then visit `http://localhost:8000/restaurants` thru your browser.

available endpoints:

GET   http://localhost:8000/restaurants
GET    http://localhost:8000/restaurants/:restaurant_id/reservations
POST http://localhost:8000/restaurants/:restaurant_id/reservations


## Test

Use following longitude and latitude for test get restaurants
```
-6.209786, 106.816454
-6.176488, 106.841203
-6.1949306 106.8172151
```
