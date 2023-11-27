import { Controller, Post, Body, Req, Get, Param } from '@nestjs/common';
import { BookingService } from '../services/booking.service';
import { CreateBookingDto } from '../domains/dtos/booking.dto';
import { AuthGuard } from '../guards/auth.guard';
import { UseGuards, ParseUUIDPipe } from '@nestjs/common';

@Controller('booking')
@UseGuards(AuthGuard)
export class BookingController {
  constructor(private readonly bookingService: BookingService) {}

  @Post()
  async create(@Body() createBookingDto: CreateBookingDto, @Req() req: any) {
    const userId = req.user.id;
    return this.bookingService.create(userId, createBookingDto.seatId);
  }

  @Get()
  async findAll(@Req() req: any) {
    const userId = req.user.id;
    return this.bookingService.findAll(userId);
  }

  @Get(':id')
  async findOne(@Param('id', new ParseUUIDPipe()) id: string, @Req() req: any) {
    const userId = req.user.id;
    return this.bookingService.findOne(id, userId);
  }
}
