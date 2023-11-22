import {
  Injectable,
  NotFoundException,
  ConflictException,
} from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { User } from '../domains/entitites/user.entity';
import { CreateUserDto, UpdateUserDto } from '../domains/dtos/user.dto';
import * as bcrypt from 'bcrypt';

@Injectable()
export class UserService {
  constructor(
    @InjectRepository(User)
    private userRepository: Repository<User>,
  ) {}

  async create(createUserDto: CreateUserDto) {
    const user = await this.userRepository.findOne({
      where: { username: createUserDto.username },
    });

    if (user) {
      throw new ConflictException('Username already exists');
    }

    // Hash the password
    createUserDto.password = await this.hashPassword(createUserDto.password);

    return this.userRepository.save(createUserDto);
  }

  findAll() {
    return this.userRepository
      .createQueryBuilder('user')
      .where('user.deleted_at IS NULL')
      .getMany();
  }

  async findOne(id: string) {
    const user = await this.userRepository
      .createQueryBuilder('user')
      .where('user.id = :id', { id })
      .andWhere('user.deleted_at IS NULL')
      .getOne();
    if (!user) {
      throw new NotFoundException('User not found');
    }
    return user;
  }

  async update(id: string, updateUserDto: UpdateUserDto) {
    const user = await this.findOne(id);

    // Hash the new password
    if (updateUserDto.password) {
      updateUserDto.password = await this.hashPassword(updateUserDto.password);
    }

    return this.userRepository.update(id, updateUserDto);
  }

  async remove(id: string) {
    const user = await this.findOne(id);
    user.deleted_at = new Date();
    return this.userRepository.save(user);
  }

  private async hashPassword(password: string): Promise<string> {
    const saltOrRounds = 10;
    return await bcrypt.hash(password, saltOrRounds);
  }
}
