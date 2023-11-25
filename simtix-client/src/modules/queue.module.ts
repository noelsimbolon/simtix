import { Module } from '@nestjs/common';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { QueueController } from '../controllers/queue.controller';

@Module({
  imports: [
    ClientsModule.register([
      {
        name: 'CLIENT_QUEUE',
        transport: Transport.RMQ,
        options: {
          urls: ['amqp://simtix-rabbitmq:5672'],
          queue: 'client_queue',
          queueOptions: {
            durable: false,
          },
        },
      },
    ]),
  ],
  controllers: [QueueController],
})
export class QueueModule {}
