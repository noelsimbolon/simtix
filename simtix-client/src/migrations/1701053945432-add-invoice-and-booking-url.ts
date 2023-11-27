import { MigrationInterface, QueryRunner } from "typeorm";

export class AddInvoiceAndBookingUrl1701053945432 implements MigrationInterface {
    name = 'AddInvoiceAndBookingUrl1701053945432'

    public async up(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.query(`ALTER TABLE "bookings" ADD "invoice_number" uuid NOT NULL`);
        await queryRunner.query(`ALTER TABLE "bookings" ADD "invoice_url" text NOT NULL`);
        await queryRunner.query(`ALTER TABLE "bookings" ADD "booking_url" text NOT NULL`);
    }

    public async down(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.query(`ALTER TABLE "bookings" DROP COLUMN "booking_url"`);
        await queryRunner.query(`ALTER TABLE "bookings" DROP COLUMN "invoice_url"`);
        await queryRunner.query(`ALTER TABLE "bookings" DROP COLUMN "invoice_number"`);
    }

}
