export const RABBITMQ_URL =
  process.env.RABBITMQ_URL || 'amqp://simtix-rabbitmq:5672';

export const QUEUE_NAME = 'client_queue';

export const APP_PORT = process.env.APP_PORT || 8000;
