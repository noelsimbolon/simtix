import { Module } from '@nestjs/common';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { QueueController } from '../controllers/queue.controller';
import process from 'process';

const RABBITMQ_URL = process.env.RABBITMQ_URL || 'amqp://simtix-rabbitmq:5672';

@Module({
  imports: [
    ClientsModule.register([
      {
        name: 'CLIENT_QUEUE',
        transport: Transport.RMQ,
        options: {
          urls: [RABBITMQ_URL],
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
