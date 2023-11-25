import { DataSourceOptions } from 'typeorm/data-source/DataSourceOptions';
import * as dotenv from 'dotenv';
import * as fs from 'fs';

const env = dotenv.parse(fs.readFileSync(`.env`));

const typeormConfig: DataSourceOptions = {
  type: 'postgres',
  host: env.PSQL_HOST,
  port: 5432,
  username: env.PSQL_USERNAME,
  password: env.PSQL_PASSWORD,
  database: env.PSQL_DBNAME,
  entities: [__dirname + '/../**/*.entity{.ts,.js}'],
  synchronize: false,
  migrationsTableName: 'migration',
  migrations: [__dirname + '/../migrations/*.ts'],
};

export default typeormConfig;
