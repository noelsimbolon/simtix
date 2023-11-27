import { Controller } from '@nestjs/common';
import { EventPattern, Payload, Ctx, RmqContext } from '@nestjs/microservices';

@Controller()
export class QueueController {
  @EventPattern('booking_process')
  async consumeMessage(@Payload() data: any, @Ctx() context: RmqContext) {
    console.log(`Received message: ${JSON.stringify(data)}`);
  }
}
