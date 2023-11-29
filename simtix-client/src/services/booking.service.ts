import {ConflictException, Injectable, NotFoundException} from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { Booking } from '../domains/entitites/booking.entity';
import { UserService } from './user.service';
import { BookingStatus } from '../domains/entitites/booking.entity';
import { lastValueFrom } from 'rxjs';
import { HttpService } from '@nestjs/axios';

interface IHoldSeat {
  status: number,
  data: {
    seat: {
      status: string
      id: string
    },
    invoice: {
      id: string
      paymentUrl: string
    }
  }
}

@Injectable()
export class BookingService {
  constructor(
    private userService: UserService,
    @InjectRepository(Booking)
    private bookingRepository: Repository<Booking>,
    private httpService: HttpService
  ) {}

  async create(userId: string, seatId: string) {
    const user = await this.userService.findOne(userId);

    const booking = await this.bookingRepository.save({
      user,
      seatId
    });

    const holdSeat: IHoldSeat = await lastValueFrom(
        this.httpService.patch(`${process.env.SIMTIX_TICKETING_URL}/seat`, {
          bookingID: booking.id,
          seatID: seatId,
        }),
    );

    if (holdSeat.status != 200) {
      await this.remove(booking.id)
      throw new ConflictException('Seat not available!')
    }

    const {id: invoiceId, paymentUrl} = holdSeat.data.invoice

    booking.invoiceNumber = invoiceId;
    booking.invoiceUrl = paymentUrl;

    await this.bookingRepository.save(booking);

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

  async remove(id: string) {
    const booking = await this.findOne(id);

    if (!booking) {
      throw new NotFoundException('Booking not found');
    }

    return await this.bookingRepository.softDelete(id);
  }
}
