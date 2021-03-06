import { Message } from 'discord.js';
import { Command } from 'discord-akairo';

export default class PingCommand extends Command {
  public constructor() {
    super('ping', {
      aliases: ['ping'],
      category: 'general',
      ownerOnly: false,
      description: {
        content: 'Ping!',
        usage: '',
        examples: ['']
      }
    });
  }

  public async exec(message: Message): Promise<void> {
    const sent = (await message.util!.send('Pong!') as Message);
    const timeDiff =
      (sent.editedTimestamp || sent.createdTimestamp) -
      (message.editedTimestamp || message.createdTimestamp);
    message.util.reply([
      'Pong!',
      `Responce time: \`${timeDiff} ms\``,
      `Gateway latency: \`${Math.round(this.client.ws.ping)} ms\``
    ]);
  }
}
