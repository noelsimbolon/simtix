import { MigrationInterface, QueryRunner } from "typeorm";

export class RemoveUsernameAddEmail1700985605033 implements MigrationInterface {
    name = 'RemoveUsernameAddEmail1700985605033'

    public async up(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.query(`ALTER TABLE "users" RENAME COLUMN "username" TO "email"`);
    }

    public async down(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.query(`ALTER TABLE "users" RENAME COLUMN "email" TO "username"`);
    }

}
