## Install
1. Clone
```
git clone https://github.com/egig/resto-rest-api.git
```
2. Install dependency
```
cd resto-rest-api
git get
```
3. Create MySQL Database
4. Import file `resto-rest-api.sql`
5. Edit file `config.json`, according to your system

## Run
```
go build && ./resto-rest-api
```
then visit `http://localhost:8000` thru your browser.

## Test

Use following longitude and latitude for test
```
-6.209786, 106.816454
-6.176488, 106.841203
-6.1949306 106.8172151
```