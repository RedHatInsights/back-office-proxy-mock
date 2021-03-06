# BOP User Service Mock

Provides a mocked set of endpoints backed by a simple Go server and JSON data, to mimic responses from the back-office proxy to use locally.

### Usage

```
go run insights-services.go
```
Visit any of the supported endpoints below, now available on `http://localhost:8090/`:

### Supported Endpoints

  - Status:
    - `http://localhost:8090/insights-services/` - nothing is gonna be resulted
  - Users for an account (V1 endpoint):
    - `http://localhost:8090/insights-services/v1/accounts/:id/users` - note that the account ID does not matter in this case.
  - Users for an account (V2 endpoint):
    - `http://localhost:8090/insights-services/v2/accounts/:id/users` - note that the account ID does not matter in this case.
  - Details for a user: 
    - `http://localhost:8090/insights-services/v1/auth`

### Future Endpoints

  - https://backoffice-proxy-insights-services.ext.us-east.aws.preprod.paas.redhat.com/docs/#/

### Generating Fake Data

User data exists by default in `data/{files}.json`, but to populate new/different,
use the following:

```
$ gem install faker
$ ./bin/populate_{name}.rb
```

### Integrating with RBAC

One usecase is to have this mock running for RBAC to use locally. Setting the
following in the `.env` file in [RBAC](https://github.com/RedHatInsights/insights-rbac/)
will allow for this:

```
PRINCIPAL_PROXY_SERVICE_PROTOCOL=http
PRINCIPAL_PROXY_SERVICE_HOST=localhost
PRINCIPAL_PROXY_SERVICE_PORT=8090
PRINCIPAL_PROXY_SERVICE_PATH=/insights-services
PRINCIPAL_PROXY_SERVICE_SSL_VERIFY=False
```

Currently this will support being able to use the `/api/rbac/v1/principals/`
GET endpoint.
