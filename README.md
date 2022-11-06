# Backend build with golang


## Installation

clone this repository

```bash
git clone https://github.com/krifik/test_fullstack_frontend.git
```
cd to directory that been cloned
```bash
cd test_fullstack_backend
```
install depency
```
go mod tidy
```

## Usage
rename .env.example to .env and edit up to you

edit what you want
make sure every variables env in accordance to your system  

run migration
```
go run main.go db:migrate
```
run seeder
```
go run main.go db:seeder
```

exec in root dir and run
```
go run main.go
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
