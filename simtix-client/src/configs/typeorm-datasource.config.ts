import typeormConfig from './typeorm.config';
import { DataSource } from 'typeorm';

export default new DataSource({
  ...typeormConfig,
});
