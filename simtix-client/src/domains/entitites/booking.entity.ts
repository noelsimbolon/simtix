import {
  Entity,
  PrimaryGeneratedColumn,
  Column,
  CreateDateColumn,
  UpdateDateColumn,
  DeleteDateColumn,
  ManyToOne,
  JoinColumn,
} from 'typeorm';
import { User } from './user.entity';

export enum BookingStatus {
  ONGOING = 'ONGOING',
  SUCCESS = 'SUCCESS',
  FAILED = 'FAILED'
}

@Entity({ name: 'bookings' })
export class Booking {
  @PrimaryGeneratedColumn('uuid')
  id: string;

  @ManyToOne(() => User, (user) => user.bookings)
  @JoinColumn({ name: 'user_id' })
  user: User;

  @Column({ type: 'uuid', name: 'seat_id' })
  seatId: string;

  @Column({
    type: 'enum',
    enum: BookingStatus,
    default: BookingStatus.ONGOING,
    name: 'status',
  })
  status: BookingStatus;

  @Column({ type: 'uuid', name: 'invoice_number', nullable: true })
  invoiceNumber: string;

  @Column({ type: 'text', name: 'invoice_url', nullable: true })
  invoiceUrl: string;

  @Column({ type: 'text', name: 'booking_url', nullable: true })
  bookingUrl: string;

  @CreateDateColumn({
    type: 'timestamptz',
    default: () => 'CURRENT_TIMESTAMP',
    name: 'created_at',
  })
  createdAt: Date;

  @UpdateDateColumn({
    type: 'timestamptz',
    default: () => 'CURRENT_TIMESTAMP',
    name: 'updated_at',
  })
  updatedAt: Date;

  @DeleteDateColumn({ type: 'timestamptz', nullable: true, name: 'deleted_at' })
  deletedAt: Date;
}
