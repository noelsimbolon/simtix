import { MigrationInterface, QueryRunner } from "typeorm";

export class SetInvoiceNullable1701196422849 implements MigrationInterface {
    name = 'SetInvoiceNullable1701196422849'

    public async up(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.query(`ALTER TABLE "bookings" ALTER COLUMN "invoice_number" DROP NOT NULL`);
        await queryRunner.query(`ALTER TABLE "bookings" ALTER COLUMN "invoice_url" DROP NOT NULL`);
    }

    public async down(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.query(`ALTER TABLE "bookings" ALTER COLUMN "invoice_url" SET NOT NULL`);
        await queryRunner.query(`ALTER TABLE "bookings" ALTER COLUMN "invoice_number" SET NOT NULL`);
    }

}
