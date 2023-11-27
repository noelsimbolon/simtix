import { Module } from '@nestjs/common';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { QueueController } from '../controllers/queue.controller';
import { RABBITMQ_URL, QUEUE_NAME } from '../configs/config';
import { BookingModule } from './booking.module';

@Module({
  imports: [
    ClientsModule.register([
      {
        name: 'CLIENT_QUEUE',
        transport: Transport.RMQ,
        options: {
          urls: [RABBITMQ_URL],
          queue: QUEUE_NAME,
          queueOptions: {
            durable: false,
          },
        },
      },
    ]),
    BookingModule,
  ],
  controllers: [QueueController],
})
export class QueueModule {}
