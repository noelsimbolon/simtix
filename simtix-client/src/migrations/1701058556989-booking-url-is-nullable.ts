import { MigrationInterface, QueryRunner } from "typeorm";

export class BookingUrlIsNullable1701058556989 implements MigrationInterface {
    name = 'BookingUrlIsNullable1701058556989'

    public async up(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.query(`ALTER TABLE "bookings" ALTER COLUMN "booking_url" DROP NOT NULL`);
    }

    public async down(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.query(`ALTER TABLE "bookings" ALTER COLUMN "booking_url" SET NOT NULL`);
    }

}
