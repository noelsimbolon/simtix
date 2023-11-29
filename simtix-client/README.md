## Overview
Client service.

## API Docs
### REST API
| HTTP Method | Endpoint                       | Description       |
|-------------|--------------------------------|-------------------|
| GET         | /api/client                    | Health Check      |
| GET         | /api/client/users              | Get User          |
| PUT         | /api/client/users              | Update User       |
| DELETE      | /api/client/users              | Delete User       |
| POST        | /api/client/auth/login         | Login             |
| POST        | /api/client/auth/register      | Register          |
| POST        | /api/client/booking            | Create Booking    |
| GET         | /api/client/booking            | Get All Bookings  |
| GET         | /api/client/booking/:bookingID | Get Booking By ID |

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