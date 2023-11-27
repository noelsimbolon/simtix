// booking.service.ts
import { Injectable, NotFoundException } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { Booking } from '../domains/entitites/booking.entity';
import { UserService } from './user.service';
import { BookingStatus } from '../domains/entitites/booking.entity';

@Injectable()
export class BookingService {
  constructor(
    private userService: UserService,
    @InjectRepository(Booking)
    private bookingRepository: Repository<Booking>,
  ) {}

  async create(userId: string, seatId: string) {
    const user = await this.userService.findOne(userId);

    // Call ticket service API for booking, constants below are only dummy
    const invoiceNumber = 'f1f5b363-446d-4185-a872-66dadfb31153';
    const invoiceUrl =
      'https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf';

    const booking = await this.bookingRepository.save({
      user,
      seatId,
      invoiceNumber,
      invoiceUrl,
    });

    const { bookingUrl, deletedAt, ...returnData } = booking;

    return {
      ...returnData,
    };
  }

  async findAll(userId: string) {
    const bookings = await this.bookingRepository
      .createQueryBuilder('booking')
      .where('booking.user.id = :userId', { userId })
      .getMany();

    return bookings.map((booking) => ({
      id: booking.id,
      seatId: booking.seatId,
      status: booking.status,
      updatedAt: booking.updatedAt,
    }));
  }

  async findOne(id: string, userId?: string) {
    let query = this.bookingRepository
      .createQueryBuilder('booking')
      .where('booking.id = :id', { id });

    if (userId) {
      query = query.andWhere('booking.user.id = :userId', { userId });
    }

    const booking = await query.getOne();

    if (!booking) {
      throw new NotFoundException('Booking not found');
    }

    const { deletedAt, ...returnData } = booking;

    return {
      ...returnData,
    };
  }

  async updateStatus(id: string, status: string, bookingUrl: string) {
    const booking = await this.findOne(id);

    booking.status = status as BookingStatus;
    booking.bookingUrl = bookingUrl;

    await this.bookingRepository.save(booking);

    return booking;
  }
}
