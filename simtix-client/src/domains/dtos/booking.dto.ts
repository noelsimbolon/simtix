import { IsUUID } from 'class-validator';

export class BookDto {
  @IsUUID()
  seatId: string;
}
