# BOP User Service Mock

Provides a mocked set of endpoints backed by a simple Go server and JSON data, to mimic responses from the back-office proxy to use locally.

### Usage

```
go run insights-services.go
```
Visit any of the supported endpoints below, now available on `http://localhost:8090/`:

### Supported Endpoints

  - Users for an account: `http://localhost:8090/insights-services/v1/accounts/:id/users` - note that the account ID does not matter in this case.

### Future Endpoints

  - Details for a user: `http://localhost:8090/insights-services/v1/auth`

### Generating Fake Data

Populate `data/users.json` with the following:

```
$ ./bin/populate_users.rb
```