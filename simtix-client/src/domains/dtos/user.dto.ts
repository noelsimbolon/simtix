import { IsString, Length, IsOptional } from 'class-validator';

export class CreateUserDto {
  @IsString()
  @Length(1, 255)
  username: string;

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
  username?: string;

  @IsString()
  @IsOptional()
  @Length(1, 255)
  name?: string;

  @IsString()
  @IsOptional()
  @Length(8, 128)
  password?: string;
}
