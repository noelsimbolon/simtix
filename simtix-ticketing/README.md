## Overview
Ticketing service.

## API Docs
### REST API
| HTTP Method | Endpoint                             | Description          |
|-------------|--------------------------------------|----------------------|
| GET         | /api/ticketing/events                | Get All Events       |
| GET         | /api/ticketing/events/:eventID       | Get Event By ID      |
| POST        | /api/ticketing/events                | Create New Event     |
| POST        | /api/ticketing/seat                  | Create New Seat      |
| POST        | /api/ticketing/seat/:seatID          | Get Seat By ID       |
| GET         | /api/ticketing/seat/?eventID=<value> | Get Seat By Event ID |
| PATCH       | /api/ticketing/seat                  | Hold Seat            |
| POST        | /api/ticketing/seat/webhook          | Seat Webhook         |
