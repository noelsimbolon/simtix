import { IsString, Length, IsOptional, IsEmail } from 'class-validator';

export class CreateUserDto {
  @IsEmail()
  @Length(1, 255)
  email: string;

  @IsString()
  @Length(1, 255)
  name: string;

  @IsString()
  @Length(8, 128)
  password: string;
}

export class UpdateUserDto {
  @IsString()
  @IsOptional()
  @Length(1, 255)
  name?: string;

  @IsString()
  @IsOptional()
  @Length(8, 128)
  password?: string;
}

export class LoginDto {
  @IsEmail()
  @Length(1, 255)
  email: string;

  @IsString()
  @Length(8, 128)
  password: string;
}
