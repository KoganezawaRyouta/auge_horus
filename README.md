# devepolments
## go install

```
gvm install go1.7.1
gvm use go1.7.1
```

## app set up
```
make deps
make build
```

## db set up
```
bundle exec rake db:create or mysql -u root < db/create_database.sql
bundle exec rake db:migrate
```

## make other commands
```
make help
```

## import rate of bit coint by coin check
```
make run_import
```

## run api server
```
make run_server
```

##  server url

* http://localhost:8080/monitor
* http://localhost:8080/tickers
* http://localhost:8080/trades

## ec2 provisioning
@see https://github.com/KoganezawaRyouta/private_ansible
