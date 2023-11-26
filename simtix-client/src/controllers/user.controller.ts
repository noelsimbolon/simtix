import { Controller, Get, Put, Delete, Body, Req } from '@nestjs/common';
import { UserService } from '../services/user.service';
import { UpdateUserDto } from '../domains/dtos/user.dto';
import { Request } from 'express';
import { UseGuards } from '@nestjs/common';
import { AuthGuard } from '../guards/auth.guard';

interface IRequestWithUser extends Request {
  user: {
    id: string;
  };
}

@Controller('users')
@UseGuards(AuthGuard)
export class UserController {
  constructor(private readonly userService: UserService) {}

  @Get()
  async findOne(@Req() req: IRequestWithUser) {
    const userId = req.user.id;
    return await this.userService.findOne(userId);
  }

  @Put()
  async update(
    @Req() req: IRequestWithUser,
    @Body() updateUserDto: UpdateUserDto,
  ) {
    const userId = req.user.id;
    await this.userService.update(userId, updateUserDto);
    return {
      message: 'User updated successfully',
      id: userId,
      changes: updateUserDto,
    };
  }

  @Delete()
  async remove(@Req() req: IRequestWithUser) {
    const userId = req.user.id;
    await this.userService.remove(userId);
    return { message: 'User deleted successfully', id: userId };
  }
}
