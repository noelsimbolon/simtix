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

    // Call ticket service API for booking

    const booking = await this.bookingRepository.save({ user, seatId });

    return {
      id: booking.id,
      seatId: booking.seatId,
      status: booking.status,
      createdAt: booking.createdAt,
      userId: booking.user.id,
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
      createdAt: booking.createdAt,
      updatedAt: booking.updatedAt,
    }));
  }

  async findOne(id: string, userId: string) {
    const booking = await this.bookingRepository
      .createQueryBuilder('booking')
      .where('booking.id = :id AND booking.user.id = :userId', { id, userId })
      .getOne();

    if (!booking) {
      throw new NotFoundException('Booking not found');
    }

    return {
      id: booking.id,
      seatId: booking.seatId,
      status: booking.status,
      createdAt: booking.createdAt,
      updatedAt: booking.updatedAt,
    };
  }

  async updateStatus(id: string, status: BookingStatus) {
    // const booking = await this.findOne(id);
    //
    // if (!booking) {
    //   throw new Error('Booking not found');
    // }
    //
    // booking.status = status;
    // return this.bookingRepository.save(booking);
  }
}
