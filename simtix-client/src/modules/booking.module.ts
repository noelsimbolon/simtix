import { Module } from '@nestjs/common';
import { BookingService } from '../services/booking.service';
import { BookingController } from '../controllers/booking.controller';
import { UserModule } from './user.module';
import { JwtService } from '@nestjs/jwt';
import { TypeOrmModule } from '@nestjs/typeorm';
import { Booking } from '../domains/entitites/booking.entity';

@Module({
  imports: [TypeOrmModule.forFeature([Booking]), UserModule],
  controllers: [BookingController],
  providers: [BookingService, JwtService],
  exports: [BookingService],
})
export class BookingModule {}
