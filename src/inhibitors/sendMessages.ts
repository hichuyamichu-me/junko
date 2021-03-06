import { Message, TextChannel } from 'discord.js';
import { Inhibitor } from 'discord-akairo';

export default class SendMessagesInhibitor extends Inhibitor {
  public constructor() {
    super('sendMessages', {
      reason: 'sendMessages'
    });
  }

  public exec(message: Message): boolean {
    if (!message.guild) return false;
    return !(message.channel as TextChannel)
      .permissionsFor(this.client.user)!
      .has('SEND_MESSAGES');
  }
}
