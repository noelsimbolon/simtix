import { Controller, Post, Body } from '@nestjs/common';
import { BookingService } from '../services/booking.service';
import { BookDto } from '../domains/dtos/booking.dto';

@Controller('booking')
export class BookingController {
  constructor(private readonly bookingService: BookingService) {}

  @Post()
  async book(@Body() bookDto: BookDto) {}
}
