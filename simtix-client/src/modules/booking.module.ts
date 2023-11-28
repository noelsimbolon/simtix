import { Module } from '@nestjs/common';
import { BookingService } from '../services/booking.service';
import { BookingController } from '../controllers/booking.controller';
import { UserModule } from './user.module';
import { JwtService } from '@nestjs/jwt';
import { TypeOrmModule } from '@nestjs/typeorm';
import { Booking } from '../domains/entitites/booking.entity';
import {HttpModule} from "@nestjs/axios";

@Module({
  imports: [TypeOrmModule.forFeature([Booking]), UserModule, HttpModule],
  controllers: [BookingController],
  providers: [BookingService, JwtService],
  exports: [BookingService],
})
export class BookingModule {}
