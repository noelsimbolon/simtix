import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { ConfigModule } from '@nestjs/config';
import { TypeOrmModule } from '@nestjs/typeorm';
import typeormConfig from './configs/typeorm.config';
import { UserModule } from './modules/user.module';
import { QueueModule } from './modules/queue.module';
import { AuthModule } from './modules/auth.module';
import { BookingModule } from './modules/booking.module';

@Module({
  imports: [
    ConfigModule.forRoot(),
    TypeOrmModule.forRoot(typeormConfig),
    QueueModule,
    UserModule,
    AuthModule,
    BookingModule,
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
