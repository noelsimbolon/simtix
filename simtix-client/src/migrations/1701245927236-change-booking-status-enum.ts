import { MigrationInterface, QueryRunner } from "typeorm";

export class ChangeBookingStatusEnum1701245927236 implements MigrationInterface {
    name = 'ChangeBookingStatusEnum1701245927236'

    public async up(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.query(`ALTER TYPE "public"."bookings_status_enum" RENAME TO "bookings_status_enum_old"`);
        await queryRunner.query(`CREATE TYPE "public"."bookings_status_enum" AS ENUM('ONGOING', 'SUCCESS', 'FAILED')`);
        await queryRunner.query(`ALTER TABLE "bookings" ALTER COLUMN "status" DROP DEFAULT`);
        await queryRunner.query(`ALTER TABLE "bookings" ALTER COLUMN "status" TYPE "public"."bookings_status_enum" USING "status"::"text"::"public"."bookings_status_enum"`);
        await queryRunner.query(`ALTER TABLE "bookings" ALTER COLUMN "status" SET DEFAULT 'ONGOING'`);
        await queryRunner.query(`DROP TYPE "public"."bookings_status_enum_old"`);
    }

    public async down(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.query(`CREATE TYPE "public"."bookings_status_enum_old" AS ENUM('ONGOING', 'PAID', 'FAILED', 'CANCELLED')`);
        await queryRunner.query(`ALTER TABLE "bookings" ALTER COLUMN "status" DROP DEFAULT`);
        await queryRunner.query(`ALTER TABLE "bookings" ALTER COLUMN "status" TYPE "public"."bookings_status_enum_old" USING "status"::"text"::"public"."bookings_status_enum_old"`);
        await queryRunner.query(`ALTER TABLE "bookings" ALTER COLUMN "status" SET DEFAULT 'ONGOING'`);
        await queryRunner.query(`DROP TYPE "public"."bookings_status_enum"`);
        await queryRunner.query(`ALTER TYPE "public"."bookings_status_enum_old" RENAME TO "bookings_status_enum"`);
    }

}
