import { Controller } from '@nestjs/common';
import { EventPattern, Payload, Ctx, RmqContext } from '@nestjs/microservices';
import { BookingService } from '../services/booking.service';

@Controller()
export class QueueController {
  constructor(private readonly bookingService: BookingService) {}

  /**
   * Consumer for booking_process pattern on CLIENT_QUEUE.
   * This method consumes messages from the 'booking_process' queue, updates the status of a booking
   *
   * Payload contract, e.g.
   * {
   *   pattern: 'booking_process';
   *   data: {
   *     id: UUID in string;
   *     status: 'PAID' | 'FAILED' | 'CANCELLED';
   *     bookingUrl: string;
   *   };
   * }
   * @param payload
   * @param context
   */
  @EventPattern('booking_process')
  async consumeMessage(@Payload() payload: any, @Ctx() context: RmqContext) {
    console.log(`Received message: ${JSON.stringify(payload)}`);

    const { id, status, bookingUrl } = payload;

    try {
      await this.bookingService.updateStatus(id, status, bookingUrl);
    } catch (error) {
      console.error(`Error processing message: ${error.message}`);
    }
  }
}
