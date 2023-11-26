import { Controller, Post, Body, Req, Get, Param } from '@nestjs/common';
import { BookingService } from '../services/booking.service';
import {
  CreateBookingDto,
  FindOneBookingDto,
} from '../domains/dtos/booking.dto';
import { AuthGuard } from '../guards/auth.guard';
import { UseGuards } from '@nestjs/common';

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
  async findOne(@Param() params: FindOneBookingDto, @Req() req: any) {
    const userId = req.user.id;
    return this.bookingService.findOne(params.id, userId);
  }
}
