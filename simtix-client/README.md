## Overview
Client service.

## API Docs
### REST API
| HTTP Method | Endpoint    | Description |
|-------------|-------------|-------------|
| GET         | /v1/example |             |

## How To Start
### Migrations

#### Generate Migration
```bash
npm run migration:generate -d src/migrations/{migration-name}
```

#### Run Migration
```bash
npm run migration:run
```

### Run the Service
```bash
npm run start:dev
```