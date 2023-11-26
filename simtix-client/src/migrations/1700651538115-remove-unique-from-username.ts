import { MigrationInterface, QueryRunner } from "typeorm";

export class RemoveUniqueFromUsername1700651538115 implements MigrationInterface {
    name = 'RemoveUniqueFromUsername1700651538115'

    public async up(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.query(`ALTER TABLE "users" DROP CONSTRAINT "UQ_fe0bb3f6520ee0469504521e710"`);
    }

    public async down(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.query(`ALTER TABLE "users" ADD CONSTRAINT "UQ_fe0bb3f6520ee0469504521e710" UNIQUE ("username")`);
    }

}
