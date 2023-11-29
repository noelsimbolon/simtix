## Overview
Simtix is a ticketing and payment simulation application implemented using the microservice architecture.

## Pre-Run
- For each folder, copy .env.docker to .env

## How to Run
### Create the network if `simtix` network doesn't exist
```bash
make create-network
```

### To Run the Services
```bash
make
```

### To Seed the Data
```bash
make seed
```

### To Stop
```bash
make stop-all
```

## Tech Stack
### Client Service
| System Entity | Frameworks/Tools |
|---------------|------------------|
| Web Framework | NestJS           |
| Database      | PostgreSQL       |

### Ticketing Service
| System Entity | Frameworks/Tools                                  |
|---------------|---------------------------------------------------|
| Web Framework | Gin                                               |
| Database      | PostgreSQL                                        |
| Task Queue    | [Asynq](https://github.com/hibiken/asynq) & Redis |

### Payment Service
| System Entity | Frameworks/Tools                                  |
|---------------|---------------------------------------------------|
| Web Framework | Gin                                               |
| Database      | PostgreSQL                                        |
| Task Queue    | [Asynq](https://github.com/hibiken/asynq) & Redis |

### Miscellaneous
| System Entity | Frameworks/Tools |
|---------------|------------------|
| API Gateway   | NGINX            |
| Message Queue | RabbitMQ         |

## Authors
| Name                 | GitHub                                             |
|----------------------|----------------------------------------------------|
| Noel Simbolon        | [noelsimbolon](https://github.com/noelsimbolon)    |
| Rachel Gabriela Chen | [chaerla](https://github.com/chaerla)              |
| Jeffrey Chow         | [JeffreyChow19](https://github.com/JeffreyChow19)  |
